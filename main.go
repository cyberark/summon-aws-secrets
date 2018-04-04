package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func RetrieveSecret(variableName string) {
	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.Must(session.NewSession())

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.

	svc := secretsmanager.New(sess)

	// Uploads the object to S3. The Context will interrupt the request if the
	// timeout expires.
	req, resp := svc.GetSecretValueRequest(&secretsmanager.GetSecretValueInput{
		SecretId: &variableName,
	})

	err := req.Send()
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
