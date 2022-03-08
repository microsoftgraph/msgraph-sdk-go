package me

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4827ba32335285b89a92c436cf648bf167912072becfb2cf9792b7e694bfe68c "github.com/microsoftgraph/msgraph-sdk-go/education/me/taughtclasses"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i75d26f07b5741abbb041934b7ff07c14802e8373402b2fe0175141ce976d99f1 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments"
    i8a6dd866ad5ea518840cdc5c8b7348b010903fad7cccc3473ccbc41bd0f2f5c5 "github.com/microsoftgraph/msgraph-sdk-go/education/me/rubrics"
    idc612964c10975ba0294061975f8fd346f82666fc4c4a4245e1b4ea27187f3e0 "github.com/microsoftgraph/msgraph-sdk-go/education/me/classes"
    if174dba299b5249c431f058ba35f021f7837f7d6553d65eaa9c8f39e0c8d2f25 "github.com/microsoftgraph/msgraph-sdk-go/education/me/user"
    if8f2e52b9fa261cdafc93406e1f0deaaa9c3444674bb9bac26886d2161875b88 "github.com/microsoftgraph/msgraph-sdk-go/education/me/schools"
    i05b947f4d440080ac6b413aae5547627313333329bb760c803e6c27f60db5c06 "github.com/microsoftgraph/msgraph-sdk-go/education/me/assignments/item"
    i434c17782aff494ea747f7c3919ee493165cf0cebf127d5797e6c3c1b044f21e "github.com/microsoftgraph/msgraph-sdk-go/education/me/schools/item"
    i4a61e633c19f80889fc1151239c34c623c4ada576c2613365f8634b2142f493d "github.com/microsoftgraph/msgraph-sdk-go/education/me/taughtclasses/item"
    ib528428124f1c371970dddb56f26c3bb4c42e706f192473c9d8d84f45c312a67 "github.com/microsoftgraph/msgraph-sdk-go/education/me/classes/item"
    ic96595483f91910c603469ce138a672deaf059c7a8c295304a9c2acf2ce6134b "github.com/microsoftgraph/msgraph-sdk-go/education/me/rubrics/item"
)

// MeRequestBuilder provides operations to manage the me property of the microsoft.graph.educationRoot entity.
type MeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// MeRequestBuilderDeleteOptions options for Delete
type MeRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MeRequestBuilderGetOptions options for Get
type MeRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *MeRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// MeRequestBuilderGetQueryParameters get me from education
type MeRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// MeRequestBuilderPatchOptions options for Patch
type MeRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationUserable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *MeRequestBuilder) Assignments()(*i75d26f07b5741abbb041934b7ff07c14802e8373402b2fe0175141ce976d99f1.AssignmentsRequestBuilder) {
    return i75d26f07b5741abbb041934b7ff07c14802e8373402b2fe0175141ce976d99f1.NewAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.assignments.item collection
func (m *MeRequestBuilder) AssignmentsById(id string)(*i05b947f4d440080ac6b413aae5547627313333329bb760c803e6c27f60db5c06.EducationAssignmentItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignment_id"] = id
    }
    return i05b947f4d440080ac6b413aae5547627313333329bb760c803e6c27f60db5c06.NewEducationAssignmentItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) Classes()(*idc612964c10975ba0294061975f8fd346f82666fc4c4a4245e1b4ea27187f3e0.ClassesRequestBuilder) {
    return idc612964c10975ba0294061975f8fd346f82666fc4c4a4245e1b4ea27187f3e0.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ClassesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.classes.item collection
func (m *MeRequestBuilder) ClassesById(id string)(*ib528428124f1c371970dddb56f26c3bb4c42e706f192473c9d8d84f45c312a67.EducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass_id"] = id
    }
    return ib528428124f1c371970dddb56f26c3bb4c42e706f192473c9d8d84f45c312a67.NewEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewMeRequestBuilderInternal instantiates a new MeRequestBuilder and sets the default values.
func NewMeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MeRequestBuilder) {
    m := &MeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/education/me{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewMeRequestBuilder instantiates a new MeRequestBuilder and sets the default values.
func NewMeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*MeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewMeRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property me for education
func (m *MeRequestBuilder) CreateDeleteRequestInformation(options *MeRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get me from education
func (m *MeRequestBuilder) CreateGetRequestInformation(options *MeRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property me in education
func (m *MeRequestBuilder) CreatePatchRequestInformation(options *MeRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property me for education
func (m *MeRequestBuilder) Delete(options *MeRequestBuilderDeleteOptions)(error) {
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
// Get get me from education
func (m *MeRequestBuilder) Get(options *MeRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationUserable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateEducationUserFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationUserable), nil
}
// Patch update the navigation property me in education
func (m *MeRequestBuilder) Patch(options *MeRequestBuilderPatchOptions)(error) {
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
func (m *MeRequestBuilder) Rubrics()(*i8a6dd866ad5ea518840cdc5c8b7348b010903fad7cccc3473ccbc41bd0f2f5c5.RubricsRequestBuilder) {
    return i8a6dd866ad5ea518840cdc5c8b7348b010903fad7cccc3473ccbc41bd0f2f5c5.NewRubricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RubricsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.rubrics.item collection
func (m *MeRequestBuilder) RubricsById(id string)(*ic96595483f91910c603469ce138a672deaf059c7a8c295304a9c2acf2ce6134b.EducationRubricItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationRubric_id"] = id
    }
    return ic96595483f91910c603469ce138a672deaf059c7a8c295304a9c2acf2ce6134b.NewEducationRubricItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) Schools()(*if8f2e52b9fa261cdafc93406e1f0deaaa9c3444674bb9bac26886d2161875b88.SchoolsRequestBuilder) {
    return if8f2e52b9fa261cdafc93406e1f0deaaa9c3444674bb9bac26886d2161875b88.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SchoolsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.schools.item collection
func (m *MeRequestBuilder) SchoolsById(id string)(*i434c17782aff494ea747f7c3919ee493165cf0cebf127d5797e6c3c1b044f21e.EducationSchoolItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSchool_id"] = id
    }
    return i434c17782aff494ea747f7c3919ee493165cf0cebf127d5797e6c3c1b044f21e.NewEducationSchoolItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) TaughtClasses()(*i4827ba32335285b89a92c436cf648bf167912072becfb2cf9792b7e694bfe68c.TaughtClassesRequestBuilder) {
    return i4827ba32335285b89a92c436cf648bf167912072becfb2cf9792b7e694bfe68c.NewTaughtClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TaughtClassesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.me.taughtClasses.item collection
func (m *MeRequestBuilder) TaughtClassesById(id string)(*i4a61e633c19f80889fc1151239c34c623c4ada576c2613365f8634b2142f493d.EducationClassItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationClass_id"] = id
    }
    return i4a61e633c19f80889fc1151239c34c623c4ada576c2613365f8634b2142f493d.NewEducationClassItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *MeRequestBuilder) User()(*if174dba299b5249c431f058ba35f021f7837f7d6553d65eaa9c8f39e0c8d2f25.UserRequestBuilder) {
    return if174dba299b5249c431f058ba35f021f7837f7d6553d65eaa9c8f39e0c8d2f25.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
