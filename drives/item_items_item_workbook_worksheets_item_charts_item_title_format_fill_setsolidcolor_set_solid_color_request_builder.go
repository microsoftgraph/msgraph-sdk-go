package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder provides operations to call the setSolidColor method.
type ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/title/format/fill/setSolidColor", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilderInternal(urlParams, requestAdapter)
}
// Post sets the fill formatting of a chart element to a uniform color.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chartfill-setsolidcolor?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation sets the fill formatting of a chart element to a uniform color.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItemTitleFormatFillSetsolidcolorSetSolidColorRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
