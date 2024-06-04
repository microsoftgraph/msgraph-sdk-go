package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.deviceManagement entity.
type DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetQueryParameters list properties and relationships of the deviceEnrollmentLimitConfiguration objects.
type DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetQueryParameters struct {
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
// DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetQueryParameters
}
// DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDeviceEnrollmentConfigurationId provides operations to manage the deviceEnrollmentConfigurations property of the microsoft.graph.deviceManagement entity.
// returns a *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder when successful
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) ByDeviceEnrollmentConfigurationId(deviceEnrollmentConfigurationId string)(*DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if deviceEnrollmentConfigurationId != "" {
        urlTplParams["deviceEnrollmentConfiguration%2Did"] = deviceEnrollmentConfigurationId
    }
    return NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderInternal instantiates a new DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder and sets the default values.
func NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) {
    m := &DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/deviceEnrollmentConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder instantiates a new DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder and sets the default values.
func NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DeviceenrollmentconfigurationsCountRequestBuilder when successful
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) Count()(*DeviceenrollmentconfigurationsCountRequestBuilder) {
    return NewDeviceenrollmentconfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list properties and relationships of the deviceEnrollmentLimitConfiguration objects.
// returns a DeviceEnrollmentConfigurationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-onboarding-deviceenrollmentlimitconfiguration-list?view=graph-rest-1.0
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceEnrollmentConfigurationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceEnrollmentConfigurationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceEnrollmentConfigurationCollectionResponseable), nil
}
// Post create a new deviceEnrollmentLimitConfiguration object.
// returns a DeviceEnrollmentConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-onboarding-deviceenrollmentlimitconfiguration-create?view=graph-rest-1.0
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceEnrollmentConfigurationable, requestConfiguration *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceEnrollmentConfigurationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceEnrollmentConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceEnrollmentConfigurationable), nil
}
// ToGetRequestInformation list properties and relationships of the deviceEnrollmentLimitConfiguration objects.
// returns a *RequestInformation when successful
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new deviceEnrollmentLimitConfiguration object.
// returns a *RequestInformation when successful
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceEnrollmentConfigurationable, requestConfiguration *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder when successful
func (m *DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) WithUrl(rawUrl string)(*DeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder) {
    return NewDeviceenrollmentconfigurationsDeviceEnrollmentConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
