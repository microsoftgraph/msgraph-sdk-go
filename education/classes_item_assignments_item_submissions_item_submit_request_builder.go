package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilder provides operations to call the submit method.
type ClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilderInternal instantiates a new SubmitRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilder) {
    m := &ClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass%2Did}/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/microsoft.graph.submit";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilder instantiates a new SubmitRequestBuilder and sets the default values.
func NewClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation indicate that a student is done with the work and is ready to hand in the assignment.  This method changes the status of the submission from `working` to `submitted`. During the submit process, all the resources are copied to the **submittedResources** bucket. The teacher will be looking at the submitted resources list for grading. A teacher can also submit a student's assignment on their behalf.
func (m *ClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post indicate that a student is done with the work and is ready to hand in the assignment.  This method changes the status of the submission from `working` to `submitted`. During the submit process, all the resources are copied to the **submittedResources** bucket. The teacher will be looking at the submitted resources list for grading. A teacher can also submit a student's assignment on their behalf.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/educationsubmission-submit?view=graph-rest-1.0
func (m *ClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilder) Post(ctx context.Context, requestConfiguration *ClassesItemAssignmentsItemSubmissionsItemSubmitRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, error) {
    requestInfo, err := m.CreatePostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationSubmissionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable), nil
}
