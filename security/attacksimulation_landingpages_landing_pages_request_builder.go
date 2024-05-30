package security

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// AttacksimulationLandingpagesLandingPagesRequestBuilder provides operations to manage the landingPages property of the microsoft.graph.attackSimulationRoot entity.
type AttacksimulationLandingpagesLandingPagesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// AttacksimulationLandingpagesLandingPagesRequestBuilderGetQueryParameters get a list of the landingPage objects and their properties.
type AttacksimulationLandingpagesLandingPagesRequestBuilderGetQueryParameters struct {
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
// AttacksimulationLandingpagesLandingPagesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationLandingpagesLandingPagesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *AttacksimulationLandingpagesLandingPagesRequestBuilderGetQueryParameters
}
// AttacksimulationLandingpagesLandingPagesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type AttacksimulationLandingpagesLandingPagesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByLandingPageId provides operations to manage the landingPages property of the microsoft.graph.attackSimulationRoot entity.
// returns a *AttacksimulationLandingpagesLandingPageItemRequestBuilder when successful
func (m *AttacksimulationLandingpagesLandingPagesRequestBuilder) ByLandingPageId(landingPageId string)(*AttacksimulationLandingpagesLandingPageItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if landingPageId != "" {
        urlTplParams["landingPage%2Did"] = landingPageId
    }
    return NewAttacksimulationLandingpagesLandingPageItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewAttacksimulationLandingpagesLandingPagesRequestBuilderInternal instantiates a new AttacksimulationLandingpagesLandingPagesRequestBuilder and sets the default values.
func NewAttacksimulationLandingpagesLandingPagesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationLandingpagesLandingPagesRequestBuilder) {
    m := &AttacksimulationLandingpagesLandingPagesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/security/attackSimulation/landingPages{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewAttacksimulationLandingpagesLandingPagesRequestBuilder instantiates a new AttacksimulationLandingpagesLandingPagesRequestBuilder and sets the default values.
func NewAttacksimulationLandingpagesLandingPagesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*AttacksimulationLandingpagesLandingPagesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewAttacksimulationLandingpagesLandingPagesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *AttacksimulationLandingpagesCountRequestBuilder when successful
func (m *AttacksimulationLandingpagesLandingPagesRequestBuilder) Count()(*AttacksimulationLandingpagesCountRequestBuilder) {
    return NewAttacksimulationLandingpagesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the landingPage objects and their properties.
// returns a LandingPageCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/attacksimulationroot-list-landingpage?view=graph-rest-1.0
func (m *AttacksimulationLandingpagesLandingPagesRequestBuilder) Get(ctx context.Context, requestConfiguration *AttacksimulationLandingpagesLandingPagesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LandingPageCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLandingPageCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LandingPageCollectionResponseable), nil
}
// Post create new navigation property to landingPages for security
// returns a LandingPageable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *AttacksimulationLandingpagesLandingPagesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LandingPageable, requestConfiguration *AttacksimulationLandingpagesLandingPagesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LandingPageable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateLandingPageFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LandingPageable), nil
}
// ToGetRequestInformation get a list of the landingPage objects and their properties.
// returns a *RequestInformation when successful
func (m *AttacksimulationLandingpagesLandingPagesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *AttacksimulationLandingpagesLandingPagesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to landingPages for security
// returns a *RequestInformation when successful
func (m *AttacksimulationLandingpagesLandingPagesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.LandingPageable, requestConfiguration *AttacksimulationLandingpagesLandingPagesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *AttacksimulationLandingpagesLandingPagesRequestBuilder when successful
func (m *AttacksimulationLandingpagesLandingPagesRequestBuilder) WithUrl(rawUrl string)(*AttacksimulationLandingpagesLandingPagesRequestBuilder) {
    return NewAttacksimulationLandingpagesLandingPagesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
