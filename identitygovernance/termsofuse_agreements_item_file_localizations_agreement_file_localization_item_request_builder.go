package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder provides operations to manage the localizations property of the microsoft.graph.agreementFile entity.
type TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetQueryParameters the localized version of the terms of use agreement files attached to the agreement.
type TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetQueryParameters
}
// TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderInternal instantiates a new TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder and sets the default values.
func NewTermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) {
    m := &TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/termsOfUse/agreements/{agreement%2Did}/file/localizations/{agreementFileLocalization%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder instantiates a new TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder and sets the default values.
func NewTermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property localizations for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the localized version of the terms of use agreement files attached to the agreement.
// returns a AgreementFileLocalizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationable, error) {
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
// Patch update the navigation property localizations in identityGovernance
// returns a AgreementFileLocalizationable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationable, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationable, error) {
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
// ToDeleteRequestInformation delete navigation property localizations for identityGovernance
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the localized version of the terms of use agreement files attached to the agreement.
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property localizations in identityGovernance
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileLocalizationable, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder when successful
func (m *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) Versions()(*TermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilder) {
    return NewTermsofuseAgreementsItemFileLocalizationsItemVersionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder when successful
func (m *TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) WithUrl(rawUrl string)(*TermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder) {
    return NewTermsofuseAgreementsItemFileLocalizationsAgreementFileLocalizationItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
