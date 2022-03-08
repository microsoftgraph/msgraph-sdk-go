package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i00bc066ede9dc6ca954a36ce6326ed9aa0ec8991ec3062f5ff8f5417f29deb8d "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item/children"
    i0291314f7d5303db7720c9d99394001ae2ec132fce0c31984e7d0bacd7474dc6 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item/versions"
    i7fa4b06eaab528b0b0cd904b722cb5484263db9cfad3c356df2e5a913cc528fd "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item/thumbnails"
    i8f977d2963cd53e113fb6b6fec8233d5fa51f87e71999c632c8d29012540d6b4 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item/listitem"
    ib3e26ae23f8c994f5814e9bf2c8ed62bf39e8d1b5c5110ce2f8a94d92d590c13 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item/analytics"
    iba5ca14914afc5b01afe019739b4ec73cc10e7e35893da784148471d6869ff71 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item/subscriptions"
    ic35bc26f986b92ba2b8123b04d70b7e11b7d7a1439e641e598c9881631c4bd05 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item/permissions"
    ic6a59f2d268c2e11cd934d013493304cb6ccccabb99ca15f63e8cb181abb042b "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item/content"
    i12c60312870dac84927d981a87bdbe85d38443f5abbfcd7aa5b6f129f67d94fc "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item/thumbnails/item"
    i42074b984122600fdbaf2472d779563fe4aefe6457d427e8f750a2f7165be212 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item/children/item"
    i43d5666ad62028b7b8f2f1c99542ebdb81664f15eb69aabab309ae57aa462d27 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item/subscriptions/item"
    ia3530d4912dee4b9f05ea051dd15c14672b24d4dd4e53e8d856fe027b97d7bd4 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item/versions/item"
    ic4e2091a1b4e1d9d0a34bafe7fa4f58db36ffb67ca1992819b05dccc584c6d75 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/bundles/item/permissions/item"
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
func (m *DriveItemItemRequestBuilder) Analytics()(*ib3e26ae23f8c994f5814e9bf2c8ed62bf39e8d1b5c5110ce2f8a94d92d590c13.AnalyticsRequestBuilder) {
    return ib3e26ae23f8c994f5814e9bf2c8ed62bf39e8d1b5c5110ce2f8a94d92d590c13.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*i00bc066ede9dc6ca954a36ce6326ed9aa0ec8991ec3062f5ff8f5417f29deb8d.ChildrenRequestBuilder) {
    return i00bc066ede9dc6ca954a36ce6326ed9aa0ec8991ec3062f5ff8f5417f29deb8d.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.bundles.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*i42074b984122600fdbaf2472d779563fe4aefe6457d427e8f750a2f7165be212.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return i42074b984122600fdbaf2472d779563fe4aefe6457d427e8f750a2f7165be212.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/bundles/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*ic6a59f2d268c2e11cd934d013493304cb6ccccabb99ca15f63e8cb181abb042b.ContentRequestBuilder) {
    return ic6a59f2d268c2e11cd934d013493304cb6ccccabb99ca15f63e8cb181abb042b.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property bundles for drives
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
// CreatePatchRequestInformation update the navigation property bundles in drives
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
// Delete delete navigation property bundles for drives
func (m *DriveItemItemRequestBuilder) Delete(options *DriveItemItemRequestBuilderDeleteOptions)(error) {
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
// Get collection of [bundles][bundle] (albums and multi-select-shared sets of items). Only in personal OneDrive.
func (m *DriveItemItemRequestBuilder) Get(options *DriveItemItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable), nil
}
func (m *DriveItemItemRequestBuilder) ListItem()(*i8f977d2963cd53e113fb6b6fec8233d5fa51f87e71999c632c8d29012540d6b4.ListItemRequestBuilder) {
    return i8f977d2963cd53e113fb6b6fec8233d5fa51f87e71999c632c8d29012540d6b4.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property bundles in drives
func (m *DriveItemItemRequestBuilder) Patch(options *DriveItemItemRequestBuilderPatchOptions)(error) {
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
func (m *DriveItemItemRequestBuilder) Permissions()(*ic35bc26f986b92ba2b8123b04d70b7e11b7d7a1439e641e598c9881631c4bd05.PermissionsRequestBuilder) {
    return ic35bc26f986b92ba2b8123b04d70b7e11b7d7a1439e641e598c9881631c4bd05.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.bundles.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*ic4e2091a1b4e1d9d0a34bafe7fa4f58db36ffb67ca1992819b05dccc584c6d75.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return ic4e2091a1b4e1d9d0a34bafe7fa4f58db36ffb67ca1992819b05dccc584c6d75.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*iba5ca14914afc5b01afe019739b4ec73cc10e7e35893da784148471d6869ff71.SubscriptionsRequestBuilder) {
    return iba5ca14914afc5b01afe019739b4ec73cc10e7e35893da784148471d6869ff71.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.bundles.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i43d5666ad62028b7b8f2f1c99542ebdb81664f15eb69aabab309ae57aa462d27.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i43d5666ad62028b7b8f2f1c99542ebdb81664f15eb69aabab309ae57aa462d27.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i7fa4b06eaab528b0b0cd904b722cb5484263db9cfad3c356df2e5a913cc528fd.ThumbnailsRequestBuilder) {
    return i7fa4b06eaab528b0b0cd904b722cb5484263db9cfad3c356df2e5a913cc528fd.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.bundles.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i12c60312870dac84927d981a87bdbe85d38443f5abbfcd7aa5b6f129f67d94fc.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i12c60312870dac84927d981a87bdbe85d38443f5abbfcd7aa5b6f129f67d94fc.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i0291314f7d5303db7720c9d99394001ae2ec132fce0c31984e7d0bacd7474dc6.VersionsRequestBuilder) {
    return i0291314f7d5303db7720c9d99394001ae2ec132fce0c31984e7d0bacd7474dc6.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.bundles.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*ia3530d4912dee4b9f05ea051dd15c14672b24d4dd4e53e8d856fe027b97d7bd4.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return ia3530d4912dee4b9f05ea051dd15c14672b24d4dd4e53e8d856fe027b97d7bd4.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
