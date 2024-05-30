package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TermsofuseAgreementsItemFilesRequestBuilder provides operations to manage the files property of the microsoft.graph.agreement entity.
type TermsofuseAgreementsItemFilesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TermsofuseAgreementsItemFilesRequestBuilderGetQueryParameters pDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead. Supports $expand.
type TermsofuseAgreementsItemFilesRequestBuilderGetQueryParameters struct {
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
// TermsofuseAgreementsItemFilesRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementsItemFilesRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermsofuseAgreementsItemFilesRequestBuilderGetQueryParameters
}
// TermsofuseAgreementsItemFilesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementsItemFilesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ByAgreementFileLocalizationId provides operations to manage the files property of the microsoft.graph.agreement entity.
// returns a *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder when successful
func (m *TermsofuseAgreementsItemFilesRequestBuilder) ByAgreementFileLocalizationId(agreementFileLocalizationId string)(*TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.BaseRequestBuilder.PathParameters {
        urlTplParams[idx] = item
    }
    if agreementFileLocalizationId != "" {
        urlTplParams["agreementFileLocalization%2Did"] = agreementFileLocalizationId
    }
    return NewTermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderInternal(urlTplParams, m.BaseRequestBuilder.RequestAdapter)
}
// NewTermsofuseAgreementsItemFilesRequestBuilderInternal instantiates a new TermsofuseAgreementsItemFilesRequestBuilder and sets the default values.
func NewTermsofuseAgreementsItemFilesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsofuseAgreementsItemFilesRequestBuilder) {
    m := &TermsofuseAgreementsItemFilesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/termsOfUse/agreements/{agreement%2Did}/files{?%24count,%24expand,%24filter,%24orderby,%24search,%24select,%24skip,%24top}", pathParameters),
    }
    return m
}
// NewTermsofuseAgreementsItemFilesRequestBuilder instantiates a new TermsofuseAgreementsItemFilesRequestBuilder and sets the default values.
func NewTermsofuseAgreementsItemFilesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsofuseAgreementsItemFilesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsofuseAgreementsItemFilesRequestBuilderInternal(urlParams, requestAdapter)
}
// Count provides operations to count the resources in the collection.
// returns a *TermsofuseAgreementsItemFilesCountRequestBuilder when successful
func (m *TermsofuseAgreementsItemFilesRequestBuilder) Count()(*TermsofuseAgreementsItemFilesCountRequestBuilder) {
    return NewTermsofuseAgreementsItemFilesCountRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get pDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead. Supports $expand.
// returns a AgreementFileLocalizationCollectionResponseable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementsItemFilesRequestBuilder) Get(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFilesRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationCollectionResponseable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAgreementFileLocalizationCollectionResponseFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationCollectionResponseable), nil
}
// Post create a new localized agreement file.
// returns a AgreementFileLocalizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/agreement-post-files?view=graph-rest-1.0
func (m *TermsofuseAgreementsItemFilesRequestBuilder) Post(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationable, requestConfiguration *TermsofuseAgreementsItemFilesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAgreementFileLocalizationFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationable), nil
}
// ToGetRequestInformation pDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead. Supports $expand.
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementsItemFilesRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFilesRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPostRequestInformation create a new localized agreement file.
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementsItemFilesRequestBuilder) ToPostRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationable, requestConfiguration *TermsofuseAgreementsItemFilesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TermsofuseAgreementsItemFilesRequestBuilder when successful
func (m *TermsofuseAgreementsItemFilesRequestBuilder) WithUrl(rawUrl string)(*TermsofuseAgreementsItemFilesRequestBuilder) {
    return NewTermsofuseAgreementsItemFilesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
