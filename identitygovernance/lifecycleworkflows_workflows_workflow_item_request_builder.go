package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder provides operations to manage the workflows property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
type LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderGetQueryParameters read the properties and relationships of a workflow object.
type LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderGetQueryParameters
}
// LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLifecycleworkflowsWorkflowsWorkflowItemRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsWorkflowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) {
    m := &LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowsWorkflowItemRequestBuilder instantiates a new LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsWorkflowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowsWorkflowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedBy provides operations to manage the createdBy property of the microsoft.graph.identityGovernance.workflowBase entity.
// returns a *LifecycleworkflowsWorkflowsItemCreatedbyCreatedByRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) CreatedBy()(*LifecycleworkflowsWorkflowsItemCreatedbyCreatedByRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemCreatedbyCreatedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a workflow object and its associated tasks, taskProcessingResults and versions. You can restore a deleted workflow and its associated objects within 30 days of deletion.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflow-delete?view=graph-rest-1.0
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// ExecutionScope provides operations to manage the executionScope property of the microsoft.graph.identityGovernance.workflow entity.
// returns a *LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) ExecutionScope()(*LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a workflow object.
// returns a Workflowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflow-get?view=graph-rest-1.0
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Workflowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// returns a *LifecycleworkflowsWorkflowsItemLastmodifiedbyLastModifiedByRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) LastModifiedBy()(*LifecycleworkflowsWorkflowsItemLastmodifiedbyLastModifiedByRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemLastmodifiedbyLastModifiedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphIdentityGovernanceActivate provides operations to call the activate method.
// returns a *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) MicrosoftGraphIdentityGovernanceActivate()(*LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphIdentityGovernanceCreateNewVersion provides operations to call the createNewVersion method.
// returns a *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) MicrosoftGraphIdentityGovernanceCreateNewVersion()(*LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphIdentityGovernanceRestore provides operations to call the restore method.
// returns a *LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) MicrosoftGraphIdentityGovernanceRestore()(*LifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of a workflow object. Only the properties listed in the request body table can be updated. To update any other workflow properties, see workflow: createNewVersion.
// returns a Workflowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflow-update?view=graph-rest-1.0
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) Patch(ctx context.Context, body ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Workflowable, requestConfiguration *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Workflowable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// returns a *LifecycleworkflowsWorkflowsItemRunsRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) Runs()(*LifecycleworkflowsWorkflowsItemRunsRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemRunsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaskReports provides operations to manage the taskReports property of the microsoft.graph.identityGovernance.workflow entity.
// returns a *LifecycleworkflowsWorkflowsItemTaskreportsTaskReportsRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) TaskReports()(*LifecycleworkflowsWorkflowsItemTaskreportsTaskReportsRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemTaskreportsTaskReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.identityGovernance.workflowBase entity.
// returns a *LifecycleworkflowsWorkflowsItemTasksRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) Tasks()(*LifecycleworkflowsWorkflowsItemTasksRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a workflow object and its associated tasks, taskProcessingResults and versions. You can restore a deleted workflow and its associated objects within 30 days of deletion.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a workflow object.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a workflow object. Only the properties listed in the request body table can be updated. To update any other workflow properties, see workflow: createNewVersion.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Workflowable, requestConfiguration *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// UserProcessingResults provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.workflow entity.
// returns a *LifecycleworkflowsWorkflowsItemUserprocessingresultsUserProcessingResultsRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) UserProcessingResults()(*LifecycleworkflowsWorkflowsItemUserprocessingresultsUserProcessingResultsRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemUserprocessingresultsUserProcessingResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
// returns a *LifecycleworkflowsWorkflowsItemVersionsRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) Versions()(*LifecycleworkflowsWorkflowsItemVersionsRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowsWorkflowItemRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsWorkflowItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
