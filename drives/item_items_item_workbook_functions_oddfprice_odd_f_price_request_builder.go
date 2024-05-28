package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder provides operations to call the oddFPrice method.
type ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilderInternal instantiates a new ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions/oddFPrice", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder instantiates a new ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action oddFPrice
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookFunctionsOddfpriceOddFPricePostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookFunctionResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable), nil
}
// ToPostRequestInformation invoke action oddFPrice
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookFunctionsOddfpriceOddFPricePostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder when successful
func (m *ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsOddfpriceOddFPriceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
