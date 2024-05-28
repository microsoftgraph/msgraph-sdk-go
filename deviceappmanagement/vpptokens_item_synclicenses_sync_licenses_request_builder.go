package deviceappmanagement

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// VpptokensItemSynclicensesSyncLicensesRequestBuilder provides operations to call the syncLicenses method.
type VpptokensItemSynclicensesSyncLicensesRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// VpptokensItemSynclicensesSyncLicensesRequestBuilderPostRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type VpptokensItemSynclicensesSyncLicensesRequestBuilderPostRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// NewVpptokensItemSynclicensesSyncLicensesRequestBuilderInternal instantiates a new VpptokensItemSynclicensesSyncLicensesRequestBuilder and sets the default values.
func NewVpptokensItemSynclicensesSyncLicensesRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VpptokensItemSynclicensesSyncLicensesRequestBuilder) {
    m := &VpptokensItemSynclicensesSyncLicensesRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/deviceAppManagement/vppTokens/{vppToken%2Did}/syncLicenses", pathParameters),
    }
    return m
}
// NewVpptokensItemSynclicensesSyncLicensesRequestBuilder instantiates a new VpptokensItemSynclicensesSyncLicensesRequestBuilder and sets the default values.
func NewVpptokensItemSynclicensesSyncLicensesRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*VpptokensItemSynclicensesSyncLicensesRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewVpptokensItemSynclicensesSyncLicensesRequestBuilderInternal(urlParams, requestAdapter)
}
// Post syncs licenses associated with a specific appleVolumePurchaseProgramToken
// returns a VppTokenable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/intune-onboarding-vpptoken-synclicenses?view=graph-rest-1.0
func (m *VpptokensItemSynclicensesSyncLicensesRequestBuilder) Post(ctx context.Context, requestConfiguration *VpptokensItemSynclicensesSyncLicensesRequestBuilderPostRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VppTokenable, error) {
    requestInfo, err := m.ToPostRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateVppTokenFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.VppTokenable), nil
}
// ToPostRequestInformation syncs licenses associated with a specific appleVolumePurchaseProgramToken
// returns a *RequestInformation when successful
func (m *VpptokensItemSynclicensesSyncLicensesRequestBuilder) ToPostRequestInformation(ctx context.Context, requestConfiguration *VpptokensItemSynclicensesSyncLicensesRequestBuilderPostRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.POST, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *VpptokensItemSynclicensesSyncLicensesRequestBuilder when successful
func (m *VpptokensItemSynclicensesSyncLicensesRequestBuilder) WithUrl(rawUrl string)(*VpptokensItemSynclicensesSyncLicensesRequestBuilder) {
    return NewVpptokensItemSynclicensesSyncLicensesRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
