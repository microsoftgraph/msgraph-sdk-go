package serviceprincipals

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemRemovekeyRemoveKeyRequestBuilder provides operations to call the removeKey method.
type ItemRemovekeyRemoveKeyRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemRemovekeyRemoveKeyRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemRemovekeyRemoveKeyRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemRemovekeyRemoveKeyRequestBuilderInternal instantiates a new ItemRemovekeyRemoveKeyRequestBuilder and sets the default values.
func NewItemRemovekeyRemoveKeyRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemovekeyRemoveKeyRequestBuilder) {
    m := &ItemRemovekeyRemoveKeyRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/servicePrincipals/{servicePrincipal%2Did}/removeKey", pathParameters),
    }
    return m
}
// NewItemRemovekeyRemoveKeyRequestBuilder instantiates a new ItemRemovekeyRemoveKeyRequestBuilder and sets the default values.
func NewItemRemovekeyRemoveKeyRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemRemovekeyRemoveKeyRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemRemovekeyRemoveKeyRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove a key credential from a servicePrincipal. This method along with addKey can be used by a servicePrincipal to automate rolling its expiring keys. As part of the request validation for this method, a proof of possession of an existing key is verified before the action can be performed.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/serviceprincipal-removekey?view=graph-rest-1.0
func (m *ItemRemovekeyRemoveKeyRequestBuilder) Post(ctx context.Context, body ItemRemovekeyRemoveKeyPostRequestBodyable, requestConfiguration *ItemRemovekeyRemoveKeyRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation remove a key credential from a servicePrincipal. This method along with addKey can be used by a servicePrincipal to automate rolling its expiring keys. As part of the request validation for this method, a proof of possession of an existing key is verified before the action can be performed.
// returns a *RequestInformation when successful
func (m *ItemRemovekeyRemoveKeyRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemRemovekeyRemoveKeyPostRequestBodyable, requestConfiguration *ItemRemovekeyRemoveKeyRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemRemovekeyRemoveKeyRequestBuilder when successful
func (m *ItemRemovekeyRemoveKeyRequestBuilder) WithUrl(rawUrl string)(*ItemRemovekeyRemoveKeyRequestBuilder) {
    return NewItemRemovekeyRemoveKeyRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
