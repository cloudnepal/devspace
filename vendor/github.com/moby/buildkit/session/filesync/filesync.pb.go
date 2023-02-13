// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filesync.proto

package filesync

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/tonistiigi/fsutil/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BytesMessage contains a chunk of byte data
type BytesMessage struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *BytesMessage) Reset()      { *m = BytesMessage{} }
func (*BytesMessage) ProtoMessage() {}
func (*BytesMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1042549f1f24495, []int{0}
}
func (m *BytesMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BytesMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BytesMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BytesMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BytesMessage.Merge(m, src)
}
func (m *BytesMessage) XXX_Size() int {
	return m.Size()
}
func (m *BytesMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BytesMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BytesMessage proto.InternalMessageInfo

func (m *BytesMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*BytesMessage)(nil), "moby.filesync.v1.BytesMessage")
}

func init() { proto.RegisterFile("filesync.proto", fileDescriptor_d1042549f1f24495) }

var fileDescriptor_d1042549f1f24495 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xcb, 0xcc, 0x49,
	0x2d, 0xae, 0xcc, 0x4b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc8, 0xcd, 0x4f, 0xaa,
	0xd4, 0x83, 0x0b, 0x96, 0x19, 0x4a, 0xe9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0x97, 0xe4, 0xe7, 0x65, 0x16, 0x97, 0x64, 0x66, 0xa6, 0x67, 0xea, 0xa7, 0x15, 0x97,
	0x96, 0x64, 0xe6, 0xe8, 0x97, 0x54, 0x16, 0xa4, 0x16, 0xeb, 0x97, 0x67, 0x16, 0xa5, 0x42, 0x0c,
	0x50, 0x52, 0xe2, 0xe2, 0x71, 0xaa, 0x2c, 0x49, 0x2d, 0xf6, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f,
	0x15, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02,
	0xb3, 0x8d, 0x9a, 0x19, 0xb9, 0x38, 0xdc, 0x32, 0x73, 0x52, 0x83, 0x2b, 0xf3, 0x92, 0x85, 0xac,
	0xb8, 0x38, 0x5c, 0x32, 0xd3, 0xd2, 0x9c, 0xf3, 0x0b, 0x2a, 0x85, 0x44, 0xf4, 0x20, 0xc6, 0xea,
	0x81, 0x8d, 0xd5, 0x0b, 0x48, 0x4c, 0xce, 0x4e, 0x2d, 0x91, 0xc2, 0x2a, 0xaa, 0xc1, 0x68, 0xc0,
	0x28, 0x64, 0xcd, 0xc5, 0x19, 0x92, 0x58, 0x14, 0x5c, 0x52, 0x94, 0x9a, 0x98, 0x4b, 0xaa, 0x66,
	0xa3, 0x28, 0xa8, 0x23, 0x52, 0xf3, 0x52, 0x84, 0xfc, 0x90, 0x1c, 0x21, 0xa7, 0x87, 0x1e, 0x06,
	0x7a, 0xc8, 0x3e, 0x92, 0x22, 0x20, 0x0f, 0x32, 0xdb, 0xc9, 0xee, 0xc2, 0x43, 0x39, 0x86, 0x1b,
	0x0f, 0xe5, 0x18, 0x3e, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3,
	0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0xf8, 0xe2, 0x91, 0x1c,
	0xc3, 0x87, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1,
	0x1c, 0x43, 0x14, 0x07, 0xcc, 0xcc, 0x24, 0x36, 0x70, 0x60, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xe6, 0x17, 0x63, 0x59, 0x9f, 0x01, 0x00, 0x00,
}

func (this *BytesMessage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BytesMessage)
	if !ok {
		that2, ok := that.(BytesMessage)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Data, that1.Data) {
		return false
	}
	return true
}
func (this *BytesMessage) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&filesync.BytesMessage{")
	s = append(s, "Data: "+fmt.Sprintf("%#v", this.Data)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringFilesync(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileSyncClient is the client API for FileSync service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileSyncClient interface {
	DiffCopy(ctx context.Context, opts ...grpc.CallOption) (FileSync_DiffCopyClient, error)
	TarStream(ctx context.Context, opts ...grpc.CallOption) (FileSync_TarStreamClient, error)
}

type fileSyncClient struct {
	cc *grpc.ClientConn
}

func NewFileSyncClient(cc *grpc.ClientConn) FileSyncClient {
	return &fileSyncClient{cc}
}

func (c *fileSyncClient) DiffCopy(ctx context.Context, opts ...grpc.CallOption) (FileSync_DiffCopyClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileSync_serviceDesc.Streams[0], "/moby.filesync.v1.FileSync/DiffCopy", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSyncDiffCopyClient{stream}
	return x, nil
}

type FileSync_DiffCopyClient interface {
	Send(*types.Packet) error
	Recv() (*types.Packet, error)
	grpc.ClientStream
}

type fileSyncDiffCopyClient struct {
	grpc.ClientStream
}

func (x *fileSyncDiffCopyClient) Send(m *types.Packet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileSyncDiffCopyClient) Recv() (*types.Packet, error) {
	m := new(types.Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSyncClient) TarStream(ctx context.Context, opts ...grpc.CallOption) (FileSync_TarStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileSync_serviceDesc.Streams[1], "/moby.filesync.v1.FileSync/TarStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSyncTarStreamClient{stream}
	return x, nil
}

type FileSync_TarStreamClient interface {
	Send(*types.Packet) error
	Recv() (*types.Packet, error)
	grpc.ClientStream
}

type fileSyncTarStreamClient struct {
	grpc.ClientStream
}

func (x *fileSyncTarStreamClient) Send(m *types.Packet) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileSyncTarStreamClient) Recv() (*types.Packet, error) {
	m := new(types.Packet)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileSyncServer is the server API for FileSync service.
type FileSyncServer interface {
	DiffCopy(FileSync_DiffCopyServer) error
	TarStream(FileSync_TarStreamServer) error
}

// UnimplementedFileSyncServer can be embedded to have forward compatible implementations.
type UnimplementedFileSyncServer struct {
}

func (*UnimplementedFileSyncServer) DiffCopy(srv FileSync_DiffCopyServer) error {
	return status.Errorf(codes.Unimplemented, "method DiffCopy not implemented")
}
func (*UnimplementedFileSyncServer) TarStream(srv FileSync_TarStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method TarStream not implemented")
}

func RegisterFileSyncServer(s *grpc.Server, srv FileSyncServer) {
	s.RegisterService(&_FileSync_serviceDesc, srv)
}

func _FileSync_DiffCopy_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileSyncServer).DiffCopy(&fileSyncDiffCopyServer{stream})
}

type FileSync_DiffCopyServer interface {
	Send(*types.Packet) error
	Recv() (*types.Packet, error)
	grpc.ServerStream
}

type fileSyncDiffCopyServer struct {
	grpc.ServerStream
}

func (x *fileSyncDiffCopyServer) Send(m *types.Packet) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileSyncDiffCopyServer) Recv() (*types.Packet, error) {
	m := new(types.Packet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileSync_TarStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileSyncServer).TarStream(&fileSyncTarStreamServer{stream})
}

type FileSync_TarStreamServer interface {
	Send(*types.Packet) error
	Recv() (*types.Packet, error)
	grpc.ServerStream
}

type fileSyncTarStreamServer struct {
	grpc.ServerStream
}

func (x *fileSyncTarStreamServer) Send(m *types.Packet) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileSyncTarStreamServer) Recv() (*types.Packet, error) {
	m := new(types.Packet)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FileSync_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moby.filesync.v1.FileSync",
	HandlerType: (*FileSyncServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DiffCopy",
			Handler:       _FileSync_DiffCopy_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "TarStream",
			Handler:       _FileSync_TarStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "filesync.proto",
}

// FileSendClient is the client API for FileSend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileSendClient interface {
	DiffCopy(ctx context.Context, opts ...grpc.CallOption) (FileSend_DiffCopyClient, error)
}

type fileSendClient struct {
	cc *grpc.ClientConn
}

func NewFileSendClient(cc *grpc.ClientConn) FileSendClient {
	return &fileSendClient{cc}
}

func (c *fileSendClient) DiffCopy(ctx context.Context, opts ...grpc.CallOption) (FileSend_DiffCopyClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileSend_serviceDesc.Streams[0], "/moby.filesync.v1.FileSend/DiffCopy", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSendDiffCopyClient{stream}
	return x, nil
}

type FileSend_DiffCopyClient interface {
	Send(*BytesMessage) error
	Recv() (*BytesMessage, error)
	grpc.ClientStream
}

type fileSendDiffCopyClient struct {
	grpc.ClientStream
}

func (x *fileSendDiffCopyClient) Send(m *BytesMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileSendDiffCopyClient) Recv() (*BytesMessage, error) {
	m := new(BytesMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FileSendServer is the server API for FileSend service.
type FileSendServer interface {
	DiffCopy(FileSend_DiffCopyServer) error
}

// UnimplementedFileSendServer can be embedded to have forward compatible implementations.
type UnimplementedFileSendServer struct {
}

func (*UnimplementedFileSendServer) DiffCopy(srv FileSend_DiffCopyServer) error {
	return status.Errorf(codes.Unimplemented, "method DiffCopy not implemented")
}

func RegisterFileSendServer(s *grpc.Server, srv FileSendServer) {
	s.RegisterService(&_FileSend_serviceDesc, srv)
}

func _FileSend_DiffCopy_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileSendServer).DiffCopy(&fileSendDiffCopyServer{stream})
}

type FileSend_DiffCopyServer interface {
	Send(*BytesMessage) error
	Recv() (*BytesMessage, error)
	grpc.ServerStream
}

type fileSendDiffCopyServer struct {
	grpc.ServerStream
}

func (x *fileSendDiffCopyServer) Send(m *BytesMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileSendDiffCopyServer) Recv() (*BytesMessage, error) {
	m := new(BytesMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _FileSend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moby.filesync.v1.FileSend",
	HandlerType: (*FileSendServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DiffCopy",
			Handler:       _FileSend_DiffCopy_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "filesync.proto",
}

func (m *BytesMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BytesMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BytesMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintFilesync(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFilesync(dAtA []byte, offset int, v uint64) int {
	offset -= sovFilesync(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BytesMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovFilesync(uint64(l))
	}
	return n
}

func sovFilesync(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFilesync(x uint64) (n int) {
	return sovFilesync(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *BytesMessage) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&BytesMessage{`,
		`Data:` + fmt.Sprintf("%v", this.Data) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringFilesync(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *BytesMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFilesync
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BytesMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BytesMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilesync
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFilesync
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFilesync
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFilesync(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFilesync
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFilesync(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFilesync
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFilesync
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFilesync
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthFilesync
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFilesync
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFilesync
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFilesync        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFilesync          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFilesync = fmt.Errorf("proto: unexpected end of group")
)