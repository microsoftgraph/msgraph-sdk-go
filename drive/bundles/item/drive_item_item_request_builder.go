package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i196a00f504cfcc6a9f69cef6ccb9b8a5412ee4bea114b8d6797615dc73aaa66c "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item/children"
    i61fb769ced9e59d2925fa83d5918db3045bd56b690fcfb92e92bd36e52a0453e "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item/analytics"
    ia05955119e155abf2b26e50ebc55da0f7085d6c22feb85011b6ac8d28350c267 "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item/versions"
    ibb083680698bfadb1b9ca3abdc030e7f9c822e8c1f6be855140f5f7242ade92e "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item/subscriptions"
    ic5a3c66f7d6eff455463fd0119a5e0cc8849b531238ed938071be216b6a3cd45 "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item/listitem"
    id33754ddf1402bf93373f2d10f1c524fc6f4f4dc25c30c3b8e0adcbed34faff7 "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item/thumbnails"
    id57c1f561fd7120e715af680ab9d67e3a72d92b54657bd3f3ac2432fcaec9d32 "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item/content"
    ie600c6a3a6a376788df2ca6e3fba668c4863d07e3d2ce7b32ff07eb7398e6b9d "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item/permissions"
    i5a39324ec2b645cb3cc090699e273fc292730c2bdf133601f5bb90f30ffb14ce "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item/subscriptions/item"
    i95f0d1a7b61d34fa7dc05db16c703139d0efe63fb3fa7cfe378630fe15aa5abb "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item/permissions/item"
    ia7f645faec80e7b6dfc84845aaa82c09e6e9ac1860011c4ae9ad93b8e59bf162 "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item/versions/item"
    ic0da1aab6993b0f8df8ad44bb5cfd16a3a2b19955e52a64023bc0455032638e7 "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item/children/item"
    ie52b86252bd16c7291b21a3b857f55a5f87fc9fb8cf5e76ef82bee358810df3b "github.com/microsoftgraph/msgraph-sdk-go/drive/bundles/item/thumbnails/item"
)

// DriveItemItemRequestBuilder provides operations to manage the bundles property of the microsoft.graph.drive entity.
type DriveItemItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DriveItemItemRequestBuilderDeleteOptions options for Delete
type DriveItemItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemItemRequestBuilderGetOptions options for Get
type DriveItemItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DriveItemItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemItemRequestBuilderGetQueryParameters collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
type DriveItemItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DriveItemItemRequestBuilderPatchOptions options for Patch
type DriveItemItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DriveItemItemRequestBuilder) Analytics()(*i61fb769ced9e59d2925fa83d5918db3045bd56b690fcfb92e92bd36e52a0453e.AnalyticsRequestBuilder) {
    return i61fb769ced9e59d2925fa83d5918db3045bd56b690fcfb92e92bd36e52a0453e.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i196a00f504cfcc6a9f69cef6ccb9b8a5412ee4bea114b8d6797615dc73aaa66c.ChildrenRequestBuilder) {
    return i196a00f504cfcc6a9f69cef6ccb9b8a5412ee4bea114b8d6797615dc73aaa66c.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.bundles.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*ic0da1aab6993b0f8df8ad44bb5cfd16a3a2b19955e52a64023bc0455032638e7.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return ic0da1aab6993b0f8df8ad44bb5cfd16a3a2b19955e52a64023bc0455032638e7.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/bundles/{driveItem_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemItemRequestBuilder instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemItemRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *DriveItemItemRequestBuilder) Content()(*id57c1f561fd7120e715af680ab9d67e3a72d92b54657bd3f3ac2432fcaec9d32.ContentRequestBuilder) {
    return id57c1f561fd7120e715af680ab9d67e3a72d92b54657bd3f3ac2432fcaec9d32.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property bundles for drive
func (m *DriveItemItemRequestBuilder) CreateDeleteRequestInformation(options *DriveItemItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *DriveItemItemRequestBuilder) CreateGetRequestInformation(options *DriveItemItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property bundles in drive
func (m *DriveItemItemRequestBuilder) CreatePatchRequestInformation(options *DriveItemItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property bundles for drive
func (m *DriveItemItemRequestBuilder) Delete(options *DriveItemItemRequestBuilderDeleteOptions)(error) {
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
// Get collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *DriveItemItemRequestBuilder) Get(options *DriveItemItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable), nil
}
func (m *DriveItemItemRequestBuilder) ListItem()(*ic5a3c66f7d6eff455463fd0119a5e0cc8849b531238ed938071be216b6a3cd45.ListItemRequestBuilder) {
    return ic5a3c66f7d6eff455463fd0119a5e0cc8849b531238ed938071be216b6a3cd45.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property bundles in drive
func (m *DriveItemItemRequestBuilder) Patch(options *DriveItemItemRequestBuilderPatchOptions)(error) {
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
func (m *DriveItemItemRequestBuilder) Permissions()(*ie600c6a3a6a376788df2ca6e3fba668c4863d07e3d2ce7b32ff07eb7398e6b9d.PermissionsRequestBuilder) {
    return ie600c6a3a6a376788df2ca6e3fba668c4863d07e3d2ce7b32ff07eb7398e6b9d.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.bundles.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i95f0d1a7b61d34fa7dc05db16c703139d0efe63fb3fa7cfe378630fe15aa5abb.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i95f0d1a7b61d34fa7dc05db16c703139d0efe63fb3fa7cfe378630fe15aa5abb.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*ibb083680698bfadb1b9ca3abdc030e7f9c822e8c1f6be855140f5f7242ade92e.SubscriptionsRequestBuilder) {
    return ibb083680698bfadb1b9ca3abdc030e7f9c822e8c1f6be855140f5f7242ade92e.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.bundles.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i5a39324ec2b645cb3cc090699e273fc292730c2bdf133601f5bb90f30ffb14ce.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i5a39324ec2b645cb3cc090699e273fc292730c2bdf133601f5bb90f30ffb14ce.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*id33754ddf1402bf93373f2d10f1c524fc6f4f4dc25c30c3b8e0adcbed34faff7.ThumbnailsRequestBuilder) {
    return id33754ddf1402bf93373f2d10f1c524fc6f4f4dc25c30c3b8e0adcbed34faff7.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.bundles.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*ie52b86252bd16c7291b21a3b857f55a5f87fc9fb8cf5e76ef82bee358810df3b.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return ie52b86252bd16c7291b21a3b857f55a5f87fc9fb8cf5e76ef82bee358810df3b.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*ia05955119e155abf2b26e50ebc55da0f7085d6c22feb85011b6ac8d28350c267.VersionsRequestBuilder) {
    return ia05955119e155abf2b26e50ebc55da0f7085d6c22feb85011b6ac8d28350c267.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.bundles.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*ia7f645faec80e7b6dfc84845aaa82c09e6e9ac1860011c4ae9ad93b8e59bf162.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return ia7f645faec80e7b6dfc84845aaa82c09e6e9ac1860011c4ae9ad93b8e59bf162.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
