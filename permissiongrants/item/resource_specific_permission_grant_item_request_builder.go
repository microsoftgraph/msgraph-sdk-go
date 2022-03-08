package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i4a7063fe0d09620e5d6522f433c1e9cf3e5d296667e87f9dfecfc4af1b3aae9d "github.com/microsoftgraph/msgraph-sdk-go/permissiongrants/item/getmemberobjects"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i8eeb73da3d89f8a760cd4adc6d53215b2c688043311f9d987c3f034cc4014646 "github.com/microsoftgraph/msgraph-sdk-go/permissiongrants/item/checkmembergroups"
    ia2efcc0d1f3eacb9f6edf91b52e6f227f398341a4af363e1ddc708ad6cade478 "github.com/microsoftgraph/msgraph-sdk-go/permissiongrants/item/restore"
    ia62a23d20d63f22d8eaaf57757dfb7d2ffbe1f8f51f38621c112384a083cf3a8 "github.com/microsoftgraph/msgraph-sdk-go/permissiongrants/item/getmembergroups"
    if0a1607357c7ac08b0f7369dd31ca3dae44d9d88c78c88afa398d08a8a0fe97e "github.com/microsoftgraph/msgraph-sdk-go/permissiongrants/item/checkmemberobjects"
)

// ResourceSpecificPermissionGrantItemRequestBuilder provides operations to manage the collection of resourceSpecificPermissionGrant entities.
type ResourceSpecificPermissionGrantItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// ResourceSpecificPermissionGrantItemRequestBuilderDeleteOptions options for Delete
type ResourceSpecificPermissionGrantItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ResourceSpecificPermissionGrantItemRequestBuilderGetOptions options for Get
type ResourceSpecificPermissionGrantItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *ResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// ResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters get entity from permissionGrants by key
type ResourceSpecificPermissionGrantItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// ResourceSpecificPermissionGrantItemRequestBuilderPatchOptions options for Patch
type ResourceSpecificPermissionGrantItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResourceSpecificPermissionGrantable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CheckMemberGroups()(*i8eeb73da3d89f8a760cd4adc6d53215b2c688043311f9d987c3f034cc4014646.CheckMemberGroupsRequestBuilder) {
    return i8eeb73da3d89f8a760cd4adc6d53215b2c688043311f9d987c3f034cc4014646.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CheckMemberObjects()(*if0a1607357c7ac08b0f7369dd31ca3dae44d9d88c78c88afa398d08a8a0fe97e.CheckMemberObjectsRequestBuilder) {
    return if0a1607357c7ac08b0f7369dd31ca3dae44d9d88c78c88afa398d08a8a0fe97e.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewResourceSpecificPermissionGrantItemRequestBuilderInternal instantiates a new ResourceSpecificPermissionGrantItemRequestBuilder and sets the default values.
func NewResourceSpecificPermissionGrantItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceSpecificPermissionGrantItemRequestBuilder) {
    m := &ResourceSpecificPermissionGrantItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/permissionGrants/{resourceSpecificPermissionGrant_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewResourceSpecificPermissionGrantItemRequestBuilder instantiates a new ResourceSpecificPermissionGrantItemRequestBuilder and sets the default values.
func NewResourceSpecificPermissionGrantItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*ResourceSpecificPermissionGrantItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewResourceSpecificPermissionGrantItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from permissionGrants
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CreateDeleteRequestInformation(options *ResourceSpecificPermissionGrantItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from permissionGrants by key
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CreateGetRequestInformation(options *ResourceSpecificPermissionGrantItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in permissionGrants
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) CreatePatchRequestInformation(options *ResourceSpecificPermissionGrantItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from permissionGrants
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) Delete(options *ResourceSpecificPermissionGrantItemRequestBuilderDeleteOptions)(error) {
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
// Get get entity from permissionGrants by key
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) Get(options *ResourceSpecificPermissionGrantItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResourceSpecificPermissionGrantable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateResourceSpecificPermissionGrantFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.ResourceSpecificPermissionGrantable), nil
}
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) GetMemberGroups()(*ia62a23d20d63f22d8eaaf57757dfb7d2ffbe1f8f51f38621c112384a083cf3a8.GetMemberGroupsRequestBuilder) {
    return ia62a23d20d63f22d8eaaf57757dfb7d2ffbe1f8f51f38621c112384a083cf3a8.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) GetMemberObjects()(*i4a7063fe0d09620e5d6522f433c1e9cf3e5d296667e87f9dfecfc4af1b3aae9d.GetMemberObjectsRequestBuilder) {
    return i4a7063fe0d09620e5d6522f433c1e9cf3e5d296667e87f9dfecfc4af1b3aae9d.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in permissionGrants
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) Patch(options *ResourceSpecificPermissionGrantItemRequestBuilderPatchOptions)(error) {
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
func (m *ResourceSpecificPermissionGrantItemRequestBuilder) Restore()(*ia2efcc0d1f3eacb9f6edf91b52e6f227f398341a4af363e1ddc708ad6cade478.RestoreRequestBuilder) {
    return ia2efcc0d1f3eacb9f6edf91b52e6f227f398341a4af363e1ddc708ad6cade478.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
