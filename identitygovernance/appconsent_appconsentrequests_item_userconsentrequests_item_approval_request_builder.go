package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder provides operations to manage the approval property of the microsoft.graph.userConsentRequest entity.
type AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderGetQueryParameters approval decisions associated with a request.
type AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderGetQueryParameters
}
// AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderInternal instantiates a new AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder and sets the default values.
func NewAppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder) {
    m := &AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/appConsent/appConsentRequests/{appConsentRequest%2Did}/userConsentRequests/{userConsentRequest%2Did}/approval{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder instantiates a new AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder and sets the default values.
func NewAppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property approval for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get approval decisions associated with a request.
// returns a Approvalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder) Get(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Approvalable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateApprovalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Approvalable), nil
}
// Patch update the navigation property approval in identityGovernance
// returns a Approvalable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Approvalable, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Approvalable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateApprovalFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Approvalable), nil
}
// Stages provides operations to manage the stages property of the microsoft.graph.approval entity.
// returns a *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalStagesRequestBuilder when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder) Stages()(*AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalStagesRequestBuilder) {
    return NewAppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalStagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property approval for identityGovernance
// returns a *RequestInformation when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation approval decisions associated with a request.
// returns a *RequestInformation when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property approval in identityGovernance
// returns a *RequestInformation when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Approvalable, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder) WithUrl(rawUrl string)(*AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder) {
    return NewAppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
