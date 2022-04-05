package deviceappmanagement

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
    i03d64ffdecc8f8ee64e8bdbbd21cd792ea749ccc7d6632ed57a9368c023aca62 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedapppolicies"
    i18bf226bfeead5544b159f6b6a278878b7180022db16e17df961b8a912ac109f "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/defaultmanagedappprotections"
    i26923530adcc2c194b34fa3104b74fd0fd1849f0fcb7b138a8585bde57bca4d6 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedebooks"
    i27d43f0c2c303b4996cf211399e808bbe9d9cd0c447c2d98061c3e3157225ee2 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappcategories"
    i2adc870208a84a1c715e640a35a322f60d13c7132faecab92cbdc856af0e50aa "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/androidmanagedappprotections"
    i32739bc83be1b44d9a0e5bef6898e52a1d08cff6911a2ca8739fb5aaf580c634 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/iosmanagedappprotections"
    i3c79f5992014bafe598fd20b6e9fd4b5f377e28ceccdbceb4002f89b8068136b "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/windowsinformationprotectionpolicies"
    i489c24d5b60b7229ec09f24405ed28097801dccb5fc9e368e362132b825c2808 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations"
    i6ca5a8c664cfb8f06d3237f44f8e2ec59140505f80b90c4cb4a1af3a15276d77 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/syncmicrosoftstoreforbusinessapps"
    ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a "github.com/microsoftgraph/msgraph-sdk-go/models/odataerrors"
    ia84055a082d525b002332712c0df79f5545cff95fe4a4ba7647a1d711273c026 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations"
    iacf1c27653093eb9d818263eb22e73f2d1d6e2a5e231803fa86fe531b6e3f037 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/vpptokens"
    id47f41c3fc661963eced1051494f915ce9e3fc4f22d73157e3a706e3abf7c3c0 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileapps"
    id8e9dee048ce49d5be720313759ba7ff0e5561292831c5866a1251d1b6f75017 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappconfigurations"
    ie5dc06c3359896dbdaa15cf264f2382a825efa77f089fedc5fd0e259ac122847 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappstatuses"
    if1c061f5229e920421d3fbb9fc720eab490b00af016a531ace0cd07293f04818 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mdmwindowsinformationprotectionpolicies"
    i0076e81bdc67eb89b0f75e1d90ce618e38e43bf24fdd9abc2d2dd7b7b418df1c "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/targetedmanagedappconfigurations/item"
    i090e5453c2fdf074fa2d7cb168a2f22ffe8f16816d015832f6cb242b663410b3 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/managedappregistrations/item"
    i0f78fb508810d8046e0126c99b4e2f29dfbc726d712f29fb4c961a92e8630256 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/mobileappcategories/item"
    i3ab3bf13acf28406514cba3b07ab5235a9c5c97ff398576b733a15152cfb62d3 "github.com/microsoftgraph/msgraph-sdk-go/deviceappmanagement/windowsinformationprotectionpolicies/item"
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

// DeviceAppManagementRequestBuilder provides operations to manage the deviceAppManagement singleton.
type DeviceAppManagementRequestBuilder struct {
    // Path parameters for the request
    pathParameters map[string]string;
    // The request adapter to use to execute the requests.
    requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter;
    // Url template to use to build the URL for the current request builder
    urlTemplate string;
}
// DeviceAppManagementRequestBuilderGetOptions options for Get
type DeviceAppManagementRequestBuilderGetOptions struct {
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Request query parameters
    QueryParameters *DeviceAppManagementRequestBuilderGetQueryParameters;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// DeviceAppManagementRequestBuilderGetQueryParameters get deviceAppManagement
type DeviceAppManagementRequestBuilderGetQueryParameters struct {
    // Expand related entities
    Expand []string;
    // Select properties to be returned
    Select []string;
}
// DeviceAppManagementRequestBuilderPatchOptions options for Patch
type DeviceAppManagementRequestBuilderPatchOptions struct {
    // 
    Body iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceAppManagementable;
    // Request headers
    Headers map[string]string;
    // Request options
    Options []i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestOption;
    // Response handler to use in place of the default response handling provided by the core service
    ResponseHandler i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ResponseHandler;
}
// AndroidManagedAppProtections the androidManagedAppProtections property
func (m *DeviceAppManagementRequestBuilder) AndroidManagedAppProtections()(*i2adc870208a84a1c715e640a35a322f60d13c7132faecab92cbdc856af0e50aa.AndroidManagedAppProtectionsRequestBuilder) {
    return i2adc870208a84a1c715e640a35a322f60d13c7132faecab92cbdc856af0e50aa.NewAndroidManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// AndroidManagedAppProtectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.androidManagedAppProtections.item collection
func (m *DeviceAppManagementRequestBuilder) AndroidManagedAppProtectionsById(id string)(*i89d4341c9169ba82f4022fde15796f9a0594e04c5cc3dab59767eaa7b1fdf35c.AndroidManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["androidManagedAppProtection_id"] = id
    }
    return i89d4341c9169ba82f4022fde15796f9a0594e04c5cc3dab59767eaa7b1fdf35c.NewAndroidManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// NewDeviceAppManagementRequestBuilderInternal instantiates a new DeviceAppManagementRequestBuilder and sets the default values.
func NewDeviceAppManagementRequestBuilderInternal(pathParameters map[string]string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementRequestBuilder) {
    m := &DeviceAppManagementRequestBuilder{
    }
    m.urlTemplate = "{+baseurl}/deviceAppManagement{?select,expand}";
    urlTplParams := make(map[string]string)
    for idx, item := range pathParameters {
        urlTplParams[idx] = item
    }
    m.pathParameters = urlTplParams;
    m.requestAdapter = requestAdapter;
    return m
}
// NewDeviceAppManagementRequestBuilder instantiates a new DeviceAppManagementRequestBuilder and sets the default values.
func NewDeviceAppManagementRequestBuilder(rawUrl string, requestAdapter i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestAdapter)(*DeviceAppManagementRequestBuilder) {
    urlParams := make(map[string]string)
    urlParams["request-raw-url"] = rawUrl
    return NewDeviceAppManagementRequestBuilderInternal(urlParams, requestAdapter)
}
// CreateGetRequestInformation get deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) CreateGetRequestInformation(options *DeviceAppManagementRequestBuilderGetOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.GET
    if options != nil && options.QueryParameters != nil {
        requestInfo.AddQueryParameters(*(options.QueryParameters))
    }
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// CreatePatchRequestInformation update deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) CreatePatchRequestInformation(options *DeviceAppManagementRequestBuilderPatchOptions)(*i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.RequestInformation, error) {
    requestInfo := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.NewRequestInformation()
    requestInfo.UrlTemplate = m.urlTemplate
    requestInfo.PathParameters = m.pathParameters
    requestInfo.Method = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.PATCH
    requestInfo.SetContentFromParsable(m.requestAdapter, "application/json", options.Body)
    if options != nil && options.Headers != nil {
        requestInfo.Headers = options.Headers
    }
    if options != nil && len(options.Options) != 0 {
        err := requestInfo.AddRequestOptions(options.Options...)
        if err != nil {
            return nil, err
        }
    }
    return requestInfo, nil
}
// DefaultManagedAppProtections the defaultManagedAppProtections property
func (m *DeviceAppManagementRequestBuilder) DefaultManagedAppProtections()(*i18bf226bfeead5544b159f6b6a278878b7180022db16e17df961b8a912ac109f.DefaultManagedAppProtectionsRequestBuilder) {
    return i18bf226bfeead5544b159f6b6a278878b7180022db16e17df961b8a912ac109f.NewDefaultManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// DefaultManagedAppProtectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.defaultManagedAppProtections.item collection
func (m *DeviceAppManagementRequestBuilder) DefaultManagedAppProtectionsById(id string)(*ia3c7a3f771ad6210672d21e0ac5e24781d4c319a31e3c1088a61103e6a6c19b8.DefaultManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["defaultManagedAppProtection_id"] = id
    }
    return ia3c7a3f771ad6210672d21e0ac5e24781d4c319a31e3c1088a61103e6a6c19b8.NewDefaultManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Get get deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) Get(options *DeviceAppManagementRequestBuilderGetOptions)(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceAppManagementable, error) {
    requestInfo, err := m.CreateGetRequestInformation(options);
    if err != nil {
        return nil, err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    res, err := m.requestAdapter.SendAsync(requestInfo, iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateDeviceAppManagementFromDiscriminatorValue, nil, errorMapping)
    if err != nil {
        return nil, err
    }
    return res.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.DeviceAppManagementable), nil
}
// IosManagedAppProtections the iosManagedAppProtections property
func (m *DeviceAppManagementRequestBuilder) IosManagedAppProtections()(*i32739bc83be1b44d9a0e5bef6898e52a1d08cff6911a2ca8739fb5aaf580c634.IosManagedAppProtectionsRequestBuilder) {
    return i32739bc83be1b44d9a0e5bef6898e52a1d08cff6911a2ca8739fb5aaf580c634.NewIosManagedAppProtectionsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// IosManagedAppProtectionsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.iosManagedAppProtections.item collection
func (m *DeviceAppManagementRequestBuilder) IosManagedAppProtectionsById(id string)(*ib2f42e9c4d8f657233687ebcb7ea2a7be72112a45159e6921784148695f33cc8.IosManagedAppProtectionItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["iosManagedAppProtection_id"] = id
    }
    return ib2f42e9c4d8f657233687ebcb7ea2a7be72112a45159e6921784148695f33cc8.NewIosManagedAppProtectionItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppPolicies the managedAppPolicies property
func (m *DeviceAppManagementRequestBuilder) ManagedAppPolicies()(*i03d64ffdecc8f8ee64e8bdbbd21cd792ea749ccc7d6632ed57a9368c023aca62.ManagedAppPoliciesRequestBuilder) {
    return i03d64ffdecc8f8ee64e8bdbbd21cd792ea749ccc7d6632ed57a9368c023aca62.NewManagedAppPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedAppPolicies.item collection
func (m *DeviceAppManagementRequestBuilder) ManagedAppPoliciesById(id string)(*i560cf4a9f430034859b2e6c353602d2fd9280316fd3f2f1265995ae7216d4dc3.ManagedAppPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppPolicy_id"] = id
    }
    return i560cf4a9f430034859b2e6c353602d2fd9280316fd3f2f1265995ae7216d4dc3.NewManagedAppPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppRegistrations the managedAppRegistrations property
func (m *DeviceAppManagementRequestBuilder) ManagedAppRegistrations()(*i489c24d5b60b7229ec09f24405ed28097801dccb5fc9e368e362132b825c2808.ManagedAppRegistrationsRequestBuilder) {
    return i489c24d5b60b7229ec09f24405ed28097801dccb5fc9e368e362132b825c2808.NewManagedAppRegistrationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppRegistrationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedAppRegistrations.item collection
func (m *DeviceAppManagementRequestBuilder) ManagedAppRegistrationsById(id string)(*i090e5453c2fdf074fa2d7cb168a2f22ffe8f16816d015832f6cb242b663410b3.ManagedAppRegistrationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppRegistration_id"] = id
    }
    return i090e5453c2fdf074fa2d7cb168a2f22ffe8f16816d015832f6cb242b663410b3.NewManagedAppRegistrationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedAppStatuses the managedAppStatuses property
func (m *DeviceAppManagementRequestBuilder) ManagedAppStatuses()(*ie5dc06c3359896dbdaa15cf264f2382a825efa77f089fedc5fd0e259ac122847.ManagedAppStatusesRequestBuilder) {
    return ie5dc06c3359896dbdaa15cf264f2382a825efa77f089fedc5fd0e259ac122847.NewManagedAppStatusesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedAppStatusesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedAppStatuses.item collection
func (m *DeviceAppManagementRequestBuilder) ManagedAppStatusesById(id string)(*iea05d6821455c74468d20c06a0c7e47fe33a08b96f01d52bce3d6d5aebb6e52d.ManagedAppStatusItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedAppStatus_id"] = id
    }
    return iea05d6821455c74468d20c06a0c7e47fe33a08b96f01d52bce3d6d5aebb6e52d.NewManagedAppStatusItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// ManagedEBooks the managedEBooks property
func (m *DeviceAppManagementRequestBuilder) ManagedEBooks()(*i26923530adcc2c194b34fa3104b74fd0fd1849f0fcb7b138a8585bde57bca4d6.ManagedEBooksRequestBuilder) {
    return i26923530adcc2c194b34fa3104b74fd0fd1849f0fcb7b138a8585bde57bca4d6.NewManagedEBooksRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// ManagedEBooksById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.managedEBooks.item collection
func (m *DeviceAppManagementRequestBuilder) ManagedEBooksById(id string)(*iddcc8f3790be69a76d73a881b2a5249b825b8ed1e8dcd570bba76b931dfeb702.ManagedEBookItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedEBook_id"] = id
    }
    return iddcc8f3790be69a76d73a881b2a5249b825b8ed1e8dcd570bba76b931dfeb702.NewManagedEBookItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MdmWindowsInformationProtectionPolicies the mdmWindowsInformationProtectionPolicies property
func (m *DeviceAppManagementRequestBuilder) MdmWindowsInformationProtectionPolicies()(*if1c061f5229e920421d3fbb9fc720eab490b00af016a531ace0cd07293f04818.MdmWindowsInformationProtectionPoliciesRequestBuilder) {
    return if1c061f5229e920421d3fbb9fc720eab490b00af016a531ace0cd07293f04818.NewMdmWindowsInformationProtectionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MdmWindowsInformationProtectionPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mdmWindowsInformationProtectionPolicies.item collection
func (m *DeviceAppManagementRequestBuilder) MdmWindowsInformationProtectionPoliciesById(id string)(*ic9cf4eefa0a637dd860228da8f11a3a1b0e6c19dcc2c41b740a8ff4b6ce24681.MdmWindowsInformationProtectionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mdmWindowsInformationProtectionPolicy_id"] = id
    }
    return ic9cf4eefa0a637dd860228da8f11a3a1b0e6c19dcc2c41b740a8ff4b6ce24681.NewMdmWindowsInformationProtectionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppCategories the mobileAppCategories property
func (m *DeviceAppManagementRequestBuilder) MobileAppCategories()(*i27d43f0c2c303b4996cf211399e808bbe9d9cd0c447c2d98061c3e3157225ee2.MobileAppCategoriesRequestBuilder) {
    return i27d43f0c2c303b4996cf211399e808bbe9d9cd0c447c2d98061c3e3157225ee2.NewMobileAppCategoriesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppCategoriesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileAppCategories.item collection
func (m *DeviceAppManagementRequestBuilder) MobileAppCategoriesById(id string)(*i0f78fb508810d8046e0126c99b4e2f29dfbc726d712f29fb4c961a92e8630256.MobileAppCategoryItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileAppCategory_id"] = id
    }
    return i0f78fb508810d8046e0126c99b4e2f29dfbc726d712f29fb4c961a92e8630256.NewMobileAppCategoryItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileAppConfigurations the mobileAppConfigurations property
func (m *DeviceAppManagementRequestBuilder) MobileAppConfigurations()(*id8e9dee048ce49d5be720313759ba7ff0e5561292831c5866a1251d1b6f75017.MobileAppConfigurationsRequestBuilder) {
    return id8e9dee048ce49d5be720313759ba7ff0e5561292831c5866a1251d1b6f75017.NewMobileAppConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileAppConfigurations.item collection
func (m *DeviceAppManagementRequestBuilder) MobileAppConfigurationsById(id string)(*ib471248babfa75d07aa6e6560ccb2ae999e37717a38cabb105d8b72b51a59fdb.ManagedDeviceMobileAppConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["managedDeviceMobileAppConfiguration_id"] = id
    }
    return ib471248babfa75d07aa6e6560ccb2ae999e37717a38cabb105d8b72b51a59fdb.NewManagedDeviceMobileAppConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// MobileApps the mobileApps property
func (m *DeviceAppManagementRequestBuilder) MobileApps()(*id47f41c3fc661963eced1051494f915ce9e3fc4f22d73157e3a706e3abf7c3c0.MobileAppsRequestBuilder) {
    return id47f41c3fc661963eced1051494f915ce9e3fc4f22d73157e3a706e3abf7c3c0.NewMobileAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// MobileAppsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.mobileApps.item collection
func (m *DeviceAppManagementRequestBuilder) MobileAppsById(id string)(*i4d54ef094838af59d772f617ee78a54ffbb57f062a912b927b484571e3b93ca5.MobileAppItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["mobileApp_id"] = id
    }
    return i4d54ef094838af59d772f617ee78a54ffbb57f062a912b927b484571e3b93ca5.NewMobileAppItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// Patch update deviceAppManagement
func (m *DeviceAppManagementRequestBuilder) Patch(options *DeviceAppManagementRequestBuilderPatchOptions)(error) {
    requestInfo, err := m.CreatePatchRequestInformation(options);
    if err != nil {
        return err
    }
    errorMapping := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.ErrorMappings {
        "4XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
        "5XX": ia572726a95efa92ddd544552cd950653dc691023836923576b2f4bf716cf204a.CreateODataErrorFromDiscriminatorValue,
    }
    err = m.requestAdapter.SendNoContentAsync(requestInfo, nil, errorMapping)
    if err != nil {
        return err
    }
    return nil
}
// SyncMicrosoftStoreForBusinessApps the syncMicrosoftStoreForBusinessApps property
func (m *DeviceAppManagementRequestBuilder) SyncMicrosoftStoreForBusinessApps()(*i6ca5a8c664cfb8f06d3237f44f8e2ec59140505f80b90c4cb4a1af3a15276d77.SyncMicrosoftStoreForBusinessAppsRequestBuilder) {
    return i6ca5a8c664cfb8f06d3237f44f8e2ec59140505f80b90c4cb4a1af3a15276d77.NewSyncMicrosoftStoreForBusinessAppsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TargetedManagedAppConfigurations the targetedManagedAppConfigurations property
func (m *DeviceAppManagementRequestBuilder) TargetedManagedAppConfigurations()(*ia84055a082d525b002332712c0df79f5545cff95fe4a4ba7647a1d711273c026.TargetedManagedAppConfigurationsRequestBuilder) {
    return ia84055a082d525b002332712c0df79f5545cff95fe4a4ba7647a1d711273c026.NewTargetedManagedAppConfigurationsRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// TargetedManagedAppConfigurationsById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.targetedManagedAppConfigurations.item collection
func (m *DeviceAppManagementRequestBuilder) TargetedManagedAppConfigurationsById(id string)(*i0076e81bdc67eb89b0f75e1d90ce618e38e43bf24fdd9abc2d2dd7b7b418df1c.TargetedManagedAppConfigurationItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["targetedManagedAppConfiguration_id"] = id
    }
    return i0076e81bdc67eb89b0f75e1d90ce618e38e43bf24fdd9abc2d2dd7b7b418df1c.NewTargetedManagedAppConfigurationItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// VppTokens the vppTokens property
func (m *DeviceAppManagementRequestBuilder) VppTokens()(*iacf1c27653093eb9d818263eb22e73f2d1d6e2a5e231803fa86fe531b6e3f037.VppTokensRequestBuilder) {
    return iacf1c27653093eb9d818263eb22e73f2d1d6e2a5e231803fa86fe531b6e3f037.NewVppTokensRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// VppTokensById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.vppTokens.item collection
func (m *DeviceAppManagementRequestBuilder) VppTokensById(id string)(*idbad9dd444740f91fdc2deb224c3442bfd8ca9dbec92cd68d18bc8830f6cc2bf.VppTokenItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["vppToken_id"] = id
    }
    return idbad9dd444740f91fdc2deb224c3442bfd8ca9dbec92cd68d18bc8830f6cc2bf.NewVppTokenItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
// WindowsInformationProtectionPolicies the windowsInformationProtectionPolicies property
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionPolicies()(*i3c79f5992014bafe598fd20b6e9fd4b5f377e28ceccdbceb4002f89b8068136b.WindowsInformationProtectionPoliciesRequestBuilder) {
    return i3c79f5992014bafe598fd20b6e9fd4b5f377e28ceccdbceb4002f89b8068136b.NewWindowsInformationProtectionPoliciesRequestBuilderInternal(m.pathParameters, m.requestAdapter);
}
// WindowsInformationProtectionPoliciesById gets an item from the github.com/microsoftgraph/msgraph-sdk-go/.deviceAppManagement.windowsInformationProtectionPolicies.item collection
func (m *DeviceAppManagementRequestBuilder) WindowsInformationProtectionPoliciesById(id string)(*i3ab3bf13acf28406514cba3b07ab5235a9c5c97ff398576b733a15152cfb62d3.WindowsInformationProtectionPolicyItemRequestBuilder) {
    urlTplParams := make(map[string]string)
    for idx, item := range m.pathParameters {
        urlTplParams[idx] = item
    }
    if id != "" {
        urlTplParams["windowsInformationProtectionPolicy_id"] = id
    }
    return i3ab3bf13acf28406514cba3b07ab5235a9c5c97ff398576b733a15152cfb62d3.NewWindowsInformationProtectionPolicyItemRequestBuilderInternal(urlTplParams, m.requestAdapter);
}
