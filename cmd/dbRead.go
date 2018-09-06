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

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/spf13/cobra"
)

// dbReadCmd represents the dbRead command
var dbReadCmd = &cobra.Command{
	Use:   "dbRead",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create structs to hold info about new item

		// Initialize a session in us-west-2 that the SDK will use to load
		// credentials from the shared credentials file ~/.aws/credentials.
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-west-2")},
		)

		// Create DynamoDB client
		svc := dynamodb.New(sess)

		result, err := svc.GetItem(&dynamodb.GetItemInput{
			TableName: aws.String("Movies"),
			Key: map[string]*dynamodb.AttributeValue{
				"year": {
					N: aws.String("2015"),
				},
				"title": {
					S: aws.String("The Big New Movie"),
				},
			},
		})

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		item := Item{}

		err = dynamodbattribute.UnmarshalMap(result.Item, &item)

		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
		}

		if item.Title == "" {
			fmt.Println("Could not find 'The Big New Movie' (2015)")
			return
		}

		fmt.Println("Found item:")
		fmt.Println("Year:  ", item.Year)
		fmt.Println("Title: ", item.Title)
		fmt.Println("Plot:  ", item.Info.Plot)
		fmt.Println("Rating:", item.Info.Rating)
	},
}

func init() {
	rootCmd.AddCommand(dbReadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbReadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbReadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
