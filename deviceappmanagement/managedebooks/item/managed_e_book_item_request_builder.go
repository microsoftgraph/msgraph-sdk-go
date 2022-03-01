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

// ManagedEBookItemRequestBuilder builds and executes requests for operations under \deviceAppManagement\managedEBooks\{managedEBook-id}
type ManagedEBookItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ManagedEBookItemRequestBuilderDeleteOptions options for Delete
type ManagedEBookItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedEBookItemRequestBuilderGetOptions options for Get
type ManagedEBookItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ManagedEBookItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ManagedEBookItemRequestBuilderGetQueryParameters the Managed eBook.
type ManagedEBookItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ManagedEBookItemRequestBuilderPatchOptions options for Patch
type ManagedEBookItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedEBook;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ManagedEBookItemRequestBuilder) Assign()(*idc6ae391c23f1dd68e1a8ab2a8216079a7e5b650c41a8ff3082a6f856ddacad3.AssignRequestBuilder) {
    return idc6ae391c23f1dd68e1a8ab2a8216079a7e5b650c41a8ff3082a6f856ddacad3.NewAssignRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ManagedEBookItemRequestBuilder) Assignments()(*i51b72059dfed2f421aeb43f36dd6cb3a75014ea68e29bfe9f0763a1e28dc32bd.AssignmentsRequestBuilder) {
    return i51b72059dfed2f421aeb43f36dd6cb3a75014ea68e29bfe9f0763a1e28dc32bd.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedEBooks.item.assignments.item collection
func (m *ManagedEBookItemRequestBuilder) AssignmentsById(id string)(*i162ff3d0f43f21953d43cb94a636ae231382fdb4a6c3f80ff27c8fb165d745cb.ManagedEBookAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBookAssignment_id"] = id
    }
    return i162ff3d0f43f21953d43cb94a636ae231382fdb4a6c3f80ff27c8fb165d745cb.NewManagedEBookAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewManagedEBookItemRequestBuilderInternal instantiates a new ManagedEBookItemRequestBuilder and sets the default values.
func NewManagedEBookItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedEBookItemRequestBuilder) {
    m := &ManagedEBookItemRequestBuilder{
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
// NewManagedEBookItemRequestBuilder instantiates a new ManagedEBookItemRequestBuilder and sets the default values.
func NewManagedEBookItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ManagedEBookItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewManagedEBookItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) CreateDeleteRequestInformation(options *ManagedEBookItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) CreateGetRequestInformation(options *ManagedEBookItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) CreatePatchRequestInformation(options *ManagedEBookItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) Delete(options *ManagedEBookItemRequestBuilderDeleteOptions)(error) {
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
func (m *ManagedEBookItemRequestBuilder) DeviceStates()(*i4099812734ce25dbdeab4692470c91900c048e8dee3becb3e5c54a0326239224.DeviceStatesRequestBuilder) {
    return i4099812734ce25dbdeab4692470c91900c048e8dee3becb3e5c54a0326239224.NewDeviceStatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceStatesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedEBooks.item.deviceStates.item collection
func (m *ManagedEBookItemRequestBuilder) DeviceStatesById(id string)(*i9ce747e428e80c804ab60cf526324b5166b04a2d275fe59526d2b02eb3bc9118.DeviceInstallStateItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceInstallState_id"] = id
    }
    return i9ce747e428e80c804ab60cf526324b5166b04a2d275fe59526d2b02eb3bc9118.NewDeviceInstallStateItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) Get(options *ManagedEBookItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedEBook, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewManagedEBook() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ManagedEBook), nil
}
func (m *ManagedEBookItemRequestBuilder) InstallSummary()(*i571f8457b84b2b26f1f46b07c4fa37d92628066190b56c6d7b9dc33cbb426950.InstallSummaryRequestBuilder) {
    return i571f8457b84b2b26f1f46b07c4fa37d92628066190b56c6d7b9dc33cbb426950.NewInstallSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch the Managed eBook.
func (m *ManagedEBookItemRequestBuilder) Patch(options *ManagedEBookItemRequestBuilderPatchOptions)(error) {
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
func (m *ManagedEBookItemRequestBuilder) UserStateSummary()(*ib4714ec089652eb60ee2444506e153b0e8d535e7937b3b83850dc25a80de4be4.UserStateSummaryRequestBuilder) {
    return ib4714ec089652eb60ee2444506e153b0e8d535e7937b3b83850dc25a80de4be4.NewUserStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UserStateSummaryById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedEBooks.item.userStateSummary.item collection
func (m *ManagedEBookItemRequestBuilder) UserStateSummaryById(id string)(*if265c839ceb44280882d17ad922716116a39c1ce469abc935e86c7c536ccf9e0.UserInstallStateSummaryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["userInstallStateSummary_id"] = id
    }
    return if265c839ceb44280882d17ad922716116a39c1ce469abc935e86c7c536ccf9e0.NewUserInstallStateSummaryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
