package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder provides operations to manage the registeredUsers property of the microsoft.graph.device entity.
type MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilderGetQueryParameters collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
type MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilderGetQueryParameters
}
// AppRoleAssignment casts the previous resource to appRoleAssignment.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder) AppRoleAssignment()(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersItemAppRoleAssignmentRequestBuilder) {
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersItemAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder) {
    m := &MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredUsers/{directoryObject%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder) Endpoint()(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersItemEndpointRequestBuilder) {
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersItemEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get collection of registered users of the device. For cloud joined devices and registered personal devices, registered users are set to the same value as registered owners at the time of registration. Read-only. Nullable. Supports $expand.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder) Get(ctx context.Context, requestConfiguration *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DirectoryObjectable, error) {
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
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersItemServicePrincipalRequestBuilder) {
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersDirectoryObjectItemRequestBuilder) User()(*MeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersItemUserRequestBuilder) {
    return NewMeAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredUsersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
