package ingest

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

type lmv1Token struct {
	AccessID  string
	Signature string
	Epoch     time.Time
}

func (t *lmv1Token) String() string {
	builder := strings.Builder{}
	append := func(s string) {
		if _, err := builder.WriteString(s); err != nil {
			panic(err)
		}
	}
	append("LMv1 ")
	append(t.AccessID)
	append(":")
	append(t.Signature)
	append(":")
	append(strconv.FormatInt(t.Epoch.UnixNano()/1000000, 10))

	return builder.String()
}

func generateLMv1Token(accessId, accessKey string, body []byte) *lmv1Token {

	epochTime := time.Now()
	epoch := strconv.FormatInt(epochTime.UnixNano()/1000000, 10)

	methodUpper := strings.ToUpper(method)

	h := hmac.New(sha256.New, []byte(accessKey))

	writeOrPanic := func(bs []byte) {
		if _, err := h.Write(bs); err != nil {
			panic(err)
		}
	}
	writeOrPanic([]byte(methodUpper))
	writeOrPanic([]byte(epoch))
	writeOrPanic(body)
	writeOrPanic([]byte(resourcePath))

	hash := h.Sum(nil)
	hexString := hex.EncodeToString(hash)
	signature := base64.StdEncoding.EncodeToString([]byte(hexString))
	return &lmv1Token{
		AccessID:  accessId,
		Signature: signature,
		Epoch:     epochTime,
	}
}
