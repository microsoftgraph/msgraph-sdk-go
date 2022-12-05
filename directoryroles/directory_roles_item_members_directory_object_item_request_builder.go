package directoryroles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DirectoryRolesItemMembersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \directoryRoles\{directoryRole-id}\members\{directoryObject-id}
type DirectoryRolesItemMembersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Application casts the previous resource to application.
func (m *DirectoryRolesItemMembersDirectoryObjectItemRequestBuilder) Application()(*DirectoryRolesItemMembersItemApplicationRequestBuilder) {
    return NewDirectoryRolesItemMembersItemApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryRolesItemMembersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryRolesItemMembersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRolesItemMembersDirectoryObjectItemRequestBuilder) {
    m := &DirectoryRolesItemMembersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directoryRoles/{directoryRole%2Did}/members/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryRolesItemMembersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryRolesItemMembersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryRolesItemMembersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryRolesItemMembersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Device casts the previous resource to device.
func (m *DirectoryRolesItemMembersDirectoryObjectItemRequestBuilder) Device()(*DirectoryRolesItemMembersItemDeviceRequestBuilder) {
    return NewDirectoryRolesItemMembersItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Group casts the previous resource to group.
func (m *DirectoryRolesItemMembersDirectoryObjectItemRequestBuilder) Group()(*DirectoryRolesItemMembersItemGroupRequestBuilder) {
    return NewDirectoryRolesItemMembersItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *DirectoryRolesItemMembersDirectoryObjectItemRequestBuilder) OrgContact()(*DirectoryRolesItemMembersItemOrgContactRequestBuilder) {
    return NewDirectoryRolesItemMembersItemOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of directoryRole entities.
func (m *DirectoryRolesItemMembersDirectoryObjectItemRequestBuilder) Ref()(*DirectoryRolesItemMembersItemRefRequestBuilder) {
    return NewDirectoryRolesItemMembersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *DirectoryRolesItemMembersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*DirectoryRolesItemMembersItemServicePrincipalRequestBuilder) {
    return NewDirectoryRolesItemMembersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *DirectoryRolesItemMembersDirectoryObjectItemRequestBuilder) User()(*DirectoryRolesItemMembersItemUserRequestBuilder) {
    return NewDirectoryRolesItemMembersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
