package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder provides operations to manage the installedApps property of the microsoft.graph.chat entity.
type ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderGetQueryParameters a collection of all the apps in the chat. Nullable.
type ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderGetQueryParameters
}
// ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderInternal instantiates a new ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder and sets the default values.
func NewItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder) {
    m := &ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/chats/{chat%2Did}/installedApps/{teamsAppInstallation%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder instantiates a new ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder and sets the default values.
func NewItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property installedApps for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get a collection of all the apps in the chat. Nullable.
// returns a TeamsAppInstallationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppInstallationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamsAppInstallationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppInstallationable), nil
}
// Patch update the navigation property installedApps in users
// returns a TeamsAppInstallationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppInstallationable, requestConfiguration *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppInstallationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamsAppInstallationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppInstallationable), nil
}
// TeamsApp provides operations to manage the teamsApp property of the microsoft.graph.teamsAppInstallation entity.
// returns a *ItemChatsItemInstalledappsItemTeamsappTeamsAppRequestBuilder when successful
func (m *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder) TeamsApp()(*ItemChatsItemInstalledappsItemTeamsappTeamsAppRequestBuilder) {
    return NewItemChatsItemInstalledappsItemTeamsappTeamsAppRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TeamsAppDefinition provides operations to manage the teamsAppDefinition property of the microsoft.graph.teamsAppInstallation entity.
// returns a *ItemChatsItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder when successful
func (m *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder) TeamsAppDefinition()(*ItemChatsItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder) {
    return NewItemChatsItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property installedApps for users
// returns a *RequestInformation when successful
func (m *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation a collection of all the apps in the chat. Nullable.
// returns a *RequestInformation when successful
func (m *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property installedApps in users
// returns a *RequestInformation when successful
func (m *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppInstallationable, requestConfiguration *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Upgrade provides operations to call the upgrade method.
// returns a *ItemChatsItemInstalledappsItemUpgradeRequestBuilder when successful
func (m *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder) Upgrade()(*ItemChatsItemInstalledappsItemUpgradeRequestBuilder) {
    return NewItemChatsItemInstalledappsItemUpgradeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder when successful
func (m *ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder) WithUrl(rawUrl string)(*ItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder) {
    return NewItemChatsItemInstalledappsTeamsAppInstallationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
