package builders

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GroupsItemMembersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \groups\{group-id}\members\{directoryObject-id}
type GroupsItemMembersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Application casts the previous resource to application.
func (m *GroupsItemMembersDirectoryObjectItemRequestBuilder) Application()(*GroupsItemMembersItemApplicationRequestBuilder) {
    return NewGroupsItemMembersItemApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupsItemMembersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewGroupsItemMembersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemMembersDirectoryObjectItemRequestBuilder) {
    m := &GroupsItemMembersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/members/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemMembersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewGroupsItemMembersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemMembersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemMembersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Device casts the previous resource to device.
func (m *GroupsItemMembersDirectoryObjectItemRequestBuilder) Device()(*GroupsItemMembersItemDeviceRequestBuilder) {
    return NewGroupsItemMembersItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Group casts the previous resource to group.
func (m *GroupsItemMembersDirectoryObjectItemRequestBuilder) Group()(*GroupsItemMembersItemGroupRequestBuilder) {
    return NewGroupsItemMembersItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *GroupsItemMembersDirectoryObjectItemRequestBuilder) OrgContact()(*GroupsItemMembersItemOrgContactRequestBuilder) {
    return NewGroupsItemMembersItemOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of group entities.
func (m *GroupsItemMembersDirectoryObjectItemRequestBuilder) Ref()(*GroupsItemMembersItemRefRequestBuilder) {
    return NewGroupsItemMembersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *GroupsItemMembersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*GroupsItemMembersItemServicePrincipalRequestBuilder) {
    return NewGroupsItemMembersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *GroupsItemMembersDirectoryObjectItemRequestBuilder) User()(*GroupsItemMembersItemUserRequestBuilder) {
    return NewGroupsItemMembersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
