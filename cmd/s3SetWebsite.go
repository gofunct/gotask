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
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

// s3SetWebsiteCmd represents the s3SetWebsite command
var s3SetWebsiteCmd = &cobra.Command{
	Use:   "s3SetWebsite",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) != 4 {
			exitErrorf("bucket name and index suffix page required\nUsage: %s bucket_name index_page [error_page]",
				filepath.Base(os.Args[0]))
		}

		bucket := fromArgs(os.Args, 1)
		indexSuffix := fromArgs(os.Args, 2)
		errorPage := fromArgs(os.Args, 3)

		// Initialize a session in us-west-2 that the SDK will use to load
		// credentials from the shared credentials file ~/.aws/credentials.
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-west-2")},
		)

		// Create S3 service client
		svc := s3.New(sess)

		// Create SetBucketWebsite parameters based on CLI input
		params := s3.PutBucketWebsiteInput{
			Bucket: aws.String(bucket),
			WebsiteConfiguration: &s3.WebsiteConfiguration{
				IndexDocument: &s3.IndexDocument{
					Suffix: aws.String(indexSuffix),
				},
			},
		}

		// Add the error page if set on CLI
		if len(errorPage) > 0 {
			params.WebsiteConfiguration.ErrorDocument = &s3.ErrorDocument{
				Key: aws.String(errorPage),
			}
		}

		// Set the website configuration on the bucket. Replacing any existing
		// configuration.
		_, err = svc.PutBucketWebsite(&params)
		if err != nil {
			exitErrorf("Unable to set bucket %q website configuration, %v",
				bucket, err)
		}

		fmt.Printf("Successfully set bucket %q website configuration\n", bucket)

	},
}

func init() {
	rootCmd.AddCommand(s3SetWebsiteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// s3SetWebsiteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// s3SetWebsiteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func fromArgs(args []string, idx int) string {
	if len(args) > idx {
		return args[idx]
	}
	return ""
}
