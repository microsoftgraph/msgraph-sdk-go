package users

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemOutlookRequestBuilder provides operations to manage the outlook property of the microsoft.graph.user entity.
type ItemOutlookRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemOutlookRequestBuilderGetQueryParameters get outlook from users
type ItemOutlookRequestBuilderGetQueryParameters struct {
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemOutlookRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemOutlookRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemOutlookRequestBuilderGetQueryParameters
}
// NewItemOutlookRequestBuilderInternal instantiates a new OutlookRequestBuilder and sets the default values.
func NewItemOutlookRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookRequestBuilder) {
    m := &ItemOutlookRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/users/{user%2Did}/outlook{?%24select}", pathParameters),
    }
    return m
}
// NewItemOutlookRequestBuilder instantiates a new OutlookRequestBuilder and sets the default values.
func NewItemOutlookRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemOutlookRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemOutlookRequestBuilderInternal(urlParams, requestAdapter)
}
// Get get outlook from users
func (m *ItemOutlookRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemOutlookRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OutlookUserable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOutlookUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OutlookUserable), nil
}
// MasterCategories provides operations to manage the masterCategories property of the microsoft.graph.outlookUser entity.
func (m *ItemOutlookRequestBuilder) MasterCategories()(*ItemOutlookMasterCategoriesRequestBuilder) {
    return NewItemOutlookMasterCategoriesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SupportedLanguages provides operations to call the supportedLanguages method.
func (m *ItemOutlookRequestBuilder) SupportedLanguages()(*ItemOutlookSupportedLanguagesRequestBuilder) {
    return NewItemOutlookSupportedLanguagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SupportedTimeZones provides operations to call the supportedTimeZones method.
func (m *ItemOutlookRequestBuilder) SupportedTimeZones()(*ItemOutlookSupportedTimeZonesRequestBuilder) {
    return NewItemOutlookSupportedTimeZonesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// SupportedTimeZonesWithTimeZoneStandard provides operations to call the supportedTimeZones method.
func (m *ItemOutlookRequestBuilder) SupportedTimeZonesWithTimeZoneStandard(timeZoneStandard *string)(*ItemOutlookSupportedTimeZonesWithTimeZoneStandardRequestBuilder) {
    return NewItemOutlookSupportedTimeZonesWithTimeZoneStandardRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter, timeZoneStandard)
}
// ToGetRequestInformation get outlook from users
func (m *ItemOutlookRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemOutlookRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.BaseRequestBuilder.UrlTemplate
    requestInfo.PathParameters = m.BaseRequestBuilder.PathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    requestInfo.Headers.Add("Accept", "application/json")
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
