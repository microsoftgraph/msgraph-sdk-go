package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i830d3ff070fd8ee773495b3c0281bb745b2cb51d722269cf3375470d7bde5582 "github.com/microsoftgraph/msgraph-sdk-go/informationprotection/threatassessmentrequests/item/results"
    ic55cf4dfb7b61ba1f1a933e1461e5ca2dd101263eebcd32b51393ae518540c31 "github.com/microsoftgraph/msgraph-sdk-go/informationprotection/threatassessmentrequests/item/results/item"
)

// ThreatAssessmentRequestItemRequestBuilder builds and executes requests for operations under \informationProtection\threatAssessmentRequests\{threatAssessmentRequest-id}
type ThreatAssessmentRequestItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ThreatAssessmentRequestItemRequestBuilderDeleteOptions options for Delete
type ThreatAssessmentRequestItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ThreatAssessmentRequestItemRequestBuilderGetOptions options for Get
type ThreatAssessmentRequestItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ThreatAssessmentRequestItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ThreatAssessmentRequestItemRequestBuilderGetQueryParameters get threatAssessmentRequests from informationProtection
type ThreatAssessmentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ThreatAssessmentRequestItemRequestBuilderPatchOptions options for Patch
type ThreatAssessmentRequestItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ThreatAssessmentRequest;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewThreatAssessmentRequestItemRequestBuilderInternal instantiates a new ThreatAssessmentRequestItemRequestBuilder and sets the default values.
func NewThreatAssessmentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ThreatAssessmentRequestItemRequestBuilder) {
    m := &ThreatAssessmentRequestItemRequestBuilder{
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
// NewThreatAssessmentRequestItemRequestBuilder instantiates a new ThreatAssessmentRequestItemRequestBuilder and sets the default values.
func NewThreatAssessmentRequestItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ThreatAssessmentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatAssessmentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property threatAssessmentRequests for informationProtection
func (m *ThreatAssessmentRequestItemRequestBuilder) CreateDeleteRequestInformation(options *ThreatAssessmentRequestItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ThreatAssessmentRequestItemRequestBuilder) CreateGetRequestInformation(options *ThreatAssessmentRequestItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ThreatAssessmentRequestItemRequestBuilder) CreatePatchRequestInformation(options *ThreatAssessmentRequestItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ThreatAssessmentRequestItemRequestBuilder) Delete(options *ThreatAssessmentRequestItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get threatAssessmentRequests from informationProtection
func (m *ThreatAssessmentRequestItemRequestBuilder) Get(options *ThreatAssessmentRequestItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ThreatAssessmentRequest, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewThreatAssessmentRequest() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ThreatAssessmentRequest), nil
}
// Patch update the navigation property threatAssessmentRequests in informationProtection
func (m *ThreatAssessmentRequestItemRequestBuilder) Patch(options *ThreatAssessmentRequestItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ThreatAssessmentRequestItemRequestBuilder) Results()(*i830d3ff070fd8ee773495b3c0281bb745b2cb51d722269cf3375470d7bde5582.ResultsRequestBuilder) {
    return i830d3ff070fd8ee773495b3c0281bb745b2cb51d722269cf3375470d7bde5582.NewResultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResultsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.informationProtection.threatAssessmentRequests.item.results.item collection
func (m *ThreatAssessmentRequestItemRequestBuilder) ResultsById(id string)(*ic55cf4dfb7b61ba1f1a933e1461e5ca2dd101263eebcd32b51393ae518540c31.ThreatAssessmentResultItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["threatAssessmentResult_id"] = id
    }
    return ic55cf4dfb7b61ba1f1a933e1461e5ca2dd101263eebcd32b51393ae518540c31.NewThreatAssessmentResultItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
