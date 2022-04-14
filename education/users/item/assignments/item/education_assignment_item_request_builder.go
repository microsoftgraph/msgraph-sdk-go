package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i4a60e659b3e1427ff2f92b71f44e26c8ac1ac819ee646533021117deeeb6b88a "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/rubric"
    i7b1c280606a6564228514ab9e1b7f963f46c753907436ec516230878d54adc8f "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/categories"
    i88086f8017be5762df8cc58a527cb8b9ece6f57e0c9ce6c6a7b1682cb8a238cc "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/resources"
    ib979357b11252b944c152b3aa6f90a6415f1f60345dd67d0804b319a1f2123ba "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions"
    ic8a76de5ffb5d6e2de378abe7f4341d25a11d4ea88cad7b35c31b72f63d17e7a "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/setupresourcesfolder"
    ied86aa882920629ba539d3c8055e54c93f45dfa9b4f250108d4b4d3bbc80975a "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/publish"
    i2d4596e7585085aa1bcbd1dfc3b2ce4826b6afe4b02ed3c80e40d2c1de2dd769 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/submissions/item"
    i37bd40aa3d07e5fb03511b981c264fa6dcad03889d85b2483191049178483b0d "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/resources/item"
    i5914468bbfd7cf156e785aa5b7c676dd4a5c8f64f204e13bca5907a261b89c6c "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item/categories/item"
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
// EducationAssignmentItemRequestBuilderDeleteOptions options for Delete
type EducationAssignmentItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EducationAssignmentItemRequestBuilderGetOptions options for Get
type EducationAssignmentItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EducationAssignmentItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EducationAssignmentItemRequestBuilderGetQueryParameters assignments belonging to the user.
type EducationAssignmentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EducationAssignmentItemRequestBuilderPatchOptions options for Patch
type EducationAssignmentItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Categories the categories property
func (m *EducationAssignmentItemRequestBuilder) Categories()(*i7b1c280606a6564228514ab9e1b7f963f46c753907436ec516230878d54adc8f.CategoriesRequestBuilder) {
    return i7b1c280606a6564228514ab9e1b7f963f46c753907436ec516230878d54adc8f.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CategoriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item.categories.item collection
func (m *EducationAssignmentItemRequestBuilder) CategoriesById(id string)(*i5914468bbfd7cf156e785aa5b7c676dd4a5c8f64f204e13bca5907a261b89c6c.EducationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory%2Did"] = id
    }
    return i5914468bbfd7cf156e785aa5b7c676dd4a5c8f64f204e13bca5907a261b89c6c.NewEducationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationAssignmentItemRequestBuilderInternal instantiates a new EducationAssignmentItemRequestBuilder and sets the default values.
func NewEducationAssignmentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationAssignmentItemRequestBuilder) {
    m := &EducationAssignmentItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser%2Did}/assignments/{educationAssignment%2Did}{?%24select,%24expand}";
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
func (m *EducationAssignmentItemRequestBuilder) CreateDeleteRequestInformation(options *EducationAssignmentItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation assignments belonging to the user.
func (m *EducationAssignmentItemRequestBuilder) CreateGetRequestInformation(options *EducationAssignmentItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property assignments in education
func (m *EducationAssignmentItemRequestBuilder) CreatePatchRequestInformation(options *EducationAssignmentItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property assignments for education
func (m *EducationAssignmentItemRequestBuilder) Delete(options *EducationAssignmentItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get assignments belonging to the user.
func (m *EducationAssignmentItemRequestBuilder) Get(options *EducationAssignmentItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationAssignmentFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationAssignmentable), nil
}
// Patch update the navigation property assignments in education
func (m *EducationAssignmentItemRequestBuilder) Patch(options *EducationAssignmentItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Publish the publish property
func (m *EducationAssignmentItemRequestBuilder) Publish()(*ied86aa882920629ba539d3c8055e54c93f45dfa9b4f250108d4b4d3bbc80975a.PublishRequestBuilder) {
    return ied86aa882920629ba539d3c8055e54c93f45dfa9b4f250108d4b4d3bbc80975a.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Resources the resources property
func (m *EducationAssignmentItemRequestBuilder) Resources()(*i88086f8017be5762df8cc58a527cb8b9ece6f57e0c9ce6c6a7b1682cb8a238cc.ResourcesRequestBuilder) {
    return i88086f8017be5762df8cc58a527cb8b9ece6f57e0c9ce6c6a7b1682cb8a238cc.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourcesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item.resources.item collection
func (m *EducationAssignmentItemRequestBuilder) ResourcesById(id string)(*i37bd40aa3d07e5fb03511b981c264fa6dcad03889d85b2483191049178483b0d.EducationAssignmentResourceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource%2Did"] = id
    }
    return i37bd40aa3d07e5fb03511b981c264fa6dcad03889d85b2483191049178483b0d.NewEducationAssignmentResourceItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Rubric the rubric property
func (m *EducationAssignmentItemRequestBuilder) Rubric()(*i4a60e659b3e1427ff2f92b71f44e26c8ac1ac819ee646533021117deeeb6b88a.RubricRequestBuilder) {
    return i4a60e659b3e1427ff2f92b71f44e26c8ac1ac819ee646533021117deeeb6b88a.NewRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SetUpResourcesFolder the setUpResourcesFolder property
func (m *EducationAssignmentItemRequestBuilder) SetUpResourcesFolder()(*ic8a76de5ffb5d6e2de378abe7f4341d25a11d4ea88cad7b35c31b72f63d17e7a.SetUpResourcesFolderRequestBuilder) {
    return ic8a76de5ffb5d6e2de378abe7f4341d25a11d4ea88cad7b35c31b72f63d17e7a.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Submissions the submissions property
func (m *EducationAssignmentItemRequestBuilder) Submissions()(*ib979357b11252b944c152b3aa6f90a6415f1f60345dd67d0804b319a1f2123ba.SubmissionsRequestBuilder) {
    return ib979357b11252b944c152b3aa6f90a6415f1f60345dd67d0804b319a1f2123ba.NewSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubmissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item.submissions.item collection
func (m *EducationAssignmentItemRequestBuilder) SubmissionsById(id string)(*i2d4596e7585085aa1bcbd1dfc3b2ce4826b6afe4b02ed3c80e40d2c1de2dd769.EducationSubmissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission%2Did"] = id
    }
    return i2d4596e7585085aa1bcbd1dfc3b2ce4826b6afe4b02ed3c80e40d2c1de2dd769.NewEducationSubmissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
