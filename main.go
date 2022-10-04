package main

import (
	"fmt"
	"net"
	"regexp"
	"time"

	"github.com/hideckies/pingo/sub"
)

type Pingo struct {
	Count    int
	Host     string
	Interval time.Duration
	IPAddr   *net.IPAddr
	Packet   sub.Packet
}

func (p *Pingo) Resolve(addr string) (*net.IPAddr, error) {
	ipaddr, err := net.ResolveIPAddr("ip4", addr)
	if err != nil {
		fmt.Println("ip address not resolved")
		return nil, err
	}

	return ipaddr, nil
}

func (p *Pingo) ValidInterval(interval string) bool {
	r, _ := regexp.Compile(`^([1-9][0-9]*|0)`)
	return r.MatchString(interval)
}

func NewPingo(f sub.Flag) Pingo {
	var p Pingo
	p.Count = f.Count
	p.Host = f.Target

	if !p.ValidInterval(f.Interval) {
		fmt.Println(sub.ERROR_INCORRECT_VALUE_INTERVAL)
		f.Interval = "1"
	}
	interval, err := time.ParseDuration(f.Interval + "s")
	if err == nil {
		p.Interval = interval
	} else {
		fmt.Println(sub.ERROR_INCORRECT_VALUE_INTERVAL)
		p.Interval = 1 * time.Second
	}

	return p
}

func main() {
	var f sub.Flag

	err := f.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	p := NewPingo(f)
	p.IPAddr, err = p.Resolve(f.Target)
	if err != nil {
		return
	}

	fmt.Printf("pingo %s (%v)\n", p.Host, p.IPAddr)

	c := 1
	for range time.Tick(p.Interval) {
		fmt.Printf("pingo %d times\n", c)
		c++
		if p.Count != 0 && c > p.Count {
			break
		}
	}
}
