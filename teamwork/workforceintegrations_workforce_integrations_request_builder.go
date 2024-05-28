package teamwork

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// WorkforceintegrationsWorkforceIntegrationsRequestBuilder provides operations to manage the workforceIntegrations property of the microsoft.graph.teamwork entity.
type WorkforceintegrationsWorkforceIntegrationsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// WorkforceintegrationsWorkforceIntegrationsRequestBuilderGetQueryParameters retrieve a list of workforceIntegration objects.
type WorkforceintegrationsWorkforceIntegrationsRequestBuilderGetQueryParameters struct {
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
// WorkforceintegrationsWorkforceIntegrationsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WorkforceintegrationsWorkforceIntegrationsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *WorkforceintegrationsWorkforceIntegrationsRequestBuilderGetQueryParameters
}
// WorkforceintegrationsWorkforceIntegrationsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type WorkforceintegrationsWorkforceIntegrationsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByWorkforceIntegrationId provides operations to manage the workforceIntegrations property of the microsoft.graph.teamwork entity.
// returns a *WorkforceintegrationsWorkforceIntegrationItemRequestBuilder when successful
func (m *WorkforceintegrationsWorkforceIntegrationsRequestBuilder) ByWorkforceIntegrationId(workforceIntegrationId string)(*WorkforceintegrationsWorkforceIntegrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if workforceIntegrationId != "" {
        urlTplParams["workforceIntegration%2Did"] = workforceIntegrationId
    }
    return NewWorkforceintegrationsWorkforceIntegrationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewWorkforceintegrationsWorkforceIntegrationsRequestBuilderInternal instantiates a new WorkforceintegrationsWorkforceIntegrationsRequestBuilder and sets the default values.
func NewWorkforceintegrationsWorkforceIntegrationsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WorkforceintegrationsWorkforceIntegrationsRequestBuilder) {
    m := &WorkforceintegrationsWorkforceIntegrationsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/teamwork/workforceIntegrations{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewWorkforceintegrationsWorkforceIntegrationsRequestBuilder instantiates a new WorkforceintegrationsWorkforceIntegrationsRequestBuilder and sets the default values.
func NewWorkforceintegrationsWorkforceIntegrationsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*WorkforceintegrationsWorkforceIntegrationsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewWorkforceintegrationsWorkforceIntegrationsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *WorkforceintegrationsCountRequestBuilder when successful
func (m *WorkforceintegrationsWorkforceIntegrationsRequestBuilder) Count()(*WorkforceintegrationsCountRequestBuilder) {
    return NewWorkforceintegrationsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get retrieve a list of workforceIntegration objects.
// returns a WorkforceIntegrationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/workforceintegration-list?view=graph-rest-1.0
func (m *WorkforceintegrationsWorkforceIntegrationsRequestBuilder) Get(ctx context.Context, requestConfiguration *WorkforceintegrationsWorkforceIntegrationsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkforceIntegrationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkforceIntegrationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkforceIntegrationCollectionResponseable), nil
}
// Post create a new workforceIntegration object.You can set up which entities you want to receive Shifts synchronous change notifications on and set entities to configure filtering by WFM rules eligibility for, including swap requests.
// returns a WorkforceIntegrationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/workforceintegration-post?view=graph-rest-1.0
func (m *WorkforceintegrationsWorkforceIntegrationsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkforceIntegrationable, requestConfiguration *WorkforceintegrationsWorkforceIntegrationsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkforceIntegrationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateWorkforceIntegrationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkforceIntegrationable), nil
}
// ToGetRequestInformation retrieve a list of workforceIntegration objects.
// returns a *RequestInformation when successful
func (m *WorkforceintegrationsWorkforceIntegrationsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *WorkforceintegrationsWorkforceIntegrationsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new workforceIntegration object.You can set up which entities you want to receive Shifts synchronous change notifications on and set entities to configure filtering by WFM rules eligibility for, including swap requests.
// returns a *RequestInformation when successful
func (m *WorkforceintegrationsWorkforceIntegrationsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.WorkforceIntegrationable, requestConfiguration *WorkforceintegrationsWorkforceIntegrationsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *WorkforceintegrationsWorkforceIntegrationsRequestBuilder when successful
func (m *WorkforceintegrationsWorkforceIntegrationsRequestBuilder) WithUrl(rawUrl string)(*WorkforceintegrationsWorkforceIntegrationsRequestBuilder) {
    return NewWorkforceintegrationsWorkforceIntegrationsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
