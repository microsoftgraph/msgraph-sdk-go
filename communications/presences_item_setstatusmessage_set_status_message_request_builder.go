package communications

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// PresencesItemSetstatusmessageSetStatusMessageRequestBuilder provides operations to call the setStatusMessage method.
type PresencesItemSetstatusmessageSetStatusMessageRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// PresencesItemSetstatusmessageSetStatusMessageRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type PresencesItemSetstatusmessageSetStatusMessageRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewPresencesItemSetstatusmessageSetStatusMessageRequestBuilderInternal instantiates a new PresencesItemSetstatusmessageSetStatusMessageRequestBuilder and sets the default values.
func NewPresencesItemSetstatusmessageSetStatusMessageRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemSetstatusmessageSetStatusMessageRequestBuilder) {
    m := &PresencesItemSetstatusmessageSetStatusMessageRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/communications/presences/{presence%2Did}/setStatusMessage", pathParameters),
    }
    return m
}
// NewPresencesItemSetstatusmessageSetStatusMessageRequestBuilder instantiates a new PresencesItemSetstatusmessageSetStatusMessageRequestBuilder and sets the default values.
func NewPresencesItemSetstatusmessageSetStatusMessageRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PresencesItemSetstatusmessageSetStatusMessageRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPresencesItemSetstatusmessageSetStatusMessageRequestBuilderInternal(urlParams, requestAdapter)
}
// Post set a presence status message for a user. An optional expiration date and time can be supplied.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/presence-setstatusmessage?view=graph-rest-1.0
func (m *PresencesItemSetstatusmessageSetStatusMessageRequestBuilder) Post(ctx context.Context, body PresencesItemSetstatusmessageSetStatusMessagePostRequestBodyable, requestConfiguration *PresencesItemSetstatusmessageSetStatusMessageRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation set a presence status message for a user. An optional expiration date and time can be supplied.
// returns a *RequestInformation when successful
func (m *PresencesItemSetstatusmessageSetStatusMessageRequestBuilder) ToPostRequestInformation(ctx context.Context, body PresencesItemSetstatusmessageSetStatusMessagePostRequestBodyable, requestConfiguration *PresencesItemSetstatusmessageSetStatusMessageRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *PresencesItemSetstatusmessageSetStatusMessageRequestBuilder when successful
func (m *PresencesItemSetstatusmessageSetStatusMessageRequestBuilder) WithUrl(rawUrl string)(*PresencesItemSetstatusmessageSetStatusMessageRequestBuilder) {
    return NewPresencesItemSetstatusmessageSetStatusMessageRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
