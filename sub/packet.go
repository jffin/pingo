package sub

type Packet struct {
	IPHeader   IPHeader
	ICMPHeader ICMPHeader
}

type IPHeader struct {
	Version        []byte
	IHL            []byte
	ToS            []byte
	TotalLengh     []byte
	Identification []byte
	Flags          []byte
	FragmentOffset []byte
	TTL            []byte
	Protocol       []byte
	Checksum       []byte
	SrcAddr        []byte
	DestAddr       []byte
}

type ICMPHeader struct {
	TypeOfMessage []byte
	Code          []byte
	Checksum      []byte
	Identifier    []byte
	SeqNum        []byte
}
