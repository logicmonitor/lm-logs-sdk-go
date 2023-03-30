package apitoken

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var accessId = "someAccessId"
var accessKey = "someAccessKey"
var bodyByte = []byte("Body string")

func TestGetAuthToken(t *testing.T) {
	timeNow := time.Now().UnixMilli()

	lmv1token := GenerateLMv1Token(accessId, accessKey, bodyByte)
	assert.NotNil(t, lmv1token)
	assert.Equal(t, accessId, lmv1token.AccessID)
	assert.Equal(t, timeNow, lmv1token.Epoch.UnixMilli())
}

func TestString(t *testing.T) {

	lmv1token := GenerateLMv1Token(accessId, accessKey, bodyByte)
	tokenString := lmv1token.String()

	assert.True(t, strings.Contains(tokenString, accessId))
	assert.True(t, strings.Contains(tokenString, "LMv1"))

}
