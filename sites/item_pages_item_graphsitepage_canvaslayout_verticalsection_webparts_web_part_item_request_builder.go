package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder provides operations to manage the webparts property of the microsoft.graph.verticalSection entity.
type ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderGetQueryParameters the set of web parts in this section.
type ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderGetQueryParameters
}
// ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderInternal instantiates a new ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder) {
    m := &ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/verticalSection/webparts/{webPart%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder instantiates a new ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property webparts for sites
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the set of web parts in this section.
// returns a WebPartable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWebPartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable), nil
}
// GetPositionOfWebPart provides operations to call the getPositionOfWebPart method.
// returns a *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder) GetPositionOfWebPart()(*ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilder) {
    return NewItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsItemGetpositionofwebpartGetPositionOfWebPartRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property webparts in sites
// returns a WebPartable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWebPartFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable), nil
}
// ToDeleteRequestInformation delete navigation property webparts for sites
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the set of web parts in this section.
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property webparts in sites
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WebPartable, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder) WithUrl(rawUrl string)(*ItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder) {
    return NewItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsWebPartItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
