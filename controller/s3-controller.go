package controller

import (
	"encoding/json"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-s3/command"
	"github.com/Appkube-awsx/awsx-s3/command/bucketcmd"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func GetS3BucketListByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) (*s3.ListBucketsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetS3BucketListByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetS3BucketListByUserCreds(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) (*s3.ListBucketsOutput, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
	return GetS3BucketListByFlagAndClientAuth(authFlag, clientAuth, err)
}
func GetS3BucketListByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) (*s3.ListBucketsOutput, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := S3BucketListController(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func S3BucketListController(clientAuth client.Auth) (*s3.ListBucketsOutput, error) {
	log.Println("Request to get s3 bucket list")
	response, err := command.GetBucketList(clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func S3BucketListWithDetailsController(clientAuth client.Auth) ([]*s3.ListObjectsV2Output, error) {
	log.Println("Request to get s3 bucket list with details")
	response, err := bucketcmd.GetBucketListWithDetails(clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func S3BucketDetailController(bucketName string, clientAuth client.Auth) (*s3.ListObjectsV2Output, error) {
	log.Println("Request to get detail of bucket: ", bucketName)
	response, err := bucketcmd.GetBucketDetail(bucketName, clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func S3BucketListWithTagsController(auth client.Auth) (string, error) {
	log.Println("getting s3 bucket list with tags")
	response, err := command.GetBucketList(auth)
	if err != nil {
		log.Println("Error:in getting  bucket list", err)
		return "", err
	}
	client := client.GetClient(auth, client.S3_CLIENT).(*s3.S3)
	allBuckets := []bucketcmd.S3Bucket{}
	for _, bucket := range response.Buckets {
		input := &s3.GetBucketTaggingInput{
			Bucket: bucket.Name,
		}
		tagOutput, err := client.GetBucketTagging(input)
		if err != nil {
			log.Println("Error in getting bucket tag", err)
			continue
		}
		s3Bucket := bucketcmd.S3Bucket{
			Bucket: bucket,
			Tags:   tagOutput,
		}
		allBuckets = append(allBuckets, s3Bucket)
	}
	jsonData, err := json.Marshal(allBuckets)
	log.Println(string(jsonData))
	return string(jsonData), err
}

func S3BucketWithTagsController(bucketName string, auth client.Auth) (string, error) {
	log.Println("getting s3 bucket with tags. bucket: ", bucketName)
	bucketDetail, err := S3BucketDetailController(bucketName, auth)
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
	s3Bucket := bucketcmd.S3Bucket{
		Bucket: bucketDetail,
		Tags:   tagOutput,
	}
	jsonData, err := json.Marshal(s3Bucket)
	if err != nil {
		log.Println("Error in json marshal of bucket: ", err)
		return "", err
	}
	log.Println(string(jsonData))
	return string(jsonData), err
}
