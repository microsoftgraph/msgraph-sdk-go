package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i1da16a4b620766997fcb2773cf780ba737b8dd64438f6bd1b46feb6cf67dfaf2 "github.com/microsoftgraph/msgraph-sdk-go/devices/item/memberof"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i6d9ea1deb1d3cfaac0245615a52427d153dc3b9897295d48d6cb94b90f2b2bb6 "github.com/microsoftgraph/msgraph-sdk-go/devices/item/extensions"
    i7f0d40d46e36fb24c086809fc583560c84c2d578585740e724f14d6bab94b6ac "github.com/microsoftgraph/msgraph-sdk-go/devices/item/restore"
    i802cbfeae38ef2d92747a6d0a312423fd41be33f3e3fb965305e643ae538adfa "github.com/microsoftgraph/msgraph-sdk-go/devices/item/getmemberobjects"
    iad59335f11b6e4f287fa8e2535c4c39e6d1968fb30e6d9729d87370eaf1a741c "github.com/microsoftgraph/msgraph-sdk-go/devices/item/registeredowners"
    id382f452535c2e36f3f851dd3aeecca79c528c74c33433e51efa465034698ec2 "github.com/microsoftgraph/msgraph-sdk-go/devices/item/transitivememberof"
    id4c71a00690f15e834a677b658e9a8fc4ab7659038027b563d609b48b71d76bc "github.com/microsoftgraph/msgraph-sdk-go/devices/item/registeredusers"
    ie01af685f90d1d6390f1cca50e6d6a4a0c0d1400e061efa09108ed2cc4f121ce "github.com/microsoftgraph/msgraph-sdk-go/devices/item/checkmemberobjects"
    ie61880210a289f885eb584ae4ffb8775884acae6cbcc7f3ef5c31226565fbb49 "github.com/microsoftgraph/msgraph-sdk-go/devices/item/checkmembergroups"
    ifad36aa7c4e70f6ef3f9d729bcb5622ac992a86d8192af5b3f76265c38bc1854 "github.com/microsoftgraph/msgraph-sdk-go/devices/item/getmembergroups"
    i3510737f0d62ae1028132f9eb8460789b2745e3f0425ffac6d00ba945e3ad462 "github.com/microsoftgraph/msgraph-sdk-go/devices/item/memberof/item"
    i81b3cccc98b5736dc9541adae22c6bfe0b07307fa1ee207c5666690a14f4d776 "github.com/microsoftgraph/msgraph-sdk-go/devices/item/transitivememberof/item"
    ib133a87f1f6aa82508a13fd26e99937ae943d6e84c9447f4c022ffc647194620 "github.com/microsoftgraph/msgraph-sdk-go/devices/item/extensions/item"
    ib4628a5555deb2b60e49a55e75b33d25c1aa9f9e4a77c4c8fd1022fda5ae1027 "github.com/microsoftgraph/msgraph-sdk-go/devices/item/registeredowners/item"
    ifc731e509397c6e1502bcc9cb553695d7c0475072d83d49e638a8991044f9333 "github.com/microsoftgraph/msgraph-sdk-go/devices/item/registeredusers/item"
)

// DeviceItemRequestBuilder provides operations to manage the collection of device entities.
type DeviceItemRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceItemRequestBuilderDeleteOptions options for Delete
type DeviceItemRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceItemRequestBuilderGetOptions options for Get
type DeviceItemRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceItemRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceItemRequestBuilderGetQueryParameters get entity from devices by key
type DeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceItemRequestBuilderPatchOptions options for Patch
type DeviceItemRequestBuilderPatchOptions struct {
    // 
    Body i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Deviceable;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceItemRequestBuilder) CheckMemberGroups()(*ie61880210a289f885eb584ae4ffb8775884acae6cbcc7f3ef5c31226565fbb49.CheckMemberGroupsRequestBuilder) {
    return ie61880210a289f885eb584ae4ffb8775884acae6cbcc7f3ef5c31226565fbb49.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) CheckMemberObjects()(*ie01af685f90d1d6390f1cca50e6d6a4a0c0d1400e061efa09108ed2cc4f121ce.CheckMemberObjectsRequestBuilder) {
    return ie01af685f90d1d6390f1cca50e6d6a4a0c0d1400e061efa09108ed2cc4f121ce.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceItemRequestBuilderInternal instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceItemRequestBuilder) {
    m := &DeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/devices/{device_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceItemRequestBuilder instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from devices
func (m *DeviceItemRequestBuilder) CreateDeleteRequestInformation(options *DeviceItemRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreateGetRequestInformation get entity from devices by key
func (m *DeviceItemRequestBuilder) CreateGetRequestInformation(options *DeviceItemRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update entity in devices
func (m *DeviceItemRequestBuilder) CreatePatchRequestInformation(options *DeviceItemRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// Delete delete entity from devices
func (m *DeviceItemRequestBuilder) Delete(options *DeviceItemRequestBuilderDeleteOptions)(error) {
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
func (m *DeviceItemRequestBuilder) Extensions()(*i6d9ea1deb1d3cfaac0245615a52427d153dc3b9897295d48d6cb94b90f2b2bb6.ExtensionsRequestBuilder) {
    return i6d9ea1deb1d3cfaac0245615a52427d153dc3b9897295d48d6cb94b90f2b2bb6.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.devices.item.extensions.item collection
func (m *DeviceItemRequestBuilder) ExtensionsById(id string)(*ib133a87f1f6aa82508a13fd26e99937ae943d6e84c9447f4c022ffc647194620.ExtensionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ib133a87f1f6aa82508a13fd26e99937ae943d6e84c9447f4c022ffc647194620.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from devices by key
func (m *DeviceItemRequestBuilder) Get(options *DeviceItemRequestBuilderGetOptions)(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Deviceable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ErrorMappings {
        "4XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
        "5XX": i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.CreateDeviceFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Deviceable), nil
}
func (m *DeviceItemRequestBuilder) GetMemberGroups()(*ifad36aa7c4e70f6ef3f9d729bcb5622ac992a86d8192af5b3f76265c38bc1854.GetMemberGroupsRequestBuilder) {
    return ifad36aa7c4e70f6ef3f9d729bcb5622ac992a86d8192af5b3f76265c38bc1854.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) GetMemberObjects()(*i802cbfeae38ef2d92747a6d0a312423fd41be33f3e3fb965305e643ae538adfa.GetMemberObjectsRequestBuilder) {
    return i802cbfeae38ef2d92747a6d0a312423fd41be33f3e3fb965305e643ae538adfa.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) MemberOf()(*i1da16a4b620766997fcb2773cf780ba737b8dd64438f6bd1b46feb6cf67dfaf2.MemberOfRequestBuilder) {
    return i1da16a4b620766997fcb2773cf780ba737b8dd64438f6bd1b46feb6cf67dfaf2.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MemberOfById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.devices.item.memberOf.item collection
func (m *DeviceItemRequestBuilder) MemberOfById(id string)(*i3510737f0d62ae1028132f9eb8460789b2745e3f0425ffac6d00ba945e3ad462.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i3510737f0d62ae1028132f9eb8460789b2745e3f0425ffac6d00ba945e3ad462.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update entity in devices
func (m *DeviceItemRequestBuilder) Patch(options *DeviceItemRequestBuilderPatchOptions)(error) {
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
func (m *DeviceItemRequestBuilder) RegisteredOwners()(*iad59335f11b6e4f287fa8e2535c4c39e6d1968fb30e6d9729d87370eaf1a741c.RegisteredOwnersRequestBuilder) {
    return iad59335f11b6e4f287fa8e2535c4c39e6d1968fb30e6d9729d87370eaf1a741c.NewRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredOwnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.devices.item.registeredOwners.item collection
func (m *DeviceItemRequestBuilder) RegisteredOwnersById(id string)(*ib4628a5555deb2b60e49a55e75b33d25c1aa9f9e4a77c4c8fd1022fda5ae1027.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return ib4628a5555deb2b60e49a55e75b33d25c1aa9f9e4a77c4c8fd1022fda5ae1027.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) RegisteredUsers()(*id4c71a00690f15e834a677b658e9a8fc4ab7659038027b563d609b48b71d76bc.RegisteredUsersRequestBuilder) {
    return id4c71a00690f15e834a677b658e9a8fc4ab7659038027b563d609b48b71d76bc.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RegisteredUsersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.devices.item.registeredUsers.item collection
func (m *DeviceItemRequestBuilder) RegisteredUsersById(id string)(*ifc731e509397c6e1502bcc9cb553695d7c0475072d83d49e638a8991044f9333.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return ifc731e509397c6e1502bcc9cb553695d7c0475072d83d49e638a8991044f9333.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) Restore()(*i7f0d40d46e36fb24c086809fc583560c84c2d578585740e724f14d6bab94b6ac.RestoreRequestBuilder) {
    return i7f0d40d46e36fb24c086809fc583560c84c2d578585740e724f14d6bab94b6ac.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceItemRequestBuilder) TransitiveMemberOf()(*id382f452535c2e36f3f851dd3aeecca79c528c74c33433e51efa465034698ec2.TransitiveMemberOfRequestBuilder) {
    return id382f452535c2e36f3f851dd3aeecca79c528c74c33433e51efa465034698ec2.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOfById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.devices.item.transitiveMemberOf.item collection
func (m *DeviceItemRequestBuilder) TransitiveMemberOfById(id string)(*i81b3cccc98b5736dc9541adae22c6bfe0b07307fa1ee207c5666690a14f4d776.DirectoryObjectItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["directoryObject_id"] = id
    }
    return i81b3cccc98b5736dc9541adae22c6bfe0b07307fa1ee207c5666690a14f4d776.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
