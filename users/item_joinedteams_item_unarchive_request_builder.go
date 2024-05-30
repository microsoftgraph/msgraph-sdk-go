package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemUnarchiveRequestBuilder provides operations to call the unarchive method.
type ItemJoinedteamsItemUnarchiveRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemUnarchiveRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemUnarchiveRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemJoinedteamsItemUnarchiveRequestBuilderInternal instantiates a new ItemJoinedteamsItemUnarchiveRequestBuilder and sets the default values.
func NewItemJoinedteamsItemUnarchiveRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemUnarchiveRequestBuilder) {
    m := &ItemJoinedteamsItemUnarchiveRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/unarchive", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemUnarchiveRequestBuilder instantiates a new ItemJoinedteamsItemUnarchiveRequestBuilder and sets the default values.
func NewItemJoinedteamsItemUnarchiveRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemUnarchiveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemUnarchiveRequestBuilderInternal(urlParams, requestAdapter)
}
// Post restore an archived team. This API restores users' ability to send messages and edit the team, abiding by tenant and team settings. A Team is archived using the archive API. Unarchiving is an async operation. A team is unarchived once the async operation completes successfully, which might occur subsequent to a response from this API.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/team-unarchive?view=graph-rest-1.0
func (m *ItemJoinedteamsItemUnarchiveRequestBuilder) Post(ctx context.Context, requestConfiguration *ItemJoinedteamsItemUnarchiveRequestBuilderPostRequestConfiguration)(error) {
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
// ToPostRequestInformation restore an archived team. This API restores users' ability to send messages and edit the team, abiding by tenant and team settings. A Team is archived using the archive API. Unarchiving is an async operation. A team is unarchived once the async operation completes successfully, which might occur subsequent to a response from this API.
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemUnarchiveRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemUnarchiveRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemJoinedteamsItemUnarchiveRequestBuilder when successful
func (m *ItemJoinedteamsItemUnarchiveRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemUnarchiveRequestBuilder) {
    return NewItemJoinedteamsItemUnarchiveRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
