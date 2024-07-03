package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder provides operations to call the getActivitiesByInterval method.
type FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetQueryParameters invoke function getActivitiesByInterval
type FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetQueryParameters struct {
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
// FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetQueryParameters
}
// NewFileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *string, interval *string, startDateTime *string)(*FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/getActivitiesByInterval(startDateTime='{startDateTime}',endDateTime='{endDateTime}',interval='{interval}'){?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    if endDateTime != nil {
        m.BaseRequestBuilder.PathParameters["endDateTime"] = *endDateTime
    }
    if interval != nil {
        m.BaseRequestBuilder.PathParameters["interval"] = *interval
    }
    if startDateTime != nil {
        m.BaseRequestBuilder.PathParameters["startDateTime"] = *startDateTime
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Get invoke function getActivitiesByInterval
// Deprecated: This method is obsolete. Use GetAsGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse instead.
// returns a FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration)(FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable), nil
}
// GetAsGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse invoke function getActivitiesByInterval
// returns a FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) GetAsGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration)(FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable), nil
}
// ToGetRequestInformation invoke function getActivitiesByInterval
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
