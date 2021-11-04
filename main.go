package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func ExtractSecretPathAndArguments(input string) (secretPath string, keyName string, secretVersion string) {
	keyName = ""
	secretVersion = ""
	secretPath = ""

	if strings.ContainsRune(input, 35) && strings.ContainsRune(input, 94) {
		fields := strings.FieldsFunc(input, func(r rune) bool {
			return r == '#' || r == '^'
		})
		return fields[0], fields[1], fields[2]

	} else if strings.ContainsRune(input, 35) {
		fields := strings.SplitN(input, "#", 2)
		return fields[0], fields[1], ""
	} else if strings.ContainsRune(input, 94) {
		fields := strings.SplitN(input, "^", 2)
		return fields[0], "", fields[1]
	}

	return input, "", ""
}

func RetrieveSecret(variableName string) {
	// By default get secrets that contain the version AWSCURRENT
	secretVersion := "AWSCURRENT"
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

	// AWS Go SDK does not currently support automatic fetching of region from ec2metadata.
	// If the region could not be found in an environment variable or a shared config file,
	// create metaSession to fetch the ec2 instance region and pass to the regular Session.
	if *sess.Config.Region == "" {
		metaSession, err := session.NewSession()
		if err != nil {
			printAndExit(err)
		}

		metaClient := ec2metadata.New(metaSession)
		// If running on an EC2 instance, the metaClient will be available and we can set the region to match the instance
		// If not on an EC2 instance, the region will remain blank and AWS returns a "MissingRegion: ..." error
		if metaClient.Available() {
			if region, err := metaClient.Region(); err == nil {
				sess.Config.Region = aws.String(region)
			} else {
				printAndExit(err)
			}
		}
	}

	// Create a new instance of the service's client with a Session.
	// Optional aws.Config values can also be provided as variadic arguments
	// to the New function. This option allows you to provide service
	// specific configuration.

	svc := secretsmanager.New(sess)

	//Check if arguments (key, version) have been specified
	secretName, keyName, secretVersion := ExtractSecretPathAndArguments(variableName)

	secretValueInput := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	}

	// If a secret version has been specified
	if len(secretVersion) > 0 {
		secretValueInput.VersionStage = aws.String(secretVersion)
	}

	// Get secret value
	req, resp := svc.GetSecretValueRequest(secretValueInput)

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
	// If key has been specified for retrieval
	if len(keyName) > 0 {
		secretBytes, err = getValueByKey(keyName, secretBytes)
		if err != nil {
			printAndExit(err)
		}
	}

	os.Stdout.Write(secretBytes)
}

func main() {
	if len(os.Args) != 2 {
		os.Stderr.Write([]byte("A variable ID or version flag must be given as the first and only argument!\n"))
		os.Exit(-1)
	}

	// Get the secret and key name from the argument
	singleArgmument := os.Args[1]

	switch singleArgmument {
	case "-v", "--version":
		os.Stdout.Write([]byte(VERSION))
	default:
		RetrieveSecret(singleArgmument)
	}
}

func printAndExit(err error) {
	os.Stderr.Write([]byte(err.Error()))
	os.Exit(1)
}

func getValueByKey(keyName string, secretBytes []byte) (secret []byte, err error) {
	var secrets map[string]interface{}
	var secretValue string

	if err := json.Unmarshal(secretBytes, &secrets); err != nil {
		return nil, err
	}

	secretValue = fmt.Sprint(secrets[keyName])

	return []byte(secretValue), nil
}
