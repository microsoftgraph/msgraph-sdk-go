package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder provides operations to call the revokeSignInSessions method.
type ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemRevokesigninsessionsRevokeSignInSessionsRequestBuilderInternal instantiates a new ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder and sets the default values.
func NewItemRevokesigninsessionsRevokeSignInSessionsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder) {
    m := &ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/revokeSignInSessions", pathParameters),
    }
    return m
}
// NewItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder instantiates a new ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder and sets the default values.
func NewItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRevokesigninsessionsRevokeSignInSessionsRequestBuilderInternal(urlParams, requestAdapter)
}
// Post invalidates all the refresh tokens issued to applications for a user (as well as session cookies in a user's browser), by resetting the signInSessionsValidFromDateTime user property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the user has a lost or stolen device. This operation prevents access to the organization's data through applications on the device by requiring the user to sign in again to all applications that they have previously consented to, independent of device.
// Deprecated: This method is obsolete. Use PostAsRevokeSignInSessionsPostResponse instead.
// returns a ItemRevokesigninsessionsRevokeSignInSessionsResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-revokesigninsessions?view=graph-rest-1.0
func (m *ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilderPostRequestConfiguration)(ItemRevokesigninsessionsRevokeSignInSessionsResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemRevokesigninsessionsRevokeSignInSessionsResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemRevokesigninsessionsRevokeSignInSessionsResponseable), nil
}
// PostAsRevokeSignInSessionsPostResponse invalidates all the refresh tokens issued to applications for a user (as well as session cookies in a user's browser), by resetting the signInSessionsValidFromDateTime user property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the user has a lost or stolen device. This operation prevents access to the organization's data through applications on the device by requiring the user to sign in again to all applications that they have previously consented to, independent of device.
// returns a ItemRevokesigninsessionsRevokeSignInSessionsPostResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/user-revokesigninsessions?view=graph-rest-1.0
func (m *ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder) PostAsRevokeSignInSessionsPostResponse(ctx context.Context, requestConfiguration *ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilderPostRequestConfiguration)(ItemRevokesigninsessionsRevokeSignInSessionsPostResponseable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemRevokesigninsessionsRevokeSignInSessionsPostResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemRevokesigninsessionsRevokeSignInSessionsPostResponseable), nil
}
// ToPostRequestInformation invalidates all the refresh tokens issued to applications for a user (as well as session cookies in a user's browser), by resetting the signInSessionsValidFromDateTime user property to the current date-time. Typically, this operation is performed (by the user or an administrator) if the user has a lost or stolen device. This operation prevents access to the organization's data through applications on the device by requiring the user to sign in again to all applications that they have previously consented to, independent of device.
// returns a *RequestInformation when successful
func (m *ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder when successful
func (m *ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder) WithUrl(rawUrl string)(*ItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder) {
    return NewItemRevokesigninsessionsRevokeSignInSessionsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
