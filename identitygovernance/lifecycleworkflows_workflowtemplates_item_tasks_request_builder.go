package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder provides operations to manage the tasks property of the microsoft.graph.identityGovernance.workflowTemplate entity.
type LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilderGetQueryParameters represents the configured tasks to execute and their execution sequence within a workflow. This relationship is expanded by default.
type LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilderGetQueryParameters
}
// ByTaskId provides operations to manage the tasks property of the microsoft.graph.identityGovernance.workflowTemplate entity.
// returns a *LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder) ByTaskId(taskId string)(*LifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if taskId != "" {
        urlTplParams["task%2Did"] = taskId
    }
    return NewLifecycleworkflowsWorkflowtemplatesItemTasksTaskItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder) {
    m := &LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflowTemplates/{workflowTemplate%2Did}/tasks{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder instantiates a new LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *LifecycleworkflowsWorkflowtemplatesItemTasksCountRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder) Count()(*LifecycleworkflowsWorkflowtemplatesItemTasksCountRequestBuilder) {
    return NewLifecycleworkflowsWorkflowtemplatesItemTasksCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents the configured tasks to execute and their execution sequence within a workflow. This relationship is expanded by default.
// returns a TaskCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.TaskCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateTaskCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.TaskCollectionResponseable), nil
}
// ToGetRequestInformation represents the configured tasks to execute and their execution sequence within a workflow. This relationship is expanded by default.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder) {
    return NewLifecycleworkflowsWorkflowtemplatesItemTasksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
