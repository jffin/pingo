package sub

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Flag struct {
	Count    int
	Help     bool
	Interval string
	Target   string
	TTL      int
	UseIPv4  bool
	UseIPv6  bool
	Verbose  bool
	Version  bool
}

var version = "0.1.0"

var usage = fmt.Sprintf(`pingo v%s - ping in Go

USAGE:
	-c	ping <count> times (default: infinite)
	-h	show usage
	-i	interval per ping (default: 1)
	-t	set TTL (time to live) of the packet (default: 64)
	-v	verbose mode
	-4	use IPv4 (default: true)
	-6	use IPv6 (default: false)

	help	show usage
	version	show version

EXAMPLES:
	pingo example.com
	pingo -c 5 example.com
	pingo -i 2 example.com
`, version)

func (f *Flag) Parse() error {
	flag.IntVar(&f.Count, "c", 0, "ping <count> times")
	flag.BoolVar(&f.Help, "h", false, "show usage")
	flag.StringVar(&f.Interval, "i", "1", "interval (second) per ping")
	flag.IntVar(&f.TTL, "t", 64, "set TTL (time to live) of the packet")
	flag.BoolVar(&f.Verbose, "v", false, "verbose mode")
	flag.BoolVar(&f.UseIPv4, "4", true, "use IPv4")
	flag.BoolVar(&f.UseIPv6, "6", false, "use IPv6")
	flag.Parse()

	if f.Help || (len(os.Args) == 2 && os.Args[1] == "help") {
		fmt.Println(usage)
		os.Exit(0)
	} else if f.Version || (len(os.Args) == 2 && os.Args[1] == "version") {
		fmt.Printf("pingo v%s\n", version)
		os.Exit(0)
	}

	if len(flag.Args()) != 1 {
		return errors.New(usage)
	}

	f.Target = flag.Arg(0)

	return nil
}
