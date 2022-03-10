package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/odataerrors"
    i40c8d5da7fc9a6217b6efac58c3740cd58572b63d41abf7ff231e644f35845bf "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/items"
    i6048965dbea1b9d4b091a4401571b9e4939fc3dcaa6066a36fadd4ad0d88ebe2 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/following"
    i792e0b31993e4efcc79acff25a9598f5972b2ddf9c703fe62d86508fdf135dfe "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/root"
    i84f47258d3d9d84f919cffd7beaf67d9220276155cd1e9600f024e1c38fc1ac1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/special"
    ief3ab11045cc2316a7986e50ca5252ef1bf5247bdd7b782ebc616fd9dedc8c6a "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/list"
    if274b7c54959658e48996ff8d954ee4de4e1313851545174efe44dd6d0573768 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/bundles"
    i1b0986cf2bb3f8cdd5c9626af60c220b8ed74834d8de43da091f289373bb0f7e "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/following/item"
    i7549c03bbbc8a5eeb871e3d1e7e255b35d87f970c4a55017c135ddac1bc09a5f "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/items/item"
    i7eaa82d7e156c1f8fe0101ef8557ffce59e5382dc33b38bf5c25f50084ccf3a1 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/special/item"
    i8b396b63c854a201e00ff83a78d93f8292a7b1a3895c9c05cf8b735f3efe3138 "github.com/microsoftgraph/msgraph-sdk-go/groups/item/drives/item/bundles/item"
)

// DriveItemRequestBuilder provides operations to manage the drives property of the microsoft.graph.group entity.
type DriveItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DriveItemRequestBuilderDeleteOptions options for Delete
type DriveItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemRequestBuilderGetOptions options for Get
type DriveItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DriveItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DriveItemRequestBuilderGetQueryParameters the group's drives. Read-only.
type DriveItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DriveItemRequestBuilderPatchOptions options for Patch
type DriveItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Driveable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DriveItemRequestBuilder) Bundles()(*if274b7c54959658e48996ff8d954ee4de4e1313851545174efe44dd6d0573768.BundlesRequestBuilder) {
    return if274b7c54959658e48996ff8d954ee4de4e1313851545174efe44dd6d0573768.NewBundlesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// BundlesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.drives.item.bundles.item collection
func (m *DriveItemRequestBuilder) BundlesById(id string)(*i8b396b63c854a201e00ff83a78d93f8292a7b1a3895c9c05cf8b735f3efe3138.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i8b396b63c854a201e00ff83a78d93f8292a7b1a3895c9c05cf8b735f3efe3138.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDriveItemRequestBuilderInternal instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemRequestBuilder) {
    m := &DriveItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/groups/{group_id}/drives/{drive_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDriveItemRequestBuilder instantiates a new DriveItemRequestBuilder and sets the default values.
func NewDriveItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DriveItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDriveItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete navigation property drives for groups
func (m *DriveItemRequestBuilder) CreateDeleteRequestInformation(options *DriveItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation the group's drives. Read-only.
func (m *DriveItemRequestBuilder) CreateGetRequestInformation(options *DriveItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update the navigation property drives in groups
func (m *DriveItemRequestBuilder) CreatePatchRequestInformation(options *DriveItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete navigation property drives for groups
func (m *DriveItemRequestBuilder) Delete(options *DriveItemRequestBuilderDeleteOptions)(error) {
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
func (m *DriveItemRequestBuilder) Following()(*i6048965dbea1b9d4b091a4401571b9e4939fc3dcaa6066a36fadd4ad0d88ebe2.FollowingRequestBuilder) {
    return i6048965dbea1b9d4b091a4401571b9e4939fc3dcaa6066a36fadd4ad0d88ebe2.NewFollowingRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// FollowingById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.drives.item.following.item collection
func (m *DriveItemRequestBuilder) FollowingById(id string)(*i1b0986cf2bb3f8cdd5c9626af60c220b8ed74834d8de43da091f289373bb0f7e.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i1b0986cf2bb3f8cdd5c9626af60c220b8ed74834d8de43da091f289373bb0f7e.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get the group's drives. Read-only.
func (m *DriveItemRequestBuilder) Get(options *DriveItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Driveable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
        "5XX": i7df4e557a1198b9abe14a17b40c7ac7db49b0d3050c749c3169541cb6f012b8b.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDriveFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Driveable), nil
}
func (m *DriveItemRequestBuilder) Items()(*i40c8d5da7fc9a6217b6efac58c3740cd58572b63d41abf7ff231e644f35845bf.ItemsRequestBuilder) {
    return i40c8d5da7fc9a6217b6efac58c3740cd58572b63d41abf7ff231e644f35845bf.NewItemsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ItemsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.drives.item.items.item collection
func (m *DriveItemRequestBuilder) ItemsById(id string)(*i7549c03bbbc8a5eeb871e3d1e7e255b35d87f970c4a55017c135ddac1bc09a5f.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i7549c03bbbc8a5eeb871e3d1e7e255b35d87f970c4a55017c135ddac1bc09a5f.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) List()(*ief3ab11045cc2316a7986e50ca5252ef1bf5247bdd7b782ebc616fd9dedc8c6a.ListRequestBuilder) {
    return ief3ab11045cc2316a7986e50ca5252ef1bf5247bdd7b782ebc616fd9dedc8c6a.NewListRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update the navigation property drives in groups
func (m *DriveItemRequestBuilder) Patch(options *DriveItemRequestBuilderPatchOptions)(error) {
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
func (m *DriveItemRequestBuilder) Root()(*i792e0b31993e4efcc79acff25a9598f5972b2ddf9c703fe62d86508fdf135dfe.RootRequestBuilder) {
    return i792e0b31993e4efcc79acff25a9598f5972b2ddf9c703fe62d86508fdf135dfe.NewRootRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DriveItemRequestBuilder) Special()(*i84f47258d3d9d84f919cffd7beaf67d9220276155cd1e9600f024e1c38fc1ac1.SpecialRequestBuilder) {
    return i84f47258d3d9d84f919cffd7beaf67d9220276155cd1e9600f024e1c38fc1ac1.NewSpecialRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// SpecialById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.groups.item.drives.item.special.item collection
func (m *DriveItemRequestBuilder) SpecialById(id string)(*i7eaa82d7e156c1f8fe0101ef8557ffce59e5382dc33b38bf5c25f50084ccf3a1.DriveItemItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["driveItem_id"] = id
    }
    return i7eaa82d7e156c1f8fe0101ef8557ffce59e5382dc33b38bf5c25f50084ccf3a1.NewDriveItemItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
