package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.entitlementManagement entity.
type EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderGetQueryParameters in Microsoft Entra entitlement management, retrieve the properties and relationships of an accessPackageAssignment object.
type EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderGetQueryParameters
}
// EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessPackage provides operations to manage the accessPackage property of the microsoft.graph.accessPackageAssignment entity.
// returns a *EntitlementmanagementAssignmentsItemAccesspackageAccessPackageRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) AccessPackage()(*EntitlementmanagementAssignmentsItemAccesspackageAccessPackageRequestBuilder) {
    return NewEntitlementmanagementAssignmentsItemAccesspackageAccessPackageRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// AssignmentPolicy provides operations to manage the assignmentPolicy property of the microsoft.graph.accessPackageAssignment entity.
// returns a *EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) AssignmentPolicy()(*EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder) {
    return NewEntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewEntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) {
    m := &EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignments/{accessPackageAssignment%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder instantiates a new EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignments for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get in Microsoft Entra entitlement management, retrieve the properties and relationships of an accessPackageAssignment object.
// returns a AccessPackageAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignment-get?view=graph-rest-1.0
func (m *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable), nil
}
// Patch update the navigation property assignments in identityGovernance
// returns a AccessPackageAssignmentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable, requestConfiguration *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable), nil
}
// Reprocess provides operations to call the reprocess method.
// returns a *EntitlementmanagementAssignmentsItemReprocessRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) Reprocess()(*EntitlementmanagementAssignmentsItemReprocessRequestBuilder) {
    return NewEntitlementmanagementAssignmentsItemReprocessRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Target provides operations to manage the target property of the microsoft.graph.accessPackageAssignment entity.
// returns a *EntitlementmanagementAssignmentsItemTargetRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) Target()(*EntitlementmanagementAssignmentsItemTargetRequestBuilder) {
    return NewEntitlementmanagementAssignmentsItemTargetRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property assignments for identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation in Microsoft Entra entitlement management, retrieve the properties and relationships of an accessPackageAssignment object.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignments in identityGovernance
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentable, requestConfiguration *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder) {
    return NewEntitlementmanagementAssignmentsAccessPackageAssignmentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
