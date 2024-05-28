package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder provides operations to manage the deviceStatuses property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
type MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetQueryParameters list properties and relationships of the managedDeviceMobileAppConfigurationDeviceStatus objects.
type MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetQueryParameters struct {
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
// MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetQueryParameters
}
// MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByManagedDeviceMobileAppConfigurationDeviceStatusId provides operations to manage the deviceStatuses property of the microsoft.graph.managedDeviceMobileAppConfiguration entity.
// returns a *MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder when successful
func (m *MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) ByManagedDeviceMobileAppConfigurationDeviceStatusId(managedDeviceMobileAppConfigurationDeviceStatusId string)(*MobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if managedDeviceMobileAppConfigurationDeviceStatusId != "" {
        urlTplParams["managedDeviceMobileAppConfigurationDeviceStatus%2Did"] = managedDeviceMobileAppConfigurationDeviceStatusId
    }
    return NewMobileappconfigurationsItemDevicestatusesManagedDeviceMobileAppConfigurationDeviceStatusItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewMobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderInternal instantiates a new MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder and sets the default values.
func NewMobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) {
    m := &MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/mobileAppConfigurations/{managedDeviceMobileAppConfiguration%2Did}/deviceStatuses{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewMobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder instantiates a new MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder and sets the default values.
func NewMobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *MobileappconfigurationsItemDevicestatusesCountRequestBuilder when successful
func (m *MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) Count()(*MobileappconfigurationsItemDevicestatusesCountRequestBuilder) {
    return NewMobileappconfigurationsItemDevicestatusesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list properties and relationships of the managedDeviceMobileAppConfigurationDeviceStatus objects.
// returns a ManagedDeviceMobileAppConfigurationDeviceStatusCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-apps-manageddevicemobileappconfigurationdevicestatus-list?view=graph-rest-1.0
func (m *MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) Get(ctx context.Context, requestConfiguration *MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationDeviceStatusCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceMobileAppConfigurationDeviceStatusCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationDeviceStatusCollectionResponseable), nil
}
// Post create a new managedDeviceMobileAppConfigurationDeviceStatus object.
// returns a ManagedDeviceMobileAppConfigurationDeviceStatusable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-apps-manageddevicemobileappconfigurationdevicestatus-create?view=graph-rest-1.0
func (m *MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationDeviceStatusable, requestConfiguration *MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationDeviceStatusable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateManagedDeviceMobileAppConfigurationDeviceStatusFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationDeviceStatusable), nil
}
// ToGetRequestInformation list properties and relationships of the managedDeviceMobileAppConfigurationDeviceStatus objects.
// returns a *RequestInformation when successful
func (m *MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new managedDeviceMobileAppConfigurationDeviceStatus object.
// returns a *RequestInformation when successful
func (m *MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ManagedDeviceMobileAppConfigurationDeviceStatusable, requestConfiguration *MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder when successful
func (m *MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) WithUrl(rawUrl string)(*MobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder) {
    return NewMobileappconfigurationsItemDevicestatusesDeviceStatusesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
