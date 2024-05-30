package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder provides operations to manage the checklistItems property of the microsoft.graph.todoTask entity.
type ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderGetQueryParameters a collection of checklistItems linked to a task.
type ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderGetQueryParameters struct {
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
// ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderGetQueryParameters
}
// ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByChecklistItemId provides operations to manage the checklistItems property of the microsoft.graph.todoTask entity.
// returns a *ItemTodoListsItemTasksItemChecklistitemsChecklistItemItemRequestBuilder when successful
func (m *ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder) ByChecklistItemId(checklistItemId string)(*ItemTodoListsItemTasksItemChecklistitemsChecklistItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if checklistItemId != "" {
        urlTplParams["checklistItem%2Did"] = checklistItemId
    }
    return NewItemTodoListsItemTasksItemChecklistitemsChecklistItemItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderInternal instantiates a new ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder and sets the default values.
func NewItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder) {
    m := &ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/todo/lists/{todoTaskList%2Did}/tasks/{todoTask%2Did}/checklistItems{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder instantiates a new ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder and sets the default values.
func NewItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ItemTodoListsItemTasksItemChecklistitemsCountRequestBuilder when successful
func (m *ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder) Count()(*ItemTodoListsItemTasksItemChecklistitemsCountRequestBuilder) {
    return NewItemTodoListsItemTasksItemChecklistitemsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get a collection of checklistItems linked to a task.
// returns a ChecklistItemCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChecklistItemCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChecklistItemCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChecklistItemCollectionResponseable), nil
}
// Post create new navigation property to checklistItems for users
// returns a ChecklistItemable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChecklistItemable, requestConfiguration *ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChecklistItemable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateChecklistItemFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChecklistItemable), nil
}
// ToGetRequestInformation a collection of checklistItems linked to a task.
// returns a *RequestInformation when successful
func (m *ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to checklistItems for users
// returns a *RequestInformation when successful
func (m *ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ChecklistItemable, requestConfiguration *ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder when successful
func (m *ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder) WithUrl(rawUrl string)(*ItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder) {
    return NewItemTodoListsItemTasksItemChecklistitemsChecklistItemsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
