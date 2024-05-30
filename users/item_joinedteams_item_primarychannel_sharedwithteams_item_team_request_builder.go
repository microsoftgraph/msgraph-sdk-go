package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder provides operations to manage the team property of the microsoft.graph.teamInfo entity.
type ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilderGetQueryParameters get team from users
type ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilderGetQueryParameters
}
// NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilderInternal instantiates a new ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder) {
    m := &ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/joinedTeams/{team%2Did}/primaryChannel/sharedWithTeams/{sharedWithChannelTeamInfo%2Did}/team{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder instantiates a new ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder and sets the default values.
func NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get team from users
// returns a Teamable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTeamFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Teamable), nil
}
// ToGetRequestInformation get team from users
// returns a *RequestInformation when successful
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder when successful
func (m *ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder) WithUrl(rawUrl string)(*ItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder) {
    return NewItemJoinedteamsItemPrimarychannelSharedwithteamsItemTeamRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
