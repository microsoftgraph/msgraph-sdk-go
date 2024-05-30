package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder provides operations to manage the verticalSection property of the microsoft.graph.canvasLayout entity.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderGetQueryParameters vertical section on the SharePoint page.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderGetQueryParameters
}
// ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderInternal instantiates a new ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder) {
    m := &ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/verticalSection{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder instantiates a new ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property verticalSection for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get vertical section on the SharePoint page.
// returns a VerticalSectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VerticalSectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateVerticalSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VerticalSectionable), nil
}
// Patch update the navigation property verticalSection in groups
// returns a VerticalSectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VerticalSectionable, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VerticalSectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateVerticalSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VerticalSectionable), nil
}
// ToDeleteRequestInformation delete navigation property verticalSection for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation vertical section on the SharePoint page.
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property verticalSection in groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VerticalSectionable, requestConfiguration *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Webparts provides operations to manage the webparts property of the microsoft.graph.verticalSection entity.
// returns a *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder) Webparts()(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsRequestBuilder) {
    return NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionWebpartsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder when successful
func (m *ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder) {
    return NewItemSitesItemPagesItemGraphsitepageCanvaslayoutVerticalsectionVerticalSectionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
