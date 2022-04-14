package planner

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i2e30aee32e30369c8ecc6bc38d177e6784a40fa62474eccf2569928795aaa39d "github.com/microsoftgraph/msgraph-sdk-go/planner/plans"
    i986fc893717836028d97e76e59d0d0f8dcd31663c6dae0666baf59f70e2d3160 "github.com/microsoftgraph/msgraph-sdk-go/planner/tasks"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ib9a05c969511c90134a6d709165bc74e1713b24df408cc128f1e0dd7ead93ed5 "github.com/microsoftgraph/msgraph-sdk-go/planner/buckets"
    i61fc96c998992d6a9f34876bca0e09cf8fd042f82ea4e76e5bcd7ec6488b05fb "github.com/microsoftgraph/msgraph-sdk-go/planner/buckets/item"
    ic1589f257c7e9d50ebdff1741b3e8c9aaea2fa2dbe14575d8dfdfaa82594b4be "github.com/microsoftgraph/msgraph-sdk-go/planner/tasks/item"
    if0be1a8c4a5c143221ab5041a7fe062880699c2169cf3b86e89b735af5f493e2 "github.com/microsoftgraph/msgraph-sdk-go/planner/plans/item"
)

// PlannerRequestBuilder provides operations to manage the planner singleton.
type PlannerRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// PlannerRequestBuilderGetOptions options for Get
type PlannerRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *PlannerRequestBuilderGetQueryParameters
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// PlannerRequestBuilderGetQueryParameters get planner
type PlannerRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// PlannerRequestBuilderPatchOptions options for Patch
type PlannerRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Plannerable
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler
}
// Buckets the buckets property
func (m *PlannerRequestBuilder) Buckets()(*ib9a05c969511c90134a6d709165bc74e1713b24df408cc128f1e0dd7ead93ed5.BucketsRequestBuilder) {
    return ib9a05c969511c90134a6d709165bc74e1713b24df408cc128f1e0dd7ead93ed5.NewBucketsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BucketsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.planner.buckets.item collection
func (m *PlannerRequestBuilder) BucketsById(id string)(*i61fc96c998992d6a9f34876bca0e09cf8fd042f82ea4e76e5bcd7ec6488b05fb.PlannerBucketItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerBucket%2Did"] = id
    }
    return i61fc96c998992d6a9f34876bca0e09cf8fd042f82ea4e76e5bcd7ec6488b05fb.NewPlannerBucketItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPlannerRequestBuilderInternal instantiates a new PlannerRequestBuilder and sets the default values.
func NewPlannerRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerRequestBuilder) {
    m := &PlannerRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/planner{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPlannerRequestBuilder instantiates a new PlannerRequestBuilder and sets the default values.
func NewPlannerRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get planner
func (m *PlannerRequestBuilder) CreateGetRequestInformation(options *PlannerRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update planner
func (m *PlannerRequestBuilder) CreatePatchRequestInformation(options *PlannerRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get get planner
func (m *PlannerRequestBuilder) Get(options *PlannerRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Plannerable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePlannerFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Plannerable), nil
}
// Patch update planner
func (m *PlannerRequestBuilder) Patch(options *PlannerRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Plans the plans property
func (m *PlannerRequestBuilder) Plans()(*i2e30aee32e30369c8ecc6bc38d177e6784a40fa62474eccf2569928795aaa39d.PlansRequestBuilder) {
    return i2e30aee32e30369c8ecc6bc38d177e6784a40fa62474eccf2569928795aaa39d.NewPlansRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PlansById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.planner.plans.item collection
func (m *PlannerRequestBuilder) PlansById(id string)(*if0be1a8c4a5c143221ab5041a7fe062880699c2169cf3b86e89b735af5f493e2.PlannerPlanItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerPlan%2Did"] = id
    }
    return if0be1a8c4a5c143221ab5041a7fe062880699c2169cf3b86e89b735af5f493e2.NewPlannerPlanItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Tasks the tasks property
func (m *PlannerRequestBuilder) Tasks()(*i986fc893717836028d97e76e59d0d0f8dcd31663c6dae0666baf59f70e2d3160.TasksRequestBuilder) {
    return i986fc893717836028d97e76e59d0d0f8dcd31663c6dae0666baf59f70e2d3160.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.planner.tasks.item collection
func (m *PlannerRequestBuilder) TasksById(id string)(*ic1589f257c7e9d50ebdff1741b3e8c9aaea2fa2dbe14575d8dfdfaa82594b4be.PlannerTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask%2Did"] = id
    }
    return ic1589f257c7e9d50ebdff1741b3e8c9aaea2fa2dbe14575d8dfdfaa82594b4be.NewPlannerTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
