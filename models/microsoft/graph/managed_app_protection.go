package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ManagedAppProtection struct {
    ManagedAppPolicy
    allowedDataStorageLocations []ManagedAppDataStorageLocation;
    allowedInboundDataTransferSources *ManagedAppDataTransferLevel;
    allowedOutboundClipboardSharingLevel *ManagedAppClipboardSharingLevel;
    allowedOutboundDataTransferDestinations *ManagedAppDataTransferLevel;
    contactSyncBlocked *bool;
    dataBackupBlocked *bool;
    deviceComplianceRequired *bool;
    disableAppPinIfDevicePinIsSet *bool;
    fingerprintBlocked *bool;
    managedBrowser *ManagedBrowserType;
    managedBrowserToOpenLinksRequired *bool;
    maximumPinRetries *int32;
    minimumPinLength *int32;
    minimumRequiredAppVersion *string;
    minimumRequiredOsVersion *string;
    minimumWarningAppVersion *string;
    minimumWarningOsVersion *string;
    organizationalCredentialsRequired *bool;
    periodBeforePinReset *string;
    periodOfflineBeforeAccessCheck *string;
    periodOfflineBeforeWipeIsEnforced *string;
    periodOnlineBeforeAccessCheck *string;
    pinCharacterSet *ManagedAppPinCharacterSet;
    pinRequired *bool;
    printBlocked *bool;
    saveAsBlocked *bool;
    simplePinBlocked *bool;
}
func NewManagedAppProtection()(*ManagedAppProtection) {
    m := &ManagedAppProtection{
        ManagedAppPolicy: *NewManagedAppPolicy(),
    }
    return m
}
func (m *ManagedAppProtection) GetAllowedDataStorageLocations()([]ManagedAppDataStorageLocation) {
    if m == nil {
        return nil
    } else {
        return m.allowedDataStorageLocations
    }
}
func (m *ManagedAppProtection) GetAllowedInboundDataTransferSources()(*ManagedAppDataTransferLevel) {
    if m == nil {
        return nil
    } else {
        return m.allowedInboundDataTransferSources
    }
}
func (m *ManagedAppProtection) GetAllowedOutboundClipboardSharingLevel()(*ManagedAppClipboardSharingLevel) {
    if m == nil {
        return nil
    } else {
        return m.allowedOutboundClipboardSharingLevel
    }
}
func (m *ManagedAppProtection) GetAllowedOutboundDataTransferDestinations()(*ManagedAppDataTransferLevel) {
    if m == nil {
        return nil
    } else {
        return m.allowedOutboundDataTransferDestinations
    }
}
func (m *ManagedAppProtection) GetContactSyncBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.contactSyncBlocked
    }
}
func (m *ManagedAppProtection) GetDataBackupBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.dataBackupBlocked
    }
}
func (m *ManagedAppProtection) GetDeviceComplianceRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deviceComplianceRequired
    }
}
func (m *ManagedAppProtection) GetDisableAppPinIfDevicePinIsSet()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableAppPinIfDevicePinIsSet
    }
}
func (m *ManagedAppProtection) GetFingerprintBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.fingerprintBlocked
    }
}
func (m *ManagedAppProtection) GetManagedBrowser()(*ManagedBrowserType) {
    if m == nil {
        return nil
    } else {
        return m.managedBrowser
    }
}
func (m *ManagedAppProtection) GetManagedBrowserToOpenLinksRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.managedBrowserToOpenLinksRequired
    }
}
func (m *ManagedAppProtection) GetMaximumPinRetries()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.maximumPinRetries
    }
}
func (m *ManagedAppProtection) GetMinimumPinLength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.minimumPinLength
    }
}
func (m *ManagedAppProtection) GetMinimumRequiredAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredAppVersion
    }
}
func (m *ManagedAppProtection) GetMinimumRequiredOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredOsVersion
    }
}
func (m *ManagedAppProtection) GetMinimumWarningAppVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningAppVersion
    }
}
func (m *ManagedAppProtection) GetMinimumWarningOsVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningOsVersion
    }
}
func (m *ManagedAppProtection) GetOrganizationalCredentialsRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.organizationalCredentialsRequired
    }
}
func (m *ManagedAppProtection) GetPeriodBeforePinReset()(*string) {
    if m == nil {
        return nil
    } else {
        return m.periodBeforePinReset
    }
}
func (m *ManagedAppProtection) GetPeriodOfflineBeforeAccessCheck()(*string) {
    if m == nil {
        return nil
    } else {
        return m.periodOfflineBeforeAccessCheck
    }
}
func (m *ManagedAppProtection) GetPeriodOfflineBeforeWipeIsEnforced()(*string) {
    if m == nil {
        return nil
    } else {
        return m.periodOfflineBeforeWipeIsEnforced
    }
}
func (m *ManagedAppProtection) GetPeriodOnlineBeforeAccessCheck()(*string) {
    if m == nil {
        return nil
    } else {
        return m.periodOnlineBeforeAccessCheck
    }
}
func (m *ManagedAppProtection) GetPinCharacterSet()(*ManagedAppPinCharacterSet) {
    if m == nil {
        return nil
    } else {
        return m.pinCharacterSet
    }
}
func (m *ManagedAppProtection) GetPinRequired()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.pinRequired
    }
}
func (m *ManagedAppProtection) GetPrintBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.printBlocked
    }
}
func (m *ManagedAppProtection) GetSaveAsBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.saveAsBlocked
    }
}
func (m *ManagedAppProtection) GetSimplePinBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.simplePinBlocked
    }
}
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
func (m *ManagedAppProtection) SetAllowedDataStorageLocations(value []ManagedAppDataStorageLocation)() {
    m.allowedDataStorageLocations = value
}
func (m *ManagedAppProtection) SetAllowedInboundDataTransferSources(value *ManagedAppDataTransferLevel)() {
    m.allowedInboundDataTransferSources = value
}
func (m *ManagedAppProtection) SetAllowedOutboundClipboardSharingLevel(value *ManagedAppClipboardSharingLevel)() {
    m.allowedOutboundClipboardSharingLevel = value
}
func (m *ManagedAppProtection) SetAllowedOutboundDataTransferDestinations(value *ManagedAppDataTransferLevel)() {
    m.allowedOutboundDataTransferDestinations = value
}
func (m *ManagedAppProtection) SetContactSyncBlocked(value *bool)() {
    m.contactSyncBlocked = value
}
func (m *ManagedAppProtection) SetDataBackupBlocked(value *bool)() {
    m.dataBackupBlocked = value
}
func (m *ManagedAppProtection) SetDeviceComplianceRequired(value *bool)() {
    m.deviceComplianceRequired = value
}
func (m *ManagedAppProtection) SetDisableAppPinIfDevicePinIsSet(value *bool)() {
    m.disableAppPinIfDevicePinIsSet = value
}
func (m *ManagedAppProtection) SetFingerprintBlocked(value *bool)() {
    m.fingerprintBlocked = value
}
func (m *ManagedAppProtection) SetManagedBrowser(value *ManagedBrowserType)() {
    m.managedBrowser = value
}
func (m *ManagedAppProtection) SetManagedBrowserToOpenLinksRequired(value *bool)() {
    m.managedBrowserToOpenLinksRequired = value
}
func (m *ManagedAppProtection) SetMaximumPinRetries(value *int32)() {
    m.maximumPinRetries = value
}
func (m *ManagedAppProtection) SetMinimumPinLength(value *int32)() {
    m.minimumPinLength = value
}
func (m *ManagedAppProtection) SetMinimumRequiredAppVersion(value *string)() {
    m.minimumRequiredAppVersion = value
}
func (m *ManagedAppProtection) SetMinimumRequiredOsVersion(value *string)() {
    m.minimumRequiredOsVersion = value
}
func (m *ManagedAppProtection) SetMinimumWarningAppVersion(value *string)() {
    m.minimumWarningAppVersion = value
}
func (m *ManagedAppProtection) SetMinimumWarningOsVersion(value *string)() {
    m.minimumWarningOsVersion = value
}
func (m *ManagedAppProtection) SetOrganizationalCredentialsRequired(value *bool)() {
    m.organizationalCredentialsRequired = value
}
func (m *ManagedAppProtection) SetPeriodBeforePinReset(value *string)() {
    m.periodBeforePinReset = value
}
func (m *ManagedAppProtection) SetPeriodOfflineBeforeAccessCheck(value *string)() {
    m.periodOfflineBeforeAccessCheck = value
}
func (m *ManagedAppProtection) SetPeriodOfflineBeforeWipeIsEnforced(value *string)() {
    m.periodOfflineBeforeWipeIsEnforced = value
}
func (m *ManagedAppProtection) SetPeriodOnlineBeforeAccessCheck(value *string)() {
    m.periodOnlineBeforeAccessCheck = value
}
func (m *ManagedAppProtection) SetPinCharacterSet(value *ManagedAppPinCharacterSet)() {
    m.pinCharacterSet = value
}
func (m *ManagedAppProtection) SetPinRequired(value *bool)() {
    m.pinRequired = value
}
func (m *ManagedAppProtection) SetPrintBlocked(value *bool)() {
    m.printBlocked = value
}
func (m *ManagedAppProtection) SetSaveAsBlocked(value *bool)() {
    m.saveAsBlocked = value
}
func (m *ManagedAppProtection) SetSimplePinBlocked(value *bool)() {
    m.simplePinBlocked = value
}
