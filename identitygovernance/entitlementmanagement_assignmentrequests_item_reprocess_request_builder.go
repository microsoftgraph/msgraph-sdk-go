package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder provides operations to call the reprocess method.
type EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementAssignmentrequestsItemReprocessRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentrequestsItemReprocessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder) {
    m := &EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignmentRequests/{accessPackageAssignmentRequest%2Did}/reprocess", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder instantiates a new EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentrequestsItemReprocessRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Microsoft Entra entitlement management, callers can automatically retry a user's request for access to an access package. It's performed on an accessPackageAssignmentRequest object whose requestState is in a DeliveryFailed or PartiallyDelivered state.  You can only reprocess a request within 14 days from the time the original request was completed. For requests completed more than 14 days, you will need to ask the users to cancel the request(s) and make a new request in the MyAccess portal.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignmentrequest-reprocess?view=graph-rest-1.0
func (m *EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation in Microsoft Entra entitlement management, callers can automatically retry a user's request for access to an access package. It's performed on an accessPackageAssignmentRequest object whose requestState is in a DeliveryFailed or PartiallyDelivered state.  You can only reprocess a request within 14 days from the time the original request was completed. For requests completed more than 14 days, you will need to ask the users to cancel the request(s) and make a new request in the MyAccess portal.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder when successful
func (m *EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsItemReprocessRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
