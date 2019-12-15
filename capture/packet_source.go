package capture

import (
	"time"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type PacketSourceOpt struct{
	Device string
	BPFFilter string
	Promiscous bool
	SnapshotLen int32 
	Timeout int
	File string
}

// CreatePacketSource creates packet source
func CreatePacketSource(opt *PacketSourceOpt) (packetSource *gopacket.PacketSource,err error){
	var handle *pcap.Handle
	timeout := time.Duration(opt.Timeout) * time.Second
	if opt.File == ""{
		
		handle,err = pcap.OpenLive(opt.Device,opt.SnapshotLen,opt.Promiscous,timeout)
	}else{
		handle , err = pcap.OpenOffline(opt.File)
	}
	if err !=  nil{
		return
	}
	if opt.BPFFilter !=  ""{
		err = handle.SetBPFFilter(opt.BPFFilter)
		if err != nil{
			return
		}
	}
	packetSource = gopacket.NewPacketSource(handle, handle.LinkType())
	return
}
