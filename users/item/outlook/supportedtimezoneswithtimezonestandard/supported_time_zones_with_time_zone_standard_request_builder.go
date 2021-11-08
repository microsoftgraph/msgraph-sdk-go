package supportedtimezoneswithtimezonestandard

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Builds and executes requests for operations under \users\{user-id}\outlook\microsoft.graph.supportedTimeZones(TimeZoneStandard={TimeZoneStandard})
type SupportedTimeZonesWithTimeZoneStandardRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type SupportedTimeZonesWithTimeZoneStandardRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new SupportedTimeZonesWithTimeZoneStandardRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
//  - TimeZoneStandard : Usage: TimeZoneStandard={TimeZoneStandard}
func NewSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter, timeZoneStandard *string)(*SupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    m := &SupportedTimeZonesWithTimeZoneStandardRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/outlook/microsoft.graph.supportedTimeZones(TimeZoneStandard={TimeZoneStandard})";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    if timeZoneStandard != nil {
        urlTplParams["TimeZoneStandard"] = *timeZoneStandard
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SupportedTimeZonesWithTimeZoneStandardRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSupportedTimeZonesWithTimeZoneStandardRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Invoke function supportedTimeZones
// Parameters:
//  - options : Options for the request
func (m *SupportedTimeZonesWithTimeZoneStandardRequestBuilder) CreateGetRequestInformation(options *SupportedTimeZonesWithTimeZoneStandardRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
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
// Invoke function supportedTimeZones
// Parameters:
//  - options : Options for the request
func (m *SupportedTimeZonesWithTimeZoneStandardRequestBuilder) Get(options *SupportedTimeZonesWithTimeZoneStandardRequestBuilderGetOptions)([]SupportedTimeZonesWithTimeZoneStandard, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendCollectionAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSupportedTimeZonesWithTimeZoneStandard() }, nil)
    if err != nil {
        return nil, err
    }
    val := make([]SupportedTimeZonesWithTimeZoneStandard, len(res))
    for i, v := range res {
        val[i] = *(v.(*SupportedTimeZonesWithTimeZoneStandard))
    }
    return val, nil
}
