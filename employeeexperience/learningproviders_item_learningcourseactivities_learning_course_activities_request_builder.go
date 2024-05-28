package employeeexperience

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder provides operations to manage the learningCourseActivities property of the microsoft.graph.learningProvider entity.
type LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderGetQueryParameters get the specified learningCourseActivity object using either an ID or an externalCourseActivityId of the learning provider, or a courseActivityId of a user.
type LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderGetQueryParameters struct {
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
// LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderGetQueryParameters
}
// LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByLearningCourseActivityId provides operations to manage the learningCourseActivities property of the microsoft.graph.learningProvider entity.
// returns a *LearningprovidersItemLearningcourseactivitiesLearningCourseActivityItemRequestBuilder when successful
func (m *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder) ByLearningCourseActivityId(learningCourseActivityId string)(*LearningprovidersItemLearningcourseactivitiesLearningCourseActivityItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if learningCourseActivityId != "" {
        urlTplParams["learningCourseActivity%2Did"] = learningCourseActivityId
    }
    return NewLearningprovidersItemLearningcourseactivitiesLearningCourseActivityItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewLearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderInternal instantiates a new LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder and sets the default values.
func NewLearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder) {
    m := &LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/employeeExperience/learningProviders/{learningProvider%2Did}/learningCourseActivities{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewLearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder instantiates a new LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder and sets the default values.
func NewLearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *LearningprovidersItemLearningcourseactivitiesCountRequestBuilder when successful
func (m *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder) Count()(*LearningprovidersItemLearningcourseactivitiesCountRequestBuilder) {
    return NewLearningprovidersItemLearningcourseactivitiesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get the specified learningCourseActivity object using either an ID or an externalCourseActivityId of the learning provider, or a courseActivityId of a user.
// returns a LearningCourseActivityCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/learningcourseactivity-get?view=graph-rest-1.0
func (m *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder) Get(ctx context.Context, requestConfiguration *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLearningCourseActivityCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityCollectionResponseable), nil
}
// Post create a new learningCourseActivity object. A learning course activity can be one of two types: - Assignment- Self-initiated Use this method to create either type of activity.
// returns a LearningCourseActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/employeeexperienceuser-post-learningcourseactivities?view=graph-rest-1.0
func (m *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityable, requestConfiguration *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLearningCourseActivityFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityable), nil
}
// ToGetRequestInformation get the specified learningCourseActivity object using either an ID or an externalCourseActivityId of the learning provider, or a courseActivityId of a user.
// returns a *RequestInformation when successful
func (m *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new learningCourseActivity object. A learning course activity can be one of two types: - Assignment- Self-initiated Use this method to create either type of activity.
// returns a *RequestInformation when successful
func (m *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityable, requestConfiguration *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder when successful
func (m *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder) WithUrl(rawUrl string)(*LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder) {
    return NewLearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
