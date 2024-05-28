package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder provides operations to manage the includeApplications property of the microsoft.graph.authenticationConditionsApplications entity.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetQueryParameters get includeApplications from identity
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetQueryParameters
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderInternal instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) {
    m := &AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/conditions/applications/includeApplications/{authenticationConditionApplication%2DappId}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property includeApplications for identity
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get includeApplications from identity
// returns a AuthenticationConditionApplicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property includeApplications in identity
// returns a AuthenticationConditionApplicationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationable, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property includeApplications for identity
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get includeApplications from identity
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property includeApplications in identity
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AuthenticationConditionApplicationable, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) WithUrl(rawUrl string)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowConditionsApplicationsIncludeapplicationsAuthenticationConditionApplicationAppItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
