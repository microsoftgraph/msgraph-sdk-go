package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder provides operations to manage the combinationConfigurations property of the microsoft.graph.authenticationStrengthPolicy entity.
type ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderGetQueryParameters get the authenticationCombinationConfiguration objects for an authentication strength policy. authenticationCombinationConfiguration represents requirements placed on specific authentication method combinations that require specified variants of those authentication methods to be used when authenticating. Currently, only fido2combinationConfigurations objects are supported. authenticationCombinationConfiguration objects are supported only for custom authentication strengths.
type ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderGetQueryParameters struct {
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
// ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderGetQueryParameters
}
// ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAuthenticationCombinationConfigurationId provides operations to manage the combinationConfigurations property of the microsoft.graph.authenticationStrengthPolicy entity.
// returns a *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder) ByAuthenticationCombinationConfigurationId(authenticationCombinationConfigurationId string)(*ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if authenticationCombinationConfigurationId != "" {
        urlTplParams["authenticationCombinationConfiguration%2Did"] = authenticationCombinationConfigurationId
    }
    return NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderInternal instantiates a new ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder) {
    m := &ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationStrength/policies/{authenticationStrengthPolicy%2Did}/combinationConfigurations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder instantiates a new ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCountRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder) Count()(*ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCountRequestBuilder) {
    return NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the authenticationCombinationConfiguration objects for an authentication strength policy. authenticationCombinationConfiguration represents requirements placed on specific authentication method combinations that require specified variants of those authentication methods to be used when authenticating. Currently, only fido2combinationConfigurations objects are supported. authenticationCombinationConfiguration objects are supported only for custom authentication strengths.
// returns a AuthenticationCombinationConfigurationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-list-combinationconfigurations?view=graph-rest-1.0
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationCombinationConfigurationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationCollectionResponseable), nil
}
// Post create a new authenticationCombinationConfiguration object. In use, only fido2combinationConfigurations may be created, and these may only be created for custom authentication strength policies.
// returns a AuthenticationCombinationConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-post-combinationconfigurations?view=graph-rest-1.0
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationCombinationConfigurationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable), nil
}
// ToGetRequestInformation get the authenticationCombinationConfiguration objects for an authentication strength policy. authenticationCombinationConfiguration represents requirements placed on specific authentication method combinations that require specified variants of those authentication methods to be used when authenticating. Currently, only fido2combinationConfigurations objects are supported. authenticationCombinationConfiguration objects are supported only for custom authentication strengths.
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new authenticationCombinationConfiguration object. In use, only fido2combinationConfigurations may be created, and these may only be created for custom authentication strength policies.
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder) WithUrl(rawUrl string)(*ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder) {
    return NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsCombinationConfigurationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
