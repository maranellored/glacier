package main

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/rdwilliamson/aws"
	"github.com/rdwilliamson/aws/glacier"
	"gopkg.in/yaml.v2"
)

func main() {

	fmt.Println("Enter absolute path of file to pick up AWS credentials from:")
	absPath := ""
	fmt.Scanln(&absPath)

	credsMap, err := getCredentials(absPath)
	if err != nil {
		panic(err)
	}
	accessKey := credsMap["access-key"]
	secretKey := credsMap["secret-key"]

	connection := glacier.NewConnection(secretKey, accessKey, aws.USEast1)

	err = connection.CreateVault("Example")
	if err != nil {
		panic(err)
	}

	_, err = connection.DescribeVault("Example")
	if err != nil {
		panic(err)
	}

	err = connection.DeleteVault("Example")
	if err != nil {
		panic(err)
	}

}

func getCredentials(filePath string) (map[string]string, error) {
	buffer, err := ioutil.ReadFile(filePath)
	if err != nil {
		errors.New("Cannot open file for reading")
	}

	credsMap := make(map[string]string)
	err = yaml.Unmarshal(buffer, credsMap)
	if err != nil {
		return nil, errors.New("Cannot unmarshal data from file")
	}

	return credsMap, err
}
