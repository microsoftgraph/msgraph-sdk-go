package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ManagedAppProtection struct {
    ManagedAppPolicy
    // Data storage locations where a user may store managed data.
    allowedDataStorageLocations []ManagedAppDataStorageLocation;
    // Sources from which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
    allowedInboundDataTransferSources *ManagedAppDataTransferLevel;
    // The level to which the clipboard may be shared between apps on the managed device. Possible values are: allApps, managedAppsWithPasteIn, managedApps, blocked.
    allowedOutboundClipboardSharingLevel *ManagedAppClipboardSharingLevel;
    // Destinations to which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
    allowedOutboundDataTransferDestinations *ManagedAppDataTransferLevel;
    // Indicates whether contacts can be synced to the user's device.
    contactSyncBlocked *bool;
    // Indicates whether the backup of a managed app's data is blocked.
    dataBackupBlocked *bool;
    // Indicates whether device compliance is required.
    deviceComplianceRequired *bool;
    // Indicates whether use of the app pin is required if the device pin is set.
    disableAppPinIfDevicePinIsSet *bool;
    // Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
    fingerprintBlocked *bool;
    // Indicates in which managed browser(s) that internet links should be opened. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. Possible values are: notConfigured, microsoftEdge.
    managedBrowser *ManagedBrowserType;
    // Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
    managedBrowserToOpenLinksRequired *bool;
    // Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
    maximumPinRetries *int32;
    // Minimum pin length required for an app-level pin if PinRequired is set to True
    minimumPinLength *int32;
    // Versions less than the specified version will block the managed app from accessing company data.
    minimumRequiredAppVersion *string;
    // Versions less than the specified version will block the managed app from accessing company data.
    minimumRequiredOsVersion *string;
    // Versions less than the specified version will result in warning message on the managed app.
    minimumWarningAppVersion *string;
    // Versions less than the specified version will result in warning message on the managed app from accessing company data.
    minimumWarningOsVersion *string;
    // Indicates whether organizational credentials are required for app use.
    organizationalCredentialsRequired *bool;
    // TimePeriod before the all-level pin must be reset if PinRequired is set to True.
    periodBeforePinReset *string;
    // The period after which access is checked when the device is not connected to the internet.
    periodOfflineBeforeAccessCheck *string;
    // The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
    periodOfflineBeforeWipeIsEnforced *string;
    // The period after which access is checked when the device is connected to the internet.
    periodOnlineBeforeAccessCheck *string;
    // Character set which may be used for an app-level pin if PinRequired is set to True. Possible values are: numeric, alphanumericAndSymbol.
    pinCharacterSet *ManagedAppPinCharacterSet;
    // Indicates whether an app-level pin is required.
    pinRequired *bool;
    // Indicates whether printing is allowed from managed apps.
    printBlocked *bool;
    // Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
    saveAsBlocked *bool;
    // Indicates whether simplePin is blocked.
    simplePinBlocked *bool;
}
// Instantiates a new managedAppProtection and sets the default values.
func NewManagedAppProtection()(*ManagedAppProtection) {
    m := &ManagedAppProtection{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    return m
}
// Gets the allowedDataStorageLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) GetAllowedDataStorageLocations()([]ManagedAppDataStorageLocation) {
    if m == nil {
        return nil
    } else {
        return m.allowedDataStorageLocations
    }
}
// Gets the allowedInboundDataTransferSources property value. Sources from which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
func (m *ManagedAppProtection) GetAllowedInboundDataTransferSources()(*ManagedAppDataTransferLevel) {
    if m == nil {
        return nil
    } else {
        return m.allowedInboundDataTransferSources
    }
}
// Gets the allowedOutboundClipboardSharingLevel property value. The level to which the clipboard may be shared between apps on the managed device. Possible values are: allApps, managedAppsWithPasteIn, managedApps, blocked.
func (m *ManagedAppProtection) GetAllowedOutboundClipboardSharingLevel()(*ManagedAppClipboardSharingLevel) {
    if m == nil {
        return nil
    } else {
        return m.allowedOutboundClipboardSharingLevel
    }
}
// Gets the allowedOutboundDataTransferDestinations property value. Destinations to which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
func (m *ManagedAppProtection) GetAllowedOutboundDataTransferDestinations()(*ManagedAppDataTransferLevel) {
    if m == nil {
        return nil
    } else {
        return m.allowedOutboundDataTransferDestinations
    }
}
// Gets the contactSyncBlocked property value. Indicates whether contacts can be synced to the user's device.
func (m *ManagedAppProtection) GetContactSyncBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contactSyncBlocked
    }
}
// Gets the dataBackupBlocked property value. Indicates whether the backup of a managed app's data is blocked.
func (m *ManagedAppProtection) GetDataBackupBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.dataBackupBlocked
    }
}
// Gets the deviceComplianceRequired property value. Indicates whether device compliance is required.
func (m *ManagedAppProtection) GetDeviceComplianceRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceRequired
    }
}
// Gets the disableAppPinIfDevicePinIsSet property value. Indicates whether use of the app pin is required if the device pin is set.
func (m *ManagedAppProtection) GetDisableAppPinIfDevicePinIsSet()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableAppPinIfDevicePinIsSet
    }
}
// Gets the fingerprintBlocked property value. Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
func (m *ManagedAppProtection) GetFingerprintBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fingerprintBlocked
    }
}
// Gets the managedBrowser property value. Indicates in which managed browser(s) that internet links should be opened. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. Possible values are: notConfigured, microsoftEdge.
func (m *ManagedAppProtection) GetManagedBrowser()(*ManagedBrowserType) {
    if m == nil {
        return nil
    } else {
        return m.managedBrowser
    }
}
// Gets the managedBrowserToOpenLinksRequired property value. Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
func (m *ManagedAppProtection) GetManagedBrowserToOpenLinksRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.managedBrowserToOpenLinksRequired
    }
}
// Gets the maximumPinRetries property value. Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
func (m *ManagedAppProtection) GetMaximumPinRetries()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumPinRetries
    }
}
// Gets the minimumPinLength property value. Minimum pin length required for an app-level pin if PinRequired is set to True
func (m *ManagedAppProtection) GetMinimumPinLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumPinLength
    }
}
// Gets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumRequiredAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredAppVersion
    }
}
// Gets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumRequiredOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredOsVersion
    }
}
// Gets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app.
func (m *ManagedAppProtection) GetMinimumWarningAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningAppVersion
    }
}
// Gets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumWarningOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningOsVersion
    }
}
// Gets the organizationalCredentialsRequired property value. Indicates whether organizational credentials are required for app use.
func (m *ManagedAppProtection) GetOrganizationalCredentialsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.organizationalCredentialsRequired
    }
}
// Gets the periodBeforePinReset property value. TimePeriod before the all-level pin must be reset if PinRequired is set to True.
func (m *ManagedAppProtection) GetPeriodBeforePinReset()(*string) {
    if m == nil {
        return nil
    } else {
        return m.periodBeforePinReset
    }
}
// Gets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet.
func (m *ManagedAppProtection) GetPeriodOfflineBeforeAccessCheck()(*string) {
    if m == nil {
        return nil
    } else {
        return m.periodOfflineBeforeAccessCheck
    }
}
// Gets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
func (m *ManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced()(*string) {
    if m == nil {
        return nil
    } else {
        return m.periodOfflineBeforeWipeIsEnforced
    }
}
// Gets the periodOnlineBeforeAccessCheck property value. The period after which access is checked when the device is connected to the internet.
func (m *ManagedAppProtection) GetPeriodOnlineBeforeAccessCheck()(*string) {
    if m == nil {
        return nil
    } else {
        return m.periodOnlineBeforeAccessCheck
    }
}
// Gets the pinCharacterSet property value. Character set which may be used for an app-level pin if PinRequired is set to True. Possible values are: numeric, alphanumericAndSymbol.
func (m *ManagedAppProtection) GetPinCharacterSet()(*ManagedAppPinCharacterSet) {
    if m == nil {
        return nil
    } else {
        return m.pinCharacterSet
    }
}
// Gets the pinRequired property value. Indicates whether an app-level pin is required.
func (m *ManagedAppProtection) GetPinRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.pinRequired
    }
}
// Gets the printBlocked property value. Indicates whether printing is allowed from managed apps.
func (m *ManagedAppProtection) GetPrintBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.printBlocked
    }
}
// Gets the saveAsBlocked property value. Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
func (m *ManagedAppProtection) GetSaveAsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.saveAsBlocked
    }
}
// Gets the simplePinBlocked property value. Indicates whether simplePin is blocked.
func (m *ManagedAppProtection) GetSimplePinBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.simplePinBlocked
    }
}
// The deserialization information for the current model
func (m *ManagedAppProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagedAppPolicy.GetFieldDeserializers()
    res["allowedDataStorageLocations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseManagedAppDataStorageLocation)
        if err != nil {
            return err
        }
        res := make([]ManagedAppDataStorageLocation, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedAppDataStorageLocation))
        }
        m.SetAllowedDataStorageLocations(res)
        return nil
    }
    res["allowedInboundDataTransferSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        cast := val.(ManagedAppDataTransferLevel)
        m.SetAllowedInboundDataTransferSources(&cast)
        return nil
    }
    res["allowedOutboundClipboardSharingLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppClipboardSharingLevel)
        if err != nil {
            return err
        }
        cast := val.(ManagedAppClipboardSharingLevel)
        m.SetAllowedOutboundClipboardSharingLevel(&cast)
        return nil
    }
    res["allowedOutboundDataTransferDestinations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        cast := val.(ManagedAppDataTransferLevel)
        m.SetAllowedOutboundDataTransferDestinations(&cast)
        return nil
    }
    res["contactSyncBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetContactSyncBlocked(val)
        return nil
    }
    res["dataBackupBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDataBackupBlocked(val)
        return nil
    }
    res["deviceComplianceRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDeviceComplianceRequired(val)
        return nil
    }
    res["disableAppPinIfDevicePinIsSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDisableAppPinIfDevicePinIsSet(val)
        return nil
    }
    res["fingerprintBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetFingerprintBlocked(val)
        return nil
    }
    res["managedBrowser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedBrowserType)
        if err != nil {
            return err
        }
        cast := val.(ManagedBrowserType)
        m.SetManagedBrowser(&cast)
        return nil
    }
    res["managedBrowserToOpenLinksRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetManagedBrowserToOpenLinksRequired(val)
        return nil
    }
    res["maximumPinRetries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMaximumPinRetries(val)
        return nil
    }
    res["minimumPinLength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetMinimumPinLength(val)
        return nil
    }
    res["minimumRequiredAppVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumRequiredAppVersion(val)
        return nil
    }
    res["minimumRequiredOsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumRequiredOsVersion(val)
        return nil
    }
    res["minimumWarningAppVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumWarningAppVersion(val)
        return nil
    }
    res["minimumWarningOsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMinimumWarningOsVersion(val)
        return nil
    }
    res["organizationalCredentialsRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetOrganizationalCredentialsRequired(val)
        return nil
    }
    res["periodBeforePinReset"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPeriodBeforePinReset(val)
        return nil
    }
    res["periodOfflineBeforeAccessCheck"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPeriodOfflineBeforeAccessCheck(val)
        return nil
    }
    res["periodOfflineBeforeWipeIsEnforced"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPeriodOfflineBeforeWipeIsEnforced(val)
        return nil
    }
    res["periodOnlineBeforeAccessCheck"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPeriodOnlineBeforeAccessCheck(val)
        return nil
    }
    res["pinCharacterSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppPinCharacterSet)
        if err != nil {
            return err
        }
        cast := val.(ManagedAppPinCharacterSet)
        m.SetPinCharacterSet(&cast)
        return nil
    }
    res["pinRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPinRequired(val)
        return nil
    }
    res["printBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetPrintBlocked(val)
        return nil
    }
    res["saveAsBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSaveAsBlocked(val)
        return nil
    }
    res["simplePinBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetSimplePinBlocked(val)
        return nil
    }
    return res
}
func (m *ManagedAppProtection) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ManagedAppProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ManagedAppPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("allowedDataStorageLocations", SerializeManagedAppDataStorageLocation(m.GetAllowedDataStorageLocations()))
        if err != nil {
            return err
        }
    }
    if m.GetAllowedInboundDataTransferSources() != nil {
        cast := m.GetAllowedInboundDataTransferSources().String()
        err = writer.WriteStringValue("allowedInboundDataTransferSources", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowedOutboundClipboardSharingLevel() != nil {
        cast := m.GetAllowedOutboundClipboardSharingLevel().String()
        err = writer.WriteStringValue("allowedOutboundClipboardSharingLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowedOutboundDataTransferDestinations() != nil {
        cast := m.GetAllowedOutboundDataTransferDestinations().String()
        err = writer.WriteStringValue("allowedOutboundDataTransferDestinations", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("contactSyncBlocked", m.GetContactSyncBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("dataBackupBlocked", m.GetDataBackupBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deviceComplianceRequired", m.GetDeviceComplianceRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableAppPinIfDevicePinIsSet", m.GetDisableAppPinIfDevicePinIsSet())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("fingerprintBlocked", m.GetFingerprintBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetManagedBrowser() != nil {
        cast := m.GetManagedBrowser().String()
        err = writer.WriteStringValue("managedBrowser", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("managedBrowserToOpenLinksRequired", m.GetManagedBrowserToOpenLinksRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("maximumPinRetries", m.GetMaximumPinRetries())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("minimumPinLength", m.GetMinimumPinLength())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredAppVersion", m.GetMinimumRequiredAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredOsVersion", m.GetMinimumRequiredOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWarningAppVersion", m.GetMinimumWarningAppVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWarningOsVersion", m.GetMinimumWarningOsVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("organizationalCredentialsRequired", m.GetOrganizationalCredentialsRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("periodBeforePinReset", m.GetPeriodBeforePinReset())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("periodOfflineBeforeAccessCheck", m.GetPeriodOfflineBeforeAccessCheck())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("periodOfflineBeforeWipeIsEnforced", m.GetPeriodOfflineBeforeWipeIsEnforced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("periodOnlineBeforeAccessCheck", m.GetPeriodOnlineBeforeAccessCheck())
        if err != nil {
            return err
        }
    }
    if m.GetPinCharacterSet() != nil {
        cast := m.GetPinCharacterSet().String()
        err = writer.WriteStringValue("pinCharacterSet", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("pinRequired", m.GetPinRequired())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("printBlocked", m.GetPrintBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("saveAsBlocked", m.GetSaveAsBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("simplePinBlocked", m.GetSimplePinBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the allowedDataStorageLocations property value. Data storage locations where a user may store managed data.
// Parameters:
//  - value : Value to set for the allowedDataStorageLocations property.
func (m *ManagedAppProtection) SetAllowedDataStorageLocations(value []ManagedAppDataStorageLocation)() {
    m.allowedDataStorageLocations = value
}
// Sets the allowedInboundDataTransferSources property value. Sources from which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
// Parameters:
//  - value : Value to set for the allowedInboundDataTransferSources property.
func (m *ManagedAppProtection) SetAllowedInboundDataTransferSources(value *ManagedAppDataTransferLevel)() {
    m.allowedInboundDataTransferSources = value
}
// Sets the allowedOutboundClipboardSharingLevel property value. The level to which the clipboard may be shared between apps on the managed device. Possible values are: allApps, managedAppsWithPasteIn, managedApps, blocked.
// Parameters:
//  - value : Value to set for the allowedOutboundClipboardSharingLevel property.
func (m *ManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(value *ManagedAppClipboardSharingLevel)() {
    m.allowedOutboundClipboardSharingLevel = value
}
// Sets the allowedOutboundDataTransferDestinations property value. Destinations to which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
// Parameters:
//  - value : Value to set for the allowedOutboundDataTransferDestinations property.
func (m *ManagedAppProtection) SetAllowedOutboundDataTransferDestinations(value *ManagedAppDataTransferLevel)() {
    m.allowedOutboundDataTransferDestinations = value
}
// Sets the contactSyncBlocked property value. Indicates whether contacts can be synced to the user's device.
// Parameters:
//  - value : Value to set for the contactSyncBlocked property.
func (m *ManagedAppProtection) SetContactSyncBlocked(value *bool)() {
    m.contactSyncBlocked = value
}
// Sets the dataBackupBlocked property value. Indicates whether the backup of a managed app's data is blocked.
// Parameters:
//  - value : Value to set for the dataBackupBlocked property.
func (m *ManagedAppProtection) SetDataBackupBlocked(value *bool)() {
    m.dataBackupBlocked = value
}
// Sets the deviceComplianceRequired property value. Indicates whether device compliance is required.
// Parameters:
//  - value : Value to set for the deviceComplianceRequired property.
func (m *ManagedAppProtection) SetDeviceComplianceRequired(value *bool)() {
    m.deviceComplianceRequired = value
}
// Sets the disableAppPinIfDevicePinIsSet property value. Indicates whether use of the app pin is required if the device pin is set.
// Parameters:
//  - value : Value to set for the disableAppPinIfDevicePinIsSet property.
func (m *ManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(value *bool)() {
    m.disableAppPinIfDevicePinIsSet = value
}
// Sets the fingerprintBlocked property value. Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
// Parameters:
//  - value : Value to set for the fingerprintBlocked property.
func (m *ManagedAppProtection) SetFingerprintBlocked(value *bool)() {
    m.fingerprintBlocked = value
}
// Sets the managedBrowser property value. Indicates in which managed browser(s) that internet links should be opened. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. Possible values are: notConfigured, microsoftEdge.
// Parameters:
//  - value : Value to set for the managedBrowser property.
func (m *ManagedAppProtection) SetManagedBrowser(value *ManagedBrowserType)() {
    m.managedBrowser = value
}
// Sets the managedBrowserToOpenLinksRequired property value. Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
// Parameters:
//  - value : Value to set for the managedBrowserToOpenLinksRequired property.
func (m *ManagedAppProtection) SetManagedBrowserToOpenLinksRequired(value *bool)() {
    m.managedBrowserToOpenLinksRequired = value
}
// Sets the maximumPinRetries property value. Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
// Parameters:
//  - value : Value to set for the maximumPinRetries property.
func (m *ManagedAppProtection) SetMaximumPinRetries(value *int32)() {
    m.maximumPinRetries = value
}
// Sets the minimumPinLength property value. Minimum pin length required for an app-level pin if PinRequired is set to True
// Parameters:
//  - value : Value to set for the minimumPinLength property.
func (m *ManagedAppProtection) SetMinimumPinLength(value *int32)() {
    m.minimumPinLength = value
}
// Sets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data.
// Parameters:
//  - value : Value to set for the minimumRequiredAppVersion property.
func (m *ManagedAppProtection) SetMinimumRequiredAppVersion(value *string)() {
    m.minimumRequiredAppVersion = value
}
// Sets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data.
// Parameters:
//  - value : Value to set for the minimumRequiredOsVersion property.
func (m *ManagedAppProtection) SetMinimumRequiredOsVersion(value *string)() {
    m.minimumRequiredOsVersion = value
}
// Sets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app.
// Parameters:
//  - value : Value to set for the minimumWarningAppVersion property.
func (m *ManagedAppProtection) SetMinimumWarningAppVersion(value *string)() {
    m.minimumWarningAppVersion = value
}
// Sets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
// Parameters:
//  - value : Value to set for the minimumWarningOsVersion property.
func (m *ManagedAppProtection) SetMinimumWarningOsVersion(value *string)() {
    m.minimumWarningOsVersion = value
}
// Sets the organizationalCredentialsRequired property value. Indicates whether organizational credentials are required for app use.
// Parameters:
//  - value : Value to set for the organizationalCredentialsRequired property.
func (m *ManagedAppProtection) SetOrganizationalCredentialsRequired(value *bool)() {
    m.organizationalCredentialsRequired = value
}
// Sets the periodBeforePinReset property value. TimePeriod before the all-level pin must be reset if PinRequired is set to True.
// Parameters:
//  - value : Value to set for the periodBeforePinReset property.
func (m *ManagedAppProtection) SetPeriodBeforePinReset(value *string)() {
    m.periodBeforePinReset = value
}
// Sets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet.
// Parameters:
//  - value : Value to set for the periodOfflineBeforeAccessCheck property.
func (m *ManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(value *string)() {
    m.periodOfflineBeforeAccessCheck = value
}
// Sets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
// Parameters:
//  - value : Value to set for the periodOfflineBeforeWipeIsEnforced property.
func (m *ManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(value *string)() {
    m.periodOfflineBeforeWipeIsEnforced = value
}
// Sets the periodOnlineBeforeAccessCheck property value. The period after which access is checked when the device is connected to the internet.
// Parameters:
//  - value : Value to set for the periodOnlineBeforeAccessCheck property.
func (m *ManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(value *string)() {
    m.periodOnlineBeforeAccessCheck = value
}
// Sets the pinCharacterSet property value. Character set which may be used for an app-level pin if PinRequired is set to True. Possible values are: numeric, alphanumericAndSymbol.
// Parameters:
//  - value : Value to set for the pinCharacterSet property.
func (m *ManagedAppProtection) SetPinCharacterSet(value *ManagedAppPinCharacterSet)() {
    m.pinCharacterSet = value
}
// Sets the pinRequired property value. Indicates whether an app-level pin is required.
// Parameters:
//  - value : Value to set for the pinRequired property.
func (m *ManagedAppProtection) SetPinRequired(value *bool)() {
    m.pinRequired = value
}
// Sets the printBlocked property value. Indicates whether printing is allowed from managed apps.
// Parameters:
//  - value : Value to set for the printBlocked property.
func (m *ManagedAppProtection) SetPrintBlocked(value *bool)() {
    m.printBlocked = value
}
// Sets the saveAsBlocked property value. Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
// Parameters:
//  - value : Value to set for the saveAsBlocked property.
func (m *ManagedAppProtection) SetSaveAsBlocked(value *bool)() {
    m.saveAsBlocked = value
}
// Sets the simplePinBlocked property value. Indicates whether simplePin is blocked.
// Parameters:
//  - value : Value to set for the simplePinBlocked property.
func (m *ManagedAppProtection) SetSimplePinBlocked(value *bool)() {
    m.simplePinBlocked = value
}
