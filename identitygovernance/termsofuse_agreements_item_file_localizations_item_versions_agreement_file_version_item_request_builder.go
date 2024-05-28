package identitygovernance

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder provides operations to manage the versions property of the microsoft.graph.agreementFileLocalization entity.
type TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetQueryParameters read-only. Customized versions of the terms of use agreement in the Microsoft Entra tenant.
type TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetQueryParameters
}
// TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewTermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderInternal instantiates a new TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder and sets the default values.
func NewTermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) {
    m := &TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/identityGovernance/termsOfUse/agreements/{agreement%2Did}/file/localizations/{agreementFileLocalization%2Did}/versions/{agreementFileVersion%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewTermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder instantiates a new TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder and sets the default values.
func NewTermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewTermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete delete navigation property versions for identityGovernance
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get read-only. Customized versions of the terms of use agreement in the Microsoft Entra tenant.
// returns a AgreementFileVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) Get(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileVersionable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAgreementFileVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileVersionable), nil
}
// Patch update the navigation property versions in identityGovernance
// returns a AgreementFileVersionable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileVersionable, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileVersionable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAgreementFileVersionFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileVersionable), nil
}
// ToDeleteRequestInformation delete navigation property versions for identityGovernance
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation read-only. Customized versions of the terms of use agreement in the Microsoft Entra tenant.
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property versions in identityGovernance
// returns a *RequestInformation when successful
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AgreementFileVersionable, requestConfiguration *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder when successful
func (m *TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) WithUrl(rawUrl string)(*TermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder) {
    return NewTermsofuseAgreementsItemFileLocalizationsItemVersionsAgreementFileVersionItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
