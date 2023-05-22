package services

import (
	"log"

	"github.com/aws/aws-sdk-go/service/kms"
)

// KMSDetail retrieves details of a specific key with help of keyId from AWS Key Management Service (KMS)

func KMSDetail(kmsClient *kms.KMS, keyId string) *kms.DescribeKeyOutput {
	log.Println("Getting aws config resource summary")

	// Prepare the DescribeKeyInput request

	kmsRequest := &kms.DescribeKeyInput{
		KeyId: &keyId,
	}

	// Send the DescribeKey request to AWS KMS

	kmsResponse, err := kmsClient.DescribeKey(kmsRequest)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	log.Println(kmsResponse)
	log.Println(kmsResponse)


	return kmsResponse
}
