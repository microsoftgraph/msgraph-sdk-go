package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder provides operations to manage the activities property of the microsoft.graph.itemActivityStat entity.
type FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderGetQueryParameters exposes the itemActivities represented in this itemActivityStat resource.
type FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/items/{driveItem%2Did}/analytics/itemActivityStats/{itemActivityStat%2Did}/activities/{itemActivity%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property activities for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder) DriveItem()(*FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get exposes the itemActivities represented in this itemActivityStat resource.
// returns a ItemActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityable, error) {
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
func (m *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityable, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityable, error) {
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
func (m *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityable, requestConfiguration *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemAnalyticsItemactivitystatsItemActivitiesItemActivityItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
