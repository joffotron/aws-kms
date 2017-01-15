package main

import (
	"fmt"
	"os"

	"github.com/joffotron/aws-kms/kms"
)

func main() {
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")

	if len(os.Args) != 3 {
		usageAndExit()
	}

	switch os.Args[1] {
	case "encrypt":
		fmt.Println(kms.Encrypt(os.Args[2]))
	case "decrypt":
		fmt.Println(kms.Decrypt(os.Args[2]))
	default:
		usageAndExit()
	}

}

func usageAndExit() {
	fmt.Println("USAGE: ")
	fmt.Println("  aws-kms encrypt DATA")
	fmt.Println("  aws-kms decrypt BASE64_ENCODED_DATA")
	os.Exit(1)
}
