package applications

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ApplicationsItemOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \applications\{application-id}\owners\{directoryObject-id}
type ApplicationsItemOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// AppRoleAssignment casts the previous resource to appRoleAssignment.
func (m *ApplicationsItemOwnersDirectoryObjectItemRequestBuilder) AppRoleAssignment()(*ApplicationsItemOwnersItemAppRoleAssignmentRequestBuilder) {
    return NewApplicationsItemOwnersItemAppRoleAssignmentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewApplicationsItemOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewApplicationsItemOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsItemOwnersDirectoryObjectItemRequestBuilder) {
    m := &ApplicationsItemOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/applications/{application%2Did}/owners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewApplicationsItemOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewApplicationsItemOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ApplicationsItemOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewApplicationsItemOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Endpoint casts the previous resource to endpoint.
func (m *ApplicationsItemOwnersDirectoryObjectItemRequestBuilder) Endpoint()(*ApplicationsItemOwnersItemEndpointRequestBuilder) {
    return NewApplicationsItemOwnersItemEndpointRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of application entities.
func (m *ApplicationsItemOwnersDirectoryObjectItemRequestBuilder) Ref()(*ApplicationsItemOwnersItemRefRequestBuilder) {
    return NewApplicationsItemOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *ApplicationsItemOwnersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*ApplicationsItemOwnersItemServicePrincipalRequestBuilder) {
    return NewApplicationsItemOwnersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *ApplicationsItemOwnersDirectoryObjectItemRequestBuilder) User()(*ApplicationsItemOwnersItemUserRequestBuilder) {
    return NewApplicationsItemOwnersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
