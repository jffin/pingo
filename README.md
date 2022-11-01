# Pingo

ping in Go.

<br />

## Usage

```sh
pingo example.com

# sends pingo 5 times
pingo -c 5 example.com

# 2 seconds per pingo
pingo -i 2
```

<br />

## Capabilities

If you feel annoying to 'sudo' every time you run, it encourages to set the capabilities as follow.

```sh
setcap cap_net_raw+ep ./pingo
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