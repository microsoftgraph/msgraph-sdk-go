package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
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
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CallItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CallItemRequestBuilderGetQueryParameters get calls from communications
type CallItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CallItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CallItemRequestBuilderGetQueryParameters
}
// CallItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CallItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Answer the answer property
func (m *CallItemRequestBuilder) Answer()(*ia71afd593ae0c607d4acf1137eef2551b24d75551eb2631d48f91f4361dc4882.AnswerRequestBuilder) {
    return ia71afd593ae0c607d4acf1137eef2551b24d75551eb2631d48f91f4361dc4882.NewAnswerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AudioRoutingGroups the audioRoutingGroups property
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
        urlTplParams["audioRoutingGroup%2Did"] = id
    }
    return i5809389356a44ad33cb003cb59127db82d834418c35f9604ed259b7fe5afd5ce.NewAudioRoutingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CancelMediaProcessing the cancelMediaProcessing property
func (m *CallItemRequestBuilder) CancelMediaProcessing()(*i9597d4ee7219b291ed6ed4ac94f7ebb5f0666e8551606a8052e7ef400c7bf200.CancelMediaProcessingRequestBuilder) {
    return i9597d4ee7219b291ed6ed4ac94f7ebb5f0666e8551606a8052e7ef400c7bf200.NewCancelMediaProcessingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangeScreenSharingRole the changeScreenSharingRole property
func (m *CallItemRequestBuilder) ChangeScreenSharingRole()(*ie9da2770859e3dca5a85a6b4849aa09b988da20f525af2a9495d23bd3bfd686b.ChangeScreenSharingRoleRequestBuilder) {
    return ie9da2770859e3dca5a85a6b4849aa09b988da20f525af2a9495d23bd3bfd686b.NewChangeScreenSharingRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCallItemRequestBuilderInternal instantiates a new CallItemRequestBuilder and sets the default values.
func NewCallItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallItemRequestBuilder) {
    m := &CallItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/communications/calls/{call%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewCallItemRequestBuilder instantiates a new CallItemRequestBuilder and sets the default values.
func NewCallItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CallItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCallItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property calls for communications
func (m *CallItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property calls for communications
func (m *CallItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *CallItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get calls from communications
func (m *CallItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get calls from communications
func (m *CallItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *CallItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property calls in communications
func (m *CallItemRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property calls in communications
func (m *CallItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable, requestConfiguration *CallItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property calls for communications
func (m *CallItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property calls for communications
func (m *CallItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *CallItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get calls from communications
func (m *CallItemRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler get calls from communications
func (m *CallItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *CallItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCallFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable), nil
}
// KeepAlive the keepAlive property
func (m *CallItemRequestBuilder) KeepAlive()(*ie2968de31b227a313ec19bfb3d49e7376b707c33a3ef61ee430fc608fc92b15d.KeepAliveRequestBuilder) {
    return ie2968de31b227a313ec19bfb3d49e7376b707c33a3ef61ee430fc608fc92b15d.NewKeepAliveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Mute the mute property
func (m *CallItemRequestBuilder) Mute()(*i16df7b37097565742989cf0afde62825b1acba1907c5e4f5f010c3897c194dc4.MuteRequestBuilder) {
    return i16df7b37097565742989cf0afde62825b1acba1907c5e4f5f010c3897c194dc4.NewMuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Operations the operations property
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
        urlTplParams["commsOperation%2Did"] = id
    }
    return ied141ef3c3b82e58ad49b195ba5c23ec1d4cc768e035310e410fa372b3a1fc49.NewCommsOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Participants the participants property
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
        urlTplParams["participant%2Did"] = id
    }
    return ib16dc40ca9627c2186a56c4d731808da49ba6e8a34c17a076a47f881856366d7.NewParticipantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calls in communications
func (m *CallItemRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property calls in communications
func (m *CallItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable, requestConfiguration *CallItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// PlayPrompt the playPrompt property
func (m *CallItemRequestBuilder) PlayPrompt()(*i1dad219920e40389a4ebc6644b092da3826fe4c3784204c96119e935c730a5e9.PlayPromptRequestBuilder) {
    return i1dad219920e40389a4ebc6644b092da3826fe4c3784204c96119e935c730a5e9.NewPlayPromptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecordResponse the recordResponse property
func (m *CallItemRequestBuilder) RecordResponse()(*ic39f6d760e484c6cb7e55f7c08819553f2edb8148ca9b866c4410d6f117ebec5.RecordResponseRequestBuilder) {
    return ic39f6d760e484c6cb7e55f7c08819553f2edb8148ca9b866c4410d6f117ebec5.NewRecordResponseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Redirect the redirect property
func (m *CallItemRequestBuilder) Redirect()(*i187513783d2415a284900df1ba8249d112b88ac196e6a9d030d285c03b187903.RedirectRequestBuilder) {
    return i187513783d2415a284900df1ba8249d112b88ac196e6a9d030d285c03b187903.NewRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reject the reject property
func (m *CallItemRequestBuilder) Reject()(*ide3c7ebfce3fcf67decd28a86fe72e18f04a31dbafdecea60efdc8ec16540ce8.RejectRequestBuilder) {
    return ide3c7ebfce3fcf67decd28a86fe72e18f04a31dbafdecea60efdc8ec16540ce8.NewRejectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeToTone the subscribeToTone property
func (m *CallItemRequestBuilder) SubscribeToTone()(*i300aa200f9ff1ec2bb7fce7b1636cb382887168d515417ba91f3f35ca26c15af.SubscribeToToneRequestBuilder) {
    return i300aa200f9ff1ec2bb7fce7b1636cb382887168d515417ba91f3f35ca26c15af.NewSubscribeToToneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Transfer the transfer property
func (m *CallItemRequestBuilder) Transfer()(*i6c297ccbb463969476936203bb9956529571241de2ba0f09753156e86f7c8519.TransferRequestBuilder) {
    return i6c297ccbb463969476936203bb9956529571241de2ba0f09753156e86f7c8519.NewTransferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unmute the unmute property
func (m *CallItemRequestBuilder) Unmute()(*i282368a08e95a72a9c387ceb899bceeb76a2496e0734a9ecccbb87afbaa0f571.UnmuteRequestBuilder) {
    return i282368a08e95a72a9c387ceb899bceeb76a2496e0734a9ecccbb87afbaa0f571.NewUnmuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateRecordingStatus the updateRecordingStatus property
func (m *CallItemRequestBuilder) UpdateRecordingStatus()(*i6f0caf07b9d5829ba084ad4149cb4986ad69a13c935041d3da78c3cf1a5e3699.UpdateRecordingStatusRequestBuilder) {
    return i6f0caf07b9d5829ba084ad4149cb4986ad69a13c935041d3da78c3cf1a5e3699.NewUpdateRecordingStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
