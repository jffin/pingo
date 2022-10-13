package sub

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
)

type Flag struct {
	Count        int
	Data         string
	Help         bool
	Interval     string
	Target       string
	TTL          int
	Unprivileged bool
	UseIPv4      bool
	UseIPv6      bool
	Verbose      bool
	Version      bool
}

var version = "0.1.0"

var usage = fmt.Sprintf(`pingo v%s - ping in Go

USAGE:
  -c <int>	ping <count> times (default: 0 -> infinite)
  -d <str>	custom data string (default: "PINGO")
  -h		show usage
  -i <int>	interval per ping (default: 1)
  -t <int>	set TTL (time to live) of the packet (default: 64)
  -u		unprivileged (UDP) ping (default: false)
  -v		verbose mode
  -4		use IPv4 (default: true)
  -6		use IPv6 (default: false)

  help		show usage
  version	show version

EXAMPLES:
  pingo example.com
  pingo -c 5 example.com
  pingo -i 2 example.com
`, version)

func (f *Flag) Parse() error {
	flag.IntVar(&f.Count, "c", 0, "ping <count> times")
	flag.StringVar(&f.Data, "d", "PINGO", "given data string")
	flag.BoolVar(&f.Help, "h", false, "show usage")
	flag.StringVar(&f.Interval, "i", "1", "interval (second) per ping")
	flag.IntVar(&f.TTL, "t", 64, "set TTL (time to live) of the packet")
	flag.BoolVar(&f.Unprivileged, "u", false, "unprivileged (UDP) ping")
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

	// validate interval
	if !validateInterval(f.Interval) {
		fmt.Println(ERROR_INCORRECT_VALUE_INTERVAL)
		f.Interval = "1"
	}

	f.Target = flag.Arg(0)

	return nil
}

// Validate interval
func validateInterval(interval string) bool {
	r, _ := regexp.Compile(`^([1-9][0-9]*|0)`)
	return r.MatchString(interval)
}
