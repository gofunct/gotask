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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/spf13/cobra"
)

// dbLoadCmd represents the dbLoad command
var dbLoadCmd = &cobra.Command{
	Use:   "dbLoad",
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

		if err != nil {
			fmt.Println("Error creating session:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Create DynamoDB client
		svc := dynamodb.New(sess)

		// Get table items from movie_data.json
		items := getItems()

		// Add each item to Movies table:
		for _, item := range items {
			av, err := dynamodbattribute.MarshalMap(item)

			if err != nil {
				fmt.Println("Got error marshalling map:")
				fmt.Println(err.Error())
				os.Exit(1)
			}

			// Create item in table Movies
			input := &dynamodb.PutItemInput{
				Item:      av,
				TableName: aws.String("Movies"),
			}

			_, err = svc.PutItem(input)

			if err != nil {
				fmt.Println("Got error calling PutItem:")
				fmt.Println(err.Error())
				os.Exit(1)
			}

			fmt.Println("Successfully added '", item.Title, "' (", item.Year, ") to Movies table")
		}

	},
}

func init() {
	rootCmd.AddCommand(dbLoadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbLoadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbLoadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type ItemInfo struct {
	Plot   string  `json:"plot"`
	Rating float64 `json:"rating"`
}

type Item struct {
	Year  int      `json:"year"`
	Title string   `json:"title"`
	Info  ItemInfo `json:"info"`
}

func getItems() []Item {
	raw, err := ioutil.ReadFile("./movie_data.json")

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var items []Item
	json.Unmarshal(raw, &items)
	return items
}
