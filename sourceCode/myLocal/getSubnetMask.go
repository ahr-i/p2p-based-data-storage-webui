package myLocal

import "fmt"

func GetSubnetMask(cidr_prefix int) string {
	/* Error */
	if cidr_prefix < 0 || cidr_prefix > 32 {
		return ""
	}

	mask := uint32((1 << uint32(cidr_prefix)) - 1)
	mask = mask << uint32(32-cidr_prefix)

	/* 4개의 옥텟 반환 */
	return fmt.Sprintf("%d.%d.%d.%d",
		byte((mask>>24)&0xFF),
		byte((mask>>16)&0xFF),
		byte((mask>>8)&0xFF),
		byte(mask&0xFF))
}
