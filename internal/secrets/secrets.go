package secrets

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

var secretMap = make(map[string]string)

// LoadSecret loads a secret by name (or ARN) and caches it under the given key.
func LoadSecret(key, secretName string) error {
	ctx := context.Background()
	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("unable to load AWS SDK config: %w", err)
	}

	client := secretsmanager.NewFromConfig(awsCfg)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: &secretName,
	}

	result, err := client.GetSecretValue(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to retrieve secret %s: %w", secretName, err)
	}

	if result.SecretString == nil {
		return fmt.Errorf("secret string is nil for secret %s", secretName)
	}

	// Cache the secret using the provided key.
	secretMap[key] = *result.SecretString
	log.Printf("✅ Secret loaded and cached under key: %s", key)
	return nil
}

// LoadSecrets loads multiple secrets at once.
// secretsToLoad is a map where the key is your identifier and the value is the secret name/ARN in AWS.
func LoadSecrets(secretsToLoad map[string]string) error {
	for key, secretName := range secretsToLoad {
		if err := LoadSecret(key, secretName); err != nil {
			return err
		}
	}
	return nil
}

// GetSecret returns the cached secret for a given key.
func GetSecret(key string) string {
	return secretMap[key]
}
