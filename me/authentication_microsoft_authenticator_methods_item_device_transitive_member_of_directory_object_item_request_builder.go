package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder provides operations to manage the transitiveMemberOf property of the microsoft.graph.device entity.
type AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
type AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetQueryParameters
}
// Application casts the previous resource to application.
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) Application()(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemApplicationRequestBuilder) {
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    m := &AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/transitiveMemberOf/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Device casts the previous resource to device.
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) Device()(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemDeviceRequestBuilder) {
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get groups and administrative units that the device is a member of. This operation is transitive. Supports $expand.
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable), nil
}
// Group casts the previous resource to group.
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) Group()(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemGroupRequestBuilder) {
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) OrgContact()(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemOrgContactRequestBuilder) {
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) ServicePrincipal()(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemServicePrincipalRequestBuilder) {
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfDirectoryObjectItemRequestBuilder) User()(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemUserRequestBuilder) {
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceTransitiveMemberOfItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
