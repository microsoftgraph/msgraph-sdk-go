package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemCloneRequestBuilder provides operations to call the clone method.
type ItemCloneRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemCloneRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemCloneRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemCloneRequestBuilderInternal instantiates a new CloneRequestBuilder and sets the default values.
func NewItemCloneRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloneRequestBuilder) {
    m := &ItemCloneRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/clone", pathParameters),
    }
    return m
}
// NewItemCloneRequestBuilder instantiates a new CloneRequestBuilder and sets the default values.
func NewItemCloneRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemCloneRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemCloneRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a copy of a team. This operation also creates a copy of the corresponding group.You can specify which parts of the team to clone: When tabs are cloned, they aren't configured. The tabs are displayed on the tab bar in Microsoft Teams, and the first time a user opens them, they must go through the configuration screen. If the user who opens the tab doesn't have permission to configure apps, they'll see a message that says that the tab isn't configured. Cloning is a long-running operation. After the POST clone returns, you need to GET the operation returned by the Location: header to see if it's running, succeeded, or failed. You should continue to GET until the status isn't running. The recommended delay between GETs is 5 seconds.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-clone?view=graph-rest-1.0
func (m *ItemCloneRequestBuilder) Post(ctx context.Context, body ItemClonePostRequestBodyable, requestConfiguration *ItemCloneRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// ToPostRequestInformation create a copy of a team. This operation also creates a copy of the corresponding group.You can specify which parts of the team to clone: When tabs are cloned, they aren't configured. The tabs are displayed on the tab bar in Microsoft Teams, and the first time a user opens them, they must go through the configuration screen. If the user who opens the tab doesn't have permission to configure apps, they'll see a message that says that the tab isn't configured. Cloning is a long-running operation. After the POST clone returns, you need to GET the operation returned by the Location: header to see if it's running, succeeded, or failed. You should continue to GET until the status isn't running. The recommended delay between GETs is 5 seconds.
func (m *ItemCloneRequestBuilder) ToPostRequestInformation(ctx context.Context, body ItemClonePostRequestBodyable, requestConfiguration *ItemCloneRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemCloneRequestBuilder) WithUrl(rawUrl string)(*ItemCloneRequestBuilder) {
    return NewItemCloneRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
