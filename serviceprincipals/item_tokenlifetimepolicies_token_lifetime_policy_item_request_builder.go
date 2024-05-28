package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder provides operations to manage the tokenLifetimePolicies property of the microsoft.graph.servicePrincipal entity.
type ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetQueryParameters the tokenLifetimePolicies assigned to this service principal.
type ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetQueryParameters
}
// NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderInternal instantiates a new ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder and sets the default values.
func NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) {
    m := &ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/tokenLifetimePolicies/{tokenLifetimePolicy%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder instantiates a new ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder and sets the default values.
func NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get the tokenLifetimePolicies assigned to this service principal.
// returns a TokenLifetimePolicyable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TokenLifetimePolicyable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTokenLifetimePolicyFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TokenLifetimePolicyable), nil
}
// ToGetRequestInformation the tokenLifetimePolicies assigned to this service principal.
// returns a *RequestInformation when successful
func (m *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder when successful
func (m *ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) WithUrl(rawUrl string)(*ItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder) {
    return NewItemTokenlifetimepoliciesTokenLifetimePolicyItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
