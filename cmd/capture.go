package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/jakhax/packet-analyzer/capture"
)

var captureCmd = &cobra.Command{
	Use:"capture",
	Short: "Capture packets",
	RunE:func(cmd *cobra.Command, args []string)(err error){
		devices,err := cmd.Flags().GetStringArray("devices")
		if err != nil{
			return
		}
		file, err := cmd.Flags().GetString("file")
		if err != nil{
			return
		}
		filter, err := cmd.Flags().GetString("filter")
		if err != nil{
			return
		}
		maxPackets,err := cmd.Flags().GetInt("max_packets")
		if err !=  nil{
			return
		}
		if maxPackets < 0  {
			err = fmt.Errorf("Max packets cannot be <= 0")
			return 
		}
		captureOpt := &capture.Opt{
			Devices:devices,
			File: file,
			BPFFilter: filter,
			MaxPackets: maxPackets,
		}
		err = captureOpt.OK()
		if err != nil{
			return
		}
		err = capture.Capture(captureOpt)
		return
	},
}

func init(){
	rootCmd.AddCommand(captureCmd)
	captureCmd.Flags().StringArrayP("devices","d",[]string{},"Devices to capture on")
	captureCmd.Flags().StringP("file","f","","Pcap file")
	captureCmd.Flags().StringP("filter","F","","BPF filter")
	captureCmd.Flags().IntP("max_packets","N",0,"Maximum number of packets to capture")
}