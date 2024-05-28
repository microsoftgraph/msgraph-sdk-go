package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder provides operations to manage the taskProcessingResults property of the microsoft.graph.identityGovernance.userProcessingResult entity.
type LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilderGetQueryParameters the associated individual task execution.
type LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilderGetQueryParameters
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilderInternal instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder) {
    m := &LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/userProcessingResults/{userProcessingResult%2Did}/taskProcessingResults/{taskProcessingResult%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the associated individual task execution.
// returns a TaskProcessingResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.TaskProcessingResultable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemMicrosoftgraphidentitygovernanceresumeMicrosoftGraphIdentityGovernanceResumeRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder) MicrosoftGraphIdentityGovernanceResume()(*LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemMicrosoftgraphidentitygovernanceresumeMicrosoftGraphIdentityGovernanceResumeRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemMicrosoftgraphidentitygovernanceresumeMicrosoftGraphIdentityGovernanceResumeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Subject provides operations to manage the subject property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder) Subject()(*LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemSubjectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Task provides operations to manage the task property of the microsoft.graph.identityGovernance.taskProcessingResult entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder) Task()(*LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsItemTaskRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation the associated individual task execution.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemUserprocessingresultsItemTaskprocessingresultsTaskProcessingResultItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
