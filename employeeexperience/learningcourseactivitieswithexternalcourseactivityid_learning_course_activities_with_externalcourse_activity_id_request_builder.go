package employeeexperience

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder provides operations to manage the learningCourseActivities property of the microsoft.graph.employeeExperience entity.
type LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetQueryParameters get the specified learningCourseActivity object using either an ID or an externalCourseActivityId of the learning provider, or a courseActivityId of a user.
type LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetQueryParameters
}
// LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderInternal instantiates a new LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder and sets the default values.
func NewLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, externalcourseActivityId *string)(*LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) {
    m := &LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/employeeExperience/learningCourseActivities(externalcourseActivityId='{externalcourseActivityId}'){?%24expand,%24select}", pathParameters),
    }
    if externalcourseActivityId != nil {
        m.BaseRequestBuilder.PathParameters["externalcourseActivityId"] = *externalcourseActivityId
    }
    return m
}
// NewLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder instantiates a new LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder and sets the default values.
func NewLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property learningCourseActivities for employeeExperience
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) Delete(ctx context.Context, requestConfiguration *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the specified learningCourseActivity object using either an ID or an externalCourseActivityId of the learning provider, or a courseActivityId of a user.
// returns a LearningCourseActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/learningcourseactivity-get?view=graph-rest-1.0
func (m *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) Get(ctx context.Context, requestConfiguration *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property learningCourseActivities in employeeExperience
// returns a LearningCourseActivityable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityable, requestConfiguration *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property learningCourseActivities for employeeExperience
// returns a *RequestInformation when successful
func (m *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the specified learningCourseActivity object using either an ID or an externalCourseActivityId of the learning provider, or a courseActivityId of a user.
// returns a *RequestInformation when successful
func (m *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property learningCourseActivities in employeeExperience
// returns a *RequestInformation when successful
func (m *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningCourseActivityable, requestConfiguration *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder when successful
func (m *LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) WithUrl(rawUrl string)(*LearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) {
    return NewLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
