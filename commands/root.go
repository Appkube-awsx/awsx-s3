/*
Copyright © 2023 Manoj Sharma manoj.sharma@synectiks.com
*/
package commands

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-s3/authenticater"
	"github.com/Appkube-awsx/awsx-s3/client"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

// AwsxCloudElementsCmd represents the base command when called without any subcommands
var AwsxS3Cmd = &cobra.Command{
	Use:   "s3 Buckets info",
	Short: "get s3 Details command gets resource counts",
	Long:  `get s3 Details command gets resource counts details of an AWS account`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Command s3 started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()

		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			ListS3(region, acKey, secKey, crossAccountRoleArn, externalId)
		}

	},
}

func ListS3(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) *s3.ListBucketsOutput {
	log.Println("Getting s3 buckets  list")
	s3Client := client.GetClient(region, accessKey, secretKey, crossAccountRoleArn, externalId)
	input := &s3.ListBucketsInput{}
	result, err := s3Client.ListBuckets(input)
	if err != nil {
		log.Println("Error listing clusters:", err)
		return nil
	}

	// print the cluster ARNs to console
	for _, bucket := range result.Buckets {
		fmt.Println(aws.StringValue(bucket.Name))
	}

	log.Println(result)

	// return the result object
	return result
}

func Execute() {
	err := AwsxS3Cmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		return
	}
}

func init() {

	// Register command with root command
	//AwsxS3Cmd.AddCommand(GetCostDataCmd)
	AwsxS3Cmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxS3Cmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxS3Cmd.PersistentFlags().String("zone", "", "aws region")
	AwsxS3Cmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxS3Cmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxS3Cmd.PersistentFlags().String("crossAccountRoleArn", "", "aws cross account role arn")
	AwsxS3Cmd.PersistentFlags().String("granularity", "", " auth")
	AwsxS3Cmd.PersistentFlags().String("externalId", "", "aws external id auth")
	AwsxS3Cmd.PersistentFlags().String("startDate", "", "start date/time")
	AwsxS3Cmd.PersistentFlags().String("endDate", "", "end date/time")
	AwsxS3Cmd.PersistentFlags().String("serviceName", "", "aws service name")
}
