// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package solutions

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder provides operations to manage the siteProtectionUnitsBulkAdditionJobs property of the microsoft.graph.sharePointProtectionPolicy entity.
type BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetQueryParameters get a siteProtectionUnitsBulkAdditionJob object by the ID associated with a sharePointProtectionPolicy.
type BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetQueryParameters
}
// NewBackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderInternal instantiates a new BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder and sets the default values.
func NewBackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) {
    m := &BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/solutions/backupRestore/sharePointProtectionPolicies/{sharePointProtectionPolicy%2Did}/siteProtectionUnitsBulkAdditionJobs/{siteProtectionUnitsBulkAdditionJob%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewBackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder instantiates a new BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder and sets the default values.
func NewBackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewBackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get a siteProtectionUnitsBulkAdditionJob object by the ID associated with a sharePointProtectionPolicy.
// returns a SiteProtectionUnitsBulkAdditionJobable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
// [Find more info here]
// 
// [Find more info here]: https://learn.microsoft.com/graph/api/siteprotectionunitsbulkadditionjobs-get?view=graph-rest-1.0
func (m *BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) Get(ctx context.Context, requestConfiguration *BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SiteProtectionUnitsBulkAdditionJobable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSiteProtectionUnitsBulkAdditionJobFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.SiteProtectionUnitsBulkAdditionJobable), nil
}
// ToGetRequestInformation get a siteProtectionUnitsBulkAdditionJob object by the ID associated with a sharePointProtectionPolicy.
// returns a *RequestInformation when successful
func (m *BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder when successful
func (m *BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) WithUrl(rawUrl string)(*BackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder) {
    return NewBackupRestoreSharePointProtectionPoliciesItemSiteProtectionUnitsBulkAdditionJobsSiteProtectionUnitsBulkAdditionJobItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
