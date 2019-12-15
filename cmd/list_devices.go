package cmd

import (
	"github.com/spf13/cobra"
	"github.com/jakhax/packet-analyzer/capture"
)

var listDevicesCmd = &cobra.Command{
	Use:"list_devices",
	Short: "Listen network devices",
	RunE:func(cmd *cobra.Command, args []string) (err error){
		capture.ListDevices()
		return
	},
}


func init(){
	rootCmd.AddCommand(listDevicesCmd)
}