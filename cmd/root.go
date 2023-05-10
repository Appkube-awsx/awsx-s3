package cmd

import (
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-s3/authenticator"
	"github.com/Appkube-awsx/awsx-s3/client"
	"github.com/Appkube-awsx/awsx-s3/cmd/bucketcmd"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var AwsxS3Cmd = &cobra.Command{
	Use:   "getListbucketMetaDataDetails",
	Short: "getListbucketMetaDataDetails command gets resource counts",
	Long:  `getListbucketMetaDataDetails command gets resource counts details of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Command s3  getElementDetails started")
		vaultUrl := cmd.PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			getListBucket(region, crossAccountRoleArn, acKey, secKey, externalId)
		}
	},
}


// json.Unmarshal
func getListBucket(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string) (*s3.ListBucketsOutput, error) {
	log.Println("getting s3 bucket metadata list summary")

	listbucketClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	nameRequest := &s3.ListBucketsInput{}
	listbucketResponse, err := listbucketClient.ListBuckets(nameRequest)
	if err != nil {
		log.Fatalln("Error:in getting  bucket list", err)
	}
	log.Println(listbucketResponse)
	return listbucketResponse, err
}


func Execute() {
	err := AwsxS3Cmd.Execute()
	if err != nil {
		log.Fatal("There was some error while executing the CLI: ", err)
		os.Exit(1)
	}
}

func init() {
	AwsxS3Cmd.AddCommand(bucketcmd.GetConfigDataCmd)

	AwsxS3Cmd.PersistentFlags().String("vaultUrl", "", "vault end point")
	AwsxS3Cmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxS3Cmd.PersistentFlags().String("zone", "", "aws region")
	AwsxS3Cmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxS3Cmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxS3Cmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn is required")
	AwsxS3Cmd.PersistentFlags().String("externalId", "", "aws external id auth")

}
