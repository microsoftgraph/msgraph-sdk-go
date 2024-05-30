package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// LabelsRetentionlabelsRetentionLabelItemRequestBuilder provides operations to manage the retentionLabels property of the microsoft.graph.security.labelsRoot entity.
type LabelsRetentionlabelsRetentionLabelItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// LabelsRetentionlabelsRetentionLabelItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsRetentionlabelsRetentionLabelItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// LabelsRetentionlabelsRetentionLabelItemRequestBuilderGetQueryParameters represents how customers can manage their data, whether and for how long to retain or delete it.
type LabelsRetentionlabelsRetentionLabelItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// LabelsRetentionlabelsRetentionLabelItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsRetentionlabelsRetentionLabelItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *LabelsRetentionlabelsRetentionLabelItemRequestBuilderGetQueryParameters
}
// LabelsRetentionlabelsRetentionLabelItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type LabelsRetentionlabelsRetentionLabelItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewLabelsRetentionlabelsRetentionLabelItemRequestBuilderInternal instantiates a new LabelsRetentionlabelsRetentionLabelItemRequestBuilder and sets the default values.
func NewLabelsRetentionlabelsRetentionLabelItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsRetentionlabelsRetentionLabelItemRequestBuilder) {
    m := &LabelsRetentionlabelsRetentionLabelItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/labels/retentionLabels/{retentionLabel%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewLabelsRetentionlabelsRetentionLabelItemRequestBuilder instantiates a new LabelsRetentionlabelsRetentionLabelItemRequestBuilder and sets the default values.
func NewLabelsRetentionlabelsRetentionLabelItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*LabelsRetentionlabelsRetentionLabelItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewLabelsRetentionlabelsRetentionLabelItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete a retentionLabel object.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-retentionlabel-delete?view=graph-rest-1.0
func (m *LabelsRetentionlabelsRetentionLabelItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *LabelsRetentionlabelsRetentionLabelItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Descriptors provides operations to manage the descriptors property of the microsoft.graph.security.retentionLabel entity.
// returns a *LabelsRetentionlabelsItemDescriptorsRequestBuilder when successful
func (m *LabelsRetentionlabelsRetentionLabelItemRequestBuilder) Descriptors()(*LabelsRetentionlabelsItemDescriptorsRequestBuilder) {
    return NewLabelsRetentionlabelsItemDescriptorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// DispositionReviewStages provides operations to manage the dispositionReviewStages property of the microsoft.graph.security.retentionLabel entity.
// returns a *LabelsRetentionlabelsItemDispositionreviewstagesDispositionReviewStagesRequestBuilder when successful
func (m *LabelsRetentionlabelsRetentionLabelItemRequestBuilder) DispositionReviewStages()(*LabelsRetentionlabelsItemDispositionreviewstagesDispositionReviewStagesRequestBuilder) {
    return NewLabelsRetentionlabelsItemDispositionreviewstagesDispositionReviewStagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents how customers can manage their data, whether and for how long to retain or delete it.
// returns a RetentionLabelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *LabelsRetentionlabelsRetentionLabelItemRequestBuilder) Get(ctx context.Context, requestConfiguration *LabelsRetentionlabelsRetentionLabelItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionLabelable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateRetentionLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionLabelable), nil
}
// Patch update the properties of a retentionLabel object. To update a disposition review stage, include the actionAfterRetentionPeriod property in the request body with one of the possible values specified.
// returns a RetentionLabelable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-retentionlabel-update?view=graph-rest-1.0
func (m *LabelsRetentionlabelsRetentionLabelItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionLabelable, requestConfiguration *LabelsRetentionlabelsRetentionLabelItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionLabelable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateRetentionLabelFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionLabelable), nil
}
// RetentionEventType provides operations to manage the retentionEventType property of the microsoft.graph.security.retentionLabel entity.
// returns a *LabelsRetentionlabelsItemRetentioneventtypeRetentionEventTypeRequestBuilder when successful
func (m *LabelsRetentionlabelsRetentionLabelItemRequestBuilder) RetentionEventType()(*LabelsRetentionlabelsItemRetentioneventtypeRetentionEventTypeRequestBuilder) {
    return NewLabelsRetentionlabelsItemRetentioneventtypeRetentionEventTypeRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete a retentionLabel object.
// returns a *RequestInformation when successful
func (m *LabelsRetentionlabelsRetentionLabelItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *LabelsRetentionlabelsRetentionLabelItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation represents how customers can manage their data, whether and for how long to retain or delete it.
// returns a *RequestInformation when successful
func (m *LabelsRetentionlabelsRetentionLabelItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *LabelsRetentionlabelsRetentionLabelItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a retentionLabel object. To update a disposition review stage, include the actionAfterRetentionPeriod property in the request body with one of the possible values specified.
// returns a *RequestInformation when successful
func (m *LabelsRetentionlabelsRetentionLabelItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.RetentionLabelable, requestConfiguration *LabelsRetentionlabelsRetentionLabelItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *LabelsRetentionlabelsRetentionLabelItemRequestBuilder when successful
func (m *LabelsRetentionlabelsRetentionLabelItemRequestBuilder) WithUrl(rawUrl string)(*LabelsRetentionlabelsRetentionLabelItemRequestBuilder) {
    return NewLabelsRetentionlabelsRetentionLabelItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
