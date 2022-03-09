package root

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i032fb091fa50ded67bd8413cc227e15a89212f22579d11ba5fcda786303b3d88 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root/versions"
    i03ddf9ed2e8d64e93fc44a7258e3589c38d6f455c1358c4d75d1614935a98e46 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root/children"
    i1009a2ff258bf0b86c5aa9b40c49762ddd855f338f1947b3a8f825c3d8524721 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root/content"
    i404bc6886d1c9123a28011fced42055517ed97d7f967f514fb167d219b2c0db0 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root/listitem"
    i7669f738788c84c7ecf0251b8089bc20ff3a6e5b410b8045267fa73a34277705 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root/analytics"
    i80b1e91bf7b013cdf8dbdf9b8d8771f6a1b6e697578eefaecea07952add6388c "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root/permissions"
    ief5ea07106b57464e0d9313bab8da87f6427ea4ce6317fef25f88f0475243573 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root/thumbnails"
    ifce4a3fdda3e56bdf741ded90a576a9c0e9f6480576f685c4c8e5eceb911a516 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root/subscriptions"
    i689f37510f9fad8e0be6242df80f8b79fff25b25c310cfaf55e7b93c51c780af "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root/subscriptions/item"
    ia16000547e3bbe576a5a946472bea8c98307843b62bbdb98e519f2ae8166f2c2 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root/versions/item"
    ie2d9fc455cd20ad77bcd5ffeb7d799dfb360cc67d9482d6709d02097def6a274 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root/children/item"
    ieeb69da49c1200fd91a53dded9fdde76e190b3cb09517e3aab7c813ac910de6d "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root/permissions/item"
    if641b83445d13c546e420bdc7e9b23cb86e18512438d6cde4308543c0d7c8cb4 "github.com/microsoftgraph/msgraph-sdk-go/me/drives/item/root/thumbnails/item"
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
func (m *RootRequestBuilder) Analytics()(*i7669f738788c84c7ecf0251b8089bc20ff3a6e5b410b8045267fa73a34277705.AnalyticsRequestBuilder) {
    return i7669f738788c84c7ecf0251b8089bc20ff3a6e5b410b8045267fa73a34277705.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *RootRequestBuilder) Children()(*i03ddf9ed2e8d64e93fc44a7258e3589c38d6f455c1358c4d75d1614935a98e46.ChildrenRequestBuilder) {
    return i03ddf9ed2e8d64e93fc44a7258e3589c38d6f455c1358c4d75d1614935a98e46.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.root.children.item collection
func (m *RootRequestBuilder) ChildrenById(id string)(*ie2d9fc455cd20ad77bcd5ffeb7d799dfb360cc67d9482d6709d02097def6a274.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return ie2d9fc455cd20ad77bcd5ffeb7d799dfb360cc67d9482d6709d02097def6a274.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewRootRequestBuilderInternal instantiates a new RootRequestBuilder and sets the default values.
func NewRootRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*RootRequestBuilder) {
    m := &RootRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/me/drives/{drive_id}/root{?select,expand}";
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
func (m *RootRequestBuilder) Content()(*i1009a2ff258bf0b86c5aa9b40c49762ddd855f338f1947b3a8f825c3d8524721.ContentRequestBuilder) {
    return i1009a2ff258bf0b86c5aa9b40c49762ddd855f338f1947b3a8f825c3d8524721.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property root for me
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
// CreatePatchRequestInformation update the navigation property root in me
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
// Delete delete navigation property root for me
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
func (m *RootRequestBuilder) ListItem()(*i404bc6886d1c9123a28011fced42055517ed97d7f967f514fb167d219b2c0db0.ListItemRequestBuilder) {
    return i404bc6886d1c9123a28011fced42055517ed97d7f967f514fb167d219b2c0db0.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property root in me
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
func (m *RootRequestBuilder) Permissions()(*i80b1e91bf7b013cdf8dbdf9b8d8771f6a1b6e697578eefaecea07952add6388c.PermissionsRequestBuilder) {
    return i80b1e91bf7b013cdf8dbdf9b8d8771f6a1b6e697578eefaecea07952add6388c.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.root.permissions.item collection
func (m *RootRequestBuilder) PermissionsById(id string)(*ieeb69da49c1200fd91a53dded9fdde76e190b3cb09517e3aab7c813ac910de6d.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return ieeb69da49c1200fd91a53dded9fdde76e190b3cb09517e3aab7c813ac910de6d.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Subscriptions()(*ifce4a3fdda3e56bdf741ded90a576a9c0e9f6480576f685c4c8e5eceb911a516.SubscriptionsRequestBuilder) {
    return ifce4a3fdda3e56bdf741ded90a576a9c0e9f6480576f685c4c8e5eceb911a516.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.root.subscriptions.item collection
func (m *RootRequestBuilder) SubscriptionsById(id string)(*i689f37510f9fad8e0be6242df80f8b79fff25b25c310cfaf55e7b93c51c780af.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i689f37510f9fad8e0be6242df80f8b79fff25b25c310cfaf55e7b93c51c780af.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Thumbnails()(*ief5ea07106b57464e0d9313bab8da87f6427ea4ce6317fef25f88f0475243573.ThumbnailsRequestBuilder) {
    return ief5ea07106b57464e0d9313bab8da87f6427ea4ce6317fef25f88f0475243573.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.root.thumbnails.item collection
func (m *RootRequestBuilder) ThumbnailsById(id string)(*if641b83445d13c546e420bdc7e9b23cb86e18512438d6cde4308543c0d7c8cb4.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return if641b83445d13c546e420bdc7e9b23cb86e18512438d6cde4308543c0d7c8cb4.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *RootRequestBuilder) Versions()(*i032fb091fa50ded67bd8413cc227e15a89212f22579d11ba5fcda786303b3d88.VersionsRequestBuilder) {
    return i032fb091fa50ded67bd8413cc227e15a89212f22579d11ba5fcda786303b3d88.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.me.drives.item.root.versions.item collection
func (m *RootRequestBuilder) VersionsById(id string)(*ia16000547e3bbe576a5a946472bea8c98307843b62bbdb98e519f2ae8166f2c2.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return ia16000547e3bbe576a5a946472bea8c98307843b62bbdb98e519f2ae8166f2c2.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
