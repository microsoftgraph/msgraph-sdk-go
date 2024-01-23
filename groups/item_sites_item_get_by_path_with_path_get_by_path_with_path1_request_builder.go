package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder provides operations to call the getByPath method.
type ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// Analytics provides operations to manage the analytics property of the microsoft.graph.site entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) Analytics()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1AnalyticsRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1AnalyticsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.site entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) Columns()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1ColumnsRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1ColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilderInternal instantiates a new GetByPathWithPath1RequestBuilder and sets the default values.
func NewItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter, path1 *string)(*ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) {
    m := &ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/getByPath(path='{path}')/getByPath(path='{path1}')", pathParameters),
    }
    if path1 != nil {
        m.BaseRequestBuilder.PathParameters["path1"] = *path1
    }
    return m
}
// NewItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder instantiates a new GetByPathWithPath1RequestBuilder and sets the default values.
func NewItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilderInternal(urlParams, requestAdapter, nil)
}
// ContentTypes provides operations to manage the contentTypes property of the microsoft.graph.site entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) ContentTypes()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1ContentTypesRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1ContentTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// CreatedByUser provides operations to manage the createdByUser property of the microsoft.graph.baseItem entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) CreatedByUser()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1CreatedByUserRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1CreatedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Drive provides operations to manage the drive property of the microsoft.graph.site entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) Drive()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1DriveRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Drives provides operations to manage the drives property of the microsoft.graph.site entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) Drives()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1DrivesRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1DrivesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ExternalColumns provides operations to manage the externalColumns property of the microsoft.graph.site entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) ExternalColumns()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1ExternalColumnsRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1ExternalColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Get invoke function getByPath
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Siteable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
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
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) Items()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1ItemsRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1ItemsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// LastModifiedByUser provides operations to manage the lastModifiedByUser property of the microsoft.graph.baseItem entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) LastModifiedByUser()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1LastModifiedByUserRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Lists provides operations to manage the lists property of the microsoft.graph.site entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) Lists()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1ListsRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1ListsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Onenote provides operations to manage the onenote property of the microsoft.graph.site entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) Onenote()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1OnenoteRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1OnenoteRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Operations provides operations to manage the operations property of the microsoft.graph.site entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) Operations()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1OperationsRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1OperationsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Permissions provides operations to manage the permissions property of the microsoft.graph.site entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) Permissions()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1PermissionsRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1PermissionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Sites provides operations to manage the sites property of the microsoft.graph.site entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) Sites()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1SitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermStore provides operations to manage the termStore property of the microsoft.graph.site entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) TermStore()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1TermStoreRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// TermStores provides operations to manage the termStores property of the microsoft.graph.site entity.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) TermStores()(*ItemSitesItemGetByPathWithPathGetByPathWithPath1TermStoresRequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1TermStoresRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToGetRequestInformation invoke function getByPath
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
func (m *ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder) {
    return NewItemSitesItemGetByPathWithPathGetByPathWithPath1RequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
