package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder provides operations to call the gamma_Dist method.
type FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/workbook/functions/gamma_Dist", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action gamma_Dist
// returns a WorkbookFunctionResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder) Post(ctx context.Context, body FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkbookFunctionResultable, error) {
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
// ToPostRequestInformation invoke action gamma_Dist
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemWorkbookFunctionsGamma_distGamma_DistRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
