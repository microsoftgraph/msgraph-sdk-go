package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder provides operations to call the closeSession method.
type ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookClosesessionCloseSessionRequestBuilderInternal instantiates a new ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookClosesessionCloseSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder) {
    m := &ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/closeSession", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder instantiates a new ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookClosesessionCloseSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post use this API to close an existing workbook session.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/workbook-closesession?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation use this API to close an existing workbook session.
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder when successful
func (m *ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder) {
    return NewItemItemsItemWorkbookClosesessionCloseSessionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
