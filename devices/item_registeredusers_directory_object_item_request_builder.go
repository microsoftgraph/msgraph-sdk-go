package devices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRegisteredusersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \devices\{device-id}\registeredUsers\{directoryObject-id}
type ItemRegisteredusersDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRegisteredusersDirectoryObjectItemRequestBuilderInternal instantiates a new ItemRegisteredusersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRegisteredusersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredusersDirectoryObjectItemRequestBuilder) {
    m := &ItemRegisteredusersDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/devices/{device%2Did}/registeredUsers/{directoryObject%2Did}", pathParameters),
    }
    return m
}
// NewItemRegisteredusersDirectoryObjectItemRequestBuilder instantiates a new ItemRegisteredusersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRegisteredusersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredusersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRegisteredusersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// GraphAppRoleAssignment casts the previous resource to appRoleAssignment.
// returns a *ItemRegisteredusersItemGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder when successful
func (m *ItemRegisteredusersDirectoryObjectItemRequestBuilder) GraphAppRoleAssignment()(*ItemRegisteredusersItemGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder) {
    return NewItemRegisteredusersItemGraphapproleassignmentGraphAppRoleAssignmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphEndpoint casts the previous resource to endpoint.
// returns a *ItemRegisteredusersItemGraphendpointGraphEndpointRequestBuilder when successful
func (m *ItemRegisteredusersDirectoryObjectItemRequestBuilder) GraphEndpoint()(*ItemRegisteredusersItemGraphendpointGraphEndpointRequestBuilder) {
    return NewItemRegisteredusersItemGraphendpointGraphEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemRegisteredusersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *ItemRegisteredusersDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*ItemRegisteredusersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewItemRegisteredusersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemRegisteredusersItemGraphuserGraphUserRequestBuilder when successful
func (m *ItemRegisteredusersDirectoryObjectItemRequestBuilder) GraphUser()(*ItemRegisteredusersItemGraphuserGraphUserRequestBuilder) {
    return NewItemRegisteredusersItemGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of device entities.
// returns a *ItemRegisteredusersItemRefRequestBuilder when successful
func (m *ItemRegisteredusersDirectoryObjectItemRequestBuilder) Ref()(*ItemRegisteredusersItemRefRequestBuilder) {
    return NewItemRegisteredusersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
