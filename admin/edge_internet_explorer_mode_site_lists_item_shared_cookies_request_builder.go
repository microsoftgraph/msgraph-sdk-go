package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder provides operations to manage the sharedCookies property of the microsoft.graph.browserSiteList entity.
type EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderGetQueryParameters get a list of the browserSharedCookie objects and their properties.
type EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderGetQueryParameters struct {
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
// EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderGetQueryParameters
}
// EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByBrowserSharedCookieId provides operations to manage the sharedCookies property of the microsoft.graph.browserSiteList entity.
func (m *EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder) ByBrowserSharedCookieId(browserSharedCookieId string)(*EdgeInternetExplorerModeSiteListsItemSharedCookiesBrowserSharedCookieItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if browserSharedCookieId != "" {
        urlTplParams["browserSharedCookie%2Did"] = browserSharedCookieId
    }
    return NewEdgeInternetExplorerModeSiteListsItemSharedCookiesBrowserSharedCookieItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewEdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderInternal instantiates a new SharedCookiesRequestBuilder and sets the default values.
func NewEdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder) {
    m := &EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/edge/internetExplorerMode/siteLists/{browserSiteList%2Did}/sharedCookies{?%24top,%24skip,%24search,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewEdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder instantiates a new SharedCookiesRequestBuilder and sets the default values.
func NewEdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder) Count()(*EdgeInternetExplorerModeSiteListsItemSharedCookiesCountRequestBuilder) {
    return NewEdgeInternetExplorerModeSiteListsItemSharedCookiesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the browserSharedCookie objects and their properties.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/browsersitelist-list-sharedcookies?view=graph-rest-1.0
func (m *EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder) Get(ctx context.Context, requestConfiguration *EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSharedCookieCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBrowserSharedCookieCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSharedCookieCollectionResponseable), nil
}
// Post create a new browserSharedCookie object in a browserSiteList.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/browsersitelist-post-sharedcookies?view=graph-rest-1.0
func (m *EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSharedCookieable, requestConfiguration *EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSharedCookieable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBrowserSharedCookieFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSharedCookieable), nil
}
// ToGetRequestInformation get a list of the browserSharedCookie objects and their properties.
func (m *EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToPostRequestInformation create a new browserSharedCookie object in a browserSiteList.
func (m *EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSharedCookieable, requestConfiguration *EdgeInternetExplorerModeSiteListsItemSharedCookiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
