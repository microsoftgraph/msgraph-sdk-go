package directoryroles

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// ItemMembersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \directoryRoles\{directoryRole-id}\members\{directoryObject-id}
type ItemMembersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Application casts the previous resource to application.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) Application()(*ItemMembersItemApplicationRequestBuilder) {
    return NewItemMembersItemApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewItemMembersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMembersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersDirectoryObjectItemRequestBuilder) {
    m := &ItemMembersDirectoryObjectItemRequestBuilder{
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
// NewItemMembersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewItemMembersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemMembersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemMembersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Device casts the previous resource to device.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) Device()(*ItemMembersItemDeviceRequestBuilder) {
    return NewItemMembersItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Group casts the previous resource to group.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) Group()(*ItemMembersItemGroupRequestBuilder) {
    return NewItemMembersItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) OrgContact()(*ItemMembersItemOrgContactRequestBuilder) {
    return NewItemMembersItemOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of directoryRole entities.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) Ref()(*ItemMembersItemRefRequestBuilder) {
    return NewItemMembersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*ItemMembersItemServicePrincipalRequestBuilder) {
    return NewItemMembersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *ItemMembersDirectoryObjectItemRequestBuilder) User()(*ItemMembersItemUserRequestBuilder) {
    return NewItemMembersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
