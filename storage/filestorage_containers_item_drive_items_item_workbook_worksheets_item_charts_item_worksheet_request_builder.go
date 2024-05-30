package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder provides operations to manage the worksheet property of the microsoft.graph.workbookChart entity.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilderGetQueryParameters the worksheet containing the current chart. Read-only.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilderGetQueryParameters
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/worksheet{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the worksheet containing the current chart. Read-only.
// returns a WorkbookWorksheetable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookWorksheetFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookWorksheetable), nil
}
// ToGetRequestInformation the worksheet containing the current chart. Read-only.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemWorksheetRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
