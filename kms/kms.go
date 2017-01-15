package kms

//noinspection SpellCheckingInspection
import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws/session"

	awskms "github.com/aws/aws-sdk-go/service/kms"
)

func Encrypt(data, keyId string) string {
	kmsApi := kmsApi()
	params := &awskms.EncryptInput{
		Plaintext: []byte(data),
		KeyId:     &keyId,
	}

	encrypted, err := kmsApi.Encrypt(params)
	if err != nil {
		fail(err)
	}

	encoded := base64.StdEncoding.EncodeToString(encrypted.CiphertextBlob)

	return encoded
}

func Decrypt(data string) string {
	cipherBytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		fail(err)
	}

	kmsApi := kmsApi()
	params := &awskms.DecryptInput{CiphertextBlob: cipherBytes}
	decrypted, decryptErr := kmsApi.Decrypt(params)
	if decryptErr != nil {
		fail(decryptErr)
	}

	return string(decrypted.Plaintext)
}

func fail(err error) {
	panic(err)
}

func kmsApi() *awskms.KMS {
	awsSession := session.Must(session.NewSession())
	return awskms.New(awsSession)
}
