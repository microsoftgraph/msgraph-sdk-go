package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
type EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdgeInternetexplorermodeSitelistsItemSitesRequestBuilderGetQueryParameters get a list of the browserSite objects and their properties.
type EdgeInternetexplorermodeSitelistsItemSitesRequestBuilderGetQueryParameters struct {
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
// EdgeInternetexplorermodeSitelistsItemSitesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetexplorermodeSitelistsItemSitesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilderGetQueryParameters
}
// EdgeInternetexplorermodeSitelistsItemSitesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetexplorermodeSitelistsItemSitesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByBrowserSiteId provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
// returns a *EdgeInternetexplorermodeSitelistsItemSitesBrowserSiteItemRequestBuilder when successful
func (m *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder) ByBrowserSiteId(browserSiteId string)(*EdgeInternetexplorermodeSitelistsItemSitesBrowserSiteItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if browserSiteId != "" {
        urlTplParams["browserSite%2Did"] = browserSiteId
    }
    return NewEdgeInternetexplorermodeSitelistsItemSitesBrowserSiteItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEdgeInternetexplorermodeSitelistsItemSitesRequestBuilderInternal instantiates a new EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder and sets the default values.
func NewEdgeInternetexplorermodeSitelistsItemSitesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder) {
    m := &EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/edge/internetExplorerMode/siteLists/{browserSiteList%2Did}/sites{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewEdgeInternetexplorermodeSitelistsItemSitesRequestBuilder instantiates a new EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder and sets the default values.
func NewEdgeInternetexplorermodeSitelistsItemSitesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdgeInternetexplorermodeSitelistsItemSitesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *EdgeInternetexplorermodeSitelistsItemSitesCountRequestBuilder when successful
func (m *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder) Count()(*EdgeInternetexplorermodeSitelistsItemSitesCountRequestBuilder) {
    return NewEdgeInternetexplorermodeSitelistsItemSitesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the browserSite objects and their properties.
// returns a BrowserSiteCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/browsersitelist-list-sites?view=graph-rest-1.0
func (m *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder) Get(ctx context.Context, requestConfiguration *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSiteCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBrowserSiteCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSiteCollectionResponseable), nil
}
// Post create a new browserSite object in a browserSiteList.
// returns a BrowserSiteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/browsersitelist-post-sites?view=graph-rest-1.0
func (m *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSiteable, requestConfiguration *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSiteable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBrowserSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSiteable), nil
}
// ToGetRequestInformation get a list of the browserSite objects and their properties.
// returns a *RequestInformation when successful
func (m *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new browserSite object in a browserSiteList.
// returns a *RequestInformation when successful
func (m *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSiteable, requestConfiguration *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder when successful
func (m *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder) WithUrl(rawUrl string)(*EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder) {
    return NewEdgeInternetexplorermodeSitelistsItemSitesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
