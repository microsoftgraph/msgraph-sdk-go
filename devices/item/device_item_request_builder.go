package item

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    i1da16a4b620766997fcb2773cf780ba737b8dd64438f6bd1b46feb6cf67dfaf2 "github.com/microsoftgraph/msgraph-sdk-go/devices/item/memberof"
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
    pathParameters map[string]string
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter
    // Url template to use to build the URL for the current request builder
    urlTemplate string
}
// DeviceItemRequestBuilderDeleteRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceItemRequestBuilderDeleteRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// DeviceItemRequestBuilderGetQueryParameters get device
type DeviceItemRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string `uriparametername:"%24expand"`
    // Select properties to be returned
    Select []string `uriparametername:"%24select"`
}
// DeviceItemRequestBuilderGetRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceItemRequestBuilderGetRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
    // Request query parameters
    QueryParameters *DeviceItemRequestBuilderGetQueryParameters
}
// DeviceItemRequestBuilderPatchRequestConfiguration configuration for the request such as headers, query parameters, and middleware options.
type DeviceItemRequestBuilderPatchRequestConfiguration struct {
    // Request headers
    Headers map[string]string
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption
}
// CheckMemberGroups the checkMemberGroups property
func (m *DeviceItemRequestBuilder) CheckMemberGroups()(*ie61880210a289f885eb584ae4ffb8775884acae6cbcc7f3ef5c31226565fbb49.CheckMemberGroupsRequestBuilder) {
    return ie61880210a289f885eb584ae4ffb8775884acae6cbcc7f3ef5c31226565fbb49.NewCheckMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// CheckMemberObjects the checkMemberObjects property
func (m *DeviceItemRequestBuilder) CheckMemberObjects()(*ie01af685f90d1d6390f1cca50e6d6a4a0c0d1400e061efa09108ed2cc4f121ce.CheckMemberObjectsRequestBuilder) {
    return ie01af685f90d1d6390f1cca50e6d6a4a0c0d1400e061efa09108ed2cc4f121ce.NewCheckMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceItemRequestBuilderInternal instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceItemRequestBuilder) {
    m := &DeviceItemRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/devices/{device%2Did}{?%24select,%24expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceItemRequestBuilder instantiates a new DeviceItemRequestBuilder and sets the default values.
func NewDeviceItemRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceItemRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceItemRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateDeleteRequestInformation delete device
func (m *DeviceItemRequestBuilder) CreateDeleteRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateDeleteRequestInformationWithRequestConfiguration(nil);
}
// CreateDeleteRequestInformationWithRequestConfiguration delete device
func (m *DeviceItemRequestBuilder) CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration *DeviceItemRequestBuilderDeleteRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.DELETE
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreateGetRequestInformation get device
func (m *DeviceItemRequestBuilder) CreateGetRequestInformation()(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreateGetRequestInformationWithRequestConfiguration(nil);
}
// CreateGetRequestInformationWithRequestConfiguration get device
func (m *DeviceItemRequestBuilder) CreateGetRequestInformationWithRequestConfiguration(requestConfiguration *DeviceItemRequestBuilderGetRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if requestConfiguration != nil {
        if requestConfiguration.QueryParameters != nil {
            requestInfo.AddQueryParameters(*(requestConfiguration.QueryParameters))
        }
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update device
func (m *DeviceItemRequestBuilder) CreatePatchRequestInformation(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    return m.CreatePatchRequestInformationWithRequestConfiguration(body, nil);
}
// CreatePatchRequestInformationWithRequestConfiguration update device
func (m *DeviceItemRequestBuilder) CreatePatchRequestInformationWithRequestConfiguration(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable, requestConfiguration *DeviceItemRequestBuilderPatchRequestConfiguration)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", body)
    if requestConfiguration != nil {
        requestInfo.AddRequestHeaders(requestConfiguration.Headers)
        requestInfo.AddRequestOptions(requestConfiguration.Options)
    }
    return requestInfo, nil
}
// Delete delete device
func (m *DeviceItemRequestBuilder) Delete()(error) {
    return m.DeleteWithRequestConfigurationAndResponseHandler(nil, nil);
}
// DeleteWithRequestConfigurationAndResponseHandler delete device
func (m *DeviceItemRequestBuilder) DeleteWithRequestConfigurationAndResponseHandler(requestConfiguration *DeviceItemRequestBuilderDeleteRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreateDeleteRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// Extensions the extensions property
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
        urlTplParams["extension%2Did"] = id
    }
    return ib133a87f1f6aa82508a13fd26e99937ae943d6e84c9447f4c022ffc647194620.NewExtensionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get device
func (m *DeviceItemRequestBuilder) Get()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable, error) {
    return m.GetWithRequestConfigurationAndResponseHandler(nil, nil);
}
// GetMemberGroups the getMemberGroups property
func (m *DeviceItemRequestBuilder) GetMemberGroups()(*ifad36aa7c4e70f6ef3f9d729bcb5622ac992a86d8192af5b3f76265c38bc1854.GetMemberGroupsRequestBuilder) {
    return ifad36aa7c4e70f6ef3f9d729bcb5622ac992a86d8192af5b3f76265c38bc1854.NewGetMemberGroupsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetMemberObjects the getMemberObjects property
func (m *DeviceItemRequestBuilder) GetMemberObjects()(*i802cbfeae38ef2d92747a6d0a312423fd41be33f3e3fb965305e643ae538adfa.GetMemberObjectsRequestBuilder) {
    return i802cbfeae38ef2d92747a6d0a312423fd41be33f3e3fb965305e643ae538adfa.NewGetMemberObjectsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// GetWithRequestConfigurationAndResponseHandler get device
func (m *DeviceItemRequestBuilder) GetWithRequestConfigurationAndResponseHandler(requestConfiguration *DeviceItemRequestBuilderGetRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable, error) {
    requestInfo, err := m.CreateGetRequestInformationWithRequestConfiguration(requestConfiguration);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceFromDiscriminatorValue, responseHandler, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable), nil
}
// MemberOf the memberOf property
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
        urlTplParams["directoryObject%2Did"] = id
    }
    return i3510737f0d62ae1028132f9eb8460789b2745e3f0425ffac6d00ba945e3ad462.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update device
func (m *DeviceItemRequestBuilder) Patch(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable)(error) {
    return m.PatchWithRequestConfigurationAndResponseHandler(body, nil, nil);
}
// PatchWithRequestConfigurationAndResponseHandler update device
func (m *DeviceItemRequestBuilder) PatchWithRequestConfigurationAndResponseHandler(body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Deviceable, requestConfiguration *DeviceItemRequestBuilderPatchRequestConfiguration, responseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler)(error) {
    requestInfo, err := m.CreatePatchRequestInformationWithRequestConfiguration(body, requestConfiguration);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, responseHandler, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// RegisteredOwners the registeredOwners property
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
        urlTplParams["directoryObject%2Did"] = id
    }
    return ib4628a5555deb2b60e49a55e75b33d25c1aa9f9e4a77c4c8fd1022fda5ae1027.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// RegisteredUsers the registeredUsers property
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
        urlTplParams["directoryObject%2Did"] = id
    }
    return ifc731e509397c6e1502bcc9cb553695d7c0475072d83d49e638a8991044f9333.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Restore the restore property
func (m *DeviceItemRequestBuilder) Restore()(*i7f0d40d46e36fb24c086809fc583560c84c2d578585740e724f14d6bab94b6ac.RestoreRequestBuilder) {
    return i7f0d40d46e36fb24c086809fc583560c84c2d578585740e724f14d6bab94b6ac.NewRestoreRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TransitiveMemberOf the transitiveMemberOf property
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
        urlTplParams["directoryObject%2Did"] = id
    }
    return i81b3cccc98b5736dc9541adae22c6bfe0b07307fa1ee207c5666690a14f4d776.NewDirectoryObjectItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
