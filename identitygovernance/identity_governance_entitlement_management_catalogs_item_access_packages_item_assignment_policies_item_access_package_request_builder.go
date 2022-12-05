package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilder provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignmentPolicy entity.
type IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilderGetQueryParameters access package containing this policy. Read-only.
type IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilderGetQueryParameters
}
// NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilderInternal instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/assignmentPolicies/{accessPackageAssignmentPolicy%2Did}/accessPackage{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilder instantiates a new AccessPackageRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation access package containing this policy. Read-only.
func (m *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get access package containing this policy. Read-only.
func (m *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemAssignmentPoliciesItemAccessPackageRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable), nil
}
