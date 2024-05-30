package identity

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ConditionalaccessNamedlocationsNamedLocationsRequestBuilder provides operations to manage the namedLocations property of the microsoft.graph.conditionalAccessRoot entity.
type ConditionalaccessNamedlocationsNamedLocationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ConditionalaccessNamedlocationsNamedLocationsRequestBuilderGetQueryParameters get a list of namedLocation objects.
type ConditionalaccessNamedlocationsNamedLocationsRequestBuilderGetQueryParameters struct {
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
// ConditionalaccessNamedlocationsNamedLocationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessNamedlocationsNamedLocationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ConditionalaccessNamedlocationsNamedLocationsRequestBuilderGetQueryParameters
}
// ConditionalaccessNamedlocationsNamedLocationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ConditionalaccessNamedlocationsNamedLocationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByNamedLocationId provides operations to manage the namedLocations property of the microsoft.graph.conditionalAccessRoot entity.
// returns a *ConditionalaccessNamedlocationsNamedLocationItemRequestBuilder when successful
func (m *ConditionalaccessNamedlocationsNamedLocationsRequestBuilder) ByNamedLocationId(namedLocationId string)(*ConditionalaccessNamedlocationsNamedLocationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if namedLocationId != "" {
        urlTplParams["namedLocation%2Did"] = namedLocationId
    }
    return NewConditionalaccessNamedlocationsNamedLocationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewConditionalaccessNamedlocationsNamedLocationsRequestBuilderInternal instantiates a new ConditionalaccessNamedlocationsNamedLocationsRequestBuilder and sets the default values.
func NewConditionalaccessNamedlocationsNamedLocationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessNamedlocationsNamedLocationsRequestBuilder) {
    m := &ConditionalaccessNamedlocationsNamedLocationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identity/conditionalAccess/namedLocations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewConditionalaccessNamedlocationsNamedLocationsRequestBuilder instantiates a new ConditionalaccessNamedlocationsNamedLocationsRequestBuilder and sets the default values.
func NewConditionalaccessNamedlocationsNamedLocationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ConditionalaccessNamedlocationsNamedLocationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewConditionalaccessNamedlocationsNamedLocationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *ConditionalaccessNamedlocationsCountRequestBuilder when successful
func (m *ConditionalaccessNamedlocationsNamedLocationsRequestBuilder) Count()(*ConditionalaccessNamedlocationsCountRequestBuilder) {
    return NewConditionalaccessNamedlocationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of namedLocation objects.
// returns a NamedLocationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conditionalaccessroot-list-namedlocations?view=graph-rest-1.0
func (m *ConditionalaccessNamedlocationsNamedLocationsRequestBuilder) Get(ctx context.Context, requestConfiguration *ConditionalaccessNamedlocationsNamedLocationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NamedLocationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateNamedLocationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NamedLocationCollectionResponseable), nil
}
// Post create a new namedLocation object. Named locations can be either ipNamedLocation or countryNamedLocation objects.
// returns a NamedLocationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/conditionalaccessroot-post-namedlocations?view=graph-rest-1.0
func (m *ConditionalaccessNamedlocationsNamedLocationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NamedLocationable, requestConfiguration *ConditionalaccessNamedlocationsNamedLocationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NamedLocationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateNamedLocationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NamedLocationable), nil
}
// ToGetRequestInformation get a list of namedLocation objects.
// returns a *RequestInformation when successful
func (m *ConditionalaccessNamedlocationsNamedLocationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ConditionalaccessNamedlocationsNamedLocationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new namedLocation object. Named locations can be either ipNamedLocation or countryNamedLocation objects.
// returns a *RequestInformation when successful
func (m *ConditionalaccessNamedlocationsNamedLocationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NamedLocationable, requestConfiguration *ConditionalaccessNamedlocationsNamedLocationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ConditionalaccessNamedlocationsNamedLocationsRequestBuilder when successful
func (m *ConditionalaccessNamedlocationsNamedLocationsRequestBuilder) WithUrl(rawUrl string)(*ConditionalaccessNamedlocationsNamedLocationsRequestBuilder) {
    return NewConditionalaccessNamedlocationsNamedLocationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
