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
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/spf13/cobra"
)

// lambdaFuncCmd represents the lambdaFunc command
var lambdaFuncCmd = &cobra.Command{
	Use:   "lambdaFunc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		zipFilePtr := flag.String("z", "", "The name of the ZIP file (without the .zip extension)")
		bucketPtr := flag.String("b", "", "the name of bucket to which the ZIP file is uploaded")
		functionPtr := flag.String("f", "", "The name of the Lambda function")
		handlerPtr := flag.String("h", "", "The name of the package.class handling the call")
		resourcePtr := flag.String("a", "", "The ARN of the role that calls the function")
		runtimePtr := flag.String("r", "", "The runtime for the function.")

		flag.Parse()

		zipFile := *zipFilePtr
		bucketName := *bucketPtr
		functionName := *functionPtr
		handler := *handlerPtr
		resourceArn := *resourcePtr
		runtime := *runtimePtr

		if zipFile == "" || bucketName == "" || functionName == "" || handler == "" || resourceArn == "" || runtime == "" {
			fmt.Println("You must supply a zip file name, bucket name, function name, handler, ARN, and runtime.")
			os.Exit(0)
		}

		createFunction(zipFile, bucketName, functionName, handler, resourceArn, runtime)
	},
}

func init() {
	rootCmd.AddCommand(lambdaFuncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// lambdaFuncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// lambdaFuncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func createFunction(zipFileName string, bucketName string, functionName string, handler string, resourceArn string, runtime string) {
	// Initialize a session
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create Lambda service client
	svc := lambda.New(sess, &aws.Config{Region: aws.String("us-west-2")})

	contents, err := ioutil.ReadFile(zipFileName + ".zip")

	if err != nil {
		fmt.Println("Could not read " + zipFileName + ".zip")
		os.Exit(0)
	}

	createCode := &lambda.FunctionCode{
		S3Bucket:        aws.String(bucketName),
		S3Key:           aws.String(zipFileName),
		S3ObjectVersion: aws.String(""),
		ZipFile:         contents,
	}

	createArgs := &lambda.CreateFunctionInput{
		Code:         createCode,
		FunctionName: aws.String(functionName),
		Handler:      aws.String(handler),
		Role:         aws.String(resourceArn),
		Runtime:      aws.String(runtime),
	}

	result, err := svc.CreateFunction(createArgs)

	if err != nil {
		fmt.Println("Cannot create function: " + err.Error())
	} else {
		fmt.Println(result)
	}
}
