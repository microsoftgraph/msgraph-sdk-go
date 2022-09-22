package msgraphsdkgo

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	az "github.com/microsoft/kiota-authentication-azure-go"
)

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
