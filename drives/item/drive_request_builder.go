package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i179944f90cbf5506a527dd2030c3cef54ae3ad4acbf50eb045881c9bd8549265 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/searchwithq"
    i2e66e6af0433d92d462a30e13b39a28fb7dee1f8ee90a449b4e31ca54ccec7f4 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/sharedwithme"
    i4205fb8ccb30dea17fb8ba72024db369db4bc1d5e093c337fcf552e72e3571c4 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i53698fe752ade34739ebb810a835811924df4d95e4b93d4e5a47d1449a7c9181 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/recent"
    iae8ee83127418dc4b9186e79ac698ea94d3e30eecc4269743c54c60cf314172b "github.com/microsoftgraph/msgraph-sdk-go/drives/item/items"
    ib2a4039148f5be30653b6dd294c5ab7aeb4ea50c5b8cb75c9ee23c4bdcd6b502 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root"
    ic658e9fc1679f0fe33489f42aae3b91cbf7d5382b5acf0f9fa109add5f3e4b2b "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special"
    id81d2362f07493ce3357d9624466a24f691c09e5f6d78806ecabf5b6f62f8ab8 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/list"
    iee46aa3df6c96b156c19a7dd468c660bb9bf053a22e578da189bbbaee28b0cdc "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following"
    i1ccf6e5f8701f6397ef714c413639aae46e1c415cb2263eeee0443f0155a743b "github.com/microsoftgraph/msgraph-sdk-go/drives/item/following/item"
    iafa8d5d3e1080d6766c6b71c082fdab35a828b1d0cc84e17edcfefd65ff9b021 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/items/item"
    ibb5bf0dffefc66fdb863261eb4a45d386965bbca07ce1fa075249ee0e414aba9 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/special/item"
    id7a7f8969d72823f194725db5971b674a1dd6dda550bc0b0d5bec20719ad9ee5 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item"
)

// Builds and executes requests for operations under \drives\{drive-id}
type DriveRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Get entity from drives by key
type DriveRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
func (m *DriveRequestBuilder) Bundles()(*i4205fb8ccb30dea17fb8ba72024db369db4bc1d5e093c337fcf552e72e3571c4.BundlesRequestBuilder) {
    return i4205fb8ccb30dea17fb8ba72024db369db4bc1d5e093c337fcf552e72e3571c4.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.bundles.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DriveRequestBuilder) BundlesById(id string)(*id7a7f8969d72823f194725db5971b674a1dd6dda550bc0b0d5bec20719ad9ee5.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return id7a7f8969d72823f194725db5971b674a1dd6dda550bc0b0d5bec20719ad9ee5.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new DriveRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDriveRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveRequestBuilder) {
    m := &DriveRequestBuilder{
    }
    m.urlTemplate = "https://graph.microsoft.com/v1.0/drives/{drive_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new DriveRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDriveRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete entity from drives
// Parameters:
//  - h : Request headers
//  - o : Request options
func (m *DriveRequestBuilder) CreateDeleteRequestInformation(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.DELETE
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Get entity from drives by key
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
func (m *DriveRequestBuilder) CreateGetRequestInformation(q func (value *DriveRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if q != nil {
        qParams := new(DriveRequestBuilderGetQueryParameters)
        err := q(qParams)
        if err != nil {
            return nil, err
        }
        err = qParams.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
    }
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Update entity in drives
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
func (m *DriveRequestBuilder) CreatePatchRequestInformation(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Drive, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if h != nil {
        err := h(requestInfo.Headers)
        if err != nil {
            return nil, err
        }
    }
    if o != nil {
        err := requestInfo.AddRequestOptions(o)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// Delete entity from drives
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *DriveRequestBuilder) Delete(h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
func (m *DriveRequestBuilder) Following()(*iee46aa3df6c96b156c19a7dd468c660bb9bf053a22e578da189bbbaee28b0cdc.FollowingRequestBuilder) {
    return iee46aa3df6c96b156c19a7dd468c660bb9bf053a22e578da189bbbaee28b0cdc.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.following.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DriveRequestBuilder) FollowingById(id string)(*i1ccf6e5f8701f6397ef714c413639aae46e1c415cb2263eeee0443f0155a743b.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i1ccf6e5f8701f6397ef714c413639aae46e1c415cb2263eeee0443f0155a743b.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get entity from drives by key
// Parameters:
//  - h : Request headers
//  - o : Request options
//  - q : Request query parameters
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *DriveRequestBuilder) Get(q func (value *DriveRequestBuilderGetQueryParameters) (err error), h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Drive, error) {
    requestInfo, err := m.CreateGetRequestInformation(q, h, o);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDrive() }, responseHandler)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Drive), nil
}
func (m *DriveRequestBuilder) Items()(*iae8ee83127418dc4b9186e79ac698ea94d3e30eecc4269743c54c60cf314172b.ItemsRequestBuilder) {
    return iae8ee83127418dc4b9186e79ac698ea94d3e30eecc4269743c54c60cf314172b.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.items.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DriveRequestBuilder) ItemsById(id string)(*iafa8d5d3e1080d6766c6b71c082fdab35a828b1d0cc84e17edcfefd65ff9b021.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return iafa8d5d3e1080d6766c6b71c082fdab35a828b1d0cc84e17edcfefd65ff9b021.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveRequestBuilder) List()(*id81d2362f07493ce3357d9624466a24f691c09e5f6d78806ecabf5b6f62f8ab8.ListRequestBuilder) {
    return id81d2362f07493ce3357d9624466a24f691c09e5f6d78806ecabf5b6f62f8ab8.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update entity in drives
// Parameters:
//  - body : 
//  - h : Request headers
//  - o : Request options
//  - responseHandler : Response handler to use in place of the default response handling provided by the core service
func (m *DriveRequestBuilder) Patch(body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Drive, h func (value map[string]string) (err error), o []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption, responseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(body, h, o);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, responseHandler)
    if err != nil {
        return err
    }
    return nil
}
// Builds and executes requests for operations under \drives\{drive-id}\microsoft.graph.recent()
func (m *DriveRequestBuilder) Recent()(*i53698fe752ade34739ebb810a835811924df4d95e4b93d4e5a47d1449a7c9181.RecentRequestBuilder) {
    return i53698fe752ade34739ebb810a835811924df4d95e4b93d4e5a47d1449a7c9181.NewRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) Root()(*ib2a4039148f5be30653b6dd294c5ab7aeb4ea50c5b8cb75c9ee23c4bdcd6b502.RootRequestBuilder) {
    return ib2a4039148f5be30653b6dd294c5ab7aeb4ea50c5b8cb75c9ee23c4bdcd6b502.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Builds and executes requests for operations under \drives\{drive-id}\microsoft.graph.search(q='{q}')
// Parameters:
//  - q : Usage: q={q}
func (m *DriveRequestBuilder) SearchWithQ(q *string)(*i179944f90cbf5506a527dd2030c3cef54ae3ad4acbf50eb045881c9bd8549265.SearchWithQRequestBuilder) {
    return i179944f90cbf5506a527dd2030c3cef54ae3ad4acbf50eb045881c9bd8549265.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// Builds and executes requests for operations under \drives\{drive-id}\microsoft.graph.sharedWithMe()
func (m *DriveRequestBuilder) SharedWithMe()(*i2e66e6af0433d92d462a30e13b39a28fb7dee1f8ee90a449b4e31ca54ccec7f4.SharedWithMeRequestBuilder) {
    return i2e66e6af0433d92d462a30e13b39a28fb7dee1f8ee90a449b4e31ca54ccec7f4.NewSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) Special()(*ic658e9fc1679f0fe33489f42aae3b91cbf7d5382b5acf0f9fa109add5f3e4b2b.SpecialRequestBuilder) {
    return ic658e9fc1679f0fe33489f42aae3b91cbf7d5382b5acf0f9fa109add5f3e4b2b.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.special.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DriveRequestBuilder) SpecialById(id string)(*ibb5bf0dffefc66fdb863261eb4a45d386965bbca07ce1fa075249ee0e414aba9.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return ibb5bf0dffefc66fdb863261eb4a45d386965bbca07ce1fa075249ee0e414aba9.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
