package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder provides operations to manage the incompatibleAccessPackages property of the microsoft.graph.accessPackage entity.
type IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetQueryParameters retrieve a list of the accessPackage objects that have been marked as incompatible on an accessPackage.  
type IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetQueryParameters
}
// NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderInternal instantiates a new IncompatibleAccessPackagesRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement/catalogs/{accessPackageCatalog%2Did}/accessPackages/{accessPackage%2Did}/incompatibleAccessPackages{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder instantiates a new IncompatibleAccessPackagesRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) Count()(*IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesCountRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesCountRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateGetRequestInformation retrieve a list of the accessPackage objects that have been marked as incompatible on an accessPackage.  
func (m *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Get retrieve a list of the accessPackage objects that have been marked as incompatible on an accessPackage.  
func (m *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCollectionResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageCollectionResponseable), nil
}
// Ref provides operations to manage the collection of identityGovernance entities.
func (m *IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRequestBuilder) Ref()(*IdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRefRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementCatalogsItemAccessPackagesItemIncompatibleAccessPackagesRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
