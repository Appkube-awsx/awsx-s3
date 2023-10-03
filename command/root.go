package command

import (
	"log"
	"os"

	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-s3/command/bucketcmd"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

type S3Bucket struct {
	Bucket interface{} `json:"bucket"`
	Tags   interface{} `json:"tags"`
}

var AwsxS3Cmd = &cobra.Command{
	Use:   "getS3BucketList",
	Short: "getS3BucketList command gets list of S3 buckets",
	Long:  `getS3BucketList command gets list of S3 buckets of an AWS account`,

	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Command s3 getS3BucketList started")

		authFlag, clientAuth, err := authenticate.CommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}
		if authFlag {
			GetBucketList(*clientAuth)
		} else {
			cmd.Help()
			return
		}

	},
}

func GetListBucketWithBucketDetail(auth client.Auth) ([]*s3.ListObjectsV2Output, error) {
	log.Println("getting s3 bucket list")

	client := client.GetClient(auth, client.S3_CLIENT).(*s3.S3)
	request := &s3.ListBucketsInput{}
	response, err := client.ListBuckets(request)
	if err != nil {
		log.Fatalln("Error:in getting  bucket list", err)

	}
	allBuckets := []*s3.ListObjectsV2Output{}
	for _, bucket := range response.Buckets {
		bucketDetail, err := bucketcmd.GetBucketDetails(*bucket.Name, auth)
		if err == nil {
			allBuckets = append(allBuckets, bucketDetail)
		}
		//fmt.Println(*bucketName.Name)
	}
	log.Println(allBuckets)
	return allBuckets, err
}

func GetBucketList(auth client.Auth) ([]S3Bucket, error) {
	log.Println("getting s3 bucket list")

	client := client.GetClient(auth, client.S3_CLIENT).(*s3.S3)
	request := &s3.ListBucketsInput{}
	response, err := client.ListBuckets(request)
	if err != nil {
		log.Println("Error:in getting  bucket list", err)
	}
	allBuckets := []S3Bucket{}
	for _, bucket := range response.Buckets {
		input := &s3.GetBucketTaggingInput{
			Bucket: bucket.Name,
		}
		tagOutput, err := client.GetBucketTagging(input)
		if err != nil {
			log.Println("Error in getting bucket tag", err)
			continue
		}
		s3Bucket := S3Bucket{
			Bucket: bucket,
			Tags:   tagOutput,
		}
		allBuckets = append(allBuckets, s3Bucket)
	}
	log.Println(allBuckets)
	return allBuckets, err
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
	AwsxS3Cmd.PersistentFlags().String("vaultToken", "", "vault token")
	AwsxS3Cmd.PersistentFlags().String("accountId", "", "aws account number")
	AwsxS3Cmd.PersistentFlags().String("zone", "", "aws region")
	AwsxS3Cmd.PersistentFlags().String("accessKey", "", "aws access key")
	AwsxS3Cmd.PersistentFlags().String("secretKey", "", "aws secret key")
	AwsxS3Cmd.PersistentFlags().String("crossAccountRoleArn", "", "aws crossAccountRoleArn is required")
	AwsxS3Cmd.PersistentFlags().String("externalId", "", "aws external id auth")

}
