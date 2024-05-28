package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder provides operations to manage the workflowTemplates property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
type LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilderGetQueryParameters get a list of the workflowTemplate objects and their properties.
type LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilderGetQueryParameters struct {
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
// LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilderGetQueryParameters
}
// ByWorkflowTemplateId provides operations to manage the workflowTemplates property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
// returns a *LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder) ByWorkflowTemplateId(workflowTemplateId string)(*LifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if workflowTemplateId != "" {
        urlTplParams["workflowTemplate%2Did"] = workflowTemplateId
    }
    return NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplateItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilderInternal instantiates a new LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder) {
    m := &LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/workflowTemplates{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder instantiates a new LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder and sets the default values.
func NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *LifecycleworkflowsWorkflowtemplatesCountRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder) Count()(*LifecycleworkflowsWorkflowtemplatesCountRequestBuilder) {
    return NewLifecycleworkflowsWorkflowtemplatesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the workflowTemplate objects and their properties.
// returns a WorkflowTemplateCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-lifecycleworkflowscontainer-list-workflowtemplates?view=graph-rest-1.0
func (m *LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.WorkflowTemplateCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateWorkflowTemplateCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.WorkflowTemplateCollectionResponseable), nil
}
// ToGetRequestInformation get a list of the workflowTemplate objects and their properties.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder when successful
func (m *LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder) {
    return NewLifecycleworkflowsWorkflowtemplatesWorkflowTemplatesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
