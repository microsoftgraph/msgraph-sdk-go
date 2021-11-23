package logteleconferencedevicequality

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
)

// LogTeleconferenceDeviceQualityRequestBuilder builds and executes requests for operations under \communications\calls\microsoft.graph.logTeleconferenceDeviceQuality
type LogTeleconferenceDeviceQualityRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// LogTeleconferenceDeviceQualityRequestBuilderPostOptions options for Post
type LogTeleconferenceDeviceQualityRequestBuilderPostOptions struct {
    // 
    Body *LogTeleconferenceDeviceQualityRequestBody;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewLogTeleconferenceDeviceQualityRequestBuilderInternal instantiates a new LogTeleconferenceDeviceQualityRequestBuilder and sets the default values.
func NewLogTeleconferenceDeviceQualityRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LogTeleconferenceDeviceQualityRequestBuilder) {
    m := &LogTeleconferenceDeviceQualityRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/microsoft.graph.logTeleconferenceDeviceQuality";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewLogTeleconferenceDeviceQualityRequestBuilder instantiates a new LogTeleconferenceDeviceQualityRequestBuilder and sets the default values.
func NewLogTeleconferenceDeviceQualityRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*LogTeleconferenceDeviceQualityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLogTeleconferenceDeviceQualityRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation invoke action logTeleconferenceDeviceQuality
func (m *LogTeleconferenceDeviceQualityRequestBuilder) CreatePostRequestInformation(options *LogTeleconferenceDeviceQualityRequestBuilderPostOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.POST
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Post invoke action logTeleconferenceDeviceQuality
func (m *LogTeleconferenceDeviceQualityRequestBuilder) Post(options *LogTeleconferenceDeviceQualityRequestBuilderPostOptions)(error) {
    requestInfo, err := m.CreatePostRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
