package devicemanagement

import (
    i0266cdabed6f940e582824f4953f60b1645b8dfea310dab6d5889d756b92f4b0 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/windowsinformationprotectionapplearningsummaries"
    i0681bab6d041a7f3f1a8ffa28dad2cd4973508c77be81d76cf2a90ef33eff0a0 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies"
    i06eed0fc70646ea797036b22a3bd208dcf64b957f60fbda2b7ac9d010b7feed0 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecategories"
    i12358d3d226408664d7e188cc82af4e061e0ba8f4ebdc93218010060e37e15e3 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/mobilethreatdefenseconnectors"
    i18050da269eccd608dfc5031d3fb417f936e4af9b764b9d4afb41d4ecadb9c5f "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/termsandconditions"
    i1974d70a31b48f13ed4aac419349cec2f881b01919f37dc352295fa6abc079f1 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/conditionalaccesssettings"
    i1ec63f108149e3b4b045393e48eccc7378122447f0242c3db6494322d329aff1 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/applepushnotificationcertificate"
    i270dc3fe87992ef7b99e2d60053e24d7e2bab1d9be5017e134d110bdeaaa32ad "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/compliancemanagementpartners"
    i2c2e90063f13b9e4a427c1b86eaf3e4609a5d3204f44f9a85d5b85e11c1df77e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/iosupdatestatuses"
    i2eb7d01d7645dcb50edbcee3c098f918d1036928f83217ac1ac02f2fd20a2c72 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddeviceoverview"
    i389fcbbe90a623cbb16365be54b2030e89b2f6a0e72f74e34fda6cdaeda3ed12 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/windowsautopilotdeviceidentities"
    i460ebd1d2d6c5576fc5e23bb59098a0eb4d16ebe0bf39c0b86508c13edbb1c20 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/resourceoperations"
    i4eb3be0854666ed4ad5a0068cb9a3dd50a4d96ac7cb6d48a532cf4053c9e7002 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/windowsinformationprotectionnetworklearningsummaries"
    i53b822adbf609100b4c0e4ed0fa96d2a112352ee03fab86eef01c0ffdf52bb08 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/detectedapps"
    i65c9665b2d065131ce96a9c18f503bfb2e2adc753f3c09bf535a5f2579bbf49b "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicysettingstatesummaries"
    i6a7a598519256477a4e2b50bf57c67a59f312f1cbb168239d5feb82a1ee9146e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/notificationmessagetemplates"
    i6b79e19af80218ca443b0983adbe0f6cdef30c9d057c37be983e59c74074ca29 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/geteffectivepermissionswithscope"
    i73c8f465ab045c16e61d98d4347d7f15ca49da6c2a5cb12e22c023adc65d9bbc "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/roleassignments"
    i8f89db33a5485e8a8a34a95327cdfc0a7a1b58b0952eae07eddb4d11f79c493d "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/troubleshootingevents"
    i92707a6c0ad5ca796fd92da8575f9bdad9b5f090ed01e195bd0330b2ea0d0e5e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/softwareupdatestatussummary"
    ib5725579489417002b6356eaf6b46c1969fd077fbc933471d444b2c79382db89 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/reports"
    ibdcf6dbbc28e3b36e6fdfb9649205e10fa11d567a3050f0921b23fecaeed9306 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/remoteassistancepartners"
    ic0282b591819c9e922d4e2dcb24aea19571278605e12e63f9c5f2e43a4012aca "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations"
    ic0a58ac6928c42577a17fd25a407afbfa185dc576ccf5d2cfbd84c4e49df93cf "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/exchangeconnectors"
    ic0b66042facca8076e7f66ab7a38bd7270fa31999f666eed03c79fe9e751c55c "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceenrollmentconfigurations"
    icb8e98ee51c1d5e9dbd9f090345642c6d2cdaefb80750cbd8acf8765769c5b3b "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/telecomexpensemanagementpartners"
    id29dcfda2bb98b4b7cf7bd4f6740f350afac6e45624c646636f5fc440f552695 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/importedwindowsautopilotdeviceidentities"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie4810305716cc404b170d5b62f311334bd3f6860e0ed7090a75dddd98b296a7c "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/roledefinitions"
    iedc86b21fb6133d4138c5e70da6ac898141cde17633742a054dc9c2b4d56ccdc "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices"
    ieed97de3e0515fd5b8f5666f0a60dba2083d4943c18cc0bfe46efaf6882e12db "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurationdevicestatesummaries"
    ifacc6e78462b6602604fdba95c2c687304c070fc8161931add6511708842e0df "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/verifywindowsenrollmentautodiscoverywithdomainname"
    ifeb490582be8577fe0d05594823855d35edb357366160e9ad7c2198d3cbbb980 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicemanagementpartners"
    iff68b83db7285c0f9536fd151184fbcf3100e0776163008e6628f613e5e93c83 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicydevicestatesummary"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i0e012f8f6f74902658fef78becd57ef59cb2e332f0de1d45c2eb7b0da3bc8d8d "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicemanagementpartners/item"
    i0fb841dc54b62513276e2131192e7fde71d44b2cc7e9032d7989e15cac59cfd5 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/compliancemanagementpartners/item"
    i17dd1ea085fa58e7adff79de1f08a20011914cbfc90ae645da10d72117831db1 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecategories/item"
    i1fc4baac6aed0f5be3c916f9eb09379061c3a0c2c15a9b52bc4922b4bc48c6e1 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/detectedapps/item"
    i30ffad8da959c490ac0034c48052fa3e14da33a4e5d2bc2fac41196f8e97e9aa "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/remoteassistancepartners/item"
    i33db959724eee2740c1918f5227e5ccdeb2c6a3b9d32f0aa3121e2994bcf87c6 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/resourceoperations/item"
    i348e6a7178caf047ba7d4e11d4755f76b9aebf8f8a3a4158d608d5a6ce273a2a "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/notificationmessagetemplates/item"
    i48632af9e52103d7e924df3e7b9a0f847596ec0c7c5ee7bd34e14d9d8bb1a93a "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/windowsautopilotdeviceidentities/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4d3a26b3bf54ee27026f84d5aed96158baa8d422d5fa52c1cd1a2174a2ef12c3 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/iosupdatestatuses/item"
    i554fa2b496bcd866a45704e8e6e4128b82ee5ee0b1181d0028299b802dd9cbdd "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/telecomexpensemanagementpartners/item"
    i7030d242a3cf4c114bf570a5d6552762b18b61c6c977d48cb287624088477d42 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/mobilethreatdefenseconnectors/item"
    i81293f477529880dfafa8305dbf38d6a41623848e064daba38d327cc34868b65 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicysettingstatesummaries/item"
    i95a086ba2ba8d75990807add25a462bbd72691c366888f645759e03b6f39cc65 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/roledefinitions/item"
    ia6325065d45340cbac90dee05565ded24dc65b6c714bd5d8bf52762b7aa42658 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/windowsinformationprotectionapplearningsummaries/item"
    iaad2c116436cc4b77a7ef65200737e8e7a867f90393f50d4c0f57fa1b022faab "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/windowsinformationprotectionnetworklearningsummaries/item"
    iad7b7c15ef9e96a638e142140a635d85409af90a1940df3693fa29e45e2009f1 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/importedwindowsautopilotdeviceidentities/item"
    ib3773734d93a499f579b7c8c42e59796a00cf61c2984fc1d034ceff70f84d6cb "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceconfigurations/item"
    ib49f8e333a9996485568d4de1f3efae787c0dbde2d33bdbd2a863ae5e110192c "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/troubleshootingevents/item"
    ibffb66fdbfcbe5e9b502701309e758010a29f241735933337a460c17e62d9581 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/devicecompliancepolicies/item"
    ic17628b2ada33156bca6f6158e18f62f1027784aa76d9f9a69a4d0cee7e57f3e "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/manageddevices/item"
    ic74e3750aadb0193930fedc0db367bfd59bac511f5b7371ed13d262768611421 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/deviceenrollmentconfigurations/item"
    idd0a88efd8948b1f5076671630001feacd82e5f59287364da79d7d9ea14e73b1 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/roleassignments/item"
    ie2787738d6d6382b33609c61cb384a808f64e6df1c1b03c2b26417c950b5b448 "github.com/microsoftgraph/msgraph-sdk-go/devicemanagement/exchangeconnectors/item"
)

// DeviceManagementRequestBuilder builds and executes requests for operations under \deviceManagement
type DeviceManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceManagementRequestBuilderGetOptions options for Get
type DeviceManagementRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceManagementRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// DeviceManagementRequestBuilderGetQueryParameters get deviceManagement
type DeviceManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceManagementRequestBuilderPatchOptions options for Patch
type DeviceManagementRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceManagement;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceManagementRequestBuilder) ApplePushNotificationCertificate()(*i1ec63f108149e3b4b045393e48eccc7378122447f0242c3db6494322d329aff1.ApplePushNotificationCertificateRequestBuilder) {
    return i1ec63f108149e3b4b045393e48eccc7378122447f0242c3db6494322d329aff1.NewApplePushNotificationCertificateRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) ComplianceManagementPartners()(*i270dc3fe87992ef7b99e2d60053e24d7e2bab1d9be5017e134d110bdeaaa32ad.ComplianceManagementPartnersRequestBuilder) {
    return i270dc3fe87992ef7b99e2d60053e24d7e2bab1d9be5017e134d110bdeaaa32ad.NewComplianceManagementPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ComplianceManagementPartnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.complianceManagementPartners.item collection
func (m *DeviceManagementRequestBuilder) ComplianceManagementPartnersById(id string)(*i0fb841dc54b62513276e2131192e7fde71d44b2cc7e9032d7989e15cac59cfd5.ComplianceManagementPartnerRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["complianceManagementPartner_id"] = id
    }
    return i0fb841dc54b62513276e2131192e7fde71d44b2cc7e9032d7989e15cac59cfd5.NewComplianceManagementPartnerRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) ConditionalAccessSettings()(*i1974d70a31b48f13ed4aac419349cec2f881b01919f37dc352295fa6abc079f1.ConditionalAccessSettingsRequestBuilder) {
    return i1974d70a31b48f13ed4aac419349cec2f881b01919f37dc352295fa6abc079f1.NewConditionalAccessSettingsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NewDeviceManagementRequestBuilderInternal instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementRequestBuilder) {
    m := &DeviceManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceManagement{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceManagementRequestBuilder instantiates a new DeviceManagementRequestBuilder and sets the default values.
func NewDeviceManagementRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get deviceManagement
func (m *DeviceManagementRequestBuilder) CreateGetRequestInformation(options *DeviceManagementRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
// CreatePatchRequestInformation update deviceManagement
func (m *DeviceManagementRequestBuilder) CreatePatchRequestInformation(options *DeviceManagementRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceManagementRequestBuilder) DetectedApps()(*i53b822adbf609100b4c0e4ed0fa96d2a112352ee03fab86eef01c0ffdf52bb08.DetectedAppsRequestBuilder) {
    return i53b822adbf609100b4c0e4ed0fa96d2a112352ee03fab86eef01c0ffdf52bb08.NewDetectedAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DetectedAppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.detectedApps.item collection
func (m *DeviceManagementRequestBuilder) DetectedAppsById(id string)(*i1fc4baac6aed0f5be3c916f9eb09379061c3a0c2c15a9b52bc4922b4bc48c6e1.DetectedAppRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["detectedApp_id"] = id
    }
    return i1fc4baac6aed0f5be3c916f9eb09379061c3a0c2c15a9b52bc4922b4bc48c6e1.NewDetectedAppRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) DeviceCategories()(*i06eed0fc70646ea797036b22a3bd208dcf64b957f60fbda2b7ac9d010b7feed0.DeviceCategoriesRequestBuilder) {
    return i06eed0fc70646ea797036b22a3bd208dcf64b957f60fbda2b7ac9d010b7feed0.NewDeviceCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCategoriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCategories.item collection
func (m *DeviceManagementRequestBuilder) DeviceCategoriesById(id string)(*i17dd1ea085fa58e7adff79de1f08a20011914cbfc90ae645da10d72117831db1.DeviceCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCategory_id"] = id
    }
    return i17dd1ea085fa58e7adff79de1f08a20011914cbfc90ae645da10d72117831db1.NewDeviceCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicies()(*i0681bab6d041a7f3f1a8ffa28dad2cd4973508c77be81d76cf2a90ef33eff0a0.DeviceCompliancePoliciesRequestBuilder) {
    return i0681bab6d041a7f3f1a8ffa28dad2cd4973508c77be81d76cf2a90ef33eff0a0.NewDeviceCompliancePoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicies.item collection
func (m *DeviceManagementRequestBuilder) DeviceCompliancePoliciesById(id string)(*ibffb66fdbfcbe5e9b502701309e758010a29f241735933337a460c17e62d9581.DeviceCompliancePolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicy_id"] = id
    }
    return ibffb66fdbfcbe5e9b502701309e758010a29f241735933337a460c17e62d9581.NewDeviceCompliancePolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicyDeviceStateSummary()(*iff68b83db7285c0f9536fd151184fbcf3100e0776163008e6628f613e5e93c83.DeviceCompliancePolicyDeviceStateSummaryRequestBuilder) {
    return iff68b83db7285c0f9536fd151184fbcf3100e0776163008e6628f613e5e93c83.NewDeviceCompliancePolicyDeviceStateSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicySettingStateSummaries()(*i65c9665b2d065131ce96a9c18f503bfb2e2adc753f3c09bf535a5f2579bbf49b.DeviceCompliancePolicySettingStateSummariesRequestBuilder) {
    return i65c9665b2d065131ce96a9c18f503bfb2e2adc753f3c09bf535a5f2579bbf49b.NewDeviceCompliancePolicySettingStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceCompliancePolicySettingStateSummariesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceCompliancePolicySettingStateSummaries.item collection
func (m *DeviceManagementRequestBuilder) DeviceCompliancePolicySettingStateSummariesById(id string)(*i81293f477529880dfafa8305dbf38d6a41623848e064daba38d327cc34868b65.DeviceCompliancePolicySettingStateSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceCompliancePolicySettingStateSummary_id"] = id
    }
    return i81293f477529880dfafa8305dbf38d6a41623848e064daba38d327cc34868b65.NewDeviceCompliancePolicySettingStateSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) DeviceConfigurationDeviceStateSummaries()(*ieed97de3e0515fd5b8f5666f0a60dba2083d4943c18cc0bfe46efaf6882e12db.DeviceConfigurationDeviceStateSummariesRequestBuilder) {
    return ieed97de3e0515fd5b8f5666f0a60dba2083d4943c18cc0bfe46efaf6882e12db.NewDeviceConfigurationDeviceStateSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) DeviceConfigurations()(*ic0282b591819c9e922d4e2dcb24aea19571278605e12e63f9c5f2e43a4012aca.DeviceConfigurationsRequestBuilder) {
    return ic0282b591819c9e922d4e2dcb24aea19571278605e12e63f9c5f2e43a4012aca.NewDeviceConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceConfigurations.item collection
func (m *DeviceManagementRequestBuilder) DeviceConfigurationsById(id string)(*ib3773734d93a499f579b7c8c42e59796a00cf61c2984fc1d034ceff70f84d6cb.DeviceConfigurationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceConfiguration_id"] = id
    }
    return ib3773734d93a499f579b7c8c42e59796a00cf61c2984fc1d034ceff70f84d6cb.NewDeviceConfigurationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) DeviceEnrollmentConfigurations()(*ic0b66042facca8076e7f66ab7a38bd7270fa31999f666eed03c79fe9e751c55c.DeviceEnrollmentConfigurationsRequestBuilder) {
    return ic0b66042facca8076e7f66ab7a38bd7270fa31999f666eed03c79fe9e751c55c.NewDeviceEnrollmentConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceEnrollmentConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceEnrollmentConfigurations.item collection
func (m *DeviceManagementRequestBuilder) DeviceEnrollmentConfigurationsById(id string)(*ic74e3750aadb0193930fedc0db367bfd59bac511f5b7371ed13d262768611421.DeviceEnrollmentConfigurationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceEnrollmentConfiguration_id"] = id
    }
    return ic74e3750aadb0193930fedc0db367bfd59bac511f5b7371ed13d262768611421.NewDeviceEnrollmentConfigurationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) DeviceManagementPartners()(*ifeb490582be8577fe0d05594823855d35edb357366160e9ad7c2198d3cbbb980.DeviceManagementPartnersRequestBuilder) {
    return ifeb490582be8577fe0d05594823855d35edb357366160e9ad7c2198d3cbbb980.NewDeviceManagementPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DeviceManagementPartnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.deviceManagementPartners.item collection
func (m *DeviceManagementRequestBuilder) DeviceManagementPartnersById(id string)(*i0e012f8f6f74902658fef78becd57ef59cb2e332f0de1d45c2eb7b0da3bc8d8d.DeviceManagementPartnerRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementPartner_id"] = id
    }
    return i0e012f8f6f74902658fef78becd57ef59cb2e332f0de1d45c2eb7b0da3bc8d8d.NewDeviceManagementPartnerRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) ExchangeConnectors()(*ic0a58ac6928c42577a17fd25a407afbfa185dc576ccf5d2cfbd84c4e49df93cf.ExchangeConnectorsRequestBuilder) {
    return ic0a58ac6928c42577a17fd25a407afbfa185dc576ccf5d2cfbd84c4e49df93cf.NewExchangeConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ExchangeConnectorsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.exchangeConnectors.item collection
func (m *DeviceManagementRequestBuilder) ExchangeConnectorsById(id string)(*ie2787738d6d6382b33609c61cb384a808f64e6df1c1b03c2b26417c950b5b448.DeviceManagementExchangeConnectorRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementExchangeConnector_id"] = id
    }
    return ie2787738d6d6382b33609c61cb384a808f64e6df1c1b03c2b26417c950b5b448.NewDeviceManagementExchangeConnectorRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get deviceManagement
func (m *DeviceManagementRequestBuilder) Get(options *DeviceManagementRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceManagement, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDeviceManagement() }, nil, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceManagement), nil
}
// GetEffectivePermissionsWithScope builds and executes requests for operations under \deviceManagement\microsoft.graph.getEffectivePermissions(scope='{scope}')
func (m *DeviceManagementRequestBuilder) GetEffectivePermissionsWithScope(scope *string)(*i6b79e19af80218ca443b0983adbe0f6cdef30c9d057c37be983e59c74074ca29.GetEffectivePermissionsWithScopeRequestBuilder) {
    return i6b79e19af80218ca443b0983adbe0f6cdef30c9d057c37be983e59c74074ca29.NewGetEffectivePermissionsWithScopeRequestBuilderInternal(m.pathParameters, m.requestAdapter, scope);
}
func (m *DeviceManagementRequestBuilder) ImportedWindowsAutopilotDeviceIdentities()(*id29dcfda2bb98b4b7cf7bd4f6740f350afac6e45624c646636f5fc440f552695.ImportedWindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return id29dcfda2bb98b4b7cf7bd4f6740f350afac6e45624c646636f5fc440f552695.NewImportedWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ImportedWindowsAutopilotDeviceIdentitiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.importedWindowsAutopilotDeviceIdentities.item collection
func (m *DeviceManagementRequestBuilder) ImportedWindowsAutopilotDeviceIdentitiesById(id string)(*iad7b7c15ef9e96a638e142140a635d85409af90a1940df3693fa29e45e2009f1.ImportedWindowsAutopilotDeviceIdentityRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["importedWindowsAutopilotDeviceIdentity_id"] = id
    }
    return iad7b7c15ef9e96a638e142140a635d85409af90a1940df3693fa29e45e2009f1.NewImportedWindowsAutopilotDeviceIdentityRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) IosUpdateStatuses()(*i2c2e90063f13b9e4a427c1b86eaf3e4609a5d3204f44f9a85d5b85e11c1df77e.IosUpdateStatusesRequestBuilder) {
    return i2c2e90063f13b9e4a427c1b86eaf3e4609a5d3204f44f9a85d5b85e11c1df77e.NewIosUpdateStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IosUpdateStatusesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.iosUpdateStatuses.item collection
func (m *DeviceManagementRequestBuilder) IosUpdateStatusesById(id string)(*i4d3a26b3bf54ee27026f84d5aed96158baa8d422d5fa52c1cd1a2174a2ef12c3.IosUpdateDeviceStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosUpdateDeviceStatus_id"] = id
    }
    return i4d3a26b3bf54ee27026f84d5aed96158baa8d422d5fa52c1cd1a2174a2ef12c3.NewIosUpdateDeviceStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) ManagedDeviceOverview()(*i2eb7d01d7645dcb50edbcee3c098f918d1036928f83217ac1ac02f2fd20a2c72.ManagedDeviceOverviewRequestBuilder) {
    return i2eb7d01d7645dcb50edbcee3c098f918d1036928f83217ac1ac02f2fd20a2c72.NewManagedDeviceOverviewRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) ManagedDevices()(*iedc86b21fb6133d4138c5e70da6ac898141cde17633742a054dc9c2b4d56ccdc.ManagedDevicesRequestBuilder) {
    return iedc86b21fb6133d4138c5e70da6ac898141cde17633742a054dc9c2b4d56ccdc.NewManagedDevicesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedDevicesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.managedDevices.item collection
func (m *DeviceManagementRequestBuilder) ManagedDevicesById(id string)(*ic17628b2ada33156bca6f6158e18f62f1027784aa76d9f9a69a4d0cee7e57f3e.ManagedDeviceRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDevice_id"] = id
    }
    return ic17628b2ada33156bca6f6158e18f62f1027784aa76d9f9a69a4d0cee7e57f3e.NewManagedDeviceRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) MobileThreatDefenseConnectors()(*i12358d3d226408664d7e188cc82af4e061e0ba8f4ebdc93218010060e37e15e3.MobileThreatDefenseConnectorsRequestBuilder) {
    return i12358d3d226408664d7e188cc82af4e061e0ba8f4ebdc93218010060e37e15e3.NewMobileThreatDefenseConnectorsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileThreatDefenseConnectorsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.mobileThreatDefenseConnectors.item collection
func (m *DeviceManagementRequestBuilder) MobileThreatDefenseConnectorsById(id string)(*i7030d242a3cf4c114bf570a5d6552762b18b61c6c977d48cb287624088477d42.MobileThreatDefenseConnectorRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileThreatDefenseConnector_id"] = id
    }
    return i7030d242a3cf4c114bf570a5d6552762b18b61c6c977d48cb287624088477d42.NewMobileThreatDefenseConnectorRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) NotificationMessageTemplates()(*i6a7a598519256477a4e2b50bf57c67a59f312f1cbb168239d5feb82a1ee9146e.NotificationMessageTemplatesRequestBuilder) {
    return i6a7a598519256477a4e2b50bf57c67a59f312f1cbb168239d5feb82a1ee9146e.NewNotificationMessageTemplatesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// NotificationMessageTemplatesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.notificationMessageTemplates.item collection
func (m *DeviceManagementRequestBuilder) NotificationMessageTemplatesById(id string)(*i348e6a7178caf047ba7d4e11d4755f76b9aebf8f8a3a4158d608d5a6ce273a2a.NotificationMessageTemplateRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["notificationMessageTemplate_id"] = id
    }
    return i348e6a7178caf047ba7d4e11d4755f76b9aebf8f8a3a4158d608d5a6ce273a2a.NewNotificationMessageTemplateRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update deviceManagement
func (m *DeviceManagementRequestBuilder) Patch(options *DeviceManagementRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    err = m.requestAdapter.SendNoContentAsync(*requestInfo, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
func (m *DeviceManagementRequestBuilder) RemoteAssistancePartners()(*ibdcf6dbbc28e3b36e6fdfb9649205e10fa11d567a3050f0921b23fecaeed9306.RemoteAssistancePartnersRequestBuilder) {
    return ibdcf6dbbc28e3b36e6fdfb9649205e10fa11d567a3050f0921b23fecaeed9306.NewRemoteAssistancePartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RemoteAssistancePartnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.remoteAssistancePartners.item collection
func (m *DeviceManagementRequestBuilder) RemoteAssistancePartnersById(id string)(*i30ffad8da959c490ac0034c48052fa3e14da33a4e5d2bc2fac41196f8e97e9aa.RemoteAssistancePartnerRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["remoteAssistancePartner_id"] = id
    }
    return i30ffad8da959c490ac0034c48052fa3e14da33a4e5d2bc2fac41196f8e97e9aa.NewRemoteAssistancePartnerRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) Reports()(*ib5725579489417002b6356eaf6b46c1969fd077fbc933471d444b2c79382db89.ReportsRequestBuilder) {
    return ib5725579489417002b6356eaf6b46c1969fd077fbc933471d444b2c79382db89.NewReportsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) ResourceOperations()(*i460ebd1d2d6c5576fc5e23bb59098a0eb4d16ebe0bf39c0b86508c13edbb1c20.ResourceOperationsRequestBuilder) {
    return i460ebd1d2d6c5576fc5e23bb59098a0eb4d16ebe0bf39c0b86508c13edbb1c20.NewResourceOperationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ResourceOperationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.resourceOperations.item collection
func (m *DeviceManagementRequestBuilder) ResourceOperationsById(id string)(*i33db959724eee2740c1918f5227e5ccdeb2c6a3b9d32f0aa3121e2994bcf87c6.ResourceOperationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["resourceOperation_id"] = id
    }
    return i33db959724eee2740c1918f5227e5ccdeb2c6a3b9d32f0aa3121e2994bcf87c6.NewResourceOperationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) RoleAssignments()(*i73c8f465ab045c16e61d98d4347d7f15ca49da6c2a5cb12e22c023adc65d9bbc.RoleAssignmentsRequestBuilder) {
    return i73c8f465ab045c16e61d98d4347d7f15ca49da6c2a5cb12e22c023adc65d9bbc.NewRoleAssignmentsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleAssignmentsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.roleAssignments.item collection
func (m *DeviceManagementRequestBuilder) RoleAssignmentsById(id string)(*idd0a88efd8948b1f5076671630001feacd82e5f59287364da79d7d9ea14e73b1.DeviceAndAppManagementRoleAssignmentRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceAndAppManagementRoleAssignment_id"] = id
    }
    return idd0a88efd8948b1f5076671630001feacd82e5f59287364da79d7d9ea14e73b1.NewDeviceAndAppManagementRoleAssignmentRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) RoleDefinitions()(*ie4810305716cc404b170d5b62f311334bd3f6860e0ed7090a75dddd98b296a7c.RoleDefinitionsRequestBuilder) {
    return ie4810305716cc404b170d5b62f311334bd3f6860e0ed7090a75dddd98b296a7c.NewRoleDefinitionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// RoleDefinitionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.roleDefinitions.item collection
func (m *DeviceManagementRequestBuilder) RoleDefinitionsById(id string)(*i95a086ba2ba8d75990807add25a462bbd72691c366888f645759e03b6f39cc65.RoleDefinitionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["roleDefinition_id"] = id
    }
    return i95a086ba2ba8d75990807add25a462bbd72691c366888f645759e03b6f39cc65.NewRoleDefinitionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) SoftwareUpdateStatusSummary()(*i92707a6c0ad5ca796fd92da8575f9bdad9b5f090ed01e195bd0330b2ea0d0e5e.SoftwareUpdateStatusSummaryRequestBuilder) {
    return i92707a6c0ad5ca796fd92da8575f9bdad9b5f090ed01e195bd0330b2ea0d0e5e.NewSoftwareUpdateStatusSummaryRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) TelecomExpenseManagementPartners()(*icb8e98ee51c1d5e9dbd9f090345642c6d2cdaefb80750cbd8acf8765769c5b3b.TelecomExpenseManagementPartnersRequestBuilder) {
    return icb8e98ee51c1d5e9dbd9f090345642c6d2cdaefb80750cbd8acf8765769c5b3b.NewTelecomExpenseManagementPartnersRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TelecomExpenseManagementPartnersById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.telecomExpenseManagementPartners.item collection
func (m *DeviceManagementRequestBuilder) TelecomExpenseManagementPartnersById(id string)(*i554fa2b496bcd866a45704e8e6e4128b82ee5ee0b1181d0028299b802dd9cbdd.TelecomExpenseManagementPartnerRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["telecomExpenseManagementPartner_id"] = id
    }
    return i554fa2b496bcd866a45704e8e6e4128b82ee5ee0b1181d0028299b802dd9cbdd.NewTelecomExpenseManagementPartnerRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) TermsAndConditions()(*i18050da269eccd608dfc5031d3fb417f936e4af9b764b9d4afb41d4ecadb9c5f.TermsAndConditionsRequestBuilder) {
    return i18050da269eccd608dfc5031d3fb417f936e4af9b764b9d4afb41d4ecadb9c5f.NewTermsAndConditionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TermsAndConditionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.termsAndConditions.item collection
func (m *DeviceManagementRequestBuilder) TermsAndConditionsById(id string)(*i18050da269eccd608dfc5031d3fb417f936e4af9b764b9d4afb41d4ecadb9c5f.TermsAndConditionsRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["termsAndConditions_id"] = id
    }
    return i18050da269eccd608dfc5031d3fb417f936e4af9b764b9d4afb41d4ecadb9c5f.NewTermsAndConditionsRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) TroubleshootingEvents()(*i8f89db33a5485e8a8a34a95327cdfc0a7a1b58b0952eae07eddb4d11f79c493d.TroubleshootingEventsRequestBuilder) {
    return i8f89db33a5485e8a8a34a95327cdfc0a7a1b58b0952eae07eddb4d11f79c493d.NewTroubleshootingEventsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TroubleshootingEventsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.troubleshootingEvents.item collection
func (m *DeviceManagementRequestBuilder) TroubleshootingEventsById(id string)(*ib49f8e333a9996485568d4de1f3efae787c0dbde2d33bdbd2a863ae5e110192c.DeviceManagementTroubleshootingEventRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["deviceManagementTroubleshootingEvent_id"] = id
    }
    return ib49f8e333a9996485568d4de1f3efae787c0dbde2d33bdbd2a863ae5e110192c.NewDeviceManagementTroubleshootingEventRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VerifyWindowsEnrollmentAutoDiscoveryWithDomainName builds and executes requests for operations under \deviceManagement\microsoft.graph.verifyWindowsEnrollmentAutoDiscovery(domainName='{domainName}')
func (m *DeviceManagementRequestBuilder) VerifyWindowsEnrollmentAutoDiscoveryWithDomainName(domainName *string)(*ifacc6e78462b6602604fdba95c2c687304c070fc8161931add6511708842e0df.VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilder) {
    return ifacc6e78462b6602604fdba95c2c687304c070fc8161931add6511708842e0df.NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameRequestBuilderInternal(m.pathParameters, m.requestAdapter, domainName);
}
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeviceIdentities()(*i389fcbbe90a623cbb16365be54b2030e89b2f6a0e72f74e34fda6cdaeda3ed12.WindowsAutopilotDeviceIdentitiesRequestBuilder) {
    return i389fcbbe90a623cbb16365be54b2030e89b2f6a0e72f74e34fda6cdaeda3ed12.NewWindowsAutopilotDeviceIdentitiesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsAutopilotDeviceIdentitiesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.windowsAutopilotDeviceIdentities.item collection
func (m *DeviceManagementRequestBuilder) WindowsAutopilotDeviceIdentitiesById(id string)(*i48632af9e52103d7e924df3e7b9a0f847596ec0c7c5ee7bd34e14d9d8bb1a93a.WindowsAutopilotDeviceIdentityRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsAutopilotDeviceIdentity_id"] = id
    }
    return i48632af9e52103d7e924df3e7b9a0f847596ec0c7c5ee7bd34e14d9d8bb1a93a.NewWindowsAutopilotDeviceIdentityRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionAppLearningSummaries()(*i0266cdabed6f940e582824f4953f60b1645b8dfea310dab6d5889d756b92f4b0.WindowsInformationProtectionAppLearningSummariesRequestBuilder) {
    return i0266cdabed6f940e582824f4953f60b1645b8dfea310dab6d5889d756b92f4b0.NewWindowsInformationProtectionAppLearningSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionAppLearningSummariesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.windowsInformationProtectionAppLearningSummaries.item collection
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionAppLearningSummariesById(id string)(*ia6325065d45340cbac90dee05565ded24dc65b6c714bd5d8bf52762b7aa42658.WindowsInformationProtectionAppLearningSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionAppLearningSummary_id"] = id
    }
    return ia6325065d45340cbac90dee05565ded24dc65b6c714bd5d8bf52762b7aa42658.NewWindowsInformationProtectionAppLearningSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionNetworkLearningSummaries()(*i4eb3be0854666ed4ad5a0068cb9a3dd50a4d96ac7cb6d48a532cf4053c9e7002.WindowsInformationProtectionNetworkLearningSummariesRequestBuilder) {
    return i4eb3be0854666ed4ad5a0068cb9a3dd50a4d96ac7cb6d48a532cf4053c9e7002.NewWindowsInformationProtectionNetworkLearningSummariesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionNetworkLearningSummariesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceManagement.windowsInformationProtectionNetworkLearningSummaries.item collection
func (m *DeviceManagementRequestBuilder) WindowsInformationProtectionNetworkLearningSummariesById(id string)(*iaad2c116436cc4b77a7ef65200737e8e7a867f90393f50d4c0f57fa1b022faab.WindowsInformationProtectionNetworkLearningSummaryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionNetworkLearningSummary_id"] = id
    }
    return iaad2c116436cc4b77a7ef65200737e8e7a867f90393f50d4c0f57fa1b022faab.NewWindowsInformationProtectionNetworkLearningSummaryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
