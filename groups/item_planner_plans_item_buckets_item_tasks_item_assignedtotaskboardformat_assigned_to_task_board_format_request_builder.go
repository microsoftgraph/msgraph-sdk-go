package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder provides operations to manage the assignedToTaskBoardFormat property of the microsoft.graph.plannerTask entity.
type ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetQueryParameters read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
type ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetQueryParameters
}
// ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderInternal instantiates a new ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder and sets the default values.
func NewItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) {
    m := &ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/planner/plans/{plannerPlan%2Did}/buckets/{plannerBucket%2Did}/tasks/{plannerTask%2Did}/assignedToTaskBoardFormat{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder instantiates a new ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder and sets the default values.
func NewItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property assignedToTaskBoardFormat for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
// returns a PlannerAssignedToTaskBoardTaskFormatable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerAssignedToTaskBoardTaskFormatable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePlannerAssignedToTaskBoardTaskFormatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerAssignedToTaskBoardTaskFormatable), nil
}
// Patch update the navigation property assignedToTaskBoardFormat in groups
// returns a PlannerAssignedToTaskBoardTaskFormatable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerAssignedToTaskBoardTaskFormatable, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerAssignedToTaskBoardTaskFormatable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePlannerAssignedToTaskBoardTaskFormatFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerAssignedToTaskBoardTaskFormatable), nil
}
// ToDeleteRequestInformation delete navigation property assignedToTaskBoardFormat for groups
// returns a *RequestInformation when successful
func (m *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read-only. Nullable. Used to render the task correctly in the task board view when grouped by assignedTo.
// returns a *RequestInformation when successful
func (m *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property assignedToTaskBoardFormat in groups
// returns a *RequestInformation when successful
func (m *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PlannerAssignedToTaskBoardTaskFormatable, requestConfiguration *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder when successful
func (m *ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) WithUrl(rawUrl string)(*ItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder) {
    return NewItemPlannerPlansItemBucketsItemTasksItemAssignedtotaskboardformatAssignedToTaskBoardFormatRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
