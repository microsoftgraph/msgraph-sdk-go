package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder provides operations to manage the taskDefinition property of the microsoft.graph.identityGovernance.taskReport entity.
type LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetQueryParameters the taskDefinition associated with the related lifecycle workflow task.Supports $filter(eq, ne) and $expand.
type LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder) {
    m := &LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/taskReports/{taskReport%2Did}/taskDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder instantiates a new LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the taskDefinition associated with the related lifecycle workflow task.Supports $filter(eq, ne) and $expand.
// returns a TaskDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.TaskDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateTaskDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.TaskDefinitionable), nil
}
// ToGetRequestInformation the taskDefinition associated with the related lifecycle workflow task.Supports $filter(eq, ne) and $expand.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemTaskreportsItemTaskdefinitionTaskDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
