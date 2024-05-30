package drives

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder provides operations to call the image method.
type ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, width *int32)(*ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/image(width={width})", pathParameters),
    }
    if width != nil {
        m.BaseRequestBuilder.PathParameters["width"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*width), 10)
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function image
// Deprecated: This method is obsolete. Use GetAsImageWithWidthGetResponse instead.
// returns a ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilderGetRequestConfiguration)(ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthResponseable), nil
}
// GetAsImageWithWidthGetResponse invoke function image
// returns a ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder) GetAsImageWithWidthGetResponse(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilderGetRequestConfiguration)(ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthGetResponseable), nil
}
// ToGetRequestInformation invoke function image
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthImageWithWidthRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
