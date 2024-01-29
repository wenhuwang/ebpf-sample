package main

import (
	"errors"
	"log"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/rlimit"
)

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -cc clang tcPing ./ebpf/tc_ping.c

func main() {
	// Remove resource limits for kernels <5.11.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal("Removing memlock:", err)
	}

	flags := parseFlags()
	devs := flags.getDevices()

	specTc, err := loadTcPing()
	if err != nil {
		log.Fatalf("Failed to load bpf spec: %v", err)
	}

	for idx := range devs {
		ifindex, ifname := idx, devs[idx]

		var obj tcPingObjects
		if specTc.LoadAndAssign(&obj, &ebpf.CollectionOptions{}); err != nil {
			var ve *ebpf.VerifierError
			if errors.As(err, &ve) {
				log.Printf("Failed to load bpf obj for if@%d:%s: %v\n%+v", ifindex, ifname, err, ve)
			}
			log.Printf("Failed to load bpf obj for if@%d:%s: %v\n%+v", ifindex, ifname, err, ve)
		}
		defer obj.Close()

		// 清理tc qdisc && filter
		deleteTcQdisc(ifindex)
		deleteTcFilterIngress(ifindex, obj.Pingpong)

		if flags.CleanUp {
			continue
		}

		// 创建或者更新tc qdisc
		if err := replaceTcQdisc(ifindex); err != nil {
			log.Printf("Failed to replace tc-qdisc for if@%d:%s: %v", ifindex, ifname, err)
			return
		}
		log.Printf("Success to replace tc-qdisc for if@%d:%s", ifindex, ifname)

		// 添加tc filter
		if err := addTcFilterIngress(ifindex, obj.Pingpong); err != nil {
			log.Printf("Failed to add tc-filter ingress for if@%d:%s: %v", ifindex, ifname, err)
			return
		}
		log.Printf("Success to add tc-filter ingress for if@%d:%s", ifindex, ifname)
	}

}
