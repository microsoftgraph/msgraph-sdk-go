package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilder provides operations to count the resources in the collection.
type UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilder) {
    m := &UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredUsers/microsoft.graph.user/$count";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the number of the resource
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "text/plain"
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get get the number of the resource
func (m *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersUserCountRequestBuilderGetRequestConfiguration)(*int32, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendPrimitiveAsync(ctx, requestInfo, "int32", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(*int32), nil
}
