- [What is awsx-s3](#awsx-s3)
- [How to write plugin subcommand](#how-to-write-plugin-subcommand)
- [How to build / Test](#how-to-build--test)
- [what it does ](#what-it-does)
- [command input](#command-input)
- [command output](#command-output)
- [How to run ](#how-to-run)

# awsx-s3

This is a plugin subcommand for awsx cli ( https://github.com/Appkube-awsx/awsx#awsx ) cli.

For details about awsx commands and how its used in Appkube platform , please refer to the diagram below:

![alt text](https://raw.githubusercontent.com/AppkubeCloud/appkube-architectures/main/LayeredArchitecture-phase2.svg)

This plugin subcommand will implement the Apis' related to S3 services , primarily the following API's:

- getConfigData

This cli collect data from metric / logs / traces of the S3 services and produce the data in a form that Appkube Platform expects.

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
            - go install command creates an exe with the name of the module (e.g. awsx-s3) and save it in the GOPATH
            - Now we can execute this command on command prompt as below
           awsx-s3 getConfigData --zone=us-east-1 --accessKey=xxxxxxxxxx --secretKey=xxxxxxxxxx --crossAccountRoleArn=xxxxxxxxxx  --externalId=xxxxxxxxxx

# what it does

This subcommand implement the following functionalities -
getConfigData - It will get the resource count summary for a given AWS account id and region.

# command input

1. --valutURL = URL location of vault - that stores credentials to call API
2. --acountId = The AWS account id.
3. --zone = AWS region
4. --accessKey = Access key for the AWS account
5. --secretKey = Secret Key for the Aws Account
6. --crossAccountRoleArn = Cross Acount Rols Arn for the account.
7. --external Id = The AWS External id.
8. --bucketName= Insert your bucket name which you craeted in aws account.

# command output

Buckets: [
{
CreationDate: 2021-10-26 13:20:59 +0000 UTC,
Name: "26oct-guardduty-bucket"
},
{
CreationDate: 2022-04-27 10:24:59 +0000 UTC,
Name: "acc-request"
},
{
CreationDate: 2019-09-30 08:05:09 +0000 UTC,
Name: "albfirewallogs"
},
{
CreationDate: 2021-11-12 12:09:53 +0000 UTC,
Name: "amplify-quick-notes-devd-120945-deployment"
},
{
CreationDate: 2021-11-12 12:13:58 +0000 UTC,
Name: "amplify-quick-notes-devp-121350-deployment"
},
{
CreationDate: 2021-11-12 12:30:51 +0000 UTC,
Name: "amplify-quick-notes-devs-123044-deployment"
},
],
Owner: {
DisplayName: "papu.bhattacharya",
ID: "de94b6df89e78dfa8b78c500a7180a64798d0a79e7ccd7e18f8797bb87e3f06f"
}

# How to run

From main awsx command , it is called as follows:

```bash
awsx-s3  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<>  --externalId=<>
```

If you build it locally , you can simply run it as standalone command as:

```bash
go run main.go  --zone=us-east-1 --accessKey=<> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
```

# awsx-s3

s3 extension

# AWSX Commands for AWSX-S3 Cli's :

1. CMD used to get list of s3 instance's :

```bash
./awsx-s3 --zone=us-east-1 --accessKey=<6f> --secretKey=<> --crossAccountRoleArn=<> --externalId=<>
```

2. CMD used to get Config data (metadata) of AWS S3 instances :

```bash
./awsx-s3 --zone=us-east-1 --accessKey=<#6f> --secretKey=<> --crossAccountRoleArn=<> --externalId=<> getConfigData --bucketName=<>
```
