package employeeexperience

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// LearningprovidersLearningProviderItemRequestBuilder provides operations to manage the learningProviders property of the microsoft.graph.employeeExperience entity.
type LearningprovidersLearningProviderItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LearningprovidersLearningProviderItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningprovidersLearningProviderItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LearningprovidersLearningProviderItemRequestBuilderGetQueryParameters read the properties and relationships of a learningProvider object.
type LearningprovidersLearningProviderItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LearningprovidersLearningProviderItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningprovidersLearningProviderItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LearningprovidersLearningProviderItemRequestBuilderGetQueryParameters
}
// LearningprovidersLearningProviderItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LearningprovidersLearningProviderItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLearningprovidersLearningProviderItemRequestBuilderInternal instantiates a new LearningprovidersLearningProviderItemRequestBuilder and sets the default values.
func NewLearningprovidersLearningProviderItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LearningprovidersLearningProviderItemRequestBuilder) {
    m := &LearningprovidersLearningProviderItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/employeeExperience/learningProviders/{learningProvider%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLearningprovidersLearningProviderItemRequestBuilder instantiates a new LearningprovidersLearningProviderItemRequestBuilder and sets the default values.
func NewLearningprovidersLearningProviderItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LearningprovidersLearningProviderItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLearningprovidersLearningProviderItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a learningProvider resource and remove its registration in Viva Learning for a tenant.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/employeeexperience-delete-learningproviders?view=graph-rest-1.0
func (m *LearningprovidersLearningProviderItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *LearningprovidersLearningProviderItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a learningProvider object.
// returns a LearningProviderable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/learningprovider-get?view=graph-rest-1.0
func (m *LearningprovidersLearningProviderItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LearningprovidersLearningProviderItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningProviderable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLearningProviderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningProviderable), nil
}
// LearningContents provides operations to manage the learningContents property of the microsoft.graph.learningProvider entity.
// returns a *LearningprovidersItemLearningcontentsLearningContentsRequestBuilder when successful
func (m *LearningprovidersLearningProviderItemRequestBuilder) LearningContents()(*LearningprovidersItemLearningcontentsLearningContentsRequestBuilder) {
    return NewLearningprovidersItemLearningcontentsLearningContentsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LearningContentsWithExternalId provides operations to manage the learningContents property of the microsoft.graph.learningProvider entity.
// returns a *LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder when successful
func (m *LearningprovidersLearningProviderItemRequestBuilder) LearningContentsWithExternalId(externalId *string)(*LearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilder) {
    return NewLearningprovidersItemLearningcontentswithexternalidLearningContentsWithExternalIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, externalId)
}
// LearningCourseActivities provides operations to manage the learningCourseActivities property of the microsoft.graph.learningProvider entity.
// returns a *LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder when successful
func (m *LearningprovidersLearningProviderItemRequestBuilder) LearningCourseActivities()(*LearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilder) {
    return NewLearningprovidersItemLearningcourseactivitiesLearningCourseActivitiesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LearningCourseActivitiesWithExternalcourseActivityId provides operations to manage the learningCourseActivities property of the microsoft.graph.learningProvider entity.
// returns a *LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder when successful
func (m *LearningprovidersLearningProviderItemRequestBuilder) LearningCourseActivitiesWithExternalcourseActivityId(externalcourseActivityId *string)(*LearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilder) {
    return NewLearningprovidersItemLearningcourseactivitieswithexternalcourseactivityidLearningCourseActivitiesWithExternalcourseActivityIdRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, externalcourseActivityId)
}
// Patch update the properties of a learningProvider object.
// returns a LearningProviderable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/learningprovider-update?view=graph-rest-1.0
func (m *LearningprovidersLearningProviderItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningProviderable, requestConfiguration *LearningprovidersLearningProviderItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningProviderable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLearningProviderFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningProviderable), nil
}
// ToDeleteRequestInformation delete a learningProvider resource and remove its registration in Viva Learning for a tenant.
// returns a *RequestInformation when successful
func (m *LearningprovidersLearningProviderItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LearningprovidersLearningProviderItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a learningProvider object.
// returns a *RequestInformation when successful
func (m *LearningprovidersLearningProviderItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LearningprovidersLearningProviderItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a learningProvider object.
// returns a *RequestInformation when successful
func (m *LearningprovidersLearningProviderItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LearningProviderable, requestConfiguration *LearningprovidersLearningProviderItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LearningprovidersLearningProviderItemRequestBuilder when successful
func (m *LearningprovidersLearningProviderItemRequestBuilder) WithUrl(rawUrl string)(*LearningprovidersLearningProviderItemRequestBuilder) {
    return NewLearningprovidersLearningProviderItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
