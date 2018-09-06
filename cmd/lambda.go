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
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/spf13/cobra"
)

// lambdaCmd represents the lambda command
var lambdaCmd = &cobra.Command{
	Use:   "lambda",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		functionName := flag.String("f", "", "The name of the Lambda function")
		sourceArn := flag.String("a", "", "The ARN of the entity invoking the function")

		addNotification(functionName, sourceArn)
	},
}

func init() {
	rootCmd.AddCommand(lambdaCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lambdaCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lambdaCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addNotification(functionName *string, sourceArn *string) {
	// Initialize a session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create Lambda service client
	svc := lambda.New(sess, &aws.Config{Region: aws.String("us-west-2")})

	permArgs := &lambda.AddPermissionInput{
		Action:       aws.String("lambda:InvokeFunction"),
		FunctionName: functionName,
		Principal:    aws.String("s3.amazonaws.com"),
		SourceArn:    sourceArn,
		StatementId:  aws.String("lambda_s3_notification"),
	}

	result, err := svc.AddPermission(permArgs)

	if err != nil {
		fmt.Println("Cannot configure function for notifications")
		os.Exit(0)
	} else {
		fmt.Println(result)
	}
}
