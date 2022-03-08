package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1b63f23b60a34f4a2123b7bb0bda1a81794c3af78660cac406cd4a3b1073a36e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/termsandconditions/item/acceptancestatuses"
    i311c330be28d96ff5663f2910da29ba0b1a78acd2fc73a31a5135e243041cb6e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/termsandconditions/item/assignments"
    i2fc7a2614754963997f99635ae3e50b962e993d7cc3fdeb1c79dc9fa08d7674c "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/termsandconditions/item/assignments/item"
    i4153c8e1ab57c8baaa98c7558ada5d29fa34dd83a3f935d56846eea107484e94 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/termsandconditions/item/acceptancestatuses/item"
)

// TermsAndConditionsItemRequestBuilder provides operations to manage the termsAndConditions property of the microsoft.graph.deviceManagement entity.
type TermsAndConditionsItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// TermsAndConditionsItemRequestBuilderDeleteOptions options for Delete
type TermsAndConditionsItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermsAndConditionsItemRequestBuilderGetOptions options for Get
type TermsAndConditionsItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *TermsAndConditionsItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// TermsAndConditionsItemRequestBuilderGetQueryParameters the terms and conditions associated with device management of the company.
type TermsAndConditionsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// TermsAndConditionsItemRequestBuilderPatchOptions options for Patch
type TermsAndConditionsItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TermsAndConditionsable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *TermsAndConditionsItemRequestBuilder) AcceptanceStatuses()(*i1b63f23b60a34f4a2123b7bb0bda1a81794c3af78660cac406cd4a3b1073a36e.AcceptanceStatusesRequestBuilder) {
    return i1b63f23b60a34f4a2123b7bb0bda1a81794c3af78660cac406cd4a3b1073a36e.NewAcceptanceStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AcceptanceStatusesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.termsAndConditions.item.acceptanceStatuses.item collection
func (m *TermsAndConditionsItemRequestBuilder) AcceptanceStatusesById(id string)(*i4153c8e1ab57c8baaa98c7558ada5d29fa34dd83a3f935d56846eea107484e94.TermsAndConditionsAcceptanceStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsAcceptanceStatus_id"] = id
    }
    return i4153c8e1ab57c8baaa98c7558ada5d29fa34dd83a3f935d56846eea107484e94.NewTermsAndConditionsAcceptanceStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *TermsAndConditionsItemRequestBuilder) Assignments()(*i311c330be28d96ff5663f2910da29ba0b1a78acd2fc73a31a5135e243041cb6e.AssignmentsRequestBuilder) {
    return i311c330be28d96ff5663f2910da29ba0b1a78acd2fc73a31a5135e243041cb6e.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.termsAndConditions.item.assignments.item collection
func (m *TermsAndConditionsItemRequestBuilder) AssignmentsById(id string)(*i2fc7a2614754963997f99635ae3e50b962e993d7cc3fdeb1c79dc9fa08d7674c.TermsAndConditionsAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditionsAssignment_id"] = id
    }
    return i2fc7a2614754963997f99635ae3e50b962e993d7cc3fdeb1c79dc9fa08d7674c.NewTermsAndConditionsAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewTermsAndConditionsItemRequestBuilderInternal instantiates a new TermsAndConditionsItemRequestBuilder and sets the default values.
func NewTermsAndConditionsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermsAndConditionsItemRequestBuilder) {
    m := &TermsAndConditionsItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement/termsAndConditions/{termsAndConditions_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTermsAndConditionsItemRequestBuilder instantiates a new TermsAndConditionsItemRequestBuilder and sets the default values.
func NewTermsAndConditionsItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*TermsAndConditionsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsAndConditionsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property termsAndConditions for deviceManagement
func (m *TermsAndConditionsItemRequestBuilder) CreateDeleteRequestInformation(options *TermsAndConditionsItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *TermsAndConditionsItemRequestBuilder) CreateGetRequestInformation(options *TermsAndConditionsItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property termsAndConditions in deviceManagement
func (m *TermsAndConditionsItemRequestBuilder) CreatePatchRequestInformation(options *TermsAndConditionsItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property termsAndConditions for deviceManagement
func (m *TermsAndConditionsItemRequestBuilder) Delete(options *TermsAndConditionsItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the terms and conditions associated with device management of the company.
func (m *TermsAndConditionsItemRequestBuilder) Get(options *TermsAndConditionsItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TermsAndConditionsable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateTermsAndConditionsFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.TermsAndConditionsable), nil
}
// Patch update the navigation property termsAndConditions in deviceManagement
func (m *TermsAndConditionsItemRequestBuilder) Patch(options *TermsAndConditionsItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
