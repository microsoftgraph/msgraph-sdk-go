package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveSearchWithQRequestBuilder provides operations to call the search method.
type FileStorageContainersItemDriveSearchWithQRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveSearchWithQRequestBuilderGetQueryParameters search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
type FileStorageContainersItemDriveSearchWithQRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// FileStorageContainersItemDriveSearchWithQRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveSearchWithQRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemDriveSearchWithQRequestBuilderGetQueryParameters
}
// NewFileStorageContainersItemDriveSearchWithQRequestBuilderInternal instantiates a new FileStorageContainersItemDriveSearchWithQRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveSearchWithQRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, q *string)(*FileStorageContainersItemDriveSearchWithQRequestBuilder) {
    m := &FileStorageContainersItemDriveSearchWithQRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/search(q='{q}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if q != nil {
        m.BaseRequestBuilder.PathParameters["q"] = *q
    }
    return m
}
// NewFileStorageContainersItemDriveSearchWithQRequestBuilder instantiates a new FileStorageContainersItemDriveSearchWithQRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveSearchWithQRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveSearchWithQRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveSearchWithQRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Get search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
// Deprecated: This method is obsolete. Use GetAsSearchWithQGetResponse instead.
// returns a FileStorageContainersItemDriveSearchWithQResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-search?view=graph-rest-1.0
func (m *FileStorageContainersItemDriveSearchWithQRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveSearchWithQRequestBuilderGetRequestConfiguration)(FileStorageContainersItemDriveSearchWithQResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageContainersItemDriveSearchWithQResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageContainersItemDriveSearchWithQResponseable), nil
}
// GetAsSearchWithQGetResponse search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
// returns a FileStorageContainersItemDriveSearchWithQGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/driveitem-search?view=graph-rest-1.0
func (m *FileStorageContainersItemDriveSearchWithQRequestBuilder) GetAsSearchWithQGetResponse(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveSearchWithQRequestBuilderGetRequestConfiguration)(FileStorageContainersItemDriveSearchWithQGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageContainersItemDriveSearchWithQGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageContainersItemDriveSearchWithQGetResponseable), nil
}
// ToGetRequestInformation search the hierarchy of items for items matching a query.You can search within a folder hierarchy, a whole drive, or files shared with the current user.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveSearchWithQRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveSearchWithQRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FileStorageContainersItemDriveSearchWithQRequestBuilder when successful
func (m *FileStorageContainersItemDriveSearchWithQRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveSearchWithQRequestBuilder) {
    return NewFileStorageContainersItemDriveSearchWithQRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
