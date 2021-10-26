package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i5af8d16af1dad61fda6b4a358f4be036e851168cc0c513919262939cfd659d80 "github.com/microsoftgraph/msgraph-sdk-go/me/planner/tasks/item/assignedtotaskboardformat"
    ia5b3f1e7520441c7435be0a6ae5dae101af2c3ceb6e3a91756368698d3726cae "github.com/microsoftgraph/msgraph-sdk-go/me/planner/tasks/item/buckettaskboardformat"
    ic6fcfd0881f9296d786a474cbc46f4eca4aaf2274ffe94b844cbccf2daf510d6 "github.com/microsoftgraph/msgraph-sdk-go/me/planner/tasks/item/details"
    if3edfa04ca1f46739b28008c9096861864ccefc9bd284dc8cf84ec07190ab653 "github.com/microsoftgraph/msgraph-sdk-go/me/planner/tasks/item/progresstaskboardformat"
)

// Builds and executes requests for operations under \me\planner\tasks\{plannerTask-id}
type PlannerTaskRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Read-only. Nullable. Returns the plannerPlans shared with the user.
type PlannerTaskRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *PlannerTaskRequestBuilder) AssignedToTaskBoardFormat()(*i5af8d16af1dad61fda6b4a358f4be036e851168cc0c513919262939cfd659d80.AssignedToTaskBoardFormatRequestBuilder) {
    return i5af8d16af1dad61fda6b4a358f4be036e851168cc0c513919262939cfd659d80.NewAssignedToTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerTaskRequestBuilder) BucketTaskBoardFormat()(*ia5b3f1e7520441c7435be0a6ae5dae101af2c3ceb6e3a91756368698d3726cae.BucketTaskBoardFormatRequestBuilder) {
    return ia5b3f1e7520441c7435be0a6ae5dae101af2c3ceb6e3a91756368698d3726cae.NewBucketTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Instantiates a new PlannerTaskRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPlannerTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerTaskRequestBuilder) {
    m := &PlannerTaskRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/planner/tasks/{plannerTask_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new PlannerTaskRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPlannerTaskRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerTaskRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerTaskRequestBuilderInternal(urlParams, requestAdapter)
}
// Read-only. Nullable. Returns the plannerPlans shared with the user.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *PlannerTaskRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable. Returns the plannerPlans shared with the user.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *PlannerTaskRequestBuilder) CreateGetRequestInformation(q func (value *PlannerTaskRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PlannerTaskRequestBuilderGetQueryParameters)
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
// Read-only. Nullable. Returns the plannerPlans shared with the user.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *PlannerTaskRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerTask, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable. Returns the plannerPlans shared with the user.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *PlannerTaskRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PlannerTaskRequestBuilder) Details()(*ic6fcfd0881f9296d786a474cbc46f4eca4aaf2274ffe94b844cbccf2daf510d6.DetailsRequestBuilder) {
    return ic6fcfd0881f9296d786a474cbc46f4eca4aaf2274ffe94b844cbccf2daf510d6.NewDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Read-only. Nullable. Returns the plannerPlans shared with the user.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *PlannerTaskRequestBuilder) Get(q func (value *PlannerTaskRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerTask, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPlannerTask() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerTask), nil
}
// Read-only. Nullable. Returns the plannerPlans shared with the user.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *PlannerTaskRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerTask, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PlannerTaskRequestBuilder) ProgressTaskBoardFormat()(*if3edfa04ca1f46739b28008c9096861864ccefc9bd284dc8cf84ec07190ab653.ProgressTaskBoardFormatRequestBuilder) {
    return if3edfa04ca1f46739b28008c9096861864ccefc9bd284dc8cf84ec07190ab653.NewProgressTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
