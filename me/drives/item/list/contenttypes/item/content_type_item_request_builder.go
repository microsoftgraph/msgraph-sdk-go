package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i40b40c14afe20846df7697563943ff4147fcf3513ce68277ca2de88e464c473d "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/columnlinks"
    i53f314f0c3675cac45d21567fca93ab1edb06078259cfaea249d2ec6d2c1fdc3 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/copytodefaultcontentlocation"
    i9ce8000fb091cd3a5804314f05a98b6a927489500c5f2f1144c7f99e190c4067 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/associatewithhubsites"
    iaa8801033902add70125689951f2f36de2346e26374e4686d5f9c9de194cc05c "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/columns"
    iab9d147d4fb3d030a6b4abf8d305a14336264a5196d363e98b75cbf009a5334b "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/ispublished"
    iad6947c553ec5a4ce509828e095618d07a6d717e1de82221d777154d9192e208 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/unpublish"
    ic7b6bd6de956f12a8a38cb22bd58e71544ae32b7980a25b0536b301ca19a7bed "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/basetypes"
    idc99e113446b35e57939736e74d537663450beb1feaebdb6606c46591d5b5fb9 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/base"
    ide567c8b90b031fe74c6884dcbedb4274d796cf26b54c0e2d34d12cfef78aa17 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/columnpositions"
    ifebd77c70e2b2b28ca8e8b783ee20b065141c868e8a8be9109b2e89c805d5cf1 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/publish"
    i46bd691a897439fa677fdb51d5324696838d658e10fc675de7973db14cd34306 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/columnpositions/item"
    i637aca2c30344f22f5f4c7820f1574d5fe884d7b3bdb14a043cdcf1fd79e1dbb "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/columns/item"
    i6e6310aa8e8c7f0e53fc51025248ae3a6a4a30cb8bc899c5eae4a34bcf3b7dfb "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/columnlinks/item"
    if0d46cb9718fba761345686cf114220325cd78ab599073b77cfccc0543a3036f "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/list/contenttypes/item/basetypes/item"
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i9ce8000fb091cd3a5804314f05a98b6a927489500c5f2f1144c7f99e190c4067.AssociateWithHubSitesRequestBuilder) {
    return i9ce8000fb091cd3a5804314f05a98b6a927489500c5f2f1144c7f99e190c4067.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base the base property
func (m *ContentTypeItemRequestBuilder) Base()(*idc99e113446b35e57939736e74d537663450beb1feaebdb6606c46591d5b5fb9.BaseRequestBuilder) {
    return idc99e113446b35e57939736e74d537663450beb1feaebdb6606c46591d5b5fb9.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes the baseTypes property
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*ic7b6bd6de956f12a8a38cb22bd58e71544ae32b7980a25b0536b301ca19a7bed.BaseTypesRequestBuilder) {
    return ic7b6bd6de956f12a8a38cb22bd58e71544ae32b7980a25b0536b301ca19a7bed.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*if0d46cb9718fba761345686cf114220325cd78ab599073b77cfccc0543a3036f.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return if0d46cb9718fba761345686cf114220325cd78ab599073b77cfccc0543a3036f.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks the columnLinks property
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i40b40c14afe20846df7697563943ff4147fcf3513ce68277ca2de88e464c473d.ColumnLinksRequestBuilder) {
    return i40b40c14afe20846df7697563943ff4147fcf3513ce68277ca2de88e464c473d.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i6e6310aa8e8c7f0e53fc51025248ae3a6a4a30cb8bc899c5eae4a34bcf3b7dfb.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i6e6310aa8e8c7f0e53fc51025248ae3a6a4a30cb8bc899c5eae4a34bcf3b7dfb.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions the columnPositions property
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*ide567c8b90b031fe74c6884dcbedb4274d796cf26b54c0e2d34d12cfef78aa17.ColumnPositionsRequestBuilder) {
    return ide567c8b90b031fe74c6884dcbedb4274d796cf26b54c0e2d34d12cfef78aa17.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i46bd691a897439fa677fdb51d5324696838d658e10fc675de7973db14cd34306.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i46bd691a897439fa677fdb51d5324696838d658e10fc675de7973db14cd34306.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns the columns property
func (m *ContentTypeItemRequestBuilder) Columns()(*iaa8801033902add70125689951f2f36de2346e26374e4686d5f9c9de194cc05c.ColumnsRequestBuilder) {
    return iaa8801033902add70125689951f2f36de2346e26374e4686d5f9c9de194cc05c.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i637aca2c30344f22f5f4c7820f1574d5fe884d7b3bdb14a043cdcf1fd79e1dbb.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i637aca2c30344f22f5f4c7820f1574d5fe884d7b3bdb14a043cdcf1fd79e1dbb.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/list/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*i53f314f0c3675cac45d21567fca93ab1edb06078259cfaea249d2ec6d2c1fdc3.CopyToDefaultContentLocationRequestBuilder) {
    return i53f314f0c3675cac45d21567fca93ab1edb06078259cfaea249d2ec6d2c1fdc3.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for me
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
// CreatePatchRequestInformation update the navigation property contentTypes in me
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
// Delete delete navigation property contentTypes for me
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
func (m *ContentTypeItemRequestBuilder) IsPublished()(*iab9d147d4fb3d030a6b4abf8d305a14336264a5196d363e98b75cbf009a5334b.IsPublishedRequestBuilder) {
    return iab9d147d4fb3d030a6b4abf8d305a14336264a5196d363e98b75cbf009a5334b.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in me
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
func (m *ContentTypeItemRequestBuilder) Publish()(*ifebd77c70e2b2b28ca8e8b783ee20b065141c868e8a8be9109b2e89c805d5cf1.PublishRequestBuilder) {
    return ifebd77c70e2b2b28ca8e8b783ee20b065141c868e8a8be9109b2e89c805d5cf1.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *ContentTypeItemRequestBuilder) Unpublish()(*iad6947c553ec5a4ce509828e095618d07a6d717e1de82221d777154d9192e208.UnpublishRequestBuilder) {
    return iad6947c553ec5a4ce509828e095618d07a6d717e1de82221d777154d9192e208.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
