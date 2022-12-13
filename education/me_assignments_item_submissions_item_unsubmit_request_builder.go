package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MeAssignmentsItemSubmissionsItemUnsubmitRequestBuilder provides operations to call the unsubmit method.
type MeAssignmentsItemSubmissionsItemUnsubmitRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// MeAssignmentsItemSubmissionsItemUnsubmitRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemSubmissionsItemUnsubmitRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeAssignmentsItemSubmissionsItemUnsubmitRequestBuilderInternal instantiates a new UnsubmitRequestBuilder and sets the default values.
func NewMeAssignmentsItemSubmissionsItemUnsubmitRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSubmissionsItemUnsubmitRequestBuilder) {
    m := &MeAssignmentsItemSubmissionsItemUnsubmitRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/submissions/{educationSubmission%2Did}/microsoft.graph.unsubmit";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeAssignmentsItemSubmissionsItemUnsubmitRequestBuilder instantiates a new UnsubmitRequestBuilder and sets the default values.
func NewMeAssignmentsItemSubmissionsItemUnsubmitRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSubmissionsItemUnsubmitRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsItemSubmissionsItemUnsubmitRequestBuilderInternal(urlParams, requestAdapter)
}
// CreatePostRequestInformation indicate that a student wants to work on the submission of the assignment after it was turned in.  This method changes the status of the submission from `submitted` to `working`. During the submit process, all the resources are copied from **submittedResources** to  **workingResources**. The teacher will be looking at the working resources list for grading. A teacher can also unsubmit a student's assignment on their behalf.
func (m *MeAssignmentsItemSubmissionsItemUnsubmitRequestBuilder) CreatePostRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsItemUnsubmitRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers["Accept"] = "application/json"
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Post indicate that a student wants to work on the submission of the assignment after it was turned in.  This method changes the status of the submission from `submitted` to `working`. During the submit process, all the resources are copied from **submittedResources** to  **workingResources**. The teacher will be looking at the working resources list for grading. A teacher can also unsubmit a student's assignment on their behalf.
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/educationsubmission-unsubmit?view=graph-rest-1.0
func (m *MeAssignmentsItemSubmissionsItemUnsubmitRequestBuilder) Post(ctx context.Context, requestConfiguration *MeAssignmentsItemSubmissionsItemUnsubmitRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationSubmissionable, error) {
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
