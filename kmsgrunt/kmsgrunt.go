package kmsgrunt

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/gruntwork-io/terragrunt/errors"
	"github.com/spf13/viper"
)

// CreateKmsClient returns a new AWS KMS client
func CreateKmsClient() (*kms.KMS, error) {
	viper.SetConfigType("toml")
	viper.SetDefault("awsRegion", "eu-west-1")
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	if len(accessKey) == 0 {
		awsProfileEnv := os.Getenv("AWS_PROFILE")
		if len(awsProfileEnv) != 0 {
			viper.SetDefault("awsProfile", awsProfileEnv)
		} else {
			viper.SetDefault("awsProfile", "default")
		}
	}
	viper.SetConfigName(".kmsgrunt")
	viper.AddConfigPath("/etc/kmsgrunt/") // path to look for the config file in
	viper.AddConfigPath("$HOME/")         // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file

		fmt.Println(errors.WithStackTraceAndPrefix(err, "fatal error. Expecting config file: .kmsgrunt.toml"))
		fmt.Println(errors.WithStackTraceAndPrefix(err, "Using defaults eu-west-1 and default profile"))
	}
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String(viper.GetString("awsRegion"))},
		Profile:           viper.GetString("awsProfile"),
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return nil, errors.WithStackTraceAndPrefix(err, "Error initializing session")
	}

	if viper.IsSet("iamRoleArn") {
		sess.Config.Credentials = stscreds.NewCredentials(sess, viper.GetString("iamRoleArn"))
	}

	_, err = sess.Config.Credentials.Get()
	if err != nil {
		return nil, errors.WithStackTraceAndPrefix(err, "Error finding AWS credentials (did you set the AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY environment variables?)")
	}

	return kms.New(sess), nil
}
