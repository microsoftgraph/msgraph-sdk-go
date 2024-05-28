package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder provides operations to manage the includeApplications property of the microsoft.graph.authenticationConditionsApplications entity.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetQueryParameters get includeApplications from identity
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetQueryParameters struct {
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
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetQueryParameters
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAuthenticationConditionApplicationAppId provides operations to manage the includeApplications property of the microsoft.graph.authenticationConditionsApplications entity.
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) ByAuthenticationConditionApplicationAppId(authenticationConditionApplicationAppId string)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if authenticationConditionApplicationAppId != "" {
        urlTplParams["authenticationConditionApplication%2DappId"] = authenticationConditionApplicationAppId
    }
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderInternal instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) {
    m := &AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/conditions/applications/includeApplications{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsCountRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) Count()(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsCountRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get includeApplications from identity
// returns a AuthenticationConditionApplicationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationConditionApplicationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationCollectionResponseable), nil
}
// Post create new navigation property to includeApplications for identity
// returns a AuthenticationConditionApplicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationable, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAuthenticationConditionApplicationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationable), nil
}
// ToGetRequestInformation get includeApplications from identity
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to includeApplications for identity
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationable, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) WithUrl(rawUrl string)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsIncludeApplicationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
