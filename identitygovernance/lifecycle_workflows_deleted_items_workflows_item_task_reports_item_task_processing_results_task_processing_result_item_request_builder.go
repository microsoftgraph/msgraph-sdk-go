package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.taskReport entity.
type LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetQueryParameters the related lifecycle workflow taskProcessingResults.
type LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal instantiates a new TaskProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) {
    m := &LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/taskReports/{taskReport%2Did}/taskProcessingResults/{taskProcessingResult%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder instantiates a new TaskProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the related lifecycle workflow taskProcessingResults.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.TaskProcessingResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateTaskProcessingResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.TaskProcessingResultable), nil
}
// MicrosoftGraphIdentityGovernanceResume provides operations to call the resume method.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) MicrosoftGraphIdentityGovernanceResume()(*LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemMicrosoftGraphIdentityGovernanceResumeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subject provides operations to manage the subject property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) Subject()(*LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemSubjectRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemSubjectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Task provides operations to manage the task property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) Task()(*LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsItemTaskRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the related lifecycle workflow taskProcessingResults.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemTaskReportsItemTaskProcessingResultsTaskProcessingResultItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
