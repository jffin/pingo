# Pingo

ping in Go.

<br />

## Usage

```
USAGE:
  -c <int>      ping <count> times (default: 0 -> infinite)
  -d <str>      custom data string (default: "PINGO")
  -h            show usage
  -i <int>      interval per ping (default: 1)
  -t <int>      set TTL (time to live) of the packet (default: 64)
  -u            unprivileged (UDP) ping (default: false)
  -v            verbose mode
  -4            use IPv4 (default: true)
  -6            use IPv6 (default: false)

  help          show usage
  version       show version

EXAMPLES:
  pingo example.com
  pingo -c 5 example.com
  pingo -i 2 example.com
```

<br />

## Unprivileged Ping

If you want to the unprivileged (UDP) ping in Linux, set the following sysctl command.

```sh
sudo sysctl -w net.ipv4.ping_group_range="0 2147483647"
```

Then run the pingo with the "-u" flag.

```sh
pingo -u example.com
```

<br />

## Install & Build

```sh
go get ; go build
```