package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder provides operations to manage the userConsentRequests property of the microsoft.graph.appConsentRequest entity.
type AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderGetQueryParameters read the properties and relationships of a userConsentRequest object.
type AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderGetQueryParameters
}
// AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Approval provides operations to manage the approval property of the microsoft.graph.userConsentRequest entity.
// returns a *AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder) Approval()(*AppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilder) {
    return NewAppconsentAppconsentrequestsItemUserconsentrequestsItemApprovalRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderInternal instantiates a new AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder and sets the default values.
func NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder) {
    m := &AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/appConsent/appConsentRequests/{appConsentRequest%2Did}/userConsentRequests/{userConsentRequest%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder instantiates a new AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder and sets the default values.
func NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property userConsentRequests for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a userConsentRequest object.
// returns a UserConsentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/userconsentrequest-get?view=graph-rest-1.0
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder) Get(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserConsentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable), nil
}
// Patch update the navigation property userConsentRequests in identityGovernance
// returns a UserConsentRequestable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserConsentRequestFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable), nil
}
// ToDeleteRequestInformation delete navigation property userConsentRequests for identityGovernance
// returns a *RequestInformation when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a userConsentRequest object.
// returns a *RequestInformation when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property userConsentRequests in identityGovernance
// returns a *RequestInformation when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UserConsentRequestable, requestConfiguration *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder when successful
func (m *AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder) WithUrl(rawUrl string)(*AppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder) {
    return NewAppconsentAppconsentrequestsItemUserconsentrequestsUserConsentRequestItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
