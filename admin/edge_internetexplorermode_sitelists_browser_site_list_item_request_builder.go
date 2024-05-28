package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder provides operations to manage the siteLists property of the microsoft.graph.internetExplorerMode entity.
type EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderGetQueryParameters get a browserSiteList that contains browserSite and browserSharedCookie resources.
type EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderGetQueryParameters
}
// EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderInternal instantiates a new EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder and sets the default values.
func NewEdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder) {
    m := &EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/edge/internetExplorerMode/siteLists/{browserSiteList%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder instantiates a new EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder and sets the default values.
func NewEdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a browserSiteList object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/internetexplorermode-delete-sitelists?view=graph-rest-1.0
func (m *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get a browserSiteList that contains browserSite and browserSharedCookie resources.
// returns a BrowserSiteListable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/browsersitelist-get?view=graph-rest-1.0
func (m *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSiteListable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBrowserSiteListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSiteListable), nil
}
// Patch update the properties of a browserSiteList object.
// returns a BrowserSiteListable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/browsersitelist-update?view=graph-rest-1.0
func (m *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSiteListable, requestConfiguration *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSiteListable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateBrowserSiteListFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSiteListable), nil
}
// Publish provides operations to call the publish method.
// returns a *EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder when successful
func (m *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder) Publish()(*EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder) {
    return NewEdgeInternetexplorermodeSitelistsItemPublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SharedCookies provides operations to manage the sharedCookies property of the microsoft.graph.browserSiteList entity.
// returns a *EdgeInternetexplorermodeSitelistsItemSharedcookiesSharedCookiesRequestBuilder when successful
func (m *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder) SharedCookies()(*EdgeInternetexplorermodeSitelistsItemSharedcookiesSharedCookiesRequestBuilder) {
    return NewEdgeInternetexplorermodeSitelistsItemSharedcookiesSharedCookiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sites provides operations to manage the sites property of the microsoft.graph.browserSiteList entity.
// returns a *EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder when successful
func (m *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder) Sites()(*EdgeInternetexplorermodeSitelistsItemSitesRequestBuilder) {
    return NewEdgeInternetexplorermodeSitelistsItemSitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a browserSiteList object.
// returns a *RequestInformation when successful
func (m *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get a browserSiteList that contains browserSite and browserSharedCookie resources.
// returns a *RequestInformation when successful
func (m *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a browserSiteList object.
// returns a *RequestInformation when successful
func (m *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSiteListable, requestConfiguration *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder when successful
func (m *EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder) WithUrl(rawUrl string)(*EdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder) {
    return NewEdgeInternetexplorermodeSitelistsBrowserSiteListItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
