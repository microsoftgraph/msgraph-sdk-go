package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4099812734ce25dbdeab4692470c91900c048e8dee3becb3e5c54a0326239224 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/devicestates"
    i51b72059dfed2f421aeb43f36dd6cb3a75014ea68e29bfe9f0763a1e28dc32bd "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/assignments"
    i571f8457b84b2b26f1f46b07c4fa37d92628066190b56c6d7b9dc33cbb426950 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/installsummary"
    ib4714ec089652eb60ee2444506e153b0e8d535e7937b3b83850dc25a80de4be4 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/userstatesummary"
    idc6ae391c23f1dd68e1a8ab2a8216079a7e5b650c41a8ff3082a6f856ddacad3 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/assign"
    i162ff3d0f43f21953d43cb94a636ae231382fdb4a6c3f80ff27c8fb165d745cb "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/assignments/item"
    i9ce747e428e80c804ab60cf526324b5166b04a2d275fe59526d2b02eb3bc9118 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/devicestates/item"
    if265c839ceb44280882d17ad922716116a39c1ce469abc935e86c7c536ccf9e0 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item/userstatesummary/item"
)

// Builds and executes requests for operations under \deviceAppManagement\managedEBooks\{managedEBook-id}
type ManagedEBookRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ManagedEBookRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ManagedEBookRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedEBookRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The Managed eBook.
type ManagedEBookRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ManagedEBookRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedEBook;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagedEBookRequestBuilder) Assign()(*idc6ae391c23f1dd68e1a8ab2a8216079a7e5b650c41a8ff3082a6f856ddacad3.AssignRequestBuilder) {
    return idc6ae391c23f1dd68e1a8ab2a8216079a7e5b650c41a8ff3082a6f856ddacad3.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedEBookRequestBuilder) Assignments()(*i51b72059dfed2f421aeb43f36dd6cb3a75014ea68e29bfe9f0763a1e28dc32bd.AssignmentsRequestBuilder) {
    return i51b72059dfed2f421aeb43f36dd6cb3a75014ea68e29bfe9f0763a1e28dc32bd.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedEBooks.item.assignments.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ManagedEBookRequestBuilder) AssignmentsById(id string)(*i162ff3d0f43f21953d43cb94a636ae231382fdb4a6c3f80ff27c8fb165d745cb.ManagedEBookAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBookAssignment_id"] = id
    }
    return i162ff3d0f43f21953d43cb94a636ae231382fdb4a6c3f80ff27c8fb165d745cb.NewManagedEBookAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new ManagedEBookRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagedEBookRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedEBookRequestBuilder) {
    m := &ManagedEBookRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement/managedEBooks/{managedEBook_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ManagedEBookRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewManagedEBookRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedEBookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedEBookRequestBuilderInternal(urlParams, requestAdapter)
}
// The Managed eBook.
// Parameters:
//  - options : Options for the request
func (m *ManagedEBookRequestBuilder) CreateDeleteRequestInformation(options *ManagedEBookRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The Managed eBook.
// Parameters:
//  - options : Options for the request
func (m *ManagedEBookRequestBuilder) CreateGetRequestInformation(options *ManagedEBookRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The Managed eBook.
// Parameters:
//  - options : Options for the request
func (m *ManagedEBookRequestBuilder) CreatePatchRequestInformation(options *ManagedEBookRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The Managed eBook.
// Parameters:
//  - options : Options for the request
func (m *ManagedEBookRequestBuilder) Delete(options *ManagedEBookRequestBuilderDeleteOptions)(error) {
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
func (m *ManagedEBookRequestBuilder) DeviceStates()(*i4099812734ce25dbdeab4692470c91900c048e8dee3becb3e5c54a0326239224.DeviceStatesRequestBuilder) {
    return i4099812734ce25dbdeab4692470c91900c048e8dee3becb3e5c54a0326239224.NewDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedEBooks.item.deviceStates.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ManagedEBookRequestBuilder) DeviceStatesById(id string)(*i9ce747e428e80c804ab60cf526324b5166b04a2d275fe59526d2b02eb3bc9118.DeviceInstallStateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceInstallState_id"] = id
    }
    return i9ce747e428e80c804ab60cf526324b5166b04a2d275fe59526d2b02eb3bc9118.NewDeviceInstallStateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The Managed eBook.
// Parameters:
//  - options : Options for the request
func (m *ManagedEBookRequestBuilder) Get(options *ManagedEBookRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedEBook, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewManagedEBook() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedEBook), nil
}
func (m *ManagedEBookRequestBuilder) InstallSummary()(*i571f8457b84b2b26f1f46b07c4fa37d92628066190b56c6d7b9dc33cbb426950.InstallSummaryRequestBuilder) {
    return i571f8457b84b2b26f1f46b07c4fa37d92628066190b56c6d7b9dc33cbb426950.NewInstallSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The Managed eBook.
// Parameters:
//  - options : Options for the request
func (m *ManagedEBookRequestBuilder) Patch(options *ManagedEBookRequestBuilderPatchOptions)(error) {
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
func (m *ManagedEBookRequestBuilder) UserStateSummary()(*ib4714ec089652eb60ee2444506e153b0e8d535e7937b3b83850dc25a80de4be4.UserStateSummaryRequestBuilder) {
    return ib4714ec089652eb60ee2444506e153b0e8d535e7937b3b83850dc25a80de4be4.NewUserStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedEBooks.item.userStateSummary.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ManagedEBookRequestBuilder) UserStateSummaryById(id string)(*if265c839ceb44280882d17ad922716116a39c1ce469abc935e86c7c536ccf9e0.UserInstallStateSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userInstallStateSummary_id"] = id
    }
    return if265c839ceb44280882d17ad922716116a39c1ce469abc935e86c7c536ccf9e0.NewUserInstallStateSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
