package directory

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder provides operations to manage the subscriptions property of the microsoft.graph.directory entity.
type SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderGetQueryParameters get subscriptions from directory
type SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderGetQueryParameters
}
// SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewSubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderInternal instantiates a new SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder and sets the default values.
func NewSubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, commerceSubscriptionId *string)(*SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder) {
    m := &SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/directory/subscriptions(commerceSubscriptionId='{commerceSubscriptionId}'){?%24expand,%24select}", pathParameters),
    }
    if commerceSubscriptionId != nil {
        m.BaseRequestBuilder.PathParameters["commerceSubscriptionId"] = *commerceSubscriptionId
    }
    return m
}
// NewSubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder instantiates a new SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder and sets the default values.
func NewSubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderInternal(urlParams, requestAdapter, nil)
}
// Delete delete navigation property subscriptions for directory
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder) Delete(ctx context.Context, requestConfiguration *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get get subscriptions from directory
// returns a CompanySubscriptionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder) Get(ctx context.Context, requestConfiguration *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CompanySubscriptionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCompanySubscriptionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CompanySubscriptionable), nil
}
// Patch update the navigation property subscriptions in directory
// returns a CompanySubscriptionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CompanySubscriptionable, requestConfiguration *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CompanySubscriptionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateCompanySubscriptionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CompanySubscriptionable), nil
}
// ToDeleteRequestInformation delete navigation property subscriptions for directory
// returns a *RequestInformation when successful
func (m *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation get subscriptions from directory
// returns a *RequestInformation when successful
func (m *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property subscriptions in directory
// returns a *RequestInformation when successful
func (m *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CompanySubscriptionable, requestConfiguration *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder when successful
func (m *SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder) WithUrl(rawUrl string)(*SubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder) {
    return NewSubscriptionswithcommercesubscriptionidSubscriptionsWithCommerceSubscriptionIdRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
