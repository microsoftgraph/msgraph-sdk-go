package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i31b3cf6cbbc7ebca2bad02cd9933ecb2bf3ccbf186efcbf5a9bac253b27a73f6 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/schools"
    i36a91c3cde131d885c49a7e8a965b98d9d4f243edffd1923ccbe125acab7c772 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignmentcategories"
    i42fbe0832799af6aa7a5885188a6c903dc87216edac7b8a194c043607700ec89 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments"
    i63a93640e2827ef945d2e110c4b1ac757012187dc92f53a4ff00468359ca90f5 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/members"
    i67f45afa87d43202cad6adf1fc1ae035e41256f0f8e7dee433a1541e905db3ab "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignmentsettings"
    ic4dc10e52534928362dffdbe18eac0369c86942de53b8d20f7fbf80483a50cd2 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignmentdefaults"
    ic516469149971cdf008fcb769390c2ccbe41fd313cb712c9badce208aa5d4d41 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/group"
    ifc98e6bb2c8fdd13d4f834e2e5e8ade09edf25227ad3c95c215307ff11544424 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/teachers"
    i82d47a387f18ab75fa126f3d2f3697a6ff95af00e66e3604a99d5edbdd0dd896 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item"
    if858bdb47736adc9d1c8aed87379aeb8381ad64cf3e897cb2b7ea10ddc52a603 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignmentcategories/item"
)

// educationClassRequestBuilder builds and executes requests for operations under \education\classes\{educationClass-id}
type EducationClassRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationClassRequestBuilderDeleteOptions options for Delete
type EducationClassRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationClassRequestBuilderGetOptions options for Get
type EducationClassRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationClassRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// educationClassRequestBuilderGetQueryParameters get classes from education
type EducationClassRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// EducationClassRequestBuilderPatchOptions options for Patch
type EducationClassRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationClass;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EducationClassRequestBuilder) AssignmentCategories()(*i36a91c3cde131d885c49a7e8a965b98d9d4f243edffd1923ccbe125acab7c772.AssignmentCategoriesRequestBuilder) {
    return i36a91c3cde131d885c49a7e8a965b98d9d4f243edffd1923ccbe125acab7c772.NewAssignmentCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentCategoriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignmentCategories.item collection
func (m *EducationClassRequestBuilder) AssignmentCategoriesById(id string)(*if858bdb47736adc9d1c8aed87379aeb8381ad64cf3e897cb2b7ea10ddc52a603.EducationCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory_id"] = id
    }
    return if858bdb47736adc9d1c8aed87379aeb8381ad64cf3e897cb2b7ea10ddc52a603.NewEducationCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationClassRequestBuilder) AssignmentDefaults()(*ic4dc10e52534928362dffdbe18eac0369c86942de53b8d20f7fbf80483a50cd2.AssignmentDefaultsRequestBuilder) {
    return ic4dc10e52534928362dffdbe18eac0369c86942de53b8d20f7fbf80483a50cd2.NewAssignmentDefaultsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationClassRequestBuilder) Assignments()(*i42fbe0832799af6aa7a5885188a6c903dc87216edac7b8a194c043607700ec89.AssignmentsRequestBuilder) {
    return i42fbe0832799af6aa7a5885188a6c903dc87216edac7b8a194c043607700ec89.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignments.item collection
func (m *EducationClassRequestBuilder) AssignmentsById(id string)(*i82d47a387f18ab75fa126f3d2f3697a6ff95af00e66e3604a99d5edbdd0dd896.EducationAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignment_id"] = id
    }
    return i82d47a387f18ab75fa126f3d2f3697a6ff95af00e66e3604a99d5edbdd0dd896.NewEducationAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationClassRequestBuilder) AssignmentSettings()(*i67f45afa87d43202cad6adf1fc1ae035e41256f0f8e7dee433a1541e905db3ab.AssignmentSettingsRequestBuilder) {
    return i67f45afa87d43202cad6adf1fc1ae035e41256f0f8e7dee433a1541e905db3ab.NewAssignmentSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewEducationClassRequestBuilderInternal instantiates a new EducationClassRequestBuilder and sets the default values.
func NewEducationClassRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationClassRequestBuilder) {
    m := &EducationClassRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/classes/{educationClass_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationClassRequestBuilder instantiates a new EducationClassRequestBuilder and sets the default values.
func NewEducationClassRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationClassRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationClassRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property classes for education
func (m *EducationClassRequestBuilder) CreateDeleteRequestInformation(options *EducationClassRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationClassRequestBuilder) CreateGetRequestInformation(options *EducationClassRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationClassRequestBuilder) CreatePatchRequestInformation(options *EducationClassRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *EducationClassRequestBuilder) Delete(options *EducationClassRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// Get get classes from education
func (m *EducationClassRequestBuilder) Get(options *EducationClassRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationClass, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEducationClass() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationClass), nil
}
func (m *EducationClassRequestBuilder) Group()(*ic516469149971cdf008fcb769390c2ccbe41fd313cb712c9badce208aa5d4d41.GroupRequestBuilder) {
    return ic516469149971cdf008fcb769390c2ccbe41fd313cb712c9badce208aa5d4d41.NewGroupRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationClassRequestBuilder) Members()(*i63a93640e2827ef945d2e110c4b1ac757012187dc92f53a4ff00468359ca90f5.MembersRequestBuilder) {
    return i63a93640e2827ef945d2e110c4b1ac757012187dc92f53a4ff00468359ca90f5.NewMembersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property classes in education
func (m *EducationClassRequestBuilder) Patch(options *EducationClassRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *EducationClassRequestBuilder) Schools()(*i31b3cf6cbbc7ebca2bad02cd9933ecb2bf3ccbf186efcbf5a9bac253b27a73f6.SchoolsRequestBuilder) {
    return i31b3cf6cbbc7ebca2bad02cd9933ecb2bf3ccbf186efcbf5a9bac253b27a73f6.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationClassRequestBuilder) Teachers()(*ifc98e6bb2c8fdd13d4f834e2e5e8ade09edf25227ad3c95c215307ff11544424.TeachersRequestBuilder) {
    return ifc98e6bb2c8fdd13d4f834e2e5e8ade09edf25227ad3c95c215307ff11544424.NewTeachersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
