package applications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder provides operations to call the setVerifiedPublisher method.
type ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemSetverifiedpublisherSetVerifiedPublisherRequestBuilderInternal instantiates a new ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder and sets the default values.
func NewItemSetverifiedpublisherSetVerifiedPublisherRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder) {
    m := &ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/applications/{application%2Did}/setVerifiedPublisher", pathParameters),
    }
    return m
}
// NewItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder instantiates a new ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder and sets the default values.
func NewItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSetverifiedpublisherSetVerifiedPublisherRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set the verifiedPublisher on an application. For more information, including prerequisites to setting a verified publisher, see Publisher verification.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/application-setverifiedpublisher?view=graph-rest-1.0
func (m *ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder) Post(ctx context.Context, body ItemSetverifiedpublisherSetVerifiedPublisherPostRequestBodyable, requestConfiguration *ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation set the verifiedPublisher on an application. For more information, including prerequisites to setting a verified publisher, see Publisher verification.
// returns a *RequestInformation when successful
func (m *ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemSetverifiedpublisherSetVerifiedPublisherPostRequestBodyable, requestConfiguration *ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder when successful
func (m *ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder) WithUrl(rawUrl string)(*ItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder) {
    return NewItemSetverifiedpublisherSetVerifiedPublisherRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
