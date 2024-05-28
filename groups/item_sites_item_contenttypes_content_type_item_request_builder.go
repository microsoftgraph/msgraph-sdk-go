package groups

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// ItemSitesItemContenttypesContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.site entity.
type ItemSitesItemContenttypesContentTypeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// ItemSitesItemContenttypesContentTypeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContenttypesContentTypeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ItemSitesItemContenttypesContentTypeItemRequestBuilderGetQueryParameters the collection of content types defined for this site.
type ItemSitesItemContenttypesContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ItemSitesItemContenttypesContentTypeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContenttypesContentTypeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ItemSitesItemContenttypesContentTypeItemRequestBuilderGetQueryParameters
}
// ItemSitesItemContenttypesContentTypeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ItemSitesItemContenttypesContentTypeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssociateWithHubSites provides operations to call the associateWithHubSites method.
// returns a *ItemSitesItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) AssociateWithHubSites()(*ItemSitesItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) {
    return NewItemSitesItemContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Base provides operations to manage the base property of the microsoft.graph.contentType entity.
// returns a *ItemSitesItemContenttypesItemBaseRequestBuilder when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) Base()(*ItemSitesItemContenttypesItemBaseRequestBuilder) {
    return NewItemSitesItemContenttypesItemBaseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BaseTypes provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
// returns a *ItemSitesItemContenttypesItemBasetypesBaseTypesRequestBuilder when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) BaseTypes()(*ItemSitesItemContenttypesItemBasetypesBaseTypesRequestBuilder) {
    return NewItemSitesItemContenttypesItemBasetypesBaseTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ColumnLinks provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
// returns a *ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) ColumnLinks()(*ItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilder) {
    return NewItemSitesItemContenttypesItemColumnlinksColumnLinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ColumnPositions provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
// returns a *ItemSitesItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) ColumnPositions()(*ItemSitesItemContenttypesItemColumnpositionsColumnPositionsRequestBuilder) {
    return NewItemSitesItemContenttypesItemColumnpositionsColumnPositionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.contentType entity.
// returns a *ItemSitesItemContenttypesItemColumnsRequestBuilder when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) Columns()(*ItemSitesItemContenttypesItemColumnsRequestBuilder) {
    return NewItemSitesItemContenttypesItemColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewItemSitesItemContenttypesContentTypeItemRequestBuilderInternal instantiates a new ItemSitesItemContenttypesContentTypeItemRequestBuilder and sets the default values.
func NewItemSitesItemContenttypesContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContenttypesContentTypeItemRequestBuilder) {
    m := &ItemSitesItemContenttypesContentTypeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/groups/{group%2Did}/sites/{site%2Did}/contentTypes/{contentType%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewItemSitesItemContenttypesContentTypeItemRequestBuilder instantiates a new ItemSitesItemContenttypesContentTypeItemRequestBuilder and sets the default values.
func NewItemSitesItemContenttypesContentTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ItemSitesItemContenttypesContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewItemSitesItemContenttypesContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToDefaultContentLocation provides operations to call the copyToDefaultContentLocation method.
// returns a *ItemSitesItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*ItemSitesItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) {
    return NewItemSitesItemContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property contentTypes for groups
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *ItemSitesItemContenttypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
// Get the collection of content types defined for this site.
// returns a ContentTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *ItemSitesItemContenttypesContentTypeItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
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
// returns a *ItemSitesItemContenttypesItemIspublishedIsPublishedRequestBuilder when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) IsPublished()(*ItemSitesItemContenttypesItemIspublishedIsPublishedRequestBuilder) {
    return NewItemSitesItemContenttypesItemIspublishedIsPublishedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property contentTypes in groups
// returns a ContentTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ItemSitesItemContenttypesContentTypeItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
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
// returns a *ItemSitesItemContenttypesItemPublishRequestBuilder when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) Publish()(*ItemSitesItemContenttypesItemPublishRequestBuilder) {
    return NewItemSitesItemContenttypesItemPublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property contentTypes for groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemContenttypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformationWithMethodAndUrlTemplateAndPathParameters(i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE, m.BaseRequestBuilder.UrlTemplate, m.BaseRequestBuilder.PathParameters)
    if requestConfiguration != nil {
        requestInfo.Headers.AddAll(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    requestInfo.Headers.TryAdd("Accept", "application/json")
    return requestInfo, nil
}
// ToGetRequestInformation the collection of content types defined for this site.
// returns a *RequestInformation when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *ItemSitesItemContenttypesContentTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property contentTypes in groups
// returns a *RequestInformation when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ItemSitesItemContenttypesContentTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *ItemSitesItemContenttypesItemUnpublishRequestBuilder when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) Unpublish()(*ItemSitesItemContenttypesItemUnpublishRequestBuilder) {
    return NewItemSitesItemContenttypesItemUnpublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *ItemSitesItemContenttypesContentTypeItemRequestBuilder when successful
func (m *ItemSitesItemContenttypesContentTypeItemRequestBuilder) WithUrl(rawUrl string)(*ItemSitesItemContenttypesContentTypeItemRequestBuilder) {
    return NewItemSitesItemContenttypesContentTypeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
