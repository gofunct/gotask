// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/spf13/cobra"
)

// dbScanCmd represents the dbScan command
var dbScanCmd = &cobra.Command{
	Use:   "dbScan",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get the movies with a minimum rating of 8.0 in 2011
		min_rating := 8.0
		year := 2011

		// Initialize a session in us-west-2 that the SDK will use to load
		// credentials from the shared credentials file ~/.aws/credentials.
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-west-2")},
		)

		if err != nil {
			fmt.Println("Got error creating session:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Create DynamoDB client
		svc := dynamodb.New(sess)

		// Create the Expression to fill the input struct with.
		// Get all movies in that year; we'll pull out those with a higher rating later
		filt := expression.Name("year").Equal(expression.Value(year))

		// Or we could get by ratings and pull out those with the right year later
		//    filt := expression.Name("info.rating").GreaterThan(expression.Value(min_rating))

		// Get back the title, year, and rating
		proj := expression.NamesList(expression.Name("title"), expression.Name("year"), expression.Name("info.rating"))

		expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

		if err != nil {
			fmt.Println("Got error building expression:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Build the query input parameters
		params := &dynamodb.ScanInput{
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			FilterExpression:          expr.Filter(),
			ProjectionExpression:      expr.Projection(),
			TableName:                 aws.String("Movies"),
		}

		// Make the DynamoDB Query API call
		result, err := svc.Scan(params)

		if err != nil {
			fmt.Println("Query API call failed:")
			fmt.Println((err.Error()))
			os.Exit(1)
		}

		num_items := 0

		for _, i := range result.Items {
			item := Item{}

			err = dynamodbattribute.UnmarshalMap(i, &item)

			if err != nil {
				fmt.Println("Got error unmarshalling:")
				fmt.Println(err.Error())
				os.Exit(1)
			}

			// Which ones had a higher rating?
			if item.Info.Rating > min_rating {
				// Or it we had filtered by rating previously:
				//   if item.Year == year {
				num_items += 1

				fmt.Println("Title: ", item.Title)
				fmt.Println("Rating:", item.Info.Rating)
				fmt.Println()
			}
		}

		fmt.Println("Found", num_items, "movie(s) with a rating above", min_rating, "in", year)
	},
}

func init() {
	rootCmd.AddCommand(dbScanCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbScanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbScanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
