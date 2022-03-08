package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0453d8f2831f761719a2d96702e4c863ff529132a101dbe684210882f2721d5c "github.com/microsoftgraph/msgraph-sdk-go/chats/item/messages"
    i4829a1a874097739700639a97e70426ca9839240f97e9abd24a5c555929e9d7f "github.com/microsoftgraph/msgraph-sdk-go/chats/item/tabs"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i6e9937a477483db52c918f11ae44d2375f40192b050c02f9aaa34371ac1a12b8 "github.com/microsoftgraph/msgraph-sdk-go/chats/item/installedapps"
    i6feff4d6e3d0767acb48ac48f31e0afd6402b4d1449e0aa23e2bbb0c98addc1a "github.com/microsoftgraph/msgraph-sdk-go/chats/item/sendactivitynotification"
    i75aa69acd6a8f7bab80f3af14760426f17ce1597f5155737243d02bac7efbc9b "github.com/microsoftgraph/msgraph-sdk-go/chats/item/members"
    i837ff1257c18874e8d41607d659eb84e1e4dde73a2ec63b89aa3896598c0d0b1 "github.com/microsoftgraph/msgraph-sdk-go/chats/item/tabs/item"
    ia1765451c54b694370a995b8293695056771707eceb99a55ae5f1a319717a9cd "github.com/microsoftgraph/msgraph-sdk-go/chats/item/members/item"
    ie48eb983af14f89a84a71353f7c41fc8847c3001c54421e947a07c7e5017ae2b "github.com/microsoftgraph/msgraph-sdk-go/chats/item/installedapps/item"
    if01cbefdd30a84f3f8ae04a8237bc738cd6d90ab68b9c2ccb4f8f2e8b18e5f34 "github.com/microsoftgraph/msgraph-sdk-go/chats/item/messages/item"
)

// ChatItemRequestBuilder provides operations to manage the collection of chat entities.
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
// ChatItemRequestBuilderGetQueryParameters get entity from chats by key
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
    m.urlTemplate = "{+baseurl}/chats/{chat_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete entity from chats
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
// CreateGetRequestInformation get entity from chats by key
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
// CreatePatchRequestInformation update entity in chats
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
// Delete delete entity from chats
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
// Get get entity from chats by key
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
func (m *ChatItemRequestBuilder) InstalledApps()(*i6e9937a477483db52c918f11ae44d2375f40192b050c02f9aaa34371ac1a12b8.InstalledAppsRequestBuilder) {
    return i6e9937a477483db52c918f11ae44d2375f40192b050c02f9aaa34371ac1a12b8.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.installedApps.item collection
func (m *ChatItemRequestBuilder) InstalledAppsById(id string)(*ie48eb983af14f89a84a71353f7c41fc8847c3001c54421e947a07c7e5017ae2b.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation_id"] = id
    }
    return ie48eb983af14f89a84a71353f7c41fc8847c3001c54421e947a07c7e5017ae2b.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Members()(*i75aa69acd6a8f7bab80f3af14760426f17ce1597f5155737243d02bac7efbc9b.MembersRequestBuilder) {
    return i75aa69acd6a8f7bab80f3af14760426f17ce1597f5155737243d02bac7efbc9b.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.members.item collection
func (m *ChatItemRequestBuilder) MembersById(id string)(*ia1765451c54b694370a995b8293695056771707eceb99a55ae5f1a319717a9cd.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return ia1765451c54b694370a995b8293695056771707eceb99a55ae5f1a319717a9cd.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Messages()(*i0453d8f2831f761719a2d96702e4c863ff529132a101dbe684210882f2721d5c.MessagesRequestBuilder) {
    return i0453d8f2831f761719a2d96702e4c863ff529132a101dbe684210882f2721d5c.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.messages.item collection
func (m *ChatItemRequestBuilder) MessagesById(id string)(*if01cbefdd30a84f3f8ae04a8237bc738cd6d90ab68b9c2ccb4f8f2e8b18e5f34.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return if01cbefdd30a84f3f8ae04a8237bc738cd6d90ab68b9c2ccb4f8f2e8b18e5f34.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in chats
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
func (m *ChatItemRequestBuilder) SendActivityNotification()(*i6feff4d6e3d0767acb48ac48f31e0afd6402b4d1449e0aa23e2bbb0c98addc1a.SendActivityNotificationRequestBuilder) {
    return i6feff4d6e3d0767acb48ac48f31e0afd6402b4d1449e0aa23e2bbb0c98addc1a.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Tabs()(*i4829a1a874097739700639a97e70426ca9839240f97e9abd24a5c555929e9d7f.TabsRequestBuilder) {
    return i4829a1a874097739700639a97e70426ca9839240f97e9abd24a5c555929e9d7f.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.tabs.item collection
func (m *ChatItemRequestBuilder) TabsById(id string)(*i837ff1257c18874e8d41607d659eb84e1e4dde73a2ec63b89aa3896598c0d0b1.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return i837ff1257c18874e8d41607d659eb84e1e4dde73a2ec63b89aa3896598c0d0b1.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
