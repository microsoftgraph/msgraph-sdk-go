package identitygovernance

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
type LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilderGetQueryParameters the workflow versions that are available.
type LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilderGetQueryParameters struct {
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
// LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilderGetQueryParameters
}
// ByWorkflowVersionVersionNumber provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
// Deprecated: This indexer is deprecated and will be removed in the next major version. Use the one with the typed parameter instead.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder) ByWorkflowVersionVersionNumber(workflowVersionVersionNumber string)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if workflowVersionVersionNumber != "" {
        urlTplParams["workflowVersion%2DversionNumber"] = workflowVersionVersionNumber
    }
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// ByWorkflowVersionVersionNumberInteger provides operations to manage the versions property of the microsoft.graph.identityGovernance.workflow entity.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder) ByWorkflowVersionVersionNumberInteger(workflowVersionVersionNumber int32)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    urlTplParams["workflowVersion%2DversionNumber"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(workflowVersionVersionNumber), 10)
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsWorkflowVersionVersionNumberItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilderInternal instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder) {
    m := &LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/deletedItems/workflows/{workflow%2Did}/versions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder instantiates a new LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder and sets the default values.
func NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsCountRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder) Count()(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsCountRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the workflow versions that are available.
// returns a WorkflowVersionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.WorkflowVersionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateWorkflowVersionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.WorkflowVersionCollectionResponseable), nil
}
// ToGetRequestInformation the workflow versions that are available.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder when successful
func (m *LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder) {
    return NewLifecycleworkflowsDeleteditemsWorkflowsItemVersionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
