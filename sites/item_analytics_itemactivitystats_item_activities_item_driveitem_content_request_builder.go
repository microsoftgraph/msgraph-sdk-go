package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder provides operations to manage the media for the site entity.
type ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderGetQueryParameters the content stream, if the item represents a file.
type ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderGetQueryParameters struct {
    // Format of the content
    Format *string `uriparametername:"%24format"`
}
// ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderGetQueryParameters
}
// ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderPutRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderPutRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderInternal instantiates a new ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder and sets the default values.
func NewItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder) {
    m := &ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/analytics/itemActivityStats/{itemActivityStat%2Did}/activities/{itemActivity%2Did}/driveItem/content{?%24format*}", pathParameters),
    }
    return m
}
// NewItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder instantiates a new ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder and sets the default values.
func NewItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the content stream, if the item represents a file.
// returns a []byte when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderGetRequestConfiguration)([]byte, error) {
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
// Put the content stream, if the item represents a file.
// returns a DriveItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder) Put(ctx context.Context, body []byte, requestConfiguration *ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderPutRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DriveItemable, error) {
    requestInfo, err := m.ToPutRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the content stream, if the item represents a file.
// returns a *RequestInformation when successful
func (m *ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/octet-stream, application/json")
    return requestInfo, nil
}
// ToPutRequestInformation the content stream, if the item represents a file.
// returns a *RequestInformation when successful
func (m *ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder) ToPutRequestInformation(ctx context.Context, body []byte, requestConfiguration *ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilderPutRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder when successful
func (m *ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder) WithUrl(rawUrl string)(*ItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder) {
    return NewItemAnalyticsItemactivitystatsItemActivitiesItemDriveitemContentRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
