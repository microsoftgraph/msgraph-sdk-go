package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementEntitlementManagementRequestBuilder provides operations to manage the entitlementManagement property of the microsoft.graph.identityGovernance entity.
type EntitlementmanagementEntitlementManagementRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementEntitlementManagementRequestBuilderGetQueryParameters get entitlementManagement from identityGovernance
type EntitlementmanagementEntitlementManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementEntitlementManagementRequestBuilderGetQueryParameters
}
// EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackageAssignmentApprovals provides operations to manage the accessPackageAssignmentApprovals property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackageassignmentapprovalsAccessPackageAssignmentApprovalsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AccessPackageAssignmentApprovals()(*EntitlementmanagementAccesspackageassignmentapprovalsAccessPackageAssignmentApprovalsRequestBuilder) {
    return NewEntitlementmanagementAccesspackageassignmentapprovalsAccessPackageAssignmentApprovalsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AccessPackages provides operations to manage the accessPackages property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAccesspackagesAccessPackagesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AccessPackages()(*EntitlementmanagementAccesspackagesAccessPackagesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesAccessPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentPolicies provides operations to manage the assignmentPolicies property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAssignmentpoliciesAssignmentPoliciesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AssignmentPolicies()(*EntitlementmanagementAssignmentpoliciesAssignmentPoliciesRequestBuilder) {
    return NewEntitlementmanagementAssignmentpoliciesAssignmentPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentRequests provides operations to manage the assignmentRequests property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAssignmentrequestsAssignmentRequestsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) AssignmentRequests()(*EntitlementmanagementAssignmentrequestsAssignmentRequestsRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsAssignmentRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Assignments provides operations to manage the assignments property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementAssignmentsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Assignments()(*EntitlementmanagementAssignmentsRequestBuilder) {
    return NewEntitlementmanagementAssignmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Catalogs provides operations to manage the catalogs property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementCatalogsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Catalogs()(*EntitlementmanagementCatalogsRequestBuilder) {
    return NewEntitlementmanagementCatalogsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ConnectedOrganizations provides operations to manage the connectedOrganizations property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ConnectedOrganizations()(*EntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilder) {
    return NewEntitlementmanagementConnectedorganizationsConnectedOrganizationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementEntitlementManagementRequestBuilderInternal instantiates a new EntitlementmanagementEntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementmanagementEntitlementManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementEntitlementManagementRequestBuilder) {
    m := &EntitlementmanagementEntitlementManagementRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementEntitlementManagementRequestBuilder instantiates a new EntitlementmanagementEntitlementManagementRequestBuilder and sets the default values.
func NewEntitlementmanagementEntitlementManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementEntitlementManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementEntitlementManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property entitlementManagement for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get entitlementManagement from identityGovernance
// returns a EntitlementManagementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EntitlementManagementable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEntitlementManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EntitlementManagementable), nil
}
// Patch update the navigation property entitlementManagement in identityGovernance
// returns a EntitlementManagementable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EntitlementManagementable, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EntitlementManagementable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEntitlementManagementFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EntitlementManagementable), nil
}
// ResourceEnvironments provides operations to manage the resourceEnvironments property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ResourceEnvironments()(*EntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilder) {
    return NewEntitlementmanagementResourceenvironmentsResourceEnvironmentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResourceRequests provides operations to manage the resourceRequests property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementResourcerequestsResourceRequestsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ResourceRequests()(*EntitlementmanagementResourcerequestsResourceRequestsRequestBuilder) {
    return NewEntitlementmanagementResourcerequestsResourceRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ResourceRoleScopes provides operations to manage the resourceRoleScopes property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementResourcerolescopesResourceRoleScopesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ResourceRoleScopes()(*EntitlementmanagementResourcerolescopesResourceRoleScopesRequestBuilder) {
    return NewEntitlementmanagementResourcerolescopesResourceRoleScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Resources provides operations to manage the resources property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementResourcesRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Resources()(*EntitlementmanagementResourcesRequestBuilder) {
    return NewEntitlementmanagementResourcesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Settings provides operations to manage the settings property of the microsoft.graph.entitlementManagement entity.
// returns a *EntitlementmanagementSettingsRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) Settings()(*EntitlementmanagementSettingsRequestBuilder) {
    return NewEntitlementmanagementSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property entitlementManagement for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get entitlementManagement from identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property entitlementManagement in identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EntitlementManagementable, requestConfiguration *EntitlementmanagementEntitlementManagementRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementEntitlementManagementRequestBuilder when successful
func (m *EntitlementmanagementEntitlementManagementRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementEntitlementManagementRequestBuilder) {
    return NewEntitlementmanagementEntitlementManagementRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
