package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilder provides operations to call the cancel method.
type RoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// RoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilderInternal instantiates a new CancelRequestBuilder and sets the default values.
func NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilder) {
    m := &RoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/roleManagement/entitlementManagement/roleEligibilityScheduleRequests/{unifiedRoleEligibilityScheduleRequest%2Did}/microsoft.graph.cancel";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilder instantiates a new CancelRequestBuilder and sets the default values.
func NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation immediately cancel a unifiedRoleEligibilityScheduleRequest object whose status is `Granted` and have the system automatically delete the cancelled request after 30 days. After calling this action, the **status** of the cancelled **unifiedRoleEligibilityScheduleRequest** changes to `Revoked`.
func (m *RoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post immediately cancel a unifiedRoleEligibilityScheduleRequest object whose status is `Granted` and have the system automatically delete the cancelled request after 30 days. After calling this action, the **status** of the cancelled **unifiedRoleEligibilityScheduleRequest** changes to `Revoked`.
func (m *RoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *RoleManagementEntitlementManagementRoleEligibilityScheduleRequestsItemCancelRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
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
