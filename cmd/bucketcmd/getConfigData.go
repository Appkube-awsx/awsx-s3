package bucketcmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-s3/authenticator"
	"github.com/Appkube-awsx/awsx-s3/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetConfigDataCmd = &cobra.Command{
	Use:   "getConfigData",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		authFlag := authenticator.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
		// print(authFlag)
		// authFlag := true
		if authFlag {
			bucketName, _ := cmd.Flags().GetString("bucketName")
			if bucketName != "" {
				getBucketDetails(region, crossAccountRoleArn, acKey, secKey, bucketName, externalId)
			} else {
				log.Fatalln("bucketName not provided. Program exit")
			}
		}
	},
}

func getBucketDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string, bucketName string, externalId string) *s3.ListObjectsV2Output {
	log.Println("Getting aws bucket data")
	listbucketClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}
	bucketDetailsResponse, err := listbucketClient.ListObjectsV2(input)
	log.Println(bucketDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
	}
	return bucketDetailsResponse
}

func init() {
	GetConfigDataCmd.Flags().StringP("bucketName", "t", "", "Cluster name")

	if err := GetConfigDataCmd.MarkFlagRequired("bucketName"); err != nil {
		fmt.Println(err)
	}
}
