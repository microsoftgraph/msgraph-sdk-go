package deviceappmanagement

import (
    i03d64ffdecc8f8ee64e8bdbbd21cd792ea749ccc7d6632ed57a9368c023aca62 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies"
    i18bf226bfeead5544b159f6b6a278878b7180022db16e17df961b8a912ac109f "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/defaultmanagedappprotections"
    i26923530adcc2c194b34fa3104b74fd0fd1849f0fcb7b138a8585bde57bca4d6 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks"
    i27d43f0c2c303b4996cf211399e808bbe9d9cd0c447c2d98061c3e3157225ee2 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappcategories"
    i2adc870208a84a1c715e640a35a322f60d13c7132faecab92cbdc856af0e50aa "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/androidmanagedappprotections"
    i32739bc83be1b44d9a0e5bef6898e52a1d08cff6911a2ca8739fb5aaf580c634 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/iosmanagedappprotections"
    i3c79f5992014bafe598fd20b6e9fd4b5f377e28ceccdbceb4002f89b8068136b "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/windowsinformationprotectionpolicies"
    i489c24d5b60b7229ec09f24405ed28097801dccb5fc9e368e362132b825c2808 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations"
    i6ca5a8c664cfb8f06d3237f44f8e2ec59140505f80b90c4cb4a1af3a15276d77 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/syncmicrosoftstoreforbusinessapps"
    ia84055a082d525b002332712c0df79f5545cff95fe4a4ba7647a1d711273c026 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations"
    iacf1c27653093eb9d818263eb22e73f2d1d6e2a5e231803fa86fe531b6e3f037 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/vpptokens"
    id47f41c3fc661963eced1051494f915ce9e3fc4f22d73157e3a706e3abf7c3c0 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileapps"
    id8e9dee048ce49d5be720313759ba7ff0e5561292831c5866a1251d1b6f75017 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations"
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9 "github.com/microsoft/kiota/abstractions/go"
    ie5dc06c3359896dbdaa15cf264f2382a825efa77f089fedc5fd0e259ac122847 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappstatuses"
    if1c061f5229e920421d3fbb9fc720eab490b00af016a531ace0cd07293f04818 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mdmwindowsinformationprotectionpolicies"
    i0076e81bdc67eb89b0f75e1d90ce618e38e43bf24fdd9abc2d2dd7b7b418df1c "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i090e5453c2fdf074fa2d7cb168a2f22ffe8f16816d015832f6cb242b663410b3 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item"
    i0f78fb508810d8046e0126c99b4e2f29dfbc726d712f29fb4c961a92e8630256 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappcategories/item"
    i3ab3bf13acf28406514cba3b07ab5235a9c5c97ff398576b733a15152cfb62d3 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/windowsinformationprotectionpolicies/item"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
    i4d54ef094838af59d772f617ee78a54ffbb57f062a912b927b484571e3b93ca5 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileapps/item"
    i560cf4a9f430034859b2e6c353602d2fd9280316fd3f2f1265995ae7216d4dc3 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies/item"
    i89d4341c9169ba82f4022fde15796f9a0594e04c5cc3dab59767eaa7b1fdf35c "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/androidmanagedappprotections/item"
    ia3c7a3f771ad6210672d21e0ac5e24781d4c319a31e3c1088a61103e6a6c19b8 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/defaultmanagedappprotections/item"
    ib2f42e9c4d8f657233687ebcb7ea2a7be72112a45159e6921784148695f33cc8 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/iosmanagedappprotections/item"
    ib471248babfa75d07aa6e6560ccb2ae999e37717a38cabb105d8b72b51a59fdb "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations/item"
    ic9cf4eefa0a637dd860228da8f11a3a1b0e6c19dcc2c41b740a8ff4b6ce24681 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mdmwindowsinformationprotectionpolicies/item"
    idbad9dd444740f91fdc2deb224c3442bfd8ca9dbec92cd68d18bc8830f6cc2bf "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/vpptokens/item"
    iddcc8f3790be69a76d73a881b2a5249b825b8ed1e8dcd570bba76b931dfeb702 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks/item"
    iea05d6821455c74468d20c06a0c7e47fe33a08b96f01d52bce3d6d5aebb6e52d "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappstatuses/item"
)

// Builds and executes requests for operations under \deviceAppManagement
type DeviceAppManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// Options for Get
type DeviceAppManagementRequestBuilderGetOptions struct {
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Request query parameters
    Q *DeviceAppManagementRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
// Get deviceAppManagement
type DeviceAppManagementRequestBuilderGetQueryParameters struct {
    ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.QueryParametersBase
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select_escaped []string;
}
// Options for Patch
type DeviceAppManagementRequestBuilderPatchOptions struct {
    // 
    Body *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceAppManagement;
    // Request headers
    H map[string]string;
    // Request options
    O []ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.ResponseHandler;
}
func (m *DeviceAppManagementRequestBuilder) AndroidManagedAppProtections()(*i2adc870208a84a1c715e640a35a322f60d13c7132faecab92cbdc856af0e50aa.AndroidManagedAppProtectionsRequestBuilder) {
    return i2adc870208a84a1c715e640a35a322f60d13c7132faecab92cbdc856af0e50aa.NewAndroidManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.androidManagedAppProtections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) AndroidManagedAppProtectionsById(id string)(*i89d4341c9169ba82f4022fde15796f9a0594e04c5cc3dab59767eaa7b1fdf35c.AndroidManagedAppProtectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["androidManagedAppProtection_id"] = id
    }
    return i89d4341c9169ba82f4022fde15796f9a0594e04c5cc3dab59767eaa7b1fdf35c.NewAndroidManagedAppProtectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Instantiates a new DeviceAppManagementRequestBuilder and sets the default values.
// Parameters:
//  - pathParameters : Path parameters for the request
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDeviceAppManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceAppManagementRequestBuilder) {
    m := &DeviceAppManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = pathParameters;
    m.requestAdapter = requestAdapter;
    return m
}
// Instantiates a new DeviceAppManagementRequestBuilder and sets the default values.
// Parameters:
//  - rawUrl : The raw URL to use for the request builder.
//  - requestAdapter : The request adapter to use to execute the requests.
func NewDeviceAppManagementRequestBuilder(rawUrl string, requestAdapter ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestAdapter)(*DeviceAppManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// Get deviceAppManagement
// Parameters:
//  - options : Options for the request
func (m *DeviceAppManagementRequestBuilder) CreateGetRequestInformation(options *DeviceAppManagementRequestBuilderGetOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
    requestInfo := ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.GET
    if options != nil && options.Q != nil {
        err := options.Q.AddQueryParameters(requestInfo.QueryParameters)
        if err != nil {
            return nil, err
        }
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
// Update deviceAppManagement
// Parameters:
//  - options : Options for the request
func (m *DeviceAppManagementRequestBuilder) CreatePatchRequestInformation(options *DeviceAppManagementRequestBuilderPatchOptions)(*ida96af0f171bb75f894a4013a6b3146a4397c58f11adb81a2b7cbea9314783a9.RequestInformation, error) {
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
func (m *DeviceAppManagementRequestBuilder) DefaultManagedAppProtections()(*i18bf226bfeead5544b159f6b6a278878b7180022db16e17df961b8a912ac109f.DefaultManagedAppProtectionsRequestBuilder) {
    return i18bf226bfeead5544b159f6b6a278878b7180022db16e17df961b8a912ac109f.NewDefaultManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.defaultManagedAppProtections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) DefaultManagedAppProtectionsById(id string)(*ia3c7a3f771ad6210672d21e0ac5e24781d4c319a31e3c1088a61103e6a6c19b8.DefaultManagedAppProtectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["defaultManagedAppProtection_id"] = id
    }
    return ia3c7a3f771ad6210672d21e0ac5e24781d4c319a31e3c1088a61103e6a6c19b8.NewDefaultManagedAppProtectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get deviceAppManagement
// Parameters:
//  - options : Options for the request
func (m *DeviceAppManagementRequestBuilder) Get(options *DeviceAppManagementRequestBuilderGetOptions)(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceAppManagement, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    res, err := m.requestAdapter.SendAsync(*requestInfo, func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewDeviceAppManagement() }, nil)
    if err != nil {
        return nil, err
    }
    return res.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.DeviceAppManagement), nil
}
func (m *DeviceAppManagementRequestBuilder) IosManagedAppProtections()(*i32739bc83be1b44d9a0e5bef6898e52a1d08cff6911a2ca8739fb5aaf580c634.IosManagedAppProtectionsRequestBuilder) {
    return i32739bc83be1b44d9a0e5bef6898e52a1d08cff6911a2ca8739fb5aaf580c634.NewIosManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.iosManagedAppProtections.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) IosManagedAppProtectionsById(id string)(*ib2f42e9c4d8f657233687ebcb7ea2a7be72112a45159e6921784148695f33cc8.IosManagedAppProtectionRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosManagedAppProtection_id"] = id
    }
    return ib2f42e9c4d8f657233687ebcb7ea2a7be72112a45159e6921784148695f33cc8.NewIosManagedAppProtectionRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceAppManagementRequestBuilder) ManagedAppPolicies()(*i03d64ffdecc8f8ee64e8bdbbd21cd792ea749ccc7d6632ed57a9368c023aca62.ManagedAppPoliciesRequestBuilder) {
    return i03d64ffdecc8f8ee64e8bdbbd21cd792ea749ccc7d6632ed57a9368c023aca62.NewManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedAppPolicies.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) ManagedAppPoliciesById(id string)(*i560cf4a9f430034859b2e6c353602d2fd9280316fd3f2f1265995ae7216d4dc3.ManagedAppPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy_id"] = id
    }
    return i560cf4a9f430034859b2e6c353602d2fd9280316fd3f2f1265995ae7216d4dc3.NewManagedAppPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceAppManagementRequestBuilder) ManagedAppRegistrations()(*i489c24d5b60b7229ec09f24405ed28097801dccb5fc9e368e362132b825c2808.ManagedAppRegistrationsRequestBuilder) {
    return i489c24d5b60b7229ec09f24405ed28097801dccb5fc9e368e362132b825c2808.NewManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedAppRegistrations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) ManagedAppRegistrationsById(id string)(*i090e5453c2fdf074fa2d7cb168a2f22ffe8f16816d015832f6cb242b663410b3.ManagedAppRegistrationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppRegistration_id"] = id
    }
    return i090e5453c2fdf074fa2d7cb168a2f22ffe8f16816d015832f6cb242b663410b3.NewManagedAppRegistrationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceAppManagementRequestBuilder) ManagedAppStatuses()(*ie5dc06c3359896dbdaa15cf264f2382a825efa77f089fedc5fd0e259ac122847.ManagedAppStatusesRequestBuilder) {
    return ie5dc06c3359896dbdaa15cf264f2382a825efa77f089fedc5fd0e259ac122847.NewManagedAppStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedAppStatuses.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) ManagedAppStatusesById(id string)(*iea05d6821455c74468d20c06a0c7e47fe33a08b96f01d52bce3d6d5aebb6e52d.ManagedAppStatusRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppStatus_id"] = id
    }
    return iea05d6821455c74468d20c06a0c7e47fe33a08b96f01d52bce3d6d5aebb6e52d.NewManagedAppStatusRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceAppManagementRequestBuilder) ManagedEBooks()(*i26923530adcc2c194b34fa3104b74fd0fd1849f0fcb7b138a8585bde57bca4d6.ManagedEBooksRequestBuilder) {
    return i26923530adcc2c194b34fa3104b74fd0fd1849f0fcb7b138a8585bde57bca4d6.NewManagedEBooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedEBooks.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) ManagedEBooksById(id string)(*iddcc8f3790be69a76d73a881b2a5249b825b8ed1e8dcd570bba76b931dfeb702.ManagedEBookRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBook_id"] = id
    }
    return iddcc8f3790be69a76d73a881b2a5249b825b8ed1e8dcd570bba76b931dfeb702.NewManagedEBookRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceAppManagementRequestBuilder) MdmWindowsInformationProtectionPolicies()(*if1c061f5229e920421d3fbb9fc720eab490b00af016a531ace0cd07293f04818.MdmWindowsInformationProtectionPoliciesRequestBuilder) {
    return if1c061f5229e920421d3fbb9fc720eab490b00af016a531ace0cd07293f04818.NewMdmWindowsInformationProtectionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mdmWindowsInformationProtectionPolicies.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) MdmWindowsInformationProtectionPoliciesById(id string)(*ic9cf4eefa0a637dd860228da8f11a3a1b0e6c19dcc2c41b740a8ff4b6ce24681.MdmWindowsInformationProtectionPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mdmWindowsInformationProtectionPolicy_id"] = id
    }
    return ic9cf4eefa0a637dd860228da8f11a3a1b0e6c19dcc2c41b740a8ff4b6ce24681.NewMdmWindowsInformationProtectionPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceAppManagementRequestBuilder) MobileAppCategories()(*i27d43f0c2c303b4996cf211399e808bbe9d9cd0c447c2d98061c3e3157225ee2.MobileAppCategoriesRequestBuilder) {
    return i27d43f0c2c303b4996cf211399e808bbe9d9cd0c447c2d98061c3e3157225ee2.NewMobileAppCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileAppCategories.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) MobileAppCategoriesById(id string)(*i0f78fb508810d8046e0126c99b4e2f29dfbc726d712f29fb4c961a92e8630256.MobileAppCategoryRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppCategory_id"] = id
    }
    return i0f78fb508810d8046e0126c99b4e2f29dfbc726d712f29fb4c961a92e8630256.NewMobileAppCategoryRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceAppManagementRequestBuilder) MobileAppConfigurations()(*id8e9dee048ce49d5be720313759ba7ff0e5561292831c5866a1251d1b6f75017.MobileAppConfigurationsRequestBuilder) {
    return id8e9dee048ce49d5be720313759ba7ff0e5561292831c5866a1251d1b6f75017.NewMobileAppConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileAppConfigurations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) MobileAppConfigurationsById(id string)(*ib471248babfa75d07aa6e6560ccb2ae999e37717a38cabb105d8b72b51a59fdb.ManagedDeviceMobileAppConfigurationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfiguration_id"] = id
    }
    return ib471248babfa75d07aa6e6560ccb2ae999e37717a38cabb105d8b72b51a59fdb.NewManagedDeviceMobileAppConfigurationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceAppManagementRequestBuilder) MobileApps()(*id47f41c3fc661963eced1051494f915ce9e3fc4f22d73157e3a706e3abf7c3c0.MobileAppsRequestBuilder) {
    return id47f41c3fc661963eced1051494f915ce9e3fc4f22d73157e3a706e3abf7c3c0.NewMobileAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileApps.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) MobileAppsById(id string)(*i4d54ef094838af59d772f617ee78a54ffbb57f062a912b927b484571e3b93ca5.MobileAppRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileApp_id"] = id
    }
    return i4d54ef094838af59d772f617ee78a54ffbb57f062a912b927b484571e3b93ca5.NewMobileAppRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Update deviceAppManagement
// Parameters:
//  - options : Options for the request
func (m *DeviceAppManagementRequestBuilder) Patch(options *DeviceAppManagementRequestBuilderPatchOptions)(error) {
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
func (m *DeviceAppManagementRequestBuilder) SyncMicrosoftStoreForBusinessApps()(*i6ca5a8c664cfb8f06d3237f44f8e2ec59140505f80b90c4cb4a1af3a15276d77.SyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    return i6ca5a8c664cfb8f06d3237f44f8e2ec59140505f80b90c4cb4a1af3a15276d77.NewSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
func (m *DeviceAppManagementRequestBuilder) TargetedManagedAppConfigurations()(*ia84055a082d525b002332712c0df79f5545cff95fe4a4ba7647a1d711273c026.TargetedManagedAppConfigurationsRequestBuilder) {
    return ia84055a082d525b002332712c0df79f5545cff95fe4a4ba7647a1d711273c026.NewTargetedManagedAppConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.targetedManagedAppConfigurations.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) TargetedManagedAppConfigurationsById(id string)(*i0076e81bdc67eb89b0f75e1d90ce618e38e43bf24fdd9abc2d2dd7b7b418df1c.TargetedManagedAppConfigurationRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["targetedManagedAppConfiguration_id"] = id
    }
    return i0076e81bdc67eb89b0f75e1d90ce618e38e43bf24fdd9abc2d2dd7b7b418df1c.NewTargetedManagedAppConfigurationRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceAppManagementRequestBuilder) VppTokens()(*iacf1c27653093eb9d818263eb22e73f2d1d6e2a5e231803fa86fe531b6e3f037.VppTokensRequestBuilder) {
    return iacf1c27653093eb9d818263eb22e73f2d1d6e2a5e231803fa86fe531b6e3f037.NewVppTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.vppTokens.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) VppTokensById(id string)(*idbad9dd444740f91fdc2deb224c3442bfd8ca9dbec92cd68d18bc8830f6cc2bf.VppTokenRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["vppToken_id"] = id
    }
    return idbad9dd444740f91fdc2deb224c3442bfd8ca9dbec92cd68d18bc8830f6cc2bf.NewVppTokenRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionPolicies()(*i3c79f5992014bafe598fd20b6e9fd4b5f377e28ceccdbceb4002f89b8068136b.WindowsInformationProtectionPoliciesRequestBuilder) {
    return i3c79f5992014bafe598fd20b6e9fd4b5f377e28ceccdbceb4002f89b8068136b.NewWindowsInformationProtectionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// Gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.windowsInformationProtectionPolicies.item collection
// Parameters:
//  - id : Unique identifier of the item
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionPoliciesById(id string)(*i3ab3bf13acf28406514cba3b07ab5235a9c5c97ff398576b733a15152cfb62d3.WindowsInformationProtectionPolicyRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionPolicy_id"] = id
    }
    return i3ab3bf13acf28406514cba3b07ab5235a9c5c97ff398576b733a15152cfb62d3.NewWindowsInformationProtectionPolicyRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
