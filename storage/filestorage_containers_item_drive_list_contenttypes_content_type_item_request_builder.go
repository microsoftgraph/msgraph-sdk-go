package storage

import (
    "context"
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
)

// FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder struct {
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.BaseRequestBuilder
}
// FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderGetQueryParameters
}
// FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestHeaders
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssociateWithHubSites provides operations to call the associateWithHubSites method.
// returns a *FilestorageContainersItemDriveListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) AssociateWithHubSites()(*FilestorageContainersItemDriveListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesItemAssociatewithhubsitesAssociateWithHubSitesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Base provides operations to manage the base property of the microsoft.graph.contentType entity.
// returns a *FilestorageContainersItemDriveListContenttypesItemBaseRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) Base()(*FilestorageContainersItemDriveListContenttypesItemBaseRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesItemBaseRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// BaseTypes provides operations to manage the baseTypes property of the microsoft.graph.contentType entity.
// returns a *FilestorageContainersItemDriveListContenttypesItemBasetypesBaseTypesRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) BaseTypes()(*FilestorageContainersItemDriveListContenttypesItemBasetypesBaseTypesRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesItemBasetypesBaseTypesRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ColumnLinks provides operations to manage the columnLinks property of the microsoft.graph.contentType entity.
// returns a *FilestorageContainersItemDriveListContenttypesItemColumnlinksColumnLinksRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) ColumnLinks()(*FilestorageContainersItemDriveListContenttypesItemColumnlinksColumnLinksRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesItemColumnlinksColumnLinksRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ColumnPositions provides operations to manage the columnPositions property of the microsoft.graph.contentType entity.
// returns a *FilestorageContainersItemDriveListContenttypesItemColumnpositionsColumnPositionsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) ColumnPositions()(*FilestorageContainersItemDriveListContenttypesItemColumnpositionsColumnPositionsRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesItemColumnpositionsColumnPositionsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Columns provides operations to manage the columns property of the microsoft.graph.contentType entity.
// returns a *FilestorageContainersItemDriveListContenttypesItemColumnsRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) Columns()(*FilestorageContainersItemDriveListContenttypesItemColumnsRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesItemColumnsRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// NewFilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderInternal instantiates a new FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) {
    m := &FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder{
        BaseRequestBuilder: *i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewBaseRequestBuilder(requestAdapter, "{+baseurl}/storage/fileStorage/containers/{fileStorageContainer%2Did}/drive/list/contentTypes/{contentType%2Did}{?%24expand,%24select}", pathParameters),
    }
    return m
}
// NewFilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder instantiates a new FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder and sets the default values.
func NewFilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewFilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToDefaultContentLocation provides operations to call the copyToDefaultContentLocation method.
// returns a *FilestorageContainersItemDriveListContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*FilestorageContainersItemDriveListContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesItemCopytodefaultcontentlocationCopyToDefaultContentLocationRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Delete delete navigation property contentTypes for storage
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) Delete(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(error) {
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
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) Get(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderGetRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
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
// returns a *FilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) IsPublished()(*FilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesItemIspublishedIsPublishedRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// Patch update the navigation property contentTypes in storage
// returns a ContentTypeable when successful
// returns a ODataError error when the service returns a 4XX or 5XX status code
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) Patch(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderPatchRequestConfiguration)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
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
// returns a *FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) Publish()(*FilestorageContainersItemDriveListContenttypesItemPublishRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesItemPublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// ToDeleteRequestInformation delete navigation property contentTypes for storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) ToDeleteRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) ToGetRequestInformation(ctx context.Context, requestConfiguration *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// ToPatchRequestInformation update the navigation property contentTypes in storage
// returns a *RequestInformation when successful
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) ToPatchRequestInformation(ctx context.Context, body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
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
// returns a *FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) Unpublish()(*FilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesItemUnpublishRequestBuilderInternal(m.BaseRequestBuilder.PathParameters, m.BaseRequestBuilder.RequestAdapter)
}
// WithUrl returns a request builder with the provided arbitrary URL. Using this method means any other path or query parameters are ignored.
// returns a *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder when successful
func (m *FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) WithUrl(rawUrl string)(*FilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder) {
    return NewFilestorageContainersItemDriveListContenttypesContentTypeItemRequestBuilder(rawUrl, m.BaseRequestBuilder.RequestAdapter);
}
