package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder provides operations to call the createUploadSession method.
type FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/createUploadSession", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createUploadSession
// returns a UploadSessionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder) Post(ctx context.Context, body FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UploadSessionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUploadSessionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UploadSessionable), nil
}
// ToPostRequestInformation invoke action createUploadSession
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemCreateuploadsessionCreateUploadSessionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
