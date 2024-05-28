package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder provides operations to manage the combinationConfigurations property of the microsoft.graph.authenticationStrengthPolicy entity.
type ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetQueryParameters read the properties and relationships of an authenticationCombinationConfiguration object.
type ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetQueryParameters
}
// ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderInternal instantiates a new ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) {
    m := &ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationStrength/policies/{authenticationStrengthPolicy%2Did}/combinationConfigurations/{authenticationCombinationConfiguration%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder instantiates a new ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete an authenticationCombinationConfiguration  for a custom authenticationStrengthPolicy object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthpolicy-delete-combinationconfigurations?view=graph-rest-1.0
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an authenticationCombinationConfiguration object.
// returns a AuthenticationCombinationConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationcombinationconfiguration-get?view=graph-rest-1.0
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the properties of an authenticationCombinationConfiguration object. In use, only fido2combinationConfigurations may be updated as they're the only type of authenticationCombinationConfiguration that may be created.
// returns a AuthenticationCombinationConfigurationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationcombinationconfiguration-update?view=graph-rest-1.0
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete an authenticationCombinationConfiguration  for a custom authenticationStrengthPolicy object.
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an authenticationCombinationConfiguration object.
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an authenticationCombinationConfiguration object. In use, only fido2combinationConfigurations may be updated as they're the only type of authenticationCombinationConfiguration that may be created.
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationCombinationConfigurationable, requestConfiguration *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) WithUrl(rawUrl string)(*ConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder) {
    return NewConditionalaccessAuthenticationstrengthPoliciesItemCombinationconfigurationsAuthenticationCombinationConfigurationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
