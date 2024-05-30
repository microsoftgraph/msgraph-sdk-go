package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder provides operations to call the cancel method.
type PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilderInternal instantiates a new PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder) {
    m := &PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/privilegedAccess/group/eligibilityScheduleRequests/{privilegedAccessGroupEligibilityScheduleRequest%2Did}/cancel", pathParameters),
    }
    return m
}
// NewPrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder instantiates a new PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder and sets the default values.
func NewPrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post cancel an eligibility assignment request to a group whose membership and ownership are governed by PIM.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/privilegedaccessgroupeligibilityschedulerequest-cancel?view=graph-rest-1.0
func (m *PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder) Post(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation cancel an eligibility assignment request to a group whose membership and ownership are governed by PIM.
// returns a *RequestInformation when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder when successful
func (m *PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder) WithUrl(rawUrl string)(*PrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder) {
    return NewPrivilegedaccessGroupEligibilityschedulerequestsItemCancelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
