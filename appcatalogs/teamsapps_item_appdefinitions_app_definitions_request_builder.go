package appcatalogs

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder provides operations to manage the appDefinitions property of the microsoft.graph.teamsApp entity.
type TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderGetQueryParameters the details for each version of the app.
type TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderGetQueryParameters
}
// TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTeamsAppDefinitionId provides operations to manage the appDefinitions property of the microsoft.graph.teamsApp entity.
// returns a *TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder when successful
func (m *TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder) ByTeamsAppDefinitionId(teamsAppDefinitionId string)(*TeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if teamsAppDefinitionId != "" {
        urlTplParams["teamsAppDefinition%2Did"] = teamsAppDefinitionId
    }
    return NewTeamsappsItemAppdefinitionsTeamsAppDefinitionItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderInternal instantiates a new TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder and sets the default values.
func NewTeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder) {
    m := &TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/appCatalogs/teamsApps/{teamsApp%2Did}/appDefinitions{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder instantiates a new TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder and sets the default values.
func NewTeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TeamsappsItemAppdefinitionsCountRequestBuilder when successful
func (m *TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder) Count()(*TeamsappsItemAppdefinitionsCountRequestBuilder) {
    return NewTeamsappsItemAppdefinitionsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the details for each version of the app.
// returns a TeamsAppDefinitionCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder) Get(ctx context.Context, requestConfiguration *TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppDefinitionCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamsAppDefinitionCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppDefinitionCollectionResponseable), nil
}
// Post update an app previously published to the Microsoft Teams app catalog. To update an app, the distributionMethod property for the app must be set to organization. This API specifically updates an app published to your organization's app catalog (the tenant app catalog).
// returns a TeamsAppDefinitionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/teamsapp-update?view=graph-rest-1.0
func (m *TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppDefinitionable, requestConfiguration *TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppDefinitionable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation the details for each version of the app.
// returns a *RequestInformation when successful
func (m *TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation update an app previously published to the Microsoft Teams app catalog. To update an app, the distributionMethod property for the app must be set to organization. This API specifically updates an app published to your organization's app catalog (the tenant app catalog).
// returns a *RequestInformation when successful
func (m *TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TeamsAppDefinitionable, requestConfiguration *TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder when successful
func (m *TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder) WithUrl(rawUrl string)(*TeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder) {
    return NewTeamsappsItemAppdefinitionsAppDefinitionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
