package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder provides operations to manage the complianceManagementPartners property of the microsoft.graph.deviceManagement entity.
type CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// CompliancemanagementpartnersComplianceManagementPartnersRequestBuilderGetQueryParameters list properties and relationships of the complianceManagementPartner objects.
type CompliancemanagementpartnersComplianceManagementPartnersRequestBuilderGetQueryParameters struct {
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
// CompliancemanagementpartnersComplianceManagementPartnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancemanagementpartnersComplianceManagementPartnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilderGetQueryParameters
}
// CompliancemanagementpartnersComplianceManagementPartnersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type CompliancemanagementpartnersComplianceManagementPartnersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByComplianceManagementPartnerId provides operations to manage the complianceManagementPartners property of the microsoft.graph.deviceManagement entity.
// returns a *CompliancemanagementpartnersComplianceManagementPartnerItemRequestBuilder when successful
func (m *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder) ByComplianceManagementPartnerId(complianceManagementPartnerId string)(*CompliancemanagementpartnersComplianceManagementPartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if complianceManagementPartnerId != "" {
        urlTplParams["complianceManagementPartner%2Did"] = complianceManagementPartnerId
    }
    return NewCompliancemanagementpartnersComplianceManagementPartnerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewCompliancemanagementpartnersComplianceManagementPartnersRequestBuilderInternal instantiates a new CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder and sets the default values.
func NewCompliancemanagementpartnersComplianceManagementPartnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder) {
    m := &CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/complianceManagementPartners{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewCompliancemanagementpartnersComplianceManagementPartnersRequestBuilder instantiates a new CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder and sets the default values.
func NewCompliancemanagementpartnersComplianceManagementPartnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewCompliancemanagementpartnersComplianceManagementPartnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *CompliancemanagementpartnersCountRequestBuilder when successful
func (m *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder) Count()(*CompliancemanagementpartnersCountRequestBuilder) {
    return NewCompliancemanagementpartnersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list properties and relationships of the complianceManagementPartner objects.
// returns a ComplianceManagementPartnerCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-onboarding-compliancemanagementpartner-list?view=graph-rest-1.0
func (m *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder) Get(ctx context.Context, requestConfiguration *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ComplianceManagementPartnerCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateComplianceManagementPartnerCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ComplianceManagementPartnerCollectionResponseable), nil
}
// Post create a new complianceManagementPartner object.
// returns a ComplianceManagementPartnerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-onboarding-compliancemanagementpartner-create?view=graph-rest-1.0
func (m *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ComplianceManagementPartnerable, requestConfiguration *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ComplianceManagementPartnerable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateComplianceManagementPartnerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ComplianceManagementPartnerable), nil
}
// ToGetRequestInformation list properties and relationships of the complianceManagementPartner objects.
// returns a *RequestInformation when successful
func (m *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new complianceManagementPartner object.
// returns a *RequestInformation when successful
func (m *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ComplianceManagementPartnerable, requestConfiguration *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder when successful
func (m *CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder) WithUrl(rawUrl string)(*CompliancemanagementpartnersComplianceManagementPartnersRequestBuilder) {
    return NewCompliancemanagementpartnersComplianceManagementPartnersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
