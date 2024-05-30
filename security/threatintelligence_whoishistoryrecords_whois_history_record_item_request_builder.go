package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder provides operations to manage the whoisHistoryRecords property of the microsoft.graph.security.threatIntelligence entity.
type ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderGetQueryParameters retrieve details about whoisHistoryRecord objects.Note: List retrieval is not yet supported.
type ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderGetQueryParameters
}
// ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderInternal instantiates a new ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder and sets the default values.
func NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder) {
    m := &ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/whoisHistoryRecords/{whoisHistoryRecord%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder instantiates a new ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder and sets the default values.
func NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property whoisHistoryRecords for security
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get retrieve details about whoisHistoryRecord objects.Note: List retrieval is not yet supported.
// returns a WhoisHistoryRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisHistoryRecordable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateWhoisHistoryRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisHistoryRecordable), nil
}
// Host provides operations to manage the host property of the microsoft.graph.security.whoisBaseRecord entity.
// returns a *ThreatintelligenceWhoishistoryrecordsItemHostRequestBuilder when successful
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder) Host()(*ThreatintelligenceWhoishistoryrecordsItemHostRequestBuilder) {
    return NewThreatintelligenceWhoishistoryrecordsItemHostRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property whoisHistoryRecords in security
// returns a WhoisHistoryRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder) Patch(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisHistoryRecordable, requestConfiguration *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderPatchRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisHistoryRecordable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateWhoisHistoryRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisHistoryRecordable), nil
}
// ToDeleteRequestInformation delete navigation property whoisHistoryRecords for security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation retrieve details about whoisHistoryRecord objects.Note: List retrieval is not yet supported.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property whoisHistoryRecords in security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisHistoryRecordable, requestConfiguration *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder when successful
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder) {
    return NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
