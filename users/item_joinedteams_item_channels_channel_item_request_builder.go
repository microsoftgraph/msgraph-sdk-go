package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemChannelsChannelItemRequestBuilder provides operations to manage the channels property of the microsoft.graph.team entity.
type ItemJoinedteamsItemChannelsChannelItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemJoinedteamsItemChannelsChannelItemRequestBuilderGetQueryParameters the collection of channels and messages associated with the team.
type ItemJoinedteamsItemChannelsChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemJoinedteamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedteamsItemChannelsChannelItemRequestBuilderGetQueryParameters
}
// ItemJoinedteamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration provides operations to call the completeMigration method.
// returns a *ItemJoinedteamsItemChannelsItemCompletemigrationCompleteMigrationRequestBuilder when successful
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) CompleteMigration()(*ItemJoinedteamsItemChannelsItemCompletemigrationCompleteMigrationRequestBuilder) {
    return NewItemJoinedteamsItemChannelsItemCompletemigrationCompleteMigrationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemJoinedteamsItemChannelsChannelItemRequestBuilderInternal instantiates a new ItemJoinedteamsItemChannelsChannelItemRequestBuilder and sets the default values.
func NewItemJoinedteamsItemChannelsChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemChannelsChannelItemRequestBuilder) {
    m := &ItemJoinedteamsItemChannelsChannelItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/channels/{channel%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemChannelsChannelItemRequestBuilder instantiates a new ItemJoinedteamsItemChannelsChannelItemRequestBuilder and sets the default values.
func NewItemJoinedteamsItemChannelsChannelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemChannelsChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemChannelsChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property channels for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemJoinedteamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName provides operations to call the doesUserHaveAccess method.
// returns a *ItemJoinedteamsItemChannelsItemDoesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalnameDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder when successful
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*ItemJoinedteamsItemChannelsItemDoesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalnameDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return NewItemJoinedteamsItemChannelsItemDoesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalnameDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilesFolder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
// returns a *ItemJoinedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder when successful
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) FilesFolder()(*ItemJoinedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder) {
    return NewItemJoinedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the collection of channels and messages associated with the team.
// returns a Channelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedteamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChannelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable), nil
}
// Members provides operations to manage the members property of the microsoft.graph.channel entity.
// returns a *ItemJoinedteamsItemChannelsItemMembersRequestBuilder when successful
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) Members()(*ItemJoinedteamsItemChannelsItemMembersRequestBuilder) {
    return NewItemJoinedteamsItemChannelsItemMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.channel entity.
// returns a *ItemJoinedteamsItemChannelsItemMessagesRequestBuilder when successful
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) Messages()(*ItemJoinedteamsItemChannelsItemMessagesRequestBuilder) {
    return NewItemJoinedteamsItemChannelsItemMessagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property channels in users
// returns a Channelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, requestConfiguration *ItemJoinedteamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChannelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable), nil
}
// ProvisionEmail provides operations to call the provisionEmail method.
// returns a *ItemJoinedteamsItemChannelsItemProvisionemailProvisionEmailRequestBuilder when successful
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) ProvisionEmail()(*ItemJoinedteamsItemChannelsItemProvisionemailProvisionEmailRequestBuilder) {
    return NewItemJoinedteamsItemChannelsItemProvisionemailProvisionEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveEmail provides operations to call the removeEmail method.
// returns a *ItemJoinedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder when successful
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) RemoveEmail()(*ItemJoinedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder) {
    return NewItemJoinedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SharedWithTeams provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
// returns a *ItemJoinedteamsItemChannelsItemSharedwithteamsSharedWithTeamsRequestBuilder when successful
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) SharedWithTeams()(*ItemJoinedteamsItemChannelsItemSharedwithteamsSharedWithTeamsRequestBuilder) {
    return NewItemJoinedteamsItemChannelsItemSharedwithteamsSharedWithTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.channel entity.
// returns a *ItemJoinedteamsItemChannelsItemTabsRequestBuilder when successful
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) Tabs()(*ItemJoinedteamsItemChannelsItemTabsRequestBuilder) {
    return NewItemJoinedteamsItemChannelsItemTabsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property channels for users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of channels and messages associated with the team.
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property channels in users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, requestConfiguration *ItemJoinedteamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemJoinedteamsItemChannelsChannelItemRequestBuilder when successful
func (m *ItemJoinedteamsItemChannelsChannelItemRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemChannelsChannelItemRequestBuilder) {
    return NewItemJoinedteamsItemChannelsChannelItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
