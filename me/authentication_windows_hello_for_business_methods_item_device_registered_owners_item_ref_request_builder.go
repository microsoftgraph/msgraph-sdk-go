package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder provides operations to manage the collection of user entities.
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteQueryParameters delete ref of navigation property registeredOwners for me
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteQueryParameters struct {
    // Delete Uri
    Id *string `uriparametername:"%40id"`
}
// AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteQueryParameters
}
// NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderInternal instantiates a new RefRequestBuilder and sets the default values.
func NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) {
    m := &AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}/$ref{?%40id*}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder instantiates a new RefRequestBuilder and sets the default values.
func NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete ref of navigation property registeredOwners for me
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToDeleteRequestInformation delete ref of navigation property registeredOwners for me
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
