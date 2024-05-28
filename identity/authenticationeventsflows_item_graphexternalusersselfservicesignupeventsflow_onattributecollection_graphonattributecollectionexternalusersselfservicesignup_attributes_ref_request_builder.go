package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder provides operations to manage the collection of identityContainer entities.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderDeleteQueryParameters delete ref of navigation property attributes for identity
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderDeleteQueryParameters struct {
    // The delete Uri
    Id *string `uriparametername:"%40id"`
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderDeleteQueryParameters
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderGetQueryParameters get an identityUserFlowAttribute collection associated with an external identities self-service user flow represented by an externalUsersSelfServiceSignupEventsFlow object. These attributes are collected from the user during the authentication experience defined by the user flow.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderGetQueryParameters
}
// AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderInternal instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder) {
    m := &AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/graph.onAttributeCollectionExternalUsersSelfServiceSignUp/attributes/$ref?@id={%40id}{&%24count,%24filter,%24orderby,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder instantiates a new AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder and sets the default values.
func NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete ref of navigation property attributes for identity
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder) Delete(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get an identityUserFlowAttribute collection associated with an external identities self-service user flow represented by an externalUsersSelfServiceSignupEventsFlow object. These attributes are collected from the user during the authentication experience defined by the user flow.
// returns a StringCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onattributecollectionexternalusersselfservicesignup-list-attributes?view=graph-rest-1.0
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder) Get(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StringCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateStringCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.StringCollectionResponseable), nil
}
// Post add an attribute to an external identities self-service user flow represented by an externalUsersSelfServiceSignupEventsFlow object. You can add both custom and built-in attributes to a user flow. Prerequisite: PATCH the user flow with the updated onAttributeCollection event that includes the new attribute.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/onattributecollectionexternalusersselfservicesignup-post-attributes?view=graph-rest-1.0
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceCreateable, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete ref of navigation property attributes for identity
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/graph.onAttributeCollectionExternalUsersSelfServiceSignUp/attributes/$ref?@id={%40id}", m.BaseRequestBuilder.PathParameters)
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
// ToGetRequestInformation get an identityUserFlowAttribute collection associated with an external identities self-service user flow represented by an externalUsersSelfServiceSignupEventsFlow object. These attributes are collected from the user during the authentication experience defined by the user flow.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/graph.onAttributeCollectionExternalUsersSelfServiceSignUp/attributes/$ref{?%24count,%24filter,%24orderby,%24search,%24skip,%24top}", m.BaseRequestBuilder.PathParameters)
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
// ToPostRequestInformation add an attribute to an external identities self-service user flow represented by an externalUsersSelfServiceSignupEventsFlow object. You can add both custom and built-in attributes to a user flow. Prerequisite: PATCH the user flow with the updated onAttributeCollection event that includes the new attribute.
// returns a *RequestInformation when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ReferenceCreateable, requestConfiguration *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, "{+baseurl}/identity/authenticationEventsFlows/{authenticationEventsFlow%2Did}/graph.externalUsersSelfServiceSignUpEventsFlow/onAttributeCollection/graph.onAttributeCollectionExternalUsersSelfServiceSignUp/attributes/$ref", m.BaseRequestBuilder.PathParameters)
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
// returns a *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder when successful
func (m *AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder) WithUrl(rawUrl string)(*AuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder) {
    return NewAuthenticationeventsflowsItemGraphexternalusersselfservicesignupeventsflowOnattributecollectionGraphonattributecollectionexternalusersselfservicesignupAttributesRefRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
