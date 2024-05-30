package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae "github.com/microsoftgraph/msgraph-sdk-go/models/security"
)

// ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder provides operations to manage the intelligenceProfileIndicators property of the microsoft.graph.security.threatIntelligence entity.
type ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderGetQueryParameters read the properties and relationships of a intelligenceProfileIndicator object.
type ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderGetQueryParameters struct {
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
// ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderGetQueryParameters
}
// ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByIntelligenceProfileIndicatorId provides operations to manage the intelligenceProfileIndicators property of the microsoft.graph.security.threatIntelligence entity.
// returns a *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder when successful
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder) ByIntelligenceProfileIndicatorId(intelligenceProfileIndicatorId string)(*ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if intelligenceProfileIndicatorId != "" {
        urlTplParams["intelligenceProfileIndicator%2Did"] = intelligenceProfileIndicatorId
    }
    return NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderInternal instantiates a new ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder and sets the default values.
func NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder) {
    m := &ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/threatIntelligence/intelligenceProfileIndicators{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder instantiates a new ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder and sets the default values.
func NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ThreatintelligenceIntelligenceprofileindicatorsCountRequestBuilder when successful
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder) Count()(*ThreatintelligenceIntelligenceprofileindicatorsCountRequestBuilder) {
    return NewThreatintelligenceIntelligenceprofileindicatorsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get read the properties and relationships of a intelligenceProfileIndicator object.
// returns a IntelligenceProfileIndicatorCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder) Get(ctx context.Context, requestConfiguration *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderGetRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.IntelligenceProfileIndicatorCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateIntelligenceProfileIndicatorCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.IntelligenceProfileIndicatorCollectionResponseable), nil
}
// Post create new navigation property to intelligenceProfileIndicators for security
// returns a IntelligenceProfileIndicatorable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder) Post(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.IntelligenceProfileIndicatorable, requestConfiguration *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderPostRequestConfiguration)(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.IntelligenceProfileIndicatorable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.CreateIntelligenceProfileIndicatorFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.IntelligenceProfileIndicatorable), nil
}
// ToGetRequestInformation read the properties and relationships of a intelligenceProfileIndicator object.
// returns a *RequestInformation when successful
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to intelligenceProfileIndicators for security
// returns a *RequestInformation when successful
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder) ToPostRequestInformation(ctx context.Context, body idd6d442c3cc83a389b8f0b8dd7ac355916e813c2882ff3aaa23331424ba827ae.IntelligenceProfileIndicatorable, requestConfiguration *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder when successful
func (m *ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder) WithUrl(rawUrl string)(*ThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder) {
    return NewThreatintelligenceIntelligenceprofileindicatorsIntelligenceProfileIndicatorsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
