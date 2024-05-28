package devices

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemRegisteredownersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \devices\{device-id}\registeredOwners\{directoryObject-id}
type ItemRegisteredownersDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemRegisteredownersDirectoryObjectItemRequestBuilderInternal instantiates a new ItemRegisteredownersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRegisteredownersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredownersDirectoryObjectItemRequestBuilder) {
    m := &ItemRegisteredownersDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/devices/{device%2Did}/registeredOwners/{directoryObject%2Did}", pathParameters),
    }
    return m
}
// NewItemRegisteredownersDirectoryObjectItemRequestBuilder instantiates a new ItemRegisteredownersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemRegisteredownersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRegisteredownersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRegisteredownersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// GraphAppRoleAssignment casts the previous resource to appRoleAssignment.
// returns a *ItemRegisteredownersItemGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder when successful
func (m *ItemRegisteredownersDirectoryObjectItemRequestBuilder) GraphAppRoleAssignment()(*ItemRegisteredownersItemGraphapproleassignmentGraphAppRoleAssignmentRequestBuilder) {
    return NewItemRegisteredownersItemGraphapproleassignmentGraphAppRoleAssignmentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphEndpoint casts the previous resource to endpoint.
// returns a *ItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder when successful
func (m *ItemRegisteredownersDirectoryObjectItemRequestBuilder) GraphEndpoint()(*ItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilder) {
    return NewItemRegisteredownersItemGraphendpointGraphEndpointRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemRegisteredownersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *ItemRegisteredownersDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*ItemRegisteredownersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewItemRegisteredownersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemRegisteredownersItemGraphuserGraphUserRequestBuilder when successful
func (m *ItemRegisteredownersDirectoryObjectItemRequestBuilder) GraphUser()(*ItemRegisteredownersItemGraphuserGraphUserRequestBuilder) {
    return NewItemRegisteredownersItemGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of device entities.
// returns a *ItemRegisteredownersItemRefRequestBuilder when successful
func (m *ItemRegisteredownersDirectoryObjectItemRequestBuilder) Ref()(*ItemRegisteredownersItemRefRequestBuilder) {
    return NewItemRegisteredownersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
