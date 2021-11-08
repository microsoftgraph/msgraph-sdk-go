package planner

import (
    i2e30aee32e30369c8ecc6bc38d177e6784a40fa62474eccf2569928795aaa39d "github.com/microsoftgraph/msgraph-sdk-go/planner/plans"
    i986fc893717836028d97e76e59d0d0f8dcd31663c6dae0666baf59f70e2d3160 "github.com/microsoftgraph/msgraph-sdk-go/planner/tasks"
    ib9a05c969511c90134a6d709165bc74e1713b24df408cc128f1e0dd7ead93ed5 "github.com/microsoftgraph/msgraph-sdk-go/planner/buckets"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i61fc96c998992d6a9f34876bca0e09cf8fd042f82ea4e76e5bcd7ec6488b05fb "github.com/microsoftgraph/msgraph-sdk-go/planner/buckets/item"
    ic1589f257c7e9d50ebdff1741b3e8c9aaea2fa2dbe14575d8dfdfaa82594b4be "github.com/microsoftgraph/msgraph-sdk-go/planner/tasks/item"
    if0be1a8c4a5c143221ab5041a7fe062880699c2169cf3b86e89b735af5f493e2 "github.com/microsoftgraph/msgraph-sdk-go/planner/plans/item"
)

// Builds and executes requests for operations under \planner
type PlannerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type PlannerRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *PlannerRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get planner
type PlannerRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type PlannerRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Planner;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *PlannerRequestBuilder) Buckets()(*ib9a05c969511c90134a6d709165bc74e1713b24df408cc128f1e0dd7ead93ed5.BucketsRequestBuilder) {
    return ib9a05c969511c90134a6d709165bc74e1713b24df408cc128f1e0dd7ead93ed5.NewBucketsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.planner.buckets.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PlannerRequestBuilder) BucketsById(id string)(*i61fc96c998992d6a9f34876bca0e09cf8fd042f82ea4e76e5bcd7ec6488b05fb.PlannerBucketRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerBucket_id"] = id
    }
    return i61fc96c998992d6a9f34876bca0e09cf8fd042f82ea4e76e5bcd7ec6488b05fb.NewPlannerBucketRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new PlannerRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerRequestBuilder) {
    m := &PlannerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/planner{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new PlannerRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewPlannerRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerRequestBuilderInternal(urlParams, requestAdapter)
}
// Get planner
// Parameters:
//  - options : Options for the request
func (m *PlannerRequestBuilder) CreateGetRequestInformation(options *PlannerRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Update planner
// Parameters:
//  - options : Options for the request
func (m *PlannerRequestBuilder) CreatePatchRequestInformation(options *PlannerRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get planner
// Parameters:
//  - options : Options for the request
func (m *PlannerRequestBuilder) Get(options *PlannerRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Planner, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPlanner() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Planner), nil
}
// Update planner
// Parameters:
//  - options : Options for the request
func (m *PlannerRequestBuilder) Patch(options *PlannerRequestBuilderPatchOptions)(error) {
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
func (m *PlannerRequestBuilder) Plans()(*i2e30aee32e30369c8ecc6bc38d177e6784a40fa62474eccf2569928795aaa39d.PlansRequestBuilder) {
    return i2e30aee32e30369c8ecc6bc38d177e6784a40fa62474eccf2569928795aaa39d.NewPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.planner.plans.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PlannerRequestBuilder) PlansById(id string)(*if0be1a8c4a5c143221ab5041a7fe062880699c2169cf3b86e89b735af5f493e2.PlannerPlanRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan_id"] = id
    }
    return if0be1a8c4a5c143221ab5041a7fe062880699c2169cf3b86e89b735af5f493e2.NewPlannerPlanRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PlannerRequestBuilder) Tasks()(*i986fc893717836028d97e76e59d0d0f8dcd31663c6dae0666baf59f70e2d3160.TasksRequestBuilder) {
    return i986fc893717836028d97e76e59d0d0f8dcd31663c6dae0666baf59f70e2d3160.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.planner.tasks.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *PlannerRequestBuilder) TasksById(id string)(*ic1589f257c7e9d50ebdff1741b3e8c9aaea2fa2dbe14575d8dfdfaa82594b4be.PlannerTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask_id"] = id
    }
    return ic1589f257c7e9d50ebdff1741b3e8c9aaea2fa2dbe14575d8dfdfaa82594b4be.NewPlannerTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
