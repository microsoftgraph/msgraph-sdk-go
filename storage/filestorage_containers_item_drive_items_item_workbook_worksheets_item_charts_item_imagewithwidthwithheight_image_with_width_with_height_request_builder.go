package storage

import (
    "context"
    i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274 "strconv"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder provides operations to call the image method.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, height *int32, width *int32)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/{workbookChart%2Did}/image(width={width},height={height})", pathParameters),
    }
    if height != nil {
        m.BaseRequestBuilder.PathParameters["height"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*height), 10)
    }
    if width != nil {
        m.BaseRequestBuilder.PathParameters["width"] = i53ac87e8cb3cc9276228f74d38694a208cacb99bb8ceb705eeae99fb88d4d274.FormatInt(int64(*width), 10)
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderInternal(urlParams, requestAdapter, nil, nil)
}
// Get invoke function image
// Deprecated: This method is obsolete. Use GetAsImageWithWidthWithHeightGetResponse instead.
// returns a FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderGetRequestConfiguration)(FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightResponseable), nil
}
// GetAsImageWithWidthWithHeightGetResponse invoke function image
// returns a FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) GetAsImageWithWidthWithHeightGetResponse(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderGetRequestConfiguration)(FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightGetResponseable), nil
}
// ToGetRequestInformation invoke function image
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookWorksheetsItemChartsItemImagewithwidthwithheightImageWithWidthWithHeightRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
