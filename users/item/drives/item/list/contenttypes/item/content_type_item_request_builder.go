package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i3749973de34d8060e52ebe630f753ac32100b0a12591721f9bfb9c8397cadac7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/columnlinks"
    i48885ca869aa49bb491aee86881753531e856a31dfe44ed73d78ec827a19580a "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/columnpositions"
    i6b3f6dac5da19803c9c9bae3700dc771dc5acb71fcce181487b5a51e0514130c "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/base"
    i74b1b25e89fc5e6c7f2400e6c381122683ba4ee07b0d10423eb379a087935e39 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/publish"
    i77bc0ce7c808a9aad0fa44a31f405e772cea4d26f119fcb3df329e06d3ef8a8f "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/unpublish"
    i91e9bed4c1599a1967322c3d0ca6fb756e1f9a45f662a72215b53971eb86b392 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/associatewithhubsites"
    ic0ccf81c4a33675527d412b17de42c755aeb1005c9bd432d9b11f30c6dba7686 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/basetypes"
    ic2618dd7e9ce27b25d7aa90580585e11db3e063952456fbf8b989c1dffd64991 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/copytodefaultcontentlocation"
    ic7d6d99ce726befd6cbeeabdd1892aae108a8089f497e3c147e6c93bc26c57f3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/columns"
    ie341bfa543b51d3047340f895ff26f30c5028349182908d36cd81778c67c1cd0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/ispublished"
    i5cc5bc9268317b37d92eab0a20737056353e76ab632fccc11a9d57b16392d9a0 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/basetypes/item"
    ic46149ee873d74983b0c6300dbbe07dc89518f4464d205f6e08b1f1c8efb9aa3 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/columns/item"
    idea2279221f49c916c62603b4edb2acc56c5d8004d66fd5f4f02742d0638991c "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/columnpositions/item"
    ifc7c800236d94ba4b493a8ea1fa70ae6a520b4d0b98c2dd1845e35bca8687be7 "github.com/microsoftgraph/msgraph-sdk-go/users/item/drives/item/list/contenttypes/item/columnlinks/item"
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i91e9bed4c1599a1967322c3d0ca6fb756e1f9a45f662a72215b53971eb86b392.AssociateWithHubSitesRequestBuilder) {
    return i91e9bed4c1599a1967322c3d0ca6fb756e1f9a45f662a72215b53971eb86b392.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Base the base property
func (m *ContentTypeItemRequestBuilder) Base()(*i6b3f6dac5da19803c9c9bae3700dc771dc5acb71fcce181487b5a51e0514130c.BaseRequestBuilder) {
    return i6b3f6dac5da19803c9c9bae3700dc771dc5acb71fcce181487b5a51e0514130c.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypes the baseTypes property
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*ic0ccf81c4a33675527d412b17de42c755aeb1005c9bd432d9b11f30c6dba7686.BaseTypesRequestBuilder) {
    return ic0ccf81c4a33675527d412b17de42c755aeb1005c9bd432d9b11f30c6dba7686.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.list.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i5cc5bc9268317b37d92eab0a20737056353e76ab632fccc11a9d57b16392d9a0.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return i5cc5bc9268317b37d92eab0a20737056353e76ab632fccc11a9d57b16392d9a0.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnLinks the columnLinks property
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i3749973de34d8060e52ebe630f753ac32100b0a12591721f9bfb9c8397cadac7.ColumnLinksRequestBuilder) {
    return i3749973de34d8060e52ebe630f753ac32100b0a12591721f9bfb9c8397cadac7.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.list.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*ifc7c800236d94ba4b493a8ea1fa70ae6a520b4d0b98c2dd1845e35bca8687be7.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return ifc7c800236d94ba4b493a8ea1fa70ae6a520b4d0b98c2dd1845e35bca8687be7.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ColumnPositions the columnPositions property
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i48885ca869aa49bb491aee86881753531e856a31dfe44ed73d78ec827a19580a.ColumnPositionsRequestBuilder) {
    return i48885ca869aa49bb491aee86881753531e856a31dfe44ed73d78ec827a19580a.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.list.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*idea2279221f49c916c62603b4edb2acc56c5d8004d66fd5f4f02742d0638991c.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return idea2279221f49c916c62603b4edb2acc56c5d8004d66fd5f4f02742d0638991c.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Columns the columns property
func (m *ContentTypeItemRequestBuilder) Columns()(*ic7d6d99ce726befd6cbeeabdd1892aae108a8089f497e3c147e6c93bc26c57f3.ColumnsRequestBuilder) {
    return ic7d6d99ce726befd6cbeeabdd1892aae108a8089f497e3c147e6c93bc26c57f3.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.users.item.drives.item.list.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*ic46149ee873d74983b0c6300dbbe07dc89518f4464d205f6e08b1f1c8efb9aa3.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return ic46149ee873d74983b0c6300dbbe07dc89518f4464d205f6e08b1f1c8efb9aa3.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/users/{user_id}/drives/{drive_id}/list/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*ic2618dd7e9ce27b25d7aa90580585e11db3e063952456fbf8b989c1dffd64991.CopyToDefaultContentLocationRequestBuilder) {
    return ic2618dd7e9ce27b25d7aa90580585e11db3e063952456fbf8b989c1dffd64991.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for users
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
// CreatePatchRequestInformation update the navigation property contentTypes in users
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
// Delete delete navigation property contentTypes for users
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
func (m *ContentTypeItemRequestBuilder) IsPublished()(*ie341bfa543b51d3047340f895ff26f30c5028349182908d36cd81778c67c1cd0.IsPublishedRequestBuilder) {
    return ie341bfa543b51d3047340f895ff26f30c5028349182908d36cd81778c67c1cd0.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in users
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
func (m *ContentTypeItemRequestBuilder) Publish()(*i74b1b25e89fc5e6c7f2400e6c381122683ba4ee07b0d10423eb379a087935e39.PublishRequestBuilder) {
    return i74b1b25e89fc5e6c7f2400e6c381122683ba4ee07b0d10423eb379a087935e39.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Unpublish the unpublish property
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i77bc0ce7c808a9aad0fa44a31f405e772cea4d26f119fcb3df329e06d3ef8a8f.UnpublishRequestBuilder) {
    return i77bc0ce7c808a9aad0fa44a31f405e772cea4d26f119fcb3df329e06d3ef8a8f.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
