package softwareupdatestatussummary

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i24c64eb0ec38877257404549b298e224afdd6d4611130bf9401230f4ad5cb41d "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/softwareupdatestatussummary/ref"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// Builds and executes requests for operations under \deviceManagement\softwareUpdateStatusSummary
type SoftwareUpdateStatusSummaryRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type SoftwareUpdateStatusSummaryRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SoftwareUpdateStatusSummaryRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The software update status summary.
type SoftwareUpdateStatusSummaryRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new SoftwareUpdateStatusSummaryRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSoftwareUpdateStatusSummaryRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SoftwareUpdateStatusSummaryRequestBuilder) {
    m := &SoftwareUpdateStatusSummaryRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/softwareUpdateStatusSummary{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SoftwareUpdateStatusSummaryRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSoftwareUpdateStatusSummaryRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SoftwareUpdateStatusSummaryRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSoftwareUpdateStatusSummaryRequestBuilderInternal(urlParams, requestAdapter)
}
// The software update status summary.
// Parameters:
//  - options : Options for the request
func (m *SoftwareUpdateStatusSummaryRequestBuilder) CreateGetRequestInformation(options *SoftwareUpdateStatusSummaryRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
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
// The software update status summary.
// Parameters:
//  - options : Options for the request
func (m *SoftwareUpdateStatusSummaryRequestBuilder) Get(options *SoftwareUpdateStatusSummaryRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SoftwareUpdateStatusSummary, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSoftwareUpdateStatusSummary() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SoftwareUpdateStatusSummary), nil
}
func (m *SoftwareUpdateStatusSummaryRequestBuilder) Ref()(*i24c64eb0ec38877257404549b298e224afdd6d4611130bf9401230f4ad5cb41d.RefRequestBuilder) {
    return i24c64eb0ec38877257404549b298e224afdd6d4611130bf9401230f4ad5cb41d.NewRefRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
