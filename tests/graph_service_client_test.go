package tests

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/microsoftgraph/msgraph-sdk-go"
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
	client, _ := msgraphsdkgo.NewGraphServiceClientWithCredentials(cred, []string{"Files.Read"})

	assert.NotNil(t, client)
}

func TestNewGraphServiceClientWithCredentialsAndHostsInitializesAClient(t *testing.T) {
	cred := NewSimpleCredentials("asdasd")
	client, _ := msgraphsdkgo.NewGraphServiceClientWithCredentialsAndHosts(cred, []string{"Files.Read"}, []string{"host1", "host2"})

	assert.NotNil(t, client)
}

func TestClientAccessToGeneratedMethods(t *testing.T) {
	cred := NewSimpleCredentials("asdasd")
	client, _ := msgraphsdkgo.NewGraphServiceClientWithCredentials(cred, []string{"Files.Read"})

	assert.NotNil(t, client.Groups())
}
