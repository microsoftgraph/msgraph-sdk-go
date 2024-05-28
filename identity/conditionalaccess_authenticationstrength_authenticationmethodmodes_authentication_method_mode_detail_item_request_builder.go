package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder provides operations to manage the authenticationMethodModes property of the microsoft.graph.authenticationStrengthRoot entity.
type ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderGetQueryParameters names and descriptions of all valid authentication method modes in the system.
type ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderGetQueryParameters
}
// ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderInternal instantiates a new ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder) {
    m := &ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/authenticationStrength/authenticationMethodModes/{authenticationMethodModeDetail%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder instantiates a new ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder and sets the default values.
func NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property authenticationMethodModes for identity
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get names and descriptions of all valid authentication method modes in the system.
// returns a AuthenticationMethodModeDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodModeDetailable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property authenticationMethodModes in identity
// returns a AuthenticationMethodModeDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodModeDetailable, requestConfiguration *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodModeDetailable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property authenticationMethodModes for identity
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation names and descriptions of all valid authentication method modes in the system.
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property authenticationMethodModes in identity
// returns a *RequestInformation when successful
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationMethodModeDetailable, requestConfiguration *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder when successful
func (m *ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder) WithUrl(rawUrl string)(*ConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder) {
    return NewConditionalaccessAuthenticationstrengthAuthenticationmethodmodesAuthenticationMethodModeDetailItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
