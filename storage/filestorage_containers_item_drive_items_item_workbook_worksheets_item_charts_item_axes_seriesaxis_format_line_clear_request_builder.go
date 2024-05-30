package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder provides operations to call the clear method.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/axes/seriesAxis/format/line/clear", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clear the line format of a chart element.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chartlineformat-clear?view=graph-rest-1.0
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder) Post(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation clear the line format of a chart element.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesSeriesaxisFormatLineClearRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
