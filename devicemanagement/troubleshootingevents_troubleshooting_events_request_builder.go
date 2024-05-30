package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TroubleshootingeventsTroubleshootingEventsRequestBuilder provides operations to manage the troubleshootingEvents property of the microsoft.graph.deviceManagement entity.
type TroubleshootingeventsTroubleshootingEventsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TroubleshootingeventsTroubleshootingEventsRequestBuilderGetQueryParameters the list of troubleshooting events for the tenant.
type TroubleshootingeventsTroubleshootingEventsRequestBuilderGetQueryParameters struct {
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
// TroubleshootingeventsTroubleshootingEventsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TroubleshootingeventsTroubleshootingEventsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TroubleshootingeventsTroubleshootingEventsRequestBuilderGetQueryParameters
}
// TroubleshootingeventsTroubleshootingEventsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TroubleshootingeventsTroubleshootingEventsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceManagementTroubleshootingEventId provides operations to manage the troubleshootingEvents property of the microsoft.graph.deviceManagement entity.
// returns a *TroubleshootingeventsDeviceManagementTroubleshootingEventItemRequestBuilder when successful
func (m *TroubleshootingeventsTroubleshootingEventsRequestBuilder) ByDeviceManagementTroubleshootingEventId(deviceManagementTroubleshootingEventId string)(*TroubleshootingeventsDeviceManagementTroubleshootingEventItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceManagementTroubleshootingEventId != "" {
        urlTplParams["deviceManagementTroubleshootingEvent%2Did"] = deviceManagementTroubleshootingEventId
    }
    return NewTroubleshootingeventsDeviceManagementTroubleshootingEventItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTroubleshootingeventsTroubleshootingEventsRequestBuilderInternal instantiates a new TroubleshootingeventsTroubleshootingEventsRequestBuilder and sets the default values.
func NewTroubleshootingeventsTroubleshootingEventsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TroubleshootingeventsTroubleshootingEventsRequestBuilder) {
    m := &TroubleshootingeventsTroubleshootingEventsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/troubleshootingEvents{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTroubleshootingeventsTroubleshootingEventsRequestBuilder instantiates a new TroubleshootingeventsTroubleshootingEventsRequestBuilder and sets the default values.
func NewTroubleshootingeventsTroubleshootingEventsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TroubleshootingeventsTroubleshootingEventsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTroubleshootingeventsTroubleshootingEventsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TroubleshootingeventsCountRequestBuilder when successful
func (m *TroubleshootingeventsTroubleshootingEventsRequestBuilder) Count()(*TroubleshootingeventsCountRequestBuilder) {
    return NewTroubleshootingeventsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of troubleshooting events for the tenant.
// returns a DeviceManagementTroubleshootingEventCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TroubleshootingeventsTroubleshootingEventsRequestBuilder) Get(ctx context.Context, requestConfiguration *TroubleshootingeventsTroubleshootingEventsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementTroubleshootingEventCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceManagementTroubleshootingEventCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementTroubleshootingEventCollectionResponseable), nil
}
// Post create new navigation property to troubleshootingEvents for deviceManagement
// returns a DeviceManagementTroubleshootingEventable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TroubleshootingeventsTroubleshootingEventsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementTroubleshootingEventable, requestConfiguration *TroubleshootingeventsTroubleshootingEventsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementTroubleshootingEventable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceManagementTroubleshootingEventFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementTroubleshootingEventable), nil
}
// ToGetRequestInformation the list of troubleshooting events for the tenant.
// returns a *RequestInformation when successful
func (m *TroubleshootingeventsTroubleshootingEventsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TroubleshootingeventsTroubleshootingEventsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to troubleshootingEvents for deviceManagement
// returns a *RequestInformation when successful
func (m *TroubleshootingeventsTroubleshootingEventsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceManagementTroubleshootingEventable, requestConfiguration *TroubleshootingeventsTroubleshootingEventsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *TroubleshootingeventsTroubleshootingEventsRequestBuilder when successful
func (m *TroubleshootingeventsTroubleshootingEventsRequestBuilder) WithUrl(rawUrl string)(*TroubleshootingeventsTroubleshootingEventsRequestBuilder) {
    return NewTroubleshootingeventsTroubleshootingEventsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
