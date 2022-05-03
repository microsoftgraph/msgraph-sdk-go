package supportedtimezoneswithtimezonestandard

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SupportedTimeZonesWithTimeZoneStandardRequestBuilder provides operations to call the supportedTimeZones method.
type SupportedTimeZonesWithTimeZoneStandardRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// SupportedTimeZonesWithTimeZoneStandardRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SupportedTimeZonesWithTimeZoneStandardRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal instantiates a new SupportedTimeZonesWithTimeZoneStandardRequestBuilder and sets the default values.
func NewSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, timeZoneStandard *string)(*SupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    m := &SupportedTimeZonesWithTimeZoneStandardRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/outlook/microsoft.graph.supportedTimeZones(TimeZoneStandard='{TimeZoneStandard}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if timeZoneStandard != nil {
        urlTplParams[""] = *timeZoneStandard
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSupportedTimeZonesWithTimeZoneStandardRequestBuilder instantiates a new SupportedTimeZonesWithTimeZoneStandardRequestBuilder and sets the default values.
func NewSupportedTimeZonesWithTimeZoneStandardRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// CreateGetRequestInformation invoke function supportedTimeZones
func (m *SupportedTimeZonesWithTimeZoneStandardRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration invoke function supportedTimeZones
func (m *SupportedTimeZonesWithTimeZoneStandardRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *SupportedTimeZonesWithTimeZoneStandardRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Get invoke function supportedTimeZones
func (m *SupportedTimeZonesWithTimeZoneStandardRequestBuilder) Get()(SupportedTimeZonesWithTimeZoneStandardResponseable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler invoke function supportedTimeZones
func (m *SupportedTimeZonesWithTimeZoneStandardRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *SupportedTimeZonesWithTimeZoneStandardRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(SupportedTimeZonesWithTimeZoneStandardResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateSupportedTimeZonesWithTimeZoneStandardResponseFromDiscriminatorValue, responseHandler, nil)
    if err != nil {
        return nil, err
    }
    return res.(SupportedTimeZonesWithTimeZoneStandardResponseable), nil
}
