package commands

import (
	"log"

	"github.com/Appkube-awsx/awsx-kms/authenticater"
	"github.com/Appkube-awsx/awsx-kms/client"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/spf13/cobra"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxKmsCmd = &cobra.Command{
	Use:   "GetKmsList",
	Short: "GetKmsList command gets resource Arn",
	Long:  `GetKmsList command gets resource Arn details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command getKmsData started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		env := cmd.PersistentFlags().Lookup("env").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			Listkms(region, acKey, secKey, env, crossAccountRoleArn, externalId)
		}
	},
}

func Listkms(region string, accessKey string, secretKey string, env string, crossAccountRoleArn string, externalId string) *kms.ListKeysOutput {
	log.Println("Getting aws config resource summary")
	kmsClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	kmsRequest := &kms.ListKeysInput{}
	kmsResponse, err := kmsClient.Listkms(kmsRequest)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	log.Println(kmsResponse)
	// for _, reservation := range ec2Response.Reservations {
	// 	for _, instance := range reservation.Instances {
	// 		fmt.Println("ID: ", *instance.InstanceId, " name: ", *instance.Tags[0].Value)
	// 	}
	// }

	return kmsResponse
}

func Execute() {
	err := AwsxKmsCmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {
	// AwsxKmsCmd.AddCommand(kmscmd.GetEC2ConfigCmd)
	// AwsxKmsCmd.AddCommand(kmscmd.GetCostDataCmd)
	// AwsxKmsCmd.AddCommand(kmscmd.GetCostSpikeCmd)
	AwsxKmsCmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxKmsCmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxKmsCmd.PersistentFlags().String("zone", "", "aws region")
	AwsxKmsCmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxKmsCmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxKmsCmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxKmsCmd.PersistentFlags().String("externalId", "", "aws external id auth")
	AwsxKmsCmd.PersistentFlags().String("env", "", "env")

}

// cmd used to get list of EC2 instance's :

//  ./awsx-ec2 --zone=us-east-1 --accessKey=<6f> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>
