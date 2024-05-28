package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder provides operations to call the resume method.
type EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EntitlementmanagementAssignmentrequestsItemResumeRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EntitlementmanagementAssignmentrequestsItemResumeRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEntitlementmanagementAssignmentrequestsItemResumeRequestBuilderInternal instantiates a new EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentrequestsItemResumeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder) {
    m := &EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/entitlementManagement/assignmentRequests/{accessPackageAssignmentRequest%2Did}/resume", pathParameters),
    }
    return m
}
// NewEntitlementmanagementAssignmentrequestsItemResumeRequestBuilder instantiates a new EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder and sets the default values.
func NewEntitlementmanagementAssignmentrequestsItemResumeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEntitlementmanagementAssignmentrequestsItemResumeRequestBuilderInternal(urlParams, requestAdapter)
}
// Post in Microsoft Entra entitlement management, when an access package policy has been enabled to call out a custom extension and the request processing is waiting for the callback from the customer, the customer can initiate a resume action. It is performed on an accessPackageAssignmentRequest object whose requestStatus is in a WaitingForCallback state.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/accesspackageassignmentrequest-resume?view=graph-rest-1.0
func (m *EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder) Post(ctx context.Context, body EntitlementmanagementAssignmentrequestsItemResumePostRequestBodyable, requestConfiguration *EntitlementmanagementAssignmentrequestsItemResumeRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation in Microsoft Entra entitlement management, when an access package policy has been enabled to call out a custom extension and the request processing is waiting for the callback from the customer, the customer can initiate a resume action. It is performed on an accessPackageAssignmentRequest object whose requestStatus is in a WaitingForCallback state.
// returns a *RequestInformation when successful
func (m *EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder) ToPostRequestInformation(ctx context.Context, body EntitlementmanagementAssignmentrequestsItemResumePostRequestBodyable, requestConfiguration *EntitlementmanagementAssignmentrequestsItemResumeRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder when successful
func (m *EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder) WithUrl(rawUrl string)(*EntitlementmanagementAssignmentrequestsItemResumeRequestBuilder) {
    return NewEntitlementmanagementAssignmentrequestsItemResumeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
