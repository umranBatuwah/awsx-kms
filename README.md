- [What is awsx-kms](#awsx-kms)
- [How to write plugin subcommand](#how-to-write-plugin-subcommand)
- [How to build / Test](#how-to-build--test)
- [what it does ](#what-it-does)
- [command input](#command-input)
- [command output](#command-output)
- [How to run ](#how-to-run)

# awsx-kms

This is a plugin subcommand for awsx cli ( https://github.com/Appkube-awsx/awsx#awsx ) cli.

For details about awsx commands and how its used in Appkube platform , please refer to the diagram below:

![alt text](https://raw.githubusercontent.com/AppkubeCloud/appkube-architectures/main/LayeredArchitecture-phase2.svg)

This plugin subcommand will implement the Apis' related to KMS services , primarily the following API's:

- getConfigData

This cli collect data from metric/logs/traces of the KMS services and produce the data in a form that Appkube Platform expects.

This CLI , interacts with other Appkube services like Appkube vault , Appkube cloud CMDB so that it can talk with cloud services as
well as filter and sort the information in terms of product/env/ services, so that Appkube platform gets the data that it expects from the cli.

# How to write plugin subcommand

Please refer to the instaruction -
https://github.com/Appkube-awsx/awsx#how-to-write-a-plugin-subcommand

It has detailed instruction on how to write a subcommand plugin , build/test/debug/publish and integrate into the main commmand.

# How to build / Test

            go run main.go
                - Program will print Calling aws-cloudelements on console

            Another way of testing is by running go install command
            go install
            - go install command creates an exe with the name of the module (e.g. awsx-kms) and save it in the GOPATH
            - Now we can execute this command on command prompt as below
           awsx-kms getConfigData --zone=us-east-1 --accessKey=xxxxxxxxxx --secretKey=xxxxxxxxxx --crossAccountRoleArn=xxxxxxxxxx  --externalId=xxxxxxxxxx

# what it does

This subcommand implement the following functionalities -
getConfigData - It will get the resource count summary for a given AWS account id and region.

# command input

1. --valutURL =specifies the URL of the AWS Key Management Service (KMS) customer master key (CMK) that you want to use to encrypt a table.
2. --acountId = specifies the AWS account ID that the kms belongs to.
3. --zone = specifies the AWS region where the kms is located.
4. --accessKey = specifies the AWS access key to use for authentication.
5. --secretKey = specifies the AWS secret key to use for authentication.
6. --crossAccountRoleArn = specifies the Amazon Resource Name (ARN) of the role that allows access to a kms in another account.
7. --external Id = specifies the AWS External id.
8. --keyId= Insert your key id which you craeted in aws account.

# command output

Keys: [
{
KeyArn: "arn:aws:kms:us-east-1:657907747545:key/0130671a-7764-4fe4-85db-2885d7cdfaf7",
KeyId: "0130671a-7764-4fe4-85db-2885d7cdfaf7"
},
{
KeyArn: "arn:aws:kms:us-east-1:657907747545:key/0132e072-d844-42bd-b7fa-ca8cd3ce7a7f",
KeyId: "0132e072-d844-42bd-b7fa-ca8cd3ce7a7f"
}
],

{
KeyMetadata: {
AWSAccountId: "657907747545",
Arn: "arn:aws:kms:us-east-1:657907747545:key/d7193167-a02f-4724-8710-3252c045dfd2",
CreationDate: 2021-03-20 09:59:35.655 +0000 UTC,
CustomerMasterKeySpec: "SYMMETRIC_DEFAULT",
Description: "Default master key that protects my EBS volumes when no other key is defined",
Enabled: true,
EncryptionAlgorithms: ["SYMMETRIC_DEFAULT"],
KeyId: "d7193167-a02f-4724-8710-3252c045dfd2",
KeyManager: "AWS",
KeySpec: "SYMMETRIC_DEFAULT",
KeyState: "Enabled",
KeyUsage: "ENCRYPT_DECRYPT",
MultiRegion: false,
Origin: "AWS_KMS"
}
}

# How to run

From main awsx command , it is called as follows:

```bash
awsx-kms  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>
```

If you build it locally , you can simply run it as standalone command as:

```bash
go run main.go  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
```

# awsx-kms

kms extension

# AWSX Commands for AWSX-KMS Cli's :

1. CMD used to get list of KMS instance's :

```bash
./awsx-kms --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
```

2. CMD used to get Config data (metadata) of AWS KMS instances :

```bash
./awsx-kms --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<> --externalId=<> getConfigData --keyId=<>
```
