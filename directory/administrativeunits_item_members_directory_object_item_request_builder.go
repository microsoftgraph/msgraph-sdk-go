package directory

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder builds and executes requests for operations under \directory\administrativeUnits\{administrativeUnit-id}\members\{directoryObject-id}
type AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// NewAdministrativeunitsItemMembersDirectoryObjectItemRequestBuilderInternal instantiates a new AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder and sets the default values.
func NewAdministrativeunitsItemMembersDirectoryObjectItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder) {
    m := &AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/administrativeUnits/{administrativeUnit%2Did}/members/{directoryObject%2Did}", pathParameters),
    }
    return m
}
// NewAdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder instantiates a new AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder and sets the default values.
func NewAdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAdministrativeunitsItemMembersDirectoryObjectItemRequestBuilderInternal(urlParams, requestAdapter)
}
// GraphApplication casts the previous resource to application.
// returns a *AdministrativeunitsItemMembersItemGraphapplicationGraphApplicationRequestBuilder when successful
func (m *AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder) GraphApplication()(*AdministrativeunitsItemMembersItemGraphapplicationGraphApplicationRequestBuilder) {
    return NewAdministrativeunitsItemMembersItemGraphapplicationGraphApplicationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphDevice casts the previous resource to device.
// returns a *AdministrativeunitsItemMembersItemGraphdeviceGraphDeviceRequestBuilder when successful
func (m *AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder) GraphDevice()(*AdministrativeunitsItemMembersItemGraphdeviceGraphDeviceRequestBuilder) {
    return NewAdministrativeunitsItemMembersItemGraphdeviceGraphDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphGroup casts the previous resource to group.
// returns a *AdministrativeunitsItemMembersItemGraphgroupGraphGroupRequestBuilder when successful
func (m *AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder) GraphGroup()(*AdministrativeunitsItemMembersItemGraphgroupGraphGroupRequestBuilder) {
    return NewAdministrativeunitsItemMembersItemGraphgroupGraphGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphOrgContact casts the previous resource to orgContact.
// returns a *AdministrativeunitsItemMembersItemGraphorgcontactGraphOrgContactRequestBuilder when successful
func (m *AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder) GraphOrgContact()(*AdministrativeunitsItemMembersItemGraphorgcontactGraphOrgContactRequestBuilder) {
    return NewAdministrativeunitsItemMembersItemGraphorgcontactGraphOrgContactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphServicePrincipal casts the previous resource to servicePrincipal.
// returns a *AdministrativeunitsItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder when successful
func (m *AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder) GraphServicePrincipal()(*AdministrativeunitsItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilder) {
    return NewAdministrativeunitsItemMembersItemGraphserviceprincipalGraphServicePrincipalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// GraphUser casts the previous resource to user.
// returns a *AdministrativeunitsItemMembersItemGraphuserGraphUserRequestBuilder when successful
func (m *AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder) GraphUser()(*AdministrativeunitsItemMembersItemGraphuserGraphUserRequestBuilder) {
    return NewAdministrativeunitsItemMembersItemGraphuserGraphUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Ref provides operations to manage the collection of directory entities.
// returns a *AdministrativeunitsItemMembersItemRefRequestBuilder when successful
func (m *AdministrativeunitsItemMembersDirectoryObjectItemRequestBuilder) Ref()(*AdministrativeunitsItemMembersItemRefRequestBuilder) {
    return NewAdministrativeunitsItemMembersItemRefRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
