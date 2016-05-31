package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/colonelmo/grpc-chat/client"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:   "connect ser.ver.add.ress:port your_nickname",
	Short: "Connects to a remote server",
	Long: `
Connects to a remote chat server and registers you with the given nickname.
You can then send messages to be broadcasted to other participants. Other
Participants' messages will be broadcasted to you too.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			err := client.Connect(args[0], args[1], false)
			if err != nil {
				log.Fatalln(err.Error)
			}
		} else {
			log.Fatalln("Incorrect number of args")
		}
	},
}

func init() {
	RootCmd.AddCommand(connectCmd)
}
