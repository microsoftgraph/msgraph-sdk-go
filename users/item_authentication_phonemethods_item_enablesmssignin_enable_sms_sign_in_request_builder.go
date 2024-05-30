package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder provides operations to call the enableSmsSignIn method.
type ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilderInternal instantiates a new ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder and sets the default values.
func NewItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder) {
    m := &ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/authentication/phoneMethods/{phoneAuthenticationMethod%2Did}/enableSmsSignIn", pathParameters),
    }
    return m
}
// NewItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder instantiates a new ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder and sets the default values.
func NewItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilderInternal(urlParams, requestAdapter)
}
// Post enable SMS sign-in for an existing mobile phone number registered to a user. To be successfully enabled:
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/phoneauthenticationmethod-enablesmssignin?view=graph-rest-1.0
func (m *ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation enable SMS sign-in for an existing mobile phone number registered to a user. To be successfully enabled:
// returns a *RequestInformation when successful
func (m *ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder when successful
func (m *ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder) WithUrl(rawUrl string)(*ItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder) {
    return NewItemAuthenticationPhonemethodsItemEnablesmssigninEnableSmsSignInRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
