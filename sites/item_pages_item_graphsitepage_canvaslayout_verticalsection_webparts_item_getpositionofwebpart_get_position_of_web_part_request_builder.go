package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder provides operations to call the getPositionOfWebPart method.
type ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderInternal instantiates a new ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) {
    m := &ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/verticalSection/webparts/{webPart%2Did}/getPositionOfWebPart", pathParameters),
    }
    return m
}
// NewItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder instantiates a new ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action getPositionOfWebPart
// returns a WebPartPositionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartPositionable, error) {
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
func (m *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) WithUrl(rawUrl string)(*ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) {
    return NewItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
