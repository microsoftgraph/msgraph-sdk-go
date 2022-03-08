package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i00b23f7f86a982af8a15ee36120329e64b13d88be4f288d635f6ae65ae966124 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/list"
    i05fb43f9021a0913de359c2cfdf1fcad6b4de0310980145d3c6276c460270205 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/items"
    i2cbdf6d3a4841b26281917b6289b085e9cd675537f09928d44582bcd4533f5d2 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/root"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4c63c9b02f3ed101ac621ec88441d97587c19d3efeb627897c02cbe1ca7c6af4 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/listitem"
    i8c15d6975df59e7bff68ceaff8e2e7e5fb860064ce0c3fe87881dfa452f8d0f3 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/site"
    ia21810eb35790209720396d6baefa7bad436f4af7de39d51a975ac8ed17d9b90 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/permission"
    ic21e3d9f3eaf63bbdce62fd14ebed43e30e12830422fa068b6c26780d7bb64a6 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/driveitem"
    i645f1e1a516727ccf24f1ba68b8f4137ee52e397b6e58438bccce4f0cb3d74b5 "github.com/microsoftgraph/msgraph-sdk-go/shares/item/items/item"
)

// SharedDriveItemItemRequestBuilder provides operations to manage the collection of sharedDriveItem entities.
type SharedDriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// SharedDriveItemItemRequestBuilderDeleteOptions options for Delete
type SharedDriveItemItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SharedDriveItemItemRequestBuilderGetOptions options for Get
type SharedDriveItemItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *SharedDriveItemItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// SharedDriveItemItemRequestBuilderGetQueryParameters get entity from shares by key
type SharedDriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// SharedDriveItemItemRequestBuilderPatchOptions options for Patch
type SharedDriveItemItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SharedDriveItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// NewSharedDriveItemItemRequestBuilderInternal instantiates a new SharedDriveItemItemRequestBuilder and sets the default values.
func NewSharedDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SharedDriveItemItemRequestBuilder) {
    m := &SharedDriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/shares/{sharedDriveItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewSharedDriveItemItemRequestBuilder instantiates a new SharedDriveItemItemRequestBuilder and sets the default values.
func NewSharedDriveItemItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*SharedDriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewSharedDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from shares
func (m *SharedDriveItemItemRequestBuilder) CreateDeleteRequestInformation(options *SharedDriveItemItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from shares by key
func (m *SharedDriveItemItemRequestBuilder) CreateGetRequestInformation(options *SharedDriveItemItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in shares
func (m *SharedDriveItemItemRequestBuilder) CreatePatchRequestInformation(options *SharedDriveItemItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from shares
func (m *SharedDriveItemItemRequestBuilder) Delete(options *SharedDriveItemItemRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *SharedDriveItemItemRequestBuilder) DriveItem()(*ic21e3d9f3eaf63bbdce62fd14ebed43e30e12830422fa068b6c26780d7bb64a6.DriveItemRequestBuilder) {
    return ic21e3d9f3eaf63bbdce62fd14ebed43e30e12830422fa068b6c26780d7bb64a6.NewDriveItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Get get entity from shares by key
func (m *SharedDriveItemItemRequestBuilder) Get(options *SharedDriveItemItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SharedDriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateSharedDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.SharedDriveItemable), nil
}
func (m *SharedDriveItemItemRequestBuilder) Items()(*i05fb43f9021a0913de359c2cfdf1fcad6b4de0310980145d3c6276c460270205.ItemsRequestBuilder) {
    return i05fb43f9021a0913de359c2cfdf1fcad6b4de0310980145d3c6276c460270205.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.shares.item.items.item collection
func (m *SharedDriveItemItemRequestBuilder) ItemsById(id string)(*i645f1e1a516727ccf24f1ba68b8f4137ee52e397b6e58438bccce4f0cb3d74b5.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i645f1e1a516727ccf24f1ba68b8f4137ee52e397b6e58438bccce4f0cb3d74b5.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *SharedDriveItemItemRequestBuilder) List()(*i00b23f7f86a982af8a15ee36120329e64b13d88be4f288d635f6ae65ae966124.ListRequestBuilder) {
    return i00b23f7f86a982af8a15ee36120329e64b13d88be4f288d635f6ae65ae966124.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SharedDriveItemItemRequestBuilder) ListItem()(*i4c63c9b02f3ed101ac621ec88441d97587c19d3efeb627897c02cbe1ca7c6af4.ListItemRequestBuilder) {
    return i4c63c9b02f3ed101ac621ec88441d97587c19d3efeb627897c02cbe1ca7c6af4.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in shares
func (m *SharedDriveItemItemRequestBuilder) Patch(options *SharedDriveItemItemRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *SharedDriveItemItemRequestBuilder) Permission()(*ia21810eb35790209720396d6baefa7bad436f4af7de39d51a975ac8ed17d9b90.PermissionRequestBuilder) {
    return ia21810eb35790209720396d6baefa7bad436f4af7de39d51a975ac8ed17d9b90.NewPermissionRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SharedDriveItemItemRequestBuilder) Root()(*i2cbdf6d3a4841b26281917b6289b085e9cd675537f09928d44582bcd4533f5d2.RootRequestBuilder) {
    return i2cbdf6d3a4841b26281917b6289b085e9cd675537f09928d44582bcd4533f5d2.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *SharedDriveItemItemRequestBuilder) Site()(*i8c15d6975df59e7bff68ceaff8e2e7e5fb860064ce0c3fe87881dfa452f8d0f3.SiteRequestBuilder) {
    return i8c15d6975df59e7bff68ceaff8e2e7e5fb860064ce0c3fe87881dfa452f8d0f3.NewSiteRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
