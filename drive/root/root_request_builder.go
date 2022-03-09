package root

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i089dbefd5043e5207e6e2000a1bf061c66b399725298944660bf5cd253b7264f "github.com/microsoftgraph/msgraph-sdk-go/drive/root/listitem"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7ab3c96c86b3b38921e0b8f81dcd5b872cc990c5cf8809f28bea1fb4dcf1558e "github.com/microsoftgraph/msgraph-sdk-go/drive/root/permissions"
    i87e0f9e6aaa1c9caa4c22be1bb55dea82446b84482de27f4512e49f0a4483fed "github.com/microsoftgraph/msgraph-sdk-go/drive/root/thumbnails"
    ia29571508baf63a873910b5dae4f5a73689ef617c6a0a050e827a3e260b72e75 "github.com/microsoftgraph/msgraph-sdk-go/drive/root/versions"
    ia929efb0838f1045c0b71fc7e927522ade38d9a38eaf5c4a09af5e7b841887c6 "github.com/microsoftgraph/msgraph-sdk-go/drive/root/children"
    idb5c2e560e3daac4a3ed8f616c4dd2ffd7defc33875f4bd870963f083b1c03a8 "github.com/microsoftgraph/msgraph-sdk-go/drive/root/content"
    idec39c75646fe0f514f66c621871a98ad8f286b849dea224ba584455e0d2e34e "github.com/microsoftgraph/msgraph-sdk-go/drive/root/analytics"
    if9e05debf18ae7a89d29e3f92e3d20b4bd85e1de8b0abc6c5d1464f8a2df79ae "github.com/microsoftgraph/msgraph-sdk-go/drive/root/subscriptions"
    i451951d9cef8df3c8b2a8be8726dc1285485859b8f467b1880fe9484d1f86b30 "github.com/microsoftgraph/msgraph-sdk-go/drive/root/thumbnails/item"
    i56826de7fe9540a3d79c02bb8e18e83cbe6733df09aac1d6cf30342c155743a3 "github.com/microsoftgraph/msgraph-sdk-go/drive/root/children/item"
    i5a7c91cef86c9d63b5934d1ca4411013dd41d7a5772f7cc81ff281eff124ef60 "github.com/microsoftgraph/msgraph-sdk-go/drive/root/versions/item"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i83a35aa1860d0b10df007e13131be0541e05e6b90b022f9bdfbeb8e94bcc3562 "github.com/microsoftgraph/msgraph-sdk-go/drive/root/permissions/item"
    iabb32685aec6f73966e56b90a18908830f75cc93fae47ddcb640f1ffc9a3be02 "github.com/microsoftgraph/msgraph-sdk-go/drive/root/subscriptions/item"
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
func (m *RootRequestBuilder) Analytics()(*idec39c75646fe0f514f66c621871a98ad8f286b849dea224ba584455e0d2e34e.AnalyticsRequestBuilder) {
    return idec39c75646fe0f514f66c621871a98ad8f286b849dea224ba584455e0d2e34e.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RootRequestBuilder) Children()(*ia929efb0838f1045c0b71fc7e927522ade38d9a38eaf5c4a09af5e7b841887c6.ChildrenRequestBuilder) {
    return ia929efb0838f1045c0b71fc7e927522ade38d9a38eaf5c4a09af5e7b841887c6.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*i56826de7fe9540a3d79c02bb8e18e83cbe6733df09aac1d6cf30342c155743a3.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i56826de7fe9540a3d79c02bb8e18e83cbe6733df09aac1d6cf30342c155743a3.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/root{?select,expand}";
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
func (m *RootRequestBuilder) Content()(*idb5c2e560e3daac4a3ed8f616c4dd2ffd7defc33875f4bd870963f083b1c03a8.ContentRequestBuilder) {
    return idb5c2e560e3daac4a3ed8f616c4dd2ffd7defc33875f4bd870963f083b1c03a8.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for drive
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
// CreatePatchRequestInformation update the navigation property root in drive
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
// Delete delete navigation property root for drive
func (m *RootRequestBuilder) Delete(options *RootRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
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
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDriveItemFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable), nil
}
func (m *RootRequestBuilder) ListItem()(*i089dbefd5043e5207e6e2000a1bf061c66b399725298944660bf5cd253b7264f.ListItemRequestBuilder) {
    return i089dbefd5043e5207e6e2000a1bf061c66b399725298944660bf5cd253b7264f.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in drive
func (m *RootRequestBuilder) Patch(options *RootRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
func (m *RootRequestBuilder) Permissions()(*i7ab3c96c86b3b38921e0b8f81dcd5b872cc990c5cf8809f28bea1fb4dcf1558e.PermissionsRequestBuilder) {
    return i7ab3c96c86b3b38921e0b8f81dcd5b872cc990c5cf8809f28bea1fb4dcf1558e.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*i83a35aa1860d0b10df007e13131be0541e05e6b90b022f9bdfbeb8e94bcc3562.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i83a35aa1860d0b10df007e13131be0541e05e6b90b022f9bdfbeb8e94bcc3562.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Subscriptions()(*if9e05debf18ae7a89d29e3f92e3d20b4bd85e1de8b0abc6c5d1464f8a2df79ae.SubscriptionsRequestBuilder) {
    return if9e05debf18ae7a89d29e3f92e3d20b4bd85e1de8b0abc6c5d1464f8a2df79ae.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*iabb32685aec6f73966e56b90a18908830f75cc93fae47ddcb640f1ffc9a3be02.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return iabb32685aec6f73966e56b90a18908830f75cc93fae47ddcb640f1ffc9a3be02.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Thumbnails()(*i87e0f9e6aaa1c9caa4c22be1bb55dea82446b84482de27f4512e49f0a4483fed.ThumbnailsRequestBuilder) {
    return i87e0f9e6aaa1c9caa4c22be1bb55dea82446b84482de27f4512e49f0a4483fed.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*i451951d9cef8df3c8b2a8be8726dc1285485859b8f467b1880fe9484d1f86b30.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i451951d9cef8df3c8b2a8be8726dc1285485859b8f467b1880fe9484d1f86b30.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Versions()(*ia29571508baf63a873910b5dae4f5a73689ef617c6a0a050e827a3e260b72e75.VersionsRequestBuilder) {
    return ia29571508baf63a873910b5dae4f5a73689ef617c6a0a050e827a3e260b72e75.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*i5a7c91cef86c9d63b5934d1ca4411013dd41d7a5772f7cc81ff281eff124ef60.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i5a7c91cef86c9d63b5934d1ca4411013dd41d7a5772f7cc81ff281eff124ef60.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
