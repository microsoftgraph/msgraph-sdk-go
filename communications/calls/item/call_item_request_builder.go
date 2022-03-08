package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i07f45294889a4c0baadf8d3c66171a11e824128d56c5d5922e1b42e5a85ab314 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/participants"
    i16df7b37097565742989cf0afde62825b1acba1907c5e4f5f010c3897c194dc4 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/mute"
    i187513783d2415a284900df1ba8249d112b88ac196e6a9d030d285c03b187903 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/redirect"
    i1dad219920e40389a4ebc6644b092da3826fe4c3784204c96119e935c730a5e9 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/playprompt"
    i282368a08e95a72a9c387ceb899bceeb76a2496e0734a9ecccbb87afbaa0f571 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/unmute"
    i300aa200f9ff1ec2bb7fce7b1636cb382887168d515417ba91f3f35ca26c15af "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/subscribetotone"
    i34ed9d0ef5bd801c08c976e0011aa37aa3b7cd0346f50aed0b58fbda8fa2d09e "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/operations"
    i6c297ccbb463969476936203bb9956529571241de2ba0f09753156e86f7c8519 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/transfer"
    i6f0caf07b9d5829ba084ad4149cb4986ad69a13c935041d3da78c3cf1a5e3699 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/updaterecordingstatus"
    i9597d4ee7219b291ed6ed4ac94f7ebb5f0666e8551606a8052e7ef400c7bf200 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/cancelmediaprocessing"
    ia71afd593ae0c607d4acf1137eef2551b24d75551eb2631d48f91f4361dc4882 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/answer"
    ib42114d547ae3a5efaa8e2718244399826bae33e41e3785ade01e865ab37b3dd "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/audioroutinggroups"
    ic39f6d760e484c6cb7e55f7c08819553f2edb8148ca9b866c4410d6f117ebec5 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/recordresponse"
    ide3c7ebfce3fcf67decd28a86fe72e18f04a31dbafdecea60efdc8ec16540ce8 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/reject"
    ie2968de31b227a313ec19bfb3d49e7376b707c33a3ef61ee430fc608fc92b15d "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/keepalive"
    ie9da2770859e3dca5a85a6b4849aa09b988da20f525af2a9495d23bd3bfd686b "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/changescreensharingrole"
    i5809389356a44ad33cb003cb59127db82d834418c35f9604ed259b7fe5afd5ce "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/audioroutinggroups/item"
    ib16dc40ca9627c2186a56c4d731808da49ba6e8a34c17a076a47f881856366d7 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/participants/item"
    ied141ef3c3b82e58ad49b195ba5c23ec1d4cc768e035310e410fa372b3a1fc49 "github.com/microsoftgraph/msgraph-sdk-go/communications/calls/item/operations/item"
)

// CallItemRequestBuilder provides operations to manage the calls property of the microsoft.graph.cloudCommunications entity.
type CallItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// CallItemRequestBuilderDeleteOptions options for Delete
type CallItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CallItemRequestBuilderGetOptions options for Get
type CallItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *CallItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// CallItemRequestBuilderGetQueryParameters get calls from communications
type CallItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// CallItemRequestBuilderPatchOptions options for Patch
type CallItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Callable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *CallItemRequestBuilder) Answer()(*ia71afd593ae0c607d4acf1137eef2551b24d75551eb2631d48f91f4361dc4882.AnswerRequestBuilder) {
    return ia71afd593ae0c607d4acf1137eef2551b24d75551eb2631d48f91f4361dc4882.NewAnswerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) AudioRoutingGroups()(*ib42114d547ae3a5efaa8e2718244399826bae33e41e3785ade01e865ab37b3dd.AudioRoutingGroupsRequestBuilder) {
    return ib42114d547ae3a5efaa8e2718244399826bae33e41e3785ade01e865ab37b3dd.NewAudioRoutingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AudioRoutingGroupsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.communications.calls.item.audioRoutingGroups.item collection
func (m *CallItemRequestBuilder) AudioRoutingGroupsById(id string)(*i5809389356a44ad33cb003cb59127db82d834418c35f9604ed259b7fe5afd5ce.AudioRoutingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["audioRoutingGroup_id"] = id
    }
    return i5809389356a44ad33cb003cb59127db82d834418c35f9604ed259b7fe5afd5ce.NewAudioRoutingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CallItemRequestBuilder) CancelMediaProcessing()(*i9597d4ee7219b291ed6ed4ac94f7ebb5f0666e8551606a8052e7ef400c7bf200.CancelMediaProcessingRequestBuilder) {
    return i9597d4ee7219b291ed6ed4ac94f7ebb5f0666e8551606a8052e7ef400c7bf200.NewCancelMediaProcessingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) ChangeScreenSharingRole()(*ie9da2770859e3dca5a85a6b4849aa09b988da20f525af2a9495d23bd3bfd686b.ChangeScreenSharingRoleRequestBuilder) {
    return ie9da2770859e3dca5a85a6b4849aa09b988da20f525af2a9495d23bd3bfd686b.NewChangeScreenSharingRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCallItemRequestBuilderInternal instantiates a new CallItemRequestBuilder and sets the default values.
func NewCallItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CallItemRequestBuilder) {
    m := &CallItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCallItemRequestBuilder instantiates a new CallItemRequestBuilder and sets the default values.
func NewCallItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*CallItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property calls for communications
func (m *CallItemRequestBuilder) CreateDeleteRequestInformation(options *CallItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get calls from communications
func (m *CallItemRequestBuilder) CreateGetRequestInformation(options *CallItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property calls in communications
func (m *CallItemRequestBuilder) CreatePatchRequestInformation(options *CallItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property calls for communications
func (m *CallItemRequestBuilder) Delete(options *CallItemRequestBuilderDeleteOptions)(error) {
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
// Get get calls from communications
func (m *CallItemRequestBuilder) Get(options *CallItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Callable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateCallFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Callable), nil
}
func (m *CallItemRequestBuilder) KeepAlive()(*ie2968de31b227a313ec19bfb3d49e7376b707c33a3ef61ee430fc608fc92b15d.KeepAliveRequestBuilder) {
    return ie2968de31b227a313ec19bfb3d49e7376b707c33a3ef61ee430fc608fc92b15d.NewKeepAliveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Mute()(*i16df7b37097565742989cf0afde62825b1acba1907c5e4f5f010c3897c194dc4.MuteRequestBuilder) {
    return i16df7b37097565742989cf0afde62825b1acba1907c5e4f5f010c3897c194dc4.NewMuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Operations()(*i34ed9d0ef5bd801c08c976e0011aa37aa3b7cd0346f50aed0b58fbda8fa2d09e.OperationsRequestBuilder) {
    return i34ed9d0ef5bd801c08c976e0011aa37aa3b7cd0346f50aed0b58fbda8fa2d09e.NewOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.communications.calls.item.operations.item collection
func (m *CallItemRequestBuilder) OperationsById(id string)(*ied141ef3c3b82e58ad49b195ba5c23ec1d4cc768e035310e410fa372b3a1fc49.CommsOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["commsOperation_id"] = id
    }
    return ied141ef3c3b82e58ad49b195ba5c23ec1d4cc768e035310e410fa372b3a1fc49.NewCommsOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Participants()(*i07f45294889a4c0baadf8d3c66171a11e824128d56c5d5922e1b42e5a85ab314.ParticipantsRequestBuilder) {
    return i07f45294889a4c0baadf8d3c66171a11e824128d56c5d5922e1b42e5a85ab314.NewParticipantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParticipantsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.communications.calls.item.participants.item collection
func (m *CallItemRequestBuilder) ParticipantsById(id string)(*ib16dc40ca9627c2186a56c4d731808da49ba6e8a34c17a076a47f881856366d7.ParticipantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["participant_id"] = id
    }
    return ib16dc40ca9627c2186a56c4d731808da49ba6e8a34c17a076a47f881856366d7.NewParticipantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calls in communications
func (m *CallItemRequestBuilder) Patch(options *CallItemRequestBuilderPatchOptions)(error) {
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
func (m *CallItemRequestBuilder) PlayPrompt()(*i1dad219920e40389a4ebc6644b092da3826fe4c3784204c96119e935c730a5e9.PlayPromptRequestBuilder) {
    return i1dad219920e40389a4ebc6644b092da3826fe4c3784204c96119e935c730a5e9.NewPlayPromptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) RecordResponse()(*ic39f6d760e484c6cb7e55f7c08819553f2edb8148ca9b866c4410d6f117ebec5.RecordResponseRequestBuilder) {
    return ic39f6d760e484c6cb7e55f7c08819553f2edb8148ca9b866c4410d6f117ebec5.NewRecordResponseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Redirect()(*i187513783d2415a284900df1ba8249d112b88ac196e6a9d030d285c03b187903.RedirectRequestBuilder) {
    return i187513783d2415a284900df1ba8249d112b88ac196e6a9d030d285c03b187903.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Reject()(*ide3c7ebfce3fcf67decd28a86fe72e18f04a31dbafdecea60efdc8ec16540ce8.RejectRequestBuilder) {
    return ide3c7ebfce3fcf67decd28a86fe72e18f04a31dbafdecea60efdc8ec16540ce8.NewRejectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) SubscribeToTone()(*i300aa200f9ff1ec2bb7fce7b1636cb382887168d515417ba91f3f35ca26c15af.SubscribeToToneRequestBuilder) {
    return i300aa200f9ff1ec2bb7fce7b1636cb382887168d515417ba91f3f35ca26c15af.NewSubscribeToToneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Transfer()(*i6c297ccbb463969476936203bb9956529571241de2ba0f09753156e86f7c8519.TransferRequestBuilder) {
    return i6c297ccbb463969476936203bb9956529571241de2ba0f09753156e86f7c8519.NewTransferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) Unmute()(*i282368a08e95a72a9c387ceb899bceeb76a2496e0734a9ecccbb87afbaa0f571.UnmuteRequestBuilder) {
    return i282368a08e95a72a9c387ceb899bceeb76a2496e0734a9ecccbb87afbaa0f571.NewUnmuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *CallItemRequestBuilder) UpdateRecordingStatus()(*i6f0caf07b9d5829ba084ad4149cb4986ad69a13c935041d3da78c3cf1a5e3699.UpdateRecordingStatusRequestBuilder) {
    return i6f0caf07b9d5829ba084ad4149cb4986ad69a13c935041d3da78c3cf1a5e3699.NewUpdateRecordingStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
