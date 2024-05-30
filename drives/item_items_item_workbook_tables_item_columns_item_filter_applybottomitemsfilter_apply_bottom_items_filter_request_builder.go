package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder provides operations to call the applyBottomItemsFilter method.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilderInternal instantiates a new ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder) {
    m := &ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/columns/{workbookTableColumn%2Did}/filter/applyBottomItemsFilter", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder instantiates a new ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action applyBottomItemsFilter
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation invoke action applyBottomItemsFilter
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterPostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemColumnsItemFilterApplybottomitemsfilterApplyBottomItemsFilterRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
