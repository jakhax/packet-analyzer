# Packet Analyzer
- A simple [packet analyzer/sniffer](https://en.wikipedia.org/wiki/Packet_analyzer) , intercepts and log traffic that passes over a network.
- Supports live packet capture from multiple network interfaces (using goroutines) and packet filtering. You can also save the captured traffic to a file and analyze later using a program like wireshark.
- **Why a packet sniffer ?** Just another understanding computers(networking) exercise.
## Usage
```
Available Commands:
  capture      Capture packets
  help         Help about any command
  list_devices Listen network devices

Flags:
  -h, --help   help for root
```
### Capture
```
Capture packets

Usage:
  root capture [flags]

Flags:
  -d, --devices stringArray   Devices to capture on
  -f, --file string           Pcap file
  -F, --filter string         BPF filter
  -h, --help                  help for capture
  -N, --max_packets int       Maximum number of packets to capture
```

#### Example
- To capture traffic on `eth0` and `wlan0` with [BPF](https://biot.com/capstats/bpf.html) filter `tcp and port 8000` and save to file `out.pcap`
```bash
go run main.go capture -d wlan0 -d eth0 -f out.pcap - F 'tcp and port 8000'
```

## Resources
- [https://godoc.org/github.com/google/gopacket](https://godoc.org/github.com/google/gopacket)