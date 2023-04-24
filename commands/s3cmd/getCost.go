package s3cmd

import (
	"fmt"
	"log"

	"github.com/Appkube-awsx/awsx-s3/authenticater"
	"github.com/Appkube-awsx/awsx-s3/client"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/spf13/cobra"
)

// getConfigDataCmd represents the getConfigData command
var GetCostDataCmd = &cobra.Command{
	Use:   "getCostData",
	Short: "Get cost data",
	Long:  `Retrieve cost data from AWS Cost Explorer`,
	Run: func(cmd *cobra.Command, args []string) {
		// Retrieve values of other flags as before
		vaultUrl := cmd.Parent().PersistentFlags().Lookup("vaultUrl").Value.String()
		accountNo := cmd.Parent().PersistentFlags().Lookup("accountId").Value.String()
		region := cmd.Parent().PersistentFlags().Lookup("zone").Value.String()
		acKey := cmd.Parent().PersistentFlags().Lookup("accessKey").Value.String()
		secKey := cmd.Parent().PersistentFlags().Lookup("secretKey").Value.String()
		crossAccountRoleArn := cmd.Parent().PersistentFlags().Lookup("crossAccountRoleArn").Value.String()
		externalId := cmd.Parent().PersistentFlags().Lookup("externalId").Value.String()

		// Retrieve value of granularity flag
		granularity, err := cmd.Flags().GetString("granularity")
		// Retireve values of start and end date/time
		startDate, err := cmd.Flags().GetString("startDate")
		endDate, err := cmd.Flags().GetString("endDate")
	
		if err != nil {
			log.Fatalln("Error: in getting granularity flag value", err)
		}
		authFlag := authenticater.AuthenticateData(vaultUrl, accountNo, region, acKey, secKey, crossAccountRoleArn, externalId)

		if authFlag {
			s3CostDetails(region, crossAccountRoleArn, acKey, secKey, externalId, granularity, startDate, endDate)
		}
	},
}

func s3CostDetails(region string, crossAccountRoleArn string, accessKey string, secretKey string, externalId string, granularity string, startDate string, endDate string) (*costexplorer.GetCostAndUsageOutput, error) {
	log.Printf("Getting cost data with granularity %s", granularity)

	costClient := client.GetCostClient(region, crossAccountRoleArn, accessKey, secretKey, externalId)

	var start, end string
	switch granularity {
	case "DAILY":
		start = startDate //"2023-03-01"
		end = endDate //"2023-03-02"
	case "WEEKLY":
		start = startDate //"2022-08-01"
		end = endDate //"2022-09-08"
	case "MONTHLY":
		start = startDate //"2022-08-01"
		end = endDate //"2022-10-01"
	case "HOURLY":
		start = startDate //"2022-08-01T00:00:00Z"
		end = endDate //"2022-08-01T02:00:00Z"
	default:
		return nil, fmt.Errorf("unsupported granularity: %s", granularity)
	}

	input := &costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(start),
			End:   aws.String(end),
		},
		Metrics: []*string{
			aws.String("UNBLENDED_COST"),
			aws.String("BLENDED_COST"),
			aws.String("AMORTIZED_COST"),
		},
		GroupBy: []*costexplorer.GroupDefinition{
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("SERVICE"),
			},
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("RECORD_TYPE"),
			},
		},
		Granularity: aws.String(granularity),
		Filter: &costexplorer.Expression{
			Dimensions: &costexplorer.DimensionValues{
				Key: aws.String("SERVICE"),
				Values: []*string{
					aws.String("Amazon Simple Storage Service"),
				},
			},
		},
	}

	costData, err := costClient.GetCostAndUsage(input)
	if err != nil {
		log.Fatalln("Error: in getting cost data", err)
	}
	log.Println(costData)
	return costData, err
}


func init() {

}
