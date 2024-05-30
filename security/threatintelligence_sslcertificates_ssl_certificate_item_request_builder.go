package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder provides operations to manage the sslCertificates property of the microsoft.graph.security.threatIntelligence entity.
type ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderGetQueryParameters get the properties and relationships of an sslCertificate object.
type ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderGetQueryParameters
}
// ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderInternal instantiates a new ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder and sets the default values.
func NewThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder) {
    m := &ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/sslCertificates/{sslCertificate%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder instantiates a new ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder and sets the default values.
func NewThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property sslCertificates for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get the properties and relationships of an sslCertificate object.
// returns a SslCertificateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-sslcertificate-get?view=graph-rest-1.0
func (m *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.SslCertificateable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateSslCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.SslCertificateable), nil
}
// Patch update the navigation property sslCertificates in security
// returns a SslCertificateable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.SslCertificateable, requestConfiguration *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.SslCertificateable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateSslCertificateFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.SslCertificateable), nil
}
// RelatedHosts provides operations to manage the relatedHosts property of the microsoft.graph.security.sslCertificate entity.
// returns a *ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder when successful
func (m *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder) RelatedHosts()(*ThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilder) {
    return NewThreatintelligenceSslcertificatesItemRelatedhostsRelatedHostsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property sslCertificates for security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get the properties and relationships of an sslCertificate object.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property sslCertificates in security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.SslCertificateable, requestConfiguration *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder when successful
func (m *ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder) {
    return NewThreatintelligenceSslcertificatesSslCertificateItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
