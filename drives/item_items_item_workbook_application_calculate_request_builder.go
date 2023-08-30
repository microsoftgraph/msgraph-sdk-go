package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookApplicationCalculateRequestBuilder provides operations to call the calculate method.
type ItemItemsItemWorkbookApplicationCalculateRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookApplicationCalculateRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookApplicationCalculateRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookApplicationCalculateRequestBuilderInternal instantiates a new CalculateRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookApplicationCalculateRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookApplicationCalculateRequestBuilder) {
    m := &ItemItemsItemWorkbookApplicationCalculateRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/application/calculate", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookApplicationCalculateRequestBuilder instantiates a new CalculateRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookApplicationCalculateRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookApplicationCalculateRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookApplicationCalculateRequestBuilderInternal(urlParams, requestAdapter)
}
// Post recalculate all currently opened workbooks in Excel.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/workbookapplication-calculate?view=graph-rest-1.0
func (m *ItemItemsItemWorkbookApplicationCalculateRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookApplicationCalculatePostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookApplicationCalculateRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation recalculate all currently opened workbooks in Excel.
func (m *ItemItemsItemWorkbookApplicationCalculateRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookApplicationCalculatePostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookApplicationCalculateRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemItemsItemWorkbookApplicationCalculateRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookApplicationCalculateRequestBuilder) {
    return NewItemItemsItemWorkbookApplicationCalculateRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
