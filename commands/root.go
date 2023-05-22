package commands

import (
	"log"

	"github.com/Appkube-awsx/awsx-kms/authenticater"
	"github.com/Appkube-awsx/awsx-kms/commands/kmscmd"
	"github.com/Appkube-awsx/awsx-kms/controllers"
	"github.com/spf13/cobra"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxKmsCmd = &cobra.Command{
	Use:   "kms",
	Short: "get kms Details command gets resource counts",
	Long:  `get kms Details command gets resource counts details of an AWS account`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command kms started")

		// check for cli flags
		authFlag := authenticater.RootCommandAuth(cmd)

		if authFlag {
			controllers.ListKeys(authenticater.ClientAuth)
		}
	},
}

// Execute runs the command
func Execute() {
	err := AwsxKmsCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {
	AwsxKmsCmd.AddCommand(kmscmd.GetConfigDataCmd)

	// Define persistent flags for the command

	AwsxKmsCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxKmsCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxKmsCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxKmsCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxKmsCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxKmsCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxKmsCmd.PersistentFlags().String("externalId", "", "aws external id auth")
}
