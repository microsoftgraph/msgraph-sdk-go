package builders

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// GroupsItemOwnersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \groups\{group-id}\owners\{directoryObject-id}
type GroupsItemOwnersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Application casts the previous resource to application.
func (m *GroupsItemOwnersDirectoryObjectItemRequestBuilder) Application()(*GroupsItemOwnersItemApplicationRequestBuilder) {
    return NewGroupsItemOwnersItemApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewGroupsItemOwnersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewGroupsItemOwnersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemOwnersDirectoryObjectItemRequestBuilder) {
    m := &GroupsItemOwnersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group%2Did}/owners/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewGroupsItemOwnersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewGroupsItemOwnersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*GroupsItemOwnersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewGroupsItemOwnersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Device casts the previous resource to device.
func (m *GroupsItemOwnersDirectoryObjectItemRequestBuilder) Device()(*GroupsItemOwnersItemDeviceRequestBuilder) {
    return NewGroupsItemOwnersItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Group casts the previous resource to group.
func (m *GroupsItemOwnersDirectoryObjectItemRequestBuilder) Group()(*GroupsItemOwnersItemGroupRequestBuilder) {
    return NewGroupsItemOwnersItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *GroupsItemOwnersDirectoryObjectItemRequestBuilder) OrgContact()(*GroupsItemOwnersItemOrgContactRequestBuilder) {
    return NewGroupsItemOwnersItemOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of group entities.
func (m *GroupsItemOwnersDirectoryObjectItemRequestBuilder) Ref()(*GroupsItemOwnersItemRefRequestBuilder) {
    return NewGroupsItemOwnersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *GroupsItemOwnersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*GroupsItemOwnersItemServicePrincipalRequestBuilder) {
    return NewGroupsItemOwnersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *GroupsItemOwnersDirectoryObjectItemRequestBuilder) User()(*GroupsItemOwnersItemUserRequestBuilder) {
    return NewGroupsItemOwnersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
