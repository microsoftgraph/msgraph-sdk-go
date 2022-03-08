package root

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i1c05393909ced8c970ac8182fc4a6c5cc419883d0d1475ce0773c5cf783ba989 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/listitem"
    i353c71cf4e7f00f79bca7c43c481c2b7fccfaf2bc763be7f85e457bcdb2a832b "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/content"
    i35ec8d39bca25c3225ac281fac7242810cb642a0a53e14a875e33c70b99c171e "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/analytics"
    i6332f0883e6ab0e66763410aa3cab286bc78b7b61f88fefedb1be3fb930832df "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/thumbnails"
    i6a155f272c5391daf639268076f7e78fb66358ef654a4f57db99ce00a32f3c11 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/children"
    i8586e46c09b38f90b06f86eb42fd10319ce6acb9451aefe28f39853cfb5719b9 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/versions"
    ia05f97a343c7dbf32d5d42d3b714d8d4e180ebdebfee2e83c7a95ec1ed6e9e3e "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/permissions"
    ic72fcde28010dbd679eab351c33b57107adc23c1cc48e905dd109128ec3bc4db "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/subscriptions"
    i162b49d374c91a27c7d63a8c7103a36e56411b6b3e84fc64ecbf5d828ae8a8af "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/versions/item"
    i3d94da0dddfc16aca3f8f4350e9b0d9d0a38bc1ffb67f012ae9cb026450ac2a1 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/subscriptions/item"
    i896e3d0df31741011967f273d16af8fb1abe2a50bd0bb2f742e0aead26932993 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/thumbnails/item"
    i8ca583e1745c18b27ce93c86c752cf4c03f78604e1970006d1886875c01639c7 "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/permissions/item"
    ic7d208ecab0cc1c68790901ba2ba97554ff0a0a86b46da4f6d1dba8ab740987f "github.com/microsoftgraph/msgraph-sdk-go/drives/item/root/children/item"
)

// RootRequestBuilder provides operations to manage the root property of the microsoft.graph.drive entity.
type RootRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// RootRequestBuilderDeleteOptions options for Delete
type RootRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RootRequestBuilderGetOptions options for Get
type RootRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *RootRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// RootRequestBuilderGetQueryParameters the root folder of the drive. Read-only.
type RootRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// RootRequestBuilderPatchOptions options for Patch
type RootRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *RootRequestBuilder) Analytics()(*i35ec8d39bca25c3225ac281fac7242810cb642a0a53e14a875e33c70b99c171e.AnalyticsRequestBuilder) {
    return i35ec8d39bca25c3225ac281fac7242810cb642a0a53e14a875e33c70b99c171e.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RootRequestBuilder) Children()(*i6a155f272c5391daf639268076f7e78fb66358ef654a4f57db99ce00a32f3c11.ChildrenRequestBuilder) {
    return i6a155f272c5391daf639268076f7e78fb66358ef654a4f57db99ce00a32f3c11.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*ic7d208ecab0cc1c68790901ba2ba97554ff0a0a86b46da4f6d1dba8ab740987f.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return ic7d208ecab0cc1c68790901ba2ba97554ff0a0a86b46da4f6d1dba8ab740987f.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drives/{drive_id}/root{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewRootRequestBuilder instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RootRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewRootRequestBuilderInternal(urlParams, requestAdapter)
}
func (m *RootRequestBuilder) Content()(*i353c71cf4e7f00f79bca7c43c481c2b7fccfaf2bc763be7f85e457bcdb2a832b.ContentRequestBuilder) {
    return i353c71cf4e7f00f79bca7c43c481c2b7fccfaf2bc763be7f85e457bcdb2a832b.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for drives
func (m *RootRequestBuilder) CreateDeleteRequestInformation(options *RootRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the root folder of the drive. Read-only.
func (m *RootRequestBuilder) CreateGetRequestInformation(options *RootRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property root in drives
func (m *RootRequestBuilder) CreatePatchRequestInformation(options *RootRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property root for drives
func (m *RootRequestBuilder) Delete(options *RootRequestBuilderDeleteOptions)(error) {
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
// Get the root folder of the drive. Read-only.
func (m *RootRequestBuilder) Get(options *RootRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable, error) {
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
func (m *RootRequestBuilder) ListItem()(*i1c05393909ced8c970ac8182fc4a6c5cc419883d0d1475ce0773c5cf783ba989.ListItemRequestBuilder) {
    return i1c05393909ced8c970ac8182fc4a6c5cc419883d0d1475ce0773c5cf783ba989.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in drives
func (m *RootRequestBuilder) Patch(options *RootRequestBuilderPatchOptions)(error) {
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
func (m *RootRequestBuilder) Permissions()(*ia05f97a343c7dbf32d5d42d3b714d8d4e180ebdebfee2e83c7a95ec1ed6e9e3e.PermissionsRequestBuilder) {
    return ia05f97a343c7dbf32d5d42d3b714d8d4e180ebdebfee2e83c7a95ec1ed6e9e3e.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*i8ca583e1745c18b27ce93c86c752cf4c03f78604e1970006d1886875c01639c7.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i8ca583e1745c18b27ce93c86c752cf4c03f78604e1970006d1886875c01639c7.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Subscriptions()(*ic72fcde28010dbd679eab351c33b57107adc23c1cc48e905dd109128ec3bc4db.SubscriptionsRequestBuilder) {
    return ic72fcde28010dbd679eab351c33b57107adc23c1cc48e905dd109128ec3bc4db.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i3d94da0dddfc16aca3f8f4350e9b0d9d0a38bc1ffb67f012ae9cb026450ac2a1.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i3d94da0dddfc16aca3f8f4350e9b0d9d0a38bc1ffb67f012ae9cb026450ac2a1.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Thumbnails()(*i6332f0883e6ab0e66763410aa3cab286bc78b7b61f88fefedb1be3fb930832df.ThumbnailsRequestBuilder) {
    return i6332f0883e6ab0e66763410aa3cab286bc78b7b61f88fefedb1be3fb930832df.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i896e3d0df31741011967f273d16af8fb1abe2a50bd0bb2f742e0aead26932993.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i896e3d0df31741011967f273d16af8fb1abe2a50bd0bb2f742e0aead26932993.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Versions()(*i8586e46c09b38f90b06f86eb42fd10319ce6acb9451aefe28f39853cfb5719b9.VersionsRequestBuilder) {
    return i8586e46c09b38f90b06f86eb42fd10319ce6acb9451aefe28f39853cfb5719b9.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i162b49d374c91a27c7d63a8c7103a36e56411b6b3e84fc64ecbf5d828ae8a8af.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i162b49d374c91a27c7d63a8c7103a36e56411b6b3e84fc64ecbf5d828ae8a8af.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
