package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
type LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetQueryParameters read the properties and relationships of a workflowVersion object.
type LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetQueryParameters
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal instantiates a new WorkflowVersionVersionNumberItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    m := &LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/versions/{workflowVersion%2DversionNumber}{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder instantiates a new WorkflowVersionVersionNumberItemRequestBuilder and sets the default values.
func NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedBy provides operations to manage the createdBy property of the microsoft.graph.identityGovernance.workflowBase entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) CreatedBy()(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemCreatedByRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemCreatedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a workflowVersion object.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflowversion-get?view=graph-rest-1.0
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.WorkflowVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateWorkflowVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.WorkflowVersionable), nil
}
// LastModifiedBy provides operations to manage the lastModifiedBy property of the microsoft.graph.identityGovernance.workflowBase entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) LastModifiedBy()(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemLastModifiedByRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemLastModifiedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.identityGovernance.workflowBase entity.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) Tasks()(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation read the properties and relationships of a workflowVersion object.
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    return NewLifecycleWorkflowsDeletedItemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
