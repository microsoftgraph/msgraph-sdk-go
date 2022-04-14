package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
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
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChatItemRequestBuilderDeleteOptions options for Delete
type ChatItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChatItemRequestBuilderGetOptions options for Get
type ChatItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChatItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChatItemRequestBuilderGetQueryParameters get chats from users
type ChatItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChatItemRequestBuilderPatchOptions options for Patch
type ChatItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// NewChatItemRequestBuilderInternal instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatItemRequestBuilder) {
    m := &ChatItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChatItemRequestBuilder instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property chats for users
func (m *ChatItemRequestBuilder) CreateDeleteRequestInformation(options *ChatItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get chats from users
func (m *ChatItemRequestBuilder) CreateGetRequestInformation(options *ChatItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property chats in users
func (m *ChatItemRequestBuilder) CreatePatchRequestInformation(options *ChatItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
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
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get chats from users
func (m *ChatItemRequestBuilder) Get(options *ChatItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable), nil
}
// InstalledApps the installedApps property
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
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return ice5d918650d80a2fd00a96c09c41f928c4c75fc531425a6ae4f4cc396b70c4f2.NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members the members property
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
        urlTplParams["conversationMember%2Did"] = id
    }
    return i03a5d9d6f43c34754e384f5d06895085102259e8c0fd8d72b735c4f7e27e19e3.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
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
        urlTplParams["chatMessage%2Did"] = id
    }
    return ia5fe7898b11b622f06cabbbf98b672db477f4cd5aa1bdc1f1091441663a7ac94.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property chats in users
func (m *ChatItemRequestBuilder) Patch(options *ChatItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Tabs the tabs property
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
        urlTplParams["teamsTab%2Did"] = id
    }
    return i89204df365cb0b4df35e680ee9e843167bce6597538b2e059b21aa0890ed91ae.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
