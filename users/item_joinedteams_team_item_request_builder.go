package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsTeamItemRequestBuilder provides operations to manage the joinedTeams property of the microsoft.graph.user entity.
type ItemJoinedteamsTeamItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsTeamItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsTeamItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemJoinedteamsTeamItemRequestBuilderGetQueryParameters get joinedTeams from users
type ItemJoinedteamsTeamItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemJoinedteamsTeamItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsTeamItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedteamsTeamItemRequestBuilderGetQueryParameters
}
// ItemJoinedteamsTeamItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsTeamItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AllChannels provides operations to manage the allChannels property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) AllChannels()(*ItemJoinedteamsItemAllchannelsAllChannelsRequestBuilder) {
    return NewItemJoinedteamsItemAllchannelsAllChannelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Archive provides operations to call the archive method.
// returns a *ItemJoinedteamsItemArchiveRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) Archive()(*ItemJoinedteamsItemArchiveRequestBuilder) {
    return NewItemJoinedteamsItemArchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Channels provides operations to manage the channels property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemChannelsRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) Channels()(*ItemJoinedteamsItemChannelsRequestBuilder) {
    return NewItemJoinedteamsItemChannelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Clone provides operations to call the clone method.
// returns a *ItemJoinedteamsItemCloneRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) Clone()(*ItemJoinedteamsItemCloneRequestBuilder) {
    return NewItemJoinedteamsItemCloneRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CompleteMigration provides operations to call the completeMigration method.
// returns a *ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) CompleteMigration()(*ItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilder) {
    return NewItemJoinedteamsItemCompletemigrationCompleteMigrationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemJoinedteamsTeamItemRequestBuilderInternal instantiates a new ItemJoinedteamsTeamItemRequestBuilder and sets the default values.
func NewItemJoinedteamsTeamItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsTeamItemRequestBuilder) {
    m := &ItemJoinedteamsTeamItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemJoinedteamsTeamItemRequestBuilder instantiates a new ItemJoinedteamsTeamItemRequestBuilder and sets the default values.
func NewItemJoinedteamsTeamItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsTeamItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsTeamItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property joinedTeams for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsTeamItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemJoinedteamsTeamItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get joinedTeams from users
// returns a Teamable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsTeamItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedteamsTeamItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable), nil
}
// Group provides operations to manage the group property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemGroupRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) Group()(*ItemJoinedteamsItemGroupRequestBuilder) {
    return NewItemJoinedteamsItemGroupRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// IncomingChannels provides operations to manage the incomingChannels property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemIncomingchannelsIncomingChannelsRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) IncomingChannels()(*ItemJoinedteamsItemIncomingchannelsIncomingChannelsRequestBuilder) {
    return NewItemJoinedteamsItemIncomingchannelsIncomingChannelsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// InstalledApps provides operations to manage the installedApps property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemInstalledappsInstalledAppsRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) InstalledApps()(*ItemJoinedteamsItemInstalledappsInstalledAppsRequestBuilder) {
    return NewItemJoinedteamsItemInstalledappsInstalledAppsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Members provides operations to manage the members property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemMembersRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) Members()(*ItemJoinedteamsItemMembersRequestBuilder) {
    return NewItemJoinedteamsItemMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemOperationsRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) Operations()(*ItemJoinedteamsItemOperationsRequestBuilder) {
    return NewItemJoinedteamsItemOperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property joinedTeams in users
// returns a Teamable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsTeamItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *ItemJoinedteamsTeamItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable), nil
}
// PermissionGrants provides operations to manage the permissionGrants property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemPermissiongrantsPermissionGrantsRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) PermissionGrants()(*ItemJoinedteamsItemPermissiongrantsPermissionGrantsRequestBuilder) {
    return NewItemJoinedteamsItemPermissiongrantsPermissionGrantsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Photo provides operations to manage the photo property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemPhotoRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) Photo()(*ItemJoinedteamsItemPhotoRequestBuilder) {
    return NewItemJoinedteamsItemPhotoRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// PrimaryChannel provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemPrimarychannelPrimaryChannelRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) PrimaryChannel()(*ItemJoinedteamsItemPrimarychannelPrimaryChannelRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelPrimaryChannelRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Schedule provides operations to manage the schedule property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemScheduleRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) Schedule()(*ItemJoinedteamsItemScheduleRequestBuilder) {
    return NewItemJoinedteamsItemScheduleRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SendActivityNotification provides operations to call the sendActivityNotification method.
// returns a *ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) SendActivityNotification()(*ItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilder) {
    return NewItemJoinedteamsItemSendactivitynotificationSendActivityNotificationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tags provides operations to manage the tags property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemTagsRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) Tags()(*ItemJoinedteamsItemTagsRequestBuilder) {
    return NewItemJoinedteamsItemTagsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Template provides operations to manage the template property of the microsoft.graph.team entity.
// returns a *ItemJoinedteamsItemTemplateRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) Template()(*ItemJoinedteamsItemTemplateRequestBuilder) {
    return NewItemJoinedteamsItemTemplateRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property joinedTeams for users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsTeamItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get joinedTeams from users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsTeamItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property joinedTeams in users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, requestConfiguration *ItemJoinedteamsTeamItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// Unarchive provides operations to call the unarchive method.
// returns a *ItemJoinedteamsItemUnarchiveRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) Unarchive()(*ItemJoinedteamsItemUnarchiveRequestBuilder) {
    return NewItemJoinedteamsItemUnarchiveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemJoinedteamsTeamItemRequestBuilder when successful
func (m *ItemJoinedteamsTeamItemRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsTeamItemRequestBuilder) {
    return NewItemJoinedteamsTeamItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
