package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder provides operations to manage the delegatedAdminCustomers property of the microsoft.graph.tenantRelationship entity.
type DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderGetQueryParameters read the properties of a delegatedAdminCustomer object.
type DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderGetQueryParameters
}
// DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewDelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderInternal instantiates a new DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder and sets the default values.
func NewDelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder) {
    m := &DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/delegatedAdminCustomers/{delegatedAdminCustomer%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewDelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder instantiates a new DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder and sets the default values.
func NewDelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property delegatedAdminCustomers for tenantRelationships
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read the properties of a delegatedAdminCustomer object.
// returns a DelegatedAdminCustomerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadmincustomer-get?view=graph-rest-1.0
func (m *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminCustomerable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminCustomerable), nil
}
// Patch update the navigation property delegatedAdminCustomers in tenantRelationships
// returns a DelegatedAdminCustomerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminCustomerable, requestConfiguration *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminCustomerable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminCustomerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminCustomerable), nil
}
// ServiceManagementDetails provides operations to manage the serviceManagementDetails property of the microsoft.graph.delegatedAdminCustomer entity.
// returns a *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder when successful
func (m *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder) ServiceManagementDetails()(*DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder) {
    return NewDelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property delegatedAdminCustomers for tenantRelationships
// returns a *RequestInformation when successful
func (m *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read the properties of a delegatedAdminCustomer object.
// returns a *RequestInformation when successful
func (m *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property delegatedAdminCustomers in tenantRelationships
// returns a *RequestInformation when successful
func (m *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminCustomerable, requestConfiguration *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder when successful
func (m *DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder) WithUrl(rawUrl string)(*DelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder) {
    return NewDelegatedadmincustomersDelegatedAdminCustomerItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
