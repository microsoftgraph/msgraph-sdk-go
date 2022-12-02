package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CommunicationsCallsCallItemRequestBuilder provides operations to manage the calls property of the microsoft.graph.cloudCommunications entity.
type CommunicationsCallsCallItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// CommunicationsCallsCallItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsCallsCallItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CommunicationsCallsCallItemRequestBuilderGetQueryParameters get calls from communications
type CommunicationsCallsCallItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// CommunicationsCallsCallItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsCallsCallItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CommunicationsCallsCallItemRequestBuilderGetQueryParameters
}
// CommunicationsCallsCallItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CommunicationsCallsCallItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AddLargeGalleryView provides operations to call the addLargeGalleryView method.
func (m *CommunicationsCallsCallItemRequestBuilder) AddLargeGalleryView()(*CommunicationsCallsItemAddLargeGalleryViewRequestBuilder) {
    return NewCommunicationsCallsItemAddLargeGalleryViewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Answer provides operations to call the answer method.
func (m *CommunicationsCallsCallItemRequestBuilder) Answer()(*CommunicationsCallsItemAnswerRequestBuilder) {
    return NewCommunicationsCallsItemAnswerRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AudioRoutingGroups provides operations to manage the audioRoutingGroups property of the microsoft.graph.call entity.
func (m *CommunicationsCallsCallItemRequestBuilder) AudioRoutingGroups()(*CommunicationsCallsItemAudioRoutingGroupsRequestBuilder) {
    return NewCommunicationsCallsItemAudioRoutingGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AudioRoutingGroupsById provides operations to manage the audioRoutingGroups property of the microsoft.graph.call entity.
func (m *CommunicationsCallsCallItemRequestBuilder) AudioRoutingGroupsById(id string)(*CommunicationsCallsItemAudioRoutingGroupsAudioRoutingGroupItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["audioRoutingGroup%2Did"] = id
    }
    return NewCommunicationsCallsItemAudioRoutingGroupsAudioRoutingGroupItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CancelMediaProcessing provides operations to call the cancelMediaProcessing method.
func (m *CommunicationsCallsCallItemRequestBuilder) CancelMediaProcessing()(*CommunicationsCallsItemCancelMediaProcessingRequestBuilder) {
    return NewCommunicationsCallsItemCancelMediaProcessingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChangeScreenSharingRole provides operations to call the changeScreenSharingRole method.
func (m *CommunicationsCallsCallItemRequestBuilder) ChangeScreenSharingRole()(*CommunicationsCallsItemChangeScreenSharingRoleRequestBuilder) {
    return NewCommunicationsCallsItemChangeScreenSharingRoleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewCommunicationsCallsCallItemRequestBuilderInternal instantiates a new CallItemRequestBuilder and sets the default values.
func NewCommunicationsCallsCallItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommunicationsCallsCallItemRequestBuilder) {
    m := &CommunicationsCallsCallItemRequestBuilder{
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
// NewCommunicationsCallsCallItemRequestBuilder instantiates a new CallItemRequestBuilder and sets the default values.
func NewCommunicationsCallsCallItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CommunicationsCallsCallItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCommunicationsCallsCallItemRequestBuilderInternal(urlParams, requestAdapter)
}
// ContentSharingSessions provides operations to manage the contentSharingSessions property of the microsoft.graph.call entity.
func (m *CommunicationsCallsCallItemRequestBuilder) ContentSharingSessions()(*CommunicationsCallsItemContentSharingSessionsRequestBuilder) {
    return NewCommunicationsCallsItemContentSharingSessionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ContentSharingSessionsById provides operations to manage the contentSharingSessions property of the microsoft.graph.call entity.
func (m *CommunicationsCallsCallItemRequestBuilder) ContentSharingSessionsById(id string)(*CommunicationsCallsItemContentSharingSessionsContentSharingSessionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentSharingSession%2Did"] = id
    }
    return NewCommunicationsCallsItemContentSharingSessionsContentSharingSessionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property calls for communications
func (m *CommunicationsCallsCallItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *CommunicationsCallsCallItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CommunicationsCallsCallItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *CommunicationsCallsCallItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *CommunicationsCallsCallItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable, requestConfiguration *CommunicationsCallsCallItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property calls for communications
func (m *CommunicationsCallsCallItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *CommunicationsCallsCallItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get calls from communications
func (m *CommunicationsCallsCallItemRequestBuilder) Get(ctx context.Context, requestConfiguration *CommunicationsCallsCallItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCallFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable), nil
}
// KeepAlive provides operations to call the keepAlive method.
func (m *CommunicationsCallsCallItemRequestBuilder) KeepAlive()(*CommunicationsCallsItemKeepAliveRequestBuilder) {
    return NewCommunicationsCallsItemKeepAliveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Mute provides operations to call the mute method.
func (m *CommunicationsCallsCallItemRequestBuilder) Mute()(*CommunicationsCallsItemMuteRequestBuilder) {
    return NewCommunicationsCallsItemMuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.call entity.
func (m *CommunicationsCallsCallItemRequestBuilder) Operations()(*CommunicationsCallsItemOperationsRequestBuilder) {
    return NewCommunicationsCallsItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.call entity.
func (m *CommunicationsCallsCallItemRequestBuilder) OperationsById(id string)(*CommunicationsCallsItemOperationsCommsOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["commsOperation%2Did"] = id
    }
    return NewCommunicationsCallsItemOperationsCommsOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Participants provides operations to manage the participants property of the microsoft.graph.call entity.
func (m *CommunicationsCallsCallItemRequestBuilder) Participants()(*CommunicationsCallsItemParticipantsRequestBuilder) {
    return NewCommunicationsCallsItemParticipantsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ParticipantsById provides operations to manage the participants property of the microsoft.graph.call entity.
func (m *CommunicationsCallsCallItemRequestBuilder) ParticipantsById(id string)(*CommunicationsCallsItemParticipantsParticipantItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["participant%2Did"] = id
    }
    return NewCommunicationsCallsItemParticipantsParticipantItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property calls in communications
func (m *CommunicationsCallsCallItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable, requestConfiguration *CommunicationsCallsCallItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCallFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Callable), nil
}
// PlayPrompt provides operations to call the playPrompt method.
func (m *CommunicationsCallsCallItemRequestBuilder) PlayPrompt()(*CommunicationsCallsItemPlayPromptRequestBuilder) {
    return NewCommunicationsCallsItemPlayPromptRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RecordResponse provides operations to call the recordResponse method.
func (m *CommunicationsCallsCallItemRequestBuilder) RecordResponse()(*CommunicationsCallsItemRecordResponseRequestBuilder) {
    return NewCommunicationsCallsItemRecordResponseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Redirect provides operations to call the redirect method.
func (m *CommunicationsCallsCallItemRequestBuilder) Redirect()(*CommunicationsCallsItemRedirectRequestBuilder) {
    return NewCommunicationsCallsItemRedirectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Reject provides operations to call the reject method.
func (m *CommunicationsCallsCallItemRequestBuilder) Reject()(*CommunicationsCallsItemRejectRequestBuilder) {
    return NewCommunicationsCallsItemRejectRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscribeToTone provides operations to call the subscribeToTone method.
func (m *CommunicationsCallsCallItemRequestBuilder) SubscribeToTone()(*CommunicationsCallsItemSubscribeToToneRequestBuilder) {
    return NewCommunicationsCallsItemSubscribeToToneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Transfer provides operations to call the transfer method.
func (m *CommunicationsCallsCallItemRequestBuilder) Transfer()(*CommunicationsCallsItemTransferRequestBuilder) {
    return NewCommunicationsCallsItemTransferRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unmute provides operations to call the unmute method.
func (m *CommunicationsCallsCallItemRequestBuilder) Unmute()(*CommunicationsCallsItemUnmuteRequestBuilder) {
    return NewCommunicationsCallsItemUnmuteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// UpdateRecordingStatus provides operations to call the updateRecordingStatus method.
func (m *CommunicationsCallsCallItemRequestBuilder) UpdateRecordingStatus()(*CommunicationsCallsItemUpdateRecordingStatusRequestBuilder) {
    return NewCommunicationsCallsItemUpdateRecordingStatusRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
