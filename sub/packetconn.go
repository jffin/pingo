package sub

import "golang.org/x/net/icmp"

type packetconn interface {
	Close() error
	ICMPRequestType() icmp.Type
}
