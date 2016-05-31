package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/colonelmo/grpc-chat/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve lis.ten.add.ress",
	Short: "Listens on a given ip:port",
	Long: `
Creates a chatroom to which clients can connect using 'chat connect'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			err := server.Serve(args[0], false)
			if err != nil {
				log.Fatalln("Error: %s", err.Error())
			}
		} else {
			log.Fatalln("Incorrect number of arguments")
		}
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
