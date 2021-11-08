package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i00b23f7f86a982af8a15ee36120329e64b13d88be4f288d635f6ae65ae966124 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i05fb43f9021a0913de359c2cfdf1fcad6b4de0310980145d3c6276c460270205 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/items"
    i2cbdf6d3a4841b26281917b6289b085e9cd675537f09928d44582bcd4533f5d2 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/root"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4c63c9b02f3ed101ac621ec88441d97587c19d3efeb627897c02cbe1ca7c6af4 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/listitem"
    i8c15d6975df59e7bff68ceaff8e2e7e5fb860064ce0c3fe87881dfa452f8d0f3 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/site"
    ia21810eb35790209720396d6baefa7bad436f4af7de39d51a975ac8ed17d9b90 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/permission"
    ic21e3d9f3eaf63bbdce62fd14ebed43e30e12830422fa068b6c26780d7bb64a6 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/driveitem"
    i645f1e1a516727ccf24f1ba68b8f4137ee52e397b6e58438bccce4f0cb3d74b5 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/items/item"
)

// Builds and executes requests for operations under \shares\{sharedDriveItem-id}
type SharedDriveItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Delete
type SharedDriveItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Options for Get
type SharedDriveItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SharedDriveItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get entity from shares by key
type SharedDriveItemRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type SharedDriveItemRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SharedDriveItem;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Instantiates a new SharedDriveItemRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSharedDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SharedDriveItemRequestBuilder) {
    m := &SharedDriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new SharedDriveItemRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewSharedDriveItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SharedDriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharedDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// Delete entity from shares
// Parameters:
//  - options : Options for the request
func (m *SharedDriveItemRequestBuilder) CreateDeleteRequestInformation(options *SharedDriveItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Get entity from shares by key
// Parameters:
//  - options : Options for the request
func (m *SharedDriveItemRequestBuilder) CreateGetRequestInformation(options *SharedDriveItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Update entity in shares
// Parameters:
//  - options : Options for the request
func (m *SharedDriveItemRequestBuilder) CreatePatchRequestInformation(options *SharedDriveItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete entity from shares
// Parameters:
//  - options : Options for the request
func (m *SharedDriveItemRequestBuilder) Delete(options *SharedDriveItemRequestBuilderDeleteOptions)(error) {
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
func (m *SharedDriveItemRequestBuilder) DriveItem()(*ic21e3d9f3eaf63bbdce62fd14ebed43e30e12830422fa068b6c26780d7bb64a6.DriveItemRequestBuilder) {
    return ic21e3d9f3eaf63bbdce62fd14ebed43e30e12830422fa068b6c26780d7bb64a6.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get entity from shares by key
// Parameters:
//  - options : Options for the request
func (m *SharedDriveItemRequestBuilder) Get(options *SharedDriveItemRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SharedDriveItem, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewSharedDriveItem() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SharedDriveItem), nil
}
func (m *SharedDriveItemRequestBuilder) Items()(*i05fb43f9021a0913de359c2cfdf1fcad6b4de0310980145d3c6276c460270205.ItemsRequestBuilder) {
    return i05fb43f9021a0913de359c2cfdf1fcad6b4de0310980145d3c6276c460270205.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go.shares.item.items.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *SharedDriveItemRequestBuilder) ItemsById(id string)(*i645f1e1a516727ccf24f1ba68b8f4137ee52e397b6e58438bccce4f0cb3d74b5.DriveItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i645f1e1a516727ccf24f1ba68b8f4137ee52e397b6e58438bccce4f0cb3d74b5.NewDriveItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SharedDriveItemRequestBuilder) List()(*i00b23f7f86a982af8a15ee36120329e64b13d88be4f288d635f6ae65ae966124.ListRequestBuilder) {
    return i00b23f7f86a982af8a15ee36120329e64b13d88be4f288d635f6ae65ae966124.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SharedDriveItemRequestBuilder) ListItem()(*i4c63c9b02f3ed101ac621ec88441d97587c19d3efeb627897c02cbe1ca7c6af4.ListItemRequestBuilder) {
    return i4c63c9b02f3ed101ac621ec88441d97587c19d3efeb627897c02cbe1ca7c6af4.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Update entity in shares
// Parameters:
//  - options : Options for the request
func (m *SharedDriveItemRequestBuilder) Patch(options *SharedDriveItemRequestBuilderPatchOptions)(error) {
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
func (m *SharedDriveItemRequestBuilder) Permission()(*ia21810eb35790209720396d6baefa7bad436f4af7de39d51a975ac8ed17d9b90.PermissionRequestBuilder) {
    return ia21810eb35790209720396d6baefa7bad436f4af7de39d51a975ac8ed17d9b90.NewPermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SharedDriveItemRequestBuilder) Root()(*i2cbdf6d3a4841b26281917b6289b085e9cd675537f09928d44582bcd4533f5d2.RootRequestBuilder) {
    return i2cbdf6d3a4841b26281917b6289b085e9cd675537f09928d44582bcd4533f5d2.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SharedDriveItemRequestBuilder) Site()(*i8c15d6975df59e7bff68ceaff8e2e7e5fb860064ce0c3fe87881dfa452f8d0f3.SiteRequestBuilder) {
    return i8c15d6975df59e7bff68ceaff8e2e7e5fb860064ce0c3fe87881dfa452f8d0f3.NewSiteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
