package serviceprincipals

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \servicePrincipals\{servicePrincipal-id}\owners\{directoryObject-id}
type ServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppRoleAssignment casts the previous resource to appRoleAssignment.
func (m *ServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilder) AppRoleAssignment()(*ServicePrincipalsItemOwnersItemAppRoleAssignmentRequestBuilder) {
    return NewServicePrincipalsItemOwnersItemAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilder) {
    m := &ServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/owners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Endpoint casts the previous resource to endpoint.
func (m *ServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilder) Endpoint()(*ServicePrincipalsItemOwnersItemEndpointRequestBuilder) {
    return NewServicePrincipalsItemOwnersItemEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of servicePrincipal entities.
func (m *ServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilder) Ref()(*ServicePrincipalsItemOwnersItemRefRequestBuilder) {
    return NewServicePrincipalsItemOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *ServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*ServicePrincipalsItemOwnersItemServicePrincipalRequestBuilder) {
    return NewServicePrincipalsItemOwnersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *ServicePrincipalsItemOwnersDirectoryObjectItemRequestBuilder) User()(*ServicePrincipalsItemOwnersItemUserRequestBuilder) {
    return NewServicePrincipalsItemOwnersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
