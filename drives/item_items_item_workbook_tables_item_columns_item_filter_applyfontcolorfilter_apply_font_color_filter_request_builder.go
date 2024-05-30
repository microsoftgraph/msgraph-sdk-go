package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder provides operations to call the applyFontColorFilter method.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderInternal instantiates a new ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder) {
    m := &ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/filter/applyFontColorFilter", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder instantiates a new ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action applyFontColorFilter
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderPostRequestConfiguration)(error) {
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
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplyfontcolorfilterApplyFontColorFilterRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
