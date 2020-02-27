package main

import (
	"fmt"
	"strconv"
	"strings"
)

// IPAddr - type
type IPAddr [4]byte

func (ip IPAddr) String() string {
	rt := make([]string, len(ip))

	for i, val := range ip {
		rt[i] = strconv.Itoa(int(val))
	}

	return strings.Join(rt, ".")
}

// Stringer - init func
func Stringer() {
	hosts := map[string]IPAddr{
		"loopback": {127, 0, 0, 1},
		"dns":      {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
