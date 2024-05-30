package tenantrelationships

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder provides operations to manage the serviceManagementDetails property of the microsoft.graph.delegatedAdminCustomer entity.
type DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderGetQueryParameters get a list of the delegatedAdminServiceManagementDetail objects and their properties.
type DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderGetQueryParameters struct {
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
// DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderGetQueryParameters
}
// DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByDelegatedAdminServiceManagementDetailId provides operations to manage the serviceManagementDetails property of the microsoft.graph.delegatedAdminCustomer entity.
// returns a *DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder when successful
func (m *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder) ByDelegatedAdminServiceManagementDetailId(delegatedAdminServiceManagementDetailId string)(*DelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if delegatedAdminServiceManagementDetailId != "" {
        urlTplParams["delegatedAdminServiceManagementDetail%2Did"] = delegatedAdminServiceManagementDetailId
    }
    return NewDelegatedadmincustomersItemServicemanagementdetailsDelegatedAdminServiceManagementDetailItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewDelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderInternal instantiates a new DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder and sets the default values.
func NewDelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder) {
    m := &DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/tenantRelationships/delegatedAdminCustomers/{delegatedAdminCustomer%2Did}/serviceManagementDetails{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewDelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder instantiates a new DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder and sets the default values.
func NewDelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *DelegatedadmincustomersItemServicemanagementdetailsCountRequestBuilder when successful
func (m *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder) Count()(*DelegatedadmincustomersItemServicemanagementdetailsCountRequestBuilder) {
    return NewDelegatedadmincustomersItemServicemanagementdetailsCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get get a list of the delegatedAdminServiceManagementDetail objects and their properties.
// returns a DelegatedAdminServiceManagementDetailCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/delegatedadmincustomer-list-servicemanagementdetails?view=graph-rest-1.0
func (m *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder) Get(ctx context.Context, requestConfiguration *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminServiceManagementDetailCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminServiceManagementDetailCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminServiceManagementDetailCollectionResponseable), nil
}
// Post create new navigation property to serviceManagementDetails for tenantRelationships
// returns a DelegatedAdminServiceManagementDetailable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminServiceManagementDetailable, requestConfiguration *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminServiceManagementDetailable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDelegatedAdminServiceManagementDetailFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminServiceManagementDetailable), nil
}
// ToGetRequestInformation get a list of the delegatedAdminServiceManagementDetail objects and their properties.
// returns a *RequestInformation when successful
func (m *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to serviceManagementDetails for tenantRelationships
// returns a *RequestInformation when successful
func (m *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DelegatedAdminServiceManagementDetailable, requestConfiguration *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder when successful
func (m *DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder) WithUrl(rawUrl string)(*DelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder) {
    return NewDelegatedadmincustomersItemServicemanagementdetailsServiceManagementDetailsRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
