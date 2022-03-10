package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// DefaultManagedAppProtection provides operations to manage the deviceAppManagement singleton.
type DefaultManagedAppProtection struct {
    ManagedAppProtection
    // Type of encryption which should be used for data in a managed app. (iOS Only). Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked.
    appDataEncryptionType *ManagedAppDataEncryptionType;
    // List of apps to which the policy is deployed.
    apps []ManagedMobileAppable;
    // A set of string key and string value pairs to be sent to the affected users, unalterned by this service
    customSettings []KeyValuePairable;
    // Count of apps to which the current policy is deployed.
    deployedAppCount *int32;
    // Navigation property to deployment summary of the configuration.
    deploymentSummary ManagedAppPolicyDeploymentSummaryable;
    // When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only)
    disableAppEncryptionIfDeviceEncryptionIsEnabled *bool;
    // Indicates whether managed-app data should be encrypted. (Android only)
    encryptAppData *bool;
    // Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only)
    faceIdBlocked *bool;
    // Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android only)
    minimumRequiredPatchVersion *string;
    // Versions less than the specified version will block the managed app from accessing company data. (iOS Only)
    minimumRequiredSdkVersion *string;
    // Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android only)
    minimumWarningPatchVersion *string;
    // Indicates whether screen capture is blocked. (Android only)
    screenCaptureBlocked *bool;
}
// NewDefaultManagedAppProtection instantiates a new defaultManagedAppProtection and sets the default values.
func NewDefaultManagedAppProtection()(*DefaultManagedAppProtection) {
    m := &DefaultManagedAppProtection{
        ManagedAppProtection: *NewManagedAppProtection(),
    }
    return m
}
// CreateDefaultManagedAppProtectionFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDefaultManagedAppProtectionFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewDefaultManagedAppProtection(), nil
}
// GetAppDataEncryptionType gets the appDataEncryptionType property value. Type of encryption which should be used for data in a managed app. (iOS Only). Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked.
func (m *DefaultManagedAppProtection) GetAppDataEncryptionType()(*ManagedAppDataEncryptionType) {
    if m == nil {
        return nil
    } else {
        return m.appDataEncryptionType
    }
}
// GetApps gets the apps property value. List of apps to which the policy is deployed.
func (m *DefaultManagedAppProtection) GetApps()([]ManagedMobileAppable) {
    if m == nil {
        return nil
    } else {
        return m.apps
    }
}
// GetCustomSettings gets the customSettings property value. A set of string key and string value pairs to be sent to the affected users, unalterned by this service
func (m *DefaultManagedAppProtection) GetCustomSettings()([]KeyValuePairable) {
    if m == nil {
        return nil
    } else {
        return m.customSettings
    }
}
// GetDeployedAppCount gets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *DefaultManagedAppProtection) GetDeployedAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deployedAppCount
    }
}
// GetDeploymentSummary gets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *DefaultManagedAppProtection) GetDeploymentSummary()(ManagedAppPolicyDeploymentSummaryable) {
    if m == nil {
        return nil
    } else {
        return m.deploymentSummary
    }
}
// GetDisableAppEncryptionIfDeviceEncryptionIsEnabled gets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only)
func (m *DefaultManagedAppProtection) GetDisableAppEncryptionIfDeviceEncryptionIsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.disableAppEncryptionIfDeviceEncryptionIsEnabled
    }
}
// GetEncryptAppData gets the encryptAppData property value. Indicates whether managed-app data should be encrypted. (Android only)
func (m *DefaultManagedAppProtection) GetEncryptAppData()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.encryptAppData
    }
}
// GetFaceIdBlocked gets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only)
func (m *DefaultManagedAppProtection) GetFaceIdBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.faceIdBlocked
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DefaultManagedAppProtection) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagedAppProtection.GetFieldDeserializers()
    res["appDataEncryptionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseManagedAppDataEncryptionType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppDataEncryptionType(val.(*ManagedAppDataEncryptionType))
        }
        return nil
    }
    res["apps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateManagedMobileAppFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ManagedMobileAppable, len(val))
            for i, v := range val {
                res[i] = v.(ManagedMobileAppable)
            }
            m.SetApps(res)
        }
        return nil
    }
    res["customSettings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateKeyValuePairFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]KeyValuePairable, len(val))
            for i, v := range val {
                res[i] = v.(KeyValuePairable)
            }
            m.SetCustomSettings(res)
        }
        return nil
    }
    res["deployedAppCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeployedAppCount(val)
        }
        return nil
    }
    res["deploymentSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateManagedAppPolicyDeploymentSummaryFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeploymentSummary(val.(ManagedAppPolicyDeploymentSummaryable))
        }
        return nil
    }
    res["disableAppEncryptionIfDeviceEncryptionIsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(val)
        }
        return nil
    }
    res["encryptAppData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEncryptAppData(val)
        }
        return nil
    }
    res["faceIdBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFaceIdBlocked(val)
        }
        return nil
    }
    res["minimumRequiredPatchVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredPatchVersion(val)
        }
        return nil
    }
    res["minimumRequiredSdkVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumRequiredSdkVersion(val)
        }
        return nil
    }
    res["minimumWarningPatchVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMinimumWarningPatchVersion(val)
        }
        return nil
    }
    res["screenCaptureBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetScreenCaptureBlocked(val)
        }
        return nil
    }
    return res
}
// GetMinimumRequiredPatchVersion gets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android only)
func (m *DefaultManagedAppProtection) GetMinimumRequiredPatchVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredPatchVersion
    }
}
// GetMinimumRequiredSdkVersion gets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data. (iOS Only)
func (m *DefaultManagedAppProtection) GetMinimumRequiredSdkVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumRequiredSdkVersion
    }
}
// GetMinimumWarningPatchVersion gets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android only)
func (m *DefaultManagedAppProtection) GetMinimumWarningPatchVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.minimumWarningPatchVersion
    }
}
// GetScreenCaptureBlocked gets the screenCaptureBlocked property value. Indicates whether screen capture is blocked. (Android only)
func (m *DefaultManagedAppProtection) GetScreenCaptureBlocked()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.screenCaptureBlocked
    }
}
func (m *DefaultManagedAppProtection) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *DefaultManagedAppProtection) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ManagedAppProtection.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAppDataEncryptionType() != nil {
        cast := (*m.GetAppDataEncryptionType()).String()
        err = writer.WriteStringValue("appDataEncryptionType", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetApps() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("apps", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCustomSettings() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCustomSettings()))
        for i, v := range m.GetCustomSettings() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("customSettings", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("deployedAppCount", m.GetDeployedAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploymentSummary", m.GetDeploymentSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("disableAppEncryptionIfDeviceEncryptionIsEnabled", m.GetDisableAppEncryptionIfDeviceEncryptionIsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("encryptAppData", m.GetEncryptAppData())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("faceIdBlocked", m.GetFaceIdBlocked())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredPatchVersion", m.GetMinimumRequiredPatchVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumRequiredSdkVersion", m.GetMinimumRequiredSdkVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("minimumWarningPatchVersion", m.GetMinimumWarningPatchVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("screenCaptureBlocked", m.GetScreenCaptureBlocked())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAppDataEncryptionType sets the appDataEncryptionType property value. Type of encryption which should be used for data in a managed app. (iOS Only). Possible values are: useDeviceSettings, afterDeviceRestart, whenDeviceLockedExceptOpenFiles, whenDeviceLocked.
func (m *DefaultManagedAppProtection) SetAppDataEncryptionType(value *ManagedAppDataEncryptionType)() {
    if m != nil {
        m.appDataEncryptionType = value
    }
}
// SetApps sets the apps property value. List of apps to which the policy is deployed.
func (m *DefaultManagedAppProtection) SetApps(value []ManagedMobileAppable)() {
    if m != nil {
        m.apps = value
    }
}
// SetCustomSettings sets the customSettings property value. A set of string key and string value pairs to be sent to the affected users, unalterned by this service
func (m *DefaultManagedAppProtection) SetCustomSettings(value []KeyValuePairable)() {
    if m != nil {
        m.customSettings = value
    }
}
// SetDeployedAppCount sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *DefaultManagedAppProtection) SetDeployedAppCount(value *int32)() {
    if m != nil {
        m.deployedAppCount = value
    }
}
// SetDeploymentSummary sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *DefaultManagedAppProtection) SetDeploymentSummary(value ManagedAppPolicyDeploymentSummaryable)() {
    if m != nil {
        m.deploymentSummary = value
    }
}
// SetDisableAppEncryptionIfDeviceEncryptionIsEnabled sets the disableAppEncryptionIfDeviceEncryptionIsEnabled property value. When this setting is enabled, app level encryption is disabled if device level encryption is enabled. (Android only)
func (m *DefaultManagedAppProtection) SetDisableAppEncryptionIfDeviceEncryptionIsEnabled(value *bool)() {
    if m != nil {
        m.disableAppEncryptionIfDeviceEncryptionIsEnabled = value
    }
}
// SetEncryptAppData sets the encryptAppData property value. Indicates whether managed-app data should be encrypted. (Android only)
func (m *DefaultManagedAppProtection) SetEncryptAppData(value *bool)() {
    if m != nil {
        m.encryptAppData = value
    }
}
// SetFaceIdBlocked sets the faceIdBlocked property value. Indicates whether use of the FaceID is allowed in place of a pin if PinRequired is set to True. (iOS Only)
func (m *DefaultManagedAppProtection) SetFaceIdBlocked(value *bool)() {
    if m != nil {
        m.faceIdBlocked = value
    }
}
// SetMinimumRequiredPatchVersion sets the minimumRequiredPatchVersion property value. Define the oldest required Android security patch level a user can have to gain secure access to the app. (Android only)
func (m *DefaultManagedAppProtection) SetMinimumRequiredPatchVersion(value *string)() {
    if m != nil {
        m.minimumRequiredPatchVersion = value
    }
}
// SetMinimumRequiredSdkVersion sets the minimumRequiredSdkVersion property value. Versions less than the specified version will block the managed app from accessing company data. (iOS Only)
func (m *DefaultManagedAppProtection) SetMinimumRequiredSdkVersion(value *string)() {
    if m != nil {
        m.minimumRequiredSdkVersion = value
    }
}
// SetMinimumWarningPatchVersion sets the minimumWarningPatchVersion property value. Define the oldest recommended Android security patch level a user can have for secure access to the app. (Android only)
func (m *DefaultManagedAppProtection) SetMinimumWarningPatchVersion(value *string)() {
    if m != nil {
        m.minimumWarningPatchVersion = value
    }
}
// SetScreenCaptureBlocked sets the screenCaptureBlocked property value. Indicates whether screen capture is blocked. (Android only)
func (m *DefaultManagedAppProtection) SetScreenCaptureBlocked(value *bool)() {
    if m != nil {
        m.screenCaptureBlocked = value
    }
}
