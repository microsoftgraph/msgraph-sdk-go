package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder provides operations to call the getByPath method.
type ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1AnalyticsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Analytics()(*ItemGetbypathwithpathGetbypathwithpath1AnalyticsRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1AnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1ColumnsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Columns()(*ItemGetbypathwithpathGetbypathwithpath1ColumnsRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1ColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderInternal instantiates a new ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder and sets the default values.
func NewItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, path1 *string)(*ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) {
    m := &ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/getByPath(path='{path}')/getByPath(path='{path1}')", pathParameters),
    }
    if path1 != nil {
        m.BaseRequestBuilder.PathParameters["path1"] = *path1
    }
    return m
}
// NewItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder instantiates a new ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder and sets the default values.
func NewItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderInternal(urlParams, requestAdapter, nil)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1ContenttypesContentTypesRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) ContentTypes()(*ItemGetbypathwithpathGetbypathwithpath1ContenttypesContentTypesRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1ContenttypesContentTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1CreatedbyuserCreatedByUserRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) CreatedByUser()(*ItemGetbypathwithpathGetbypathwithpath1CreatedbyuserCreatedByUserRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1CreatedbyuserCreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Drive provides operations to manage the drive property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1DriveRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Drive()(*ItemGetbypathwithpathGetbypathwithpath1DriveRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1DriveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Drives provides operations to manage the drives property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1DrivesRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Drives()(*ItemGetbypathwithpathGetbypathwithpath1DrivesRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1DrivesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExternalColumns provides operations to manage the externalColumns property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) ExternalColumns()(*ItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1ExternalcolumnsExternalColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get invoke function getByPath
// returns a Siteable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Get(ctx context.Context, requestConfiguration *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, error) {
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
// returns a *ItemGetbypathwithpathGetbypathwithpath1ItemsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Items()(*ItemGetbypathwithpathGetbypathwithpath1ItemsRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1ItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) LastModifiedByUser()(*ItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1LastmodifiedbyuserLastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Lists provides operations to manage the lists property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1ListsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Lists()(*ItemGetbypathwithpathGetbypathwithpath1ListsRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1ListsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Onenote()(*ItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1OnenoteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1OperationsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Operations()(*ItemGetbypathwithpathGetbypathwithpath1OperationsRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1OperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Pages provides operations to manage the pages property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1PagesRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Pages()(*ItemGetbypathwithpathGetbypathwithpath1PagesRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1PagesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1PermissionsRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Permissions()(*ItemGetbypathwithpathGetbypathwithpath1PermissionsRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1PermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sites provides operations to manage the sites property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1SitesRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) Sites()(*ItemGetbypathwithpathGetbypathwithpath1SitesRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1SitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermStore provides operations to manage the termStore property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1TermstoreTermStoreRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) TermStore()(*ItemGetbypathwithpathGetbypathwithpath1TermstoreTermStoreRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1TermstoreTermStoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermStores provides operations to manage the termStores property of the microsoft.graph.site entity.
// returns a *ItemGetbypathwithpathGetbypathwithpath1TermstoresTermStoresRequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) TermStores()(*ItemGetbypathwithpathGetbypathwithpath1TermstoresTermStoresRequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1TermstoresTermStoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation invoke function getByPath
// returns a *RequestInformation when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder when successful
func (m *ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) WithUrl(rawUrl string)(*ItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder) {
    return NewItemGetbypathwithpathGetbypathwithpath1GetByPathWithPath1RequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
