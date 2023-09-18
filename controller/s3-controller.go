package controller

import (
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/Appkube-awsx/awsx-s3/command"
	"github.com/Appkube-awsx/awsx-s3/command/bucketcmd"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func GetS3BucketListByAccountNo(vaultUrl string, vaultToken string, accountNo string, region string) ([]*s3.ListObjectsV2Output, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData(vaultUrl, vaultToken, accountNo, region, "", "", "", "")
	return GetS3BucketListByFlagAndClientAuth(authFlag, clientAuth, err)
}

func GetS3BucketListByUserCreds(region string, accessKey string, secretKey string, crossAccountRoleArn string, externalId string) ([]*s3.ListObjectsV2Output, error) {
	authFlag, clientAuth, err := authenticate.AuthenticateData("", "", "", region, accessKey, secretKey, crossAccountRoleArn, externalId)
	return GetS3BucketListByFlagAndClientAuth(authFlag, clientAuth, err)
}
func GetS3BucketListByFlagAndClientAuth(authFlag bool, clientAuth *client.Auth, err error) ([]*s3.ListObjectsV2Output, error) {
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	if !authFlag {
		log.Println(err.Error())
		return nil, err
	}
	response, err := GetS3BucketList(clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetS3BucketList(clientAuth *client.Auth) ([]*s3.ListObjectsV2Output, error) {
	response, err := command.GetListBucket(*clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}

func GetS3BucketDetails(bucketName string, clientAuth *client.Auth) (*s3.ListObjectsV2Output, error) {
	response, err := bucketcmd.GetBucketDetails(bucketName, *clientAuth)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return response, nil
}
