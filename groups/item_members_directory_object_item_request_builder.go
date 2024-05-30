package groups

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMembersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \groups\{group-id}\members\{directoryObject-id}
type ItemMembersDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewItemMembersDirectoryObjectItemRequestBuilderInternal instantiates a new ItemMembersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMembersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersDirectoryObjectItemRequestBuilder) {
    m := &ItemMembersDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/members/{directoryObject%2Did}", pathParameters),
    }
    return m
}
// NewItemMembersDirectoryObjectItemRequestBuilder instantiates a new ItemMembersDirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMembersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// GraphApplication casts the previous resource to application.
// returns a *ItemMembersItemGraphapplicationGraphApplicationRequestBuilder when successful
func (m *ItemMembersDirectoryObjectItemRequestBuilder) GraphApplication()(*ItemMembersItemGraphapplicationGraphApplicationRequestBuilder) {
    return NewItemMembersItemGraphapplicationGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
// returns a *ItemMembersItemGraphdeviceGraphDeviceRequestBuilder when successful
func (m *ItemMembersDirectoryObjectItemRequestBuilder) GraphDevice()(*ItemMembersItemGraphdeviceGraphDeviceRequestBuilder) {
    return NewItemMembersItemGraphdeviceGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *ItemMembersItemGraphgroupGraphGroupRequestBuilder when successful
func (m *ItemMembersDirectoryObjectItemRequestBuilder) GraphGroup()(*ItemMembersItemGraphgroupGraphGroupRequestBuilder) {
    return NewItemMembersItemGraphgroupGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphOrgContact casts the previous resource to orgContact.
// returns a *ItemMembersItemGraphorgcontactGraphOrgContactRequestBuilder when successful
func (m *ItemMembersDirectoryObjectItemRequestBuilder) GraphOrgContact()(*ItemMembersItemGraphorgcontactGraphOrgContactRequestBuilder) {
    return NewItemMembersItemGraphorgcontactGraphOrgContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *ItemMembersDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*ItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *ItemMembersItemGraphuserGraphUserRequestBuilder when successful
func (m *ItemMembersDirectoryObjectItemRequestBuilder) GraphUser()(*ItemMembersItemGraphuserGraphUserRequestBuilder) {
    return NewItemMembersItemGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of group entities.
// returns a *ItemMembersItemRefRequestBuilder when successful
func (m *ItemMembersDirectoryObjectItemRequestBuilder) Ref()(*ItemMembersItemRefRequestBuilder) {
    return NewItemMembersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
