package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemActivitiesItemHistoryitemsItemActivityRequestBuilder provides operations to manage the activity property of the microsoft.graph.activityHistoryItem entity.
type ItemActivitiesItemHistoryitemsItemActivityRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActivitiesItemHistoryitemsItemActivityRequestBuilderGetQueryParameters optional. NavigationProperty/Containment; navigation property to the associated activity.
type ItemActivitiesItemHistoryitemsItemActivityRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemActivitiesItemHistoryitemsItemActivityRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActivitiesItemHistoryitemsItemActivityRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemActivitiesItemHistoryitemsItemActivityRequestBuilderGetQueryParameters
}
// NewItemActivitiesItemHistoryitemsItemActivityRequestBuilderInternal instantiates a new ItemActivitiesItemHistoryitemsItemActivityRequestBuilder and sets the default values.
func NewItemActivitiesItemHistoryitemsItemActivityRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActivitiesItemHistoryitemsItemActivityRequestBuilder) {
    m := &ItemActivitiesItemHistoryitemsItemActivityRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/activities/{userActivity%2Did}/historyItems/{activityHistoryItem%2Did}/activity{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemActivitiesItemHistoryitemsItemActivityRequestBuilder instantiates a new ItemActivitiesItemHistoryitemsItemActivityRequestBuilder and sets the default values.
func NewItemActivitiesItemHistoryitemsItemActivityRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActivitiesItemHistoryitemsItemActivityRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActivitiesItemHistoryitemsItemActivityRequestBuilderInternal(urlParams, requestAdapter)
}
// Get optional. NavigationProperty/Containment; navigation property to the associated activity.
// returns a UserActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemActivitiesItemHistoryitemsItemActivityRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActivitiesItemHistoryitemsItemActivityRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserActivityable), nil
}
// ToGetRequestInformation optional. NavigationProperty/Containment; navigation property to the associated activity.
// returns a *RequestInformation when successful
func (m *ItemActivitiesItemHistoryitemsItemActivityRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActivitiesItemHistoryitemsItemActivityRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemActivitiesItemHistoryitemsItemActivityRequestBuilder when successful
func (m *ItemActivitiesItemHistoryitemsItemActivityRequestBuilder) WithUrl(rawUrl string)(*ItemActivitiesItemHistoryitemsItemActivityRequestBuilder) {
    return NewItemActivitiesItemHistoryitemsItemActivityRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
