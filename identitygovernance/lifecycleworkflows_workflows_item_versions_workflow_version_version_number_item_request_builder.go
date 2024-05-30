package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
type LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetQueryParameters read the properties and relationships of a workflowVersion object.
type LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    m := &LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/versions/{workflowVersion%2DversionNumber}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder instantiates a new LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatedBy provides operations to manage the createdBy property of the microsoft.graph.identityGovernance.workflowBase entity.
// returns a *LifecycleworkflowsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) CreatedBy()(*LifecycleworkflowsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemVersionsItemCreatedbyCreatedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a workflowVersion object.
// returns a WorkflowVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-workflowversion-get?view=graph-rest-1.0
func (m *LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.WorkflowVersionable, error) {
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
// returns a *LifecycleworkflowsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) LastModifiedBy()(*LifecycleworkflowsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemVersionsItemLastmodifiedbyLastModifiedByRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tasks provides operations to manage the tasks property of the microsoft.graph.identityGovernance.workflowBase entity.
// returns a *LifecycleworkflowsWorkflowsItemVersionsItemTasksRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) Tasks()(*LifecycleworkflowsWorkflowsItemVersionsItemTasksRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemVersionsItemTasksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation read the properties and relationships of a workflowVersion object.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
