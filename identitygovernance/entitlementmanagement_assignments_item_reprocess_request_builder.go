package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentsItemReprocessRequestBuilder provides operations to call the reprocess method.
type EntitlementmanagementAssignmentsItemReprocessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentsItemReprocessRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentsItemReprocessRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementAssignmentsItemReprocessRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentsItemReprocessRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentsItemReprocessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentsItemReprocessRequestBuilder) {
    m := &EntitlementmanagementAssignmentsItemReprocessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignments/{accessPackageAssignment%2Did}/reprocess", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAssignmentsItemReprocessRequestBuilder instantiates a new EntitlementmanagementAssignmentsItemReprocessRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentsItemReprocessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentsItemReprocessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentsItemReprocessRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Microsoft Entra entitlement management, callers can automatically reevaluate and enforce an accessPackageAssignment object of a user’s assignments for a specific access package. The state of the access package assignment must be Delivered for the administrator to reprocess the user's assignment. Only admins with the Access Package Assignment Manager role, or higher, in Microsoft Entra entitlement management can perform this action.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignment-reprocess?view=graph-rest-1.0
func (m *EntitlementmanagementAssignmentsItemReprocessRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsItemReprocessRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation in Microsoft Entra entitlement management, callers can automatically reevaluate and enforce an accessPackageAssignment object of a user’s assignments for a specific access package. The state of the access package assignment must be Delivered for the administrator to reprocess the user's assignment. Only admins with the Access Package Assignment Manager role, or higher, in Microsoft Entra entitlement management can perform this action.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentsItemReprocessRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentsItemReprocessRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementAssignmentsItemReprocessRequestBuilder when successful
func (m *EntitlementmanagementAssignmentsItemReprocessRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentsItemReprocessRequestBuilder) {
    return NewEntitlementmanagementAssignmentsItemReprocessRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
