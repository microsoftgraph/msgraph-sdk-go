package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder provides operations to manage the executionScope property of the microsoft.graph.identityGovernance.workflow entity.
type LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilderGetQueryParameters the unique identifier of the Microsoft Entra identity that last modified the workflow object.
type LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilderGetQueryParameters
}
// ByUserProcessingResultId provides operations to manage the executionScope property of the microsoft.graph.identityGovernance.workflow entity.
// returns a *LifecycleworkflowsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder) ByUserProcessingResultId(userProcessingResultId string)(*LifecycleworkflowsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userProcessingResultId != "" {
        urlTplParams["userProcessingResult%2Did"] = userProcessingResultId
    }
    return NewLifecycleworkflowsWorkflowsItemExecutionscopeUserProcessingResultItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder) {
    m := &LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflows/{workflow%2Did}/executionScope{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder instantiates a new LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *LifecycleworkflowsWorkflowsItemExecutionscopeCountRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder) Count()(*LifecycleworkflowsWorkflowsItemExecutionscopeCountRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemExecutionscopeCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the unique identifier of the Microsoft Entra identity that last modified the workflow object.
// returns a UserProcessingResultCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.UserProcessingResultCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateUserProcessingResultCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.UserProcessingResultCollectionResponseable), nil
}
// ToGetRequestInformation the unique identifier of the Microsoft Entra identity that last modified the workflow object.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder) {
    return NewLifecycleworkflowsWorkflowsItemExecutionscopeExecutionScopeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
