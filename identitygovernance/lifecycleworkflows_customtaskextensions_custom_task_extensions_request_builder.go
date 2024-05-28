package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430 "github.com/microsoftgraph/msgraph-sdk-go/models/identitygovernance"
)

// LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder provides operations to manage the customTaskExtensions property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
type LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderGetQueryParameters get a list of the customTaskExtension objects and their properties.
type LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderGetQueryParameters struct {
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
// LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderGetQueryParameters
}
// LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByCustomTaskExtensionId provides operations to manage the customTaskExtensions property of the microsoft.graph.identityGovernance.lifecycleWorkflowsContainer entity.
// returns a *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionItemRequestBuilder when successful
func (m *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder) ByCustomTaskExtensionId(customTaskExtensionId string)(*LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if customTaskExtensionId != "" {
        urlTplParams["customTaskExtension%2Did"] = customTaskExtensionId
    }
    return NewLifecycleworkflowsCustomtaskextensionsCustomTaskExtensionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderInternal instantiates a new LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder and sets the default values.
func NewLifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder) {
    m := &LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/lifecycleWorkflows/customTaskExtensions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewLifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder instantiates a new LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder and sets the default values.
func NewLifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *LifecycleworkflowsCustomtaskextensionsCountRequestBuilder when successful
func (m *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder) Count()(*LifecycleworkflowsCustomtaskextensionsCountRequestBuilder) {
    return NewLifecycleworkflowsCustomtaskextensionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the customTaskExtension objects and their properties.
// returns a CustomTaskExtensionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-lifecycleworkflowscontainer-list-customtaskextensions?view=graph-rest-1.0
func (m *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder) Get(ctx context.Context, requestConfiguration *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderGetRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CustomTaskExtensionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateCustomTaskExtensionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CustomTaskExtensionCollectionResponseable), nil
}
// Post create a new customTaskExtension object.
// returns a CustomTaskExtensionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/identitygovernance-lifecycleworkflowscontainer-post-customtaskextensions?view=graph-rest-1.0
func (m *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder) Post(ctx context.Context, body ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CustomTaskExtensionable, requestConfiguration *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderPostRequestConfiguration)(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CustomTaskExtensionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CreateCustomTaskExtensionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CustomTaskExtensionable), nil
}
// ToGetRequestInformation get a list of the customTaskExtension objects and their properties.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new customTaskExtension object.
// returns a *RequestInformation when successful
func (m *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body ibf6ed4fc8e373ed2600905053a507c004671ad1749cb4b6b77078a908490c430.CustomTaskExtensionable, requestConfiguration *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder when successful
func (m *LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder) WithUrl(rawUrl string)(*LifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder) {
    return NewLifecycleworkflowsCustomtaskextensionsCustomTaskExtensionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
