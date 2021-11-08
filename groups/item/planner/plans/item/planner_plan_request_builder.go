package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i71563296eb2eb02c4e099b4ac0d877faf391b44babcddac0f302d0dfa966e39b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/buckets"
    i91a46074a79025105fd6d697b9144b42ef9be801e1ac60a82fd79fa290e172c8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/tasks"
    if58d3b2a2e144a13025259faa98b7faf20c554a797c638968db7d6b76fb94bd0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/details"
    iebf76be10788d0e38135a2d0301032a64b4e80b4cd584750240147430a330b34 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/buckets/item"
    iee9a3189c1fe1484426aec41b0bf61667bd357944a598d934891b65f9d744c40 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/tasks/item"
)

// Builds and executes requests for operations under \groups\{group-id}\planner\plans\{plannerPlan-id}
type PlannerPlanRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type PlannerPlanRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type PlannerPlanRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PlannerPlanRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Read-only. Nullable. Returns the plannerPlans owned by the group.
type PlannerPlanRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type PlannerPlanRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerPlan;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PlannerPlanRequestBuilder) Buckets()(*i71563296eb2eb02c4e099b4ac0d877faf391b44babcddac0f302d0dfa966e39b.BucketsRequestBuilder) {
    return i71563296eb2eb02c4e099b4ac0d877faf391b44babcddac0f302d0dfa966e39b.NewBucketsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.groups.item.planner.plans.item.buckets.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PlannerPlanRequestBuilder) BucketsById(id string)(*iebf76be10788d0e38135a2d0301032a64b4e80b4cd584750240147430a330b34.PlannerBucketRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerBucket_id"] = id
    }
    return iebf76be10788d0e38135a2d0301032a64b4e80b4cd584750240147430a330b34.NewPlannerBucketRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new PlannerPlanRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPlannerPlanRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerPlanRequestBuilder) {
    m := &PlannerPlanRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/planner/plans/{plannerPlan_id}{?select,expand}";
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
// Read-only. Nullable. Returns the plannerPlans owned by the group.
// Parameters:
//  - options : Options for the request
func (m *PlannerPlanRequestBuilder) CreateDeleteRequestInformation(options *PlannerPlanRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable. Returns the plannerPlans owned by the group.
// Parameters:
//  - options : Options for the request
func (m *PlannerPlanRequestBuilder) CreateGetRequestInformation(options *PlannerPlanRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Read-only. Nullable. Returns the plannerPlans owned by the group.
// Parameters:
//  - options : Options for the request
func (m *PlannerPlanRequestBuilder) CreatePatchRequestInformation(options *PlannerPlanRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Read-only. Nullable. Returns the plannerPlans owned by the group.
// Parameters:
//  - options : Options for the request
func (m *PlannerPlanRequestBuilder) Delete(options *PlannerPlanRequestBuilderDeleteOptions)(error) {
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
func (m *PlannerPlanRequestBuilder) Details()(*if58d3b2a2e144a13025259faa98b7faf20c554a797c638968db7d6b76fb94bd0.DetailsRequestBuilder) {
    return if58d3b2a2e144a13025259faa98b7faf20c554a797c638968db7d6b76fb94bd0.NewDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Read-only. Nullable. Returns the plannerPlans owned by the group.
// Parameters:
//  - options : Options for the request
func (m *PlannerPlanRequestBuilder) Get(options *PlannerPlanRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerPlan, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPlannerPlan() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerPlan), nil
}
// Read-only. Nullable. Returns the plannerPlans owned by the group.
// Parameters:
//  - options : Options for the request
func (m *PlannerPlanRequestBuilder) Patch(options *PlannerPlanRequestBuilderPatchOptions)(error) {
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
func (m *PlannerPlanRequestBuilder) Tasks()(*i91a46074a79025105fd6d697b9144b42ef9be801e1ac60a82fd79fa290e172c8.TasksRequestBuilder) {
    return i91a46074a79025105fd6d697b9144b42ef9be801e1ac60a82fd79fa290e172c8.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.groups.item.planner.plans.item.tasks.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PlannerPlanRequestBuilder) TasksById(id string)(*iee9a3189c1fe1484426aec41b0bf61667bd357944a598d934891b65f9d744c40.PlannerTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask_id"] = id
    }
    return iee9a3189c1fe1484426aec41b0bf61667bd357944a598d934891b65f9d744c40.NewPlannerTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
