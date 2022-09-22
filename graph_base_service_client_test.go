package msgraphsdkgo

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewSimpleCredentials(token string) *SimpleCredentials {
	return &SimpleCredentials{
		TokenValue: token,
	}
}

type SimpleCredentials struct {
	TokenValue string
}

func (m *SimpleCredentials) GetToken(context.Context, policy.TokenRequestOptions) (azcore.AccessToken, error) {
	return azcore.AccessToken{
		Token: m.TokenValue,
	}, nil
}

func TestNewGraphServiceClientWithCredentialsInitializesAClient(t *testing.T) {
	cred := NewSimpleCredentials("asdasd")
	client, _ := NewGraphServiceClientWithCredentials(cred, []string{"Files.Read"})

	assert.NotNil(t, client)
}
