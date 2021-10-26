package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1ccb9a790e9d82e66ff780557fa3b6c9ac2107d22ea57dd2221476a3b828a42c "github.com/microsoftgraph/msgraph-sdk-go/me/planner/plans/item/details"
    i3aa95c30f5cac332a91a2c502c9d8acd486593c986b64f439a3f37faa7f6d400 "github.com/microsoftgraph/msgraph-sdk-go/me/planner/plans/item/buckets"
    i69f2355362701bc492379a1b06fa9623066e0e9fa09dd0918f3778e7898efc67 "github.com/microsoftgraph/msgraph-sdk-go/me/planner/plans/item/tasks"
    i3eaef46927d8aa78c404190c087e42549bb774eb776d36eea876b81d3a93ff78 "github.com/microsoftgraph/msgraph-sdk-go/me/planner/plans/item/buckets/item"
    i73511b20272f66b82b7aa067d7a59dd84908a7478b0d91553243df5d8bcc734e "github.com/microsoftgraph/msgraph-sdk-go/me/planner/plans/item/tasks/item"
)

// Builds and executes requests for operations under \me\planner\plans\{plannerPlan-id}
type PlannerPlanRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Read-only. Nullable. Returns the plannerTasks assigned to the user.
type PlannerPlanRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *PlannerPlanRequestBuilder) Buckets()(*i3aa95c30f5cac332a91a2c502c9d8acd486593c986b64f439a3f37faa7f6d400.BucketsRequestBuilder) {
    return i3aa95c30f5cac332a91a2c502c9d8acd486593c986b64f439a3f37faa7f6d400.NewBucketsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.planner.plans.item.buckets.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PlannerPlanRequestBuilder) BucketsById(id string)(*i3eaef46927d8aa78c404190c087e42549bb774eb776d36eea876b81d3a93ff78.PlannerBucketRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerBucket_id"] = id
    }
    return i3eaef46927d8aa78c404190c087e42549bb774eb776d36eea876b81d3a93ff78.NewPlannerBucketRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new PlannerPlanRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPlannerPlanRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerPlanRequestBuilder) {
    m := &PlannerPlanRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/me/planner/plans/{plannerPlan_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new PlannerPlanRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPlannerPlanRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerPlanRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerPlanRequestBuilderInternal(urlParams, requestAdapter)
}
// Read-only. Nullable. Returns the plannerTasks assigned to the user.
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *PlannerPlanRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable. Returns the plannerTasks assigned to the user.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *PlannerPlanRequestBuilder) CreateGetRequestInformation(q func (value *PlannerPlanRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PlannerPlanRequestBuilderGetQueryParameters)
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
// Read-only. Nullable. Returns the plannerTasks assigned to the user.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *PlannerPlanRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerPlan, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable. Returns the plannerTasks assigned to the user.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *PlannerPlanRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PlannerPlanRequestBuilder) Details()(*i1ccb9a790e9d82e66ff780557fa3b6c9ac2107d22ea57dd2221476a3b828a42c.DetailsRequestBuilder) {
    return i1ccb9a790e9d82e66ff780557fa3b6c9ac2107d22ea57dd2221476a3b828a42c.NewDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Read-only. Nullable. Returns the plannerTasks assigned to the user.
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *PlannerPlanRequestBuilder) Get(q func (value *PlannerPlanRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerPlan, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPlannerPlan() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerPlan), nil
}
// Read-only. Nullable. Returns the plannerTasks assigned to the user.
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *PlannerPlanRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerPlan, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PlannerPlanRequestBuilder) Tasks()(*i69f2355362701bc492379a1b06fa9623066e0e9fa09dd0918f3778e7898efc67.TasksRequestBuilder) {
    return i69f2355362701bc492379a1b06fa9623066e0e9fa09dd0918f3778e7898efc67.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.planner.plans.item.tasks.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PlannerPlanRequestBuilder) TasksById(id string)(*i73511b20272f66b82b7aa067d7a59dd84908a7478b0d91553243df5d8bcc734e.PlannerTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask_id"] = id
    }
    return i73511b20272f66b82b7aa067d7a59dd84908a7478b0d91553243df5d8bcc734e.NewPlannerTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
