package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ManagedAppProtection 
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
    periodBeforePinReset *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The period after which access is checked when the device is not connected to the internet.
    periodOfflineBeforeAccessCheck *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
    periodOfflineBeforeWipeIsEnforced *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // The period after which access is checked when the device is connected to the internet.
    periodOnlineBeforeAccessCheck *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
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
// NewManagedAppProtection instantiates a new managedAppProtection and sets the default values.
func NewManagedAppProtection()(*ManagedAppProtection) {
    m := &ManagedAppProtection{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    return m
}
// GetAllowedDataStorageLocations gets the allowedDataStorageLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) GetAllowedDataStorageLocations()([]ManagedAppDataStorageLocation) {
    if m == nil {
        return nil
    } else {
        return m.allowedDataStorageLocations
    }
}
// GetAllowedInboundDataTransferSources gets the allowedInboundDataTransferSources property value. Sources from which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
func (m *ManagedAppProtection) GetAllowedInboundDataTransferSources()(*ManagedAppDataTransferLevel) {
    if m == nil {
        return nil
    } else {
        return m.allowedInboundDataTransferSources
    }
}
// GetAllowedOutboundClipboardSharingLevel gets the allowedOutboundClipboardSharingLevel property value. The level to which the clipboard may be shared between apps on the managed device. Possible values are: allApps, managedAppsWithPasteIn, managedApps, blocked.
func (m *ManagedAppProtection) GetAllowedOutboundClipboardSharingLevel()(*ManagedAppClipboardSharingLevel) {
    if m == nil {
        return nil
    } else {
        return m.allowedOutboundClipboardSharingLevel
    }
}
// GetAllowedOutboundDataTransferDestinations gets the allowedOutboundDataTransferDestinations property value. Destinations to which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
func (m *ManagedAppProtection) GetAllowedOutboundDataTransferDestinations()(*ManagedAppDataTransferLevel) {
    if m == nil {
        return nil
    } else {
        return m.allowedOutboundDataTransferDestinations
    }
}
// GetContactSyncBlocked gets the contactSyncBlocked property value. Indicates whether contacts can be synced to the user's device.
func (m *ManagedAppProtection) GetContactSyncBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contactSyncBlocked
    }
}
// GetDataBackupBlocked gets the dataBackupBlocked property value. Indicates whether the backup of a managed app's data is blocked.
func (m *ManagedAppProtection) GetDataBackupBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.dataBackupBlocked
    }
}
// GetDeviceComplianceRequired gets the deviceComplianceRequired property value. Indicates whether device compliance is required.
func (m *ManagedAppProtection) GetDeviceComplianceRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceRequired
    }
}
// GetDisableAppPinIfDevicePinIsSet gets the disableAppPinIfDevicePinIsSet property value. Indicates whether use of the app pin is required if the device pin is set.
func (m *ManagedAppProtection) GetDisableAppPinIfDevicePinIsSet()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableAppPinIfDevicePinIsSet
    }
}
// GetFingerprintBlocked gets the fingerprintBlocked property value. Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
func (m *ManagedAppProtection) GetFingerprintBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fingerprintBlocked
    }
}
// GetManagedBrowser gets the managedBrowser property value. Indicates in which managed browser(s) that internet links should be opened. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. Possible values are: notConfigured, microsoftEdge.
func (m *ManagedAppProtection) GetManagedBrowser()(*ManagedBrowserType) {
    if m == nil {
        return nil
    } else {
        return m.managedBrowser
    }
}
// GetManagedBrowserToOpenLinksRequired gets the managedBrowserToOpenLinksRequired property value. Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
func (m *ManagedAppProtection) GetManagedBrowserToOpenLinksRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.managedBrowserToOpenLinksRequired
    }
}
// GetMaximumPinRetries gets the maximumPinRetries property value. Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
func (m *ManagedAppProtection) GetMaximumPinRetries()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumPinRetries
    }
}
// GetMinimumPinLength gets the minimumPinLength property value. Minimum pin length required for an app-level pin if PinRequired is set to True
func (m *ManagedAppProtection) GetMinimumPinLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumPinLength
    }
}
// GetMinimumRequiredAppVersion gets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumRequiredAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredAppVersion
    }
}
// GetMinimumRequiredOsVersion gets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumRequiredOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredOsVersion
    }
}
// GetMinimumWarningAppVersion gets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app.
func (m *ManagedAppProtection) GetMinimumWarningAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningAppVersion
    }
}
// GetMinimumWarningOsVersion gets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
func (m *ManagedAppProtection) GetMinimumWarningOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningOsVersion
    }
}
// GetOrganizationalCredentialsRequired gets the organizationalCredentialsRequired property value. Indicates whether organizational credentials are required for app use.
func (m *ManagedAppProtection) GetOrganizationalCredentialsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.organizationalCredentialsRequired
    }
}
// GetPeriodBeforePinReset gets the periodBeforePinReset property value. TimePeriod before the all-level pin must be reset if PinRequired is set to True.
func (m *ManagedAppProtection) GetPeriodBeforePinReset()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.periodBeforePinReset
    }
}
// GetPeriodOfflineBeforeAccessCheck gets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet.
func (m *ManagedAppProtection) GetPeriodOfflineBeforeAccessCheck()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.periodOfflineBeforeAccessCheck
    }
}
// GetPeriodOfflineBeforeWipeIsEnforced gets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
func (m *ManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.periodOfflineBeforeWipeIsEnforced
    }
}
// GetPeriodOnlineBeforeAccessCheck gets the periodOnlineBeforeAccessCheck property value. The period after which access is checked when the device is connected to the internet.
func (m *ManagedAppProtection) GetPeriodOnlineBeforeAccessCheck()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.periodOnlineBeforeAccessCheck
    }
}
// GetPinCharacterSet gets the pinCharacterSet property value. Character set which may be used for an app-level pin if PinRequired is set to True. Possible values are: numeric, alphanumericAndSymbol.
func (m *ManagedAppProtection) GetPinCharacterSet()(*ManagedAppPinCharacterSet) {
    if m == nil {
        return nil
    } else {
        return m.pinCharacterSet
    }
}
// GetPinRequired gets the pinRequired property value. Indicates whether an app-level pin is required.
func (m *ManagedAppProtection) GetPinRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.pinRequired
    }
}
// GetPrintBlocked gets the printBlocked property value. Indicates whether printing is allowed from managed apps.
func (m *ManagedAppProtection) GetPrintBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.printBlocked
    }
}
// GetSaveAsBlocked gets the saveAsBlocked property value. Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
func (m *ManagedAppProtection) GetSaveAsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.saveAsBlocked
    }
}
// GetSimplePinBlocked gets the simplePinBlocked property value. Indicates whether simplePin is blocked.
func (m *ManagedAppProtection) GetSimplePinBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.simplePinBlocked
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ManagedAppProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagedAppPolicy.GetFieldDeserializers()
    res["allowedDataStorageLocations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseManagedAppDataStorageLocation)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedAppDataStorageLocation, len(val))
            for i, v := range val {
                res[i] = *(v.(*ManagedAppDataStorageLocation))
            }
            m.SetAllowedDataStorageLocations(res)
        }
        return nil
    }
    res["allowedInboundDataTransferSources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedInboundDataTransferSources(val.(*ManagedAppDataTransferLevel))
        }
        return nil
    }
    res["allowedOutboundClipboardSharingLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppClipboardSharingLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundClipboardSharingLevel(val.(*ManagedAppClipboardSharingLevel))
        }
        return nil
    }
    res["allowedOutboundDataTransferDestinations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataTransferLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowedOutboundDataTransferDestinations(val.(*ManagedAppDataTransferLevel))
        }
        return nil
    }
    res["contactSyncBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetContactSyncBlocked(val)
        }
        return nil
    }
    res["dataBackupBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDataBackupBlocked(val)
        }
        return nil
    }
    res["deviceComplianceRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceComplianceRequired(val)
        }
        return nil
    }
    res["disableAppPinIfDevicePinIsSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableAppPinIfDevicePinIsSet(val)
        }
        return nil
    }
    res["fingerprintBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFingerprintBlocked(val)
        }
        return nil
    }
    res["managedBrowser"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedBrowserType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBrowser(val.(*ManagedBrowserType))
        }
        return nil
    }
    res["managedBrowserToOpenLinksRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManagedBrowserToOpenLinksRequired(val)
        }
        return nil
    }
    res["maximumPinRetries"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaximumPinRetries(val)
        }
        return nil
    }
    res["minimumPinLength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumPinLength(val)
        }
        return nil
    }
    res["minimumRequiredAppVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredAppVersion(val)
        }
        return nil
    }
    res["minimumRequiredOsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredOsVersion(val)
        }
        return nil
    }
    res["minimumWarningAppVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningAppVersion(val)
        }
        return nil
    }
    res["minimumWarningOsVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningOsVersion(val)
        }
        return nil
    }
    res["organizationalCredentialsRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOrganizationalCredentialsRequired(val)
        }
        return nil
    }
    res["periodBeforePinReset"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodBeforePinReset(val)
        }
        return nil
    }
    res["periodOfflineBeforeAccessCheck"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodOfflineBeforeAccessCheck(val)
        }
        return nil
    }
    res["periodOfflineBeforeWipeIsEnforced"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodOfflineBeforeWipeIsEnforced(val)
        }
        return nil
    }
    res["periodOnlineBeforeAccessCheck"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPeriodOnlineBeforeAccessCheck(val)
        }
        return nil
    }
    res["pinCharacterSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppPinCharacterSet)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinCharacterSet(val.(*ManagedAppPinCharacterSet))
        }
        return nil
    }
    res["pinRequired"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPinRequired(val)
        }
        return nil
    }
    res["printBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPrintBlocked(val)
        }
        return nil
    }
    res["saveAsBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSaveAsBlocked(val)
        }
        return nil
    }
    res["simplePinBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSimplePinBlocked(val)
        }
        return nil
    }
    return res
}
func (m *ManagedAppProtection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ManagedAppProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ManagedAppPolicy.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAllowedDataStorageLocations() != nil {
        err = writer.WriteCollectionOfStringValues("allowedDataStorageLocations", SerializeManagedAppDataStorageLocation(m.GetAllowedDataStorageLocations()))
        if err != nil {
            return err
        }
    }
    if m.GetAllowedInboundDataTransferSources() != nil {
        cast := (*m.GetAllowedInboundDataTransferSources()).String()
        err = writer.WriteStringValue("allowedInboundDataTransferSources", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowedOutboundClipboardSharingLevel() != nil {
        cast := (*m.GetAllowedOutboundClipboardSharingLevel()).String()
        err = writer.WriteStringValue("allowedOutboundClipboardSharingLevel", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetAllowedOutboundDataTransferDestinations() != nil {
        cast := (*m.GetAllowedOutboundDataTransferDestinations()).String()
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
        cast := (*m.GetManagedBrowser()).String()
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
        err = writer.WriteISODurationValue("periodBeforePinReset", m.GetPeriodBeforePinReset())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("periodOfflineBeforeAccessCheck", m.GetPeriodOfflineBeforeAccessCheck())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("periodOfflineBeforeWipeIsEnforced", m.GetPeriodOfflineBeforeWipeIsEnforced())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteISODurationValue("periodOnlineBeforeAccessCheck", m.GetPeriodOnlineBeforeAccessCheck())
        if err != nil {
            return err
        }
    }
    if m.GetPinCharacterSet() != nil {
        cast := (*m.GetPinCharacterSet()).String()
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
// SetAllowedDataStorageLocations sets the allowedDataStorageLocations property value. Data storage locations where a user may store managed data.
func (m *ManagedAppProtection) SetAllowedDataStorageLocations(value []ManagedAppDataStorageLocation)() {
    if m != nil {
        m.allowedDataStorageLocations = value
    }
}
// SetAllowedInboundDataTransferSources sets the allowedInboundDataTransferSources property value. Sources from which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
func (m *ManagedAppProtection) SetAllowedInboundDataTransferSources(value *ManagedAppDataTransferLevel)() {
    if m != nil {
        m.allowedInboundDataTransferSources = value
    }
}
// SetAllowedOutboundClipboardSharingLevel sets the allowedOutboundClipboardSharingLevel property value. The level to which the clipboard may be shared between apps on the managed device. Possible values are: allApps, managedAppsWithPasteIn, managedApps, blocked.
func (m *ManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(value *ManagedAppClipboardSharingLevel)() {
    if m != nil {
        m.allowedOutboundClipboardSharingLevel = value
    }
}
// SetAllowedOutboundDataTransferDestinations sets the allowedOutboundDataTransferDestinations property value. Destinations to which data is allowed to be transferred. Possible values are: allApps, managedApps, none.
func (m *ManagedAppProtection) SetAllowedOutboundDataTransferDestinations(value *ManagedAppDataTransferLevel)() {
    if m != nil {
        m.allowedOutboundDataTransferDestinations = value
    }
}
// SetContactSyncBlocked sets the contactSyncBlocked property value. Indicates whether contacts can be synced to the user's device.
func (m *ManagedAppProtection) SetContactSyncBlocked(value *bool)() {
    if m != nil {
        m.contactSyncBlocked = value
    }
}
// SetDataBackupBlocked sets the dataBackupBlocked property value. Indicates whether the backup of a managed app's data is blocked.
func (m *ManagedAppProtection) SetDataBackupBlocked(value *bool)() {
    if m != nil {
        m.dataBackupBlocked = value
    }
}
// SetDeviceComplianceRequired sets the deviceComplianceRequired property value. Indicates whether device compliance is required.
func (m *ManagedAppProtection) SetDeviceComplianceRequired(value *bool)() {
    if m != nil {
        m.deviceComplianceRequired = value
    }
}
// SetDisableAppPinIfDevicePinIsSet sets the disableAppPinIfDevicePinIsSet property value. Indicates whether use of the app pin is required if the device pin is set.
func (m *ManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(value *bool)() {
    if m != nil {
        m.disableAppPinIfDevicePinIsSet = value
    }
}
// SetFingerprintBlocked sets the fingerprintBlocked property value. Indicates whether use of the fingerprint reader is allowed in place of a pin if PinRequired is set to True.
func (m *ManagedAppProtection) SetFingerprintBlocked(value *bool)() {
    if m != nil {
        m.fingerprintBlocked = value
    }
}
// SetManagedBrowser sets the managedBrowser property value. Indicates in which managed browser(s) that internet links should be opened. When this property is configured, ManagedBrowserToOpenLinksRequired should be true. Possible values are: notConfigured, microsoftEdge.
func (m *ManagedAppProtection) SetManagedBrowser(value *ManagedBrowserType)() {
    if m != nil {
        m.managedBrowser = value
    }
}
// SetManagedBrowserToOpenLinksRequired sets the managedBrowserToOpenLinksRequired property value. Indicates whether internet links should be opened in the managed browser app, or any custom browser specified by CustomBrowserProtocol (for iOS) or CustomBrowserPackageId/CustomBrowserDisplayName (for Android)
func (m *ManagedAppProtection) SetManagedBrowserToOpenLinksRequired(value *bool)() {
    if m != nil {
        m.managedBrowserToOpenLinksRequired = value
    }
}
// SetMaximumPinRetries sets the maximumPinRetries property value. Maximum number of incorrect pin retry attempts before the managed app is either blocked or wiped.
func (m *ManagedAppProtection) SetMaximumPinRetries(value *int32)() {
    if m != nil {
        m.maximumPinRetries = value
    }
}
// SetMinimumPinLength sets the minimumPinLength property value. Minimum pin length required for an app-level pin if PinRequired is set to True
func (m *ManagedAppProtection) SetMinimumPinLength(value *int32)() {
    if m != nil {
        m.minimumPinLength = value
    }
}
// SetMinimumRequiredAppVersion sets the minimumRequiredAppVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMinimumRequiredAppVersion(value *string)() {
    if m != nil {
        m.minimumRequiredAppVersion = value
    }
}
// SetMinimumRequiredOsVersion sets the minimumRequiredOsVersion property value. Versions less than the specified version will block the managed app from accessing company data.
func (m *ManagedAppProtection) SetMinimumRequiredOsVersion(value *string)() {
    if m != nil {
        m.minimumRequiredOsVersion = value
    }
}
// SetMinimumWarningAppVersion sets the minimumWarningAppVersion property value. Versions less than the specified version will result in warning message on the managed app.
func (m *ManagedAppProtection) SetMinimumWarningAppVersion(value *string)() {
    if m != nil {
        m.minimumWarningAppVersion = value
    }
}
// SetMinimumWarningOsVersion sets the minimumWarningOsVersion property value. Versions less than the specified version will result in warning message on the managed app from accessing company data.
func (m *ManagedAppProtection) SetMinimumWarningOsVersion(value *string)() {
    if m != nil {
        m.minimumWarningOsVersion = value
    }
}
// SetOrganizationalCredentialsRequired sets the organizationalCredentialsRequired property value. Indicates whether organizational credentials are required for app use.
func (m *ManagedAppProtection) SetOrganizationalCredentialsRequired(value *bool)() {
    if m != nil {
        m.organizationalCredentialsRequired = value
    }
}
// SetPeriodBeforePinReset sets the periodBeforePinReset property value. TimePeriod before the all-level pin must be reset if PinRequired is set to True.
func (m *ManagedAppProtection) SetPeriodBeforePinReset(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.periodBeforePinReset = value
    }
}
// SetPeriodOfflineBeforeAccessCheck sets the periodOfflineBeforeAccessCheck property value. The period after which access is checked when the device is not connected to the internet.
func (m *ManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.periodOfflineBeforeAccessCheck = value
    }
}
// SetPeriodOfflineBeforeWipeIsEnforced sets the periodOfflineBeforeWipeIsEnforced property value. The amount of time an app is allowed to remain disconnected from the internet before all managed data it is wiped.
func (m *ManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.periodOfflineBeforeWipeIsEnforced = value
    }
}
// SetPeriodOnlineBeforeAccessCheck sets the periodOnlineBeforeAccessCheck property value. The period after which access is checked when the device is connected to the internet.
func (m *ManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.periodOnlineBeforeAccessCheck = value
    }
}
// SetPinCharacterSet sets the pinCharacterSet property value. Character set which may be used for an app-level pin if PinRequired is set to True. Possible values are: numeric, alphanumericAndSymbol.
func (m *ManagedAppProtection) SetPinCharacterSet(value *ManagedAppPinCharacterSet)() {
    if m != nil {
        m.pinCharacterSet = value
    }
}
// SetPinRequired sets the pinRequired property value. Indicates whether an app-level pin is required.
func (m *ManagedAppProtection) SetPinRequired(value *bool)() {
    if m != nil {
        m.pinRequired = value
    }
}
// SetPrintBlocked sets the printBlocked property value. Indicates whether printing is allowed from managed apps.
func (m *ManagedAppProtection) SetPrintBlocked(value *bool)() {
    if m != nil {
        m.printBlocked = value
    }
}
// SetSaveAsBlocked sets the saveAsBlocked property value. Indicates whether users may use the 'Save As' menu item to save a copy of protected files.
func (m *ManagedAppProtection) SetSaveAsBlocked(value *bool)() {
    if m != nil {
        m.saveAsBlocked = value
    }
}
// SetSimplePinBlocked sets the simplePinBlocked property value. Indicates whether simplePin is blocked.
func (m *ManagedAppProtection) SetSimplePinBlocked(value *bool)() {
    if m != nil {
        m.simplePinBlocked = value
    }
}
