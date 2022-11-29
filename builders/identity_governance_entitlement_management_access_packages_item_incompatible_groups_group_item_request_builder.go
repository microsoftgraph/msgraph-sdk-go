package builders

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// IdentityGovernanceEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\accessPackages\{accessPackage-id}\incompatibleGroups\{group-id}
type IdentityGovernanceEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewIdentityGovernanceEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilderInternal instantiates a new GroupItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}/incompatibleGroups/{group%2Did}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder instantiates a new GroupItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *IdentityGovernanceEntitlementManagementAccessPackagesItemIncompatibleGroupsGroupItemRequestBuilder) Ref()(*IdentityGovernanceEntitlementManagementAccessPackagesItemIncompatibleGroupsItemRefRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackagesItemIncompatibleGroupsItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
