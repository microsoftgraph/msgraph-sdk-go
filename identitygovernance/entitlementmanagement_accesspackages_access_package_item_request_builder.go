package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder provides operations to manage the accessPackages property of the microsoft.graph.entitlementManagement entity.
type EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetQueryParameters retrieve the properties and relationships of an accessPackage object.
type EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackagesIncompatibleWith provides operations to manage the accessPackagesIncompatibleWith property of the microsoft.graph.accessPackage entity.
// returns a *EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackagesIncompatibleWithRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) AccessPackagesIncompatibleWith()(*EntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackagesIncompatibleWithRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAccesspackagesincompatiblewithAccessPackagesIncompatibleWithRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentPolicies provides operations to manage the assignmentPolicies property of the microsoft.graph.accessPackage entity.
// returns a *EntitlementmanagementAccesspackagesItemAssignmentpoliciesAssignmentPoliciesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) AssignmentPolicies()(*EntitlementmanagementAccesspackagesItemAssignmentpoliciesAssignmentPoliciesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemAssignmentpoliciesAssignmentPoliciesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Catalog provides operations to manage the catalog property of the microsoft.graph.accessPackage entity.
// returns a *EntitlementmanagementAccesspackagesItemCatalogRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) Catalog()(*EntitlementmanagementAccesspackagesItemCatalogRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemCatalogRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderInternal instantiates a new EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) {
    m := &EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/accessPackages/{accessPackage%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder instantiates a new EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an accessPackage object. You cannot delete an access package if it has any accessPackageAssignment.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-delete?view=graph-rest-1.0
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of an accessPackage object.
// returns a AccessPackageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-get?view=graph-rest-1.0
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable), nil
}
// GetApplicablePolicyRequirements provides operations to call the getApplicablePolicyRequirements method.
// returns a *EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) GetApplicablePolicyRequirements()(*EntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemGetapplicablepolicyrequirementsGetApplicablePolicyRequirementsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IncompatibleAccessPackages provides operations to manage the incompatibleAccessPackages property of the microsoft.graph.accessPackage entity.
// returns a *EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) IncompatibleAccessPackages()(*EntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatibleaccesspackagesIncompatibleAccessPackagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IncompatibleGroups provides operations to manage the incompatibleGroups property of the microsoft.graph.accessPackage entity.
// returns a *EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) IncompatibleGroups()(*EntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemIncompatiblegroupsIncompatibleGroupsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update an existing accessPackage object to change one or more of its properties, such as the display name or description.
// returns a AccessPackageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackage-update?view=graph-rest-1.0
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, requestConfiguration *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable), nil
}
// ResourceRoleScopes provides operations to manage the resourceRoleScopes property of the microsoft.graph.accessPackage entity.
// returns a *EntitlementmanagementAccesspackagesItemResourcerolescopesResourceRoleScopesRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) ResourceRoleScopes()(*EntitlementmanagementAccesspackagesItemResourcerolescopesResourceRoleScopesRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesItemResourcerolescopesResourceRoleScopesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete an accessPackage object. You cannot delete an access package if it has any accessPackageAssignment.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve the properties and relationships of an accessPackage object.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update an existing accessPackage object to change one or more of its properties, such as the display name or description.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageable, requestConfiguration *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder when successful
func (m *EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder) {
    return NewEntitlementmanagementAccesspackagesAccessPackageItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
