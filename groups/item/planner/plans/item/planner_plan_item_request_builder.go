package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i71563296eb2eb02c4e099b4ac0d877faf391b44babcddac0f302d0dfa966e39b "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/buckets"
    i91a46074a79025105fd6d697b9144b42ef9be801e1ac60a82fd79fa290e172c8 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/tasks"
    if58d3b2a2e144a13025259faa98b7faf20c554a797c638968db7d6b76fb94bd0 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/details"
    iebf76be10788d0e38135a2d0301032a64b4e80b4cd584750240147430a330b34 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/buckets/item"
    iee9a3189c1fe1484426aec41b0bf61667bd357944a598d934891b65f9d744c40 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/planner/plans/item/tasks/item"
)

// PlannerPlanItemRequestBuilder provides operations to manage the plans property of the microsoft.graph.plannerGroup entity.
type PlannerPlanItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// PlannerPlanItemRequestBuilderDeleteOptions options for Delete
type PlannerPlanItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// PlannerPlanItemRequestBuilderGetOptions options for Get
type PlannerPlanItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *PlannerPlanItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// PlannerPlanItemRequestBuilderGetQueryParameters read-only. Nullable. Returns the plannerPlans owned by the group.
type PlannerPlanItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// PlannerPlanItemRequestBuilderPatchOptions options for Patch
type PlannerPlanItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerPlanable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// Buckets the buckets property
func (m *PlannerPlanItemRequestBuilder) Buckets()(*i71563296eb2eb02c4e099b4ac0d877faf391b44babcddac0f302d0dfa966e39b.BucketsRequestBuilder) {
    return i71563296eb2eb02c4e099b4ac0d877faf391b44babcddac0f302d0dfa966e39b.NewBucketsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BucketsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.planner.plans.item.buckets.item collection
func (m *PlannerPlanItemRequestBuilder) BucketsById(id string)(*iebf76be10788d0e38135a2d0301032a64b4e80b4cd584750240147430a330b34.PlannerBucketItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerBucket_id"] = id
    }
    return iebf76be10788d0e38135a2d0301032a64b4e80b4cd584750240147430a330b34.NewPlannerBucketItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewPlannerPlanItemRequestBuilderInternal instantiates a new PlannerPlanItemRequestBuilder and sets the default values.
func NewPlannerPlanItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerPlanItemRequestBuilder) {
    m := &PlannerPlanItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/planner/plans/{plannerPlan_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewPlannerPlanItemRequestBuilder instantiates a new PlannerPlanItemRequestBuilder and sets the default values.
func NewPlannerPlanItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*PlannerPlanItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewPlannerPlanItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property plans for groups
func (m *PlannerPlanItemRequestBuilder) CreateDeleteRequestInformation(options *PlannerPlanItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
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
// CreateGetRequestInformation read-only. Nullable. Returns the plannerPlans owned by the group.
func (m *PlannerPlanItemRequestBuilder) CreateGetRequestInformation(options *PlannerPlanItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property plans in groups
func (m *PlannerPlanItemRequestBuilder) CreatePatchRequestInformation(options *PlannerPlanItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Delete delete navigation property plans for groups
func (m *PlannerPlanItemRequestBuilder) Delete(options *PlannerPlanItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
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
// Details the details property
func (m *PlannerPlanItemRequestBuilder) Details()(*if58d3b2a2e144a13025259faa98b7faf20c554a797c638968db7d6b76fb94bd0.DetailsRequestBuilder) {
    return if58d3b2a2e144a13025259faa98b7faf20c554a797c638968db7d6b76fb94bd0.NewDetailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get read-only. Nullable. Returns the plannerPlans owned by the group.
func (m *PlannerPlanItemRequestBuilder) Get(options *PlannerPlanItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerPlanable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePlannerPlanFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerPlanable), nil
}
// Patch update the navigation property plans in groups
func (m *PlannerPlanItemRequestBuilder) Patch(options *PlannerPlanItemRequestBuilderPatchOptions)(error) {
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
// Tasks the tasks property
func (m *PlannerPlanItemRequestBuilder) Tasks()(*i91a46074a79025105fd6d697b9144b42ef9be801e1ac60a82fd79fa290e172c8.TasksRequestBuilder) {
    return i91a46074a79025105fd6d697b9144b42ef9be801e1ac60a82fd79fa290e172c8.NewTasksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TasksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.planner.plans.item.tasks.item collection
func (m *PlannerPlanItemRequestBuilder) TasksById(id string)(*iee9a3189c1fe1484426aec41b0bf61667bd357944a598d934891b65f9d744c40.PlannerTaskItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["plannerTask_id"] = id
    }
    return iee9a3189c1fe1484426aec41b0bf61667bd357944a598d934891b65f9d744c40.NewPlannerTaskItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
