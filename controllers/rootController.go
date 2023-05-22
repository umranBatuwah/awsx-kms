package controllers

import (
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-kms/authenticater"
	"github.com/Appkube-awsx/awsx-kms/services"
	"github.com/aws/aws-sdk-go/service/kms"
)

func ListKeys(auth client.Auth) *kms.ListKeysOutput {

	// this is Api auth and compulsory for every controller
	authenticater.ApiAuth(auth)

	// Kms client

	kmsClient := client.GetClient(auth, client.KMS_CLIENT).(*kms.KMS)

	kms := services.Listkms(kmsClient)
	return kms
}
