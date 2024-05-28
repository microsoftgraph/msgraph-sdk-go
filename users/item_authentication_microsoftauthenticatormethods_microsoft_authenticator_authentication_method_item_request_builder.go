package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder provides operations to manage the microsoftAuthenticatorMethods property of the microsoft.graph.authentication entity.
type ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetQueryParameters read the properties and relationships of a microsoftAuthenticatorAuthenticationMethod object.
type ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetQueryParameters
}
// NewItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal instantiates a new ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    m := &ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/microsoftAuthenticatorMethods/{microsoftAuthenticatorAuthenticationMethod%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder instantiates a new ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder and sets the default values.
func NewItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a microsoftAuthenticatorAuthenticationMethod object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/microsoftauthenticatorauthenticationmethod-delete?view=graph-rest-1.0
func (m *ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Device provides operations to manage the device property of the microsoft.graph.microsoftAuthenticatorAuthenticationMethod entity.
// returns a *ItemAuthenticationMicrosoftauthenticatormethodsItemDeviceRequestBuilder when successful
func (m *ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) Device()(*ItemAuthenticationMicrosoftauthenticatormethodsItemDeviceRequestBuilder) {
    return NewItemAuthenticationMicrosoftauthenticatormethodsItemDeviceRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a microsoftAuthenticatorAuthenticationMethod object.
// returns a MicrosoftAuthenticatorAuthenticationMethodable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/microsoftauthenticatorauthenticationmethod-get?view=graph-rest-1.0
func (m *ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MicrosoftAuthenticatorAuthenticationMethodable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateMicrosoftAuthenticatorAuthenticationMethodFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.MicrosoftAuthenticatorAuthenticationMethodable), nil
}
// ToDeleteRequestInformation deletes a microsoftAuthenticatorAuthenticationMethod object.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a microsoftAuthenticatorAuthenticationMethod object.
// returns a *RequestInformation when successful
func (m *ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder when successful
func (m *ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder) {
    return NewItemAuthenticationMicrosoftauthenticatormethodsMicrosoftAuthenticatorAuthenticationMethodItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
