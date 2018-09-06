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
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

// s3PolicyNewCmd represents the s3PolicyNew command
var s3PolicyNewCmd = &cobra.Command{
	Use:   "s3PolicyNew",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) != 2 {
			exitErrorf("bucket name required\nUsage: %s bucket_name",
				filepath.Base(os.Args[0]))
		}
		bucket := os.Args[1]

		// Initialize a session in us-west-2 that the SDK will use to load
		// credentials from the shared credentials file ~/.aws/credentials.
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-west-2")},
		)

		// Create S3 service client
		svc := s3.New(sess)

		// Create a policy using map interface. Filling in the bucket as the
		// resource.
		readOnlyAnonUserPolicy := map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				{
					"Sid":       "AddPerm",
					"Effect":    "Allow",
					"Principal": "*",
					"Action": []string{
						"s3:GetObject",
					},
					"Resource": []string{
						fmt.Sprintf("arn:aws:s3:::%s/*", bucket),
					},
				},
			},
		}

		// Marshal the policy into a JSON value so that it can be sent to S3.
		policy, err := json.Marshal(readOnlyAnonUserPolicy)
		if err != nil {
			exitErrorf("Failed to marshal policy, %v", err)
		}

		// Call S3 to put the policy for the bucket.
		_, err = svc.PutBucketPolicy(&s3.PutBucketPolicyInput{
			Bucket: aws.String(bucket),
			Policy: aws.String(string(policy)),
		})
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok && aerr.Code() == s3.ErrCodeNoSuchBucket {
				// Special error handling for the when the bucket doesn't
				// exists so we can give a more direct error message from the CLI.
				exitErrorf("Bucket %q does not exist", bucket)
			}
			exitErrorf("Unable to set bucket %q policy, %v", bucket, err)
		}

		fmt.Printf("Successfully set bucket %q's policy\n", bucket)

	},
}

func init() {
	rootCmd.AddCommand(s3PolicyNewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// s3PolicyNewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// s3PolicyNewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
