package main

import (
	"log"

	flag "github.com/spf13/pflag"
	"github.com/vishvananda/netlink"
)

type flags struct {
	Devices []string
	CleanUp bool
}

func parseFlags() *flags {
	var f flags

	flag.StringSliceVarP(&f.Devices, "device", "d", nil, "network devices to run tc-ping")
	flag.BoolVarP(&f.CleanUp, "cleanup", "c", false, "clean up ebpf program")

	flag.Parse()
	return &f
}

func (f *flags) getDevices() map[int]string {
	links, err := netlink.LinkList()
	if err != nil {
		log.Fatalf("Failed to list links: %v", err)
	}

	m := make(map[int]string)
	if len(f.Devices) == 0 {
		for _, l := range links {
			ifindex, ifname := l.Attrs().Index, l.Attrs().Name
			m[ifindex] = ifname
		}

		return m
	}

	target := make(map[string]struct{}, len(f.Devices))
	for _, dev := range f.Devices {
		target[dev] = struct{}{}
	}
	for _, l := range links {
		ifindex, ifname := l.Attrs().Index, l.Attrs().Name
		if _, ok := target[ifname]; ok {
			m[ifindex] = ifname
		}
	}

	return m
}
