package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i52508fdeacdff04c86ef6ed98d578ffd2c77d62a7edfb160198322b9a67e8e0a "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/channels/item/tabs"
    i793c23296f16d9688a86912d15fe2bf5d6212cd55a992ac7d34ce9ea17c9d8cb "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/channels/item/provisionemail"
    i99487f33782211bac82941f8c49dfba40147399ed21891f1ad74a96382992ae3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/channels/item/completemigration"
    ib0c11d6a85aba358cbcfc13c9be3902096670bb120397dd26d3faa62652e2524 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/channels/item/filesfolder"
    ic775254656b39a9ea58d56de7e743ecea61bb0555fd277cbb8e893134cdc519e "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/channels/item/members"
    ie5ca968c3c0027bf6b121a128bc8b2a3b8c58e61b6f1a6ccaa044003449557cf "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/channels/item/removeemail"
    if0b23a5506431d362cbdca425249e3b64c69fb21596d086069afe81af11c676d "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/channels/item/messages"
    i51bfbcbb97372642de8f2ed3fe79c93955eafb6f6ba3bfe2532d49160430b57d "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/channels/item/members/item"
    i7aaac21246fdaa52fcd18af794bbc834c1e127facfe4e64e697ea1931cbe2496 "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/channels/item/tabs/item"
    ia6c1a592dc0444820daee3899e4be3c57cdf1079f6c5d508460448df15f9926d "github.com/microsoftgraph/msgraph-sdk-go/users/item/joinedteams/item/channels/item/messages/item"
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
func (m *ChannelItemRequestBuilder) CompleteMigration()(*i99487f33782211bac82941f8c49dfba40147399ed21891f1ad74a96382992ae3.CompleteMigrationRequestBuilder) {
    return i99487f33782211bac82941f8c49dfba40147399ed21891f1ad74a96382992ae3.NewCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewChannelItemRequestBuilderInternal instantiates a new ChannelItemRequestBuilder and sets the default values.
func NewChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ChannelItemRequestBuilder) {
    m := &ChannelItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/channels/{channel%2Did}{?%24select,%24expand}";
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
// CreateDeleteRequestInformation delete navigation property channels for users
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
// CreatePatchRequestInformation update the navigation property channels in users
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
// Delete delete navigation property channels for users
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
func (m *ChannelItemRequestBuilder) FilesFolder()(*ib0c11d6a85aba358cbcfc13c9be3902096670bb120397dd26d3faa62652e2524.FilesFolderRequestBuilder) {
    return ib0c11d6a85aba358cbcfc13c9be3902096670bb120397dd26d3faa62652e2524.NewFilesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
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
func (m *ChannelItemRequestBuilder) Members()(*ic775254656b39a9ea58d56de7e743ecea61bb0555fd277cbb8e893134cdc519e.MembersRequestBuilder) {
    return ic775254656b39a9ea58d56de7e743ecea61bb0555fd277cbb8e893134cdc519e.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.channels.item.members.item collection
func (m *ChannelItemRequestBuilder) MembersById(id string)(*i51bfbcbb97372642de8f2ed3fe79c93955eafb6f6ba3bfe2532d49160430b57d.ConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return i51bfbcbb97372642de8f2ed3fe79c93955eafb6f6ba3bfe2532d49160430b57d.NewConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Messages the messages property
func (m *ChannelItemRequestBuilder) Messages()(*if0b23a5506431d362cbdca425249e3b64c69fb21596d086069afe81af11c676d.MessagesRequestBuilder) {
    return if0b23a5506431d362cbdca425249e3b64c69fb21596d086069afe81af11c676d.NewMessagesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MessagesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.channels.item.messages.item collection
func (m *ChannelItemRequestBuilder) MessagesById(id string)(*ia6c1a592dc0444820daee3899e4be3c57cdf1079f6c5d508460448df15f9926d.ChatMessageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["chatMessage%2Did"] = id
    }
    return ia6c1a592dc0444820daee3899e4be3c57cdf1079f6c5d508460448df15f9926d.NewChatMessageItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property channels in users
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
func (m *ChannelItemRequestBuilder) ProvisionEmail()(*i793c23296f16d9688a86912d15fe2bf5d6212cd55a992ac7d34ce9ea17c9d8cb.ProvisionEmailRequestBuilder) {
    return i793c23296f16d9688a86912d15fe2bf5d6212cd55a992ac7d34ce9ea17c9d8cb.NewProvisionEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoveEmail the removeEmail property
func (m *ChannelItemRequestBuilder) RemoveEmail()(*ie5ca968c3c0027bf6b121a128bc8b2a3b8c58e61b6f1a6ccaa044003449557cf.RemoveEmailRequestBuilder) {
    return ie5ca968c3c0027bf6b121a128bc8b2a3b8c58e61b6f1a6ccaa044003449557cf.NewRemoveEmailRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tabs the tabs property
func (m *ChannelItemRequestBuilder) Tabs()(*i52508fdeacdff04c86ef6ed98d578ffd2c77d62a7edfb160198322b9a67e8e0a.TabsRequestBuilder) {
    return i52508fdeacdff04c86ef6ed98d578ffd2c77d62a7edfb160198322b9a67e8e0a.NewTabsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TabsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.joinedTeams.item.channels.item.tabs.item collection
func (m *ChannelItemRequestBuilder) TabsById(id string)(*i7aaac21246fdaa52fcd18af794bbc834c1e127facfe4e64e697ea1931cbe2496.TeamsTabItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsTab%2Did"] = id
    }
    return i7aaac21246fdaa52fcd18af794bbc834c1e127facfe4e64e697ea1931cbe2496.NewTeamsTabItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
