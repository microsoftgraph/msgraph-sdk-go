package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamPrimaryChannelRequestBuilder provides operations to manage the primaryChannel property of the microsoft.graph.team entity.
type ItemTeamPrimaryChannelRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamPrimaryChannelRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimaryChannelRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamPrimaryChannelRequestBuilderGetQueryParameters get the default channel, General, of a team.
type ItemTeamPrimaryChannelRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamPrimaryChannelRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimaryChannelRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamPrimaryChannelRequestBuilderGetQueryParameters
}
// ItemTeamPrimaryChannelRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimaryChannelRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CompleteMigration provides operations to call the completeMigration method.
func (m *ItemTeamPrimaryChannelRequestBuilder) CompleteMigration()(*ItemTeamPrimaryChannelCompleteMigrationRequestBuilder) {
    return NewItemTeamPrimaryChannelCompleteMigrationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamPrimaryChannelRequestBuilderInternal instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimaryChannelRequestBuilder) {
    m := &ItemTeamPrimaryChannelRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/primaryChannel{?%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemTeamPrimaryChannelRequestBuilder instantiates a new PrimaryChannelRequestBuilder and sets the default values.
func NewItemTeamPrimaryChannelRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimaryChannelRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamPrimaryChannelRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property primaryChannel for groups
func (m *ItemTeamPrimaryChannelRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamPrimaryChannelRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName provides operations to call the doesUserHaveAccess method.
func (m *ItemTeamPrimaryChannelRequestBuilder) DoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalName()(*ItemTeamPrimaryChannelDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilder) {
    return NewItemTeamPrimaryChannelDoesUserHaveAccessuserIdUserIdTenantIdTenantIdUserPrincipalNameUserPrincipalNameRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilesFolder provides operations to manage the filesFolder property of the microsoft.graph.channel entity.
func (m *ItemTeamPrimaryChannelRequestBuilder) FilesFolder()(*ItemTeamPrimaryChannelFilesFolderRequestBuilder) {
    return NewItemTeamPrimaryChannelFilesFolderRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the default channel, General, of a team.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-get-primarychannel?view=graph-rest-1.0
func (m *ItemTeamPrimaryChannelRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamPrimaryChannelRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *ItemTeamPrimaryChannelRequestBuilder) Members()(*ItemTeamPrimaryChannelMembersRequestBuilder) {
    return NewItemTeamPrimaryChannelMembersRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Messages provides operations to manage the messages property of the microsoft.graph.channel entity.
func (m *ItemTeamPrimaryChannelRequestBuilder) Messages()(*ItemTeamPrimaryChannelMessagesRequestBuilder) {
    return NewItemTeamPrimaryChannelMessagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property primaryChannel in groups
func (m *ItemTeamPrimaryChannelRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, requestConfiguration *ItemTeamPrimaryChannelRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *ItemTeamPrimaryChannelRequestBuilder) ProvisionEmail()(*ItemTeamPrimaryChannelProvisionEmailRequestBuilder) {
    return NewItemTeamPrimaryChannelProvisionEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// RemoveEmail provides operations to call the removeEmail method.
func (m *ItemTeamPrimaryChannelRequestBuilder) RemoveEmail()(*ItemTeamPrimaryChannelRemoveEmailRequestBuilder) {
    return NewItemTeamPrimaryChannelRemoveEmailRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SharedWithTeams provides operations to manage the sharedWithTeams property of the microsoft.graph.channel entity.
func (m *ItemTeamPrimaryChannelRequestBuilder) SharedWithTeams()(*ItemTeamPrimaryChannelSharedWithTeamsRequestBuilder) {
    return NewItemTeamPrimaryChannelSharedWithTeamsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Tabs provides operations to manage the tabs property of the microsoft.graph.channel entity.
func (m *ItemTeamPrimaryChannelRequestBuilder) Tabs()(*ItemTeamPrimaryChannelTabsRequestBuilder) {
    return NewItemTeamPrimaryChannelTabsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property primaryChannel for groups
func (m *ItemTeamPrimaryChannelRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamPrimaryChannelRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// ToGetRequestInformation get the default channel, General, of a team.
func (m *ItemTeamPrimaryChannelRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamPrimaryChannelRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
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
// ToPatchRequestInformation update the navigation property primaryChannel in groups
func (m *ItemTeamPrimaryChannelRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Channelable, requestConfiguration *ItemTeamPrimaryChannelRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.Headers.Add("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemTeamPrimaryChannelRequestBuilder) WithUrl(rawUrl string)(*ItemTeamPrimaryChannelRequestBuilder) {
    return NewItemTeamPrimaryChannelRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
