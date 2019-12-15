package capture

import (
	"os"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"github.com/google/gopacket/layers"
)

// Opt opt
type Opt struct{
	Devices []string
	BPFFilter string
	Promiscous bool
	SnapshotLen int32 
	Timeout int
	File string
	MaxPackets int 
}
const (
	defaultTimeout = 30
	defaultSnapshotLen = 1024
)
// OK validates
func (opt *Opt) OK() (err error){

	if len(opt.Devices) < 1{
		err = fmt.Errorf("Must Provide device")
		return
	}
	for _,device :=  range opt.Devices{
		ok,errX := IsDeviceExists(device)
		err = errX
		if err != nil{
			return
		}
		if !ok{
			err = fmt.Errorf("Device with name %s does not exist",device)
			return
		}
	}
	if opt.SnapshotLen == 0{
		opt.SnapshotLen = defaultSnapshotLen
	}
	if opt.Timeout == 0 {
		opt.Timeout = defaultTimeout
	}
	return
}

// IsDeviceExists checks if device with name exists
func IsDeviceExists(name string)(ok bool, err error){
	devices , err :=  pcap.FindAllDevs()
	if err != nil {
		return
	}
	for _,device :=  range devices{
		if name == device.Name{
			ok = true
			return
		}
	}
	return
}

// ListDevices prints a list of founf devices
func ListDevices() (err error){
	devices , err :=  pcap.FindAllDevs()
	if err != nil {
		return
	}
	fmt.Println("Found Devices");
	for _,device := range devices{
		fmt.Println("--------------------------")
		fmt.Printf("Device Name: %s\n",device.Name)
		fmt.Printf("Device Description: %s\n",device.Description)
	}
	return
}

// CreatePacketCaptureSources creates packet source
func CreatePacketCaptureSources(opt *Opt) (packetSources []*gopacket.PacketSource,err error){
	packetSources = []*gopacket.PacketSource{}
	for _, device := range opt.Devices{
		pktSourceOpt := &PacketSourceOpt{
			Device:device,
			BPFFilter:opt.BPFFilter,
			Promiscous:opt.Promiscous,
			SnapshotLen:opt.SnapshotLen,
			Timeout:opt.Timeout,
		}
		packetSource,errX := CreatePacketSource(pktSourceOpt)
		if errX != nil{
			err =errX
			return
		} 
		packetSources = append(packetSources,packetSource)
	}
	return
}


func Capture(opt *Opt) (err error){
	packetSources, err := CreatePacketCaptureSources(opt)
	if err != nil{
		return
	}
	var w *pcapgo.Writer
	
	if opt.File != ""{
		fx,errX := os.Create(opt.File)
		defer fx.Close()
		if errX !=  nil{
			err = errX
			return
		}
		w = pcapgo.NewWriter(fx)
		w.WriteFileHeader(uint32(opt.SnapshotLen),layers.LinkTypeEthernet)
		
	}
	packetsCh := make(chan gopacket.Packet)
	stop := make(chan bool)

	getPackets := func(packetSource *gopacket.PacketSource){
		for packet := range packetSource.Packets() {
			select{
				case <- stop:
					return
				default:
					packetsCh <- packet
					
			}
		}
	}
	for _, packetSource :=  range packetSources{
		go getPackets(packetSource)
	}
	count := 1
	for{
		packet := <- packetsCh
		fmt.Println(packet)
		if opt.File != ""{
			w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		}
		count++
		if count > opt.MaxPackets && opt.MaxPackets != 0{
			// stop <- true
			break;
		}
	}
	return

}