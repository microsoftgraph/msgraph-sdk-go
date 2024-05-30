package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder provides operations to manage the whoisRecords property of the microsoft.graph.security.threatIntelligence entity.
type ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderGetQueryParameters get a list of whoisRecord objects.
type ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderGetQueryParameters struct {
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
// ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderGetQueryParameters
}
// ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWhoisRecordId provides operations to manage the whoisRecords property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder when successful
func (m *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder) ByWhoisRecordId(whoisRecordId string)(*ThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if whoisRecordId != "" {
        urlTplParams["whoisRecord%2Did"] = whoisRecordId
    }
    return NewThreatintelligenceWhoisrecordsWhoisRecordItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderInternal instantiates a new ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder and sets the default values.
func NewThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder) {
    m := &ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/whoisRecords{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder instantiates a new ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder and sets the default values.
func NewThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ThreatintelligenceWhoisrecordsCountRequestBuilder when successful
func (m *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder) Count()(*ThreatintelligenceWhoisrecordsCountRequestBuilder) {
    return NewThreatintelligenceWhoisrecordsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of whoisRecord objects.
// returns a WhoisRecordCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/security-threatintelligence-list-whoisrecords?view=graph-rest-1.0
func (m *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisRecordCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateWhoisRecordCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisRecordCollectionResponseable), nil
}
// Post create new navigation property to whoisRecords for security
// returns a WhoisRecordable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder) Post(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisRecordable, requestConfiguration *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderPostRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisRecordable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateWhoisRecordFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisRecordable), nil
}
// ToGetRequestInformation get a list of whoisRecord objects.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to whoisRecords for security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder) ToPostRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.WhoisRecordable, requestConfiguration *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder when successful
func (m *ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder) {
    return NewThreatintelligenceWhoisrecordsWhoisRecordsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
