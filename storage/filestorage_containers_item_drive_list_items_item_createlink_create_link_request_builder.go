package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder provides operations to call the createLink method.
type FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilderInternal instantiates a new FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder) {
    m := &FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/items/{listItem%2Did}/createLink", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder instantiates a new FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invoke action createLink
// returns a Permissionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder) Post(ctx context.Context, body FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Permissionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePermissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Permissionable), nil
}
// ToPostRequestInformation invoke action createLink
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkPostRequestBodyable, requestConfiguration *FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder when successful
func (m *FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder) {
    return NewFilestorageContainersItemDriveListItemsItemCreatelinkCreateLinkRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
