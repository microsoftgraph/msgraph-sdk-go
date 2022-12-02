package directory

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// DirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \directory\administrativeUnits\{administrativeUnit-id}\members\{directoryObject-id}
type DirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// Application casts the previous resource to application.
func (m *DirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder) Application()(*DirectoryAdministrativeUnitsItemMembersItemApplicationRequestBuilder) {
    return NewDirectoryAdministrativeUnitsItemMembersItemApplicationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilderInternal instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder) {
    m := &DirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/directory/administrativeUnits/{administrativeUnit%2Did}/members/{directoryObject%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder instantiates a new DirectoryObjectItemRequestBuilder and sets the default values.
func NewDirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Device casts the previous resource to device.
func (m *DirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder) Device()(*DirectoryAdministrativeUnitsItemMembersItemDeviceRequestBuilder) {
    return NewDirectoryAdministrativeUnitsItemMembersItemDeviceRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Group casts the previous resource to group.
func (m *DirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder) Group()(*DirectoryAdministrativeUnitsItemMembersItemGroupRequestBuilder) {
    return NewDirectoryAdministrativeUnitsItemMembersItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OrgContact casts the previous resource to orgContact.
func (m *DirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder) OrgContact()(*DirectoryAdministrativeUnitsItemMembersItemOrgContactRequestBuilder) {
    return NewDirectoryAdministrativeUnitsItemMembersItemOrgContactRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Ref provides operations to manage the collection of directory entities.
func (m *DirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder) Ref()(*DirectoryAdministrativeUnitsItemMembersItemRefRequestBuilder) {
    return NewDirectoryAdministrativeUnitsItemMembersItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ServicePrincipal casts the previous resource to servicePrincipal.
func (m *DirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder) ServicePrincipal()(*DirectoryAdministrativeUnitsItemMembersItemServicePrincipalRequestBuilder) {
    return NewDirectoryAdministrativeUnitsItemMembersItemServicePrincipalRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// User casts the previous resource to user.
func (m *DirectoryAdministrativeUnitsItemMembersDirectoryObjectItemRequestBuilder) User()(*DirectoryAdministrativeUnitsItemMembersItemUserRequestBuilder) {
    return NewDirectoryAdministrativeUnitsItemMembersItemUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
