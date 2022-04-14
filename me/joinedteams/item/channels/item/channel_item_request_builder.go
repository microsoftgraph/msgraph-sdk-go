package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i253a46a809c827c7298a764f8c82f551eb9c16db15abb743fffe2eba96700378 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels/item/filesfolder"
    i2ad3e3c936b6898b388b6d3cc1c2364d7851f04fd4217471b556b15e62ffff65 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels/item/provisionemail"
    i40fe54e22de0311e21811c5781924675aa201c68840b6af48bd6b80c4ccfc083 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels/item/messages"
    iaf082408b1fd7fd4e8e3c6c1befccb1388486f663eb9ad351a7f8f61d4cf7be6 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels/item/completemigration"
    ie326949afcb08e6a5436a367b925b7f68ec414a2c72fce4b6164a4b239264d52 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels/item/tabs"
    ief0cacbd8be774a1502106b209117815bb3293f805cb67532513c01f04e36513 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels/item/members"
    iefaaa288cea77216f47ed2ce48474ac53a86f88b057403e15cbe3142e3187502 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels/item/removeemail"
    i2da7ef6ccb4db37d9d9fe954811ee5ddad198aa369fa874d4f6099c7f62e0bb6 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels/item/messages/item"
    i3a81824602dda13e4a5b75e20ede850d39a8762eb95c95f58b595895ef98be06 "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels/item/tabs/item"
    if78b603399676ac0a458d83dbd0e2eb877a642ddfd394dd4905b49017bf3160d "github.com/microsoftgraph/msgraph-sdk-go/me/joinedteams/item/channels/item/members/item"
)

// ChannelItemRequestBuilder provides operations to manage the channels property of the microsoft.graph.team entity.
type ChannelItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ChannelItemRequestBuilderDeleteOptions options for Delete
type ChannelItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChannelItemRequestBuilderGetOptions options for Get
type ChannelItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ChannelItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// ChannelItemRequestBuilderGetQueryParameters the collection of channels and messages associated with the team.
type ChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ChannelItemRequestBuilderPatchOptions options for Patch
type ChannelItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// CompleteMigration the completeMigration property
func (m *ChannelItemRequestBuilder) CompleteMigration()(*iaf082408b1fd7fd4e8e3c6c1befccb1388486f663eb9ad351a7f8f61d4cf7be6.CompleteMigrationRequestBuilder) {
    return iaf082408b1fd7fd4e8e3c6c1befccb1388486f663eb9ad351a7f8f61d4cf7be6.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelItemRequestBuilder) {
    m := &ChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/joinedTeams/{team%2Did}/channels/{channel%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewChannelItemRequestBuilder instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property channels for me
func (m *ChannelItemRequestBuilder) CreateDeleteRequestInformation(options *ChannelItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) CreateGetRequestInformation(options *ChannelItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property channels in me
func (m *ChannelItemRequestBuilder) CreatePatchRequestInformation(options *ChannelItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property channels for me
func (m *ChannelItemRequestBuilder) Delete(options *ChannelItemRequestBuilderDeleteOptions)(error) {
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
// FilesFolder the filesFolder property
func (m *ChannelItemRequestBuilder) FilesFolder()(*i253a46a809c827c7298a764f8c82f551eb9c16db15abb743fffe2eba96700378.FilesFolderRequestBuilder) {
    return i253a46a809c827c7298a764f8c82f551eb9c16db15abb743fffe2eba96700378.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get the collection of channels and messages associated with the team.
func (m *ChannelItemRequestBuilder) Get(options *ChannelItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChannelFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable), nil
}
// Members the members property
func (m *ChannelItemRequestBuilder) Members()(*ief0cacbd8be774a1502106b209117815bb3293f805cb67532513c01f04e36513.MembersRequestBuilder) {
    return ief0cacbd8be774a1502106b209117815bb3293f805cb67532513c01f04e36513.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.channels.item.members.item collection
func (m *ChannelItemRequestBuilder) MembersById(id string)(*if78b603399676ac0a458d83dbd0e2eb877a642ddfd394dd4905b49017bf3160d.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return if78b603399676ac0a458d83dbd0e2eb877a642ddfd394dd4905b49017bf3160d.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *ChannelItemRequestBuilder) Messages()(*i40fe54e22de0311e21811c5781924675aa201c68840b6af48bd6b80c4ccfc083.MessagesRequestBuilder) {
    return i40fe54e22de0311e21811c5781924675aa201c68840b6af48bd6b80c4ccfc083.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.channels.item.messages.item collection
func (m *ChannelItemRequestBuilder) MessagesById(id string)(*i2da7ef6ccb4db37d9d9fe954811ee5ddad198aa369fa874d4f6099c7f62e0bb6.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return i2da7ef6ccb4db37d9d9fe954811ee5ddad198aa369fa874d4f6099c7f62e0bb6.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property channels in me
func (m *ChannelItemRequestBuilder) Patch(options *ChannelItemRequestBuilderPatchOptions)(error) {
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
// ProvisionEmail the provisionEmail property
func (m *ChannelItemRequestBuilder) ProvisionEmail()(*i2ad3e3c936b6898b388b6d3cc1c2364d7851f04fd4217471b556b15e62ffff65.ProvisionEmailRequestBuilder) {
    return i2ad3e3c936b6898b388b6d3cc1c2364d7851f04fd4217471b556b15e62ffff65.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail the removeEmail property
func (m *ChannelItemRequestBuilder) RemoveEmail()(*iefaaa288cea77216f47ed2ce48474ac53a86f88b057403e15cbe3142e3187502.RemoveEmailRequestBuilder) {
    return iefaaa288cea77216f47ed2ce48474ac53a86f88b057403e15cbe3142e3187502.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tabs the tabs property
func (m *ChannelItemRequestBuilder) Tabs()(*ie326949afcb08e6a5436a367b925b7f68ec414a2c72fce4b6164a4b239264d52.TabsRequestBuilder) {
    return ie326949afcb08e6a5436a367b925b7f68ec414a2c72fce4b6164a4b239264d52.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.joinedTeams.item.channels.item.tabs.item collection
func (m *ChannelItemRequestBuilder) TabsById(id string)(*i3a81824602dda13e4a5b75e20ede850d39a8762eb95c95f58b595895ef98be06.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return i3a81824602dda13e4a5b75e20ede850d39a8762eb95c95f58b595895ef98be06.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
