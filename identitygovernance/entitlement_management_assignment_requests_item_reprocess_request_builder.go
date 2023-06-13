package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementAssignmentRequestsItemReprocessRequestBuilder provides operations to call the reprocess method.
type EntitlementManagementAssignmentRequestsItemReprocessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAssignmentRequestsItemReprocessRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAssignmentRequestsItemReprocessRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementAssignmentRequestsItemReprocessRequestBuilderInternal instantiates a new ReprocessRequestBuilder and sets the default values.
func NewEntitlementManagementAssignmentRequestsItemReprocessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAssignmentRequestsItemReprocessRequestBuilder) {
    m := &EntitlementManagementAssignmentRequestsItemReprocessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignmentRequests/{accessPackageAssignmentRequest%2Did}/reprocess", pathParameters),
    }
    return m
}
// NewEntitlementManagementAssignmentRequestsItemReprocessRequestBuilder instantiates a new ReprocessRequestBuilder and sets the default values.
func NewEntitlementManagementAssignmentRequestsItemReprocessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAssignmentRequestsItemReprocessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAssignmentRequestsItemReprocessRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Azure AD entitlement management, callers can automatically retry a user's request for access to an access package. It is performed on an accessPackageAssignmentRequest object whose **requestState** is in a `DeliveryFailed` or `PartiallyDelivered` state.  You can only reprocess a request within 14 days from the time the original request was completed. For requests completed more than 14 days, you will need to ask the users to cancel the request(s) and make a new request in the MyAccess portal.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/accesspackageassignmentrequest-reprocess?view=graph-rest-1.0
func (m *EntitlementManagementAssignmentRequestsItemReprocessRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementManagementAssignmentRequestsItemReprocessRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation in Azure AD entitlement management, callers can automatically retry a user's request for access to an access package. It is performed on an accessPackageAssignmentRequest object whose **requestState** is in a `DeliveryFailed` or `PartiallyDelivered` state.  You can only reprocess a request within 14 days from the time the original request was completed. For requests completed more than 14 days, you will need to ask the users to cancel the request(s) and make a new request in the MyAccess portal.
func (m *EntitlementManagementAssignmentRequestsItemReprocessRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAssignmentRequestsItemReprocessRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
