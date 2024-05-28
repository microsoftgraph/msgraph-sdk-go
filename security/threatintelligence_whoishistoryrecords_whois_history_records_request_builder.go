package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder provides operations to manage the whoisHistoryRecords property of the microsoft.graph.security.threatIntelligence entity.
type ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderGetQueryParameters retrieve details about whoisHistoryRecord objects.Note: List retrieval is not yet supported.
type ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Order items by property values
    Orderby []string `uriparametername:"%24orderby"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderGetQueryParameters
}
// ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWhoisHistoryRecordId provides operations to manage the whoisHistoryRecords property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder when successful
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder) ByWhoisHistoryRecordId(whoisHistoryRecordId string)(*ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if whoisHistoryRecordId != "" {
        urlTplParams["whoisHistoryRecord%2Did"] = whoisHistoryRecordId
    }
    return NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderInternal instantiates a new ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder and sets the default values.
func NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder) {
    m := &ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/whoisHistoryRecords{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder instantiates a new ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder and sets the default values.
func NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ThreatintelligenceWhoishistoryrecordsCountRequestBuilder when successful
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder) Count()(*ThreatintelligenceWhoishistoryrecordsCountRequestBuilder) {
    return NewThreatintelligenceWhoishistoryrecordsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve details about whoisHistoryRecord objects.Note: List retrieval is not yet supported.
// returns a WhoisHistoryRecordCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisHistoryRecordCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateWhoisHistoryRecordCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisHistoryRecordCollectionResponseable), nil
}
// Post create new navigation property to whoisHistoryRecords for security
// returns a WhoisHistoryRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder) Post(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisHistoryRecordable, requestConfiguration *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderPostRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisHistoryRecordable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
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
// ToGetRequestInformation retrieve details about whoisHistoryRecord objects.Note: List retrieval is not yet supported.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to whoisHistoryRecords for security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder) ToPostRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisHistoryRecordable, requestConfiguration *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
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
// returns a *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder when successful
func (m *ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder) {
    return NewThreatintelligenceWhoishistoryrecordsWhoisHistoryRecordsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
