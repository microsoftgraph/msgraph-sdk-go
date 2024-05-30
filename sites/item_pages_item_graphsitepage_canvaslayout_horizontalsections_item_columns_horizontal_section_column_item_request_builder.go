package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder provides operations to manage the columns property of the microsoft.graph.horizontalSection entity.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetQueryParameters the set of vertical columns in this section.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetQueryParameters
}
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal instantiates a new ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) {
    m := &ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/horizontalSections/{horizontalSection%2Did}/columns/{horizontalSectionColumn%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder instantiates a new ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property columns for sites
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the set of vertical columns in this section.
// returns a HorizontalSectionColumnable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHorizontalSectionColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable), nil
}
// Patch update the navigation property columns in sites
// returns a HorizontalSectionColumnable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHorizontalSectionColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable), nil
}
// ToDeleteRequestInformation delete navigation property columns for sites
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the set of vertical columns in this section.
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property columns in sites
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionColumnable, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Webparts provides operations to manage the webparts property of the microsoft.graph.horizontalSectionColumn entity.
// returns a *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsItemWebpartsRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) Webparts()(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsItemWebpartsRequestBuilder) {
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsItemWebpartsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) WithUrl(rawUrl string)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder) {
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsHorizontalSectionColumnItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
