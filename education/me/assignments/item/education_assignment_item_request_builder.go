package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i608fa3df84918be40f4d961440af7a1df9cdc70cdcbcd01332c72421f74df332 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/categories"
    i6182b98a227da91892039c997d1742545d169f3ded56acfb9ebddc2b8754acb3 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/rubric"
    i710174f1350381dcd776625bf1b593be962f1d2c9b1e368ae0bd178d3495728f "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions"
    i74b91723ec2102e06ba61f2eec38b8a338e711092e99e7a6121ae3287f6f5788 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/publish"
    i874a2b579f5eb892e97d9c0c157220e5118544d0dca2de101640f4b07a1c3fe3 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/resources"
    ie33a7d66ff2752d51b9a63d7758fc6c00305b565d547fa3e4e42f07c7a26bdec "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/setupresourcesfolder"
    i0a97b1183e4efc9029dbfdb7442c76c2d4fc13a243fc9089be3106c63fd1b7f2 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/submissions/item"
    i26284194e56af98cb4a7e2e347a5485f31cf2c535beac0528c6c3661c312a162 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/categories/item"
    i7b969696e68843bc2d2d30122cafa6f7eaf8a9843bf04feaf8693f3f794c0d94 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item/resources/item"
)

// EducationAssignmentItemRequestBuilder provides operations to manage the assignments property of the microsoft.graph.educationUser entity.
type EducationAssignmentItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EducationAssignmentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationAssignmentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// EducationAssignmentItemRequestBuilderGetQueryParameters list of assignments for the user. Nullable.
type EducationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EducationAssignmentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationAssignmentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EducationAssignmentItemRequestBuilderGetQueryParameters
}
// EducationAssignmentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type EducationAssignmentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Categories the categories property
func (m *EducationAssignmentItemRequestBuilder) Categories()(*i608fa3df84918be40f4d961440af7a1df9cdc70cdcbcd01332c72421f74df332.CategoriesRequestBuilder) {
    return i608fa3df84918be40f4d961440af7a1df9cdc70cdcbcd01332c72421f74df332.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item.categories.item collection
func (m *EducationAssignmentItemRequestBuilder) CategoriesById(id string)(*i26284194e56af98cb4a7e2e347a5485f31cf2c535beac0528c6c3661c312a162.EducationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory%2Did"] = id
    }
    return i26284194e56af98cb4a7e2e347a5485f31cf2c535beac0528c6c3661c312a162.NewEducationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationAssignmentItemRequestBuilderInternal instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewEducationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationAssignmentItemRequestBuilder) {
    m := &EducationAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/me/assignments/{educationAssignment%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationAssignmentItemRequestBuilder instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewEducationAssignmentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationAssignmentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationAssignmentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property assignments for education
func (m *EducationAssignmentItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property assignments for education
func (m *EducationAssignmentItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *EducationAssignmentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation list of assignments for the user. Nullable.
func (m *EducationAssignmentItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration list of assignments for the user. Nullable.
func (m *EducationAssignmentItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *EducationAssignmentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property assignments in education
func (m *EducationAssignmentItemRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property assignments in education
func (m *EducationAssignmentItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, requestConfiguration *EducationAssignmentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property assignments for education
func (m *EducationAssignmentItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property assignments for education
func (m *EducationAssignmentItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *EducationAssignmentItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get list of assignments for the user. Nullable.
func (m *EducationAssignmentItemRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler list of assignments for the user. Nullable.
func (m *EducationAssignmentItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *EducationAssignmentItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationAssignmentFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable), nil
}
// Patch update the navigation property assignments in education
func (m *EducationAssignmentItemRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property assignments in education
func (m *EducationAssignmentItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, requestConfiguration *EducationAssignmentItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Publish the publish property
func (m *EducationAssignmentItemRequestBuilder) Publish()(*i74b91723ec2102e06ba61f2eec38b8a338e711092e99e7a6121ae3287f6f5788.PublishRequestBuilder) {
    return i74b91723ec2102e06ba61f2eec38b8a338e711092e99e7a6121ae3287f6f5788.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Resources the resources property
func (m *EducationAssignmentItemRequestBuilder) Resources()(*i874a2b579f5eb892e97d9c0c157220e5118544d0dca2de101640f4b07a1c3fe3.ResourcesRequestBuilder) {
    return i874a2b579f5eb892e97d9c0c157220e5118544d0dca2de101640f4b07a1c3fe3.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item.resources.item collection
func (m *EducationAssignmentItemRequestBuilder) ResourcesById(id string)(*i7b969696e68843bc2d2d30122cafa6f7eaf8a9843bf04feaf8693f3f794c0d94.EducationAssignmentResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource%2Did"] = id
    }
    return i7b969696e68843bc2d2d30122cafa6f7eaf8a9843bf04feaf8693f3f794c0d94.NewEducationAssignmentResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Rubric the rubric property
func (m *EducationAssignmentItemRequestBuilder) Rubric()(*i6182b98a227da91892039c997d1742545d169f3ded56acfb9ebddc2b8754acb3.RubricRequestBuilder) {
    return i6182b98a227da91892039c997d1742545d169f3ded56acfb9ebddc2b8754acb3.NewRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUpResourcesFolder the setUpResourcesFolder property
func (m *EducationAssignmentItemRequestBuilder) SetUpResourcesFolder()(*ie33a7d66ff2752d51b9a63d7758fc6c00305b565d547fa3e4e42f07c7a26bdec.SetUpResourcesFolderRequestBuilder) {
    return ie33a7d66ff2752d51b9a63d7758fc6c00305b565d547fa3e4e42f07c7a26bdec.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Submissions the submissions property
func (m *EducationAssignmentItemRequestBuilder) Submissions()(*i710174f1350381dcd776625bf1b593be962f1d2c9b1e368ae0bd178d3495728f.SubmissionsRequestBuilder) {
    return i710174f1350381dcd776625bf1b593be962f1d2c9b1e368ae0bd178d3495728f.NewSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item.submissions.item collection
func (m *EducationAssignmentItemRequestBuilder) SubmissionsById(id string)(*i0a97b1183e4efc9029dbfdb7442c76c2d4fc13a243fc9089be3106c63fd1b7f2.EducationSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission%2Did"] = id
    }
    return i0a97b1183e4efc9029dbfdb7442c76c2d4fc13a243fc9089be3106c63fd1b7f2.NewEducationSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
