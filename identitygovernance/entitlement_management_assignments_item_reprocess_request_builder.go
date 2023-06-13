package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementManagementAssignmentsItemReprocessRequestBuilder provides operations to call the reprocess method.
type EntitlementManagementAssignmentsItemReprocessRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementManagementAssignmentsItemReprocessRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementManagementAssignmentsItemReprocessRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementManagementAssignmentsItemReprocessRequestBuilderInternal instantiates a new ReprocessRequestBuilder and sets the default values.
func NewEntitlementManagementAssignmentsItemReprocessRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAssignmentsItemReprocessRequestBuilder) {
    m := &EntitlementManagementAssignmentsItemReprocessRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignments/{accessPackageAssignment%2Did}/reprocess", pathParameters),
    }
    return m
}
// NewEntitlementManagementAssignmentsItemReprocessRequestBuilder instantiates a new ReprocessRequestBuilder and sets the default values.
func NewEntitlementManagementAssignmentsItemReprocessRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementManagementAssignmentsItemReprocessRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementManagementAssignmentsItemReprocessRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Azure AD entitlement management, callers can automatically reevaluate and enforce an accessPackageAssignment object of a user’s assignments for a specific access package. The **state** of the access package assignment must be `Delivered` for the administrator to reprocess the user's assignment. Only admins with the Access Package Assignment Manager role, or higher, in Azure AD entitlement management can perform this action.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/accesspackageassignment-reprocess?view=graph-rest-1.0
func (m *EntitlementManagementAssignmentsItemReprocessRequestBuilder) Post(ctx context.Context, requestConfiguration *EntitlementManagementAssignmentsItemReprocessRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation in Azure AD entitlement management, callers can automatically reevaluate and enforce an accessPackageAssignment object of a user’s assignments for a specific access package. The **state** of the access package assignment must be `Delivered` for the administrator to reprocess the user's assignment. Only admins with the Access Package Assignment Manager role, or higher, in Azure AD entitlement management can perform this action.
func (m *EntitlementManagementAssignmentsItemReprocessRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *EntitlementManagementAssignmentsItemReprocessRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
