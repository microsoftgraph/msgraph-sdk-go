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

// Builds and executes requests for operations under \chats\{chat-id}
type ChatRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Get entity from chats by key
type ChatRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Instantiates a new ChatRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewChatRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatRequestBuilder) {
    m := &ChatRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/chats/{chat_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ChatRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewChatRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ChatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete entity from chats
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *ChatRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get entity from chats by key
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *ChatRequestBuilder) CreateGetRequestInformation(q func (value *ChatRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(ChatRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Update entity in chats
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *ChatRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Chat, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete entity from chats
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ChatRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
// Get entity from chats by key
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ChatRequestBuilder) Get(q func (value *ChatRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Chat, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewChat() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Chat), nil
}
func (m *ChatRequestBuilder) InstalledApps()(*i6e9937a477483db52c918f11ae44d2375f40192b050c02f9aaa34371ac1a12b8.InstalledAppsRequestBuilder) {
    return i6e9937a477483db52c918f11ae44d2375f40192b050c02f9aaa34371ac1a12b8.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.installedApps.item collection
// Parameters:
//  - id : Unique identifier of the item
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
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.members.item collection
// Parameters:
//  - id : Unique identifier of the item
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
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.messages.item collection
// Parameters:
//  - id : Unique identifier of the item
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
// Update entity in chats
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *ChatRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Chat, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
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
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.chats.item.tabs.item collection
// Parameters:
//  - id : Unique identifier of the item
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
