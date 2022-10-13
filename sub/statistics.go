package sub

import (
	"fmt"
	"net"
	"strings"
)

type Statistics struct {
	Proto string

	Transmitted int
	Received    int
	Loss        int
}

func (s *Statistics) Result(peer net.Addr, n int, packet Packet) {
	fmt.Printf(":-) from=%v::bytes=%d::id=0x%x::seq=%d::ttl=%d::protocol=%s\n", peer, n, packet.ID, packet.Seq, packet.TTL, s.Proto)
}

func (s *Statistics) FinalResult() {
	fmt.Println()
	fmt.Println("--- statistics ---")
	fmt.Printf("transmitted=%d::received=%d::loss=%d::protocol=%s\n", s.Transmitted, s.Received, s.Loss, s.Proto)
}

func NewStatistics(p *Packet) *Statistics {
	proto := "icmp"
	if strings.Contains(p.Network, "udp") {
		proto = "udp"
	}

	return &Statistics{
		Proto:       proto,
		Transmitted: 0,
		Received:    0,
		Loss:        0,
	}
}
