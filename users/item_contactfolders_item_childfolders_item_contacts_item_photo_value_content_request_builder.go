package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder provides operations to manage the media for the user entity.
type ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderInternal instantiates a new ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder and sets the default values.
func NewItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder) {
    m := &ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/contactFolders/{contactFolder%2Did}/childFolders/{contactFolder%2Did1}/contacts/{contact%2Did}/photo/$value", pathParameters),
    }
    return m
}
// NewItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder instantiates a new ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder and sets the default values.
func NewItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete optional contact picture. You can get or set a photo for a contact.
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get optional contact picture. You can get or set a photo for a contact.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderGetRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// Put optional contact picture. You can get or set a photo for a contact.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderPutRequestConfiguration)([]byte, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.SendPrimitive(ctx, requestInfo, "[]byte", errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.([]byte), nil
}
// ToDeleteRequestInformation optional contact picture. You can get or set a photo for a contact.
// returns a *RequestInformation when successful
func (m *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation optional contact picture. You can get or set a photo for a contact.
// returns a *RequestInformation when successful
func (m *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation optional contact picture. You can get or set a photo for a contact.
// returns a *RequestInformation when successful
func (m *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PUT, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    requestInfo.SetStreamContentAndContentType(body, "application/octet-stream")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder when successful
func (m *ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder) WithUrl(rawUrl string)(*ItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder) {
    return NewItemContactfoldersItemChildfoldersItemContactsItemPhotoValueContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
