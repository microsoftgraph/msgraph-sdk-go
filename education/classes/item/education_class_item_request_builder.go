package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i31b3cf6cbbc7ebca2bad02cd9933ecb2bf3ccbf186efcbf5a9bac253b27a73f6 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/schools"
    i36a91c3cde131d885c49a7e8a965b98d9d4f243edffd1923ccbe125acab7c772 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignmentcategories"
    i42fbe0832799af6aa7a5885188a6c903dc87216edac7b8a194c043607700ec89 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments"
    i63a93640e2827ef945d2e110c4b1ac757012187dc92f53a4ff00468359ca90f5 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/members"
    i67f45afa87d43202cad6adf1fc1ae035e41256f0f8e7dee433a1541e905db3ab "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignmentsettings"
    ic4dc10e52534928362dffdbe18eac0369c86942de53b8d20f7fbf80483a50cd2 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignmentdefaults"
    ic516469149971cdf008fcb769390c2ccbe41fd313cb712c9badce208aa5d4d41 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/group"
    ifc98e6bb2c8fdd13d4f834e2e5e8ade09edf25227ad3c95c215307ff11544424 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/teachers"
    i30f20691b867314f7df4463e5a84cdcabc27d3c5ee04690a1c94de7c4661e568 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/teachers/item"
    i82d47a387f18ab75fa126f3d2f3697a6ff95af00e66e3604a99d5edbdd0dd896 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item"
    ie06837c5d8d48e6cef6cb6b1bdc8fbe725c1239ff35cadfb51756b614376e50f "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/schools/item"
    if858bdb47736adc9d1c8aed87379aeb8381ad64cf3e897cb2b7ea10ddc52a603 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignmentcategories/item"
    if939da9a656c288d5389707e55630431aca3554eaf235d0e14f5f7fad55c64a8 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/members/item"
)

// EducationClassItemRequestBuilder provides operations to manage the classes property of the microsoft.graph.educationRoot entity.
type EducationClassItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// EducationClassItemRequestBuilderDeleteOptions options for Delete
type EducationClassItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EducationClassItemRequestBuilderGetOptions options for Get
type EducationClassItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *EducationClassItemRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// EducationClassItemRequestBuilderGetQueryParameters get classes from education
type EducationClassItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// EducationClassItemRequestBuilderPatchOptions options for Patch
type EducationClassItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationClassable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// AssignmentCategories the assignmentCategories property
func (m *EducationClassItemRequestBuilder) AssignmentCategories()(*i36a91c3cde131d885c49a7e8a965b98d9d4f243edffd1923ccbe125acab7c772.AssignmentCategoriesRequestBuilder) {
    return i36a91c3cde131d885c49a7e8a965b98d9d4f243edffd1923ccbe125acab7c772.NewAssignmentCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentCategoriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignmentCategories.item collection
func (m *EducationClassItemRequestBuilder) AssignmentCategoriesById(id string)(*if858bdb47736adc9d1c8aed87379aeb8381ad64cf3e897cb2b7ea10ddc52a603.EducationCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory%2Did"] = id
    }
    return if858bdb47736adc9d1c8aed87379aeb8381ad64cf3e897cb2b7ea10ddc52a603.NewEducationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignmentDefaults the assignmentDefaults property
func (m *EducationClassItemRequestBuilder) AssignmentDefaults()(*ic4dc10e52534928362dffdbe18eac0369c86942de53b8d20f7fbf80483a50cd2.AssignmentDefaultsRequestBuilder) {
    return ic4dc10e52534928362dffdbe18eac0369c86942de53b8d20f7fbf80483a50cd2.NewAssignmentDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Assignments the assignments property
func (m *EducationClassItemRequestBuilder) Assignments()(*i42fbe0832799af6aa7a5885188a6c903dc87216edac7b8a194c043607700ec89.AssignmentsRequestBuilder) {
    return i42fbe0832799af6aa7a5885188a6c903dc87216edac7b8a194c043607700ec89.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignments.item collection
func (m *EducationClassItemRequestBuilder) AssignmentsById(id string)(*i82d47a387f18ab75fa126f3d2f3697a6ff95af00e66e3604a99d5edbdd0dd896.EducationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignment%2Did"] = id
    }
    return i82d47a387f18ab75fa126f3d2f3697a6ff95af00e66e3604a99d5edbdd0dd896.NewEducationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// AssignmentSettings the assignmentSettings property
func (m *EducationClassItemRequestBuilder) AssignmentSettings()(*i67f45afa87d43202cad6adf1fc1ae035e41256f0f8e7dee433a1541e905db3ab.AssignmentSettingsRequestBuilder) {
    return i67f45afa87d43202cad6adf1fc1ae035e41256f0f8e7dee433a1541e905db3ab.NewAssignmentSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEducationClassItemRequestBuilderInternal instantiates a new EducationClassItemRequestBuilder and sets the default values.
func NewEducationClassItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationClassItemRequestBuilder) {
    m := &EducationClassItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationClassItemRequestBuilder instantiates a new EducationClassItemRequestBuilder and sets the default values.
func NewEducationClassItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*EducationClassItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationClassItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property classes for education
func (m *EducationClassItemRequestBuilder) CreateDeleteRequestInformation(options *EducationClassItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreateGetRequestInformation get classes from education
func (m *EducationClassItemRequestBuilder) CreateGetRequestInformation(options *EducationClassItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property classes in education
func (m *EducationClassItemRequestBuilder) CreatePatchRequestInformation(options *EducationClassItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property classes for education
func (m *EducationClassItemRequestBuilder) Delete(options *EducationClassItemRequestBuilderDeleteOptions)(error) {
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
// Get get classes from education
func (m *EducationClassItemRequestBuilder) Get(options *EducationClassItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationClassable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateEducationClassFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.EducationClassable), nil
}
// Group the group property
func (m *EducationClassItemRequestBuilder) Group()(*ic516469149971cdf008fcb769390c2ccbe41fd313cb712c9badce208aa5d4d41.GroupRequestBuilder) {
    return ic516469149971cdf008fcb769390c2ccbe41fd313cb712c9badce208aa5d4d41.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Members the members property
func (m *EducationClassItemRequestBuilder) Members()(*i63a93640e2827ef945d2e110c4b1ac757012187dc92f53a4ff00468359ca90f5.MembersRequestBuilder) {
    return i63a93640e2827ef945d2e110c4b1ac757012187dc92f53a4ff00468359ca90f5.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MembersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.members.item collection
func (m *EducationClassItemRequestBuilder) MembersById(id string)(*if939da9a656c288d5389707e55630431aca3554eaf235d0e14f5f7fad55c64a8.EducationUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationUser%2Did"] = id
    }
    return if939da9a656c288d5389707e55630431aca3554eaf235d0e14f5f7fad55c64a8.NewEducationUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property classes in education
func (m *EducationClassItemRequestBuilder) Patch(options *EducationClassItemRequestBuilderPatchOptions)(error) {
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
// Schools the schools property
func (m *EducationClassItemRequestBuilder) Schools()(*i31b3cf6cbbc7ebca2bad02cd9933ecb2bf3ccbf186efcbf5a9bac253b27a73f6.SchoolsRequestBuilder) {
    return i31b3cf6cbbc7ebca2bad02cd9933ecb2bf3ccbf186efcbf5a9bac253b27a73f6.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchoolsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.schools.item collection
func (m *EducationClassItemRequestBuilder) SchoolsById(id string)(*ie06837c5d8d48e6cef6cb6b1bdc8fbe725c1239ff35cadfb51756b614376e50f.EducationSchoolItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSchool%2Did"] = id
    }
    return ie06837c5d8d48e6cef6cb6b1bdc8fbe725c1239ff35cadfb51756b614376e50f.NewEducationSchoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Teachers the teachers property
func (m *EducationClassItemRequestBuilder) Teachers()(*ifc98e6bb2c8fdd13d4f834e2e5e8ade09edf25227ad3c95c215307ff11544424.TeachersRequestBuilder) {
    return ifc98e6bb2c8fdd13d4f834e2e5e8ade09edf25227ad3c95c215307ff11544424.NewTeachersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TeachersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.teachers.item collection
func (m *EducationClassItemRequestBuilder) TeachersById(id string)(*i30f20691b867314f7df4463e5a84cdcabc27d3c5ee04690a1c94de7c4661e568.EducationUserItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationUser%2Did"] = id
    }
    return i30f20691b867314f7df4463e5a84cdcabc27d3c5ee04690a1c94de7c4661e568.NewEducationUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
