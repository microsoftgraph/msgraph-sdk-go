package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \users\{user-id}\authentication\windowsHelloForBusinessMethods\{windowsHelloForBusinessAuthenticationMethod-id}\device\registeredOwners\{directoryObject-id}
type ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppRoleAssignment casts the previous resource to appRoleAssignment.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) AppRoleAssignment()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemAppRoleAssignmentRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    m := &ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/windowsHelloForBusinessMethods/{windowsHelloForBusinessAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Endpoint casts the previous resource to endpoint.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) Endpoint()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of user entities.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) Ref()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemServicePrincipalRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) User()(*ItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemUserRequestBuilder) {
    return NewItemAuthenticationWindowsHelloForBusinessMethodsItemDeviceRegisteredOwnersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
