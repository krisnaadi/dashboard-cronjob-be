package signature

import (
	"encoding/base64"
	"encoding/json"
	"errors"

	"github.com/tech-djoin/wallet-djoin-service/internal/usecase/signature"
)

func GenerateSignatureForPpob(timestamp string, requestBody map[string]interface{}, companyCode string, companyStaticKey string) (string, error) {
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return "", errors.New("Generate Signature Failed.")
	}

	message := timestamp + "|" + string(requestBodyJSON)
	code := companyCode

	key := companyStaticKey + code
	hash := signature.HmacSHA256(message, key)

	signature := base64.StdEncoding.EncodeToString(hash)

	return signature, nil
}
