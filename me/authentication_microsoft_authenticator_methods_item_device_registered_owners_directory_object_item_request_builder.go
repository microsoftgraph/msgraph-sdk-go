package me

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \me\authentication\microsoftAuthenticatorMethods\{microsoftAuthenticatorAuthenticationMethod-id}\device\registeredOwners\{directoryObject-id}
type AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppRoleAssignment casts the previous resource to appRoleAssignment.
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) AppRoleAssignment()(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemAppRoleAssignmentRequestBuilder) {
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    m := &AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Endpoint casts the previous resource to endpoint.
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) Endpoint()(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder) {
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of user entities.
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) Ref()(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) {
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemServicePrincipalRequestBuilder) {
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) User()(*AuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemUserRequestBuilder) {
    return NewAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
