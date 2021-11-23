package item

import (
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
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
    ib133a87f1f6aa82508a13fd26e99937ae943d6e84c9447f4c022ffc647194620 "github.com/microsoftgraph/msgraph-sdk-go/devices/item/extensions/item"
)

// DeviceRequestBuilder builds and executes requests for operations under \devices\{device-id}
type DeviceRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceRequestBuilderDeleteOptions options for Delete
type DeviceRequestBuilderDeleteOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceRequestBuilderGetOptions options for Get
type DeviceRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceRequestBuilderGetQueryParameters get entity from devices by key
type DeviceRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// DeviceRequestBuilderPatchOptions options for Patch
type DeviceRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Device;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceRequestBuilder) CheckMemberGroups()(*ie61880210a289f885eb584ae4ffb8775884acae6cbcc7f3ef5c31226565fbb49.CheckMemberGroupsRequestBuilder) {
    return ie61880210a289f885eb584ae4ffb8775884acae6cbcc7f3ef5c31226565fbb49.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) CheckMemberObjects()(*ie01af685f90d1d6390f1cca50e6d6a4a0c0d1400e061efa09108ed2cc4f121ce.CheckMemberObjectsRequestBuilder) {
    return ie01af685f90d1d6390f1cca50e6d6a4a0c0d1400e061efa09108ed2cc4f121ce.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceRequestBuilderInternal instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceRequestBuilder) {
    m := &DeviceRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/devices/{device_id}{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceRequestBuilder instantiates a new DeviceRequestBuilder and sets the default values.
func NewDeviceRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete entity from devices
func (m *DeviceRequestBuilder) CreateDeleteRequestInformation(options *DeviceRequestBuilderDeleteOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceRequestBuilder) CreateGetRequestInformation(options *DeviceRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceRequestBuilder) CreatePatchRequestInformation(options *DeviceRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceRequestBuilder) Delete(options *DeviceRequestBuilderDeleteOptions)(error) {
    requestInfo, err := m.CreateDeleteRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceRequestBuilder) Extensions()(*i6d9ea1deb1d3cfaac0245615a52427d153dc3b9897295d48d6cb94b90f2b2bb6.ExtensionsRequestBuilder) {
    return i6d9ea1deb1d3cfaac0245615a52427d153dc3b9897295d48d6cb94b90f2b2bb6.NewExtensionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExtensionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.devices.item.extensions.item collection
func (m *DeviceRequestBuilder) ExtensionsById(id string)(*ib133a87f1f6aa82508a13fd26e99937ae943d6e84c9447f4c022ffc647194620.ExtensionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["extension_id"] = id
    }
    return ib133a87f1f6aa82508a13fd26e99937ae943d6e84c9447f4c022ffc647194620.NewExtensionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get entity from devices by key
func (m *DeviceRequestBuilder) Get(options *DeviceRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Device, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDevice() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.Device), nil
}
func (m *DeviceRequestBuilder) GetMemberGroups()(*ifad36aa7c4e70f6ef3f9d729bcb5622ac992a86d8192af5b3f76265c38bc1854.GetMemberGroupsRequestBuilder) {
    return ifad36aa7c4e70f6ef3f9d729bcb5622ac992a86d8192af5b3f76265c38bc1854.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) GetMemberObjects()(*i802cbfeae38ef2d92747a6d0a312423fd41be33f3e3fb965305e643ae538adfa.GetMemberObjectsRequestBuilder) {
    return i802cbfeae38ef2d92747a6d0a312423fd41be33f3e3fb965305e643ae538adfa.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) MemberOf()(*i1da16a4b620766997fcb2773cf780ba737b8dd64438f6bd1b46feb6cf67dfaf2.MemberOfRequestBuilder) {
    return i1da16a4b620766997fcb2773cf780ba737b8dd64438f6bd1b46feb6cf67dfaf2.NewMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Patch update entity in devices
func (m *DeviceRequestBuilder) Patch(options *DeviceRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceRequestBuilder) RegisteredOwners()(*iad59335f11b6e4f287fa8e2535c4c39e6d1968fb30e6d9729d87370eaf1a741c.RegisteredOwnersRequestBuilder) {
    return iad59335f11b6e4f287fa8e2535c4c39e6d1968fb30e6d9729d87370eaf1a741c.NewRegisteredOwnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) RegisteredUsers()(*id4c71a00690f15e834a677b658e9a8fc4ab7659038027b563d609b48b71d76bc.RegisteredUsersRequestBuilder) {
    return id4c71a00690f15e834a677b658e9a8fc4ab7659038027b563d609b48b71d76bc.NewRegisteredUsersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) Restore()(*i7f0d40d46e36fb24c086809fc583560c84c2d578585740e724f14d6bab94b6ac.RestoreRequestBuilder) {
    return i7f0d40d46e36fb24c086809fc583560c84c2d578585740e724f14d6bab94b6ac.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceRequestBuilder) TransitiveMemberOf()(*id382f452535c2e36f3f851dd3aeecca79c528c74c33433e51efa465034698ec2.TransitiveMemberOfRequestBuilder) {
    return id382f452535c2e36f3f851dd3aeecca79c528c74c33433e51efa465034698ec2.NewTransitiveMemberOfRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
