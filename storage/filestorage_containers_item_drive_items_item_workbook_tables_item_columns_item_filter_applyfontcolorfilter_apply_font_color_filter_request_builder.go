package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder provides operations to call the applyFontColorFilter method.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/filter/applyFontColorFilter", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action applyFontColorFilter
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder) Post(ctx context.Context, body FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action applyFontColorFilter
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
