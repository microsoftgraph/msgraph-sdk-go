package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder provides operations to call the image method.
type ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilderInternal instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder) {
    m := &ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/worksheets/{workbookWorksheet%2Did}/charts/itemAt(index={index})/image()", pathParameters),
    }
    return m
}
// NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder instantiates a new ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilderInternal(urlParams, requestAdapter)
}
// Get invoke function image
// Deprecated: This method is obsolete. Use GetAsImageGetResponse instead.
// returns a ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilderGetRequestConfiguration)(ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageResponseable), nil
}
// GetAsImageGetResponse invoke function image
// returns a ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder) GetAsImageGetResponse(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilderGetRequestConfiguration)(ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageGetResponseable), nil
}
// ToGetRequestInformation invoke function image
// returns a *RequestInformation when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder when successful
func (m *ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder) {
    return NewItemItemsItemWorkbookWorksheetsItemChartsItematwithindexImageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
