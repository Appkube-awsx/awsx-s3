package bucketcmd

import (
	"encoding/json"
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"log"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

type S3Bucket struct {
	Bucket interface{} `json:"bucket"`
	Tags   interface{} `json:"tags"`
}

// GetConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "Get S3 bucket configuration",
	Long:  `Get S3 bucket configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		authFlag, clientAuth, err := authenticate.SubCommandAuth(cmd)
		if err != nil {
			cmd.Help()
			return
		}

		if authFlag {
			bucketName, _ := cmd.Flags().GetString("bucketName")
			if bucketName != "" {
				GetBucketDetail(bucketName, *clientAuth)
			} else {
				log.Fatalln("bucket name not provided. program exit")
			}
		}

	},
}

func GetBucketDetail(bucketName string, auth client.Auth) (*s3.ListObjectsV2Output, error) {
	log.Println("Getting details of bucket: ", bucketName)
	client := client.GetClient(auth, client.S3_CLIENT).(*s3.S3)
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}
	response, err := client.ListObjectsV2(input)
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}
	log.Println(response)
	return response, nil
}

func GetBucketDetailWithTags(bucketName string, auth client.Auth) (string, error) {
	log.Println("getting s3 bucket with tags. bucket: ", bucketName)
	bucketDetail, err := GetBucketDetail(bucketName, auth)
	if err != nil {
		log.Println("Error in getting bucket detail", err)
		return "", err
	}
	client := client.GetClient(auth, client.S3_CLIENT).(*s3.S3)
	input := &s3.GetBucketTaggingInput{
		Bucket: aws.String(bucketName),
	}
	tagOutput, err := client.GetBucketTagging(input)
	if err != nil {
		log.Println("Error in getting bucket tag: ", err)
		return "", err
	}
	s3Bucket := S3Bucket{
		Bucket: bucketDetail,
		Tags:   tagOutput,
	}
	jsonData, err := json.Marshal(s3Bucket)
	if err != nil {
		log.Println("Error in getting bucket tag", err)
		return "", err
	}
	log.Println(string(jsonData))
	return string(jsonData), err
}

func GetBucketListWithDetails(auth client.Auth) ([]*s3.ListObjectsV2Output, error) {
	log.Println("getting s3 bucket list with details")
	client := client.GetClient(auth, client.S3_CLIENT).(*s3.S3)
	request := &s3.ListBucketsInput{}
	response, err := client.ListBuckets(request)
	if err != nil {
		log.Println("Error:in getting  bucket list", err)
		return nil, err
	}

	allBuckets := []*s3.ListObjectsV2Output{}
	for _, bucket := range response.Buckets {
		bucketDetail, err := GetBucketDetail(*bucket.Name, auth)
		if err != nil {
			log.Println("Error in getting bucket detail ", err)
			continue
		}
		allBuckets = append(allBuckets, bucketDetail)
	}
	log.Println(allBuckets)
	return allBuckets, err
}

func init() {
	GetConfigDataCmd.Flags().StringP("bucketName", "b", "", "Bucket name")

	if err := GetConfigDataCmd.MarkFlagRequired("bucketName"); err != nil {
		fmt.Println(err)
	}
}
