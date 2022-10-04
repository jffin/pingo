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
	Version  bool
}

var version = "0.1.0"

var usage = fmt.Sprintf(`pingo v%s - ping in Go

USAGE:
	-c	ping <count> times
	-h	show usage
	-i	interval per ping
	-v	show version

EXAMPLES:
	pingo example.com
	pingo -c 5 example.com
	pingo -i 2 example.com
`, version)

func (f *Flag) Parse() error {
	flag.IntVar(&f.Count, "c", 0, "ping <count> times")
	flag.BoolVar(&f.Help, "h", false, "show usage")
	flag.StringVar(&f.Interval, "i", "1", "interval (second) per ping")
	flag.BoolVar(&f.Version, "v", false, "show version")
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
