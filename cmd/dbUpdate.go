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
	"github.com/spf13/cobra"
)

// dbUpdateCmd represents the dbUpdate command
var dbUpdateCmd = &cobra.Command{
	Use:   "dbUpdate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize a session in us-west-2 that the SDK will use to load
		// credentials from the shared credentials file ~/.aws/credentials.
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-west-2")},
		)

		// Create DynamoDB client
		svc := dynamodb.New(sess)

		// Create item in table Movies
		input := &dynamodb.UpdateItemInput{
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":r": {
					N: aws.String("0.5"),
				},
			},
			TableName: aws.String("Movies"),
			Key: map[string]*dynamodb.AttributeValue{
				"year": {
					N: aws.String("2015"),
				},
				"title": {
					S: aws.String("The Big New Movie"),
				},
			},
			ReturnValues:     aws.String("UPDATED_NEW"),
			UpdateExpression: aws.String("set info.rating = :r"),
		}

		_, err = svc.UpdateItem(input)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Successfully updated 'The Big New Movie' (2015) rating to 0.5")
	},
}

func init() {
	rootCmd.AddCommand(dbUpdateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbUpdateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbUpdateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
