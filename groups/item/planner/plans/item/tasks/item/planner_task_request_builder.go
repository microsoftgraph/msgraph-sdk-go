package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0254b76d9a912abf790fc0e01cc3681fb5772c4f076a4819c9ce5422523c2ae6 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/tasks/item/progresstaskboardformat"
    i495f9972695d085e0bdac47e8ae4c30b326357a034263710344f12b65e49fcdd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/tasks/item/details"
    i7aea6d7f015a1cab2712e5ae98974cc82abe0e2903ac814c7dfb88ca82445edb "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/tasks/item/buckettaskboardformat"
    i941f474b2a105ecf46500f512524003cc8eccbdcc58f63e19ddbd402c03f86cd "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/tasks/item/assignedtotaskboardformat"
)

// PlannerTaskRequestBuilder builds and executes requests for operations under \groups\{group-id}\planner\plans\{plannerPlan-id}\tasks\{plannerTask-id}
type PlannerTaskRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PlannerTaskRequestBuilderDeleteOptions options for Delete
type PlannerTaskRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PlannerTaskRequestBuilderGetOptions options for Get
type PlannerTaskRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PlannerTaskRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// PlannerTaskRequestBuilderGetQueryParameters collection of tasks in the plan. Read-only. Nullable.
type PlannerTaskRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PlannerTaskRequestBuilderPatchOptions options for Patch
type PlannerTaskRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerTask;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PlannerTaskRequestBuilder) AssignedToTaskBoardFormat()(*i941f474b2a105ecf46500f512524003cc8eccbdcc58f63e19ddbd402c03f86cd.AssignedToTaskBoardFormatRequestBuilder) {
    return i941f474b2a105ecf46500f512524003cc8eccbdcc58f63e19ddbd402c03f86cd.NewAssignedToTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerTaskRequestBuilder) BucketTaskBoardFormat()(*i7aea6d7f015a1cab2712e5ae98974cc82abe0e2903ac814c7dfb88ca82445edb.BucketTaskBoardFormatRequestBuilder) {
    return i7aea6d7f015a1cab2712e5ae98974cc82abe0e2903ac814c7dfb88ca82445edb.NewBucketTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewPlannerTaskRequestBuilderInternal instantiates a new PlannerTaskRequestBuilder and sets the default values.
func NewPlannerTaskRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerTaskRequestBuilder) {
    m := &PlannerTaskRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/planner/plans/{plannerPlan_id}/tasks/{plannerTask_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPlannerTaskRequestBuilder instantiates a new PlannerTaskRequestBuilder and sets the default values.
func NewPlannerTaskRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerTaskRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerTaskRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation collection of tasks in the plan. Read-only. Nullable.
func (m *PlannerTaskRequestBuilder) CreateDeleteRequestInformation(options *PlannerTaskRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation collection of tasks in the plan. Read-only. Nullable.
func (m *PlannerTaskRequestBuilder) CreateGetRequestInformation(options *PlannerTaskRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation collection of tasks in the plan. Read-only. Nullable.
func (m *PlannerTaskRequestBuilder) CreatePatchRequestInformation(options *PlannerTaskRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete collection of tasks in the plan. Read-only. Nullable.
func (m *PlannerTaskRequestBuilder) Delete(options *PlannerTaskRequestBuilderDeleteOptions)(error) {
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
func (m *PlannerTaskRequestBuilder) Details()(*i495f9972695d085e0bdac47e8ae4c30b326357a034263710344f12b65e49fcdd.DetailsRequestBuilder) {
    return i495f9972695d085e0bdac47e8ae4c30b326357a034263710344f12b65e49fcdd.NewDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get collection of tasks in the plan. Read-only. Nullable.
func (m *PlannerTaskRequestBuilder) Get(options *PlannerTaskRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerTask, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPlannerTask() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PlannerTask), nil
}
// Patch collection of tasks in the plan. Read-only. Nullable.
func (m *PlannerTaskRequestBuilder) Patch(options *PlannerTaskRequestBuilderPatchOptions)(error) {
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
func (m *PlannerTaskRequestBuilder) ProgressTaskBoardFormat()(*i0254b76d9a912abf790fc0e01cc3681fb5772c4f076a4819c9ce5422523c2ae6.ProgressTaskBoardFormatRequestBuilder) {
    return i0254b76d9a912abf790fc0e01cc3681fb5772c4f076a4819c9ce5422523c2ae6.NewProgressTaskBoardFormatRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
