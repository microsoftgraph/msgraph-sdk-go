package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i221d3f1c8592f7a9636476245f2e4a9e82296b6e5cc50822548a6744c13ec682 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/user"
    i3ff603f1915775aa8fd74380549e8fdf33447260d93eee635198c7c3ce55ef41 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments"
    ibaab995f04768ff7b941456360e066c6fd775e37b26f2c71367a3f6264b2a4c3 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/classes"
    icc9d55c2563db563fec8aa8615c0e6e26a0a9a03820b47d4b01f7ef2f53ee2e6 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/schools"
    if19c2a0949f5f27bde752502b1831a530b8da87a55d1240bdad90f98dc585e1e "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/rubrics"
    ifc2ed629d6ca53bbfc33acc06329c54a103a8e43f74a0e018639ae7bafc6020b "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/taughtclasses"
    i142cc91605bccdf7d1947e2356d3e675fabe42483abf58b911ef3417b99a7658 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/assignments/item"
    i6c43b628928247a4ca6ee2535b4cc6a110e39b8ff8155ce809de4582b7428450 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/taughtclasses/item"
    ia6a8d66cb9ce8a7587ed39000e658b45c988146e3baee26d141c9c4829ae1818 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/rubrics/item"
    ia9b5ce3f1a9b6628a35c3d194c8db39e4f5a678c68df1c7d38765c1e6983ee9f "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/schools/item"
    id83e01290b448c553c9d9d81fdf7d67ecfbcc1916e58827c310107bad3ba0ec7 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/classes/item"
)

// EducationUserItemRequestBuilder provides operations to manage the users property of the microsoft.graph.educationRoot entity.
type EducationUserItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// EducationUserItemRequestBuilderDeleteOptions options for Delete
type EducationUserItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationUserItemRequestBuilderGetOptions options for Get
type EducationUserItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *EducationUserItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// EducationUserItemRequestBuilderGetQueryParameters get users from education
type EducationUserItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// EducationUserItemRequestBuilderPatchOptions options for Patch
type EducationUserItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationUserable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *EducationUserItemRequestBuilder) Assignments()(*i3ff603f1915775aa8fd74380549e8fdf33447260d93eee635198c7c3ce55ef41.AssignmentsRequestBuilder) {
    return i3ff603f1915775aa8fd74380549e8fdf33447260d93eee635198c7c3ce55ef41.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.assignments.item collection
func (m *EducationUserItemRequestBuilder) AssignmentsById(id string)(*i142cc91605bccdf7d1947e2356d3e675fabe42483abf58b911ef3417b99a7658.EducationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignment_id"] = id
    }
    return i142cc91605bccdf7d1947e2356d3e675fabe42483abf58b911ef3417b99a7658.NewEducationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationUserItemRequestBuilder) Classes()(*ibaab995f04768ff7b941456360e066c6fd775e37b26f2c71367a3f6264b2a4c3.ClassesRequestBuilder) {
    return ibaab995f04768ff7b941456360e066c6fd775e37b26f2c71367a3f6264b2a4c3.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.classes.item collection
func (m *EducationUserItemRequestBuilder) ClassesById(id string)(*id83e01290b448c553c9d9d81fdf7d67ecfbcc1916e58827c310107bad3ba0ec7.EducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass_id"] = id
    }
    return id83e01290b448c553c9d9d81fdf7d67ecfbcc1916e58827c310107bad3ba0ec7.NewEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewEducationUserItemRequestBuilderInternal instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewEducationUserItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationUserItemRequestBuilder) {
    m := &EducationUserItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/users/{educationUser_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewEducationUserItemRequestBuilder instantiates a new EducationUserItemRequestBuilder and sets the default values.
func NewEducationUserItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationUserItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationUserItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property users for education
func (m *EducationUserItemRequestBuilder) CreateDeleteRequestInformation(options *EducationUserItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get users from education
func (m *EducationUserItemRequestBuilder) CreateGetRequestInformation(options *EducationUserItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property users in education
func (m *EducationUserItemRequestBuilder) CreatePatchRequestInformation(options *EducationUserItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property users for education
func (m *EducationUserItemRequestBuilder) Delete(options *EducationUserItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get get users from education
func (m *EducationUserItemRequestBuilder) Get(options *EducationUserItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationUserable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateEducationUserFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationUserable), nil
}
// Patch update the navigation property users in education
func (m *EducationUserItemRequestBuilder) Patch(options *EducationUserItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *EducationUserItemRequestBuilder) Rubrics()(*if19c2a0949f5f27bde752502b1831a530b8da87a55d1240bdad90f98dc585e1e.RubricsRequestBuilder) {
    return if19c2a0949f5f27bde752502b1831a530b8da87a55d1240bdad90f98dc585e1e.NewRubricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RubricsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.rubrics.item collection
func (m *EducationUserItemRequestBuilder) RubricsById(id string)(*ia6a8d66cb9ce8a7587ed39000e658b45c988146e3baee26d141c9c4829ae1818.EducationRubricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationRubric_id"] = id
    }
    return ia6a8d66cb9ce8a7587ed39000e658b45c988146e3baee26d141c9c4829ae1818.NewEducationRubricItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationUserItemRequestBuilder) Schools()(*icc9d55c2563db563fec8aa8615c0e6e26a0a9a03820b47d4b01f7ef2f53ee2e6.SchoolsRequestBuilder) {
    return icc9d55c2563db563fec8aa8615c0e6e26a0a9a03820b47d4b01f7ef2f53ee2e6.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchoolsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.schools.item collection
func (m *EducationUserItemRequestBuilder) SchoolsById(id string)(*ia9b5ce3f1a9b6628a35c3d194c8db39e4f5a678c68df1c7d38765c1e6983ee9f.EducationSchoolItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSchool_id"] = id
    }
    return ia9b5ce3f1a9b6628a35c3d194c8db39e4f5a678c68df1c7d38765c1e6983ee9f.NewEducationSchoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationUserItemRequestBuilder) TaughtClasses()(*ifc2ed629d6ca53bbfc33acc06329c54a103a8e43f74a0e018639ae7bafc6020b.TaughtClassesRequestBuilder) {
    return ifc2ed629d6ca53bbfc33acc06329c54a103a8e43f74a0e018639ae7bafc6020b.NewTaughtClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaughtClassesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.taughtClasses.item collection
func (m *EducationUserItemRequestBuilder) TaughtClassesById(id string)(*i6c43b628928247a4ca6ee2535b4cc6a110e39b8ff8155ce809de4582b7428450.EducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass_id"] = id
    }
    return i6c43b628928247a4ca6ee2535b4cc6a110e39b8ff8155ce809de4582b7428450.NewEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationUserItemRequestBuilder) User()(*i221d3f1c8592f7a9636476245f2e4a9e82296b6e5cc50822548a6744c13ec682.UserRequestBuilder) {
    return i221d3f1c8592f7a9636476245f2e4a9e82296b6e5cc50822548a6744c13ec682.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
