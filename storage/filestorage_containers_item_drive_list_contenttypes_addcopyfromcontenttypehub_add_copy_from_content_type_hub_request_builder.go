package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder provides operations to call the addCopyFromContentTypeHub method.
type FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderInternal instantiates a new FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) {
    m := &FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/contentTypes/addCopyFromContentTypeHub", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder instantiates a new FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderInternal(urlParams, requestAdapter)
}
// Post add or sync a copy of a published content type from the content type hub to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see contentType: getCompatibleHubContentTypes and the blog post Syntex Product Updates – August 2021.
// returns a ContentTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/contenttype-addcopyfromcontenttypehub?view=graph-rest-1.0
func (m *FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) Post(ctx context.Context, body FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable), nil
}
// ToPostRequestInformation add or sync a copy of a published content type from the content type hub to a target site or a list. This method is part of the content type publishing changes to optimize the syncing of published content types to sites and lists, effectively switching from a 'push everywhere' to 'pull as needed' approach. The method allows users to pull content types directly from the content type hub to a site or list. For more information, see contentType: getCompatibleHubContentTypes and the blog post Syntex Product Updates – August 2021.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesAddcopyfromcontenttypehubAddCopyFromContentTypeHubRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
