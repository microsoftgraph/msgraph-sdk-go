package sites

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemListsItemContenttypesContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ItemListsItemContenttypesContentTypeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemListsItemContenttypesContentTypeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemContenttypesContentTypeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemListsItemContenttypesContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type ItemListsItemContenttypesContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemListsItemContenttypesContentTypeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemContenttypesContentTypeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemListsItemContenttypesContentTypeItemRequestBuilderGetQueryParameters
}
// ItemListsItemContenttypesContentTypeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemListsItemContenttypesContentTypeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssociateWithHubSites provides operations to call the associateWithHubSites method.
// returns a *ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) AssociateWithHubSites()(*ItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) {
    return NewItemListsItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Base provides operations to manage the base property of the microsoft.graph.contentType entity.
// returns a *ItemListsItemContenttypesItemBaseRequestBuilder when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) Base()(*ItemListsItemContenttypesItemBaseRequestBuilder) {
    return NewItemListsItemContenttypesItemBaseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BaseTypes provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
// returns a *ItemListsItemContenttypesItemBasetypesBaseTypesRequestBuilder when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) BaseTypes()(*ItemListsItemContenttypesItemBasetypesBaseTypesRequestBuilder) {
    return NewItemListsItemContenttypesItemBasetypesBaseTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ColumnLinks provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
// returns a *ItemListsItemContenttypesItemColumnlinksColumnLinksRequestBuilder when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) ColumnLinks()(*ItemListsItemContenttypesItemColumnlinksColumnLinksRequestBuilder) {
    return NewItemListsItemContenttypesItemColumnlinksColumnLinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ColumnPositions provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
// returns a *ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) ColumnPositions()(*ItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder) {
    return NewItemListsItemContenttypesItemColumnpositionsColumnPositionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.contentType entity.
// returns a *ItemListsItemContenttypesItemColumnsRequestBuilder when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) Columns()(*ItemListsItemContenttypesItemColumnsRequestBuilder) {
    return NewItemListsItemContenttypesItemColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemListsItemContenttypesContentTypeItemRequestBuilderInternal instantiates a new ItemListsItemContenttypesContentTypeItemRequestBuilder and sets the default values.
func NewItemListsItemContenttypesContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContenttypesContentTypeItemRequestBuilder) {
    m := &ItemListsItemContenttypesContentTypeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/sites/{site%2Did}/lists/{list%2Did}/contentTypes/{contentType%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemListsItemContenttypesContentTypeItemRequestBuilder instantiates a new ItemListsItemContenttypesContentTypeItemRequestBuilder and sets the default values.
func NewItemListsItemContenttypesContentTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemListsItemContenttypesContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemListsItemContenttypesContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToDefaultContentLocation provides operations to call the copyToDefaultContentLocation method.
// returns a *ItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*ItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) {
    return NewItemListsItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property contentTypes for sites
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemListsItemContenttypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of content types present in this list.
// returns a ContentTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemListsItemContenttypesContentTypeItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
    requestInfo, err := m.ToGetRequestInformation(ctx, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
// returns a *ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) IsPublished()(*ItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilder) {
    return NewItemListsItemContenttypesItemIspublishedIsPublishedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property contentTypes in sites
// returns a ContentTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ItemListsItemContenttypesContentTypeItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
    requestInfo, err := m.ToPatchRequestInformation(ctx, body, requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "XXX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.BaseRequestBuilder.RequestAdapter.Send(ctx, requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentTypeFromDiscriminatorValue, errorMapping)
    if err != nil {
        return nil, err
    }
    if res == nil {
        return nil, nil
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable), nil
}
// Publish provides operations to call the publish method.
// returns a *ItemListsItemContenttypesItemPublishRequestBuilder when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) Publish()(*ItemListsItemContenttypesItemPublishRequestBuilder) {
    return NewItemListsItemContenttypesItemPublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property contentTypes for sites
// returns a *RequestInformation when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemContenttypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of content types present in this list.
// returns a *RequestInformation when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemListsItemContenttypesContentTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property contentTypes in sites
// returns a *RequestInformation when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ItemListsItemContenttypesContentTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// Unpublish provides operations to call the unpublish method.
// returns a *ItemListsItemContenttypesItemUnpublishRequestBuilder when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) Unpublish()(*ItemListsItemContenttypesItemUnpublishRequestBuilder) {
    return NewItemListsItemContenttypesItemUnpublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemListsItemContenttypesContentTypeItemRequestBuilder when successful
func (m *ItemListsItemContenttypesContentTypeItemRequestBuilder) WithUrl(rawUrl string)(*ItemListsItemContenttypesContentTypeItemRequestBuilder) {
    return NewItemListsItemContenttypesContentTypeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
