package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder provides operations to call the sharedWithMe method.
type FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetQueryParameters get a list of driveItem objects shared with the owner of a drive. The driveItems returned from the sharedWithMe method always include the remoteItem facet that indicates they are items from a different drive.
type FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetQueryParameters struct {
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
// FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetQueryParameters
}
// NewFilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilderInternal instantiates a new FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder) {
    m := &FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/sharedWithMe(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder instantiates a new FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a list of driveItem objects shared with the owner of a drive. The driveItems returned from the sharedWithMe method always include the remoteItem facet that indicates they are items from a different drive.
// Deprecated: This method is obsolete. Use GetAsSharedWithMeGetResponse instead.
// returns a FilestorageContainersItemDriveSharedwithmeSharedWithMeResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/drive-sharedwithme?view=graph-rest-1.0
func (m *FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration)(FilestorageContainersItemDriveSharedwithmeSharedWithMeResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveSharedwithmeSharedWithMeResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveSharedwithmeSharedWithMeResponseable), nil
}
// GetAsSharedWithMeGetResponse get a list of driveItem objects shared with the owner of a drive. The driveItems returned from the sharedWithMe method always include the remoteItem facet that indicates they are items from a different drive.
// returns a FilestorageContainersItemDriveSharedwithmeSharedWithMeGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/drive-sharedwithme?view=graph-rest-1.0
func (m *FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder) GetAsSharedWithMeGetResponse(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration)(FilestorageContainersItemDriveSharedwithmeSharedWithMeGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveSharedwithmeSharedWithMeGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveSharedwithmeSharedWithMeGetResponseable), nil
}
// ToGetRequestInformation get a list of driveItem objects shared with the owner of a drive. The driveItems returned from the sharedWithMe method always include the remoteItem facet that indicates they are items from a different drive.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder when successful
func (m *FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder) {
    return NewFilestorageContainersItemDriveSharedwithmeSharedWithMeRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
