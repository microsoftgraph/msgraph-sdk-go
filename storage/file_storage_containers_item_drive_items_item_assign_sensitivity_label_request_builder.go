package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder provides operations to call the assignSensitivityLabel method.
type FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/assignSensitivityLabel", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action assignSensitivityLabel
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder) Post(ctx context.Context, body FileStorageContainersItemDriveItemsItemAssignSensitivityLabelPostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToPostRequestInformation invoke action assignSensitivityLabel
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder) ToPostRequestInformation(ctx context.Context, body FileStorageContainersItemDriveItemsItemAssignSensitivityLabelPostRequestBodyable, requestConfiguration *FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemAssignSensitivityLabelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
