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
	"testing"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// commentedField is one string field for commentedMessage, with a leading
// comment and optional field_behavior annotations.
type commentedField struct {
	name      string
	comment   string
	behaviors []annotations.FieldBehavior
}

// commentedMessage builds a message whose fields carry leading comments, which
// is what OutputOnlyComment reads. SourceCodeInfo paths are
// [4=message_type, msgIndex, 2=field, fieldIndex].
func commentedMessage(t *testing.T, name string, fields ...commentedField) protoreflect.MessageDescriptor {
	t.Helper()
	var fdps []*descriptorpb.FieldDescriptorProto
	var locs []*descriptorpb.SourceCodeInfo_Location
	for i, f := range fields {
		fdp := &descriptorpb.FieldDescriptorProto{
			Name:   protoPtr(f.name),
			Number: protoPtr(int32(i + 1)),
			Type:   typeDescriptor(descriptorpb.FieldDescriptorProto_TYPE_STRING),
		}
		if len(f.behaviors) > 0 {
			fdp.Options = &descriptorpb.FieldOptions{}
			proto.SetExtension(fdp.Options, annotations.E_FieldBehavior, f.behaviors)
		}
		fdps = append(fdps, fdp)
		locs = append(locs, &descriptorpb.SourceCodeInfo_Location{
			Path:            []int32{4, 0, 2, int32(i)},
			Span:            []int32{int32(i), 0, 1},
			LeadingComments: protoPtr(" " + f.comment + "\n"),
		})
	}
	fd, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name:           protoPtr("commented.proto"),
		Package:        protoPtr("google.cloud.test.v1"),
		MessageType:    []*descriptorpb.DescriptorProto{{Name: protoPtr(name), Field: fdps}},
		SourceCodeInfo: &descriptorpb.SourceCodeInfo{Location: locs},
	}, nil)
	if err != nil {
		t.Fatalf("building file descriptor: %v", err)
	}
	return fd.Messages().ByName(protoreflect.Name(name))
}

// TestOutputOnlyComment covers the two prefixes in any case, and comments that
// mention output only without opening with it.
func TestOutputOnlyComment(t *testing.T) {
	for _, tc := range []struct {
		comment string
		want    bool
	}{
		{"Output only. Set by the server.", true},              // the long-standing spelling
		{"[Output Only] IP address on the Google side.", true}, // how compute writes it
		// Compute also writes "[Output only]", so the prefixes match in any case.
		{"[Output only] Number of network endpoints in the group.", true},
		{"Output Only. The overall outcome of the test.", true}, // devtools.testing
		{"The display name of the widget.", false},
		{"Set by the user. Output only in some other sense.", false}, // not at the front
		// Where the words run on into a condition, the field is only output
		// some of the time.
		{"[Output only for type PARTNER. Input only for PARTNER_PROVIDER.] Pairing key.", false}, // compute
		{"Output only for the create operation. Required for update.", false},                    // spanner-style
	} {
		t.Run(tc.comment, func(t *testing.T) {
			// Arrange
			msg := commentedMessage(t, "Widget", commentedField{name: "widget_field", comment: tc.comment})

			// Act
			_, got := OutputOnlyComment(msg.Fields().Get(0))

			// Assert
			if got != tc.want {
				t.Errorf("OutputOnlyComment(%q) matched = %v, want %v", tc.comment, got, tc.want)
			}
		})
	}
}

// The comment comes back on one line, because the queue entry that quotes it
// is one line.
func TestOutputOnlyCommentIsOneLine(t *testing.T) {
	// Arrange
	msg := commentedMessage(t, "Network", commentedField{
		name:    "gateway_ipv4",
		comment: "[Output Only] The gateway address for default routing\n out of the network.",
	})
	want := "[Output Only] The gateway address for default routing out of the network."

	// Act
	got, ok := OutputOnlyComment(msg.Fields().Get(0))

	// Assert
	if !ok || got != want {
		t.Errorf("OutputOnlyComment() = %q, %v, want %q, true", got, ok, want)
	}
}
