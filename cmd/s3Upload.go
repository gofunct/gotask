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
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/cobra"
)

// s3UploadCmd represents the s3Upload command
var s3UploadCmd = &cobra.Command{
	Use:   "s3Upload",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) != 3 {
			exitErrorf("bucket and file name required\nUsage: %s bucket_name filename",
				os.Args[0])
		}

		bucket := os.Args[1]
		filename := os.Args[2]

		file, err := os.Open(filename)
		if err != nil {
			exitErrorf("Unable to open file %q, %v", err)
		}

		defer file.Close()

		// Initialize a session in us-west-2 that the SDK will use to load
		// credentials from the shared credentials file ~/.aws/credentials.
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String("us-west-2")},
		)

		// Setup the S3 Upload Manager. Also see the SDK doc for the Upload Manager
		// for more information on configuring part size, and concurrency.
		//
		// http://docs.aws.amazon.com/sdk-for-go/api/service/s3/s3manager/#NewUploader
		uploader := s3manager.NewUploader(sess)

		// Upload the file's body to S3 bucket as an object with the key being the
		// same as the filename.
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucket),

			// Can also use the `filepath` standard library package to modify the
			// filename as need for an S3 object key. Such as turning absolute path
			// to a relative path.
			Key: aws.String(filename),

			// The file to be uploaded. io.ReadSeeker is preferred as the Uploader
			// will be able to optimize memory when uploading large content. io.Reader
			// is supported, but will require buffering of the reader's bytes for
			// each part.
			Body: file,
		})
		if err != nil {
			// Print the error and exit.
			exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
		}

		fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)

	},
}

func init() {
	rootCmd.AddCommand(s3UploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// s3UploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// s3UploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
