package users

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \users\{user-id}\authentication\microsoftAuthenticatorMethods\{microsoftAuthenticatorAuthenticationMethod-id}\device\registeredOwners\{directoryObject-id}
type ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppRoleAssignment casts the previous resource to appRoleAssignment.
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) AppRoleAssignment()(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemAppRoleAssignmentRequestBuilder) {
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    m := &ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}/device/registeredOwners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Endpoint casts the previous resource to endpoint.
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) Endpoint()(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilder) {
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of user entities.
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) Ref()(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemRefRequestBuilder) {
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemServicePrincipalRequestBuilder) {
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersDirectoryObjectItemRequestBuilder) User()(*ItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemUserRequestBuilder) {
    return NewItemAuthenticationMicrosoftAuthenticatorMethodsItemDeviceRegisteredOwnersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
