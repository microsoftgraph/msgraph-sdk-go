package employeeexperience

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LearningprovidersItemLearningcontentsLearningContentsRequestBuilder provides operations to manage the learningContents property of the microsoft.graph.learningProvider entity.
type LearningprovidersItemLearningcontentsLearningContentsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LearningprovidersItemLearningcontentsLearningContentsRequestBuilderGetQueryParameters get a list of the learningContent resources and their properties. This list represents the metadata of the specified provider's content in Viva Learning.
type LearningprovidersItemLearningcontentsLearningContentsRequestBuilderGetQueryParameters struct {
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
// LearningprovidersItemLearningcontentsLearningContentsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningprovidersItemLearningcontentsLearningContentsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LearningprovidersItemLearningcontentsLearningContentsRequestBuilderGetQueryParameters
}
// LearningprovidersItemLearningcontentsLearningContentsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningprovidersItemLearningcontentsLearningContentsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByLearningContentId provides operations to manage the learningContents property of the microsoft.graph.learningProvider entity.
// returns a *LearningprovidersItemLearningcontentsLearningContentItemRequestBuilder when successful
func (m *LearningprovidersItemLearningcontentsLearningContentsRequestBuilder) ByLearningContentId(learningContentId string)(*LearningprovidersItemLearningcontentsLearningContentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if learningContentId != "" {
        urlTplParams["learningContent%2Did"] = learningContentId
    }
    return NewLearningprovidersItemLearningcontentsLearningContentItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLearningprovidersItemLearningcontentsLearningContentsRequestBuilderInternal instantiates a new LearningprovidersItemLearningcontentsLearningContentsRequestBuilder and sets the default values.
func NewLearningprovidersItemLearningcontentsLearningContentsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LearningprovidersItemLearningcontentsLearningContentsRequestBuilder) {
    m := &LearningprovidersItemLearningcontentsLearningContentsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/employeeExperience/learningProviders/{learningProvider%2Did}/learningContents{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewLearningprovidersItemLearningcontentsLearningContentsRequestBuilder instantiates a new LearningprovidersItemLearningcontentsLearningContentsRequestBuilder and sets the default values.
func NewLearningprovidersItemLearningcontentsLearningContentsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LearningprovidersItemLearningcontentsLearningContentsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLearningprovidersItemLearningcontentsLearningContentsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *LearningprovidersItemLearningcontentsCountRequestBuilder when successful
func (m *LearningprovidersItemLearningcontentsLearningContentsRequestBuilder) Count()(*LearningprovidersItemLearningcontentsCountRequestBuilder) {
    return NewLearningprovidersItemLearningcontentsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the learningContent resources and their properties. This list represents the metadata of the specified provider's content in Viva Learning.
// returns a LearningContentCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/learningprovider-list-learningcontents?view=graph-rest-1.0
func (m *LearningprovidersItemLearningcontentsLearningContentsRequestBuilder) Get(ctx context.Context, requestConfiguration *LearningprovidersItemLearningcontentsLearningContentsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningContentCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLearningContentCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningContentCollectionResponseable), nil
}
// Post create new navigation property to learningContents for employeeExperience
// returns a LearningContentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LearningprovidersItemLearningcontentsLearningContentsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningContentable, requestConfiguration *LearningprovidersItemLearningcontentsLearningContentsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningContentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLearningContentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningContentable), nil
}
// ToGetRequestInformation get a list of the learningContent resources and their properties. This list represents the metadata of the specified provider's content in Viva Learning.
// returns a *RequestInformation when successful
func (m *LearningprovidersItemLearningcontentsLearningContentsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LearningprovidersItemLearningcontentsLearningContentsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to learningContents for employeeExperience
// returns a *RequestInformation when successful
func (m *LearningprovidersItemLearningcontentsLearningContentsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningContentable, requestConfiguration *LearningprovidersItemLearningcontentsLearningContentsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LearningprovidersItemLearningcontentsLearningContentsRequestBuilder when successful
func (m *LearningprovidersItemLearningcontentsLearningContentsRequestBuilder) WithUrl(rawUrl string)(*LearningprovidersItemLearningcontentsLearningContentsRequestBuilder) {
    return NewLearningprovidersItemLearningcontentsLearningContentsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
