package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
type ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilderGetQueryParameters identity of the user who created the item. Read-only.
type ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilderGetQueryParameters
}
// NewItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilderInternal instantiates a new ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder) {
    m := &ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/pages/{baseSitePage%2Did}/createdByUser{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder instantiates a new ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder and sets the default values.
func NewItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilderInternal(urlParams, requestAdapter)
}
// Get identity of the user who created the item. Read-only.
// returns a Userable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUserFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Userable), nil
}
// MailboxSettings the mailboxSettings property
// returns a *ItemSitesItemPagesItemCreatedbyuserMailboxsettingsMailboxSettingsRequestBuilder when successful
func (m *ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder) MailboxSettings()(*ItemSitesItemPagesItemCreatedbyuserMailboxsettingsMailboxSettingsRequestBuilder) {
    return NewItemSitesItemPagesItemCreatedbyuserMailboxsettingsMailboxSettingsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ServiceProvisioningErrors the serviceProvisioningErrors property
// returns a *ItemSitesItemPagesItemCreatedbyuserServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder when successful
func (m *ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder) ServiceProvisioningErrors()(*ItemSitesItemPagesItemCreatedbyuserServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilder) {
    return NewItemSitesItemPagesItemCreatedbyuserServiceprovisioningerrorsServiceProvisioningErrorsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation identity of the user who created the item. Read-only.
// returns a *RequestInformation when successful
func (m *ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder when successful
func (m *ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder) {
    return NewItemSitesItemPagesItemCreatedbyuserCreatedByUserRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
