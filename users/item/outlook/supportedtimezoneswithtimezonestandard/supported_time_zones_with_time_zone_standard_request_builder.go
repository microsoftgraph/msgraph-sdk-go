package supportedtimezoneswithtimezonestandard

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SupportedTimeZonesWithTimeZoneStandardRequestBuilder provides operations to call the supportedTimeZones method.
type SupportedTimeZonesWithTimeZoneStandardRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SupportedTimeZonesWithTimeZoneStandardRequestBuilderGetOptions options for Get
type SupportedTimeZonesWithTimeZoneStandardRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal instantiates a new SupportedTimeZonesWithTimeZoneStandardRequestBuilder and sets the default values.
func NewSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, timeZoneStandard *string)(*SupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    m := &SupportedTimeZonesWithTimeZoneStandardRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/outlook/microsoft.graph.supportedTimeZones(TimeZoneStandard='{TimeZoneStandard}')";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if timeZoneStandard != nil {
        urlTplParams["TimeZoneStandard"] = *timeZoneStandard
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
func (m *SupportedTimeZonesWithTimeZoneStandardRequestBuilder) CreateGetRequestInformation(options *SupportedTimeZonesWithTimeZoneStandardRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get invoke function supportedTimeZones
func (m *SupportedTimeZonesWithTimeZoneStandardRequestBuilder) Get(options *SupportedTimeZonesWithTimeZoneStandardRequestBuilderGetOptions)(SupportedTimeZonesWithTimeZoneStandardResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateSupportedTimeZonesWithTimeZoneStandardResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(SupportedTimeZonesWithTimeZoneStandardResponseable), nil
}
