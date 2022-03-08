package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i0e5c9a32179abc7285a029a80ab04c202e99a224840c8b1b2f7ac93ada87e7dc "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item/subscriptions"
    i355ec1e9ffdfa89a7a08c95fd3e83cadb95d3e0161bb1397c823ac06065ba8e0 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item/listitem"
    i5cd0d87c9e838b53c655e5f09feb60cabb4900d2924df9ce7f7af11daa7ef13f "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item/content"
    i7210ca04cf9b0d7df1b87c5a63afb9be1bb2f83dc75fa92fb85ab9cb710c0d83 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item/versions"
    i8d17fa91c4dfafc98b0037f43b0de0d488df77442eb3abcb9b213081d28842fe "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item/thumbnails"
    i8fd6c4ab257bf3f12a2c26e02e5204ea22e6be7b3c19be1fa87e8ddd00050f5a "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item/analytics"
    icb954d58897fbd96dbba37db89d3650f67197108af6d14b8f217ca0ded9c3483 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item/children"
    ice018df2b9a164d207804d00e23efd81767d95caf272f4a659e15db8bd20ce43 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item/permissions"
    i077042217a9dfd17c953af4e2c54ce2d3f6261ced510662e1a2e811c783b3a58 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item/subscriptions/item"
    i2b92c89ee1f834318fb064ccdfcc7b95ec95c002a54e631bf6804155916cc337 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item/permissions/item"
    i2eaee21a95dd2fa09ee1114fa907fdaa5ad47f2e7f9b29143236b004a061370e "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item/thumbnails/item"
    ib0fe672c94508abacdbf07932d5760e35f212cc8ddc0478071eba293cec0c9d7 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item/versions/item"
    idcfeb2bf4ef4e715518bd648dc928e50a61d800a046bbfc89aaebf59648556c5 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/special/item/children/item"
)

// DriveItemItemRequestBuilder provides operations to manage the special property of the microsoft.graph.drive entity.
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
// DriveItemItemRequestBuilderGetQueryParameters collection of common folders available in OneDrive. Read-only. Nullable.
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
func (m *DriveItemItemRequestBuilder) Analytics()(*i8fd6c4ab257bf3f12a2c26e02e5204ea22e6be7b3c19be1fa87e8ddd00050f5a.AnalyticsRequestBuilder) {
    return i8fd6c4ab257bf3f12a2c26e02e5204ea22e6be7b3c19be1fa87e8ddd00050f5a.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*icb954d58897fbd96dbba37db89d3650f67197108af6d14b8f217ca0ded9c3483.ChildrenRequestBuilder) {
    return icb954d58897fbd96dbba37db89d3650f67197108af6d14b8f217ca0ded9c3483.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.special.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*idcfeb2bf4ef4e715518bd648dc928e50a61d800a046bbfc89aaebf59648556c5.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return idcfeb2bf4ef4e715518bd648dc928e50a61d800a046bbfc89aaebf59648556c5.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/special/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*i5cd0d87c9e838b53c655e5f09feb60cabb4900d2924df9ce7f7af11daa7ef13f.ContentRequestBuilder) {
    return i5cd0d87c9e838b53c655e5f09feb60cabb4900d2924df9ce7f7af11daa7ef13f.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property special for me
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
// CreateGetRequestInformation collection of common folders available in OneDrive. Read-only. Nullable.
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
// CreatePatchRequestInformation update the navigation property special in me
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
// Delete delete navigation property special for me
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
// Get collection of common folders available in OneDrive. Read-only. Nullable.
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
func (m *DriveItemItemRequestBuilder) ListItem()(*i355ec1e9ffdfa89a7a08c95fd3e83cadb95d3e0161bb1397c823ac06065ba8e0.ListItemRequestBuilder) {
    return i355ec1e9ffdfa89a7a08c95fd3e83cadb95d3e0161bb1397c823ac06065ba8e0.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property special in me
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
func (m *DriveItemItemRequestBuilder) Permissions()(*ice018df2b9a164d207804d00e23efd81767d95caf272f4a659e15db8bd20ce43.PermissionsRequestBuilder) {
    return ice018df2b9a164d207804d00e23efd81767d95caf272f4a659e15db8bd20ce43.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.special.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i2b92c89ee1f834318fb064ccdfcc7b95ec95c002a54e631bf6804155916cc337.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i2b92c89ee1f834318fb064ccdfcc7b95ec95c002a54e631bf6804155916cc337.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*i0e5c9a32179abc7285a029a80ab04c202e99a224840c8b1b2f7ac93ada87e7dc.SubscriptionsRequestBuilder) {
    return i0e5c9a32179abc7285a029a80ab04c202e99a224840c8b1b2f7ac93ada87e7dc.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.special.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i077042217a9dfd17c953af4e2c54ce2d3f6261ced510662e1a2e811c783b3a58.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i077042217a9dfd17c953af4e2c54ce2d3f6261ced510662e1a2e811c783b3a58.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i8d17fa91c4dfafc98b0037f43b0de0d488df77442eb3abcb9b213081d28842fe.ThumbnailsRequestBuilder) {
    return i8d17fa91c4dfafc98b0037f43b0de0d488df77442eb3abcb9b213081d28842fe.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.special.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i2eaee21a95dd2fa09ee1114fa907fdaa5ad47f2e7f9b29143236b004a061370e.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i2eaee21a95dd2fa09ee1114fa907fdaa5ad47f2e7f9b29143236b004a061370e.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i7210ca04cf9b0d7df1b87c5a63afb9be1bb2f83dc75fa92fb85ab9cb710c0d83.VersionsRequestBuilder) {
    return i7210ca04cf9b0d7df1b87c5a63afb9be1bb2f83dc75fa92fb85ab9cb710c0d83.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.special.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*ib0fe672c94508abacdbf07932d5760e35f212cc8ddc0478071eba293cec0c9d7.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return ib0fe672c94508abacdbf07932d5760e35f212cc8ddc0478071eba293cec0c9d7.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
