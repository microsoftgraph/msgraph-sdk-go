package msgraphsdkgo

import (
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	az "github.com/microsoft/kiota-authentication-azure-go"
)

type GraphServiceClient struct {
	*GraphBaseServiceClient
}

func NewGraphServiceClient(adapter abstractions.RequestAdapter) *GraphServiceClient {
	client := NewBaseGraphServiceClient(adapter)
	return &GraphServiceClient{
		client,
	}
}

// NewGraphServiceClientWithCredentials instantiates a new GraphServiceClient with provided credentials and scopes
func NewGraphServiceClientWithCredentials(credential azcore.TokenCredential, scopes []string) (*GraphServiceClient, error) {
	auth, err := az.NewAzureIdentityAuthenticationProviderWithScopes(credential, scopes)
	if err != nil {
		return nil, err
	}

	adapter, err := NewGraphRequestAdapter(auth)
	if err != nil {
		return nil, err
	}

	client := NewGraphServiceClient(adapter)
	return client, nil
}

// NewGraphServiceClientWithCredentialsAndHosts instantiates a new GraphServiceClient with provided credentials , scopes and validhosts
func NewGraphServiceClientWithCredentialsAndHosts(credential azcore.TokenCredential, scopes []string, validhosts []string) (*GraphServiceClient, error) {
	auth, err := az.NewAzureIdentityAuthenticationProviderWithScopesAndValidHosts(credential, scopes, validhosts)
	if err != nil {
		return nil, err
	}

	adapter, err := NewGraphRequestAdapter(auth)
	if err != nil {
		return nil, err
	}

	client := NewGraphServiceClient(adapter)
	return client, nil
}

// GetAdapter returns the client current adapter, Method should only be called when the user is certain an adapter has been provided
func (m *GraphBaseServiceClient) GetAdapter() abstractions.RequestAdapter {
	if m.requestAdapter != nil {
		panic(errors.New("request adapter has not been initialized"))
	}
	return m.requestAdapter
}
