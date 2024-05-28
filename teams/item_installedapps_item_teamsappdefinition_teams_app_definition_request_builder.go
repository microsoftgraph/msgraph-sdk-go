package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder provides operations to manage the teamsAppDefinition property of the microsoft.graph.teamsAppInstallation entity.
type ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetQueryParameters the details of this version of the app.
type ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetQueryParameters
}
// NewItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderInternal instantiates a new ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder and sets the default values.
func NewItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder) {
    m := &ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/installedApps/{teamsAppInstallation%2Did}/teamsAppDefinition{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder instantiates a new ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder and sets the default values.
func NewItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the details of this version of the app.
// returns a TeamsAppDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppDefinitionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamsAppDefinitionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppDefinitionable), nil
}
// ToGetRequestInformation the details of this version of the app.
// returns a *RequestInformation when successful
func (m *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder when successful
func (m *ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder) WithUrl(rawUrl string)(*ItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder) {
    return NewItemInstalledappsItemTeamsappdefinitionTeamsAppDefinitionRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
