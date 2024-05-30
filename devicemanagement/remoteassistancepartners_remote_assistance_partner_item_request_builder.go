package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder provides operations to manage the remoteAssistancePartners property of the microsoft.graph.deviceManagement entity.
type RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderGetQueryParameters read properties and relationships of the remoteAssistancePartner object.
type RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderGetQueryParameters
}
// RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// BeginOnboarding provides operations to call the beginOnboarding method.
// returns a *RemoteassistancepartnersItemBeginonboardingBeginOnboardingRequestBuilder when successful
func (m *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder) BeginOnboarding()(*RemoteassistancepartnersItemBeginonboardingBeginOnboardingRequestBuilder) {
    return NewRemoteassistancepartnersItemBeginonboardingBeginOnboardingRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewRemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderInternal instantiates a new RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder and sets the default values.
func NewRemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder) {
    m := &RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/remoteAssistancePartners/{remoteAssistancePartner%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewRemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder instantiates a new RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder and sets the default values.
func NewRemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete deletes a remoteAssistancePartner.
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-remoteassistance-remoteassistancepartner-delete?view=graph-rest-1.0
func (m *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Disconnect provides operations to call the disconnect method.
// returns a *RemoteassistancepartnersItemDisconnectRequestBuilder when successful
func (m *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder) Disconnect()(*RemoteassistancepartnersItemDisconnectRequestBuilder) {
    return NewRemoteassistancepartnersItemDisconnectRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read properties and relationships of the remoteAssistancePartner object.
// returns a RemoteAssistancePartnerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-remoteassistance-remoteassistancepartner-get?view=graph-rest-1.0
func (m *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder) Get(ctx context.Context, requestConfiguration *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemoteAssistancePartnerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRemoteAssistancePartnerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemoteAssistancePartnerable), nil
}
// Patch update the properties of a remoteAssistancePartner object.
// returns a RemoteAssistancePartnerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-remoteassistance-remoteassistancepartner-update?view=graph-rest-1.0
func (m *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemoteAssistancePartnerable, requestConfiguration *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemoteAssistancePartnerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateRemoteAssistancePartnerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemoteAssistancePartnerable), nil
}
// ToDeleteRequestInformation deletes a remoteAssistancePartner.
// returns a *RequestInformation when successful
func (m *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read properties and relationships of the remoteAssistancePartner object.
// returns a *RequestInformation when successful
func (m *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the properties of a remoteAssistancePartner object.
// returns a *RequestInformation when successful
func (m *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.RemoteAssistancePartnerable, requestConfiguration *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder when successful
func (m *RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder) WithUrl(rawUrl string)(*RemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder) {
    return NewRemoteassistancepartnersRemoteAssistancePartnerItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
