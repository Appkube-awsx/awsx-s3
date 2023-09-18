package bucketcmd

import (
	"fmt"
	"github.com/Appkube-awsx/awsx-common/authenticate"
	"log"

	"github.com/Appkube-awsx/awsx-common/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

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
				GetBucketDetails(bucketName, *clientAuth)
			} else {
				log.Fatalln("bucket name not provided. program exit")
			}
		}

	},
}

func GetBucketDetails(bucketName string, auth client.Auth) (*s3.ListObjectsV2Output, error) {
	log.Println("Getting aws bucket data")
	listbucketClient := client.GetClient(auth, client.S3_CLIENT).(*s3.S3)
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}
	bucketDetailsResponse, err := listbucketClient.ListObjectsV2(input)
	log.Println(bucketDetailsResponse.String())
	if err != nil {
		log.Fatalln("Error:", err)
		return nil, err
	}
	return bucketDetailsResponse, nil
}

func init() {
	GetConfigDataCmd.Flags().StringP("bucketName", "b", "", "Bucket name")

	if err := GetConfigDataCmd.MarkFlagRequired("bucketName"); err != nil {
		fmt.Println(err)
	}
}
