package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder provides operations to call the clearFilters method.
type ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilderInternal instantiates a new ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder) {
    m := &ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/tables/{workbookTable%2Did}/clearFilters", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder instantiates a new ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilderInternal(urlParams, requestAdapter)
}
// Post clears all the filters currently applied on the table.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/table-clearfilters?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation clears all the filters currently applied on the table.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder when successful
func (m *ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder) {
    return NewItemItemsItemWorkbookTablesItemClearfiltersClearFiltersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
