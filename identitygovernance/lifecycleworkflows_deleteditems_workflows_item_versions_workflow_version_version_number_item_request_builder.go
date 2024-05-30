package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
type LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetQueryParameters the workflow versions that are available.
type LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    m := &LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/versions/{workflowVersion%2DversionNumber}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedBy provides operations to manage the createdBy property of the microsoft.graph.identityGovernance.workflowBase entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) CreatedBy()(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the workflow versions that are available.
// returns a WorkflowVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.WorkflowVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) LastModifiedBy()(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.identityGovernance.workflowBase entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemTasksRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) Tasks()(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemTasksRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the workflow versions that are available.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
