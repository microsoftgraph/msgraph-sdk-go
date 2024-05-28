package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder provides operations to manage the agreementAcceptances property of the microsoft.graph.termsOfUseContainer entity.
type TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderGetQueryParameters represents the current status of a user's response to a company's customizable terms of use agreement.
type TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderGetQueryParameters struct {
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
// TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderGetQueryParameters
}
// TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAgreementAcceptanceId provides operations to manage the agreementAcceptances property of the microsoft.graph.termsOfUseContainer entity.
// returns a *TermsofuseAgreementacceptancesAgreementAcceptanceItemRequestBuilder when successful
func (m *TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder) ByAgreementAcceptanceId(agreementAcceptanceId string)(*TermsofuseAgreementacceptancesAgreementAcceptanceItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if agreementAcceptanceId != "" {
        urlTplParams["agreementAcceptance%2Did"] = agreementAcceptanceId
    }
    return NewTermsofuseAgreementacceptancesAgreementAcceptanceItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderInternal instantiates a new TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder and sets the default values.
func NewTermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder) {
    m := &TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/termsOfUse/agreementAcceptances{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder instantiates a new TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder and sets the default values.
func NewTermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TermsofuseAgreementacceptancesCountRequestBuilder when successful
func (m *TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder) Count()(*TermsofuseAgreementacceptancesCountRequestBuilder) {
    return NewTermsofuseAgreementacceptancesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get represents the current status of a user's response to a company's customizable terms of use agreement.
// returns a AgreementAcceptanceCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder) Get(ctx context.Context, requestConfiguration *TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementAcceptanceCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAgreementAcceptanceCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementAcceptanceCollectionResponseable), nil
}
// Post create new navigation property to agreementAcceptances for identityGovernance
// returns a AgreementAcceptanceable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementAcceptanceable, requestConfiguration *TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementAcceptanceable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAgreementAcceptanceFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementAcceptanceable), nil
}
// ToGetRequestInformation represents the current status of a user's response to a company's customizable terms of use agreement.
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create new navigation property to agreementAcceptances for identityGovernance
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementAcceptanceable, requestConfiguration *TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder when successful
func (m *TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder) WithUrl(rawUrl string)(*TermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder) {
    return NewTermsofuseAgreementacceptancesAgreementAcceptancesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
