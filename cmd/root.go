package cmd;
import (
	"github.com/spf13/cobra"
	"log"
);

var rootCmd = &cobra.Command{
	Use:"root",
	Short:"Root Command",
	Run:run,
}


func run(cmd *cobra.Command, args []string){
	cmd.Usage();
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}