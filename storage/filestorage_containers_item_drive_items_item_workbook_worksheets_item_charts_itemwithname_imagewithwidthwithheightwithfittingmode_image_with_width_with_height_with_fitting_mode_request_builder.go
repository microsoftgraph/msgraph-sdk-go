package storage

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder provides operations to call the image method.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, fittingMode *string, height *int32, width *int32)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/item(name='{name}')/image(width={width},height={height},fittingMode='{fittingMode}')", pathParameters),
    }
    if fittingMode != nil {
        m.BaseRequestBuilder.PathParameters["fittingMode"] = *fittingMode
    }
    if height != nil {
        m.BaseRequestBuilder.PathParameters["height"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*height), 10)
    }
    if width != nil {
        m.BaseRequestBuilder.PathParameters["width"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*width), 10)
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Get invoke function image
// Deprecated: This method is obsolete. Use GetAsImageWithWidthWithHeightWithFittingModeGetResponse instead.
// returns a FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration)(FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeResponseable), nil
}
// GetAsImageWithWidthWithHeightWithFittingModeGetResponse invoke function image
// returns a FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder) GetAsImageWithWidthWithHeightWithFittingModeGetResponse(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration)(FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeGetResponseable), nil
}
// ToGetRequestInformation invoke function image
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemwithnameImagewithwidthwithheightwithfittingmodeImageWithWidthWithHeightWithFittingModeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
