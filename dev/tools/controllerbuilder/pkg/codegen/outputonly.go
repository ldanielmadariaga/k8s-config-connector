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
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"
)

// outputOnlyPrefixes are the markers Google API comments use to say a field is
// output only when the proto has no field_behavior annotation. They match
// regardless of case, because compute writes both "[Output Only]" and
// "[Output only]".
var outputOnlyPrefixes = []string{"Output only.", "[Output Only]"}

// OutputOnlyComment returns a field's leading comment, formatted as a single
// line, if the comment opens with one of outputOnlyPrefixes.
func OutputOnlyComment(field protoreflect.FieldDescriptor) (string, bool) {
	loc := field.ParentFile().SourceLocations().ByDescriptor(field)
	comment := strings.TrimSpace(loc.LeadingComments)
	for _, prefix := range outputOnlyPrefixes {
		if len(comment) >= len(prefix) && strings.EqualFold(comment[:len(prefix)], prefix) {
			return strings.Join(strings.Fields(comment), " "), true
		}
	}
	return "", false
}
