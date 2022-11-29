package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilder provides operations to count the resources in the collection.
type MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilderInternal instantiates a new CountRequestBuilder and sets the default values.
func NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilder) {
    m := &MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/transitiveMemberOf/microsoft.graph.orgContact/$count";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilder instantiates a new CountRequestBuilder and sets the default values.
func NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get the number of the resource
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfOrgContactCountRequestBuilderGetRequestConfiguration)(*int32, error) {
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
