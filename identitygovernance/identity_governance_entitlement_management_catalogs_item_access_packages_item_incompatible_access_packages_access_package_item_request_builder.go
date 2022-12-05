package identitygovernance

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder builds and executes requests for operations under \identityGovernance\entitlementManagement\catalogs\{accessPackageCatalog-id}\accessPackages\{accessPackage-id}\incompatibleAccessPackages\{accessPackage-id1}
type IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/incompatibleAccessPackages/{accessPackage%2Did1}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder instantiates a new AccessPackageItemRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesAccessPackageItemRequestBuilder) Ref()(*IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesItemRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
