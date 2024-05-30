package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder provides operations to call the dec2Hex method.
type FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/functions/dec2Hex", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action dec2Hex
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder) Post(ctx context.Context, body FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable, error) {
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
// ToPostRequestInformation invoke action dec2Hex
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsDec2hexDec2HexRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
