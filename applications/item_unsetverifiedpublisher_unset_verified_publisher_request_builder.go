package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder provides operations to call the unsetVerifiedPublisher method.
type ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilderInternal instantiates a new ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder and sets the default values.
func NewItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder) {
    m := &ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/unsetVerifiedPublisher", pathParameters),
    }
    return m
}
// NewItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder instantiates a new ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder and sets the default values.
func NewItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilderInternal(urlParams, requestAdapter)
}
// Post unset the verifiedPublisher previously set on an application, removing all verified publisher properties. For more information, see Publisher verification.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/application-unsetverifiedpublisher?view=graph-rest-1.0
func (m *ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation unset the verifiedPublisher previously set on an application, removing all verified publisher properties. For more information, see Publisher verification.
// returns a *RequestInformation when successful
func (m *ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder when successful
func (m *ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder) WithUrl(rawUrl string)(*ItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder) {
    return NewItemUnsetverifiedpublisherUnsetVerifiedPublisherRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
