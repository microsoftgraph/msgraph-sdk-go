package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder provides operations to call the clear method.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/axes/categoryAxis/minorGridlines/format/line/clear", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clear the line format of a chart element.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/chartlineformat-clear?view=graph-rest-1.0
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder) Post(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilderPostRequestConfiguration)(error) {
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
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemAxesCategoryaxisMinorgridlinesFormatLineClearRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
