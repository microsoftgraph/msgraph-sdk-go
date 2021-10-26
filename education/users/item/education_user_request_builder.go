package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i221d3f1c8592f7a9636476245f2e4a9e82296b6e5cc50822548a6744c13ec682 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/user"
    ibaab995f04768ff7b941456360e066c6fd775e37b26f2c71367a3f6264b2a4c3 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/classes"
    icc9d55c2563db563fec8aa8615c0e6e26a0a9a03820b47d4b01f7ef2f53ee2e6 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/schools"
    if19c2a0949f5f27bde752502b1831a530b8da87a55d1240bdad90f98dc585e1e "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/rubrics"
    ifc2ed629d6ca53bbfc33acc06329c54a103a8e43f74a0e018639ae7bafc6020b "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/taughtclasses"
    ia6a8d66cb9ce8a7587ed39000e658b45c988146e3baee26d141c9c4829ae1818 "github.com/microsoftgraph/msgraph-sdk-go/education/users/item/rubrics/item"
)

// Builds and executes requests for operations under \education\users\{educationUser-id}
type EducationUserRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Get users from education
type EducationUserRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *EducationUserRequestBuilder) Classes()(*ibaab995f04768ff7b941456360e066c6fd775e37b26f2c71367a3f6264b2a4c3.ClassesRequestBuilder) {
    return ibaab995f04768ff7b941456360e066c6fd775e37b26f2c71367a3f6264b2a4c3.NewClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new EducationUserRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEducationUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationUserRequestBuilder) {
    m := &EducationUserRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/education/users/{educationUser_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new EducationUserRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEducationUserRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete navigation property users for education
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *EducationUserRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get users from education
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *EducationUserRequestBuilder) CreateGetRequestInformation(q func (value *EducationUserRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EducationUserRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Update the navigation property users in education
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *EducationUserRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationUser, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete navigation property users for education
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *EducationUserRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
// Get users from education
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *EducationUserRequestBuilder) Get(q func (value *EducationUserRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationUser, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEducationUser() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationUser), nil
}
// Update the navigation property users in education
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *EducationUserRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationUser, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *EducationUserRequestBuilder) Rubrics()(*if19c2a0949f5f27bde752502b1831a530b8da87a55d1240bdad90f98dc585e1e.RubricsRequestBuilder) {
    return if19c2a0949f5f27bde752502b1831a530b8da87a55d1240bdad90f98dc585e1e.NewRubricsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.users.item.rubrics.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EducationUserRequestBuilder) RubricsById(id string)(*ia6a8d66cb9ce8a7587ed39000e658b45c988146e3baee26d141c9c4829ae1818.EducationRubricRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationRubric_id"] = id
    }
    return ia6a8d66cb9ce8a7587ed39000e658b45c988146e3baee26d141c9c4829ae1818.NewEducationRubricRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationUserRequestBuilder) Schools()(*icc9d55c2563db563fec8aa8615c0e6e26a0a9a03820b47d4b01f7ef2f53ee2e6.SchoolsRequestBuilder) {
    return icc9d55c2563db563fec8aa8615c0e6e26a0a9a03820b47d4b01f7ef2f53ee2e6.NewSchoolsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationUserRequestBuilder) TaughtClasses()(*ifc2ed629d6ca53bbfc33acc06329c54a103a8e43f74a0e018639ae7bafc6020b.TaughtClassesRequestBuilder) {
    return ifc2ed629d6ca53bbfc33acc06329c54a103a8e43f74a0e018639ae7bafc6020b.NewTaughtClassesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationUserRequestBuilder) User()(*i221d3f1c8592f7a9636476245f2e4a9e82296b6e5cc50822548a6744c13ec682.UserRequestBuilder) {
    return i221d3f1c8592f7a9636476245f2e4a9e82296b6e5cc50822548a6744c13ec682.NewUserRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
