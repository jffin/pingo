package sub

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"sync/atomic"
	"time"
)

type Packet struct {
	ID       int
	Proto    string
	SrcAddr  *net.IPAddr
	DestAddr *net.IPAddr
}

func NewPacket(f Flag) *Packet {
	var p Packet

	var srcIp string

	if f.UseIPv6 {
		p.Proto = "ip6"
		srcIp = "::"
	} else {
		p.Proto = "ip4"
		srcIp = "0.0.0.0"
	}

	r := rand.New(rand.NewSource(getSeed()))
	p.ID = r.Intn(math.MaxUint16)

	srcAddr, err := net.ResolveIPAddr(p.Proto, srcIp)
	if err != nil {
		fmt.Println("source address cannot be resolved")
	}
	p.SrcAddr = srcAddr

	destAddr, err := net.ResolveIPAddr(p.Proto, f.Target)
	if err != nil {
		fmt.Println("destination address cannot be resolved")
	}
	p.DestAddr = destAddr

	return &p
}

var seed int64 = time.Now().UnixNano()

func getSeed() int64 {
	return atomic.AddInt64(&seed, 1)
}
