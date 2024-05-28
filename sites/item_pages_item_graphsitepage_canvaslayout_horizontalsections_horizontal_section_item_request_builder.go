package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder provides operations to manage the horizontalSections property of the microsoft.graph.canvasLayout entity.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderGetQueryParameters collection of horizontal sections on the SharePoint page.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderGetQueryParameters
}
// ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Columns provides operations to manage the columns property of the microsoft.graph.horizontalSection entity.
// returns a *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder) Columns()(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilder) {
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsItemColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderInternal instantiates a new ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder) {
    m := &ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/pages/{baseSitePage%2Did}/graph.sitePage/canvasLayout/horizontalSections/{horizontalSection%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder instantiates a new ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder and sets the default values.
func NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property horizontalSections for sites
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get collection of horizontal sections on the SharePoint page.
// returns a HorizontalSectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHorizontalSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionable), nil
}
// Patch update the navigation property horizontalSections in sites
// returns a HorizontalSectionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionable, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateHorizontalSectionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionable), nil
}
// ToDeleteRequestInformation delete navigation property horizontalSections for sites
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation collection of horizontal sections on the SharePoint page.
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property horizontalSections in sites
// returns a *RequestInformation when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.HorizontalSectionable, requestConfiguration *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder when successful
func (m *ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder) WithUrl(rawUrl string)(*ItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder) {
    return NewItemPagesItemGraphsitepageCanvaslayoutHorizontalsectionsHorizontalSectionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
