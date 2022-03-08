package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i2ea8718568dd68e4db5b0c9b13edefe941431dfb1541af1606f4bc0a7d2db21c "github.com/microsoftgraph/msgraph-sdk-go/me/chats/item/installedapps"
    i88a4f2f2a7c4ae16c306dfc70fe5cfa8f9175e29138583b1d60f25decd49e6d9 "github.com/microsoftgraph/msgraph-sdk-go/me/chats/item/messages"
    ibfa184694112defc7dad34dc61131fc41730c513148b28b61ceae4b917fc813b "github.com/microsoftgraph/msgraph-sdk-go/me/chats/item/members"
    id387b678bda7815af798dc5b19d04d6775c9e30340b77735c65b245e5bae6867 "github.com/microsoftgraph/msgraph-sdk-go/me/chats/item/tabs"
    i25a5bc17ba5983e7d43d8a128bba975fd3706b2f2204b481e7864fe0ae9d5030 "github.com/microsoftgraph/msgraph-sdk-go/me/chats/item/tabs/item"
    i2febef273a5d29b65dd84845ed6389ff6349f765563d2c8edcf94919833a9f49 "github.com/microsoftgraph/msgraph-sdk-go/me/chats/item/members/item"
    i8d1551e7ddb5563740c721b9a2267c157137532374fa22029cbe1c43d7821a74 "github.com/microsoftgraph/msgraph-sdk-go/me/chats/item/installedapps/item"
    i9c260c1538d5b918af33a76a10a514936b0f1f71f65fedc06df9fbe2d4dc5876 "github.com/microsoftgraph/msgraph-sdk-go/me/chats/item/messages/item"
)

// ChatItemRequestBuilder provides operations to manage the chats property of the microsoft.graph.user entity.
type ChatItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ChatItemRequestBuilderDeleteOptions options for Delete
type ChatItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChatItemRequestBuilderGetOptions options for Get
type ChatItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChatItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChatItemRequestBuilderGetQueryParameters get chats from me
type ChatItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ChatItemRequestBuilderPatchOptions options for Patch
type ChatItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Chatable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewChatItemRequestBuilderInternal instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatItemRequestBuilder) {
    m := &ChatItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/chats/{chat_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChatItemRequestBuilder instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property chats for me
func (m *ChatItemRequestBuilder) CreateDeleteRequestInformation(options *ChatItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get chats from me
func (m *ChatItemRequestBuilder) CreateGetRequestInformation(options *ChatItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property chats in me
func (m *ChatItemRequestBuilder) CreatePatchRequestInformation(options *ChatItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property chats for me
func (m *ChatItemRequestBuilder) Delete(options *ChatItemRequestBuilderDeleteOptions)(error) {
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
// Get get chats from me
func (m *ChatItemRequestBuilder) Get(options *ChatItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Chatable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateChatFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Chatable), nil
}
func (m *ChatItemRequestBuilder) InstalledApps()(*i2ea8718568dd68e4db5b0c9b13edefe941431dfb1541af1606f4bc0a7d2db21c.InstalledAppsRequestBuilder) {
    return i2ea8718568dd68e4db5b0c9b13edefe941431dfb1541af1606f4bc0a7d2db21c.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.chats.item.installedApps.item collection
func (m *ChatItemRequestBuilder) InstalledAppsById(id string)(*i8d1551e7ddb5563740c721b9a2267c157137532374fa22029cbe1c43d7821a74.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation_id"] = id
    }
    return i8d1551e7ddb5563740c721b9a2267c157137532374fa22029cbe1c43d7821a74.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Members()(*ibfa184694112defc7dad34dc61131fc41730c513148b28b61ceae4b917fc813b.MembersRequestBuilder) {
    return ibfa184694112defc7dad34dc61131fc41730c513148b28b61ceae4b917fc813b.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.chats.item.members.item collection
func (m *ChatItemRequestBuilder) MembersById(id string)(*i2febef273a5d29b65dd84845ed6389ff6349f765563d2c8edcf94919833a9f49.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return i2febef273a5d29b65dd84845ed6389ff6349f765563d2c8edcf94919833a9f49.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Messages()(*i88a4f2f2a7c4ae16c306dfc70fe5cfa8f9175e29138583b1d60f25decd49e6d9.MessagesRequestBuilder) {
    return i88a4f2f2a7c4ae16c306dfc70fe5cfa8f9175e29138583b1d60f25decd49e6d9.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.chats.item.messages.item collection
func (m *ChatItemRequestBuilder) MessagesById(id string)(*i9c260c1538d5b918af33a76a10a514936b0f1f71f65fedc06df9fbe2d4dc5876.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return i9c260c1538d5b918af33a76a10a514936b0f1f71f65fedc06df9fbe2d4dc5876.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property chats in me
func (m *ChatItemRequestBuilder) Patch(options *ChatItemRequestBuilderPatchOptions)(error) {
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
func (m *ChatItemRequestBuilder) Tabs()(*id387b678bda7815af798dc5b19d04d6775c9e30340b77735c65b245e5bae6867.TabsRequestBuilder) {
    return id387b678bda7815af798dc5b19d04d6775c9e30340b77735c65b245e5bae6867.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.chats.item.tabs.item collection
func (m *ChatItemRequestBuilder) TabsById(id string)(*i25a5bc17ba5983e7d43d8a128bba975fd3706b2f2204b481e7864fe0ae9d5030.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return i25a5bc17ba5983e7d43d8a128bba975fd3706b2f2204b481e7864fe0ae9d5030.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
