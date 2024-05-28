package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder provides operations to manage the bucketTaskBoardFormat property of the microsoft.graph.plannerTask entity.
type ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderGetQueryParameters read-only. Nullable. Used to render the task correctly in the task board view when grouped by bucket.
type ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderGetQueryParameters
}
// ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderInternal instantiates a new ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder and sets the default values.
func NewItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder) {
    m := &ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/planner/plans/{plannerPlan%2Did}/buckets/{plannerBucket%2Did}/tasks/{plannerTask%2Did}/bucketTaskBoardFormat{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder instantiates a new ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder and sets the default values.
func NewItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property bucketTaskBoardFormat for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderDeleteRequestConfiguration)(error) {
    requestInfo, err := m.ToDeleteRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.BaseRequestBuilder.RequestAdapter.SendNoContent(ctx, requestInfo, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get read-only. Nullable. Used to render the task correctly in the task board view when grouped by bucket.
// returns a PlannerBucketTaskBoardTaskFormatable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerBucketTaskBoardTaskFormatable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePlannerBucketTaskBoardTaskFormatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerBucketTaskBoardTaskFormatable), nil
}
// Patch update the navigation property bucketTaskBoardFormat in groups
// returns a PlannerBucketTaskBoardTaskFormatable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerBucketTaskBoardTaskFormatable, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerBucketTaskBoardTaskFormatable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePlannerBucketTaskBoardTaskFormatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerBucketTaskBoardTaskFormatable), nil
}
// ToDeleteRequestInformation delete navigation property bucketTaskBoardFormat for groups
// returns a *RequestInformation when successful
func (m *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read-only. Nullable. Used to render the task correctly in the task board view when grouped by bucket.
// returns a *RequestInformation when successful
func (m *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToPatchRequestInformation update the navigation property bucketTaskBoardFormat in groups
// returns a *RequestInformation when successful
func (m *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerBucketTaskBoardTaskFormatable, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    err := requestInfo.SetContentFromParsable(ctx, m.BaseRequestBuilder.RequestAdapter, "application/json", body)
    if err != nil {
        return nil, err
    }
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder when successful
func (m *ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder) WithUrl(rawUrl string)(*ItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder) {
    return NewItemPlannerPlansItemBucketsItemTasksItemBuckettaskboardformatBucketTaskBoardFormatRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
