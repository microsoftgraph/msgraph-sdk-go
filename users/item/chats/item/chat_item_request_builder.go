package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i161ad19b0c7a99cae8a80283db44d458d311ea7ea0fda7b7bee6231e5a71bf34 "github.com/microsoftgraph/msgraph-sdk-go/users/item/chats/item/installedapps"
    i29250e65d3047fac7086c3da168f1dc97b716e57040772b108c92eb7e41e5082 "github.com/microsoftgraph/msgraph-sdk-go/users/item/chats/item/members"
    i747f2653d0182726c4364050b88c4c017361f81ae437c821457319c34de8bbe4 "github.com/microsoftgraph/msgraph-sdk-go/users/item/chats/item/messages"
    ic797b970aa8f13def47c8a9daa30e882f6b3ffba8f376b46454ecdca4bbfff2a "github.com/microsoftgraph/msgraph-sdk-go/users/item/chats/item/tabs"
    i03a5d9d6f43c34754e384f5d06895085102259e8c0fd8d72b735c4f7e27e19e3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/chats/item/members/item"
    i89204df365cb0b4df35e680ee9e843167bce6597538b2e059b21aa0890ed91ae "github.com/microsoftgraph/msgraph-sdk-go/users/item/chats/item/tabs/item"
    ia5fe7898b11b622f06cabbbf98b672db477f4cd5aa1bdc1f1091441663a7ac94 "github.com/microsoftgraph/msgraph-sdk-go/users/item/chats/item/messages/item"
    ice5d918650d80a2fd00a96c09c41f928c4c75fc531425a6ae4f4cc396b70c4f2 "github.com/microsoftgraph/msgraph-sdk-go/users/item/chats/item/installedapps/item"
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
// ChatItemRequestBuilderGetQueryParameters get chats from users
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
    m.urlTemplate = "{+baseurl}/users/{user_id}/chats/{chat_id}{?select,expand}";
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
// CreateDeleteRequestInformation delete navigation property chats for users
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
// CreateGetRequestInformation get chats from users
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
// CreatePatchRequestInformation update the navigation property chats in users
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
// Delete delete navigation property chats for users
func (m *ChatItemRequestBuilder) Delete(options *ChatItemRequestBuilderDeleteOptions)(error) {
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
// Get get chats from users
func (m *ChatItemRequestBuilder) Get(options *ChatItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Chatable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateChatFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Chatable), nil
}
func (m *ChatItemRequestBuilder) InstalledApps()(*i161ad19b0c7a99cae8a80283db44d458d311ea7ea0fda7b7bee6231e5a71bf34.InstalledAppsRequestBuilder) {
    return i161ad19b0c7a99cae8a80283db44d458d311ea7ea0fda7b7bee6231e5a71bf34.NewInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.chats.item.installedApps.item collection
func (m *ChatItemRequestBuilder) InstalledAppsById(id string)(*ice5d918650d80a2fd00a96c09c41f928c4c75fc531425a6ae4f4cc396b70c4f2.TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation_id"] = id
    }
    return ice5d918650d80a2fd00a96c09c41f928c4c75fc531425a6ae4f4cc396b70c4f2.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Members()(*i29250e65d3047fac7086c3da168f1dc97b716e57040772b108c92eb7e41e5082.MembersRequestBuilder) {
    return i29250e65d3047fac7086c3da168f1dc97b716e57040772b108c92eb7e41e5082.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.chats.item.members.item collection
func (m *ChatItemRequestBuilder) MembersById(id string)(*i03a5d9d6f43c34754e384f5d06895085102259e8c0fd8d72b735c4f7e27e19e3.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember_id"] = id
    }
    return i03a5d9d6f43c34754e384f5d06895085102259e8c0fd8d72b735c4f7e27e19e3.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ChatItemRequestBuilder) Messages()(*i747f2653d0182726c4364050b88c4c017361f81ae437c821457319c34de8bbe4.MessagesRequestBuilder) {
    return i747f2653d0182726c4364050b88c4c017361f81ae437c821457319c34de8bbe4.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.chats.item.messages.item collection
func (m *ChatItemRequestBuilder) MessagesById(id string)(*ia5fe7898b11b622f06cabbbf98b672db477f4cd5aa1bdc1f1091441663a7ac94.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage_id"] = id
    }
    return ia5fe7898b11b622f06cabbbf98b672db477f4cd5aa1bdc1f1091441663a7ac94.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property chats in users
func (m *ChatItemRequestBuilder) Patch(options *ChatItemRequestBuilderPatchOptions)(error) {
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
func (m *ChatItemRequestBuilder) Tabs()(*ic797b970aa8f13def47c8a9daa30e882f6b3ffba8f376b46454ecdca4bbfff2a.TabsRequestBuilder) {
    return ic797b970aa8f13def47c8a9daa30e882f6b3ffba8f376b46454ecdca4bbfff2a.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.chats.item.tabs.item collection
func (m *ChatItemRequestBuilder) TabsById(id string)(*i89204df365cb0b4df35e680ee9e843167bce6597538b2e059b21aa0890ed91ae.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab_id"] = id
    }
    return i89204df365cb0b4df35e680ee9e843167bce6597538b2e059b21aa0890ed91ae.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
