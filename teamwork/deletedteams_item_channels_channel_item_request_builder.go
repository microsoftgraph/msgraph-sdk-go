package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DeletedteamsItemChannelsChannelItemRequestBuilder provides operations to manage the channels property of the microsoft.graph.deletedTeam entity.
type DeletedteamsItemChannelsChannelItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DeletedteamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeletedteamsItemChannelsChannelItemRequestBuilderGetQueryParameters the channels that are either shared with this deleted team or created in this deleted team.
type DeletedteamsItemChannelsChannelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeletedteamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeletedteamsItemChannelsChannelItemRequestBuilderGetQueryParameters
}
// DeletedteamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeletedteamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration provides operations to call the completeMigration method.
// returns a *DeletedteamsItemChannelsItemCompletemigrationCompleteMigrationRequestBuilder when successful
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) CompleteMigration()(*DeletedteamsItemChannelsItemCompletemigrationCompleteMigrationRequestBuilder) {
    return NewDeletedteamsItemChannelsItemCompletemigrationCompleteMigrationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewDeletedteamsItemChannelsChannelItemRequestBuilderInternal instantiates a new DeletedteamsItemChannelsChannelItemRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsChannelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsChannelItemRequestBuilder) {
    m := &DeletedteamsItemChannelsChannelItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/deletedTeams/{deletedTeam%2Did}/channels/{channel%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDeletedteamsItemChannelsChannelItemRequestBuilder instantiates a new DeletedteamsItemChannelsChannelItemRequestBuilder and sets the default values.
func NewDeletedteamsItemChannelsChannelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeletedteamsItemChannelsChannelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeletedteamsItemChannelsChannelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property channels for teamwork
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// returns a *DeletedteamsItemChannelsItemDoesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalnameDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder when successful
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*DeletedteamsItemChannelsItemDoesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalnameDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return NewDeletedteamsItemChannelsItemDoesuserhaveaccessuseriduseridtenantidtenantiduserprincipalnameuserprincipalnameDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilesFolder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
// returns a *DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder when successful
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) FilesFolder()(*DeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilder) {
    return NewDeletedteamsItemChannelsItemFilesfolderFilesFolderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the channels that are either shared with this deleted team or created in this deleted team.
// returns a Channelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, error) {
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
// returns a *DeletedteamsItemChannelsItemMembersRequestBuilder when successful
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) Members()(*DeletedteamsItemChannelsItemMembersRequestBuilder) {
    return NewDeletedteamsItemChannelsItemMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.channel entity.
// returns a *DeletedteamsItemChannelsItemMessagesRequestBuilder when successful
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) Messages()(*DeletedteamsItemChannelsItemMessagesRequestBuilder) {
    return NewDeletedteamsItemChannelsItemMessagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property channels in teamwork
// returns a Channelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, requestConfiguration *DeletedteamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, error) {
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
// returns a *DeletedteamsItemChannelsItemProvisionemailProvisionEmailRequestBuilder when successful
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) ProvisionEmail()(*DeletedteamsItemChannelsItemProvisionemailProvisionEmailRequestBuilder) {
    return NewDeletedteamsItemChannelsItemProvisionemailProvisionEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveEmail provides operations to call the removeEmail method.
// returns a *DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder when successful
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) RemoveEmail()(*DeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilder) {
    return NewDeletedteamsItemChannelsItemRemoveemailRemoveEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SharedWithTeams provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
// returns a *DeletedteamsItemChannelsItemSharedwithteamsSharedWithTeamsRequestBuilder when successful
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) SharedWithTeams()(*DeletedteamsItemChannelsItemSharedwithteamsSharedWithTeamsRequestBuilder) {
    return NewDeletedteamsItemChannelsItemSharedwithteamsSharedWithTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.channel entity.
// returns a *DeletedteamsItemChannelsItemTabsRequestBuilder when successful
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) Tabs()(*DeletedteamsItemChannelsItemTabsRequestBuilder) {
    return NewDeletedteamsItemChannelsItemTabsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property channels for teamwork
// returns a *RequestInformation when successful
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsChannelItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the channels that are either shared with this deleted team or created in this deleted team.
// returns a *RequestInformation when successful
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DeletedteamsItemChannelsChannelItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property channels in teamwork
// returns a *RequestInformation when successful
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, requestConfiguration *DeletedteamsItemChannelsChannelItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DeletedteamsItemChannelsChannelItemRequestBuilder when successful
func (m *DeletedteamsItemChannelsChannelItemRequestBuilder) WithUrl(rawUrl string)(*DeletedteamsItemChannelsChannelItemRequestBuilder) {
    return NewDeletedteamsItemChannelsChannelItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
