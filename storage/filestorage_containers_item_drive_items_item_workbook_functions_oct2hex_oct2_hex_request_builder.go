package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder provides operations to call the oct2Hex method.
type FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/functions/oct2Hex", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action oct2Hex
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder) Post(ctx context.Context, body FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable, error) {
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
// ToPostRequestInformation invoke action oct2Hex
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsOct2hexOct2HexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
