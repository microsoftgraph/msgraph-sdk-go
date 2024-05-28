package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder provides operations to manage the assignmentSettings property of the microsoft.graph.educationClass entity.
type ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetQueryParameters specifies class-level assignments settings.
type ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetQueryParameters
}
// ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderInternal instantiates a new ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder and sets the default values.
func NewClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) {
    m := &ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/classes/{educationClass%2Did}/assignmentSettings{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder instantiates a new ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder and sets the default values.
func NewClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignmentSettings for education
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) Delete(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get specifies class-level assignments settings.
// returns a EducationAssignmentSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) Get(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentSettingsable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationAssignmentSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentSettingsable), nil
}
// GradingCategories provides operations to manage the gradingCategories property of the microsoft.graph.educationAssignmentSettings entity.
// returns a *ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder when successful
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) GradingCategories()(*ClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilder) {
    return NewClassesItemAssignmentsettingsGradingcategoriesGradingCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the properties of an educationAssignmentSettings object. Only teachers can update these settings.
// returns a EducationAssignmentSettingsable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationassignmentsettings-update?view=graph-rest-1.0
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentSettingsable, requestConfiguration *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentSettingsable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationAssignmentSettingsFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentSettingsable), nil
}
// ToDeleteRequestInformation delete navigation property assignmentSettings for education
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation specifies class-level assignments settings.
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of an educationAssignmentSettings object. Only teachers can update these settings.
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentSettingsable, requestConfiguration *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder when successful
func (m *ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) WithUrl(rawUrl string)(*ClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder) {
    return NewClassesItemAssignmentsettingsAssignmentSettingsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
