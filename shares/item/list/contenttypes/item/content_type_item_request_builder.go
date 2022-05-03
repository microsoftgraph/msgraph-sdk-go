package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i0df9b813acca9f52cb58191b981da2631e2ce7ec1e0f39baea115430a4d91ec7 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/columnlinks"
    i1ec18c0a4fbb8f904f34e72c387a8703eb3af7790657ffef68bd11ec2a9905ef "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/columns"
    i49bc8900bfdd8b3b1d3f1b52a9d8d36be89f0d252009782f35e12ba5c6fae675 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/associatewithhubsites"
    i4a1ce58e21fb551499684168220c8c6eee92bb5bc6a4b99639b513c5cfe71a1b "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/basetypes"
    i553035798d555384c1d78519561934573cbdb4b0f2bd9da038ce9a71c4991d05 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/columnpositions"
    i7f1730a5f677932cd121090203a3bc4480b7f148f3eae6dd7de05441a796ed7c "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/unpublish"
    i88052e6ec6afde3eb4c8a4e647edd6f1b4d0d4c30b1be1a17f71fdd90f644c90 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/ispublished"
    i934cd417c891824167d128199e27f353bdac8d75faa96747cb1f2f67634a829e "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/publish"
    ibda66b96fbe7db9cdba0b0f6fb218ff35d517fc7b3ac53ebf9bdc4c45957d5c8 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/copytodefaultcontentlocation"
    if20ce4a9c8987ec45c17d11c97d9255ead0b63c6bb3cb2a948e79d243584e4de "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/base"
    i5a14dba7acc9a67d9b71d060c0e0de130dbeaf1185def884644520254fd0a2bf "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/columns/item"
    i684c0ce66c376c68b7eda568c4a0e38ff6f9905b99e380e5c7d6735c0ce5813b "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/basetypes/item"
    i798aa8d3eb2beecb739a9f0d6422dbcd303b5e2935628722cd84d0e6fbea4bb1 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/columnpositions/item"
    i8459e0e9fdbef8a8ad4c7a315789b0cca6cc0c5d53c4d9f685c083f3c7d8381d "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list/contenttypes/item/columnlinks/item"
)

// ContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.list entity.
type ContentTypeItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// ContentTypeItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContentTypeItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types present in this list.
type ContentTypeItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// ContentTypeItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContentTypeItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *ContentTypeItemRequestBuilderGetQueryParameters
}
// ContentTypeItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type ContentTypeItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// AssociateWithHubSites the associateWithHubSites property
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i49bc8900bfdd8b3b1d3f1b52a9d8d36be89f0d252009782f35e12ba5c6fae675.AssociateWithHubSitesRequestBuilder) {
    return i49bc8900bfdd8b3b1d3f1b52a9d8d36be89f0d252009782f35e12ba5c6fae675.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base the base property
func (m *ContentTypeItemRequestBuilder) Base()(*if20ce4a9c8987ec45c17d11c97d9255ead0b63c6bb3cb2a948e79d243584e4de.BaseRequestBuilder) {
    return if20ce4a9c8987ec45c17d11c97d9255ead0b63c6bb3cb2a948e79d243584e4de.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes the baseTypes property
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*i4a1ce58e21fb551499684168220c8c6eee92bb5bc6a4b99639b513c5cfe71a1b.BaseTypesRequestBuilder) {
    return i4a1ce58e21fb551499684168220c8c6eee92bb5bc6a4b99639b513c5cfe71a1b.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i684c0ce66c376c68b7eda568c4a0e38ff6f9905b99e380e5c7d6735c0ce5813b.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType%2Did1"] = id
    }
    return i684c0ce66c376c68b7eda568c4a0e38ff6f9905b99e380e5c7d6735c0ce5813b.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks the columnLinks property
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i0df9b813acca9f52cb58191b981da2631e2ce7ec1e0f39baea115430a4d91ec7.ColumnLinksRequestBuilder) {
    return i0df9b813acca9f52cb58191b981da2631e2ce7ec1e0f39baea115430a4d91ec7.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*i8459e0e9fdbef8a8ad4c7a315789b0cca6cc0c5d53c4d9f685c083f3c7d8381d.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink%2Did"] = id
    }
    return i8459e0e9fdbef8a8ad4c7a315789b0cca6cc0c5d53c4d9f685c083f3c7d8381d.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions the columnPositions property
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i553035798d555384c1d78519561934573cbdb4b0f2bd9da038ce9a71c4991d05.ColumnPositionsRequestBuilder) {
    return i553035798d555384c1d78519561934573cbdb4b0f2bd9da038ce9a71c4991d05.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i798aa8d3eb2beecb739a9f0d6422dbcd303b5e2935628722cd84d0e6fbea4bb1.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i798aa8d3eb2beecb739a9f0d6422dbcd303b5e2935628722cd84d0e6fbea4bb1.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns the columns property
func (m *ContentTypeItemRequestBuilder) Columns()(*i1ec18c0a4fbb8f904f34e72c387a8703eb3af7790657ffef68bd11ec2a9905ef.ColumnsRequestBuilder) {
    return i1ec18c0a4fbb8f904f34e72c387a8703eb3af7790657ffef68bd11ec2a9905ef.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i5a14dba7acc9a67d9b71d060c0e0de130dbeaf1185def884644520254fd0a2bf.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition%2Did"] = id
    }
    return i5a14dba7acc9a67d9b71d060c0e0de130dbeaf1185def884644520254fd0a2bf.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem%2Did}/list/contentTypes/{contentType%2Did}{?%24select,%24expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*ibda66b96fbe7db9cdba0b0f6fb218ff35d517fc7b3ac53ebf9bdc4c45957d5c8.CopyToDefaultContentLocationRequestBuilder) {
    return ibda66b96fbe7db9cdba0b0f6fb218ff35d517fc7b3ac53ebf9bdc4c45957d5c8.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for shares
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete navigation property contentTypes for shares
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *ContentTypeItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *ContentTypeItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property contentTypes in shares
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update the navigation property contentTypes in shares
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ContentTypeItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete navigation property contentTypes for shares
func (m *ContentTypeItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete navigation property contentTypes for shares
func (m *ContentTypeItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *ContentTypeItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetWithRequestConfigurationAndResponseHandler the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *ContentTypeItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateContentTypeFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable), nil
}
// IsPublished provides operations to call the isPublished method.
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i88052e6ec6afde3eb4c8a4e647edd6f1b4d0d4c30b1be1a17f71fdd90f644c90.IsPublishedRequestBuilder) {
    return i88052e6ec6afde3eb4c8a4e647edd6f1b4d0d4c30b1be1a17f71fdd90f644c90.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in shares
func (m *ContentTypeItemRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update the navigation property contentTypes in shares
func (m *ContentTypeItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ContentTypeable, requestConfiguration *ContentTypeItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Publish the publish property
func (m *ContentTypeItemRequestBuilder) Publish()(*i934cd417c891824167d128199e27f353bdac8d75faa96747cb1f2f67634a829e.PublishRequestBuilder) {
    return i934cd417c891824167d128199e27f353bdac8d75faa96747cb1f2f67634a829e.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i7f1730a5f677932cd121090203a3bc4480b7f148f3eae6dd7de05441a796ed7c.UnpublishRequestBuilder) {
    return i7f1730a5f677932cd121090203a3bc4480b7f148f3eae6dd7de05441a796ed7c.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
