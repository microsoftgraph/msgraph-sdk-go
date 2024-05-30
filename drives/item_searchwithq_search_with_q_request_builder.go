package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSearchwithqSearchWithQRequestBuilder provides operations to call the search method.
type ItemSearchwithqSearchWithQRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSearchwithqSearchWithQRequestBuilderGetQueryParameters search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
type ItemSearchwithqSearchWithQRequestBuilderGetQueryParameters struct {
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
// ItemSearchwithqSearchWithQRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSearchwithqSearchWithQRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSearchwithqSearchWithQRequestBuilderGetQueryParameters
}
// NewItemSearchwithqSearchWithQRequestBuilderInternal instantiates a new ItemSearchwithqSearchWithQRequestBuilder and sets the default values.
func NewItemSearchwithqSearchWithQRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, q *string)(*ItemSearchwithqSearchWithQRequestBuilder) {
    m := &ItemSearchwithqSearchWithQRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/search(q='{q}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if q != nil {
        m.BaseRequestBuilder.PathParameters["q"] = *q
    }
    return m
}
// NewItemSearchwithqSearchWithQRequestBuilder instantiates a new ItemSearchwithqSearchWithQRequestBuilder and sets the default values.
func NewItemSearchwithqSearchWithQRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSearchwithqSearchWithQRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSearchwithqSearchWithQRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
// Deprecated: This method is obsolete. Use GetAsSearchWithQGetResponse instead.
// returns a ItemSearchwithqSearchWithQResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-search?view=graph-rest-1.0
func (m *ItemSearchwithqSearchWithQRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSearchwithqSearchWithQRequestBuilderGetRequestConfiguration)(ItemSearchwithqSearchWithQResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSearchwithqSearchWithQResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSearchwithqSearchWithQResponseable), nil
}
// GetAsSearchWithQGetResponse search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
// returns a ItemSearchwithqSearchWithQGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-search?view=graph-rest-1.0
func (m *ItemSearchwithqSearchWithQRequestBuilder) GetAsSearchWithQGetResponse(ctx context.Context, requestConfiguration *ItemSearchwithqSearchWithQRequestBuilderGetRequestConfiguration)(ItemSearchwithqSearchWithQGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemSearchwithqSearchWithQGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemSearchwithqSearchWithQGetResponseable), nil
}
// ToGetRequestInformation search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
// returns a *RequestInformation when successful
func (m *ItemSearchwithqSearchWithQRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSearchwithqSearchWithQRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSearchwithqSearchWithQRequestBuilder when successful
func (m *ItemSearchwithqSearchWithQRequestBuilder) WithUrl(rawUrl string)(*ItemSearchwithqSearchWithQRequestBuilder) {
    return NewItemSearchwithqSearchWithQRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
