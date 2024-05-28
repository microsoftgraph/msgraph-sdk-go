package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder provides operations to manage the passiveDnsRecords property of the microsoft.graph.security.threatIntelligence entity.
type ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderGetQueryParameters read the properties and relationships of a passiveDnsRecord object.
type ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderGetQueryParameters
}
// ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Artifact provides operations to manage the artifact property of the microsoft.graph.security.passiveDnsRecord entity.
// returns a *ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder when successful
func (m *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder) Artifact()(*ThreatintelligencePassivednsrecordsItemArtifactRequestBuilder) {
    return NewThreatintelligencePassivednsrecordsItemArtifactRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderInternal instantiates a new ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder and sets the default values.
func NewThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder) {
    m := &ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/passiveDnsRecords/{passiveDnsRecord%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder instantiates a new ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder and sets the default values.
func NewThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property passiveDnsRecords for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties and relationships of a passiveDnsRecord object.
// returns a PassiveDnsRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-passivednsrecord-get?view=graph-rest-1.0
func (m *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PassiveDnsRecordable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreatePassiveDnsRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PassiveDnsRecordable), nil
}
// ParentHost provides operations to manage the parentHost property of the microsoft.graph.security.passiveDnsRecord entity.
// returns a *ThreatintelligencePassivednsrecordsItemParenthostParentHostRequestBuilder when successful
func (m *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder) ParentHost()(*ThreatintelligencePassivednsrecordsItemParenthostParentHostRequestBuilder) {
    return NewThreatintelligencePassivednsrecordsItemParenthostParentHostRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property passiveDnsRecords in security
// returns a PassiveDnsRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PassiveDnsRecordable, requestConfiguration *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PassiveDnsRecordable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreatePassiveDnsRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PassiveDnsRecordable), nil
}
// ToDeleteRequestInformation delete navigation property passiveDnsRecords for security
// returns a *RequestInformation when successful
func (m *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties and relationships of a passiveDnsRecord object.
// returns a *RequestInformation when successful
func (m *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property passiveDnsRecords in security
// returns a *RequestInformation when successful
func (m *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.PassiveDnsRecordable, requestConfiguration *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder when successful
func (m *ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder) {
    return NewThreatintelligencePassivednsrecordsPassiveDnsRecordItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
