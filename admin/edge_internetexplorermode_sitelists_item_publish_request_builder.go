package admin

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder provides operations to call the publish method.
type EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// EdgeInternetexplorermodeSitelistsItemPublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EdgeInternetexplorermodeSitelistsItemPublishRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewEdgeInternetexplorermodeSitelistsItemPublishRequestBuilderInternal instantiates a new EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder and sets the default values.
func NewEdgeInternetexplorermodeSitelistsItemPublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder) {
    m := &EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/admin/edge/internetExplorerMode/siteLists/{browserSiteList%2Did}/publish", pathParameters),
    }
    return m
}
// NewEdgeInternetexplorermodeSitelistsItemPublishRequestBuilder instantiates a new EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder and sets the default values.
func NewEdgeInternetexplorermodeSitelistsItemPublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEdgeInternetexplorermodeSitelistsItemPublishRequestBuilderInternal(urlParams, requestAdapter)
}
// Post publish the specified browserSiteList for devices to download.
// returns a BrowserSiteListable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder) Post(ctx context.Context, body EdgeInternetexplorermodeSitelistsItemPublishPostRequestBodyable, requestConfiguration *EdgeInternetexplorermodeSitelistsItemPublishRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BrowserSiteListable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation publish the specified browserSiteList for devices to download.
// returns a *RequestInformation when successful
func (m *EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder) ToPostRequestInformation(ctx context.Context, body EdgeInternetexplorermodeSitelistsItemPublishPostRequestBodyable, requestConfiguration *EdgeInternetexplorermodeSitelistsItemPublishRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder when successful
func (m *EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder) WithUrl(rawUrl string)(*EdgeInternetexplorermodeSitelistsItemPublishRequestBuilder) {
    return NewEdgeInternetexplorermodeSitelistsItemPublishRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
