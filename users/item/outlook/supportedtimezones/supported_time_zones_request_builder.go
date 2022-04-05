package supportedtimezones

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// SupportedTimeZonesRequestBuilder provides operations to call the supportedTimeZones method.
type SupportedTimeZonesRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SupportedTimeZonesRequestBuilderGetOptions options for Get
type SupportedTimeZonesRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// NewSupportedTimeZonesRequestBuilderInternal instantiates a new SupportedTimeZonesRequestBuilder and sets the default values.
func NewSupportedTimeZonesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SupportedTimeZonesRequestBuilder) {
    m := &SupportedTimeZonesRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/outlook/microsoft.graph.supportedTimeZones()";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSupportedTimeZonesRequestBuilder instantiates a new SupportedTimeZonesRequestBuilder and sets the default values.
func NewSupportedTimeZonesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SupportedTimeZonesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSupportedTimeZonesRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation invoke function supportedTimeZones
func (m *SupportedTimeZonesRequestBuilder) CreateGetRequestInformation(options *SupportedTimeZonesRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *SupportedTimeZonesRequestBuilder) Get(options *SupportedTimeZonesRequestBuilderGetOptions)(SupportedTimeZonesResponseable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, CreateSupportedTimeZonesResponseFromDiscriminatorValue, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(SupportedTimeZonesResponseable), nil
}
