package logteleconferencedevicequality

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
)

// LogTeleconferenceDeviceQualityRequestBuilder provides operations to call the logTeleconferenceDeviceQuality method.
type LogTeleconferenceDeviceQualityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// LogTeleconferenceDeviceQualityRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LogTeleconferenceDeviceQualityRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLogTeleconferenceDeviceQualityRequestBuilderInternal instantiates a new LogTeleconferenceDeviceQualityRequestBuilder and sets the default values.
func NewLogTeleconferenceDeviceQualityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LogTeleconferenceDeviceQualityRequestBuilder) {
    m := &LogTeleconferenceDeviceQualityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/microsoft.graph.logTeleconferenceDeviceQuality";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewLogTeleconferenceDeviceQualityRequestBuilder instantiates a new LogTeleconferenceDeviceQualityRequestBuilder and sets the default values.
func NewLogTeleconferenceDeviceQualityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LogTeleconferenceDeviceQualityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLogTeleconferenceDeviceQualityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action logTeleconferenceDeviceQuality
func (m *LogTeleconferenceDeviceQualityRequestBuilder) CreatePostRequestInformation(body LogTeleconferenceDeviceQualityRequestBodyable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePostRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePostRequestInformationWithRequestConfiguration invoke action logTeleconferenceDeviceQuality
func (m *LogTeleconferenceDeviceQualityRequestBuilder) CreatePostRequestInformationWithRequestConfiguration(body LogTeleconferenceDeviceQualityRequestBodyable, requestConfiguration *LogTeleconferenceDeviceQualityRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post invoke action logTeleconferenceDeviceQuality
func (m *LogTeleconferenceDeviceQualityRequestBuilder) Post(body LogTeleconferenceDeviceQualityRequestBodyable)(error) {
    return m.PostWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PostWithRequestConfigurationAndResponseHandler invoke action logTeleconferenceDeviceQuality
func (m *LogTeleconferenceDeviceQualityRequestBuilder) PostWithRequestConfigurationAndResponseHandler(body LogTeleconferenceDeviceQualityRequestBodyable, requestConfiguration *LogTeleconferenceDeviceQualityRequestBuilderPostRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePostRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, nil)
    if err != nil {
        return err
    }
    return nil
}
