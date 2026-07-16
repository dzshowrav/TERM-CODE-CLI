package llm

import (
	"fmt"
	"os"
	"strings"
)

type AuthManager struct{}

func NewAuthManager() *AuthManager {
	return &AuthManager{}
}

func (m *AuthManager) ResolveAPIKey(providerName, storedKey string, decryptFn func(string) (string, error)) (string, error) {
	envKey := fmt.Sprintf("TC_PROVIDER_%s_API_KEY", strings.ToUpper(strings.ReplaceAll(providerName, " ", "_")))
	if envKey := os.Getenv(envKey); envKey != "" {
		return envKey, nil
	}

	if storedKey != "" {
		return decryptFn(storedKey)
	}

	return "", nil
}

func (m *AuthManager) MaskKey(key string) string {
	if len(key) <= 8 {
		return "****"
	}
	return key[:4] + "****" + key[len(key)-4:]
}
