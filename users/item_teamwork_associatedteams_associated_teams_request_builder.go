package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder provides operations to manage the associatedTeams property of the microsoft.graph.userTeamwork entity.
type ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderGetQueryParameters the list of associatedTeamInfo objects that a user is associated with.
type ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderGetQueryParameters
}
// ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAssociatedTeamInfoId provides operations to manage the associatedTeams property of the microsoft.graph.userTeamwork entity.
// returns a *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder when successful
func (m *ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder) ByAssociatedTeamInfoId(associatedTeamInfoId string)(*ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if associatedTeamInfoId != "" {
        urlTplParams["associatedTeamInfo%2Did"] = associatedTeamInfoId
    }
    return NewItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderInternal instantiates a new ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder and sets the default values.
func NewItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder) {
    m := &ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/teamwork/associatedTeams{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder instantiates a new ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder and sets the default values.
func NewItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTeamworkAssociatedteamsCountRequestBuilder when successful
func (m *ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder) Count()(*ItemTeamworkAssociatedteamsCountRequestBuilder) {
    return NewItemTeamworkAssociatedteamsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get the list of associatedTeamInfo objects that a user is associated with.
// returns a AssociatedTeamInfoCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssociatedTeamInfoCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAssociatedTeamInfoCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssociatedTeamInfoCollectionResponseable), nil
}
// Post create new navigation property to associatedTeams for users
// returns a AssociatedTeamInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssociatedTeamInfoable, requestConfiguration *ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssociatedTeamInfoable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAssociatedTeamInfoFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssociatedTeamInfoable), nil
}
// ToGetRequestInformation the list of associatedTeamInfo objects that a user is associated with.
// returns a *RequestInformation when successful
func (m *ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to associatedTeams for users
// returns a *RequestInformation when successful
func (m *ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssociatedTeamInfoable, requestConfiguration *ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder when successful
func (m *ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder) WithUrl(rawUrl string)(*ItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder) {
    return NewItemTeamworkAssociatedteamsAssociatedTeamsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
