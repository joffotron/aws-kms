package main

import (
	"os"

	"github.com/joffotron/aws-kms/kms"

	"fmt"

	"github.com/voxelbrain/goptions"
)

func main() {
	os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
	opts := struct {
		Help goptions.Help `goptions:"-h, --help, description='Show this help'"`

		goptions.Verbs
		Encrypt struct {
			KeyId string `goptions:"-k, --key-id, description='Key ID to encrypt with', obligatory"`
			Input string `goptions:"-i, --input, description='Input data', obligatory"`
		} `goptions:"encrypt"`

		Decrypt struct {
			Input string `goptions:"-i, --input, description='Input data', obligatory"`
		} `goptions:"decrypt"`
	}{}

	goptions.ParseAndFail(&opts)
	if len(os.Args) < 2 {
		goptions.PrintHelp()
	}

	if opts.Encrypt.Input != "" {
		fmt.Println(kms.Encrypt(opts.Encrypt.Input, opts.Encrypt.KeyId))
	}

	if opts.Decrypt.Input != "" {
		fmt.Println(kms.Decrypt(opts.Decrypt.Input))
	}
}
