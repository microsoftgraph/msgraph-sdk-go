package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder provides operations to manage the associatedTeams property of the microsoft.graph.userTeamwork entity.
type ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderGetQueryParameters the list of associatedTeamInfo objects that a user is associated with.
type ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderGetQueryParameters
}
// ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderInternal instantiates a new ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder and sets the default values.
func NewItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder) {
    m := &ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/teamwork/associatedTeams/{associatedTeamInfo%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder instantiates a new ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder and sets the default values.
func NewItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property associatedTeams for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
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
// Get the list of associatedTeamInfo objects that a user is associated with.
// returns a AssociatedTeamInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssociatedTeamInfoable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property associatedTeams in users
// returns a AssociatedTeamInfoable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssociatedTeamInfoable, requestConfiguration *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssociatedTeamInfoable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// Team provides operations to manage the team property of the microsoft.graph.teamInfo entity.
// returns a *ItemTeamworkAssociatedteamsItemTeamRequestBuilder when successful
func (m *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder) Team()(*ItemTeamworkAssociatedteamsItemTeamRequestBuilder) {
    return NewItemTeamworkAssociatedteamsItemTeamRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property associatedTeams for users
// returns a *RequestInformation when successful
func (m *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the list of associatedTeamInfo objects that a user is associated with.
// returns a *RequestInformation when successful
func (m *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property associatedTeams in users
// returns a *RequestInformation when successful
func (m *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssociatedTeamInfoable, requestConfiguration *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder when successful
func (m *ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder) WithUrl(rawUrl string)(*ItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder) {
    return NewItemTeamworkAssociatedteamsAssociatedTeamInfoItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
