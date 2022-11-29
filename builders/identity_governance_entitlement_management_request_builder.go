package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// IdentityGovernanceEntitlementManagementRequestBuilder provides operations to manage the entitlementManagement property of the microsoft.graph.identityGovernance entity.
type IdentityGovernanceEntitlementManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// IdentityGovernanceEntitlementManagementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// IdentityGovernanceEntitlementManagementRequestBuilderGetQueryParameters get entitlementManagement from identityGovernance
type IdentityGovernanceEntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// IdentityGovernanceEntitlementManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *IdentityGovernanceEntitlementManagementRequestBuilderGetQueryParameters
}
// IdentityGovernanceEntitlementManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type IdentityGovernanceEntitlementManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentApprovals provides operations to manage the accessPackageAssignmentApprovals property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageAssignmentApprovals()(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackageAssignmentApprovalsById provides operations to manage the accessPackageAssignmentApprovals property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackageAssignmentApprovalsById(id string)(*IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsApprovalItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["approval%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalsApprovalItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AccessPackages provides operations to manage the accessPackages property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackages()(*IdentityGovernanceEntitlementManagementAccessPackagesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAccessPackagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AccessPackagesById provides operations to manage the accessPackages property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AccessPackagesById(id string)(*IdentityGovernanceEntitlementManagementAccessPackagesAccessPackageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackage%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAccessPackagesAccessPackageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignmentPolicies provides operations to manage the assignmentPolicies property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AssignmentPolicies()(*IdentityGovernanceEntitlementManagementAssignmentPoliciesRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAssignmentPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentPoliciesById provides operations to manage the assignmentPolicies property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AssignmentPoliciesById(id string)(*IdentityGovernanceEntitlementManagementAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentPolicy%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAssignmentPoliciesAccessPackageAssignmentPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignmentRequests provides operations to manage the assignmentRequests property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AssignmentRequests()(*IdentityGovernanceEntitlementManagementAssignmentRequestsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAssignmentRequestsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentRequestsById provides operations to manage the assignmentRequests property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AssignmentRequestsById(id string)(*IdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignmentRequest%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAssignmentRequestsAccessPackageAssignmentRequestItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) Assignments()(*IdentityGovernanceEntitlementManagementAssignmentsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById provides operations to manage the assignments property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) AssignmentsById(id string)(*IdentityGovernanceEntitlementManagementAssignmentsAccessPackageAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageAssignment%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementAssignmentsAccessPackageAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Catalogs provides operations to manage the catalogs property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) Catalogs()(*IdentityGovernanceEntitlementManagementCatalogsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementCatalogsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CatalogsById provides operations to manage the catalogs property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) CatalogsById(id string)(*IdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["accessPackageCatalog%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementCatalogsAccessPackageCatalogItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ConnectedOrganizations provides operations to manage the connectedOrganizations property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) ConnectedOrganizations()(*IdentityGovernanceEntitlementManagementConnectedOrganizationsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementConnectedOrganizationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ConnectedOrganizationsById provides operations to manage the connectedOrganizations property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) ConnectedOrganizationsById(id string)(*IdentityGovernanceEntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["connectedOrganization%2Did"] = id
    }
    return NewIdentityGovernanceEntitlementManagementConnectedOrganizationsConnectedOrganizationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewIdentityGovernanceEntitlementManagementRequestBuilderInternal instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementRequestBuilder) {
    m := &IdentityGovernanceEntitlementManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/identityGovernance/entitlementManagement{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewIdentityGovernanceEntitlementManagementRequestBuilder instantiates a new EntitlementManagementRequestBuilder and sets the default values.
func NewIdentityGovernanceEntitlementManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*IdentityGovernanceEntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewIdentityGovernanceEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property entitlementManagement for identityGovernance
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get entitlementManagement from identityGovernance
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property entitlementManagement in identityGovernance
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EntitlementManagementable, requestConfiguration *IdentityGovernanceEntitlementManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property entitlementManagement for identityGovernance
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) Delete(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entitlementManagement from identityGovernance
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *IdentityGovernanceEntitlementManagementRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EntitlementManagementable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEntitlementManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EntitlementManagementable), nil
}
// Patch update the navigation property entitlementManagement in identityGovernance
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EntitlementManagementable, requestConfiguration *IdentityGovernanceEntitlementManagementRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EntitlementManagementable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEntitlementManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EntitlementManagementable), nil
}
// Settings provides operations to manage the settings property of the microsoft.graph.entitlementManagement entity.
func (m *IdentityGovernanceEntitlementManagementRequestBuilder) Settings()(*IdentityGovernanceEntitlementManagementSettingsRequestBuilder) {
    return NewIdentityGovernanceEntitlementManagementSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
