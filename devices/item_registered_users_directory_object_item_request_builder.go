package devices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRegisteredUsersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \devices\{device-id}\registeredUsers\{directoryObject-id}
type ItemRegisteredUsersDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRegisteredUsersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRegisteredUsersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredUsersDirectoryObjectItemRequestBuilder) {
    m := &ItemRegisteredUsersDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/devices/{device%2Did}/registeredUsers/{directoryObject%2Did}", pathParameters),
    }
    return m
}
// NewItemRegisteredUsersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRegisteredUsersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredUsersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRegisteredUsersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// GraphAppRoleAssignment casts the previous resource to appRoleAssignment.
func (m *ItemRegisteredUsersDirectoryObjectItemRequestBuilder) GraphAppRoleAssignment()(*ItemRegisteredUsersItemGraphAppRoleAssignmentRequestBuilder) {
    return NewItemRegisteredUsersItemGraphAppRoleAssignmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphEndpoint casts the previous resource to endpoint.
func (m *ItemRegisteredUsersDirectoryObjectItemRequestBuilder) GraphEndpoint()(*ItemRegisteredUsersItemGraphEndpointRequestBuilder) {
    return NewItemRegisteredUsersItemGraphEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemRegisteredUsersDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*ItemRegisteredUsersItemGraphServicePrincipalRequestBuilder) {
    return NewItemRegisteredUsersItemGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
func (m *ItemRegisteredUsersDirectoryObjectItemRequestBuilder) GraphUser()(*ItemRegisteredUsersItemGraphUserRequestBuilder) {
    return NewItemRegisteredUsersItemGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of device entities.
func (m *ItemRegisteredUsersDirectoryObjectItemRequestBuilder) Ref()(*ItemRegisteredUsersItemRefRequestBuilder) {
    return NewItemRegisteredUsersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
