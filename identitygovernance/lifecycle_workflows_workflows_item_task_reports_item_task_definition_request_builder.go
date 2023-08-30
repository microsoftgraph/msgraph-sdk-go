package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilder provides operations to manage the taskDefinition property of the microsoft.graph.identityGovernance.taskReport entity.
type LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilderGetQueryParameters the taskDefinition associated with the related lifecycle workflow task.Supports $filter(eq, ne) and $expand.
type LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilderInternal instantiates a new TaskDefinitionRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilder) {
    m := &LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/taskReports/{taskReport%2Did}/taskDefinition{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilder instantiates a new TaskDefinitionRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the taskDefinition associated with the related lifecycle workflow task.Supports $filter(eq, ne) and $expand.
func (m *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.TaskDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilder) WithUrl(rawUrl string)(*LifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemTaskReportsItemTaskDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
