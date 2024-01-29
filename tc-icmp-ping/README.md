## tc-icmp-ping(第一个自己写的eBPF小工具)
使用 tc 挂载点类型 eBPF 程序实现 ICMP 包响应.

## 使用说明
```
# sudo ./tc-icmp-ping --help
Usage of ./tc-icmp-ping:
  -c, --cleanup          clean up ebpf program
  -d, --device strings   network devices to run tc-ping
pflag: help requested
```

## 使用示例
```
# 指定网卡enp0s3挂载ebpf程序到内核，并在其它机器上ping网卡enp0s3对应IP
# sudo ./tc-icmp-ping -d enp0s3
2024/01/29 03:20:05 Success to replace tc-qdisc for if@2:enp0s3
2024/01/29 03:20:05 Success to add tc-filter ingress for if@2:enp0s3

# 查看ebpf程序debug信息
# make exec
sudo tc exec bpf dbg
Running! Hang up with ^C!

          <idle>-0       [003] dNs.. 1128412.221226: bpf_trace_printk: [action1] IP Packet, proto= 1, src= 20490432, dst= 1765320896
```

## 编译
```
# 更新ebpf c程序后执行
# go generate

# go build
```

## 参考
- [ebpf-icmp-ping](https://github.com/badboy/ebpf-icmp-ping)
- [tc-dump](https://github.com/Asphaltt/tc-dump)
