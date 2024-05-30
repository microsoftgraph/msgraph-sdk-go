package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder provides operations to manage the tasks property of the microsoft.graph.identityGovernance.workflowTemplate entity.
type LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilderGetQueryParameters represents the configured tasks to execute and their execution sequence within a workflow. This relationship is expanded by default.
type LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder) {
    m := &LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflowTemplates/{workflowTemplate%2Did}/tasks/{task%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder instantiates a new LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get represents the configured tasks to execute and their execution sequence within a workflow. This relationship is expanded by default.
// returns a Taskable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Taskable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateTaskFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Taskable), nil
}
// TaskProcessingResults provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.task entity.
// returns a *LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsTaskProcessingResultsRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder) TaskProcessingResults()(*LifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsTaskProcessingResultsRequestBuilder) {
    return NewLifecycleworkflowsWorkflowtemplatesItemTasksItemTaskprocessingresultsTaskProcessingResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation represents the configured tasks to execute and their execution sequence within a workflow. This relationship is expanded by default.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder) {
    return NewLifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
