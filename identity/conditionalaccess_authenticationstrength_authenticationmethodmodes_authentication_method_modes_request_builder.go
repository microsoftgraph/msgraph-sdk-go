package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder provides operations to manage the authenticationMethodModes property of the microsoft.graph.authenticationStrengthRoot entity.
type ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderGetQueryParameters get a list of all supported authentication methods, or all supported authentication method combinations as a list of authenticationMethodModes objects and their properties.
type ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderGetQueryParameters struct {
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
// ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderGetQueryParameters
}
// ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAuthenticationMethodModeDetailId provides operations to manage the authenticationMethodModes property of the microsoft.graph.authenticationStrengthRoot entity.
// returns a *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder) ByAuthenticationMethodModeDetailId(authenticationMethodModeDetailId string)(*ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if authenticationMethodModeDetailId != "" {
        urlTplParams["authenticationMethodModeDetail%2Did"] = authenticationMethodModeDetailId
    }
    return NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderInternal instantiates a new ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder) {
    m := &ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationStrength/authenticationMethodModes{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder instantiates a new ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesCountRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder) Count()(*ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesCountRequestBuilder) {
    return NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of all supported authentication methods, or all supported authentication method combinations as a list of authenticationMethodModes objects and their properties.
// returns a AuthenticationMethodModeDetailCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/authenticationstrengthroot-list-authenticationmethodmodes?view=graph-rest-1.0
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodModeDetailCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationMethodModeDetailCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodModeDetailCollectionResponseable), nil
}
// Post create new navigation property to authenticationMethodModes for identity
// returns a AuthenticationMethodModeDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodModeDetailable, requestConfiguration *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodModeDetailable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationMethodModeDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodModeDetailable), nil
}
// ToGetRequestInformation get a list of all supported authentication methods, or all supported authentication method combinations as a list of authenticationMethodModes objects and their properties.
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to authenticationMethodModes for identity
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodModeDetailable, requestConfiguration *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder) WithUrl(rawUrl string)(*ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder) {
    return NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
