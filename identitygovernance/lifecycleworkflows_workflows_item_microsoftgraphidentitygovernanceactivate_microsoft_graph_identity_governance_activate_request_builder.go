package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder provides operations to call the activate method.
type LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder) {
    m := &LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/microsoft.graph.identityGovernance.activate", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder instantiates a new LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post run a workflow object on-demand. You can run any workflow on-demand, including scheduled workflows. Workflows created from the 'Real-time employee termination' template are run on-demand only. When you run a workflow on demand, the tasks are executed regardless of whether the user state matches the scope and trigger execution conditions.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflow-activate?view=graph-rest-1.0
func (m *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder) Post(ctx context.Context, body LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateActivatePostRequestBodyable, requestConfiguration *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation run a workflow object on-demand. You can run any workflow on-demand, including scheduled workflows. Workflows created from the 'Real-time employee termination' template are run on-demand only. When you run a workflow on demand, the tasks are executed regardless of whether the user state matches the scope and trigger execution conditions.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder) ToPostRequestInformation(ctx context.Context, body LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateActivatePostRequestBodyable, requestConfiguration *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
