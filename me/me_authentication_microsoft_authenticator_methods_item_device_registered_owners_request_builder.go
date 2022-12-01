package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder provides operations to manage the registeredOwners property of the microsoft.graph.device entity.
type MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilderGetQueryParameters the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
type MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilderGetQueryParameters
}
// AppRoleAssignment casts the previous resource to appRoleAssignment.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder) AppRoleAssignment()(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersAppRoleAssignmentRequestBuilder) {
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilderInternal instantiates a new RegisteredOwnersRequestBuilder and sets the default values.
func NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder) {
    m := &MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder instantiates a new RegisteredOwnersRequestBuilder and sets the default values.
func NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder) Count()(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersCountRequestBuilder) {
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Endpoint casts the previous resource to endpoint.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder) Endpoint()(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersEndpointRequestBuilder) {
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the user that cloud joined the device or registered their personal device. The registered owner is set at the time of registration. Currently, there can be only one owner. Read-only. Nullable. Supports $expand.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDirectoryObjectCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectCollectionResponseable), nil
}
// Ref provides operations to manage the collection of user entities.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder) Ref()(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRefRequestBuilder) {
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder) ServicePrincipal()(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersServicePrincipalRequestBuilder) {
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersRequestBuilder) User()(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserRequestBuilder) {
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
