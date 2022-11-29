package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// UsersItemJoinedTeamsTeamItemRequestBuilder provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
type UsersItemJoinedTeamsTeamItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// UsersItemJoinedTeamsTeamItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemJoinedTeamsTeamItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// UsersItemJoinedTeamsTeamItemRequestBuilderGetQueryParameters get joinedTeams from users
type UsersItemJoinedTeamsTeamItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// UsersItemJoinedTeamsTeamItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemJoinedTeamsTeamItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *UsersItemJoinedTeamsTeamItemRequestBuilderGetQueryParameters
}
// UsersItemJoinedTeamsTeamItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type UsersItemJoinedTeamsTeamItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) AllChannels()(*UsersItemJoinedTeamsItemAllChannelsRequestBuilder) {
    return NewUsersItemJoinedTeamsItemAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllChannelsById provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) AllChannelsById(id string)(*UsersItemJoinedTeamsItemAllChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewUsersItemJoinedTeamsItemAllChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Archive provides operations to call the archive method.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Archive()(*UsersItemJoinedTeamsItemArchiveRequestBuilder) {
    return NewUsersItemJoinedTeamsItemArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Channels()(*UsersItemJoinedTeamsItemChannelsRequestBuilder) {
    return NewUsersItemJoinedTeamsItemChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) ChannelsById(id string)(*UsersItemJoinedTeamsItemChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewUsersItemJoinedTeamsItemChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Clone provides operations to call the clone method.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Clone()(*UsersItemJoinedTeamsItemCloneRequestBuilder) {
    return NewUsersItemJoinedTeamsItemCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteMigration provides operations to call the completeMigration method.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) CompleteMigration()(*UsersItemJoinedTeamsItemCompleteMigrationRequestBuilder) {
    return NewUsersItemJoinedTeamsItemCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewUsersItemJoinedTeamsTeamItemRequestBuilderInternal instantiates a new TeamItemRequestBuilder and sets the default values.
func NewUsersItemJoinedTeamsTeamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemJoinedTeamsTeamItemRequestBuilder) {
    m := &UsersItemJoinedTeamsTeamItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewUsersItemJoinedTeamsTeamItemRequestBuilder instantiates a new TeamItemRequestBuilder and sets the default values.
func NewUsersItemJoinedTeamsTeamItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*UsersItemJoinedTeamsTeamItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewUsersItemJoinedTeamsTeamItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property joinedTeams for users
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *UsersItemJoinedTeamsTeamItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get joinedTeams from users
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *UsersItemJoinedTeamsTeamItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property joinedTeams in users
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *UsersItemJoinedTeamsTeamItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers["Accept"] = "application/json"
    requestInfo.SetContentFromParsable(ctx, m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property joinedTeams for users
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *UsersItemJoinedTeamsTeamItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get joinedTeams from users
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Get(ctx context.Context, requestConfiguration *UsersItemJoinedTeamsTeamItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
    requestInfo, err := m.CreateGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Group()(*UsersItemJoinedTeamsItemGroupRequestBuilder) {
    return NewUsersItemJoinedTeamsItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) IncomingChannels()(*UsersItemJoinedTeamsItemIncomingChannelsRequestBuilder) {
    return NewUsersItemJoinedTeamsItemIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannelsById provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) IncomingChannelsById(id string)(*UsersItemJoinedTeamsItemIncomingChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewUsersItemJoinedTeamsItemIncomingChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) InstalledApps()(*UsersItemJoinedTeamsItemInstalledAppsRequestBuilder) {
    return NewUsersItemJoinedTeamsItemInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) InstalledAppsById(id string)(*UsersItemJoinedTeamsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewUsersItemJoinedTeamsItemInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Members()(*UsersItemJoinedTeamsItemMembersRequestBuilder) {
    return NewUsersItemJoinedTeamsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) MembersById(id string)(*UsersItemJoinedTeamsItemMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewUsersItemJoinedTeamsItemMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Operations()(*UsersItemJoinedTeamsItemOperationsRequestBuilder) {
    return NewUsersItemJoinedTeamsItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) OperationsById(id string)(*UsersItemJoinedTeamsItemOperationsTeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return NewUsersItemJoinedTeamsItemOperationsTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property joinedTeams in users
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *UsersItemJoinedTeamsTeamItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
    requestInfo, err := m.CreatePatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable), nil
}
// Photo provides operations to manage the photo property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Photo()(*UsersItemJoinedTeamsItemPhotoRequestBuilder) {
    return NewUsersItemJoinedTeamsItemPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) PrimaryChannel()(*UsersItemJoinedTeamsItemPrimaryChannelRequestBuilder) {
    return NewUsersItemJoinedTeamsItemPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Schedule()(*UsersItemJoinedTeamsItemScheduleRequestBuilder) {
    return NewUsersItemJoinedTeamsItemScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendActivityNotification provides operations to call the sendActivityNotification method.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) SendActivityNotification()(*UsersItemJoinedTeamsItemSendActivityNotificationRequestBuilder) {
    return NewUsersItemJoinedTeamsItemSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Tags()(*UsersItemJoinedTeamsItemTagsRequestBuilder) {
    return NewUsersItemJoinedTeamsItemTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) TagsById(id string)(*UsersItemJoinedTeamsItemTagsTeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return NewUsersItemJoinedTeamsItemTagsTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Template()(*UsersItemJoinedTeamsItemTemplateRequestBuilder) {
    return NewUsersItemJoinedTeamsItemTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unarchive provides operations to call the unarchive method.
func (m *UsersItemJoinedTeamsTeamItemRequestBuilder) Unarchive()(*UsersItemJoinedTeamsItemUnarchiveRequestBuilder) {
    return NewUsersItemJoinedTeamsItemUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
