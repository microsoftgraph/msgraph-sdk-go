package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder provides operations to call the extractSensitivityLabels method.
type FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/extractSensitivityLabels", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action extractSensitivityLabels
// returns a ExtractSensitivityLabelsResultable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder) Post(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ExtractSensitivityLabelsResultable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateExtractSensitivityLabelsResultFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ExtractSensitivityLabelsResultable), nil
}
// ToPostRequestInformation invoke action extractSensitivityLabels
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemExtractSensitivityLabelsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
