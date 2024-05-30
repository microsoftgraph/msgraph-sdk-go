package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder provides operations to manage the assignmentPolicy property of the microsoft.graph.accessPackageAssignment entity.
type EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilderGetQueryParameters read-only. Supports $filter (eq) on the id property and $expand query parameters.
type EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilderGetQueryParameters
}
// NewEntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder) {
    m := &EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignments/{accessPackageAssignment%2Did}/assignmentPolicy{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder instantiates a new EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilderInternal(urlParams, requestAdapter)
}
// Get read-only. Supports $filter (eq) on the id property and $expand query parameters.
// returns a AccessPackageAssignmentPolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder) Get(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessPackageAssignmentPolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessPackageAssignmentPolicyable), nil
}
// ToGetRequestInformation read-only. Supports $filter (eq) on the id property and $expand query parameters.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder) {
    return NewEntitlementmanagementAssignmentsItemAssignmentpolicyAssignmentPolicyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
