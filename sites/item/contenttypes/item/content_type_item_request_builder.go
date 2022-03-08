package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i05aa6bb50d61c9a91685ccd79551fe45c7fd5a6f41b2f9052961feb2dc394eb0 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/associatewithhubsites"
    i0cf1daad8519139f1c1dddb4eb3f61c0e4ec6e71759a998ee5a46fd4b2468890 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/columns"
    i2bd83908df726274d2c1ed9ae647d9d262c28eb0503cd31251764fddce791655 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/columnpositions"
    i2e3ff723a01a654065641276ad9c59facdf69c3d223421c5e324df2133e20d2d "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/columnlinks"
    i43a81088f78dc3b0d7c4c4386ce5e5fe9e9adeeeed6c450fa8819629e5edd1d2 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/publish"
    i44ccd150eeae5d04e5afc727774f113c375506bb3a712e0dc54818ed2d2a7173 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/ispublished"
    i5e2cbca2e66013cc685d3b90b6f6620ddfc6a3ec8bd3560353ccf1bc0243fec6 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/unpublish"
    ib8daa131ed0a2d3b42d5cabdcf429cf168fa29e9581c9b560445404fc3855531 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/basetypes"
    id257be66b68c6ec4f94dd1b16a346a3628e3e2ff5283493c88ea924e3651e546 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/base"
    ie4cb5930f587e11ddf957e2716c252f45c2bae1f0cbb3eca003f83d9478ad44d "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/copytodefaultcontentlocation"
    i4e76b560d7ab063fee73b44a70701c5d35c8cd952639ebdad225d0fd1dd440b6 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/columns/item"
    i8031712eaa64ad6bcb9a08e5abc085cfafe42aef18be156e924188e8a86b8403 "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/basetypes/item"
    i9d5cdbf190630a4b53f62d7eb59d02a008abe57fc0685c8324eec7d57bd1e98d "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/columnpositions/item"
    if03de5dd741d9ed2d25876eea430ca53f3feda55c4d1a8e9ac5b875bac63413a "github.com/microsoftgraph/msgraph-sdk-go/sites/item/contenttypes/item/columnlinks/item"
)

// ContentTypeItemRequestBuilder provides operations to manage the contentTypes property of the microsoft.graph.site entity.
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
// ContentTypeItemRequestBuilderGetQueryParameters the collection of content types defined for this site.
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
func (m *ContentTypeItemRequestBuilder) AssociateWithHubSites()(*i05aa6bb50d61c9a91685ccd79551fe45c7fd5a6f41b2f9052961feb2dc394eb0.AssociateWithHubSitesRequestBuilder) {
    return i05aa6bb50d61c9a91685ccd79551fe45c7fd5a6f41b2f9052961feb2dc394eb0.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Base()(*id257be66b68c6ec4f94dd1b16a346a3628e3e2ff5283493c88ea924e3651e546.BaseRequestBuilder) {
    return id257be66b68c6ec4f94dd1b16a346a3628e3e2ff5283493c88ea924e3651e546.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) BaseTypes()(*ib8daa131ed0a2d3b42d5cabdcf429cf168fa29e9581c9b560445404fc3855531.BaseTypesRequestBuilder) {
    return ib8daa131ed0a2d3b42d5cabdcf429cf168fa29e9581c9b560445404fc3855531.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BaseTypesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.contentTypes.item.baseTypes.item collection
func (m *ContentTypeItemRequestBuilder) BaseTypesById(id string)(*i8031712eaa64ad6bcb9a08e5abc085cfafe42aef18be156e924188e8a86b8403.ContentTypeItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["contentType_id1"] = id
    }
    return i8031712eaa64ad6bcb9a08e5abc085cfafe42aef18be156e924188e8a86b8403.NewContentTypeItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnLinks()(*i2e3ff723a01a654065641276ad9c59facdf69c3d223421c5e324df2133e20d2d.ColumnLinksRequestBuilder) {
    return i2e3ff723a01a654065641276ad9c59facdf69c3d223421c5e324df2133e20d2d.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnLinksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.contentTypes.item.columnLinks.item collection
func (m *ContentTypeItemRequestBuilder) ColumnLinksById(id string)(*if03de5dd741d9ed2d25876eea430ca53f3feda55c4d1a8e9ac5b875bac63413a.ColumnLinkItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return if03de5dd741d9ed2d25876eea430ca53f3feda55c4d1a8e9ac5b875bac63413a.NewColumnLinkItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) ColumnPositions()(*i2bd83908df726274d2c1ed9ae647d9d262c28eb0503cd31251764fddce791655.ColumnPositionsRequestBuilder) {
    return i2bd83908df726274d2c1ed9ae647d9d262c28eb0503cd31251764fddce791655.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnPositionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.contentTypes.item.columnPositions.item collection
func (m *ContentTypeItemRequestBuilder) ColumnPositionsById(id string)(*i9d5cdbf190630a4b53f62d7eb59d02a008abe57fc0685c8324eec7d57bd1e98d.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i9d5cdbf190630a4b53f62d7eb59d02a008abe57fc0685c8324eec7d57bd1e98d.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Columns()(*i0cf1daad8519139f1c1dddb4eb3f61c0e4ec6e71759a998ee5a46fd4b2468890.ColumnsRequestBuilder) {
    return i0cf1daad8519139f1c1dddb4eb3f61c0e4ec6e71759a998ee5a46fd4b2468890.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ColumnsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.sites.item.contentTypes.item.columns.item collection
func (m *ContentTypeItemRequestBuilder) ColumnsById(id string)(*i4e76b560d7ab063fee73b44a70701c5d35c8cd952639ebdad225d0fd1dd440b6.ColumnDefinitionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i4e76b560d7ab063fee73b44a70701c5d35c8cd952639ebdad225d0fd1dd440b6.NewColumnDefinitionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewContentTypeItemRequestBuilderInternal instantiates a new ContentTypeItemRequestBuilder and sets the default values.
func NewContentTypeItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeItemRequestBuilder) {
    m := &ContentTypeItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/sites/{site_id}/contentTypes/{contentType_id}{?select,expand}";
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
func (m *ContentTypeItemRequestBuilder) CopyToDefaultContentLocation()(*ie4cb5930f587e11ddf957e2716c252f45c2bae1f0cbb3eca003f83d9478ad44d.CopyToDefaultContentLocationRequestBuilder) {
    return ie4cb5930f587e11ddf957e2716c252f45c2bae1f0cbb3eca003f83d9478ad44d.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property contentTypes for sites
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
// CreateGetRequestInformation the collection of content types defined for this site.
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
// CreatePatchRequestInformation update the navigation property contentTypes in sites
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
// Delete delete navigation property contentTypes for sites
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
// Get the collection of content types defined for this site.
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
func (m *ContentTypeItemRequestBuilder) IsPublished()(*i44ccd150eeae5d04e5afc727774f113c375506bb3a712e0dc54818ed2d2a7173.IsPublishedRequestBuilder) {
    return i44ccd150eeae5d04e5afc727774f113c375506bb3a712e0dc54818ed2d2a7173.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property contentTypes in sites
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
func (m *ContentTypeItemRequestBuilder) Publish()(*i43a81088f78dc3b0d7c4c4386ce5e5fe9e9adeeeed6c450fa8819629e5edd1d2.PublishRequestBuilder) {
    return i43a81088f78dc3b0d7c4c4386ce5e5fe9e9adeeeed6c450fa8819629e5edd1d2.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeItemRequestBuilder) Unpublish()(*i5e2cbca2e66013cc685d3b90b6f6620ddfc6a3ec8bd3560353ccf1bc0243fec6.UnpublishRequestBuilder) {
    return i5e2cbca2e66013cc685d3b90b6f6620ddfc6a3ec8bd3560353ccf1bc0243fec6.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
