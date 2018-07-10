package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func RetrieveSecret(variableName string) {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.

	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		printAndExit(err)
	}

	creds := credentials.NewChainCredentials(
	[]credentials.Provider{
		&ec2rolecreds.EC2RoleProvider{
			Client: ec2metadata.New(sess),
		},
		&credentials.EnvProvider{},
	})

	// Create a new instance of the service's client with a Session.

	svc := secretsmanager.New(session.Must(session.NewSession(&aws.Config{
		Credentials: creds,
	})))

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
	case "-v", "--version":
		os.Stdout.Write([]byte(VERSION))
	default:
		RetrieveSecret(singleArgument)
	}
}

func printAndExit(err error) {
	os.Stderr.Write([]byte(err.Error()))
	os.Exit(1)
}
