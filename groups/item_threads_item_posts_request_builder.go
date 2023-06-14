package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemThreadsItemPostsRequestBuilder provides operations to manage the posts property of the microsoft.graph.conversationThread entity.
type ItemThreadsItemPostsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemThreadsItemPostsRequestBuilderGetQueryParameters get the properties and relationships of a post in a specified thread. You can specify both the parent conversation and the thread, or, you can specify the thread without referencing the parent conversation. Since the **post** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **post** instance.
type ItemThreadsItemPostsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemThreadsItemPostsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemThreadsItemPostsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemThreadsItemPostsRequestBuilderGetQueryParameters
}
// ByPostId provides operations to manage the posts property of the microsoft.graph.conversationThread entity.
func (m *ItemThreadsItemPostsRequestBuilder) ByPostId(postId string)(*ItemThreadsItemPostsPostItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if postId != "" {
        urlTplParams["post%2Did"] = postId
    }
    return NewItemThreadsItemPostsPostItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemThreadsItemPostsRequestBuilderInternal instantiates a new PostsRequestBuilder and sets the default values.
func NewItemThreadsItemPostsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsRequestBuilder) {
    m := &ItemThreadsItemPostsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/threads/{conversationThread%2Did}/posts{?%24top,%24skip,%24filter,%24count,%24orderby,%24select,%24expand}", pathParameters),
    }
    return m
}
// NewItemThreadsItemPostsRequestBuilder instantiates a new PostsRequestBuilder and sets the default values.
func NewItemThreadsItemPostsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemThreadsItemPostsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemThreadsItemPostsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
func (m *ItemThreadsItemPostsRequestBuilder) Count()(*ItemThreadsItemPostsCountRequestBuilder) {
    return NewItemThreadsItemPostsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the properties and relationships of a post in a specified thread. You can specify both the parent conversation and the thread, or, you can specify the thread without referencing the parent conversation. Since the **post** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **post** instance.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/post-get?view=graph-rest-1.0
func (m *ItemThreadsItemPostsRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemThreadsItemPostsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PostCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePostCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PostCollectionResponseable), nil
}
// ToGetRequestInformation get the properties and relationships of a post in a specified thread. You can specify both the parent conversation and the thread, or, you can specify the thread without referencing the parent conversation. Since the **post** resource supports extensions, you can also use the `GET` operation to get custom properties and extension data in a **post** instance.
func (m *ItemThreadsItemPostsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemThreadsItemPostsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
