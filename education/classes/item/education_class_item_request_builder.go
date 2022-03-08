package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
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
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationClassItemRequestBuilderDeleteOptions options for Delete
type EducationClassItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationClassItemRequestBuilderGetOptions options for Get
type EducationClassItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationClassItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationClassItemRequestBuilderGetQueryParameters get classes from education
type EducationClassItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationClassItemRequestBuilderPatchOptions options for Patch
type EducationClassItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationClassable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
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
        urlTplParams["educationCategory_id"] = id
    }
    return if858bdb47736adc9d1c8aed87379aeb8381ad64cf3e897cb2b7ea10ddc52a603.NewEducationCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationClassItemRequestBuilder) AssignmentDefaults()(*ic4dc10e52534928362dffdbe18eac0369c86942de53b8d20f7fbf80483a50cd2.AssignmentDefaultsRequestBuilder) {
    return ic4dc10e52534928362dffdbe18eac0369c86942de53b8d20f7fbf80483a50cd2.NewAssignmentDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["educationAssignment_id"] = id
    }
    return i82d47a387f18ab75fa126f3d2f3697a6ff95af00e66e3604a99d5edbdd0dd896.NewEducationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationClassItemRequestBuilder) AssignmentSettings()(*i67f45afa87d43202cad6adf1fc1ae035e41256f0f8e7dee433a1541e905db3ab.AssignmentSettingsRequestBuilder) {
    return i67f45afa87d43202cad6adf1fc1ae035e41256f0f8e7dee433a1541e905db3ab.NewAssignmentSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEducationClassItemRequestBuilderInternal instantiates a new EducationClassItemRequestBuilder and sets the default values.
func NewEducationClassItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationClassItemRequestBuilder) {
    m := &EducationClassItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationClassItemRequestBuilder instantiates a new EducationClassItemRequestBuilder and sets the default values.
func NewEducationClassItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationClassItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationClassItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property classes for education
func (m *EducationClassItemRequestBuilder) CreateDeleteRequestInformation(options *EducationClassItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get classes from education
func (m *EducationClassItemRequestBuilder) CreateGetRequestInformation(options *EducationClassItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property classes in education
func (m *EducationClassItemRequestBuilder) CreatePatchRequestInformation(options *EducationClassItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
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
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get classes from education
func (m *EducationClassItemRequestBuilder) Get(options *EducationClassItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationClassable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateEducationClassFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationClassable), nil
}
func (m *EducationClassItemRequestBuilder) Group()(*ic516469149971cdf008fcb769390c2ccbe41fd313cb712c9badce208aa5d4d41.GroupRequestBuilder) {
    return ic516469149971cdf008fcb769390c2ccbe41fd313cb712c9badce208aa5d4d41.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
        urlTplParams["educationUser_id"] = id
    }
    return if939da9a656c288d5389707e55630431aca3554eaf235d0e14f5f7fad55c64a8.NewEducationUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update the navigation property classes in education
func (m *EducationClassItemRequestBuilder) Patch(options *EducationClassItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
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
        urlTplParams["educationSchool_id"] = id
    }
    return ie06837c5d8d48e6cef6cb6b1bdc8fbe725c1239ff35cadfb51756b614376e50f.NewEducationSchoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
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
        urlTplParams["educationUser_id"] = id
    }
    return i30f20691b867314f7df4463e5a84cdcabc27d3c5ee04690a1c94de7c4661e568.NewEducationUserItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
