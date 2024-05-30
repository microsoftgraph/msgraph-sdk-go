package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder provides operations to manage the userProcessingResults property of the microsoft.graph.identityGovernance.run entity.
type LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetQueryParameters the associated individual user execution.
type LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderInternal instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) {
    m := &LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/runs/{run%2Did}/userProcessingResults/{userProcessingResult%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the associated individual user execution.
// returns a UserProcessingResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.UserProcessingResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateUserProcessingResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.UserProcessingResultable), nil
}
// Subject provides operations to manage the subject property of the microsoft.graph.identityGovernance.userProcessingResult entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsItemSubjectRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) Subject()(*LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsItemSubjectRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsItemSubjectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TaskProcessingResults provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.userProcessingResult entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) TaskProcessingResults()(*LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the associated individual user execution.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemRunsItemUserprocessingresultsUserProcessingResultItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
