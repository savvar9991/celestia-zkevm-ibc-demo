// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/groth16/v1/groth16.proto

package groth16

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ClientState defines a groth16 client that holds two keys for verifying state
// transition proofs needed to verify IBC packets
type ClientState struct {
	// latest height of the client state
	LatestHeight uint64 `protobuf:"varint,1,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height,omitempty"`
	// groth16 state transition proof verifier key. Verifies proofs on a rollups
	// state root after the state transition has been applied. Only BN254 curve is
	// supported
	StateTransitionVerifierKey []byte `protobuf:"bytes,2,opt,name=state_transition_verifier_key,json=stateTransitionVerifierKey,proto3" json:"state_transition_verifier_key,omitempty" yaml:"stp_verifier_key"`
	// Provided during initialization of the IBC Client
	CodeCommitment []byte `protobuf:"bytes,3,opt,name=code_commitment,json=codeCommitment,proto3" json:"code_commitment,omitempty"`
	// Provided during initialization of the IBC Client
	GenesisStateRoot []byte `protobuf:"bytes,4,opt,name=genesis_state_root,json=genesisStateRoot,proto3" json:"genesis_state_root,omitempty"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35b13d3299f8e80, []int{0}
}
func (m *ClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(m, src)
}
func (m *ClientState) XXX_Size() int {
	return m.Size()
}
func (m *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(m)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

// ConsensusState defines a groth16 consensus state.
type ConsensusState struct {
	// timestamp that corresponds to the block height in which the ConsensusState
	// was stored.
	Timestamp time.Time `protobuf:"bytes,1,opt,name=timestamp,proto3,stdtime" json:"timestamp"`
	// state root of the rollup
	StateRoot []byte `protobuf:"bytes,2,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty" yaml:"state_root"`
}

func (m *ConsensusState) Reset()         { *m = ConsensusState{} }
func (m *ConsensusState) String() string { return proto.CompactTextString(m) }
func (*ConsensusState) ProtoMessage()    {}
func (*ConsensusState) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35b13d3299f8e80, []int{1}
}
func (m *ConsensusState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsensusState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsensusState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsensusState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusState.Merge(m, src)
}
func (m *ConsensusState) XXX_Size() int {
	return m.Size()
}
func (m *ConsensusState) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusState.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusState proto.InternalMessageInfo

// Header defines a groth16 header for updating the trusted state root of a
// rollup
type Header struct {
	// serialized groth16 proof that the given state transition is valid
	StateTransitionProof []byte `protobuf:"bytes,1,opt,name=state_transition_proof,json=stateTransitionProof,proto3" json:"state_transition_proof,omitempty" yaml:"state_transition_proof"`
	// last verified height of the rollup. This is used to retrieve the previous
	// state root with which the proof is verified against
	TrustedHeight int64 `protobuf:"varint,2,opt,name=trusted_height,json=trustedHeight,proto3" json:"trusted_height,omitempty" yaml:"trusted_height"`
	// trusted header hash passed by the relayer
	TrustedCelestiaHeaderHash []byte `protobuf:"bytes,3,opt,name=trusted_celestia_header_hash,json=trustedCelestiaHeaderHash,proto3" json:"trusted_celestia_header_hash,omitempty" yaml:"trusted_celestia_header_hash"`
	// new state root, height and header hash of the rollup after the state transition has been
	// applied
	NewStateRoot          []byte `protobuf:"bytes,4,opt,name=new_state_root,json=newStateRoot,proto3" json:"new_state_root,omitempty" yaml:"new_state_root"`
	NewHeight             int64  `protobuf:"varint,5,opt,name=new_height,json=newHeight,proto3" json:"new_height,omitempty" yaml:"new_height"`
	NewCelestiaHeaderHash []byte `protobuf:"bytes,6,opt,name=new_celestia_header_hash,json=newCelestiaHeaderHash,proto3" json:"new_celestia_header_hash,omitempty" yaml:"new_celestia_header_hash"`
	// TODO: This is provided by the user at the moment but we can't trust them
	// with this data. We need to get all the data roots from the
	// the store.
	DataRoots [][]byte `protobuf:"bytes,7,rep,name=data_roots,json=dataRoots,proto3" json:"data_roots,omitempty" yaml:"data_roots"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_e35b13d3299f8e80, []int{2}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClientState)(nil), "celestia.ibc.lightclients.groth16.v1.ClientState")
	proto.RegisterType((*ConsensusState)(nil), "celestia.ibc.lightclients.groth16.v1.ConsensusState")
	proto.RegisterType((*Header)(nil), "celestia.ibc.lightclients.groth16.v1.Header")
}

func init() { proto.RegisterFile("proto/groth16/v1/groth16.proto", fileDescriptor_e35b13d3299f8e80) }

var fileDescriptor_e35b13d3299f8e80 = []byte{
	// 625 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0xcf, 0x6e, 0xd3, 0x4e,
	0x18, 0x8c, 0x7f, 0xcd, 0xaf, 0x90, 0x6d, 0x1a, 0xc0, 0x6a, 0xc1, 0x0d, 0xd4, 0x2e, 0x2e, 0x52,
	0x7b, 0xa0, 0xb6, 0x0a, 0x15, 0x87, 0x5e, 0x40, 0xce, 0x81, 0x4a, 0x5c, 0xd0, 0x52, 0x81, 0x84,
	0x10, 0xd6, 0xda, 0xf9, 0x6a, 0xaf, 0x6a, 0x7b, 0x23, 0xef, 0x26, 0x51, 0x79, 0x02, 0x8e, 0x7d,
	0x02, 0xc4, 0xe3, 0xf4, 0xd8, 0x23, 0x27, 0x83, 0xda, 0x07, 0x40, 0xf2, 0x13, 0x20, 0xaf, 0xed,
	0xfc, 0x69, 0x73, 0xb3, 0xbf, 0x19, 0x8f, 0x67, 0xe6, 0x5b, 0x2d, 0xd2, 0x07, 0x29, 0x13, 0xcc,
	0x0e, 0x52, 0x26, 0xc2, 0xfd, 0x57, 0xf6, 0x68, 0xbf, 0x7e, 0xb4, 0x24, 0xa0, 0x3e, 0xf3, 0x21,
	0x02, 0x2e, 0x28, 0xb1, 0xa8, 0xe7, 0x5b, 0x11, 0x0d, 0x42, 0xe1, 0x47, 0x14, 0x12, 0xc1, 0xad,
	0x9a, 0x38, 0xda, 0xef, 0xae, 0x05, 0x2c, 0x60, 0x95, 0x12, 0x0b, 0x58, 0xf9, 0x6d, 0xd7, 0x08,
	0x18, 0x0b, 0x22, 0xb0, 0xe5, 0x9b, 0x37, 0x3c, 0xb1, 0x05, 0x8d, 0x81, 0x0b, 0x12, 0x0f, 0x4a,
	0x82, 0xf9, 0x57, 0x41, 0x2b, 0x3d, 0xa9, 0xf6, 0x41, 0x10, 0x01, 0xea, 0x36, 0x5a, 0x8d, 0x88,
	0x00, 0x2e, 0xdc, 0x10, 0x8a, 0x3f, 0x69, 0xca, 0x96, 0xb2, 0xdb, 0xc4, 0xed, 0x72, 0x78, 0x24,
	0x67, 0xea, 0x57, 0xb4, 0xc9, 0x0b, 0xb6, 0x2b, 0x52, 0x92, 0x70, 0x2a, 0x28, 0x4b, 0xdc, 0x11,
	0xa4, 0xf4, 0x84, 0x42, 0xea, 0x9e, 0xc2, 0x99, 0xf6, 0xdf, 0x96, 0xb2, 0xdb, 0x76, 0x1e, 0xe7,
	0x99, 0xf1, 0xe8, 0x8c, 0xc4, 0xd1, 0xa1, 0xc9, 0xc5, 0x60, 0x8e, 0x61, 0xe2, 0xae, 0x54, 0x38,
	0x9e, 0x08, 0x7c, 0xac, 0xd0, 0x77, 0x70, 0xa6, 0xee, 0xa0, 0x7b, 0x3e, 0xeb, 0x83, 0xeb, 0xb3,
	0x38, 0xa6, 0x22, 0x86, 0x44, 0x68, 0x4b, 0x85, 0x22, 0xee, 0x14, 0xe3, 0xde, 0x64, 0xaa, 0x3e,
	0x47, 0x6a, 0x00, 0x09, 0x70, 0xca, 0xdd, 0xd2, 0x50, 0xca, 0x98, 0xd0, 0x9a, 0x92, 0x7b, 0xbf,
	0x42, 0x64, 0x2e, 0xcc, 0x98, 0x38, 0x6c, 0x7e, 0xff, 0x69, 0x34, 0xcc, 0x73, 0x05, 0x75, 0x7a,
	0x2c, 0xe1, 0x90, 0xf0, 0x61, 0x09, 0xaa, 0x0e, 0x6a, 0x4d, 0x7a, 0x91, 0x81, 0x57, 0x5e, 0x74,
	0xad, 0xb2, 0x39, 0xab, 0x6e, 0xce, 0x3a, 0xae, 0x19, 0xce, 0xdd, 0x8b, 0xcc, 0x68, 0x9c, 0xff,
	0x36, 0x14, 0x3c, 0xfd, 0x4c, 0x3d, 0x40, 0x68, 0xc6, 0x42, 0x59, 0xc0, 0x7a, 0x9e, 0x19, 0x0f,
	0xea, 0x02, 0x6a, 0xcc, 0xc4, 0x2d, 0x7e, 0xc3, 0xd2, 0x8f, 0x26, 0x5a, 0x3e, 0x02, 0xd2, 0x87,
	0x54, 0xfd, 0x84, 0x1e, 0xde, 0xaa, 0x76, 0x90, 0x32, 0x76, 0x22, 0x7d, 0xb5, 0x9d, 0xa7, 0x79,
	0x66, 0x6c, 0xce, 0x4a, 0xde, 0xe4, 0x99, 0x78, 0xed, 0x46, 0xb3, 0xef, 0x8b, 0xb1, 0xfa, 0x06,
	0x75, 0x44, 0x3a, 0xe4, 0x02, 0xfa, 0xf5, 0x66, 0x0b, 0x8f, 0x4b, 0xce, 0x46, 0x9e, 0x19, 0xeb,
	0xa5, 0xe0, 0x3c, 0x6e, 0xe2, 0xd5, 0x6a, 0x50, 0x6d, 0x3d, 0x44, 0x4f, 0x6a, 0x46, 0x7d, 0x22,
	0xdd, 0x50, 0xba, 0x76, 0x43, 0xc2, 0xc3, 0x72, 0x45, 0xce, 0x4e, 0x9e, 0x19, 0xdb, 0xf3, 0x7a,
	0x8b, 0xd8, 0x26, 0xde, 0xa8, 0xe0, 0x5e, 0x85, 0x96, 0x05, 0x1c, 0x11, 0x1e, 0xaa, 0xaf, 0x51,
	0x27, 0x81, 0xf1, 0xad, 0x95, 0xce, 0x7a, 0x9d, 0xc7, 0x4d, 0xdc, 0x4e, 0x60, 0x3c, 0xd9, 0x74,
	0xb1, 0x8c, 0x82, 0x50, 0x05, 0xfd, 0x5f, 0x06, 0x9d, 0x59, 0xc6, 0x14, 0x33, 0x71, 0x2b, 0x81,
	0x71, 0x15, 0xf0, 0x0b, 0xd2, 0x0a, 0x64, 0x61, 0xb8, 0x65, 0x69, 0x60, 0x3b, 0xcf, 0x0c, 0x63,
	0xaa, 0xb1, 0x38, 0xd8, 0x7a, 0x02, 0xe3, 0x05, 0xa1, 0x0e, 0x10, 0xea, 0x13, 0x41, 0xa4, 0x5f,
	0xae, 0xdd, 0xd9, 0x5a, 0x9a, 0x3f, 0x20, 0x53, 0xcc, 0xc4, 0xad, 0xe2, 0xa5, 0x08, 0xc2, 0xcb,
	0x03, 0xe2, 0x90, 0x8b, 0x2b, 0x5d, 0xb9, 0xbc, 0xd2, 0x95, 0x3f, 0x57, 0xba, 0x72, 0x7e, 0xad,
	0x37, 0x2e, 0xaf, 0xf5, 0xc6, 0xaf, 0x6b, 0xbd, 0xf1, 0xf9, 0x6d, 0x40, 0x45, 0x38, 0xf4, 0x2c,
	0x9f, 0xc5, 0x76, 0x6d, 0x87, 0xa5, 0xc1, 0xe4, 0x79, 0xef, 0xdb, 0x29, 0x8c, 0xe2, 0x3d, 0xea,
	0xf9, 0x7b, 0x7d, 0x88, 0x99, 0x4d, 0x3d, 0xdf, 0x9e, 0xbd, 0x42, 0xea, 0xbb, 0xc6, 0x5b, 0x96,
	0x07, 0xfd, 0xe5, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x9f, 0x69, 0x36, 0x8e, 0x04, 0x00,
	0x00,
}

func (m *ClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GenesisStateRoot) > 0 {
		i -= len(m.GenesisStateRoot)
		copy(dAtA[i:], m.GenesisStateRoot)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.GenesisStateRoot)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CodeCommitment) > 0 {
		i -= len(m.CodeCommitment)
		copy(dAtA[i:], m.CodeCommitment)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.CodeCommitment)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.StateTransitionVerifierKey) > 0 {
		i -= len(m.StateTransitionVerifierKey)
		copy(dAtA[i:], m.StateTransitionVerifierKey)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.StateTransitionVerifierKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.LatestHeight != 0 {
		i = encodeVarintGroth16(dAtA, i, uint64(m.LatestHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ConsensusState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsensusState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.StateRoot) > 0 {
		i -= len(m.StateRoot)
		copy(dAtA[i:], m.StateRoot)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.StateRoot)))
		i--
		dAtA[i] = 0x12
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Timestamp, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGroth16(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DataRoots) > 0 {
		for iNdEx := len(m.DataRoots) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.DataRoots[iNdEx])
			copy(dAtA[i:], m.DataRoots[iNdEx])
			i = encodeVarintGroth16(dAtA, i, uint64(len(m.DataRoots[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.NewCelestiaHeaderHash) > 0 {
		i -= len(m.NewCelestiaHeaderHash)
		copy(dAtA[i:], m.NewCelestiaHeaderHash)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.NewCelestiaHeaderHash)))
		i--
		dAtA[i] = 0x32
	}
	if m.NewHeight != 0 {
		i = encodeVarintGroth16(dAtA, i, uint64(m.NewHeight))
		i--
		dAtA[i] = 0x28
	}
	if len(m.NewStateRoot) > 0 {
		i -= len(m.NewStateRoot)
		copy(dAtA[i:], m.NewStateRoot)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.NewStateRoot)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TrustedCelestiaHeaderHash) > 0 {
		i -= len(m.TrustedCelestiaHeaderHash)
		copy(dAtA[i:], m.TrustedCelestiaHeaderHash)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.TrustedCelestiaHeaderHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.TrustedHeight != 0 {
		i = encodeVarintGroth16(dAtA, i, uint64(m.TrustedHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.StateTransitionProof) > 0 {
		i -= len(m.StateTransitionProof)
		copy(dAtA[i:], m.StateTransitionProof)
		i = encodeVarintGroth16(dAtA, i, uint64(len(m.StateTransitionProof)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGroth16(dAtA []byte, offset int, v uint64) int {
	offset -= sovGroth16(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LatestHeight != 0 {
		n += 1 + sovGroth16(uint64(m.LatestHeight))
	}
	l = len(m.StateTransitionVerifierKey)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	l = len(m.CodeCommitment)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	l = len(m.GenesisStateRoot)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	return n
}

func (m *ConsensusState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Timestamp)
	n += 1 + l + sovGroth16(uint64(l))
	l = len(m.StateRoot)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	return n
}

func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.StateTransitionProof)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	if m.TrustedHeight != 0 {
		n += 1 + sovGroth16(uint64(m.TrustedHeight))
	}
	l = len(m.TrustedCelestiaHeaderHash)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	l = len(m.NewStateRoot)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	if m.NewHeight != 0 {
		n += 1 + sovGroth16(uint64(m.NewHeight))
	}
	l = len(m.NewCelestiaHeaderHash)
	if l > 0 {
		n += 1 + l + sovGroth16(uint64(l))
	}
	if len(m.DataRoots) > 0 {
		for _, b := range m.DataRoots {
			l = len(b)
			n += 1 + l + sovGroth16(uint64(l))
		}
	}
	return n
}

func sovGroth16(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGroth16(x uint64) (n int) {
	return sovGroth16(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroth16
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
			return fmt.Errorf("proto: ClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHeight", wireType)
			}
			m.LatestHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LatestHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateTransitionVerifierKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
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
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateTransitionVerifierKey = append(m.StateTransitionVerifierKey[:0], dAtA[iNdEx:postIndex]...)
			if m.StateTransitionVerifierKey == nil {
				m.StateTransitionVerifierKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeCommitment", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
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
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CodeCommitment = append(m.CodeCommitment[:0], dAtA[iNdEx:postIndex]...)
			if m.CodeCommitment == nil {
				m.CodeCommitment = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisStateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
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
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisStateRoot = append(m.GenesisStateRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.GenesisStateRoot == nil {
				m.GenesisStateRoot = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroth16(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroth16
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
func (m *ConsensusState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroth16
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
			return fmt.Errorf("proto: ConsensusState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsensusState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Timestamp, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
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
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateRoot = append(m.StateRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.StateRoot == nil {
				m.StateRoot = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroth16(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroth16
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
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGroth16
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateTransitionProof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
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
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateTransitionProof = append(m.StateTransitionProof[:0], dAtA[iNdEx:postIndex]...)
			if m.StateTransitionProof == nil {
				m.StateTransitionProof = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedHeight", wireType)
			}
			m.TrustedHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TrustedHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedCelestiaHeaderHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
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
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrustedCelestiaHeaderHash = append(m.TrustedCelestiaHeaderHash[:0], dAtA[iNdEx:postIndex]...)
			if m.TrustedCelestiaHeaderHash == nil {
				m.TrustedCelestiaHeaderHash = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewStateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
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
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewStateRoot = append(m.NewStateRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.NewStateRoot == nil {
				m.NewStateRoot = []byte{}
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewHeight", wireType)
			}
			m.NewHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewCelestiaHeaderHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
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
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewCelestiaHeaderHash = append(m.NewCelestiaHeaderHash[:0], dAtA[iNdEx:postIndex]...)
			if m.NewCelestiaHeaderHash == nil {
				m.NewCelestiaHeaderHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataRoots", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGroth16
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
				return ErrInvalidLengthGroth16
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGroth16
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataRoots = append(m.DataRoots, make([]byte, postIndex-iNdEx))
			copy(m.DataRoots[len(m.DataRoots)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGroth16(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGroth16
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
func skipGroth16(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGroth16
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
					return 0, ErrIntOverflowGroth16
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
					return 0, ErrIntOverflowGroth16
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
				return 0, ErrInvalidLengthGroth16
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGroth16
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGroth16
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGroth16        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGroth16          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGroth16 = fmt.Errorf("proto: unexpected end of group")
)
