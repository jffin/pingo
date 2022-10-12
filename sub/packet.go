package sub

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"os"
	"sync/atomic"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

const (
	PROTO_NUM_ICMP_IPv4 = 1
	PROTO_NUM_ICMP_IPv6 = 58
)

type Packet struct {
	Data     string
	ICMPType icmp.Type
	ID       int
	Proto    string
	ProtoNum int
	Seq      int
	TTL      int

	SrcAddr  *net.IPAddr
	DestAddr *net.IPAddr
}

func NewPacket(f Flag) *Packet {
	var icmpType icmp.Type
	var proto string
	var protoNum int
	var srcIp string

	if f.UseIPv6 {
		icmpType = ipv6.ICMPTypeEchoRequest
		proto = "ip6"
		protoNum = PROTO_NUM_ICMP_IPv6
		srcIp = "::"
	} else {
		icmpType = ipv4.ICMPTypeEcho
		proto = "ip4"
		protoNum = PROTO_NUM_ICMP_IPv4
		srcIp = "0.0.0.0"
	}

	srcAddr, err := net.ResolveIPAddr(proto, srcIp)
	if err != nil {
		fmt.Println("source address cannot be resolved")
		os.Exit(0)
	}

	destAddr, err := net.ResolveIPAddr(proto, f.Target)
	if err != nil {
		fmt.Println("destination address cannot be resolved")
		os.Exit(0)
	}

	r := rand.New(rand.NewSource(getSeed()))

	return &Packet{
		Data:     f.Data,
		ICMPType: icmpType,
		ID:       r.Intn(math.MaxUint16),
		Proto:    proto,
		ProtoNum: protoNum,
		Seq:      0,
		TTL:      f.TTL,
		SrcAddr:  srcAddr,
		DestAddr: destAddr,
	}
}

var seed int64 = time.Now().UnixNano()

func getSeed() int64 {
	return atomic.AddInt64(&seed, 1)
}
