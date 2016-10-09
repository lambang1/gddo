// Code generated by protoc-gen-go.
// source: google.golang.org/genproto/googleapis/api/serviceconfig/log.proto
// DO NOT EDIT!

package google_api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_api1 "google.golang.org/genproto/googleapis/api/label"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A description of a log type. Example in YAML format:
//
//     - name: library.googleapis.com/activity_history
//       description: The history of borrowing and returning library items.
//       display_name: Activity
//       labels:
//       - key: /customer_id
//         description: Identifier of a library customer
type LogDescriptor struct {
	// The name of the log. It must be less than 512 characters long and can
	// include the following characters: upper- and lower-case alphanumeric
	// characters [A-Za-z0-9], and punctuation characters including
	// slash, underscore, hyphen, period [/_-.].
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The set of labels that are available to describe a specific log entry.
	// Runtime requests that contain labels not specified here are
	// considered invalid.
	Labels []*google_api1.LabelDescriptor `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty"`
	// A human-readable description of this log. This information appears in
	// the documentation and can contain details.
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	// The human-readable name for this log. This information appears on
	// the user interface and should be concise.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName" json:"display_name,omitempty"`
}

func (m *LogDescriptor) Reset()                    { *m = LogDescriptor{} }
func (m *LogDescriptor) String() string            { return proto.CompactTextString(m) }
func (*LogDescriptor) ProtoMessage()               {}
func (*LogDescriptor) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *LogDescriptor) GetLabels() []*google_api1.LabelDescriptor {
	if m != nil {
		return m.Labels
	}
	return nil
}

func init() {
	proto.RegisterType((*LogDescriptor)(nil), "google.api.LogDescriptor")
}

func init() {
	proto.RegisterFile("google.golang.org/genproto/googleapis/api/serviceconfig/log.proto", fileDescriptor10)
}

var fileDescriptor10 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x72, 0x4c, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4b, 0xcf, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0xcb, 0x2f, 0x4a, 0xd7, 0x4f, 0x4f,
	0xcd, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x87, 0x48, 0x25, 0x16, 0x64, 0x16, 0xeb, 0x03, 0x09,
	0xfd, 0xe2, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0xd4, 0xe4, 0xfc, 0xbc, 0xb4, 0xcc, 0x74, 0xfd, 0x9c,
	0xfc, 0x74, 0x3d, 0xb0, 0x32, 0x21, 0x2e, 0xa8, 0x11, 0x40, 0x35, 0x52, 0xd6, 0xc4, 0x1b, 0x97,
	0x93, 0x98, 0x94, 0x9a, 0x03, 0x21, 0x21, 0x06, 0x29, 0xcd, 0x65, 0xe4, 0xe2, 0xf5, 0xc9, 0x4f,
	0x77, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0x2f, 0x12, 0x12, 0xe2, 0x62, 0xc9, 0x4b,
	0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x8c, 0xb9, 0xd8, 0xc0,
	0x9a, 0x8a, 0x25, 0x98, 0x14, 0x98, 0x35, 0xb8, 0x8d, 0xa4, 0xf5, 0x10, 0xf6, 0xeb, 0xf9, 0x80,
	0x64, 0x10, 0x06, 0x04, 0x41, 0x95, 0x0a, 0x29, 0x70, 0x71, 0xa7, 0x40, 0x45, 0x33, 0xf3, 0xf3,
	0x24, 0x98, 0xc1, 0xe6, 0x21, 0x0b, 0x09, 0x29, 0x72, 0xf1, 0xa4, 0x64, 0x16, 0x17, 0xe4, 0x24,
	0x56, 0xc6, 0x83, 0xad, 0x64, 0x81, 0x2a, 0x81, 0x88, 0xf9, 0x01, 0x85, 0x9c, 0x94, 0xb9, 0xf8,
	0x92, 0xf3, 0x73, 0x91, 0xac, 0x73, 0xe2, 0x00, 0x3a, 0x37, 0x00, 0xe4, 0xf6, 0x00, 0xc6, 0x45,
	0x4c, 0x2c, 0xee, 0x8e, 0x01, 0x9e, 0x49, 0x6c, 0x60, 0xbf, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x32, 0x96, 0x08, 0x72, 0x59, 0x01, 0x00, 0x00,
}