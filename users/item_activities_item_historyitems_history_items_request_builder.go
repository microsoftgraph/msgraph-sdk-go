package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder provides operations to manage the historyItems property of the microsoft.graph.userActivity entity.
type ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderGetQueryParameters optional. NavigationProperty/Containment; navigation property to the activity's historyItems.
type ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderGetQueryParameters struct {
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
// ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderGetQueryParameters
}
// ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByActivityHistoryItemId provides operations to manage the historyItems property of the microsoft.graph.userActivity entity.
// returns a *ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder when successful
func (m *ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder) ByActivityHistoryItemId(activityHistoryItemId string)(*ItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if activityHistoryItemId != "" {
        urlTplParams["activityHistoryItem%2Did"] = activityHistoryItemId
    }
    return NewItemActivitiesItemHistoryitemsActivityHistoryItemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderInternal instantiates a new ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder and sets the default values.
func NewItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder) {
    m := &ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/activities/{userActivity%2Did}/historyItems{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder instantiates a new ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder and sets the default values.
func NewItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemActivitiesItemHistoryitemsCountRequestBuilder when successful
func (m *ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder) Count()(*ItemActivitiesItemHistoryitemsCountRequestBuilder) {
    return NewItemActivitiesItemHistoryitemsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get optional. NavigationProperty/Containment; navigation property to the activity's historyItems.
// returns a ActivityHistoryItemCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ActivityHistoryItemCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateActivityHistoryItemCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ActivityHistoryItemCollectionResponseable), nil
}
// Post create new navigation property to historyItems for users
// returns a ActivityHistoryItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ActivityHistoryItemable, requestConfiguration *ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ActivityHistoryItemable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateActivityHistoryItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ActivityHistoryItemable), nil
}
// ToGetRequestInformation optional. NavigationProperty/Containment; navigation property to the activity's historyItems.
// returns a *RequestInformation when successful
func (m *ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to historyItems for users
// returns a *RequestInformation when successful
func (m *ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ActivityHistoryItemable, requestConfiguration *ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder when successful
func (m *ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder) WithUrl(rawUrl string)(*ItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder) {
    return NewItemActivitiesItemHistoryitemsHistoryItemsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
