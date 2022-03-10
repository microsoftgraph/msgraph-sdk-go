package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i11526ee5699d623863f00b0b685d457eb13b19842baadbeddfd36eecf73dcdf5 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/channels/item/completemigration"
    i1370e74741c1f731fcda20e6e79b93c6cbee203716569746c1e43924df754b98 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/channels/item/filesfolder"
    iaf3119de5714740cbdb24901fa13a2700a0dc0f152b343a2c356536580c6e10b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/channels/item/tabs"
    ib9ab15dbda262e6b76749343df5d5f49e644a1926521ae149443532a84b18125 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/channels/item/removeemail"
    icb337b1772f31d93683bb7dfb1edf3c4a97b38d5449f9ca882a52e68e02d5e27 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/channels/item/members"
    iece274c950455f68d3b9284c5b46b89d9f9f8aa9c5482a9ba6d18dde70bd04f8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/channels/item/messages"
    if559d86b89add997d564b55f90c9800ec81efa7e7c483735cee45eedd11460bd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/channels/item/provisionemail"
    i045b00b932c574167b22217b96e68e85ba6f525a94bcd3b95118701b46c60f03 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/channels/item/tabs/item"
    ib543babcb7b560b08aa30ad60d5db8d65b281cf0de36ac91c72b9eb5099e576b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/channels/item/members/item"
    ic2fc4726b6ef1949b7220307cf2677355049193c9fd3815132e291979132f29b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/team/channels/item/messages/item"
)

// ChannelItemRequestBuilder provides operations to manage the channels property of the microsoft.graph.team entity.
type ChannelItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ChannelItemRequestBuilderDeleteOptions options for Delete
type ChannelItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChannelItemRequestBuilderGetOptions options for Get
type ChannelItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChannelItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChannelItemRequestBuilderGetQueryParameters the collection of channels and messages associated with the team.
type ChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ChannelItemRequestBuilderPatchOptions options for Patch
type ChannelItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channelable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ChannelItemRequestBuilder) CompleteMigration()(*i11526ee5699d623863f00b0b685d457eb13b19842baadbeddfd36eecf73dcdf5.CompleteMigrationRequestBuilder) {
    return i11526ee5699d623863f00b0b685d457eb13b19842baadbeddfd36eecf73dcdf5.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChannelItemRequestBuilder) {
    m := &ChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/team/channels/{channel_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChannelItemRequestBuilder instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property channels for groups
func (m *ChannelItemRequestBuilder) CreateDeleteRequestInformation(options *ChannelItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) CreateGetRequestInformation(options *ChannelItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property channels in groups
func (m *ChannelItemRequestBuilder) CreatePatchRequestInformation(options *ChannelItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property channels for groups
func (m *ChannelItemRequestBuilder) Delete(options *ChannelItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ChannelItemRequestBuilder) FilesFolder()(*i1370e74741c1f731fcda20e6e79b93c6cbee203716569746c1e43924df754b98.FilesFolderRequestBuilder) {
    return i1370e74741c1f731fcda20e6e79b93c6cbee203716569746c1e43924df754b98.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) Get(options *ChannelItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channelable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateChannelFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channelable), nil
}
func (m *ChannelItemRequestBuilder) Members()(*icb337b1772f31d93683bb7dfb1edf3c4a97b38d5449f9ca882a52e68e02d5e27.MembersRequestBuilder) {
    return icb337b1772f31d93683bb7dfb1edf3c4a97b38d5449f9ca882a52e68e02d5e27.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.channels.item.members.item collection
func (m *ChannelItemRequestBuilder) MembersById(id string)(*ib543babcb7b560b08aa30ad60d5db8d65b281cf0de36ac91c72b9eb5099e576b.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return ib543babcb7b560b08aa30ad60d5db8d65b281cf0de36ac91c72b9eb5099e576b.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChannelItemRequestBuilder) Messages()(*iece274c950455f68d3b9284c5b46b89d9f9f8aa9c5482a9ba6d18dde70bd04f8.MessagesRequestBuilder) {
    return iece274c950455f68d3b9284c5b46b89d9f9f8aa9c5482a9ba6d18dde70bd04f8.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.channels.item.messages.item collection
func (m *ChannelItemRequestBuilder) MessagesById(id string)(*ic2fc4726b6ef1949b7220307cf2677355049193c9fd3815132e291979132f29b.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return ic2fc4726b6ef1949b7220307cf2677355049193c9fd3815132e291979132f29b.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property channels in groups
func (m *ChannelItemRequestBuilder) Patch(options *ChannelItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ChannelItemRequestBuilder) ProvisionEmail()(*if559d86b89add997d564b55f90c9800ec81efa7e7c483735cee45eedd11460bd.ProvisionEmailRequestBuilder) {
    return if559d86b89add997d564b55f90c9800ec81efa7e7c483735cee45eedd11460bd.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChannelItemRequestBuilder) RemoveEmail()(*ib9ab15dbda262e6b76749343df5d5f49e644a1926521ae149443532a84b18125.RemoveEmailRequestBuilder) {
    return ib9ab15dbda262e6b76749343df5d5f49e644a1926521ae149443532a84b18125.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChannelItemRequestBuilder) Tabs()(*iaf3119de5714740cbdb24901fa13a2700a0dc0f152b343a2c356536580c6e10b.TabsRequestBuilder) {
    return iaf3119de5714740cbdb24901fa13a2700a0dc0f152b343a2c356536580c6e10b.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.team.channels.item.tabs.item collection
func (m *ChannelItemRequestBuilder) TabsById(id string)(*i045b00b932c574167b22217b96e68e85ba6f525a94bcd3b95118701b46c60f03.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return i045b00b932c574167b22217b96e68e85ba6f525a94bcd3b95118701b46c60f03.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
