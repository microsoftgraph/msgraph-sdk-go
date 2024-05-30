package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder provides operations to call the getByPath method.
type ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1AnalyticsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Analytics()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1AnalyticsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1AnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1ColumnsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Columns()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1ColumnsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1ColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderInternal instantiates a new ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder and sets the default values.
func NewItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, path1 *string)(*ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) {
    m := &ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/getByPath(path='{path}')/getByPath(path='{path1}')", pathParameters),
    }
    if path1 != nil {
        m.BaseRequestBuilder.PathParameters["path1"] = *path1
    }
    return m
}
// NewItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder instantiates a new ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder and sets the default values.
func NewItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderInternal(urlParams, requestAdapter, nil)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1ContenttypesContentTypesRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) ContentTypes()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1ContenttypesContentTypesRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1ContenttypesContentTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1CreatedbyuserCreatedByUserRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) CreatedByUser()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1CreatedbyuserCreatedByUserRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1CreatedbyuserCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Drive provides operations to manage the drive property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1DriveRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Drive()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1DriveRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1DriveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Drives provides operations to manage the drives property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1DrivesRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Drives()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1DrivesRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1DrivesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExternalColumns provides operations to manage the externalColumns property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) ExternalColumns()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get invoke function getByPath
// returns a Siteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateSiteFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable), nil
}
// Items provides operations to manage the items property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1ItemsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Items()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1ItemsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1ItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) LastModifiedByUser()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Lists provides operations to manage the lists property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1ListsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Lists()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1ListsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1ListsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Onenote()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1OperationsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Operations()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1OperationsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1OperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pages provides operations to manage the pages property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1PagesRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Pages()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1PagesRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1PagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1PermissionsRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Permissions()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1PermissionsRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1PermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sites provides operations to manage the sites property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1SitesRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Sites()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1SitesRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1SitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermStore provides operations to manage the termStore property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1TermstoreTermStoreRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) TermStore()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1TermstoreTermStoreRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1TermstoreTermStoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermStores provides operations to manage the termStores property of the microsoft.graph.site entity.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1TermstoresTermStoresRequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) TermStores()(*ItemSitesItemGetbypathwithpathGetbypathwithpath1TermstoresTermStoresRequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1TermstoresTermStoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation invoke function getByPath
// returns a *RequestInformation when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder when successful
func (m *ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) {
    return NewItemSitesItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
