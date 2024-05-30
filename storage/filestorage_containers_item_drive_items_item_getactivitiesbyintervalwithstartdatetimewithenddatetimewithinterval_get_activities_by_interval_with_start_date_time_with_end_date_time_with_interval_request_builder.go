package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder provides operations to call the getActivitiesByInterval method.
type FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetQueryParameters invoke function getActivitiesByInterval
type FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetQueryParameters struct {
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
// FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetQueryParameters
}
// NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal instantiates a new FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, endDateTime *string, interval *string, startDateTime *string)(*FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    m := &FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder{
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
// NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder instantiates a new FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderInternal(urlParams, requestAdapter, nil, nil, nil)
}
// Get invoke function getActivitiesByInterval
// Deprecated: This method is obsolete. Use GetAsGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse instead.
// returns a FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration)(FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalResponseable), nil
}
// GetAsGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse invoke function getActivitiesByInterval
// returns a FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) GetAsGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponse(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration)(FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalGetResponseable), nil
}
// ToGetRequestInformation invoke function getActivitiesByInterval
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder when successful
func (m *FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder) {
    return NewFilestorageContainersItemDriveItemsItemGetactivitiesbyintervalwithstartdatetimewithenddatetimewithintervalGetActivitiesByIntervalWithStartDateTimeWithEndDateTimeWithIntervalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
