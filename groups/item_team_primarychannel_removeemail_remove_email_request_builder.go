package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder provides operations to call the removeEmail method.
type ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilderInternal instantiates a new ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder and sets the default values.
func NewItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder) {
    m := &ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/team/primaryChannel/removeEmail", pathParameters),
    }
    return m
}
// NewItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder instantiates a new ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder and sets the default values.
func NewItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilderInternal(urlParams, requestAdapter)
}
// Post remove the email address of a channel. You can remove an email address only if it was provisioned using the provisionEmail method or through the Microsoft Teams client.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/channel-removeemail?view=graph-rest-1.0
func (m *ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation remove the email address of a channel. You can remove an email address only if it was provisioned using the provisionEmail method or through the Microsoft Teams client.
// returns a *RequestInformation when successful
func (m *ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder when successful
func (m *ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder) WithUrl(rawUrl string)(*ItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder) {
    return NewItemTeamPrimarychannelRemoveemailRemoveEmailRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
