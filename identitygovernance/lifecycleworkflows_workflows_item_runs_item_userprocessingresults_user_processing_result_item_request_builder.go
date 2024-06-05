package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.run entity.
type LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetQueryParameters get the user processing result of a user processing result of a run.
type LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) {
    m := &LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/runs/{run%2Did}/userProcessingResults/{userProcessingResult%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder instantiates a new LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get the user processing result of a user processing result of a run.
// returns a UserProcessingResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-userprocessingresult-get?view=graph-rest-1.0
func (m *LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.UserProcessingResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateUserProcessingResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.UserProcessingResultable), nil
}
// Subject provides operations to manage the subject property of the microsoft.graph.identityGovernance.userProcessingResult entity.
// returns a *LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsItemSubjectRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) Subject()(*LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsItemSubjectRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsItemSubjectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaskProcessingResults provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.userProcessingResult entity.
// returns a *LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) TaskProcessingResults()(*LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation get the user processing result of a user processing result of a run.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
