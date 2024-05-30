package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder provides operations to manage the appConsentRequests property of the microsoft.graph.appConsentApprovalRoute entity.
type AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderGetQueryParameters read the properties and relationships of an appConsentRequest object.
type AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderGetQueryParameters
}
// AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewAppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderInternal instantiates a new AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder and sets the default values.
func NewAppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder) {
    m := &AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/appConsent/appConsentRequests/{appConsentRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder instantiates a new AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder and sets the default values.
func NewAppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property appConsentRequests for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of an appConsentRequest object.
// returns a AppConsentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/appconsentrequest-get?view=graph-rest-1.0
func (m *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppConsentRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAppConsentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppConsentRequestable), nil
}
// Patch update the navigation property appConsentRequests in identityGovernance
// returns a AppConsentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppConsentRequestable, requestConfiguration *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppConsentRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAppConsentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppConsentRequestable), nil
}
// ToDeleteRequestInformation delete navigation property appConsentRequests for identityGovernance
// returns a *RequestInformation when successful
func (m *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of an appConsentRequest object.
// returns a *RequestInformation when successful
func (m *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property appConsentRequests in identityGovernance
// returns a *RequestInformation when successful
func (m *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AppConsentRequestable, requestConfiguration *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// UserConsentRequests provides operations to manage the userConsentRequests property of the microsoft.graph.appConsentRequest entity.
// returns a *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder when successful
func (m *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder) UserConsentRequests()(*AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilder) {
    return NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder when successful
func (m *AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder) WithUrl(rawUrl string)(*AppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder) {
    return NewAppconsentAppconsentrequestsAppConsentRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
