package services

import (
	"log"

	"github.com/aws/aws-sdk-go/service/kms"
)

// Listkms retrieves details of list of keys from AWS Key Management Service (KMS)

func Listkms(kmsClient *kms.KMS) *kms.ListKeysOutput {
	log.Println("Getting aws config resource summary")

	// Prepare the ListKeyInput request

	kmsRequest := &kms.ListKeysInput{}

	// Send the ListKey request to AWS KMS

	kmsResponse, err := kmsClient.ListKeys(kmsRequest)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	log.Println(kmsResponse)

	return kmsResponse
}
