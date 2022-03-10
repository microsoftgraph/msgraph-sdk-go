package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i3af63adf0d43d4fe53b7fe2f44a532d155b45cb427b6c05529081de55b84e0f7 "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item/permissions"
    i45c66de3d72bb8d0979aaa2542379091c3d4c3365f52d9867ba7a6dcc9e5df42 "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item/thumbnails"
    i5a016ea83f5468d3d384668d4e1b7f692e56c99df4537c4474da6357c1c96807 "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item/versions"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i7f342a44650eaf54d67337200a5d5d5a4bb27afa3086b4f49cae0ef5338c7770 "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item/listitem"
    ia6465f33760e8f7e56d0d23b7935140ceaca1da76050c18eca543431030efdbf "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item/subscriptions"
    iaac3a30c4b6c5fe7bf905435a8e107a25c1d0e91598227b2d7ccbee96f8c1a0c "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item/analytics"
    ibd5b8c197c6af68fece20ade2ff9b16b4366794693211d859e6567778bd151f5 "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item/content"
    icd9dbfb9ba416b7ebdcca39623f665287090d298a69c8ce4137a88826afc519a "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item/children"
    i4c8b4fdd4d449520ec88c6ee16d559f059a899da5b4405e59d57b14803a8bf70 "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item/thumbnails/item"
    i53f4ce758441460d7c2590ee04d34b039b97e2a7f5145228c8896c6c6e2b9d9c "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item/subscriptions/item"
    i6b9c730cb9b03b7166904de06757c7f45cec60fcdadc1b61ac7d35ab118d1705 "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item/permissions/item"
    i6e93dd5488b1f8ff104190af0e14ef7e0ca5e8c426e2f091ef84ee678588bb80 "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item/versions/item"
    iec7c656682c8d279d280cc673c823f79a674fc7e1d221f9f0a61f78e3c823367 "github.com/microsoftgraph/msgraph-sdk-go/drive/following/item/children/item"
)

// DriveItemItemRequestBuilder provides operations to manage the following property of the microsoft.graph.drive entity.
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
// DriveItemItemRequestBuilderGetQueryParameters the list of items the user is following. Only in OneDrive for Business.
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
func (m *DriveItemItemRequestBuilder) Analytics()(*iaac3a30c4b6c5fe7bf905435a8e107a25c1d0e91598227b2d7ccbee96f8c1a0c.AnalyticsRequestBuilder) {
    return iaac3a30c4b6c5fe7bf905435a8e107a25c1d0e91598227b2d7ccbee96f8c1a0c.NewAnalyticsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Children()(*icd9dbfb9ba416b7ebdcca39623f665287090d298a69c8ce4137a88826afc519a.ChildrenRequestBuilder) {
    return icd9dbfb9ba416b7ebdcca39623f665287090d298a69c8ce4137a88826afc519a.NewChildrenRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ChildrenById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.following.item.children.item collection
func (m *DriveItemItemRequestBuilder) ChildrenById(id string)(*iec7c656682c8d279d280cc673c823f79a674fc7e1d221f9f0a61f78e3c823367.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id1"] = id
    }
    return iec7c656682c8d279d280cc673c823f79a674fc7e1d221f9f0a61f78e3c823367.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemItemRequestBuilderInternal instantiates a new DriveItemItemRequestBuilder and sets the default values.
func NewDriveItemItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemItemRequestBuilder) {
    m := &DriveItemItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/drive/following/{driveItem_id}{?select,expand}";
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
func (m *DriveItemItemRequestBuilder) Content()(*ibd5b8c197c6af68fece20ade2ff9b16b4366794693211d859e6567778bd151f5.ContentRequestBuilder) {
    return ibd5b8c197c6af68fece20ade2ff9b16b4366794693211d859e6567778bd151f5.NewContentRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CreateDeleteRequestInformation delete navigation property following for drive
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
// CreateGetRequestInformation the list of items the user is following. Only in OneDrive for Business.
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
// CreatePatchRequestInformation update the navigation property following in drive
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
// Delete delete navigation property following for drive
func (m *DriveItemItemRequestBuilder) Delete(options *DriveItemItemRequestBuilderDeleteOptions)(error) {
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
// Get the list of items the user is following. Only in OneDrive for Business.
func (m *DriveItemItemRequestBuilder) Get(options *DriveItemItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DriveItemable, error) {
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
func (m *DriveItemItemRequestBuilder) ListItem()(*i7f342a44650eaf54d67337200a5d5d5a4bb27afa3086b4f49cae0ef5338c7770.ListItemRequestBuilder) {
    return i7f342a44650eaf54d67337200a5d5d5a4bb27afa3086b4f49cae0ef5338c7770.NewListItemRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property following in drive
func (m *DriveItemItemRequestBuilder) Patch(options *DriveItemItemRequestBuilderPatchOptions)(error) {
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
func (m *DriveItemItemRequestBuilder) Permissions()(*i3af63adf0d43d4fe53b7fe2f44a532d155b45cb427b6c05529081de55b84e0f7.PermissionsRequestBuilder) {
    return i3af63adf0d43d4fe53b7fe2f44a532d155b45cb427b6c05529081de55b84e0f7.NewPermissionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// PermissionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.following.item.permissions.item collection
func (m *DriveItemItemRequestBuilder) PermissionsById(id string)(*i6b9c730cb9b03b7166904de06757c7f45cec60fcdadc1b61ac7d35ab118d1705.PermissionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["permission_id"] = id
    }
    return i6b9c730cb9b03b7166904de06757c7f45cec60fcdadc1b61ac7d35ab118d1705.NewPermissionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Subscriptions()(*ia6465f33760e8f7e56d0d23b7935140ceaca1da76050c18eca543431030efdbf.SubscriptionsRequestBuilder) {
    return ia6465f33760e8f7e56d0d23b7935140ceaca1da76050c18eca543431030efdbf.NewSubscriptionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SubscriptionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.following.item.subscriptions.item collection
func (m *DriveItemItemRequestBuilder) SubscriptionsById(id string)(*i53f4ce758441460d7c2590ee04d34b039b97e2a7f5145228c8896c6c6e2b9d9c.SubscriptionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["subscription_id"] = id
    }
    return i53f4ce758441460d7c2590ee04d34b039b97e2a7f5145228c8896c6c6e2b9d9c.NewSubscriptionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Thumbnails()(*i45c66de3d72bb8d0979aaa2542379091c3d4c3365f52d9867ba7a6dcc9e5df42.ThumbnailsRequestBuilder) {
    return i45c66de3d72bb8d0979aaa2542379091c3d4c3365f52d9867ba7a6dcc9e5df42.NewThumbnailsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ThumbnailsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.following.item.thumbnails.item collection
func (m *DriveItemItemRequestBuilder) ThumbnailsById(id string)(*i4c8b4fdd4d449520ec88c6ee16d559f059a899da5b4405e59d57b14803a8bf70.ThumbnailSetItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["thumbnailSet_id"] = id
    }
    return i4c8b4fdd4d449520ec88c6ee16d559f059a899da5b4405e59d57b14803a8bf70.NewThumbnailSetItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemItemRequestBuilder) Versions()(*i5a016ea83f5468d3d384668d4e1b7f692e56c99df4537c4474da6357c1c96807.VersionsRequestBuilder) {
    return i5a016ea83f5468d3d384668d4e1b7f692e56c99df4537c4474da6357c1c96807.NewVersionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VersionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.drive.following.item.versions.item collection
func (m *DriveItemItemRequestBuilder) VersionsById(id string)(*i6e93dd5488b1f8ff104190af0e14ef7e0ca5e8c426e2f091ef84ee678588bb80.DriveItemVersionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItemVersion_id"] = id
    }
    return i6e93dd5488b1f8ff104190af0e14ef7e0ca5e8c426e2f091ef84ee678588bb80.NewDriveItemVersionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
