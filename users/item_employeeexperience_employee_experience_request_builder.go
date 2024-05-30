package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemEmployeeexperienceEmployeeExperienceRequestBuilder provides operations to manage the employeeExperience property of the microsoft.graph.user entity.
type ItemEmployeeexperienceEmployeeExperienceRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemEmployeeexperienceEmployeeExperienceRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEmployeeexperienceEmployeeExperienceRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemEmployeeexperienceEmployeeExperienceRequestBuilderGetQueryParameters get employeeExperience from users
type ItemEmployeeexperienceEmployeeExperienceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemEmployeeexperienceEmployeeExperienceRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEmployeeexperienceEmployeeExperienceRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemEmployeeexperienceEmployeeExperienceRequestBuilderGetQueryParameters
}
// ItemEmployeeexperienceEmployeeExperienceRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemEmployeeexperienceEmployeeExperienceRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemEmployeeexperienceEmployeeExperienceRequestBuilderInternal instantiates a new ItemEmployeeexperienceEmployeeExperienceRequestBuilder and sets the default values.
func NewItemEmployeeexperienceEmployeeExperienceRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEmployeeexperienceEmployeeExperienceRequestBuilder) {
    m := &ItemEmployeeexperienceEmployeeExperienceRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/employeeExperience{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemEmployeeexperienceEmployeeExperienceRequestBuilder instantiates a new ItemEmployeeexperienceEmployeeExperienceRequestBuilder and sets the default values.
func NewItemEmployeeexperienceEmployeeExperienceRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemEmployeeexperienceEmployeeExperienceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemEmployeeexperienceEmployeeExperienceRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property employeeExperience for users
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemEmployeeexperienceEmployeeExperienceRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemEmployeeexperienceEmployeeExperienceRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get employeeExperience from users
// returns a EmployeeExperienceUserable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemEmployeeexperienceEmployeeExperienceRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemEmployeeexperienceEmployeeExperienceRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EmployeeExperienceUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEmployeeExperienceUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EmployeeExperienceUserable), nil
}
// LearningCourseActivities provides operations to manage the learningCourseActivities property of the microsoft.graph.employeeExperienceUser entity.
// returns a *ItemEmployeeexperienceLearningcourseactivitiesLearningCourseActivitiesRequestBuilder when successful
func (m *ItemEmployeeexperienceEmployeeExperienceRequestBuilder) LearningCourseActivities()(*ItemEmployeeexperienceLearningcourseactivitiesLearningCourseActivitiesRequestBuilder) {
    return NewItemEmployeeexperienceLearningcourseactivitiesLearningCourseActivitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LearningCourseActivitiesWithExternalcourseActivityId provides operations to manage the learningCourseActivities property of the microsoft.graph.employeeExperienceUser entity.
// returns a *ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder when successful
func (m *ItemEmployeeexperienceEmployeeExperienceRequestBuilder) LearningCourseActivitiesWithExternalcourseActivityId(externalcourseActivityId *string)(*ItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) {
    return NewItemEmployeeexperienceLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, externalcourseActivityId)
}
// Patch update the navigation property employeeExperience in users
// returns a EmployeeExperienceUserable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemEmployeeexperienceEmployeeExperienceRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EmployeeExperienceUserable, requestConfiguration *ItemEmployeeexperienceEmployeeExperienceRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EmployeeExperienceUserable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEmployeeExperienceUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EmployeeExperienceUserable), nil
}
// ToDeleteRequestInformation delete navigation property employeeExperience for users
// returns a *RequestInformation when successful
func (m *ItemEmployeeexperienceEmployeeExperienceRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemEmployeeexperienceEmployeeExperienceRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get employeeExperience from users
// returns a *RequestInformation when successful
func (m *ItemEmployeeexperienceEmployeeExperienceRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemEmployeeexperienceEmployeeExperienceRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property employeeExperience in users
// returns a *RequestInformation when successful
func (m *ItemEmployeeexperienceEmployeeExperienceRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EmployeeExperienceUserable, requestConfiguration *ItemEmployeeexperienceEmployeeExperienceRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemEmployeeexperienceEmployeeExperienceRequestBuilder when successful
func (m *ItemEmployeeexperienceEmployeeExperienceRequestBuilder) WithUrl(rawUrl string)(*ItemEmployeeexperienceEmployeeExperienceRequestBuilder) {
    return NewItemEmployeeexperienceEmployeeExperienceRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
