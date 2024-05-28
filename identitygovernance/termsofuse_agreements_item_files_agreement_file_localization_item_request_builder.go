package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder provides operations to manage the files property of the microsoft.graph.agreement entity.
type TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderGetQueryParameters pDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead. Supports $expand.
type TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderGetQueryParameters
}
// TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderInternal instantiates a new TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder and sets the default values.
func NewTermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder) {
    m := &TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/termsOfUse/agreements/{agreement%2Did}/files/{agreementFileLocalization%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder instantiates a new TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder and sets the default values.
func NewTermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property files for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get pDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead. Supports $expand.
// returns a AgreementFileLocalizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
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
// Patch update the navigation property files in identityGovernance
// returns a AgreementFileLocalizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationable, requestConfiguration *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
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
// ToDeleteRequestInformation delete navigation property files for identityGovernance
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation pDFs linked to this agreement. This property is in the process of being deprecated. Use the  file property instead. Supports $expand.
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property files in identityGovernance
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationable, requestConfiguration *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Versions provides operations to manage the versions property of the microsoft.graph.agreementFileLocalization entity.
// returns a *TermsofuseAgreementsItemFilesItemVersionsRequestBuilder when successful
func (m *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder) Versions()(*TermsofuseAgreementsItemFilesItemVersionsRequestBuilder) {
    return NewTermsofuseAgreementsItemFilesItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder when successful
func (m *TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder) WithUrl(rawUrl string)(*TermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder) {
    return NewTermsofuseAgreementsItemFilesAgreementFileLocalizationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
