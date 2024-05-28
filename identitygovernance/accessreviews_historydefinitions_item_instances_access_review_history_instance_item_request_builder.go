package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder provides operations to manage the instances property of the microsoft.graph.accessReviewHistoryDefinition entity.
type AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderGetQueryParameters if the accessReviewHistoryDefinition is a recurring definition, instances represent each recurrence. A definition that doesn't recur will have exactly one instance.
type AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderGetQueryParameters
}
// AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderInternal instantiates a new AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder and sets the default values.
func NewAccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder) {
    m := &AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/accessReviews/historyDefinitions/{accessReviewHistoryDefinition%2Did}/instances/{accessReviewHistoryInstance%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder instantiates a new AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder and sets the default values.
func NewAccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property instances for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// GenerateDownloadUri provides operations to call the generateDownloadUri method.
// returns a *AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder when successful
func (m *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder) GenerateDownloadUri()(*AccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilder) {
    return NewAccessreviewsHistorydefinitionsItemInstancesItemGeneratedownloaduriGenerateDownloadUriRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get if the accessReviewHistoryDefinition is a recurring definition, instances represent each recurrence. A definition that doesn't recur will have exactly one instance.
// returns a AccessReviewHistoryInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewHistoryInstanceable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessReviewHistoryInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewHistoryInstanceable), nil
}
// Patch update the navigation property instances in identityGovernance
// returns a AccessReviewHistoryInstanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewHistoryInstanceable, requestConfiguration *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewHistoryInstanceable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAccessReviewHistoryInstanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewHistoryInstanceable), nil
}
// ToDeleteRequestInformation delete navigation property instances for identityGovernance
// returns a *RequestInformation when successful
func (m *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation if the accessReviewHistoryDefinition is a recurring definition, instances represent each recurrence. A definition that doesn't recur will have exactly one instance.
// returns a *RequestInformation when successful
func (m *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property instances in identityGovernance
// returns a *RequestInformation when successful
func (m *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AccessReviewHistoryInstanceable, requestConfiguration *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder when successful
func (m *AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder) WithUrl(rawUrl string)(*AccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder) {
    return NewAccessreviewsHistorydefinitionsItemInstancesAccessReviewHistoryInstanceItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
