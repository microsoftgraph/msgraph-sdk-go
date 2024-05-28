package drives

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder provides operations to manage the driveItem property of the microsoft.graph.itemActivity entity.
type ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetQueryParameters exposes the driveItem that was the target of this activity.
type ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetQueryParameters
}
// NewItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderInternal instantiates a new ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder and sets the default values.
func NewItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) {
    m := &ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/drives/{drive%2Did}/items/{driveItem%2Did}/analytics/itemActivityStats/{itemActivityStat%2Did}/activities/{itemActivity%2Did}/driveItem{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder instantiates a new ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder and sets the default values.
func NewItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the drive entity.
// returns a *ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder when successful
func (m *ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) Content()(*ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder) {
    return NewItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get exposes the driveItem that was the target of this activity.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDriveItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable), nil
}
// ToGetRequestInformation exposes the driveItem that was the target of this activity.
// returns a *RequestInformation when successful
func (m *ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder when successful
func (m *ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) WithUrl(rawUrl string)(*ItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder) {
    return NewItemItemsItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemDriveItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
