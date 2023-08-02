package ingest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	accessId    = "someAccessId"
	accessKey   = "someAccessKey"
	bearerToken = "token"
	comapny     = "company"
	bodyByte    = []byte("Body string")
)

func TestAllAuthModeSpecified(t *testing.T) {

	lmIngest, err := NewLogIngester(comapny, accessId, accessKey, bearerToken, "ls", "ver")
	assert.Nil(t, err)
	authString := lmIngest.GenerateAuthString(bodyByte)
	assert.True(t, strings.Contains(authString, fmt.Sprintf("LMv1 %s", accessId)))

}

func TestNoAuthSpecified(t *testing.T) {

	_, err := NewLogIngester(comapny, "", "", "", "ls", "ver")
	assert.NotNil(t, err)

}

func TestOnlyPartialLmv1specified(t *testing.T) {

	_, err := NewLogIngester(comapny, accessId, "", "", "ls", "ver")
	assert.NotNil(t, err)

}
func TestOnlyBearerSpecified(t *testing.T) {

	lmIngest, err := NewLogIngester(comapny, "", "", bearerToken, "ls", "ver")
	assert.Nil(t, err)
	authString := lmIngest.GenerateAuthString(bodyByte)
	assert.Equal(t, authString, fmt.Sprintf("Bearer %s", bearerToken))

}

func TestBearerWithPartialLmv1Specified(t *testing.T) {

	lmIngest, err := NewLogIngester(comapny, "", accessKey, bearerToken, "ls", "ver")
	assert.Nil(t, err)
	authString := lmIngest.GenerateAuthString(bodyByte)
	assert.Equal(t, authString, fmt.Sprintf("Bearer %s", bearerToken))

}
