package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder provides operations to manage the workflows property of the microsoft.graph.deletedItemContainer entity.
type LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderGetQueryParameters retrieve a deleted workflow object.
type LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderInternal instantiates a new LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) {
    m := &LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder instantiates a new LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedBy provides operations to manage the createdBy property of the microsoft.graph.identityGovernance.workflowBase entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemCreatedbyCreatedByRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) CreatedBy()(*LifecycleworkflowsDeleteditemsWorkflowsItemCreatedbyCreatedByRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemCreatedbyCreatedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete a workflow object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-deletedItemcontainer-delete?view=graph-rest-1.0
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) ExecutionScope()(*LifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemExecutionscopeExecutionScopeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a deleted workflow object.
// returns a Workflowable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-deleteditemcontainer-get?view=graph-rest-1.0
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.Workflowable, error) {
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemLastmodifiedbyLastModifiedByRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) LastModifiedBy()(*LifecycleworkflowsDeleteditemsWorkflowsItemLastmodifiedbyLastModifiedByRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemLastmodifiedbyLastModifiedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphIdentityGovernanceActivate provides operations to call the activate method.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) MicrosoftGraphIdentityGovernanceActivate()(*LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernanceactivateMicrosoftGraphIdentityGovernanceActivateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphIdentityGovernanceCreateNewVersion provides operations to call the createNewVersion method.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) MicrosoftGraphIdentityGovernanceCreateNewVersion()(*LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancecreatenewversionMicrosoftGraphIdentityGovernanceCreateNewVersionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// MicrosoftGraphIdentityGovernanceRestore provides operations to call the restore method.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) MicrosoftGraphIdentityGovernanceRestore()(*LifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemMicrosoftgraphidentitygovernancerestoreMicrosoftGraphIdentityGovernanceRestoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Runs provides operations to manage the runs property of the microsoft.graph.identityGovernance.workflow entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemRunsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) Runs()(*LifecycleworkflowsDeleteditemsWorkflowsItemRunsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaskReports provides operations to manage the taskReports property of the microsoft.graph.identityGovernance.workflow entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) TaskReports()(*LifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemTaskreportsTaskReportsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.identityGovernance.workflowBase entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemTasksRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) Tasks()(*LifecycleworkflowsDeleteditemsWorkflowsItemTasksRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a workflow object.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve a deleted workflow object.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserProcessingResults provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.workflow entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsUserProcessingResultsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) UserProcessingResults()(*LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsUserProcessingResultsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsUserProcessingResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Versions provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) Versions()(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsWorkflowItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
