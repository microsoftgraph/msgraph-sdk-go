package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder provides operations to manage the driveItem property of the microsoft.graph.itemActivity entity.
type ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetQueryParameters exposes the driveItem that was the target of this activity.
type ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetQueryParameters
}
// NewItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) {
    m := &ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/analytics/itemActivityStats/{itemActivityStat%2Did}/activities/{itemActivity%2Did}/driveItem{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder instantiates a new DriveItemRequestBuilder and sets the default values.
func NewItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Content provides operations to manage the media for the site entity.
func (m *ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) Content()(*ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemContentRequestBuilder) {
    return NewItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemContentRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get exposes the driveItem that was the target of this activity.
func (m *ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAnalyticsItemActivityStatsItemActivitiesItemDriveItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
