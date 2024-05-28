package reports

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder provides operations to manage the userRegistrationDetails property of the microsoft.graph.authenticationMethodsRoot entity.
type AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderGetQueryParameters read the properties and relationships of a userRegistrationDetails object.
type AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderGetQueryParameters
}
// AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderInternal instantiates a new AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder and sets the default values.
func NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder) {
    m := &AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/reports/authenticationMethods/userRegistrationDetails/{userRegistrationDetails%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder instantiates a new AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder and sets the default values.
func NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userRegistrationDetails for reports
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a userRegistrationDetails object.
// returns a UserRegistrationDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/userregistrationdetails-get?view=graph-rest-1.0
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserRegistrationDetailsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserRegistrationDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserRegistrationDetailsable), nil
}
// Patch update the navigation property userRegistrationDetails in reports
// returns a UserRegistrationDetailsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserRegistrationDetailsable, requestConfiguration *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserRegistrationDetailsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserRegistrationDetailsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserRegistrationDetailsable), nil
}
// ToDeleteRequestInformation delete navigation property userRegistrationDetails for reports
// returns a *RequestInformation when successful
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a userRegistrationDetails object.
// returns a *RequestInformation when successful
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userRegistrationDetails in reports
// returns a *RequestInformation when successful
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserRegistrationDetailsable, requestConfiguration *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder when successful
func (m *AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder) WithUrl(rawUrl string)(*AuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder) {
    return NewAuthenticationmethodsUserregistrationdetailsUserRegistrationDetailsItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
