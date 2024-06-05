package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder provides operations to manage the submittedResources property of the microsoft.graph.educationSubmission entity.
type ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderGetQueryParameters retrieve a submitted resource. Only teachers, students, and applications with application permissions can perform this operation. Resources are available to a teacher or an application with application permissions after a student submits it, and are available to the student after the teacher returns the submission. Teachers can leave notes in some resources.
type ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderGetQueryParameters
}
// ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderInternal instantiates a new ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder) {
    m := &ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/classes/{educationClass%2Did}/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/submittedResources/{educationSubmissionResource%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder instantiates a new ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property submittedResources for education
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve a submitted resource. Only teachers, students, and applications with application permissions can perform this operation. Resources are available to a teacher or an application with application permissions after a student submits it, and are available to the student after the teacher returns the submission. Teachers can leave notes in some resources.
// returns a EducationSubmissionResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/educationsubmittedsubmissionresource-get?view=graph-rest-1.0
func (m *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionResourceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationSubmissionResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionResourceable), nil
}
// Patch update the navigation property submittedResources in education
// returns a EducationSubmissionResourceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionResourceable, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionResourceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationSubmissionResourceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionResourceable), nil
}
// ToDeleteRequestInformation delete navigation property submittedResources for education
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve a submitted resource. Only teachers, students, and applications with application permissions can perform this operation. Resources are available to a teacher or an application with application permissions after a student submits it, and are available to the student after the teacher returns the submission. Teachers can leave notes in some resources.
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property submittedResources in education
// returns a *RequestInformation when successful
func (m *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionResourceable, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder when successful
func (m *ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder) WithUrl(rawUrl string)(*ClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder) {
    return NewClassesItemAssignmentsItemSubmissionsItemSubmittedresourcesEducationSubmissionResourceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
