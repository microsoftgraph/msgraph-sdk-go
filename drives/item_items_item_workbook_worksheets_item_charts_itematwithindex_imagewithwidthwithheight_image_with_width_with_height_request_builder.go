package drives

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder provides operations to call the image method.
type ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, height *int32, width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/itemAt(index={index})/image(width={width},height={height})", pathParameters),
    }
    if height != nil {
        m.BaseRequestBuilder.PathParameters["height"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*height), 10)
    }
    if width != nil {
        m.BaseRequestBuilder.PathParameters["width"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*width), 10)
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function image
// Deprecated: This method is obsolete. Use GetAsImageWithWidthWithHeightGetResponse instead.
// returns a ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderGetRequestConfiguration)(ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightResponseable), nil
}
// GetAsImageWithWidthWithHeightGetResponse invoke function image
// returns a ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) GetAsImageWithWidthWithHeightGetResponse(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderGetRequestConfiguration)(ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightGetResponseable), nil
}
// ToGetRequestInformation invoke function image
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
