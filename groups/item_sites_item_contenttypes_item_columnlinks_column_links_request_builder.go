package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
type ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderGetQueryParameters the collection of columns that are required by this content type.
type ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderGetQueryParameters struct {
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
// ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderGetQueryParameters
}
// ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByColumnLinkId provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
// returns a *ItemSitesItemContenttypesItemColumnlinksColumnLinkItemRequestBuilder when successful
func (m *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder) ByColumnLinkId(columnLinkId string)(*ItemSitesItemContenttypesItemColumnlinksColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if columnLinkId != "" {
        urlTplParams["columnLink%2Did"] = columnLinkId
    }
    return NewItemSitesItemContenttypesItemColumnlinksColumnLinkItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderInternal instantiates a new ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder and sets the default values.
func NewItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder) {
    m := &ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/contentTypes/{contentType%2Did}/columnLinks{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder instantiates a new ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder and sets the default values.
func NewItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemSitesItemContenttypesItemColumnlinksCountRequestBuilder when successful
func (m *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder) Count()(*ItemSitesItemContenttypesItemColumnlinksCountRequestBuilder) {
    return NewItemSitesItemContenttypesItemColumnlinksCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of columns that are required by this content type.
// returns a ColumnLinkCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnLinkCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateColumnLinkCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnLinkCollectionResponseable), nil
}
// Post create new navigation property to columnLinks for groups
// returns a ColumnLinkable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnLinkable, requestConfiguration *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnLinkable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateColumnLinkFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnLinkable), nil
}
// ToGetRequestInformation the collection of columns that are required by this content type.
// returns a *RequestInformation when successful
func (m *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to columnLinks for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ColumnLinkable, requestConfiguration *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder when successful
func (m *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder) {
    return NewItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
