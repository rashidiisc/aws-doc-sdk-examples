/*
   Copyright 2010-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.

   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at

    http://aws.amazon.com/apache2.0/

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/

package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

// Usage:
// go run iam_deleteuser.go <username>
func main() {
	// Initialize a session that the SDK will use to load configuration,
	// credentials, and region from the shared config file. (~/.aws/config).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create a IAM service client.
	svc := iam.New(sess)

	_, err := svc.DeleteUser(&iam.DeleteUserInput{
		UserName: &os.Args[1],
	})

	// If the user does not exist than we will log an error.
	if awserr, ok := err.(awserr.Error); ok && awserr.Code() == "NoSuchEntity" {
		fmt.Printf("User %s does not exist\n", os.Args[1])
		return
	} else if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("User %s has been deleted\n", os.Args[1])
}
