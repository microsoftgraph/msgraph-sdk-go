package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0adbaf2bb9aa63a5e604a5fa371b072519032b803f737301b0b52921cc1d3a35 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/unpublish"
    i1a7695ad2fabf538093408204c95d6189330219c052c9f612b44a85734601b1c "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/ispublished"
    i278379c29b4a8817ca06b917afe4d8c15eeafd2fd0f5640b5269d826d7333584 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columnlinks"
    i3a448b37181c7678264208ecae87fe0e94cc7ea93cf9ce950d8c69413aaec6db "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/associatewithhubsites"
    i6e3f707a321441169d0528a8f9d4e00d29f3fdc3ce146c739c6537af4b41e473 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/base"
    i91d7fb036b4d816c4d7231c78f8e9680172ba5a52ca457e5ae197cfc261d3361 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/publish"
    ib477091a8a0eb17f6038d6809d39677533b56c9feeae3f59b9fc6af836d4766e "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columnpositions"
    iba33625fbf7112f94866f86d5bf89ec9e16e3e3d0e4f825b748afa7b67c6fb15 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/copytodefaultcontentlocation"
    ida9964b630004248401c792f6f5212cb1ddb19922f1fb202be62edb471410340 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/basetypes"
    iff966f1c31329b7dc1c1986153d73717d35370de251cc9894dff12910c064f2a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columns"
    i2150d109d18cf045a75c8d5444fead62bb38570ae3bd07afea3c979c0b6f7858 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columns/item"
    i2cf1a02d48d3a21ccb17cada00470bc9d1fabf5dcd49a56a9761d5192beebeb4 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columnpositions/item"
    i541fa36eb7a1b1f0420e9bae56b716dee553af507c47eb7a9eef7659f1cef097 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/basetypes/item"
    ia4e76c4ea17d224738d3bd333eae6742adbcd53e331ba5959a9aed631af5d7c2 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/lists/item/contenttypes/item/columnlinks/item"
)

// ContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ContentTypeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContentTypeItemRequestBuilderDeleteOptions options for Delete
type ContentTypeItemRequestBuilderDeleteOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ContentTypeItemRequestBuilderGetOptions options for Get
type ContentTypeItemRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *ContentTypeItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ContentTypeItemRequestBuilderPatchOptions options for Patch
type ContentTypeItemRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AssociateWithHubSites the associateWithHubSites property
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i3a448b37181c7678264208ecae87fe0e94cc7ea93cf9ce950d8c69413aaec6db.AssociateWithHubSitesRequestBuilder) {
    return i3a448b37181c7678264208ecae87fe0e94cc7ea93cf9ce950d8c69413aaec6db.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base the base property
func (m *ContentTypeItemRequestBuilder) Base()(*i6e3f707a321441169d0528a8f9d4e00d29f3fdc3ce146c739c6537af4b41e473.BaseRequestBuilder) {
    return i6e3f707a321441169d0528a8f9d4e00d29f3fdc3ce146c739c6537af4b41e473.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes the baseTypes property
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*ida9964b630004248401c792f6f5212cb1ddb19922f1fb202be62edb471410340.BaseTypesRequestBuilder) {
    return ida9964b630004248401c792f6f5212cb1ddb19922f1fb202be62edb471410340.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.lists.item.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i541fa36eb7a1b1f0420e9bae56b716dee553af507c47eb7a9eef7659f1cef097.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return i541fa36eb7a1b1f0420e9bae56b716dee553af507c47eb7a9eef7659f1cef097.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks the columnLinks property
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i278379c29b4a8817ca06b917afe4d8c15eeafd2fd0f5640b5269d826d7333584.ColumnLinksRequestBuilder) {
    return i278379c29b4a8817ca06b917afe4d8c15eeafd2fd0f5640b5269d826d7333584.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.lists.item.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*ia4e76c4ea17d224738d3bd333eae6742adbcd53e331ba5959a9aed631af5d7c2.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return ia4e76c4ea17d224738d3bd333eae6742adbcd53e331ba5959a9aed631af5d7c2.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions the columnPositions property
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*ib477091a8a0eb17f6038d6809d39677533b56c9feeae3f59b9fc6af836d4766e.ColumnPositionsRequestBuilder) {
    return ib477091a8a0eb17f6038d6809d39677533b56c9feeae3f59b9fc6af836d4766e.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.lists.item.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i2cf1a02d48d3a21ccb17cada00470bc9d1fabf5dcd49a56a9761d5192beebeb4.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i2cf1a02d48d3a21ccb17cada00470bc9d1fabf5dcd49a56a9761d5192beebeb4.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns the columns property
func (m *ContentTypeItemRequestBuilder) Columns()(*iff966f1c31329b7dc1c1986153d73717d35370de251cc9894dff12910c064f2a.ColumnsRequestBuilder) {
    return iff966f1c31329b7dc1c1986153d73717d35370de251cc9894dff12910c064f2a.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.lists.item.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i2150d109d18cf045a75c8d5444fead62bb38570ae3bd07afea3c979c0b6f7858.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i2150d109d18cf045a75c8d5444fead62bb38570ae3bd07afea3c979c0b6f7858.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/lists/{list_id}/contentTypes/{contentType_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewContentTypeItemRequestBuilder instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CopyToDefaultContentLocation the copyToDefaultContentLocation property
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*iba33625fbf7112f94866f86d5bf89ec9e16e3e3d0e4f825b748afa7b67c6fb15.CopyToDefaultContentLocationRequestBuilder) {
    return iba33625fbf7112f94866f86d5bf89ec9e16e3e3d0e4f825b748afa7b67c6fb15.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for sites
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeItemRequestBuilderDeleteOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformation(options *ContentTypeItemRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property contentTypes in sites
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(options *ContentTypeItemRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete delete navigation property contentTypes for sites
func (m *ContentTypeItemRequestBuilder) Delete(options *ContentTypeItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get(options *ContentTypeItemRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i1a7695ad2fabf538093408204c95d6189330219c052c9f612b44a85734601b1c.IsPublishedRequestBuilder) {
    return i1a7695ad2fabf538093408204c95d6189330219c052c9f612b44a85734601b1c.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in sites
func (m *ContentTypeItemRequestBuilder) Patch(options *ContentTypeItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Publish the publish property
func (m *ContentTypeItemRequestBuilder) Publish()(*i91d7fb036b4d816c4d7231c78f8e9680172ba5a52ca457e5ae197cfc261d3361.PublishRequestBuilder) {
    return i91d7fb036b4d816c4d7231c78f8e9680172ba5a52ca457e5ae197cfc261d3361.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i0adbaf2bb9aa63a5e604a5fa371b072519032b803f737301b0b52921cc1d3a35.UnpublishRequestBuilder) {
    return i0adbaf2bb9aa63a5e604a5fa371b072519032b803f737301b0b52921cc1d3a35.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
