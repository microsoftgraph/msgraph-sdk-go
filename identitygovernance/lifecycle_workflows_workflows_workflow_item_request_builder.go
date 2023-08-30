package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder provides operations to manage the workflows property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
type LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetQueryParameters read the properties and relationships of a workflow object.
type LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetQueryParameters
}
// LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderInternal instantiates a new WorkflowItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) {
    m := &LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder instantiates a new WorkflowItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedBy provides operations to manage the createdBy property of the microsoft.graph.identityGovernance.workflowBase entity.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) CreatedBy()(*LifecycleWorkflowsWorkflowsItemCreatedByRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemCreatedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a workflow object and its associated tasks, taskProcessingResults and versions. You can restore a deleted workflow and its associated objects within 30 days of deletion.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflow-delete?view=graph-rest-1.0
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ExecutionScope provides operations to manage the executionScope property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) ExecutionScope()(*LifecycleWorkflowsWorkflowsItemExecutionScopeRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemExecutionScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a workflow object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflow-get?view=graph-rest-1.0
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Workflowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateWorkflowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Workflowable), nil
}
// LastModifiedBy provides operations to manage the lastModifiedBy property of the microsoft.graph.identityGovernance.workflowBase entity.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) LastModifiedBy()(*LifecycleWorkflowsWorkflowsItemLastModifiedByRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemLastModifiedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphIdentityGovernanceActivate provides operations to call the activate method.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) MicrosoftGraphIdentityGovernanceActivate()(*LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceActivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphIdentityGovernanceCreateNewVersion provides operations to call the createNewVersion method.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) MicrosoftGraphIdentityGovernanceCreateNewVersion()(*LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphIdentityGovernanceRestore provides operations to call the restore method.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) MicrosoftGraphIdentityGovernanceRestore()(*LifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceRestoreRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemMicrosoftGraphIdentityGovernanceRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a workflow object. Only the properties listed in the request body table can be updated. To update any other workflow properties, see workflow: createNewVersion.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflow-update?view=graph-rest-1.0
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) Patch(ctx context.Context, body ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Workflowable, requestConfiguration *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Workflowable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateWorkflowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Workflowable), nil
}
// Runs provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) Runs()(*LifecycleWorkflowsWorkflowsItemRunsRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemRunsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaskReports provides operations to manage the taskReports property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) TaskReports()(*LifecycleWorkflowsWorkflowsItemTaskReportsRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemTaskReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.identityGovernance.workflowBase entity.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) Tasks()(*LifecycleWorkflowsWorkflowsItemTasksRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a workflow object and its associated tasks, taskProcessingResults and versions. You can restore a deleted workflow and its associated objects within 30 days of deletion.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a workflow object.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a workflow object. Only the properties listed in the request body table can be updated. To update any other workflow properties, see workflow: createNewVersion.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Workflowable, requestConfiguration *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// UserProcessingResults provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) UserProcessingResults()(*LifecycleWorkflowsWorkflowsItemUserProcessingResultsRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemUserProcessingResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) Versions()(*LifecycleWorkflowsWorkflowsItemVersionsRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder) {
    return NewLifecycleWorkflowsWorkflowsWorkflowItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
