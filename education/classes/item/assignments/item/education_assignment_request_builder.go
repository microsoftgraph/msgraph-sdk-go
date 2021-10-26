package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i3d50a0849e57022c8564fae5dbbb2b471978d26a2db68975ed257f10e693e7cd "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions"
    i42850d7acdbd3a8e01158ce3f91031a284414d79a1b8423753fe9921702335f6 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/setupresourcesfolder"
    i560e0f51eb0b15ee54800eeae829fe0350e27c3694c083160844bef71444f4f0 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/rubric"
    i6f1671e5ccc67a30dbb510d7c37809232aa6c2612b707440d4a323edf5cdceb1 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/resources"
    i87cf461bbbb0a4957678adb2652f5d882f0056914766c119c39f374d5e351968 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/categories"
    i9952290c89ab227cf2447a4dcefa0a460e4f57f5f1fa1a544211747659e33ab5 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/publish"
    i1683357a3b1fcd506250383b696db7903aedda5919d87479cc06dbe504619859 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/submissions/item"
    ic1a63ca1d60be548c163492f15ab00755d812529295e8c920991227ea011d404 "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/categories/item"
    ieb0811fa685bb225bf1a7f78082f07573e816696c309fce9d12e426d809e5daa "github.com/microsoftgraph/msgraph-sdk-go/education/classes/item/assignments/item/resources/item"
)

// Builds and executes requests for operations under \education\classes\{educationClass-id}\assignments\{educationAssignment-id}
type EducationAssignmentRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// All assignments associated with this class. Nullable.
type EducationAssignmentRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *EducationAssignmentRequestBuilder) Categories()(*i87cf461bbbb0a4957678adb2652f5d882f0056914766c119c39f374d5e351968.CategoriesRequestBuilder) {
    return i87cf461bbbb0a4957678adb2652f5d882f0056914766c119c39f374d5e351968.NewCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignments.item.categories.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EducationAssignmentRequestBuilder) CategoriesById(id string)(*ic1a63ca1d60be548c163492f15ab00755d812529295e8c920991227ea011d404.EducationCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationCategory_id"] = id
    }
    return ic1a63ca1d60be548c163492f15ab00755d812529295e8c920991227ea011d404.NewEducationCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new EducationAssignmentRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEducationAssignmentRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentRequestBuilder) {
    m := &EducationAssignmentRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/education/classes/{educationClass_id}/assignments/{educationAssignment_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new EducationAssignmentRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewEducationAssignmentRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*EducationAssignmentRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewEducationAssignmentRequestBuilderInternal(urlParams, requestAdapter)
}
// All assignments associated with this class. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *EducationAssignmentRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// All assignments associated with this class. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *EducationAssignmentRequestBuilder) CreateGetRequestInformation(q func (value *EducationAssignmentRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(EducationAssignmentRequestBuilderGetQueryParameters)
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
// All assignments associated with this class. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *EducationAssignmentRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationAssignment, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// All assignments associated with this class. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *EducationAssignmentRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
// All assignments associated with this class. Nullable.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *EducationAssignmentRequestBuilder) Get(q func (value *EducationAssignmentRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationAssignment, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewEducationAssignment() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationAssignment), nil
}
// All assignments associated with this class. Nullable.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *EducationAssignmentRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.EducationAssignment, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *EducationAssignmentRequestBuilder) Publish()(*i9952290c89ab227cf2447a4dcefa0a460e4f57f5f1fa1a544211747659e33ab5.PublishRequestBuilder) {
    return i9952290c89ab227cf2447a4dcefa0a460e4f57f5f1fa1a544211747659e33ab5.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Resources()(*i6f1671e5ccc67a30dbb510d7c37809232aa6c2612b707440d4a323edf5cdceb1.ResourcesRequestBuilder) {
    return i6f1671e5ccc67a30dbb510d7c37809232aa6c2612b707440d4a323edf5cdceb1.NewResourcesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignments.item.resources.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EducationAssignmentRequestBuilder) ResourcesById(id string)(*ieb0811fa685bb225bf1a7f78082f07573e816696c309fce9d12e426d809e5daa.EducationAssignmentResourceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationAssignmentResource_id"] = id
    }
    return ieb0811fa685bb225bf1a7f78082f07573e816696c309fce9d12e426d809e5daa.NewEducationAssignmentResourceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Rubric()(*i560e0f51eb0b15ee54800eeae829fe0350e27c3694c083160844bef71444f4f0.RubricRequestBuilder) {
    return i560e0f51eb0b15ee54800eeae829fe0350e27c3694c083160844bef71444f4f0.NewRubricRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) SetUpResourcesFolder()(*i42850d7acdbd3a8e01158ce3f91031a284414d79a1b8423753fe9921702335f6.SetUpResourcesFolderRequestBuilder) {
    return i42850d7acdbd3a8e01158ce3f91031a284414d79a1b8423753fe9921702335f6.NewSetUpResourcesFolderRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *EducationAssignmentRequestBuilder) Submissions()(*i3d50a0849e57022c8564fae5dbbb2b471978d26a2db68975ed257f10e693e7cd.SubmissionsRequestBuilder) {
    return i3d50a0849e57022c8564fae5dbbb2b471978d26a2db68975ed257f10e693e7cd.NewSubmissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.education.classes.item.assignments.item.submissions.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *EducationAssignmentRequestBuilder) SubmissionsById(id string)(*i1683357a3b1fcd506250383b696db7903aedda5919d87479cc06dbe504619859.EducationSubmissionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["educationSubmission_id"] = id
    }
    return i1683357a3b1fcd506250383b696db7903aedda5919d87479cc06dbe504619859.NewEducationSubmissionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
