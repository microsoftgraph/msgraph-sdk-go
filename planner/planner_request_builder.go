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

type PlannerRequestBuilder struct {
    pathParameters map[string]string;
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    urlTemplate string;
}
type PlannerRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    Expand []string;
    Select_escpaped []string;
}
func (m *PlannerRequestBuilder) Buckets()(*ib9a05c969511c90134a6d709165bc74e1713b24df408cc128f1e0dd7ead93ed5.BucketsRequestBuilder) {
    return ib9a05c969511c90134a6d709165bc74e1713b24df408cc128f1e0dd7ead93ed5.NewBucketsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerRequestBuilder) BucketsById(id string)(*i61fc96c998992d6a9f34876bca0e09cf8fd042f82ea4e76e5bcd7ec6488b05fb.PlannerBucketRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["plannerBucket_id"] = id
    }
    return i61fc96c998992d6a9f34876bca0e09cf8fd042f82ea4e76e5bcd7ec6488b05fb.NewPlannerBucketRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func NewPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerRequestBuilder) {
    m := &PlannerRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/planner{?select,expand}";
    urlTplParams := make(map[string]string)
    if pathParameters != nil {
        for idx, item := range pathParameters {
            urlTplParams[idx] = item
        }
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
func NewPlannerRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*PlannerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *PlannerRequestBuilder) CreateGetRequestInformation(q func (value *PlannerRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(PlannerRequestBuilderGetQueryParameters)
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
func (m *PlannerRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Planner, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *PlannerRequestBuilder) Get(q func (value *PlannerRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Planner, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPlanner() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Planner), nil
}
func (m *PlannerRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Planner, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
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
func (m *PlannerRequestBuilder) Plans()(*i2e30aee32e30369c8ecc6bc38d177e6784a40fa62474eccf2569928795aaa39d.PlansRequestBuilder) {
    return i2e30aee32e30369c8ecc6bc38d177e6784a40fa62474eccf2569928795aaa39d.NewPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerRequestBuilder) PlansById(id string)(*if0be1a8c4a5c143221ab5041a7fe062880699c2169cf3b86e89b735af5f493e2.PlannerPlanRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["plannerPlan_id"] = id
    }
    return if0be1a8c4a5c143221ab5041a7fe062880699c2169cf3b86e89b735af5f493e2.NewPlannerPlanRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *PlannerRequestBuilder) Tasks()(*i986fc893717836028d97e76e59d0d0f8dcd31663c6dae0666baf59f70e2d3160.TasksRequestBuilder) {
    return i986fc893717836028d97e76e59d0d0f8dcd31663c6dae0666baf59f70e2d3160.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *PlannerRequestBuilder) TasksById(id string)(*ic1589f257c7e9d50ebdff1741b3e8c9aaea2fa2dbe14575d8dfdfaa82594b4be.PlannerTaskRequestBuilder) {
    urlTplParams := make(map[string]string)
    if m.pathParameters != nil {
        for idx, item := range m.pathParameters {
            urlTplParams[idx] = item
        }
    }
    if id != "" {
        urlTplParams["plannerTask_id"] = id
    }
    return ic1589f257c7e9d50ebdff1741b3e8c9aaea2fa2dbe14575d8dfdfaa82594b4be.NewPlannerTaskRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
