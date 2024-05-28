package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder provides operations to call the getManagedAppDiagnosticStatuses method.
type ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilderGetQueryParameters gets diagnostics validation status for a given user.
type ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilderGetQueryParameters struct {
    // Include count of items
    Count *bool `uriparametername:"%24count"`
    // Filter items by property values
    Filter *string `uriparametername:"%24filter"`
    // Search items by search phrases
    Search *string `uriparametername:"%24search"`
    // Skip the first n items
    Skip *int32 `uriparametername:"%24skip"`
    // Show only the first n items
    Top *int32 `uriparametername:"%24top"`
}
// ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilderGetQueryParameters
}
// NewItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilderInternal instantiates a new ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder and sets the default values.
func NewItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder) {
    m := &ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/getManagedAppDiagnosticStatuses(){?%24count,%24filter,%24search,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder instantiates a new ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder and sets the default values.
func NewItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilderInternal(urlParams, requestAdapter)
}
// Get gets diagnostics validation status for a given user.
// Deprecated: This method is obsolete. Use GetAsGetManagedAppDiagnosticStatusesGetResponse instead.
// returns a ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-user-getmanagedappdiagnosticstatuses?view=graph-rest-1.0
func (m *ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilderGetRequestConfiguration)(ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesResponseable), nil
}
// GetAsGetManagedAppDiagnosticStatusesGetResponse gets diagnostics validation status for a given user.
// returns a ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesGetResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-mam-user-getmanagedappdiagnosticstatuses?view=graph-rest-1.0
func (m *ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder) GetAsGetManagedAppDiagnosticStatusesGetResponse(ctx context.Context, requestConfiguration *ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilderGetRequestConfiguration)(ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesGetResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, CreateItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesGetResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesGetResponseable), nil
}
// ToGetRequestInformation gets diagnostics validation status for a given user.
// returns a *RequestInformation when successful
func (m *ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder when successful
func (m *ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder) WithUrl(rawUrl string)(*ItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder) {
    return NewItemGetmanagedappdiagnosticstatusesGetManagedAppDiagnosticStatusesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
