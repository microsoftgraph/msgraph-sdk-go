package teams

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemUnarchiveRequestBuilder provides operations to call the unarchive method.
type ItemUnarchiveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemUnarchiveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemUnarchiveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemUnarchiveRequestBuilderInternal instantiates a new UnarchiveRequestBuilder and sets the default values.
func NewItemUnarchiveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnarchiveRequestBuilder) {
    m := &ItemUnarchiveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teams/{team%2Did}/unarchive", pathParameters),
    }
    return m
}
// NewItemUnarchiveRequestBuilder instantiates a new UnarchiveRequestBuilder and sets the default values.
func NewItemUnarchiveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemUnarchiveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemUnarchiveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore an archived team. This restores users' ability to send messages and edit the team, abiding by tenant and team settings. A Team is archived using the archive API. Unarchiving is an async operation. A team is unarchived once the async operation completes successfully, which may occur subsequent to a response from this API.
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-unarchive?view=graph-rest-1.0
func (m *ItemUnarchiveRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemUnarchiveRequestBuilderPostRequestConfiguration)(error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
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
// ToPostRequestInformation restore an archived team. This restores users' ability to send messages and edit the team, abiding by tenant and team settings. A Team is archived using the archive API. Unarchiving is an async operation. A team is unarchived once the async operation completes successfully, which may occur subsequent to a response from this API.
func (m *ItemUnarchiveRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemUnarchiveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemUnarchiveRequestBuilder) WithUrl(rawUrl string)(*ItemUnarchiveRequestBuilder) {
    return NewItemUnarchiveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
