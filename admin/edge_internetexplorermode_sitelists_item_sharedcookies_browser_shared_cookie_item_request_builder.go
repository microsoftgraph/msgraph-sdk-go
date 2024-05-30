package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder provides operations to manage the sharedCookies property of the microsoft.graph.browserSiteList entity.
type EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderGetQueryParameters get a session cookie that can be shared between a Microsoft Edge process and an Internet Explorer process, while using Internet Explorer mode.
type EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderGetQueryParameters
}
// EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderInternal instantiates a new EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder and sets the default values.
func NewEdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder) {
    m := &EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/edge/internetExplorerMode/siteLists/{browserSiteList%2Did}/sharedCookies/{browserSharedCookie%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewEdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder instantiates a new EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder and sets the default values.
func NewEdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a browserSharedCookie from a browserSiteList.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/browsersitelist-delete-sharedcookies?view=graph-rest-1.0
func (m *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get a session cookie that can be shared between a Microsoft Edge process and an Internet Explorer process, while using Internet Explorer mode.
// returns a BrowserSharedCookieable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/browsersharedcookie-get?view=graph-rest-1.0
func (m *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder) Get(ctx context.Context, requestConfiguration *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSharedCookieable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// Patch update the properties of a browserSharedCookie object.
// returns a BrowserSharedCookieable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/browsersharedcookie-update?view=graph-rest-1.0
func (m *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSharedCookieable, requestConfiguration *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSharedCookieable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
// ToDeleteRequestInformation delete a browserSharedCookie from a browserSiteList.
// returns a *RequestInformation when successful
func (m *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get a session cookie that can be shared between a Microsoft Edge process and an Internet Explorer process, while using Internet Explorer mode.
// returns a *RequestInformation when successful
func (m *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a browserSharedCookie object.
// returns a *RequestInformation when successful
func (m *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSharedCookieable, requestConfiguration *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder when successful
func (m *EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder) WithUrl(rawUrl string)(*EdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder) {
    return NewEdgeInternetexplorermodeSitelistsItemSharedcookiesBrowserSharedCookieItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
