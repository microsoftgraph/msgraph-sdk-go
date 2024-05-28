package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder provides operations to call the unpublish method.
type FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilderInternal instantiates a new FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder) {
    m := &FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/contentTypes/{contentType%2Did}/unpublish", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder instantiates a new FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unpublish a contentType from a content type hub site.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-unpublish?view=graph-rest-1.0
func (m *FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder) Post(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation unpublish a contentType from a content type hub site.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
