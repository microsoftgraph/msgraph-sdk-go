package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
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
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ContentTypeItemRequestBuilderDeleteOptions options for Delete
type ContentTypeItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ContentTypeItemRequestBuilderGetOptions options for Get
type ContentTypeItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContentTypeItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
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
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentTypeable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i91e9bed4c1599a1967322c3d0ca6fb756e1f9a45f662a72215b53971eb86b392.AssociateWithHubSitesRequestBuilder) {
    return i91e9bed4c1599a1967322c3d0ca6fb756e1f9a45f662a72215b53971eb86b392.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*i6b3f6dac5da19803c9c9bae3700dc771dc5acb71fcce181487b5a51e0514130c.BaseRequestBuilder) {
    return i6b3f6dac5da19803c9c9bae3700dc771dc5acb71fcce181487b5a51e0514130c.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
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
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
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
func NewContentTypeItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*ic2618dd7e9ce27b25d7aa90580585e11db3e063952456fbf8b989c1dffd64991.CopyToDefaultContentLocationRequestBuilder) {
    return ic2618dd7e9ce27b25d7aa90580585e11db3e063952456fbf8b989c1dffd64991.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for users
func (m *ContentTypeItemRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreateGetRequestInformation the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) CreateGetRequestInformation(options *ContentTypeItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        requestInfo.AddQueryParameters(*(options.Q))
    }
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update the navigation property contentTypes in users
func (m *ContentTypeItemRequestBuilder) CreatePatchRequestInformation(options *ContentTypeItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.H != nil {
        requestInfo.Headers = options.H
    }
    if options != nil && len(options.O) != 0 {
        err := requestInfo.AddRequestOptions(options.O...)
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
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Get the collection of content types present in this list.
func (m *ContentTypeItemRequestBuilder) Get(options *ContentTypeItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentTypeable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateContentTypeFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentTypeable), nil
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
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeItemRequestBuilder) Publish()(*i74b1b25e89fc5e6c7f2400e6c381122683ba4ee07b0d10423eb379a087935e39.PublishRequestBuilder) {
    return i74b1b25e89fc5e6c7f2400e6c381122683ba4ee07b0d10423eb379a087935e39.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i77bc0ce7c808a9aad0fa44a31f405e772cea4d26f119fcb3df329e06d3ef8a8f.UnpublishRequestBuilder) {
    return i77bc0ce7c808a9aad0fa44a31f405e772cea4d26f119fcb3df329e06d3ef8a8f.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
