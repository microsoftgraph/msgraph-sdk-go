package builders

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TeamsTeamItemRequestBuilder provides operations to manage the collection of team entities.
type TeamsTeamItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// TeamsTeamItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsTeamItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TeamsTeamItemRequestBuilderGetQueryParameters retrieve the properties and relationships of the specified team.
type TeamsTeamItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TeamsTeamItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsTeamItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamsTeamItemRequestBuilderGetQueryParameters
}
// TeamsTeamItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsTeamItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) AllChannels()(*TeamsItemAllChannelsRequestBuilder) {
    return NewTeamsItemAllChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AllChannelsById provides operations to manage the allChannels property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) AllChannelsById(id string)(*TeamsItemAllChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewTeamsItemAllChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Archive provides operations to call the archive method.
func (m *TeamsTeamItemRequestBuilder) Archive()(*TeamsItemArchiveRequestBuilder) {
    return NewTeamsItemArchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) Channels()(*TeamsItemChannelsRequestBuilder) {
    return NewTeamsItemChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChannelsById provides operations to manage the channels property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) ChannelsById(id string)(*TeamsItemChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewTeamsItemChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Clone provides operations to call the clone method.
func (m *TeamsTeamItemRequestBuilder) Clone()(*TeamsItemCloneRequestBuilder) {
    return NewTeamsItemCloneRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CompleteMigration provides operations to call the completeMigration method.
func (m *TeamsTeamItemRequestBuilder) CompleteMigration()(*TeamsItemCompleteMigrationRequestBuilder) {
    return NewTeamsItemCompleteMigrationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewTeamsTeamItemRequestBuilderInternal instantiates a new TeamItemRequestBuilder and sets the default values.
func NewTeamsTeamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsTeamItemRequestBuilder) {
    m := &TeamsTeamItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/teams/{team%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewTeamsTeamItemRequestBuilder instantiates a new TeamItemRequestBuilder and sets the default values.
func NewTeamsTeamItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsTeamItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamsTeamItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from teams
func (m *TeamsTeamItemRequestBuilder) CreateDeleteRequestInformation(ctx context.Context, requestConfiguration *TeamsTeamItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation retrieve the properties and relationships of the specified team.
func (m *TeamsTeamItemRequestBuilder) CreateGetRequestInformation(ctx context.Context, requestConfiguration *TeamsTeamItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the properties of the specified team.
func (m *TeamsTeamItemRequestBuilder) CreatePatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *TeamsTeamItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete entity from teams
func (m *TeamsTeamItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TeamsTeamItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve the properties and relationships of the specified team.
func (m *TeamsTeamItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamsTeamItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
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
func (m *TeamsTeamItemRequestBuilder) Group()(*TeamsItemGroupRequestBuilder) {
    return NewTeamsItemGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) IncomingChannels()(*TeamsItemIncomingChannelsRequestBuilder) {
    return NewTeamsItemIncomingChannelsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IncomingChannelsById provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) IncomingChannelsById(id string)(*TeamsItemIncomingChannelsChannelItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["channel%2Did"] = id
    }
    return NewTeamsItemIncomingChannelsChannelItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) InstalledApps()(*TeamsItemInstalledAppsRequestBuilder) {
    return NewTeamsItemInstalledAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// InstalledAppsById provides operations to manage the installedApps property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) InstalledAppsById(id string)(*TeamsItemInstalledAppsTeamsAppInstallationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAppInstallation%2Did"] = id
    }
    return NewTeamsItemInstalledAppsTeamsAppInstallationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) Members()(*TeamsItemMembersRequestBuilder) {
    return NewTeamsItemMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById provides operations to manage the members property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) MembersById(id string)(*TeamsItemMembersConversationMemberItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["conversationMember%2Did"] = id
    }
    return NewTeamsItemMembersConversationMemberItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) Operations()(*TeamsItemOperationsRequestBuilder) {
    return NewTeamsItemOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// OperationsById provides operations to manage the operations property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) OperationsById(id string)(*TeamsItemOperationsTeamsAsyncOperationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamsAsyncOperation%2Did"] = id
    }
    return NewTeamsItemOperationsTeamsAsyncOperationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the properties of the specified team.
func (m *TeamsTeamItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *TeamsTeamItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
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
func (m *TeamsTeamItemRequestBuilder) Photo()(*TeamsItemPhotoRequestBuilder) {
    return NewTeamsItemPhotoRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) PrimaryChannel()(*TeamsItemPrimaryChannelRequestBuilder) {
    return NewTeamsItemPrimaryChannelRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) Schedule()(*TeamsItemScheduleRequestBuilder) {
    return NewTeamsItemScheduleRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SendActivityNotification provides operations to call the sendActivityNotification method.
func (m *TeamsTeamItemRequestBuilder) SendActivityNotification()(*TeamsItemSendActivityNotificationRequestBuilder) {
    return NewTeamsItemSendActivityNotificationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) Tags()(*TeamsItemTagsRequestBuilder) {
    return NewTeamsItemTagsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TagsById provides operations to manage the tags property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) TagsById(id string)(*TeamsItemTagsTeamworkTagItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["teamworkTag%2Did"] = id
    }
    return NewTeamsItemTagsTeamworkTagItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
func (m *TeamsTeamItemRequestBuilder) Template()(*TeamsItemTemplateRequestBuilder) {
    return NewTeamsItemTemplateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unarchive provides operations to call the unarchive method.
func (m *TeamsTeamItemRequestBuilder) Unarchive()(*TeamsItemUnarchiveRequestBuilder) {
    return NewTeamsItemUnarchiveRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
