package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder provides operations to call the getActivitiesByInterval method.
type FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetQueryParameters get a collection of itemActivityStats resources for the activities that took place on this resource within the specified time interval. Analytics aggregates might not be available for all action types.
type FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetQueryParameters struct {
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
// FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetQueryParameters
}
// NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/getActivitiesByInterval(){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a collection of itemActivityStats resources for the activities that took place on this resource within the specified time interval. Analytics aggregates might not be available for all action types.
// Deprecated: This method is obsolete. Use GetAsGetActivitiesByIntervalGetResponse instead.
// returns a FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/itemactivitystat-getactivitybyinterval?view=graph-rest-1.0
func (m *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration)(FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalResponseable), nil
}
// GetAsGetActivitiesByIntervalGetResponse get a collection of itemActivityStats resources for the activities that took place on this resource within the specified time interval. Analytics aggregates might not be available for all action types.
// returns a FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/itemactivitystat-getactivitybyinterval?view=graph-rest-1.0
func (m *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) GetAsGetActivitiesByIntervalGetResponse(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration)(FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalGetResponseable), nil
}
// ToGetRequestInformation get a collection of itemActivityStats resources for the activities that took place on this resource within the specified time interval. Analytics aggregates might not be available for all action types.
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalGetActivitiesByIntervalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
