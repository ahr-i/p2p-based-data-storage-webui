package myLocal

import (
	"net"
)

func GetBroadcastAddress() string {
	const CIDR = 16 // 사이더

	ip_str := GetLocalIP()
	mask_str := GetSubnetMask(CIDR)

	ip := net.ParseIP(ip_str).To4()
	if ip == nil {
		return ""
	}

	mask := net.IPMask(net.ParseIP(mask_str).To4())
	if len(mask) == 0 {
		return ""
	}

	broadcast := make(net.IP, 4)
	for i := 0; i < 4; i++ {
		broadcast[i] = ip[i] | ^mask[i]
	}

	return broadcast.String()
}
