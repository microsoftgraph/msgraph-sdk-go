package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder provides operations to call the add method.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/tables/{workbookTable%2Did}/columns/add", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilderInternal(urlParams, requestAdapter)
}
// Post adds a new column to the table.
// returns a WorkbookTableColumnable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/tablecolumncollection-add?view=graph-rest-1.0
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder) Post(ctx context.Context, body FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookTableColumnFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableColumnable), nil
}
// ToPostRequestInformation adds a new column to the table.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemTablesItemColumnsAddRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
