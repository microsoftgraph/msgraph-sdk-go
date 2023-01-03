package me

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ChatsChatItemRequestBuilder provides operations to manage the chats property of the microsoft.graph.user entity.
type ChatsChatItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChatsChatItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatsChatItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ChatsChatItemRequestBuilderGetQueryParameters get chats from me
type ChatsChatItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChatsChatItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatsChatItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChatsChatItemRequestBuilderGetQueryParameters
}
// ChatsChatItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ChatsChatItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewChatsChatItemRequestBuilderInternal instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatsChatItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatsChatItemRequestBuilder) {
    m := &ChatsChatItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/chats/{chat%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChatsChatItemRequestBuilder instantiates a new ChatItemRequestBuilder and sets the default values.
func NewChatsChatItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChatsChatItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChatsChatItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property chats for me
func (m *ChatsChatItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *ChatsChatItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get chats from me
func (m *ChatsChatItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *ChatsChatItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property chats in me
func (m *ChatsChatItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, requestConfiguration *ChatsChatItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property chats for me
func (m *ChatsChatItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ChatsChatItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get chats from me
func (m *ChatsChatItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ChatsChatItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable), nil
}
// HideForUser provides operations to call the hideForUser method.
func (m *ChatsChatItemRequestBuilder) HideForUser()(*ChatsItemHideForUserRequestBuilder) {
    return NewChatsItemHideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) InstalledApps()(*ChatsItemInstalledAppsRequestBuilder) {
    return NewChatsItemInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) InstalledAppsById(id string)(*TeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// LastMessagePreview provides operations to manage the lastMessagePreview property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) LastMessagePreview()(*ChatsItemLastMessagePreviewRequestBuilder) {
    return NewChatsItemLastMessagePreviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MarkChatReadForUser provides operations to call the markChatReadForUser method.
func (m *ChatsChatItemRequestBuilder) MarkChatReadForUser()(*ChatsItemMarkChatReadForUserRequestBuilder) {
    return NewChatsItemMarkChatReadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MarkChatUnreadForUser provides operations to call the markChatUnreadForUser method.
func (m *ChatsChatItemRequestBuilder) MarkChatUnreadForUser()(*ChatsItemMarkChatUnreadForUserRequestBuilder) {
    return NewChatsItemMarkChatUnreadForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Members provides operations to manage the members property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) Members()(*ChatsItemMembersRequestBuilder) {
    return NewChatsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) MembersById(id string)(*ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages provides operations to manage the messages property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) Messages()(*ChatsItemMessagesRequestBuilder) {
    return NewChatsItemMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById provides operations to manage the messages property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) MessagesById(id string)(*ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property chats in me
func (m *ChatsChatItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, requestConfiguration *ChatsChatItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Chatable), nil
}
// PinnedMessages provides operations to manage the pinnedMessages property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) PinnedMessages()(*ChatsItemPinnedMessagesRequestBuilder) {
    return NewChatsItemPinnedMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PinnedMessagesById provides operations to manage the pinnedMessages property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) PinnedMessagesById(id string)(*PinnedChatMessageInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["pinnedChatMessageInfo%2Did"] = id
    }
    return NewPinnedChatMessageInfoItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// SendActivityNotification provides operations to call the sendActivityNotification method.
func (m *ChatsChatItemRequestBuilder) SendActivityNotification()(*ChatsItemSendActivityNotificationRequestBuilder) {
    return NewChatsItemSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) Tabs()(*ChatsItemTabsRequestBuilder) {
    return NewChatsItemTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById provides operations to manage the tabs property of the microsoft.graph.chat entity.
func (m *ChatsChatItemRequestBuilder) TabsById(id string)(*TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// UnhideForUser provides operations to call the unhideForUser method.
func (m *ChatsChatItemRequestBuilder) UnhideForUser()(*ChatsItemUnhideForUserRequestBuilder) {
    return NewChatsItemUnhideForUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
