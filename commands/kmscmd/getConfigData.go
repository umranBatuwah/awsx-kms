package kmscmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-kms/authenticater"
	"github.com/Appkube-awsx/awsx-kms/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
		// print(authFlag)
		// authFlag := true
		if authFlag {
			keyId, _ := cmd.Flags().GetString("keyId")
			if keyId != "" {
				getKeyDetails(region, crossAccountRoleArn, acKey, secKey, keyId, externalId)
			} else {
				log.Fatalln("keyId not provided. Program exit")
			}
		}
	},
}

func getKeyDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string, keyId string, externalId string) *kms.DescribeKeyOutput {
	log.Println("Getting aws cluster data")
	listKeyClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	input := &kms.DescribeKeyInput{
		KeyId: aws.String(keyId),
		
	}
	keyDetailsResponse, err := listKeyClient.DescribeKey(input)
	log.Println(keyDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return keyDetailsResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("keyId", "t", "", "key id")

	if err := GetConfigDataCmd.MarkFlagRequired("keyId"); err != nil {
		fmt.Println(err)
	}
}
