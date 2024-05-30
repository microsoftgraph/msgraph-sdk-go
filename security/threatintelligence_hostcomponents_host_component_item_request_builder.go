package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder provides operations to manage the hostComponents property of the microsoft.graph.security.threatIntelligence entity.
type ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderGetQueryParameters read the properties and relationships of a hostComponent object.
type ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderGetQueryParameters
}
// ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewThreatintelligenceHostcomponentsHostComponentItemRequestBuilderInternal instantiates a new ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder and sets the default values.
func NewThreatintelligenceHostcomponentsHostComponentItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder) {
    m := &ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/hostComponents/{hostComponent%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatintelligenceHostcomponentsHostComponentItemRequestBuilder instantiates a new ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder and sets the default values.
func NewThreatintelligenceHostcomponentsHostComponentItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceHostcomponentsHostComponentItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property hostComponents for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a hostComponent object.
// returns a HostComponentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-hostcomponent-get?view=graph-rest-1.0
func (m *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HostComponentable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateHostComponentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HostComponentable), nil
}
// Host provides operations to manage the host property of the microsoft.graph.security.hostComponent entity.
// returns a *ThreatintelligenceHostcomponentsItemHostRequestBuilder when successful
func (m *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder) Host()(*ThreatintelligenceHostcomponentsItemHostRequestBuilder) {
    return NewThreatintelligenceHostcomponentsItemHostRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property hostComponents in security
// returns a HostComponentable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HostComponentable, requestConfiguration *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HostComponentable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateHostComponentFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HostComponentable), nil
}
// ToDeleteRequestInformation delete navigation property hostComponents for security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a hostComponent object.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property hostComponents in security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.HostComponentable, requestConfiguration *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder when successful
func (m *ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceHostcomponentsHostComponentItemRequestBuilder) {
    return NewThreatintelligenceHostcomponentsHostComponentItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
