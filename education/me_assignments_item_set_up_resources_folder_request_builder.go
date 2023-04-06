package education

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// MeAssignmentsItemSetUpResourcesFolderRequestBuilder provides operations to call the setUpResourcesFolder method.
type MeAssignmentsItemSetUpResourcesFolderRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// MeAssignmentsItemSetUpResourcesFolderRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type MeAssignmentsItemSetUpResourcesFolderRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewMeAssignmentsItemSetUpResourcesFolderRequestBuilderInternal instantiates a new SetUpResourcesFolderRequestBuilder and sets the default values.
func NewMeAssignmentsItemSetUpResourcesFolderRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSetUpResourcesFolderRequestBuilder) {
    m := &MeAssignmentsItemSetUpResourcesFolderRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/education/me/assignments/{educationAssignment%2Did}/setUpResourcesFolder", pathParameters),
    }
    return m
}
// NewMeAssignmentsItemSetUpResourcesFolderRequestBuilder instantiates a new SetUpResourcesFolderRequestBuilder and sets the default values.
func NewMeAssignmentsItemSetUpResourcesFolderRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*MeAssignmentsItemSetUpResourcesFolderRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeAssignmentsItemSetUpResourcesFolderRequestBuilderInternal(urlParams, requestAdapter)
}
// Post create a SharePoint folder to upload files for a given educationAssignment. Only teachers can perform this operation. The teacher determines the resources to upload in the assignment's folder. 
// [Find more info here]
// 
// [Find more info here]: https://docs.microsoft.com/graph/api/educationassignment-setupresourcesfolder?view=graph-rest-1.0
func (m *MeAssignmentsItemSetUpResourcesFolderRequestBuilder) Post(ctx context.Context, requestConfiguration *MeAssignmentsItemSetUpResourcesFolderRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationAssignmentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable), nil
}
// ToPostRequestInformation create a SharePoint folder to upload files for a given educationAssignment. Only teachers can perform this operation. The teacher determines the resources to upload in the assignment's folder. 
func (m *MeAssignmentsItemSetUpResourcesFolderRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *MeAssignmentsItemSetUpResourcesFolderRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
