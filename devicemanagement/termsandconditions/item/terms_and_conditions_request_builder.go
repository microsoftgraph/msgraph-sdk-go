package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1b63f23b60a34f4a2123b7bb0bda1a81794c3af78660cac406cd4a3b1073a36e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/termsandconditions/item/acceptancestatuses"
    i311c330be28d96ff5663f2910da29ba0b1a78acd2fc73a31a5135e243041cb6e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/termsandconditions/item/assignments"
    i2fc7a2614754963997f99635ae3e50b962e993d7cc3fdeb1c79dc9fa08d7674c "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/termsandconditions/item/assignments/item"
    i4153c8e1ab57c8baaa98c7558ada5d29fa34dd83a3f935d56846eea107484e94 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/termsandconditions/item/acceptancestatuses/item"
)

// TermsAndConditionsRequestBuilder builds and executes requests for operations under \deviceManagement\termsAndConditions\{termsAndConditions-id}
type TermsAndConditionsRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TermsAndConditionsRequestBuilderDeleteOptions options for Delete
type TermsAndConditionsRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermsAndConditionsRequestBuilderGetOptions options for Get
type TermsAndConditionsRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TermsAndConditionsRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermsAndConditionsRequestBuilderGetQueryParameters the terms and conditions associated with device management of the company.
type TermsAndConditionsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// TermsAndConditionsRequestBuilderPatchOptions options for Patch
type TermsAndConditionsRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TermsAndConditions;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TermsAndConditionsRequestBuilder) AcceptanceStatuses()(*i1b63f23b60a34f4a2123b7bb0bda1a81794c3af78660cac406cd4a3b1073a36e.AcceptanceStatusesRequestBuilder) {
    return i1b63f23b60a34f4a2123b7bb0bda1a81794c3af78660cac406cd4a3b1073a36e.NewAcceptanceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptanceStatusesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.termsAndConditions.item.acceptanceStatuses.item collection
func (m *TermsAndConditionsRequestBuilder) AcceptanceStatusesById(id string)(*i4153c8e1ab57c8baaa98c7558ada5d29fa34dd83a3f935d56846eea107484e94.TermsAndConditionsAcceptanceStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsAcceptanceStatus_id"] = id
    }
    return i4153c8e1ab57c8baaa98c7558ada5d29fa34dd83a3f935d56846eea107484e94.NewTermsAndConditionsAcceptanceStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TermsAndConditionsRequestBuilder) Assignments()(*i311c330be28d96ff5663f2910da29ba0b1a78acd2fc73a31a5135e243041cb6e.AssignmentsRequestBuilder) {
    return i311c330be28d96ff5663f2910da29ba0b1a78acd2fc73a31a5135e243041cb6e.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.termsAndConditions.item.assignments.item collection
func (m *TermsAndConditionsRequestBuilder) AssignmentsById(id string)(*i2fc7a2614754963997f99635ae3e50b962e993d7cc3fdeb1c79dc9fa08d7674c.TermsAndConditionsAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsAssignment_id"] = id
    }
    return i2fc7a2614754963997f99635ae3e50b962e993d7cc3fdeb1c79dc9fa08d7674c.NewTermsAndConditionsAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTermsAndConditionsRequestBuilderInternal instantiates a new TermsAndConditionsRequestBuilder and sets the default values.
func NewTermsAndConditionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermsAndConditionsRequestBuilder) {
    m := &TermsAndConditionsRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/termsAndConditions/{termsAndConditions_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTermsAndConditionsRequestBuilder instantiates a new TermsAndConditionsRequestBuilder and sets the default values.
func NewTermsAndConditionsRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermsAndConditionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsAndConditionsRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsRequestBuilder) CreateDeleteRequestInformation(options *TermsAndConditionsRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsRequestBuilder) CreateGetRequestInformation(options *TermsAndConditionsRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsRequestBuilder) CreatePatchRequestInformation(options *TermsAndConditionsRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsRequestBuilder) Delete(options *TermsAndConditionsRequestBuilderDeleteOptions)(error) {
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
// Get the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsRequestBuilder) Get(options *TermsAndConditionsRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TermsAndConditions, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewTermsAndConditions() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TermsAndConditions), nil
}
// Patch the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsRequestBuilder) Patch(options *TermsAndConditionsRequestBuilderPatchOptions)(error) {
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
