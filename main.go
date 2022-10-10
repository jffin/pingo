package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	"github.com/hideckies/pingo/sub"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

type Pingo struct {
	Count    int
	Host     string
	Interval time.Duration
	Packet   *sub.Packet

	// Channel
	done chan interface{}
	mtx  sync.Mutex
}

// Run
// Reference: https://pkg.go.dev/golang.org/x/net@v0.0.0-20221004154528-8021a29435af/icmp#example-PacketConn-NonPrivilegedPing
func (p *Pingo) Run(statistics *sub.Statistics) error {
	packetconn, err := icmp.ListenPacket(p.Packet.Proto+":"+strconv.Itoa(p.Packet.ProtoNum), p.Packet.SrcAddr.String())
	if err != nil {
		log.Fatalf("ICMP ListenPacket Error: %v\n", err)
	}
	defer packetconn.Close()

	pktconn := packetconn.IPv4PacketConn()
	if err := pktconn.SetTTL(p.Packet.TTL); err != nil {
		log.Fatalf("SetTTL Error: %v\n", err)
	}

	c := 1
	for range time.Tick(p.Interval) {
		p.Packet.Seq = c
		body := &icmp.Echo{
			ID: p.Packet.ID,
			// ID: os.Getpid() & 0xffff,
			Seq:  p.Packet.Seq,
			Data: []byte("PINGO"),
		}
		msg := &icmp.Message{
			Type: p.Packet.ICMPType,
			Code: 0,
			Body: body,
		}
		wb, err := msg.Marshal(nil)
		if err != nil {
			log.Fatalf("Marshal Error: %v\n", err)
		}
		if _, err := packetconn.WriteTo(wb, p.Packet.DestAddr); err != nil {
			log.Fatalf("WriteTo Error: %v\n", err)
		}

		rb := make([]byte, 1500)
		n, peer, err := packetconn.ReadFrom(rb)
		if err != nil {
			log.Fatalf("ReadFrom Error: %v\n", err)
		}

		rm, err := icmp.ParseMessage(p.Packet.ProtoNum, rb[:n])
		if err != nil {
			log.Fatal(err)
		}

		resultReceived := fmt.Sprintf(":-) from=%v::bytes=%d::id=0x%x::seq=%d::ttl=%d\n", peer, n, p.Packet.ID, p.Packet.Seq, p.Packet.TTL)

		switch rm.Type {
		case ipv4.ICMPTypeEchoReply:
			fmt.Print(resultReceived)
			statistics.Received++
		case ipv6.ICMPTypeEchoReply:
			fmt.Print(resultReceived)
			statistics.Received++
		default:
			fmt.Printf(":-< faled %+v\n", rm)
			statistics.Loss++
		}

		c++
		statistics.Transmitted++

		if p.Count != 0 && c > p.Count {
			break
		}
	}

	return nil
}

func (p *Pingo) Stop(statistics *sub.Statistics) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	open := true
	select {
	case _, open = <-p.done:
	default:
	}

	if open {
		close(p.done)
	}

	// display the result
	statistics.ShowResult()

	os.Exit(0)
}

func NewPingo(flag sub.Flag, packet sub.Packet) *Pingo {
	var p Pingo
	p.Count = flag.Count
	p.Host = flag.Target
	p.Packet = &packet
	p.done = make(chan interface{})

	interval, err := time.ParseDuration(flag.Interval + "s")
	if err == nil {
		p.Interval = interval
	} else {
		fmt.Println(sub.ERROR_INCORRECT_VALUE_INTERVAL)
		p.Interval = 1 * time.Second
	}

	return &p
}

func main() {
	var f sub.Flag

	err := f.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	packet := sub.NewPacket(f)
	pingo := NewPingo(f, *packet)

	fmt.Printf("pingo %s (%s)\n", pingo.Host, pingo.Packet.DestAddr.String())

	statistics := sub.NewStatistics()

	// Listen for Ctrl+c signal
	cch := make(chan os.Signal, 1)
	signal.Notify(cch, os.Interrupt)
	go func() {
		for range cch {
			pingo.Stop(statistics)
		}
	}()

	err = pingo.Run(statistics)
	if err != nil {
		fmt.Println("Error pingo")
	}

	statistics.ShowResult()
}
