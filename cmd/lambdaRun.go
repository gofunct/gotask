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
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/spf13/cobra"
)

// labdaRunCmd represents the labdaRun command
var labdaRunCmd = &cobra.Command{
	Use:   "labdaRun",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Create Lambda service client
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		client := lambda.New(sess, &aws.Config{Region: aws.String("us-west-2")})

		// Get the 10 most recent items
		request := getItemsRequest{"time", "descending", 10}

		payload, err := json.Marshal(request)

		if err != nil {
			fmt.Println("Error marshalling MyGetItemsFunction request")
			os.Exit(0)
		}

		result, err := client.Invoke(&lambda.InvokeInput{FunctionName: aws.String("MyGetItemsFunction"), Payload: payload})

		if err != nil {
			fmt.Println("Error calling MyGetItemsFunction")
			os.Exit(0)
		}

		var resp getItemsResponse

		err = json.Unmarshal(result.Payload, &resp)

		if err != nil {
			fmt.Println("Error unmarshalling MyGetItemsFunction response")
			os.Exit(0)
		}

		// If the status code is NOT 200, the call failed
		if resp.StatusCode != 200 {
			fmt.Println("Error getting items, StatusCode: " + strconv.Itoa(resp.StatusCode))
			os.Exit(0)
		}

		// If the result is failure, we got an error
		if resp.Body.Result == "failure" {
			fmt.Println("Failed to get items")
			os.Exit(0)
		}

		// Print out items
		if len(resp.Body.Data) > 0 {
			for i := range resp.Body.Data {
				fmt.Println(resp.Body.Data[i].Item)
			}
		} else {
			fmt.Println("There were no items")
		}
	},
}

func init() {
	rootCmd.AddCommand(labdaRunCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// labdaRunCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// labdaRunCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type getItemsRequest struct {
	SortBy     string
	SortOrder  string
	ItemsToGet int
}

type getItemsResponseError struct {
	Message string `json:"message"`
}

type getItemsResponseData struct {
	Item string `json:"item"`
}

type getItemsResponseBody struct {
	Result string                 `json:"result"`
	Data   []getItemsResponseData `json:"data"`
	Error  getItemsResponseError  `json:"error"`
}

type getItemsResponseHeaders struct {
	ContentType string `json:"Content-Type"`
}

type getItemsResponse struct {
	StatusCode int                     `json:"statusCode"`
	Headers    getItemsResponseHeaders `json:"headers"`
	Body       getItemsResponseBody    `json:"body"`
}
