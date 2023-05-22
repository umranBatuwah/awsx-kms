package kmscmd

import (
	"fmt"

	"github.com/Appkube-awsx/awsx-kms/authenticater"
	"github.com/Appkube-awsx/awsx-kms/controllers"
	"github.com/spf13/cobra"
)

// GetConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		authFlag := authenticater.ChildCommandAuth(cmd)

		keyId, _ := cmd.Flags().GetString("keyId")

		if authFlag {
			controllers.KMSController(keyId, authenticater.ClientAuth)
		}
	},
}

func init() {
	GetConfigDataCmd.Flags().StringP("keyId", "t", "", "Key Id")

	if err := GetConfigDataCmd.MarkFlagRequired("keyId"); err != nil {
		fmt.Println("--keyId or -t is required", err)
	}
}
