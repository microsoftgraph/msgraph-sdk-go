package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i24968b403d6ddad22f24d63ea935950ae5ca327471cd684f87d172b9041a8b41 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/channels/item/tabs"
    i5b2a4d2eaabfa3faa34dc6898f8cb9e57e0b1242b01f641ae9e0207836fe0ad4 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/channels/item/messages"
    i674af80ba48ecc53d2978ae46338b8a9f32a25fed5f35dde4ba5bc959577bbcb "github.com/microsoftgraph/msgraph-sdk-go/teams/item/channels/item/removeemail"
    i6917a550a4991ea62c7afce470df64fd772e953905bf19e5d14225e8a7d388df "github.com/microsoftgraph/msgraph-sdk-go/teams/item/channels/item/provisionemail"
    i9d35966f1cf1d79964b04fc579a16ad9d6fde5d5b6ea86ce1147f94ab0e92ab7 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/channels/item/filesfolder"
    ib1c6ce99382e50e6c2a17723af8d800e520796e30cdda5ad72c456a4be0287a3 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/channels/item/members"
    id475bce00f5dd8f4fee815fc2fb530ee5f11935ffc33cfa7e9e92780aad132dc "github.com/microsoftgraph/msgraph-sdk-go/teams/item/channels/item/completemigration"
    ida5ab63d8e96fe8dfad94c2576765817725271d852ddd41b2ada05e5cc7d6e49 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/channels/item/tabs/item"
    if5725a60ca0d27f7ade8f6fcec0dcfe3a4c988fdc95238d1b2bc5148c05df6f3 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/channels/item/messages/item"
    ifbf820637987673b919966137b926fb0a0b7c26841e2f3e1a8a08762ca2537c0 "github.com/microsoftgraph/msgraph-sdk-go/teams/item/channels/item/members/item"
)

// ChannelItemRequestBuilder builds and executes requests for operations under \teams\{team-id}\channels\{channel-id}
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
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channel;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ChannelItemRequestBuilder) CompleteMigration()(*id475bce00f5dd8f4fee815fc2fb530ee5f11935ffc33cfa7e9e92780aad132dc.CompleteMigrationRequestBuilder) {
    return id475bce00f5dd8f4fee815fc2fb530ee5f11935ffc33cfa7e9e92780aad132dc.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChannelItemRequestBuilder) {
    m := &ChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team_id}/channels/{channel_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChannelItemRequestBuilder instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation the collection of channels and messages associated with the team.
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
// CreatePatchRequestInformation the collection of channels and messages associated with the team.
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
// Delete the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) Delete(options *ChannelItemRequestBuilderDeleteOptions)(error) {
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
func (m *ChannelItemRequestBuilder) FilesFolder()(*i9d35966f1cf1d79964b04fc579a16ad9d6fde5d5b6ea86ce1147f94ab0e92ab7.FilesFolderRequestBuilder) {
    return i9d35966f1cf1d79964b04fc579a16ad9d6fde5d5b6ea86ce1147f94ab0e92ab7.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) Get(options *ChannelItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channel, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewChannel() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channel), nil
}
func (m *ChannelItemRequestBuilder) Members()(*ib1c6ce99382e50e6c2a17723af8d800e520796e30cdda5ad72c456a4be0287a3.MembersRequestBuilder) {
    return ib1c6ce99382e50e6c2a17723af8d800e520796e30cdda5ad72c456a4be0287a3.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.channels.item.members.item collection
func (m *ChannelItemRequestBuilder) MembersById(id string)(*ifbf820637987673b919966137b926fb0a0b7c26841e2f3e1a8a08762ca2537c0.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return ifbf820637987673b919966137b926fb0a0b7c26841e2f3e1a8a08762ca2537c0.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChannelItemRequestBuilder) Messages()(*i5b2a4d2eaabfa3faa34dc6898f8cb9e57e0b1242b01f641ae9e0207836fe0ad4.MessagesRequestBuilder) {
    return i5b2a4d2eaabfa3faa34dc6898f8cb9e57e0b1242b01f641ae9e0207836fe0ad4.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.channels.item.messages.item collection
func (m *ChannelItemRequestBuilder) MessagesById(id string)(*if5725a60ca0d27f7ade8f6fcec0dcfe3a4c988fdc95238d1b2bc5148c05df6f3.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return if5725a60ca0d27f7ade8f6fcec0dcfe3a4c988fdc95238d1b2bc5148c05df6f3.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) Patch(options *ChannelItemRequestBuilderPatchOptions)(error) {
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
func (m *ChannelItemRequestBuilder) ProvisionEmail()(*i6917a550a4991ea62c7afce470df64fd772e953905bf19e5d14225e8a7d388df.ProvisionEmailRequestBuilder) {
    return i6917a550a4991ea62c7afce470df64fd772e953905bf19e5d14225e8a7d388df.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChannelItemRequestBuilder) RemoveEmail()(*i674af80ba48ecc53d2978ae46338b8a9f32a25fed5f35dde4ba5bc959577bbcb.RemoveEmailRequestBuilder) {
    return i674af80ba48ecc53d2978ae46338b8a9f32a25fed5f35dde4ba5bc959577bbcb.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChannelItemRequestBuilder) Tabs()(*i24968b403d6ddad22f24d63ea935950ae5ca327471cd684f87d172b9041a8b41.TabsRequestBuilder) {
    return i24968b403d6ddad22f24d63ea935950ae5ca327471cd684f87d172b9041a8b41.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.channels.item.tabs.item collection
func (m *ChannelItemRequestBuilder) TabsById(id string)(*ida5ab63d8e96fe8dfad94c2576765817725271d852ddd41b2ada05e5cc7d6e49.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return ida5ab63d8e96fe8dfad94c2576765817725271d852ddd41b2ada05e5cc7d6e49.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
