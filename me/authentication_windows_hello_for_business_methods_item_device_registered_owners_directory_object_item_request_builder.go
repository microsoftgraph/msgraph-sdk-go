package me

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \me\authentication\windowsHelloForBusinessMethods\{windowsHelloForBusinessAuthenticationMethod-id}\device\registeredOwners\{directoryObject-id}
type AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppRoleAssignment casts the previous resource to appRoleAssignment.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) AppRoleAssignment()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemAppRoleAssignmentRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    m := &AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Endpoint casts the previous resource to endpoint.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) Endpoint()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of user entities.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) Ref()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemServicePrincipalRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) User()(*AuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemUserRequestBuilder) {
    return NewAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
