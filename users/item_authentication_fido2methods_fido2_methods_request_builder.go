package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAuthenticationFido2methodsFido2MethodsRequestBuilder provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
type ItemAuthenticationFido2methodsFido2MethodsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationFido2methodsFido2MethodsRequestBuilderGetQueryParameters represents the FIDO2 security keys registered to a user for authentication.
type ItemAuthenticationFido2methodsFido2MethodsRequestBuilderGetQueryParameters struct {
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
// ItemAuthenticationFido2methodsFido2MethodsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationFido2methodsFido2MethodsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationFido2methodsFido2MethodsRequestBuilderGetQueryParameters
}
// ByFido2AuthenticationMethodId provides operations to manage the fido2Methods property of the microsoft.graph.authentication entity.
// returns a *ItemAuthenticationFido2methodsFido2AuthenticationMethodItemRequestBuilder when successful
func (m *ItemAuthenticationFido2methodsFido2MethodsRequestBuilder) ByFido2AuthenticationMethodId(fido2AuthenticationMethodId string)(*ItemAuthenticationFido2methodsFido2AuthenticationMethodItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if fido2AuthenticationMethodId != "" {
        urlTplParams["fido2AuthenticationMethod%2Did"] = fido2AuthenticationMethodId
    }
    return NewItemAuthenticationFido2methodsFido2AuthenticationMethodItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemAuthenticationFido2methodsFido2MethodsRequestBuilderInternal instantiates a new ItemAuthenticationFido2methodsFido2MethodsRequestBuilder and sets the default values.
func NewItemAuthenticationFido2methodsFido2MethodsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationFido2methodsFido2MethodsRequestBuilder) {
    m := &ItemAuthenticationFido2methodsFido2MethodsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/fido2Methods{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemAuthenticationFido2methodsFido2MethodsRequestBuilder instantiates a new ItemAuthenticationFido2methodsFido2MethodsRequestBuilder and sets the default values.
func NewItemAuthenticationFido2methodsFido2MethodsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationFido2methodsFido2MethodsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationFido2methodsFido2MethodsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemAuthenticationFido2methodsCountRequestBuilder when successful
func (m *ItemAuthenticationFido2methodsFido2MethodsRequestBuilder) Count()(*ItemAuthenticationFido2methodsCountRequestBuilder) {
    return NewItemAuthenticationFido2methodsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents the FIDO2 security keys registered to a user for authentication.
// returns a Fido2AuthenticationMethodCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemAuthenticationFido2methodsFido2MethodsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationFido2methodsFido2MethodsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Fido2AuthenticationMethodCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateFido2AuthenticationMethodCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Fido2AuthenticationMethodCollectionResponseable), nil
}
// ToGetRequestInformation represents the FIDO2 security keys registered to a user for authentication.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationFido2methodsFido2MethodsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationFido2methodsFido2MethodsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAuthenticationFido2methodsFido2MethodsRequestBuilder when successful
func (m *ItemAuthenticationFido2methodsFido2MethodsRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationFido2methodsFido2MethodsRequestBuilder) {
    return NewItemAuthenticationFido2methodsFido2MethodsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
