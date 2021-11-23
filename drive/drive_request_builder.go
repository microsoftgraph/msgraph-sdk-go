package drive

import (
    i33074db2b8fa9be0b6a9d94cf471e252f008c27ab9bb5a8e0ca53ac0f0210cf7 "github.com/microsoftgraph/msgraph-sdk-go/drive/items"
    i56f7e69bbb7d6f10a3721658326d4d734b2ed1edb084522652d8117a07ccaaff "github.com/microsoftgraph/msgraph-sdk-go/drive/sharedwithme"
    i62e01495b26c3ef220d18dd4a2ec05f79ddb93094e7d5a328fdfc46f27cef7b6 "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles"
    i6540397555b3c095b7c6c5d4fcae5bef89c2baee0dcf1f8bba0233611b3d3f95 "github.com/microsoftgraph/msgraph-sdk-go/drive/list"
    i9316c3bd03b4192d3d91a742a5f05f6406c00c15b3b5de33a8e843dac77d53ad "github.com/microsoftgraph/msgraph-sdk-go/drive/following"
    ia82c9355f6f8d356282e506ef46bfccb28fc443731e9a05eab283791a6de8220 "github.com/microsoftgraph/msgraph-sdk-go/drive/special"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie13063c1cadad44d23aa1c8d0fa86d8d1520f918579ef541dec1ccdb6ed9a220 "github.com/microsoftgraph/msgraph-sdk-go/drive/searchwithq"
    ie72d87938961cb52922de029a87b3e51e1f879639a8ade8b937c65900a85220c "github.com/microsoftgraph/msgraph-sdk-go/drive/recent"
    ie85bb22170659be58e3c60d2b8b86f9a5d711197d7702af8a50ef22c9883d76a "github.com/microsoftgraph/msgraph-sdk-go/drive/root"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i196a6eae213268ad21678594954acf3ccd7f039d9e67578273af389beab82ac2 "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item"
    i3087ffea8c433c599359b8347be323c02e65364584d9de678689a61cc31d9d68 "github.com/microsoftgraph/msgraph-sdk-go/drive/special/item"
    i442efb2ad5fc2ef8bc831b0f78ab7125b2330176f4baf9ac638a842e8b683dc9 "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    ibd7a7a1ce01856c78689c5524ff307585d4847757b7dc445f97bff4e3fdb5a1b "github.com/microsoftgraph/msgraph-sdk-go/drive/items/item"
)

// DriveRequestBuilder builds and executes requests for operations under \drive
type DriveRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DriveRequestBuilderGetOptions options for Get
type DriveRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DriveRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveRequestBuilderGetQueryParameters get drive
type DriveRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// DriveRequestBuilderPatchOptions options for Patch
type DriveRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Drive;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DriveRequestBuilder) Bundles()(*i62e01495b26c3ef220d18dd4a2ec05f79ddb93094e7d5a328fdfc46f27cef7b6.BundlesRequestBuilder) {
    return i62e01495b26c3ef220d18dd4a2ec05f79ddb93094e7d5a328fdfc46f27cef7b6.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.bundles.item collection
func (m *DriveRequestBuilder) BundlesById(id string)(*i196a6eae213268ad21678594954acf3ccd7f039d9e67578273af389beab82ac2.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i196a6eae213268ad21678594954acf3ccd7f039d9e67578273af389beab82ac2.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveRequestBuilderInternal instantiates a new DriveRequestBuilder and sets the default values.
func NewDriveRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveRequestBuilder) {
    m := &DriveRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveRequestBuilder instantiates a new DriveRequestBuilder and sets the default values.
func NewDriveRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get drive
func (m *DriveRequestBuilder) CreateGetRequestInformation(options *DriveRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update drive
func (m *DriveRequestBuilder) CreatePatchRequestInformation(options *DriveRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DriveRequestBuilder) Following()(*i9316c3bd03b4192d3d91a742a5f05f6406c00c15b3b5de33a8e843dac77d53ad.FollowingRequestBuilder) {
    return i9316c3bd03b4192d3d91a742a5f05f6406c00c15b3b5de33a8e843dac77d53ad.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.following.item collection
func (m *DriveRequestBuilder) FollowingById(id string)(*i442efb2ad5fc2ef8bc831b0f78ab7125b2330176f4baf9ac638a842e8b683dc9.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i442efb2ad5fc2ef8bc831b0f78ab7125b2330176f4baf9ac638a842e8b683dc9.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get drive
func (m *DriveRequestBuilder) Get(options *DriveRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Drive, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDrive() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Drive), nil
}
func (m *DriveRequestBuilder) Items()(*i33074db2b8fa9be0b6a9d94cf471e252f008c27ab9bb5a8e0ca53ac0f0210cf7.ItemsRequestBuilder) {
    return i33074db2b8fa9be0b6a9d94cf471e252f008c27ab9bb5a8e0ca53ac0f0210cf7.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.items.item collection
func (m *DriveRequestBuilder) ItemsById(id string)(*ibd7a7a1ce01856c78689c5524ff307585d4847757b7dc445f97bff4e3fdb5a1b.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return ibd7a7a1ce01856c78689c5524ff307585d4847757b7dc445f97bff4e3fdb5a1b.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveRequestBuilder) List()(*i6540397555b3c095b7c6c5d4fcae5bef89c2baee0dcf1f8bba0233611b3d3f95.ListRequestBuilder) {
    return i6540397555b3c095b7c6c5d4fcae5bef89c2baee0dcf1f8bba0233611b3d3f95.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update drive
func (m *DriveRequestBuilder) Patch(options *DriveRequestBuilderPatchOptions)(error) {
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
// Recent builds and executes requests for operations under \drive\microsoft.graph.recent()
func (m *DriveRequestBuilder) Recent()(*ie72d87938961cb52922de029a87b3e51e1f879639a8ade8b937c65900a85220c.RecentRequestBuilder) {
    return ie72d87938961cb52922de029a87b3e51e1f879639a8ade8b937c65900a85220c.NewRecentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) Root()(*ie85bb22170659be58e3c60d2b8b86f9a5d711197d7702af8a50ef22c9883d76a.RootRequestBuilder) {
    return ie85bb22170659be58e3c60d2b8b86f9a5d711197d7702af8a50ef22c9883d76a.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SearchWithQ builds and executes requests for operations under \drive\microsoft.graph.search(q='{q}')
func (m *DriveRequestBuilder) SearchWithQ(q *string)(*ie13063c1cadad44d23aa1c8d0fa86d8d1520f918579ef541dec1ccdb6ed9a220.SearchWithQRequestBuilder) {
    return ie13063c1cadad44d23aa1c8d0fa86d8d1520f918579ef541dec1ccdb6ed9a220.NewSearchWithQRequestBuilderInternal(m.pathParameters, m.requestAdapter, q);
}
// SharedWithMe builds and executes requests for operations under \drive\microsoft.graph.sharedWithMe()
func (m *DriveRequestBuilder) SharedWithMe()(*i56f7e69bbb7d6f10a3721658326d4d734b2ed1edb084522652d8117a07ccaaff.SharedWithMeRequestBuilder) {
    return i56f7e69bbb7d6f10a3721658326d4d734b2ed1edb084522652d8117a07ccaaff.NewSharedWithMeRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveRequestBuilder) Special()(*ia82c9355f6f8d356282e506ef46bfccb28fc443731e9a05eab283791a6de8220.SpecialRequestBuilder) {
    return ia82c9355f6f8d356282e506ef46bfccb28fc443731e9a05eab283791a6de8220.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.special.item collection
func (m *DriveRequestBuilder) SpecialById(id string)(*i3087ffea8c433c599359b8347be323c02e65364584d9de678689a61cc31d9d68.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i3087ffea8c433c599359b8347be323c02e65364584d9de678689a61cc31d9d68.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
