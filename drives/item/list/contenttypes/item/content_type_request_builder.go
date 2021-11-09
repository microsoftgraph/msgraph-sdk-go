package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i05d9428439d35e52cb2ce2ba5c3c7be1cd010f18c1f354cb859cec016c46ad56 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/columns"
    i3060c755378bef956392441c325290f6ea178acc828301e16119a470f9192ffe "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/copytodefaultcontentlocation"
    i3b8ee80d094176753a5065e732b8fe371a09654d0bd329c3db9b51d7203fd853 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/publish"
    i43f44f7d18ca3faffb3453324604e7cc5c111263d5e22f2c445d29e2eb8218e9 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/base"
    i4f8ea4604387554b83af4225698949b2df687d40f43c743c6c46462ded20d3f9 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/ispublished"
    i582b9b1101bb3849bfd2fa23dfed26494d51af218a6e426e47df1c8cfbfa5554 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/basetypes"
    i7539bdc5f9b19a28873835c6b9f2f1c5a304eaf2c846cb1c025784c17b81241f "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/associatewithhubsites"
    ia9f43a3a505aa44f104fa23f977e28b34f2065aad158175d2503a0cc251a6f38 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/unpublish"
    ib1e7c6dac6ea9f3518585a39144c5b871a51fc839296cea92c1bd56467d4b734 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/columnpositions"
    ibb8519e802c21fcd2fd2594af481ca8371cfb5ded0d4b614dbcf185b47bf3d59 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/columnlinks"
    i47a65a6f94278006557555cc162ea76107be8a3f3deedeec2b3f8fd743a4e3ef "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/columns/item"
    i9d04d98cacfd97673fbb11e0191a5450a9dc4bb49e2376d572d9a0639d40abde "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list/contenttypes/item/columnlinks/item"
)

// Builds and executes requests for operations under \drives\{drive-id}\list\contentTypes\{contentType-id}
type ContentTypeRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type ContentTypeRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type ContentTypeRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ContentTypeRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// The collection of content types present in this list.
type ContentTypeRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type ContentTypeRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ContentTypeRequestBuilder) AssociateWithHubSites()(*i7539bdc5f9b19a28873835c6b9f2f1c5a304eaf2c846cb1c025784c17b81241f.AssociateWithHubSitesRequestBuilder) {
    return i7539bdc5f9b19a28873835c6b9f2f1c5a304eaf2c846cb1c025784c17b81241f.NewAssociateWithHubSitesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Base()(*i43f44f7d18ca3faffb3453324604e7cc5c111263d5e22f2c445d29e2eb8218e9.BaseRequestBuilder) {
    return i43f44f7d18ca3faffb3453324604e7cc5c111263d5e22f2c445d29e2eb8218e9.NewBaseRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) BaseTypes()(*i582b9b1101bb3849bfd2fa23dfed26494d51af218a6e426e47df1c8cfbfa5554.BaseTypesRequestBuilder) {
    return i582b9b1101bb3849bfd2fa23dfed26494d51af218a6e426e47df1c8cfbfa5554.NewBaseTypesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnLinks()(*ibb8519e802c21fcd2fd2594af481ca8371cfb5ded0d4b614dbcf185b47bf3d59.ColumnLinksRequestBuilder) {
    return ibb8519e802c21fcd2fd2594af481ca8371cfb5ded0d4b614dbcf185b47bf3d59.NewColumnLinksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.list.contentTypes.item.columnLinks.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContentTypeRequestBuilder) ColumnLinksById(id string)(*i9d04d98cacfd97673fbb11e0191a5450a9dc4bb49e2376d572d9a0639d40abde.ColumnLinkRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnLink_id"] = id
    }
    return i9d04d98cacfd97673fbb11e0191a5450a9dc4bb49e2376d572d9a0639d40abde.NewColumnLinkRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) ColumnPositions()(*ib1e7c6dac6ea9f3518585a39144c5b871a51fc839296cea92c1bd56467d4b734.ColumnPositionsRequestBuilder) {
    return ib1e7c6dac6ea9f3518585a39144c5b871a51fc839296cea92c1bd56467d4b734.NewColumnPositionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Columns()(*i05d9428439d35e52cb2ce2ba5c3c7be1cd010f18c1f354cb859cec016c46ad56.ColumnsRequestBuilder) {
    return i05d9428439d35e52cb2ce2ba5c3c7be1cd010f18c1f354cb859cec016c46ad56.NewColumnsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.list.contentTypes.item.columns.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *ContentTypeRequestBuilder) ColumnsById(id string)(*i47a65a6f94278006557555cc162ea76107be8a3f3deedeec2b3f8fd743a4e3ef.ColumnDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["columnDefinition_id"] = id
    }
    return i47a65a6f94278006557555cc162ea76107be8a3f3deedeec2b3f8fd743a4e3ef.NewColumnDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new ContentTypeRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContentTypeRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    m := &ContentTypeRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/list/contentTypes/{contentType_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new ContentTypeRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewContentTypeRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ContentTypeRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewContentTypeRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *ContentTypeRequestBuilder) CopyToDefaultContentLocation()(*i3060c755378bef956392441c325290f6ea178acc828301e16119a470f9192ffe.CopyToDefaultContentLocationRequestBuilder) {
    return i3060c755378bef956392441c325290f6ea178acc828301e16119a470f9192ffe.NewCopyToDefaultContentLocationRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of content types present in this list.
// Parameters:
//  - options : Options for the request
func (m *ContentTypeRequestBuilder) CreateDeleteRequestInformation(options *ContentTypeRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The collection of content types present in this list.
// Parameters:
//  - options : Options for the request
func (m *ContentTypeRequestBuilder) CreateGetRequestInformation(options *ContentTypeRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// The collection of content types present in this list.
// Parameters:
//  - options : Options for the request
func (m *ContentTypeRequestBuilder) CreatePatchRequestInformation(options *ContentTypeRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// The collection of content types present in this list.
// Parameters:
//  - options : Options for the request
func (m *ContentTypeRequestBuilder) Delete(options *ContentTypeRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
// The collection of content types present in this list.
// Parameters:
//  - options : Options for the request
func (m *ContentTypeRequestBuilder) Get(options *ContentTypeRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewContentType() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ContentType), nil
}
// Builds and executes requests for operations under \drives\{drive-id}\list\contentTypes\{contentType-id}\microsoft.graph.isPublished()
func (m *ContentTypeRequestBuilder) IsPublished()(*i4f8ea4604387554b83af4225698949b2df687d40f43c743c6c46462ded20d3f9.IsPublishedRequestBuilder) {
    return i4f8ea4604387554b83af4225698949b2df687d40f43c743c6c46462ded20d3f9.NewIsPublishedRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// The collection of content types present in this list.
// Parameters:
//  - options : Options for the request
func (m *ContentTypeRequestBuilder) Patch(options *ContentTypeRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *ContentTypeRequestBuilder) Publish()(*i3b8ee80d094176753a5065e732b8fe371a09654d0bd329c3db9b51d7203fd853.PublishRequestBuilder) {
    return i3b8ee80d094176753a5065e732b8fe371a09654d0bd329c3db9b51d7203fd853.NewPublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ContentTypeRequestBuilder) Unpublish()(*ia9f43a3a505aa44f104fa23f977e28b34f2065aad158175d2503a0cc251a6f38.UnpublishRequestBuilder) {
    return ia9f43a3a505aa44f104fa23f977e28b34f2065aad158175d2503a0cc251a6f38.NewUnpublishRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
