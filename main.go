package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws"
)

func RetrieveSecret(variableName string) {

	// AWS Go SDK does not currently support automatic fetching of region from ec2metadata
	// Create metaSession to fetch the region and supply it to the regular Session

	metaSession, _ := session.NewSession()
	metaClient := ec2metadata.New(metaSession)
	region, _ := metaClient.Region()

	conf := aws.NewConfig().WithRegion(region)

	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config: *conf,
	})
	if err != nil {
		printAndExit(err)
	}

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.

	svc := secretsmanager.New(sess)

	// Get secret value
	req, resp := svc.GetSecretValueRequest(&secretsmanager.GetSecretValueInput{
		SecretId: aws.String(variableName),
	})

	err = req.Send()
	if err != nil { // resp is now filled
		printAndExit(err)
	}

	var secretBytes []byte
	if resp.SecretString != nil {
		secretBytes = []byte(*resp.SecretString)
	} else {
		secretBytes = resp.SecretBinary
	}

	os.Stdout.Write(secretBytes)
}

func main() {
	if len(os.Args) != 2 {
		os.Stderr.Write([]byte("A variable name or version flag must be given as the first and only argument!"))
		os.Exit(-1)
	}

	singleArgument := os.Args[1]
	switch singleArgument {
	case "-v","--version":
		os.Stdout.Write([]byte(VERSION))
	default:
		RetrieveSecret(singleArgument)
	}
}

func printAndExit(err error) {
	os.Stderr.Write([]byte(err.Error()))
	os.Exit(1)
}
