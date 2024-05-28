package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder provides operations to manage the userConsentRequests property of the microsoft.graph.appConsentRequest entity.
type AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderGetQueryParameters retrieve a collection of userConsentRequest objects and their properties.
type AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderGetQueryParameters struct {
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
// AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderGetQueryParameters
}
// AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByUserConsentRequestId provides operations to manage the userConsentRequests property of the microsoft.graph.appConsentRequest entity.
// returns a *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder) ByUserConsentRequestId(userConsentRequestId string)(*AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if userConsentRequestId != "" {
        urlTplParams["userConsentRequest%2Did"] = userConsentRequestId
    }
    return NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderInternal instantiates a new AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder and sets the default values.
func NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder) {
    m := &AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/appConsent/appConsentRequests/{appConsentRequest%2Did}/userConsentRequests{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder instantiates a new AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder and sets the default values.
func NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AppconsentAppconsentrequestsItemUserconsentrequestsCountRequestBuilder when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder) Count()(*AppconsentAppconsentrequestsItemUserconsentrequestsCountRequestBuilder) {
    return NewAppconsentAppconsentrequestsItemUserconsentrequestsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// FilterByCurrentUserWithOn provides operations to call the filterByCurrentUser method.
// returns a *AppconsentAppconsentrequestsItemUserconsentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder) FilterByCurrentUserWithOn(on *string)(*AppconsentAppconsentrequestsItemUserconsentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilder) {
    return NewAppconsentAppconsentrequestsItemUserconsentrequestsFilterbycurrentuserwithonFilterByCurrentUserWithOnRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, on)
}
// Get retrieve a collection of userConsentRequest objects and their properties.
// returns a UserConsentRequestCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/appconsentrequest-list-userconsentrequests?view=graph-rest-1.0
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder) Get(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserConsentRequestCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestCollectionResponseable), nil
}
// Post create new navigation property to userConsentRequests for identityGovernance
// returns a UserConsentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserConsentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable), nil
}
// ToGetRequestInformation retrieve a collection of userConsentRequest objects and their properties.
// returns a *RequestInformation when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to userConsentRequests for identityGovernance
// returns a *RequestInformation when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder) WithUrl(rawUrl string)(*AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder) {
    return NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
