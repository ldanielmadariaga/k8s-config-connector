// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package codegen

import (
	"bytes"
	"strings"
	"testing"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"k8s.io/apimachinery/pkg/util/sets"
)

// serverSetTestFile builds two messages:
//
//	Discovery  no field_behavior anywhere, like compute's protos
//	Annotated  one field marked OUTPUT_ONLY, the rest bare
//
// Both carry the same field names, so a test can isolate the guard from the
// allowlist.
func serverSetTestFile(t *testing.T) protoreflect.FileDescriptor {
	t.Helper()

	outputOnly := &descriptorpb.FieldOptions{}
	proto.SetExtension(outputOnly, annotations.E_FieldBehavior,
		[]annotations.FieldBehavior{annotations.FieldBehavior_OUTPUT_ONLY})

	names := []string{"creation_timestamp", "self_link", "etag", "state", "status", "type", "name", "description"}
	bare := func() []*descriptorpb.FieldDescriptorProto {
		var out []*descriptorpb.FieldDescriptorProto
		for i, n := range names {
			out = append(out, &descriptorpb.FieldDescriptorProto{
				Name:   protoPtr(n),
				Number: protoPtr(int32(i + 1)),
				Type:   typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING),
			})
		}
		return out
	}

	annotated := bare()
	// Only "description" is annotated. No allowlisted field carries an
	// annotation itself, so the guard has to come from the message, not the
	// field.
	annotated[len(annotated)-1].Options = outputOnly

	fdp := &descriptorpb.FileDescriptorProto{
		Name:    protoPtr("serverset.proto"),
		Package: protoPtr("google.cloud.test.v1"),
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: protoPtr("Discovery"), Field: bare()},
			{Name: protoPtr("Annotated"), Field: annotated},
		},
	}
	fd, err := protodesc.NewFile(fdp, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	return fd
}

// TestIsServerSetField covers the allowlist and the guard: the names the rule
// moves, the names it leaves in the Spec, and the annotation that turns it off.
func TestIsServerSetField(t *testing.T) {
	fd := serverSetTestFile(t)
	discovery := fd.Messages().ByName("Discovery")
	annotated := fd.Messages().ByName("Annotated")
	on := WriteOptions{PlaceServerSetFields: true}

	tests := []struct {
		name  string
		msg   protoreflect.MessageDescriptor
		field string
		opts  WriteOptions
		want  bool
	}{
		{"allowlisted, nothing annotated", discovery, "creation_timestamp", on, true},
		{"allowlisted, nothing annotated", discovery, "self_link", on, true},
		{"etag is in the list", discovery, "etag", on, true},

		// serverSetFieldNames records why these three stay out.
		{"state stays in the Spec", discovery, "state", on, false},
		{"status stays in the Spec", discovery, "status", on, false},
		{"type stays in the Spec", discovery, "type", on, false},
		// identityFields in the scaffold package handles name instead.
		{"name is left to the identity policy", discovery, "name", on, false},
		{"an ordinary field is untouched", discovery, "description", on, false},

		// The guard: one annotation anywhere on the message turns the rule off.
		{"annotated message, allowlisted field", annotated, "creation_timestamp", on, false},
		{"annotated message, etag", annotated, "etag", on, false},

		{"off by default", discovery, "creation_timestamp", WriteOptions{}, false},
	}
	for _, tc := range tests {
		t.Run(tc.name+"/"+tc.field, func(t *testing.T) {
			// Arrange
			f := tc.msg.Fields().ByName(protoreflect.Name(tc.field))
			if f == nil {
				t.Fatalf("no field %q on %s", tc.field, tc.msg.Name())
			}

			// Act
			got := IsServerSetField(f, tc.msg, tc.opts)

			// Assert
			if got != tc.want {
				t.Errorf("IsServerSetField(%s.%s) = %v, want %v", tc.msg.Name(), tc.field, got, tc.want)
			}
		})
	}
}

// TestIsServerSetFieldOnlyAppliesToTheRootMessage pins where the rule stops.
// identifyOutputs recurses into nested messages, and a creationTimestamp found
// there is often the user's to set.
func TestIsServerSetFieldOnlyAppliesToTheRootMessage(t *testing.T) {
	// Arrange: one field of one message, asked twice. Discovery carries no
	// annotation, so the guard cannot account for a false answer here and only
	// the root check can.
	fd := serverSetTestFile(t)
	discovery := fd.Messages().ByName("Discovery")
	field := discovery.Fields().ByName("creation_timestamp")
	opts := WriteOptions{PlaceServerSetFields: true}
	visiting := &TypeGenerator{writeOptions: opts, rootMessageFQN: string(discovery.FullName())}
	elsewhere := &TypeGenerator{writeOptions: opts, rootMessageFQN: "google.cloud.test.v1.Annotated"}

	// Act
	atRoot := visiting.isServerSet(field, discovery)
	away := elsewhere.isServerSet(field, discovery)

	// Assert
	if !atRoot {
		t.Error("creationTimestamp on the visited message is not server-set, want server-set")
	}
	if away {
		t.Error("creationTimestamp below the visited message is server-set, want not server-set")
	}
}

// TestWriteFieldWritesNoteOnce covers both places a note comes from: the
// caller, which passes the placement note for a server-set field, and the
// sibling rule inside WriteField.
func TestWriteFieldWritesNoteOnce(t *testing.T) {
	// Arrange
	fd := serverSetTestFile(t)
	msg := fd.Messages().ByName("Discovery")
	creationTimestamp := msg.Fields().ByName("creation_timestamp")
	description := msg.Fields().ByName("description")

	tests := []struct {
		name   string
		field  protoreflect.FieldDescriptor
		opts   WriteOptions
		note   string
		marker string
	}{
		{
			name:   "note from the caller",
			field:  creationTimestamp,
			note:   placementNote(creationTimestamp, msg, WriteOptions{PlaceServerSetFields: true}),
			marker: "+kcc:guess=placement",
		},
		{
			name:   "note from the sibling rule",
			field:  description,
			opts:   WriteOptions{Siblings: map[string]string{"description": "TestDescription"}},
			marker: SiblingGuessMarker + "TestDescription",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			var buf bytes.Buffer
			WriteField(&buf, tc.field, msg, 0, false, tc.opts, tc.note)

			// Assert
			if got := strings.Count(buf.String(), tc.marker); got != 1 {
				t.Errorf("WriteField() wrote %q %d times, want once:\n%s", tc.marker, got, buf.String())
			}
		})
	}
}

// TestServerSetPlacement covers both rules, which one wins when both match,
// and the fields neither rule moves.
func TestServerSetPlacement(t *testing.T) {
	// Arrange: Network carries no field_behavior anywhere, like compute's
	// protos. Pipeline annotates some fields and not others.
	network := commentedMessage(t, "Network",
		commentedField{name: "creation_timestamp", comment: "[Output Only] Creation timestamp in RFC3339 text format."},
		commentedField{name: "self_link", comment: "Server-defined URL for the resource."},
		commentedField{name: "firewall_policy", comment: "[Output Only] URL of the firewall policy the network is associated with."},
		commentedField{name: "name", comment: "[Output Only] Name of the resource."},
		commentedField{name: "description", comment: "An optional description of this resource."},
	)
	pipeline := commentedMessage(t, "Pipeline",
		commentedField{name: "display_name", comment: "Required. Display name.", behaviors: []annotations.FieldBehavior{annotations.FieldBehavior_REQUIRED}},
		commentedField{name: "etag", comment: "Output only. This checksum is computed by the server."},
		commentedField{name: "create_time", comment: "Output only. The creation time.", behaviors: []annotations.FieldBehavior{annotations.FieldBehavior_OUTPUT_ONLY}},
		commentedField{name: "state", comment: "Output only. The state of the pipeline.", behaviors: []annotations.FieldBehavior{annotations.FieldBehavior_OPTIONAL}},
	)
	byComment := WriteOptions{PlaceOutputOnlyFromComments: true}
	byName := WriteOptions{PlaceServerSetFields: true}
	both := WriteOptions{PlaceOutputOnlyFromComments: true, PlaceServerSetFields: true}

	tests := []struct {
		name  string
		msg   protoreflect.MessageDescriptor
		field string
		opts  WriteOptions
		want  PlacementReason
	}{
		{"comment rule", network, "firewall_policy", byComment, PlacedByComment},
		{"comment rule ignores the name", network, "self_link", byComment, ""},
		{"name rule", network, "self_link", byName, PlacedByName},
		{"name rule ignores the comment", network, "firewall_policy", byName, ""},
		{"both match, the comment wins", network, "creation_timestamp", both, PlacedByComment},
		{"name is left to the identity policy", network, "name", both, ""},
		{"no signal", network, "description", both, ""},

		// One annotation on the message turns the name rule off. The comment
		// rule still applies to a field that carries none itself.
		{"partly annotated message, comment rule", pipeline, "etag", both, PlacedByComment},
		{"partly annotated message, name rule", pipeline, "etag", byName, ""},
		{"a field marked OUTPUT_ONLY is already an output", pipeline, "create_time", both, ""},
		{"a field's own annotation outranks its comment", pipeline, "state", both, ""},

		{"off by default", network, "creation_timestamp", WriteOptions{}, ""},
	}
	for _, tc := range tests {
		t.Run(tc.name+"/"+tc.field, func(t *testing.T) {
			// Arrange
			f := tc.msg.Fields().ByName(protoreflect.Name(tc.field))
			if f == nil {
				t.Fatalf("no field %q on %s", tc.field, tc.msg.Name())
			}

			// Act
			got := ServerSetPlacement(f, tc.msg, tc.opts)

			// Assert
			if got != tc.want {
				t.Errorf("ServerSetPlacement(%s.%s) = %q, want %q", tc.msg.Name(), tc.field, got, tc.want)
			}
			if IsServerSetField(f, tc.msg, tc.opts) != (tc.want != "") {
				t.Errorf("IsServerSetField(%s.%s) disagrees with ServerSetPlacement", tc.msg.Name(), tc.field)
			}
		})
	}
}

// TestCommentRuleOnlyAppliesToTheRootMessage pins the comment rule to the
// same scope as the name rule. A nested message can be shared by several
// resources in the package, so a move there would reach all of them.
func TestCommentRuleOnlyAppliesToTheRootMessage(t *testing.T) {
	// Arrange
	network := commentedMessage(t, "Network",
		commentedField{name: "firewall_policy", comment: "[Output Only] URL of the firewall policy."})
	field := network.Fields().ByName("firewall_policy")
	opts := WriteOptions{PlaceOutputOnlyFromComments: true}
	visiting := &TypeGenerator{writeOptions: opts, rootMessageFQN: string(network.FullName())}
	elsewhere := &TypeGenerator{writeOptions: opts, rootMessageFQN: "google.cloud.test.v1.Other"}

	// Act
	atRoot := visiting.isServerSet(field, network)
	away := elsewhere.isServerSet(field, network)

	// Assert
	if !atRoot {
		t.Error("firewallPolicy on the visited message is not placed, want placed")
	}
	if away {
		t.Error("firewallPolicy below the visited message is placed, want not placed")
	}
}

// TestPlacementNoteNamesTheRule checks that the marker says which rule moved
// the field, so a reader of the Go type knows what to check.
func TestPlacementNoteNamesTheRule(t *testing.T) {
	// Arrange
	network := commentedMessage(t, "Network",
		commentedField{name: "firewall_policy", comment: "[Output Only] URL of the firewall policy."},
		commentedField{name: "self_link", comment: "Server-defined URL for the resource."},
		commentedField{name: "description", comment: "An optional description of this resource."},
	)
	both := WriteOptions{PlaceOutputOnlyFromComments: true, PlaceServerSetFields: true}

	for _, tc := range []struct {
		field string
		want  string
	}{
		{"firewall_policy", "+kcc:guess=placement reason=output-only-in-comment"},
		{"self_link", "+kcc:guess=placement reason=no-field-behavior-on-message"},
		{"description", ""},
	} {
		t.Run(tc.field, func(t *testing.T) {
			// Act
			got := placementNote(network.Fields().ByName(protoreflect.Name(tc.field)), network, both)

			// Assert
			if got != tc.want {
				t.Errorf("placementNote(%s) = %q, want %q", tc.field, got, tc.want)
			}
		})
	}
}

// TestWriteObservedStateMessageLeavesOutPlacementNotes covers the nested
// ObservedState structs in types.generated.go. Only the resource's own
// ObservedState, which the scaffolder writes, carries a placement note.
func TestWriteObservedStateMessageLeavesOutPlacementNotes(t *testing.T) {
	// Arrange
	network := commentedMessage(t, "Network",
		commentedField{name: "firewall_policy", comment: "[Output Only] URL of the firewall policy."},
		commentedField{name: "self_link", comment: "Server-defined URL for the resource."},
	)
	details := &OutputMessageDetails{
		Message:      network,
		OutputFields: []protoreflect.FieldDescriptor{network.Fields().Get(0), network.Fields().Get(1)},
	}
	both := WriteOptions{PlaceOutputOnlyFromComments: true, PlaceServerSetFields: true}

	// Act
	var buf bytes.Buffer
	WriteObservedStateMessage(&buf, details, sets.NewString(), both)

	// Assert
	if strings.Contains(buf.String(), "+kcc:guess=placement") {
		t.Errorf("nested ObservedState carries a placement note:\n%s", buf.String())
	}
	if !strings.Contains(buf.String(), `json:"firewallPolicy,`) {
		t.Errorf("nested ObservedState is missing firewallPolicy:\n%s", buf.String())
	}
}
