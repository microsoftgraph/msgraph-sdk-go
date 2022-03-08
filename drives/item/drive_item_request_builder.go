package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
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

// DriveItemRequestBuilder provides operations to manage the collection of drive entities.
type DriveItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DriveItemRequestBuilderDeleteOptions options for Delete
type DriveItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemRequestBuilderGetOptions options for Get
type DriveItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DriveItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemRequestBuilderGetQueryParameters get entity from drives by key
type DriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DriveItemRequestBuilderPatchOptions options for Patch
type DriveItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Driveable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DriveItemRequestBuilder) Bundles()(*i4205fb8ccb30dea17fb8ba72024db369db4bc1d5e093c337fcf552e72e3571c4.BundlesRequestBuilder) {
    return i4205fb8ccb30dea17fb8ba72024db369db4bc1d5e093c337fcf552e72e3571c4.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.bundles.item collection
func (m *DriveItemRequestBuilder) BundlesById(id string)(*id7a7f8969d72823f194725db5971b674a1dd6dda550bc0b0d5bec20719ad9ee5.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return id7a7f8969d72823f194725db5971b674a1dd6dda550bc0b0d5bec20719ad9ee5.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemRequestBuilder) {
    m := &DriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemRequestBuilder instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from drives
func (m *DriveItemRequestBuilder) CreateDeleteRequestInformation(options *DriveItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from drives by key
func (m *DriveItemRequestBuilder) CreateGetRequestInformation(options *DriveItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in drives
func (m *DriveItemRequestBuilder) CreatePatchRequestInformation(options *DriveItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from drives
func (m *DriveItemRequestBuilder) Delete(options *DriveItemRequestBuilderDeleteOptions)(error) {
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
func (m *DriveItemRequestBuilder) Following()(*iee46aa3df6c96b156c19a7dd468c660bb9bf053a22e578da189bbbaee28b0cdc.FollowingRequestBuilder) {
    return iee46aa3df6c96b156c19a7dd468c660bb9bf053a22e578da189bbbaee28b0cdc.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.following.item collection
func (m *DriveItemRequestBuilder) FollowingById(id string)(*i1ccf6e5f8701f6397ef714c413639aae46e1c415cb2263eeee0443f0155a743b.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i1ccf6e5f8701f6397ef714c413639aae46e1c415cb2263eeee0443f0155a743b.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from drives by key
func (m *DriveItemRequestBuilder) Get(options *DriveItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Driveable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDriveFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Driveable), nil
}
func (m *DriveItemRequestBuilder) Items()(*iae8ee83127418dc4b9186e79ac698ea94d3e30eecc4269743c54c60cf314172b.ItemsRequestBuilder) {
    return iae8ee83127418dc4b9186e79ac698ea94d3e30eecc4269743c54c60cf314172b.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.items.item collection
func (m *DriveItemRequestBuilder) ItemsById(id string)(*iafa8d5d3e1080d6766c6b71c082fdab35a828b1d0cc84e17edcfefd65ff9b021.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return iafa8d5d3e1080d6766c6b71c082fdab35a828b1d0cc84e17edcfefd65ff9b021.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) List()(*id81d2362f07493ce3357d9624466a24f691c09e5f6d78806ecabf5b6f62f8ab8.ListRequestBuilder) {
    return id81d2362f07493ce3357d9624466a24f691c09e5f6d78806ecabf5b6f62f8ab8.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in drives
func (m *DriveItemRequestBuilder) Patch(options *DriveItemRequestBuilderPatchOptions)(error) {
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
// Recent provides operations to call the recent method.
func (m *DriveItemRequestBuilder) Recent()(*i53698fe752ade34739ebb810a835811924df4d95e4b93d4e5a47d1449a7c9181.RecentRequestBuilder) {
    return i53698fe752ade34739ebb810a835811924df4d95e4b93d4e5a47d1449a7c9181.NewRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Root()(*ib2a4039148f5be30653b6dd294c5ab7aeb4ea50c5b8cb75c9ee23c4bdcd6b502.RootRequestBuilder) {
    return ib2a4039148f5be30653b6dd294c5ab7aeb4ea50c5b8cb75c9ee23c4bdcd6b502.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ provides operations to call the search method.
func (m *DriveItemRequestBuilder) SearchWithQ(q *string)(*i179944f90cbf5506a527dd2030c3cef54ae3ad4acbf50eb045881c9bd8549265.SearchWithQRequestBuilder) {
    return i179944f90cbf5506a527dd2030c3cef54ae3ad4acbf50eb045881c9bd8549265.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// SharedWithMe provides operations to call the sharedWithMe method.
func (m *DriveItemRequestBuilder) SharedWithMe()(*i2e66e6af0433d92d462a30e13b39a28fb7dee1f8ee90a449b4e31ca54ccec7f4.SharedWithMeRequestBuilder) {
    return i2e66e6af0433d92d462a30e13b39a28fb7dee1f8ee90a449b4e31ca54ccec7f4.NewSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Special()(*ic658e9fc1679f0fe33489f42aae3b91cbf7d5382b5acf0f9fa109add5f3e4b2b.SpecialRequestBuilder) {
    return ic658e9fc1679f0fe33489f42aae3b91cbf7d5382b5acf0f9fa109add5f3e4b2b.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.special.item collection
func (m *DriveItemRequestBuilder) SpecialById(id string)(*ibb5bf0dffefc66fdb863261eb4a45d386965bbca07ce1fa075249ee0e414aba9.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return ibb5bf0dffefc66fdb863261eb4a45d386965bbca07ce1fa075249ee0e414aba9.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
