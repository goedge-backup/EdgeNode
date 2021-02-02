package iplibrary

import "github.com/TeaOSLab/EdgeNode/internal/utils"

type IPItemType = string

const (
	IPItemTypeIPv4 IPItemType = "ipv4" // IPv4
	IPItemTypeIPv6 IPItemType = "ipv6" // IPv6
	IPItemTypeAll  IPItemType = "all"  // 所有IP
)

// IP条目
type IPItem struct {
	Type      string
	Id        int64
	IPFrom    uint64
	IPTo      uint64
	ExpiredAt int64
}

// 检查是否包含某个IP
func (this *IPItem) Contains(ip uint64) bool {
	switch this.Type {
	case IPItemTypeIPv4:
		return this.containsIPv4(ip)
	case IPItemTypeIPv6:
		return this.containsIPv6(ip)
	case IPItemTypeAll:
		return this.containsAll(ip)
	default:
		return this.containsIPv4(ip)
	}
}

// 检查是否包含某个IPv4
func (this *IPItem) containsIPv4(ip uint64) bool {
	if this.IPTo == 0 {
		if this.IPFrom != ip {
			return false
		}
	} else {
		if this.IPFrom > ip || this.IPTo < ip {
			return false
		}
	}
	if this.ExpiredAt > 0 && this.ExpiredAt < utils.UnixTime() {
		return false
	}
	return true
}

// 检查是否包含某个IPv6
func (this *IPItem) containsIPv6(ip uint64) bool {
	if this.IPFrom != ip {
		return false
	}
	if this.ExpiredAt > 0 && this.ExpiredAt < utils.UnixTime() {
		return false
	}
	return true
}

// 检查是否包所有IP
func (this *IPItem) containsAll(ip uint64) bool {
	if this.ExpiredAt > 0 && this.ExpiredAt < utils.UnixTime() {
		return false
	}
	return true
}
