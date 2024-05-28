package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemPermissionsItemGrantRequestBuilder provides operations to call the grant method.
type FilestorageContainersItemPermissionsItemGrantRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemPermissionsItemGrantRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemPermissionsItemGrantRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemPermissionsItemGrantRequestBuilderInternal instantiates a new FilestorageContainersItemPermissionsItemGrantRequestBuilder and sets the default values.
func NewFilestorageContainersItemPermissionsItemGrantRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemPermissionsItemGrantRequestBuilder) {
    m := &FilestorageContainersItemPermissionsItemGrantRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/permissions/{permission%2Did}/grant", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemPermissionsItemGrantRequestBuilder instantiates a new FilestorageContainersItemPermissionsItemGrantRequestBuilder and sets the default values.
func NewFilestorageContainersItemPermissionsItemGrantRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemPermissionsItemGrantRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemPermissionsItemGrantRequestBuilderInternal(urlParams, requestAdapter)
}
// Post grant users access to a link represented by a permission.
// Deprecated: This method is obsolete. Use PostAsGrantPostResponse instead.
// returns a FilestorageContainersItemPermissionsItemGrantResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permission-grant?view=graph-rest-1.0
func (m *FilestorageContainersItemPermissionsItemGrantRequestBuilder) Post(ctx context.Context, body FilestorageContainersItemPermissionsItemGrantPostRequestBodyable, requestConfiguration *FilestorageContainersItemPermissionsItemGrantRequestBuilderPostRequestConfiguration)(FilestorageContainersItemPermissionsItemGrantResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemPermissionsItemGrantResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemPermissionsItemGrantResponseable), nil
}
// PostAsGrantPostResponse grant users access to a link represented by a permission.
// returns a FilestorageContainersItemPermissionsItemGrantPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/permission-grant?view=graph-rest-1.0
func (m *FilestorageContainersItemPermissionsItemGrantRequestBuilder) PostAsGrantPostResponse(ctx context.Context, body FilestorageContainersItemPermissionsItemGrantPostRequestBodyable, requestConfiguration *FilestorageContainersItemPermissionsItemGrantRequestBuilderPostRequestConfiguration)(FilestorageContainersItemPermissionsItemGrantPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemPermissionsItemGrantPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemPermissionsItemGrantPostResponseable), nil
}
// ToPostRequestInformation grant users access to a link represented by a permission.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemPermissionsItemGrantRequestBuilder) ToPostRequestInformation(ctx context.Context, body FilestorageContainersItemPermissionsItemGrantPostRequestBodyable, requestConfiguration *FilestorageContainersItemPermissionsItemGrantRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemPermissionsItemGrantRequestBuilder when successful
func (m *FilestorageContainersItemPermissionsItemGrantRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemPermissionsItemGrantRequestBuilder) {
    return NewFilestorageContainersItemPermissionsItemGrantRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
