package s3cmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-s3/authenticater"
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

		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)
		print(authFlag)
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

func getBucketDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string, bucketName string, externalId string) *s3.HeadObjectOutput {
	log.Println("Getting aws s3 bucket data")
	listbucketClient := client.GetClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(bucketName),
	}

	// Call the ListObjectsV2 operation to list objects in the bucket
	output, err := listbucketClient.ListObjectsV2(input)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	var headOutput *s3.HeadObjectOutput
	for _, object := range output.Contents {
		// Create an input object for the HeadObject operation
		headInput := &s3.HeadObjectInput{
			Bucket: aws.String(bucketName),
			Key:    object.Key,
		}

		headOutput, err = listbucketClient.HeadObject(headInput)

		if err != nil {
			fmt.Println("Failed to get object metadata", err)
			continue
		}
	}
	fmt.Println("  Content-Type:", aws.StringValue(headOutput.ContentType))
	fmt.Println("  Content-Encoding:", aws.StringValue(headOutput.ContentEncoding))
	fmt.Println("  Custom Metadata:", aws.StringValueMap(headOutput.Metadata))
	return headOutput
}

func init() {
	GetConfigDataCmd.Flags().StringP("bucketName", "t", "", "bucket name")

	if err := GetConfigDataCmd.MarkFlagRequired("bucketName"); err != nil {
		fmt.Println(err)
	}
}
