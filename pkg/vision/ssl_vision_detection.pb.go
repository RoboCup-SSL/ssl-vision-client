// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: ssl_vision_detection.proto

package vision

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type SSL_DetectionBall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Confidence *float32 `protobuf:"fixed32,1,req,name=confidence" json:"confidence,omitempty"`
	Area       *uint32  `protobuf:"varint,2,opt,name=area" json:"area,omitempty"`
	X          *float32 `protobuf:"fixed32,3,req,name=x" json:"x,omitempty"`
	Y          *float32 `protobuf:"fixed32,4,req,name=y" json:"y,omitempty"`
	Z          *float32 `protobuf:"fixed32,5,opt,name=z" json:"z,omitempty"`
	PixelX     *float32 `protobuf:"fixed32,6,req,name=pixel_x,json=pixelX" json:"pixel_x,omitempty"`
	PixelY     *float32 `protobuf:"fixed32,7,req,name=pixel_y,json=pixelY" json:"pixel_y,omitempty"`
}

func (x *SSL_DetectionBall) Reset() {
	*x = SSL_DetectionBall{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_vision_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSL_DetectionBall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSL_DetectionBall) ProtoMessage() {}

func (x *SSL_DetectionBall) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_vision_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSL_DetectionBall.ProtoReflect.Descriptor instead.
func (*SSL_DetectionBall) Descriptor() ([]byte, []int) {
	return file_ssl_vision_detection_proto_rawDescGZIP(), []int{0}
}

func (x *SSL_DetectionBall) GetConfidence() float32 {
	if x != nil && x.Confidence != nil {
		return *x.Confidence
	}
	return 0
}

func (x *SSL_DetectionBall) GetArea() uint32 {
	if x != nil && x.Area != nil {
		return *x.Area
	}
	return 0
}

func (x *SSL_DetectionBall) GetX() float32 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *SSL_DetectionBall) GetY() float32 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

func (x *SSL_DetectionBall) GetZ() float32 {
	if x != nil && x.Z != nil {
		return *x.Z
	}
	return 0
}

func (x *SSL_DetectionBall) GetPixelX() float32 {
	if x != nil && x.PixelX != nil {
		return *x.PixelX
	}
	return 0
}

func (x *SSL_DetectionBall) GetPixelY() float32 {
	if x != nil && x.PixelY != nil {
		return *x.PixelY
	}
	return 0
}

type SSL_DetectionRobot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Confidence  *float32 `protobuf:"fixed32,1,req,name=confidence" json:"confidence,omitempty"`
	RobotId     *uint32  `protobuf:"varint,2,opt,name=robot_id,json=robotId" json:"robot_id,omitempty"`
	X           *float32 `protobuf:"fixed32,3,req,name=x" json:"x,omitempty"`
	Y           *float32 `protobuf:"fixed32,4,req,name=y" json:"y,omitempty"`
	Orientation *float32 `protobuf:"fixed32,5,opt,name=orientation" json:"orientation,omitempty"`
	PixelX      *float32 `protobuf:"fixed32,6,req,name=pixel_x,json=pixelX" json:"pixel_x,omitempty"`
	PixelY      *float32 `protobuf:"fixed32,7,req,name=pixel_y,json=pixelY" json:"pixel_y,omitempty"`
	Height      *float32 `protobuf:"fixed32,8,opt,name=height" json:"height,omitempty"`
}

func (x *SSL_DetectionRobot) Reset() {
	*x = SSL_DetectionRobot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_vision_detection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSL_DetectionRobot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSL_DetectionRobot) ProtoMessage() {}

func (x *SSL_DetectionRobot) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_vision_detection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSL_DetectionRobot.ProtoReflect.Descriptor instead.
func (*SSL_DetectionRobot) Descriptor() ([]byte, []int) {
	return file_ssl_vision_detection_proto_rawDescGZIP(), []int{1}
}

func (x *SSL_DetectionRobot) GetConfidence() float32 {
	if x != nil && x.Confidence != nil {
		return *x.Confidence
	}
	return 0
}

func (x *SSL_DetectionRobot) GetRobotId() uint32 {
	if x != nil && x.RobotId != nil {
		return *x.RobotId
	}
	return 0
}

func (x *SSL_DetectionRobot) GetX() float32 {
	if x != nil && x.X != nil {
		return *x.X
	}
	return 0
}

func (x *SSL_DetectionRobot) GetY() float32 {
	if x != nil && x.Y != nil {
		return *x.Y
	}
	return 0
}

func (x *SSL_DetectionRobot) GetOrientation() float32 {
	if x != nil && x.Orientation != nil {
		return *x.Orientation
	}
	return 0
}

func (x *SSL_DetectionRobot) GetPixelX() float32 {
	if x != nil && x.PixelX != nil {
		return *x.PixelX
	}
	return 0
}

func (x *SSL_DetectionRobot) GetPixelY() float32 {
	if x != nil && x.PixelY != nil {
		return *x.PixelY
	}
	return 0
}

func (x *SSL_DetectionRobot) GetHeight() float32 {
	if x != nil && x.Height != nil {
		return *x.Height
	}
	return 0
}

type SSL_DetectionFrame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrameNumber  *uint32               `protobuf:"varint,1,req,name=frame_number,json=frameNumber" json:"frame_number,omitempty"`
	TCapture     *float64              `protobuf:"fixed64,2,req,name=t_capture,json=tCapture" json:"t_capture,omitempty"`
	TSent        *float64              `protobuf:"fixed64,3,req,name=t_sent,json=tSent" json:"t_sent,omitempty"`
	CameraId     *uint32               `protobuf:"varint,4,req,name=camera_id,json=cameraId" json:"camera_id,omitempty"`
	Balls        []*SSL_DetectionBall  `protobuf:"bytes,5,rep,name=balls" json:"balls,omitempty"`
	RobotsYellow []*SSL_DetectionRobot `protobuf:"bytes,6,rep,name=robots_yellow,json=robotsYellow" json:"robots_yellow,omitempty"`
	RobotsBlue   []*SSL_DetectionRobot `protobuf:"bytes,7,rep,name=robots_blue,json=robotsBlue" json:"robots_blue,omitempty"`
}

func (x *SSL_DetectionFrame) Reset() {
	*x = SSL_DetectionFrame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_vision_detection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SSL_DetectionFrame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SSL_DetectionFrame) ProtoMessage() {}

func (x *SSL_DetectionFrame) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_vision_detection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SSL_DetectionFrame.ProtoReflect.Descriptor instead.
func (*SSL_DetectionFrame) Descriptor() ([]byte, []int) {
	return file_ssl_vision_detection_proto_rawDescGZIP(), []int{2}
}

func (x *SSL_DetectionFrame) GetFrameNumber() uint32 {
	if x != nil && x.FrameNumber != nil {
		return *x.FrameNumber
	}
	return 0
}

func (x *SSL_DetectionFrame) GetTCapture() float64 {
	if x != nil && x.TCapture != nil {
		return *x.TCapture
	}
	return 0
}

func (x *SSL_DetectionFrame) GetTSent() float64 {
	if x != nil && x.TSent != nil {
		return *x.TSent
	}
	return 0
}

func (x *SSL_DetectionFrame) GetCameraId() uint32 {
	if x != nil && x.CameraId != nil {
		return *x.CameraId
	}
	return 0
}

func (x *SSL_DetectionFrame) GetBalls() []*SSL_DetectionBall {
	if x != nil {
		return x.Balls
	}
	return nil
}

func (x *SSL_DetectionFrame) GetRobotsYellow() []*SSL_DetectionRobot {
	if x != nil {
		return x.RobotsYellow
	}
	return nil
}

func (x *SSL_DetectionFrame) GetRobotsBlue() []*SSL_DetectionRobot {
	if x != nil {
		return x.RobotsBlue
	}
	return nil
}

var File_ssl_vision_detection_proto protoreflect.FileDescriptor

var file_ssl_vision_detection_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x73, 0x6c, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a,
	0x11, 0x53, 0x53, 0x4c, 0x5f, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61,
	0x6c, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x02, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x03, 0x20, 0x02, 0x28,
	0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x04, 0x20, 0x02, 0x28, 0x02, 0x52,
	0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a,
	0x12, 0x17, 0x0a, 0x07, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x5f, 0x78, 0x18, 0x06, 0x20, 0x02, 0x28,
	0x02, 0x52, 0x06, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x58, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x69, 0x78,
	0x65, 0x6c, 0x5f, 0x79, 0x18, 0x07, 0x20, 0x02, 0x28, 0x02, 0x52, 0x06, 0x70, 0x69, 0x78, 0x65,
	0x6c, 0x59, 0x22, 0xd7, 0x01, 0x0a, 0x12, 0x53, 0x53, 0x4c, 0x5f, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x02, 0x52, 0x0a, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f, 0x62,
	0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x6f, 0x62,
	0x6f, 0x74, 0x49, 0x64, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x03, 0x20, 0x02, 0x28, 0x02, 0x52,
	0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x04, 0x20, 0x02, 0x28, 0x02, 0x52, 0x01, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x5f, 0x78, 0x18, 0x06, 0x20,
	0x02, 0x28, 0x02, 0x52, 0x06, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x58, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x69, 0x78, 0x65, 0x6c, 0x5f, 0x79, 0x18, 0x07, 0x20, 0x02, 0x28, 0x02, 0x52, 0x06, 0x70, 0x69,
	0x78, 0x65, 0x6c, 0x59, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0xa2, 0x02, 0x0a,
	0x12, 0x53, 0x53, 0x4c, 0x5f, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72,
	0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x5f, 0x63, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x01, 0x52, 0x08, 0x74, 0x43, 0x61, 0x70, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x74, 0x5f, 0x73, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x01, 0x52, 0x05, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x61,
	0x6d, 0x65, 0x72, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x63,
	0x61, 0x6d, 0x65, 0x72, 0x61, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x05, 0x62, 0x61, 0x6c, 0x6c, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x53, 0x53, 0x4c, 0x5f, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x61, 0x6c, 0x6c, 0x52, 0x05, 0x62, 0x61, 0x6c, 0x6c,
	0x73, 0x12, 0x38, 0x0a, 0x0d, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x5f, 0x79, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x53, 0x53, 0x4c, 0x5f, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x52, 0x0c, 0x72,
	0x6f, 0x62, 0x6f, 0x74, 0x73, 0x59, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x34, 0x0a, 0x0b, 0x72,
	0x6f, 0x62, 0x6f, 0x74, 0x73, 0x5f, 0x62, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x53, 0x53, 0x4c, 0x5f, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x6f, 0x62, 0x6f, 0x74, 0x52, 0x0a, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x73, 0x42, 0x6c, 0x75,
	0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x52, 0x6f, 0x62, 0x6f, 0x43, 0x75, 0x70, 0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73, 0x6c, 0x2d,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
}

var (
	file_ssl_vision_detection_proto_rawDescOnce sync.Once
	file_ssl_vision_detection_proto_rawDescData = file_ssl_vision_detection_proto_rawDesc
)

func file_ssl_vision_detection_proto_rawDescGZIP() []byte {
	file_ssl_vision_detection_proto_rawDescOnce.Do(func() {
		file_ssl_vision_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_vision_detection_proto_rawDescData)
	})
	return file_ssl_vision_detection_proto_rawDescData
}

var file_ssl_vision_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ssl_vision_detection_proto_goTypes = []interface{}{
	(*SSL_DetectionBall)(nil),  // 0: SSL_DetectionBall
	(*SSL_DetectionRobot)(nil), // 1: SSL_DetectionRobot
	(*SSL_DetectionFrame)(nil), // 2: SSL_DetectionFrame
}
var file_ssl_vision_detection_proto_depIdxs = []int32{
	0, // 0: SSL_DetectionFrame.balls:type_name -> SSL_DetectionBall
	1, // 1: SSL_DetectionFrame.robots_yellow:type_name -> SSL_DetectionRobot
	1, // 2: SSL_DetectionFrame.robots_blue:type_name -> SSL_DetectionRobot
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ssl_vision_detection_proto_init() }
func file_ssl_vision_detection_proto_init() {
	if File_ssl_vision_detection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ssl_vision_detection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSL_DetectionBall); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ssl_vision_detection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSL_DetectionRobot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_ssl_vision_detection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SSL_DetectionFrame); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ssl_vision_detection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_vision_detection_proto_goTypes,
		DependencyIndexes: file_ssl_vision_detection_proto_depIdxs,
		MessageInfos:      file_ssl_vision_detection_proto_msgTypes,
	}.Build()
	File_ssl_vision_detection_proto = out.File
	file_ssl_vision_detection_proto_rawDesc = nil
	file_ssl_vision_detection_proto_goTypes = nil
	file_ssl_vision_detection_proto_depIdxs = nil
}
