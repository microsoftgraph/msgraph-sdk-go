package devicemanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder provides operations to manage the telecomExpenseManagementPartners property of the microsoft.graph.deviceManagement entity.
type TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetQueryParameters list properties and relationships of the telecomExpenseManagementPartner objects.
type TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetQueryParameters struct {
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
// TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetQueryParameters
}
// TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByTelecomExpenseManagementPartnerId provides operations to manage the telecomExpenseManagementPartners property of the microsoft.graph.deviceManagement entity.
// returns a *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnerItemRequestBuilder when successful
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) ByTelecomExpenseManagementPartnerId(telecomExpenseManagementPartnerId string)(*TelecomexpensemanagementpartnersTelecomExpenseManagementPartnerItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if telecomExpenseManagementPartnerId != "" {
        urlTplParams["telecomExpenseManagementPartner%2Did"] = telecomExpenseManagementPartnerId
    }
    return NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnerItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderInternal instantiates a new TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder and sets the default values.
func NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) {
    m := &TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceManagement/telecomExpenseManagementPartners{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder instantiates a new TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder and sets the default values.
func NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TelecomexpensemanagementpartnersCountRequestBuilder when successful
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) Count()(*TelecomexpensemanagementpartnersCountRequestBuilder) {
    return NewTelecomexpensemanagementpartnersCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get list properties and relationships of the telecomExpenseManagementPartner objects.
// returns a TelecomExpenseManagementPartnerCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-tem-telecomexpensemanagementpartner-list?view=graph-rest-1.0
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) Get(ctx context.Context, requestConfiguration *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TelecomExpenseManagementPartnerCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTelecomExpenseManagementPartnerCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TelecomExpenseManagementPartnerCollectionResponseable), nil
}
// Post create a new telecomExpenseManagementPartner object.
// returns a TelecomExpenseManagementPartnerable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-tem-telecomexpensemanagementpartner-create?view=graph-rest-1.0
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TelecomExpenseManagementPartnerable, requestConfiguration *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TelecomExpenseManagementPartnerable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateTelecomExpenseManagementPartnerFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TelecomExpenseManagementPartnerable), nil
}
// ToGetRequestInformation list properties and relationships of the telecomExpenseManagementPartner objects.
// returns a *RequestInformation when successful
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new telecomExpenseManagementPartner object.
// returns a *RequestInformation when successful
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.TelecomExpenseManagementPartnerable, requestConfiguration *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder when successful
func (m *TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) WithUrl(rawUrl string)(*TelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder) {
    return NewTelecomexpensemanagementpartnersTelecomExpenseManagementPartnersRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
