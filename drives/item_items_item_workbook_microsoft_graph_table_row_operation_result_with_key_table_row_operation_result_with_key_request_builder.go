package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilder provides operations to call the tableRowOperationResult method.
type ItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilderInternal instantiates a new TableRowOperationResultWithKeyRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, key *string)(*ItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilder) {
    m := &ItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/workbook/microsoft.graph.tableRowOperationResult(key='{key}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if key != nil {
        urlTplParams["key"] = *key
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilder instantiates a new TableRowOperationResultWithKeyRequestBuilder and sets the default values.
func NewItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get invoke function tableRowOperationResult
func (m *ItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableRowable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkbookTableRowFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookTableRowable), nil
}
// ToGetRequestInformation invoke function tableRowOperationResult
func (m *ItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemWorkbookMicrosoftGraphTableRowOperationResultWithKeyTableRowOperationResultWithKeyRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
