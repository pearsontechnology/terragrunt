package kmsgrunt

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/aws/aws-sdk-go/service/kms"
)

// Decrypt decrypts a passed argument
func Decrypt(svc *kms.KMS, encryptedVal string, encryptedValkey string) string {
	blob, _ := base64.StdEncoding.DecodeString(encryptedVal[3:])

	// Decrypt the data
	result, err := svc.Decrypt(&kms.DecryptInput{CiphertextBlob: blob})

	if err != nil {
		fmt.Println("Error decrypting data: ", err)
		fmt.Println("Using undecrypted vars")
		return ""
	}
	viper.SetConfigType("toml")
	viper.SetConfigName(".kmsgrunt")
	viper.AddConfigPath("$HOME/") // call multiple times to add many search paths
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	viper.AddConfigPath("..")
	viper.AddConfigPath("../../")
	viperr := viper.ReadInConfig() // Find and read the config file
	if viperr != nil {
		fmt.Println("kmsgrunt config read failed - undecrypted variables will be used")
		return ""
	}
	tmpSecretsPath := viper.GetString("tmpSecretsPath")
	fmt.Println("Secrets location: ", tmpSecretsPath)
	f, err := os.OpenFile(tmpSecretsPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println("Open tmpSecretsPath failed - undecrypted variables will be used", err)
		return ""
	}
	defer f.Close()
	if err == nil {
		if _, err = f.WriteString(string(encryptedValkey) + " = \"" + string(result.Plaintext) + "\"\n"); err != nil {
			fmt.Println("Error writing decrypted variable - using undecrypted value")
			return ""
		}
	}
	blobString := string(result.Plaintext)

	return (blobString)
}
