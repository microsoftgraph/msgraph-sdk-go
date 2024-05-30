package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPrimarychannelPrimaryChannelRequestBuilder provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
type ItemPrimarychannelPrimaryChannelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPrimarychannelPrimaryChannelRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimarychannelPrimaryChannelRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPrimarychannelPrimaryChannelRequestBuilderGetQueryParameters get the default channel, General, of a team.
type ItemPrimarychannelPrimaryChannelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPrimarychannelPrimaryChannelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimarychannelPrimaryChannelRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPrimarychannelPrimaryChannelRequestBuilderGetQueryParameters
}
// ItemPrimarychannelPrimaryChannelRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPrimarychannelPrimaryChannelRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration provides operations to call the completeMigration method.
// returns a *ItemPrimarychannelCompletemigrationCompleteMigrationRequestBuilder when successful
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) CompleteMigration()(*ItemPrimarychannelCompletemigrationCompleteMigrationRequestBuilder) {
    return NewItemPrimarychannelCompletemigrationCompleteMigrationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemPrimarychannelPrimaryChannelRequestBuilderInternal instantiates a new ItemPrimarychannelPrimaryChannelRequestBuilder and sets the default values.
func NewItemPrimarychannelPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimarychannelPrimaryChannelRequestBuilder) {
    m := &ItemPrimarychannelPrimaryChannelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/primaryChannel{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPrimarychannelPrimaryChannelRequestBuilder instantiates a new ItemPrimarychannelPrimaryChannelRequestBuilder and sets the default values.
func NewItemPrimarychannelPrimaryChannelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPrimarychannelPrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPrimarychannelPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property primaryChannel for teams
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPrimarychannelPrimaryChannelRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *ItemPrimarychannelDoesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalnameDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder when successful
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*ItemPrimarychannelDoesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalnameDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return NewItemPrimarychannelDoesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalnameDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilesFolder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
// returns a *ItemPrimarychannelFilesfolderFilesFolderRequestBuilder when successful
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) FilesFolder()(*ItemPrimarychannelFilesfolderFilesFolderRequestBuilder) {
    return NewItemPrimarychannelFilesfolderFilesFolderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the default channel, General, of a team.
// returns a Channelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-get-primarychannel?view=graph-rest-1.0
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPrimarychannelPrimaryChannelRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, error) {
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
// returns a *ItemPrimarychannelMembersRequestBuilder when successful
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) Members()(*ItemPrimarychannelMembersRequestBuilder) {
    return NewItemPrimarychannelMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.channel entity.
// returns a *ItemPrimarychannelMessagesRequestBuilder when successful
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) Messages()(*ItemPrimarychannelMessagesRequestBuilder) {
    return NewItemPrimarychannelMessagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property primaryChannel in teams
// returns a Channelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, requestConfiguration *ItemPrimarychannelPrimaryChannelRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, error) {
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
// returns a *ItemPrimarychannelProvisionemailProvisionEmailRequestBuilder when successful
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) ProvisionEmail()(*ItemPrimarychannelProvisionemailProvisionEmailRequestBuilder) {
    return NewItemPrimarychannelProvisionemailProvisionEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveEmail provides operations to call the removeEmail method.
// returns a *ItemPrimarychannelRemoveemailRemoveEmailRequestBuilder when successful
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) RemoveEmail()(*ItemPrimarychannelRemoveemailRemoveEmailRequestBuilder) {
    return NewItemPrimarychannelRemoveemailRemoveEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SharedWithTeams provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
// returns a *ItemPrimarychannelSharedwithteamsSharedWithTeamsRequestBuilder when successful
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) SharedWithTeams()(*ItemPrimarychannelSharedwithteamsSharedWithTeamsRequestBuilder) {
    return NewItemPrimarychannelSharedwithteamsSharedWithTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.channel entity.
// returns a *ItemPrimarychannelTabsRequestBuilder when successful
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) Tabs()(*ItemPrimarychannelTabsRequestBuilder) {
    return NewItemPrimarychannelTabsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property primaryChannel for teams
// returns a *RequestInformation when successful
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPrimarychannelPrimaryChannelRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the default channel, General, of a team.
// returns a *RequestInformation when successful
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPrimarychannelPrimaryChannelRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property primaryChannel in teams
// returns a *RequestInformation when successful
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, requestConfiguration *ItemPrimarychannelPrimaryChannelRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPrimarychannelPrimaryChannelRequestBuilder when successful
func (m *ItemPrimarychannelPrimaryChannelRequestBuilder) WithUrl(rawUrl string)(*ItemPrimarychannelPrimaryChannelRequestBuilder) {
    return NewItemPrimarychannelPrimaryChannelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
