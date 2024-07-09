package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder provides operations to manage the activities property of the microsoft.graph.itemActivityStat entity.
type FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderGetQueryParameters exposes the itemActivities represented in this itemActivityStat resource.
type FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderGetQueryParameters
}
// FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderInternal instantiates a new FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder) {
    m := &FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/analytics/itemActivityStats/{itemActivityStat%2Did}/activities/{itemActivity%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder instantiates a new FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder and sets the default values.
func NewFileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property activities for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DriveItem provides operations to manage the driveItem property of the microsoft.graph.itemActivity entity.
// returns a *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder) DriveItem()(*FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get exposes the itemActivities represented in this itemActivityStat resource.
// returns a ItemActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateItemActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityable), nil
}
// Patch update the navigation property activities in storage
// returns a ItemActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityable, requestConfiguration *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateItemActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityable), nil
}
// ToDeleteRequestInformation delete navigation property activities for storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation exposes the itemActivities represented in this itemActivityStat resource.
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property activities in storage
// returns a *RequestInformation when successful
func (m *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityable, requestConfiguration *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder when successful
func (m *FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder) WithUrl(rawUrl string)(*FileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder) {
    return NewFileStorageContainersItemDriveItemsItemAnalyticsItemActivityStatsItemActivitiesItemActivityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
