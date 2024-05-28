package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder provides operations to call the erfC_Precise method.
type ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilderInternal instantiates a new ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder) {
    m := &ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/functions/erfC_Precise", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder instantiates a new ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action erfC_Precise
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder) Post(ctx context.Context, body ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PrecisePostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable, error) {
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
// ToPostRequestInformation invoke action erfC_Precise
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PrecisePostRequestBodyable, requestConfiguration *ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder when successful
func (m *ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder) {
    return NewItemItemsItemWorkbookFunctionsErfc_preciseErfC_PreciseRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
