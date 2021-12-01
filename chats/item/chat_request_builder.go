package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i0453d8f2831f761719a2d96702e4c863ff529132a101dbe684210882f2721d5c "github.com/microsoftgraph/msgraph-sdk-go/chats/item/messages"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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

// ChatRequestBuilder builds and executes requests for operations under \chats\{chat-id}
type ChatRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ChatRequestBuilderDeleteOptions options for Delete
type ChatRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChatRequestBuilderGetOptions options for Get
type ChatRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ChatRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ChatRequestBuilderGetQueryParameters get entity from chats by key
type ChatRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ChatRequestBuilderPatchOptions options for Patch
type ChatRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Chat;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewChatRequestBuilderInternal instantiates a new ChatRequestBuilder and sets the default values.
func NewChatRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatRequestBuilder) {
    m := &ChatRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/chats/{chat_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChatRequestBuilder instantiates a new ChatRequestBuilder and sets the default values.
func NewChatRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from chats
func (m *ChatRequestBuilder) CreateDeleteRequestInformation(options *ChatRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ChatRequestBuilder) CreateGetRequestInformation(options *ChatRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ChatRequestBuilder) CreatePatchRequestInformation(options *ChatRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *ChatRequestBuilder) Delete(options *ChatRequestBuilderDeleteOptions)(error) {
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
// Get get entity from chats by key
func (m *ChatRequestBuilder) Get(options *ChatRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Chat, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewChat() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Chat), nil
}
func (m *ChatRequestBuilder) InstalledApps()(*i6e9937a477483db52c918f11ae44d2375f40192b050c02f9aaa34371ac1a12b8.InstalledAppsRequestBuilder) {
    return i6e9937a477483db52c918f11ae44d2375f40192b050c02f9aaa34371ac1a12b8.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.installedApps.item collection
func (m *ChatRequestBuilder) InstalledAppsById(id string)(*ie48eb983af14f89a84a71353f7c41fc8847c3001c54421e947a07c7e5017ae2b.TeamsAppInstallationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation_id"] = id
    }
    return ie48eb983af14f89a84a71353f7c41fc8847c3001c54421e947a07c7e5017ae2b.NewTeamsAppInstallationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatRequestBuilder) Members()(*i75aa69acd6a8f7bab80f3af14760426f17ce1597f5155737243d02bac7efbc9b.MembersRequestBuilder) {
    return i75aa69acd6a8f7bab80f3af14760426f17ce1597f5155737243d02bac7efbc9b.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.members.item collection
func (m *ChatRequestBuilder) MembersById(id string)(*ia1765451c54b694370a995b8293695056771707eceb99a55ae5f1a319717a9cd.ConversationMemberRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return ia1765451c54b694370a995b8293695056771707eceb99a55ae5f1a319717a9cd.NewConversationMemberRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatRequestBuilder) Messages()(*i0453d8f2831f761719a2d96702e4c863ff529132a101dbe684210882f2721d5c.MessagesRequestBuilder) {
    return i0453d8f2831f761719a2d96702e4c863ff529132a101dbe684210882f2721d5c.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.messages.item collection
func (m *ChatRequestBuilder) MessagesById(id string)(*if01cbefdd30a84f3f8ae04a8237bc738cd6d90ab68b9c2ccb4f8f2e8b18e5f34.ChatMessageRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return if01cbefdd30a84f3f8ae04a8237bc738cd6d90ab68b9c2ccb4f8f2e8b18e5f34.NewChatMessageRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in chats
func (m *ChatRequestBuilder) Patch(options *ChatRequestBuilderPatchOptions)(error) {
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
func (m *ChatRequestBuilder) SendActivityNotification()(*i6feff4d6e3d0767acb48ac48f31e0afd6402b4d1449e0aa23e2bbb0c98addc1a.SendActivityNotificationRequestBuilder) {
    return i6feff4d6e3d0767acb48ac48f31e0afd6402b4d1449e0aa23e2bbb0c98addc1a.NewSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ChatRequestBuilder) Tabs()(*i4829a1a874097739700639a97e70426ca9839240f97e9abd24a5c555929e9d7f.TabsRequestBuilder) {
    return i4829a1a874097739700639a97e70426ca9839240f97e9abd24a5c555929e9d7f.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.tabs.item collection
func (m *ChatRequestBuilder) TabsById(id string)(*i837ff1257c18874e8d41607d659eb84e1e4dde73a2ec63b89aa3896598c0d0b1.TeamsTabRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return i837ff1257c18874e8d41607d659eb84e1e4dde73a2ec63b89aa3896598c0d0b1.NewTeamsTabRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
