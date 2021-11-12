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

// Builds and executes requests for operations under \teams\{team-id}\channels\{channel-id}
type ChannelRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ChannelRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ChannelRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChannelRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The collection of channels & messages associated with the team.
type ChannelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ChannelRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channel;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ChannelRequestBuilder) CompleteMigration()(*id475bce00f5dd8f4fee815fc2fb530ee5f11935ffc33cfa7e9e92780aad132dc.CompleteMigrationRequestBuilder) {
    return id475bce00f5dd8f4fee815fc2fb530ee5f11935ffc33cfa7e9e92780aad132dc.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new ChannelRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChannelRequestBuilder) {
    m := &ChannelRequestBuilder{
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
// Instantiates a new ChannelRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewChannelRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelRequestBuilderInternal(urlParams, requestAdapter)
}
// The collection of channels & messages associated with the team.
// Parameters:
//  - options : Options for the request
func (m *ChannelRequestBuilder) CreateDeleteRequestInformation(options *ChannelRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The collection of channels & messages associated with the team.
// Parameters:
//  - options : Options for the request
func (m *ChannelRequestBuilder) CreateGetRequestInformation(options *ChannelRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(options.Q)
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
// The collection of channels & messages associated with the team.
// Parameters:
//  - options : Options for the request
func (m *ChannelRequestBuilder) CreatePatchRequestInformation(options *ChannelRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The collection of channels & messages associated with the team.
// Parameters:
//  - options : Options for the request
func (m *ChannelRequestBuilder) Delete(options *ChannelRequestBuilderDeleteOptions)(error) {
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
func (m *ChannelRequestBuilder) FilesFolder()(*i9d35966f1cf1d79964b04fc579a16ad9d6fde5d5b6ea86ce1147f94ab0e92ab7.FilesFolderRequestBuilder) {
    return i9d35966f1cf1d79964b04fc579a16ad9d6fde5d5b6ea86ce1147f94ab0e92ab7.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of channels & messages associated with the team.
// Parameters:
//  - options : Options for the request
func (m *ChannelRequestBuilder) Get(options *ChannelRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channel, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewChannel() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Channel), nil
}
func (m *ChannelRequestBuilder) Members()(*ib1c6ce99382e50e6c2a17723af8d800e520796e30cdda5ad72c456a4be0287a3.MembersRequestBuilder) {
    return ib1c6ce99382e50e6c2a17723af8d800e520796e30cdda5ad72c456a4be0287a3.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.channels.item.members.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ChannelRequestBuilder) MembersById(id string)(*ifbf820637987673b919966137b926fb0a0b7c26841e2f3e1a8a08762ca2537c0.ConversationMemberRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return ifbf820637987673b919966137b926fb0a0b7c26841e2f3e1a8a08762ca2537c0.NewConversationMemberRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChannelRequestBuilder) Messages()(*i5b2a4d2eaabfa3faa34dc6898f8cb9e57e0b1242b01f641ae9e0207836fe0ad4.MessagesRequestBuilder) {
    return i5b2a4d2eaabfa3faa34dc6898f8cb9e57e0b1242b01f641ae9e0207836fe0ad4.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.channels.item.messages.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ChannelRequestBuilder) MessagesById(id string)(*if5725a60ca0d27f7ade8f6fcec0dcfe3a4c988fdc95238d1b2bc5148c05df6f3.ChatMessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return if5725a60ca0d27f7ade8f6fcec0dcfe3a4c988fdc95238d1b2bc5148c05df6f3.NewChatMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// The collection of channels & messages associated with the team.
// Parameters:
//  - options : Options for the request
func (m *ChannelRequestBuilder) Patch(options *ChannelRequestBuilderPatchOptions)(error) {
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
func (m *ChannelRequestBuilder) ProvisionEmail()(*i6917a550a4991ea62c7afce470df64fd772e953905bf19e5d14225e8a7d388df.ProvisionEmailRequestBuilder) {
    return i6917a550a4991ea62c7afce470df64fd772e953905bf19e5d14225e8a7d388df.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChannelRequestBuilder) RemoveEmail()(*i674af80ba48ecc53d2978ae46338b8a9f32a25fed5f35dde4ba5bc959577bbcb.RemoveEmailRequestBuilder) {
    return i674af80ba48ecc53d2978ae46338b8a9f32a25fed5f35dde4ba5bc959577bbcb.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChannelRequestBuilder) Tabs()(*i24968b403d6ddad22f24d63ea935950ae5ca327471cd684f87d172b9041a8b41.TabsRequestBuilder) {
    return i24968b403d6ddad22f24d63ea935950ae5ca327471cd684f87d172b9041a8b41.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.teams.item.channels.item.tabs.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ChannelRequestBuilder) TabsById(id string)(*ida5ab63d8e96fe8dfad94c2576765817725271d852ddd41b2ada05e5cc7d6e49.TeamsTabRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return ida5ab63d8e96fe8dfad94c2576765817725271d852ddd41b2ada05e5cc7d6e49.NewTeamsTabRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
