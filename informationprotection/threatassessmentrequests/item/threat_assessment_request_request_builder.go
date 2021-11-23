package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i830d3ff070fd8ee773495b3c0281bb745b2cb51d722269cf3375470d7bde5582 "github.com/microsoftgraph/msgraph-sdk-go/informationprotection/threatassessmentrequests/item/results"
    ic55cf4dfb7b61ba1f1a933e1461e5ca2dd101263eebcd32b51393ae518540c31 "github.com/microsoftgraph/msgraph-sdk-go/informationprotection/threatassessmentrequests/item/results/item"
)

// ThreatAssessmentRequestRequestBuilder builds and executes requests for operations under \informationProtection\threatAssessmentRequests\{threatAssessmentRequest-id}
type ThreatAssessmentRequestRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ThreatAssessmentRequestRequestBuilderDeleteOptions options for Delete
type ThreatAssessmentRequestRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ThreatAssessmentRequestRequestBuilderGetOptions options for Get
type ThreatAssessmentRequestRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ThreatAssessmentRequestRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ThreatAssessmentRequestRequestBuilderGetQueryParameters get threatAssessmentRequests from informationProtection
type ThreatAssessmentRequestRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// ThreatAssessmentRequestRequestBuilderPatchOptions options for Patch
type ThreatAssessmentRequestRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ThreatAssessmentRequest;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewThreatAssessmentRequestRequestBuilderInternal instantiates a new ThreatAssessmentRequestRequestBuilder and sets the default values.
func NewThreatAssessmentRequestRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ThreatAssessmentRequestRequestBuilder) {
    m := &ThreatAssessmentRequestRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/informationProtection/threatAssessmentRequests/{threatAssessmentRequest_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewThreatAssessmentRequestRequestBuilder instantiates a new ThreatAssessmentRequestRequestBuilder and sets the default values.
func NewThreatAssessmentRequestRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ThreatAssessmentRequestRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatAssessmentRequestRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property threatAssessmentRequests for informationProtection
func (m *ThreatAssessmentRequestRequestBuilder) CreateDeleteRequestInformation(options *ThreatAssessmentRequestRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
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
// CreateGetRequestInformation get threatAssessmentRequests from informationProtection
func (m *ThreatAssessmentRequestRequestBuilder) CreateGetRequestInformation(options *ThreatAssessmentRequestRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property threatAssessmentRequests in informationProtection
func (m *ThreatAssessmentRequestRequestBuilder) CreatePatchRequestInformation(options *ThreatAssessmentRequestRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
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
// Delete delete navigation property threatAssessmentRequests for informationProtection
func (m *ThreatAssessmentRequestRequestBuilder) Delete(options *ThreatAssessmentRequestRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get threatAssessmentRequests from informationProtection
func (m *ThreatAssessmentRequestRequestBuilder) Get(options *ThreatAssessmentRequestRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ThreatAssessmentRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewThreatAssessmentRequest() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ThreatAssessmentRequest), nil
}
// Patch update the navigation property threatAssessmentRequests in informationProtection
func (m *ThreatAssessmentRequestRequestBuilder) Patch(options *ThreatAssessmentRequestRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ThreatAssessmentRequestRequestBuilder) Results()(*i830d3ff070fd8ee773495b3c0281bb745b2cb51d722269cf3375470d7bde5582.ResultsRequestBuilder) {
    return i830d3ff070fd8ee773495b3c0281bb745b2cb51d722269cf3375470d7bde5582.NewResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResultsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.informationProtection.threatAssessmentRequests.item.results.item collection
func (m *ThreatAssessmentRequestRequestBuilder) ResultsById(id string)(*ic55cf4dfb7b61ba1f1a933e1461e5ca2dd101263eebcd32b51393ae518540c31.ThreatAssessmentResultRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["threatAssessmentResult_id"] = id
    }
    return ic55cf4dfb7b61ba1f1a933e1461e5ca2dd101263eebcd32b51393ae518540c31.NewThreatAssessmentResultRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
