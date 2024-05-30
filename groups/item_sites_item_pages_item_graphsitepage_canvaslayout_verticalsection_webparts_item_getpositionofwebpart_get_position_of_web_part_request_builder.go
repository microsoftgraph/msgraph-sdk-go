package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder provides operations to call the getPositionOfWebPart method.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderInternal instantiates a new ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) {
    m := &ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/verticalSection/webparts/{webPart%2Did}/getPositionOfWebPart", pathParameters),
    }
    return m
}
// NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder instantiates a new ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getPositionOfWebPart
// returns a WebPartPositionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartPositionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWebPartPositionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartPositionable), nil
}
// ToPostRequestInformation invoke action getPositionOfWebPart
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) {
    return NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
