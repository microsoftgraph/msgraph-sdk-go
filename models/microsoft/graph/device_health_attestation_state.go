package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DeviceHealthAttestationState struct {
    additionalData map[string]interface{};
    attestationIdentityKey *string;
    bitLockerStatus *string;
    bootAppSecurityVersion *string;
    bootDebugging *string;
    bootManagerSecurityVersion *string;
    bootManagerVersion *string;
    bootRevisionListInfo *string;
    codeIntegrity *string;
    codeIntegrityCheckVersion *string;
    codeIntegrityPolicy *string;
    contentNamespaceUrl *string;
    contentVersion *string;
    dataExcutionPolicy *string;
    deviceHealthAttestationStatus *string;
    earlyLaunchAntiMalwareDriverProtection *string;
    healthAttestationSupportedStatus *string;
    healthStatusMismatchInfo *string;
    issuedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    lastUpdateDateTime *string;
    operatingSystemKernelDebugging *string;
    operatingSystemRevListInfo *string;
    pcr0 *string;
    pcrHashAlgorithm *string;
    resetCount *int64;
    restartCount *int64;
    safeMode *string;
    secureBoot *string;
    secureBootConfigurationPolicyFingerPrint *string;
    testSigning *string;
    tpmVersion *string;
    virtualSecureMode *string;
    windowsPE *string;
}
func NewDeviceHealthAttestationState()(*DeviceHealthAttestationState) {
    m := &DeviceHealthAttestationState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *DeviceHealthAttestationState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *DeviceHealthAttestationState) GetAttestationIdentityKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.attestationIdentityKey
    }
}
func (m *DeviceHealthAttestationState) GetBitLockerStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bitLockerStatus
    }
}
func (m *DeviceHealthAttestationState) GetBootAppSecurityVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bootAppSecurityVersion
    }
}
func (m *DeviceHealthAttestationState) GetBootDebugging()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bootDebugging
    }
}
func (m *DeviceHealthAttestationState) GetBootManagerSecurityVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bootManagerSecurityVersion
    }
}
func (m *DeviceHealthAttestationState) GetBootManagerVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bootManagerVersion
    }
}
func (m *DeviceHealthAttestationState) GetBootRevisionListInfo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bootRevisionListInfo
    }
}
func (m *DeviceHealthAttestationState) GetCodeIntegrity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.codeIntegrity
    }
}
func (m *DeviceHealthAttestationState) GetCodeIntegrityCheckVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.codeIntegrityCheckVersion
    }
}
func (m *DeviceHealthAttestationState) GetCodeIntegrityPolicy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.codeIntegrityPolicy
    }
}
func (m *DeviceHealthAttestationState) GetContentNamespaceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentNamespaceUrl
    }
}
func (m *DeviceHealthAttestationState) GetContentVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentVersion
    }
}
func (m *DeviceHealthAttestationState) GetDataExcutionPolicy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataExcutionPolicy
    }
}
func (m *DeviceHealthAttestationState) GetDeviceHealthAttestationStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceHealthAttestationStatus
    }
}
func (m *DeviceHealthAttestationState) GetEarlyLaunchAntiMalwareDriverProtection()(*string) {
    if m == nil {
        return nil
    } else {
        return m.earlyLaunchAntiMalwareDriverProtection
    }
}
func (m *DeviceHealthAttestationState) GetHealthAttestationSupportedStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.healthAttestationSupportedStatus
    }
}
func (m *DeviceHealthAttestationState) GetHealthStatusMismatchInfo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.healthStatusMismatchInfo
    }
}
func (m *DeviceHealthAttestationState) GetIssuedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.issuedDateTime
    }
}
func (m *DeviceHealthAttestationState) GetLastUpdateDateTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdateDateTime
    }
}
func (m *DeviceHealthAttestationState) GetOperatingSystemKernelDebugging()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemKernelDebugging
    }
}
func (m *DeviceHealthAttestationState) GetOperatingSystemRevListInfo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemRevListInfo
    }
}
func (m *DeviceHealthAttestationState) GetPcr0()(*string) {
    if m == nil {
        return nil
    } else {
        return m.pcr0
    }
}
func (m *DeviceHealthAttestationState) GetPcrHashAlgorithm()(*string) {
    if m == nil {
        return nil
    } else {
        return m.pcrHashAlgorithm
    }
}
func (m *DeviceHealthAttestationState) GetResetCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.resetCount
    }
}
func (m *DeviceHealthAttestationState) GetRestartCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.restartCount
    }
}
func (m *DeviceHealthAttestationState) GetSafeMode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.safeMode
    }
}
func (m *DeviceHealthAttestationState) GetSecureBoot()(*string) {
    if m == nil {
        return nil
    } else {
        return m.secureBoot
    }
}
func (m *DeviceHealthAttestationState) GetSecureBootConfigurationPolicyFingerPrint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.secureBootConfigurationPolicyFingerPrint
    }
}
func (m *DeviceHealthAttestationState) GetTestSigning()(*string) {
    if m == nil {
        return nil
    } else {
        return m.testSigning
    }
}
func (m *DeviceHealthAttestationState) GetTpmVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tpmVersion
    }
}
func (m *DeviceHealthAttestationState) GetVirtualSecureMode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.virtualSecureMode
    }
}
func (m *DeviceHealthAttestationState) GetWindowsPE()(*string) {
    if m == nil {
        return nil
    } else {
        return m.windowsPE
    }
}
func (m *DeviceHealthAttestationState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["attestationIdentityKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAttestationIdentityKey(val)
        return nil
    }
    res["bitLockerStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBitLockerStatus(val)
        return nil
    }
    res["bootAppSecurityVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBootAppSecurityVersion(val)
        return nil
    }
    res["bootDebugging"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBootDebugging(val)
        return nil
    }
    res["bootManagerSecurityVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBootManagerSecurityVersion(val)
        return nil
    }
    res["bootManagerVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBootManagerVersion(val)
        return nil
    }
    res["bootRevisionListInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetBootRevisionListInfo(val)
        return nil
    }
    res["codeIntegrity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCodeIntegrity(val)
        return nil
    }
    res["codeIntegrityCheckVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCodeIntegrityCheckVersion(val)
        return nil
    }
    res["codeIntegrityPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCodeIntegrityPolicy(val)
        return nil
    }
    res["contentNamespaceUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentNamespaceUrl(val)
        return nil
    }
    res["contentVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetContentVersion(val)
        return nil
    }
    res["dataExcutionPolicy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDataExcutionPolicy(val)
        return nil
    }
    res["deviceHealthAttestationStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDeviceHealthAttestationStatus(val)
        return nil
    }
    res["earlyLaunchAntiMalwareDriverProtection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEarlyLaunchAntiMalwareDriverProtection(val)
        return nil
    }
    res["healthAttestationSupportedStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHealthAttestationSupportedStatus(val)
        return nil
    }
    res["healthStatusMismatchInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHealthStatusMismatchInfo(val)
        return nil
    }
    res["issuedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetIssuedDateTime(val)
        return nil
    }
    res["lastUpdateDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLastUpdateDateTime(val)
        return nil
    }
    res["operatingSystemKernelDebugging"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperatingSystemKernelDebugging(val)
        return nil
    }
    res["operatingSystemRevListInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperatingSystemRevListInfo(val)
        return nil
    }
    res["pcr0"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPcr0(val)
        return nil
    }
    res["pcrHashAlgorithm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPcrHashAlgorithm(val)
        return nil
    }
    res["resetCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetResetCount(val)
        return nil
    }
    res["restartCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        m.SetRestartCount(val)
        return nil
    }
    res["safeMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSafeMode(val)
        return nil
    }
    res["secureBoot"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSecureBoot(val)
        return nil
    }
    res["secureBootConfigurationPolicyFingerPrint"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSecureBootConfigurationPolicyFingerPrint(val)
        return nil
    }
    res["testSigning"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTestSigning(val)
        return nil
    }
    res["tpmVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTpmVersion(val)
        return nil
    }
    res["virtualSecureMode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetVirtualSecureMode(val)
        return nil
    }
    res["windowsPE"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetWindowsPE(val)
        return nil
    }
    return res
}
func (m *DeviceHealthAttestationState) IsNil()(bool) {
    return m == nil
}
func (m *DeviceHealthAttestationState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("attestationIdentityKey", m.GetAttestationIdentityKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bitLockerStatus", m.GetBitLockerStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bootAppSecurityVersion", m.GetBootAppSecurityVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bootDebugging", m.GetBootDebugging())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bootManagerSecurityVersion", m.GetBootManagerSecurityVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bootManagerVersion", m.GetBootManagerVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("bootRevisionListInfo", m.GetBootRevisionListInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("codeIntegrity", m.GetCodeIntegrity())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("codeIntegrityCheckVersion", m.GetCodeIntegrityCheckVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("codeIntegrityPolicy", m.GetCodeIntegrityPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contentNamespaceUrl", m.GetContentNamespaceUrl())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("contentVersion", m.GetContentVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dataExcutionPolicy", m.GetDataExcutionPolicy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("deviceHealthAttestationStatus", m.GetDeviceHealthAttestationStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("earlyLaunchAntiMalwareDriverProtection", m.GetEarlyLaunchAntiMalwareDriverProtection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("healthAttestationSupportedStatus", m.GetHealthAttestationSupportedStatus())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("healthStatusMismatchInfo", m.GetHealthStatusMismatchInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("issuedDateTime", m.GetIssuedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("lastUpdateDateTime", m.GetLastUpdateDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operatingSystemKernelDebugging", m.GetOperatingSystemKernelDebugging())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("operatingSystemRevListInfo", m.GetOperatingSystemRevListInfo())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pcr0", m.GetPcr0())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("pcrHashAlgorithm", m.GetPcrHashAlgorithm())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("resetCount", m.GetResetCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("restartCount", m.GetRestartCount())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("safeMode", m.GetSafeMode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secureBoot", m.GetSecureBoot())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("secureBootConfigurationPolicyFingerPrint", m.GetSecureBootConfigurationPolicyFingerPrint())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("testSigning", m.GetTestSigning())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("tpmVersion", m.GetTpmVersion())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("virtualSecureMode", m.GetVirtualSecureMode())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("windowsPE", m.GetWindowsPE())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *DeviceHealthAttestationState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *DeviceHealthAttestationState) SetAttestationIdentityKey(value *string)() {
    m.attestationIdentityKey = value
}
func (m *DeviceHealthAttestationState) SetBitLockerStatus(value *string)() {
    m.bitLockerStatus = value
}
func (m *DeviceHealthAttestationState) SetBootAppSecurityVersion(value *string)() {
    m.bootAppSecurityVersion = value
}
func (m *DeviceHealthAttestationState) SetBootDebugging(value *string)() {
    m.bootDebugging = value
}
func (m *DeviceHealthAttestationState) SetBootManagerSecurityVersion(value *string)() {
    m.bootManagerSecurityVersion = value
}
func (m *DeviceHealthAttestationState) SetBootManagerVersion(value *string)() {
    m.bootManagerVersion = value
}
func (m *DeviceHealthAttestationState) SetBootRevisionListInfo(value *string)() {
    m.bootRevisionListInfo = value
}
func (m *DeviceHealthAttestationState) SetCodeIntegrity(value *string)() {
    m.codeIntegrity = value
}
func (m *DeviceHealthAttestationState) SetCodeIntegrityCheckVersion(value *string)() {
    m.codeIntegrityCheckVersion = value
}
func (m *DeviceHealthAttestationState) SetCodeIntegrityPolicy(value *string)() {
    m.codeIntegrityPolicy = value
}
func (m *DeviceHealthAttestationState) SetContentNamespaceUrl(value *string)() {
    m.contentNamespaceUrl = value
}
func (m *DeviceHealthAttestationState) SetContentVersion(value *string)() {
    m.contentVersion = value
}
func (m *DeviceHealthAttestationState) SetDataExcutionPolicy(value *string)() {
    m.dataExcutionPolicy = value
}
func (m *DeviceHealthAttestationState) SetDeviceHealthAttestationStatus(value *string)() {
    m.deviceHealthAttestationStatus = value
}
func (m *DeviceHealthAttestationState) SetEarlyLaunchAntiMalwareDriverProtection(value *string)() {
    m.earlyLaunchAntiMalwareDriverProtection = value
}
func (m *DeviceHealthAttestationState) SetHealthAttestationSupportedStatus(value *string)() {
    m.healthAttestationSupportedStatus = value
}
func (m *DeviceHealthAttestationState) SetHealthStatusMismatchInfo(value *string)() {
    m.healthStatusMismatchInfo = value
}
func (m *DeviceHealthAttestationState) SetIssuedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.issuedDateTime = value
}
func (m *DeviceHealthAttestationState) SetLastUpdateDateTime(value *string)() {
    m.lastUpdateDateTime = value
}
func (m *DeviceHealthAttestationState) SetOperatingSystemKernelDebugging(value *string)() {
    m.operatingSystemKernelDebugging = value
}
func (m *DeviceHealthAttestationState) SetOperatingSystemRevListInfo(value *string)() {
    m.operatingSystemRevListInfo = value
}
func (m *DeviceHealthAttestationState) SetPcr0(value *string)() {
    m.pcr0 = value
}
func (m *DeviceHealthAttestationState) SetPcrHashAlgorithm(value *string)() {
    m.pcrHashAlgorithm = value
}
func (m *DeviceHealthAttestationState) SetResetCount(value *int64)() {
    m.resetCount = value
}
func (m *DeviceHealthAttestationState) SetRestartCount(value *int64)() {
    m.restartCount = value
}
func (m *DeviceHealthAttestationState) SetSafeMode(value *string)() {
    m.safeMode = value
}
func (m *DeviceHealthAttestationState) SetSecureBoot(value *string)() {
    m.secureBoot = value
}
func (m *DeviceHealthAttestationState) SetSecureBootConfigurationPolicyFingerPrint(value *string)() {
    m.secureBootConfigurationPolicyFingerPrint = value
}
func (m *DeviceHealthAttestationState) SetTestSigning(value *string)() {
    m.testSigning = value
}
func (m *DeviceHealthAttestationState) SetTpmVersion(value *string)() {
    m.tpmVersion = value
}
func (m *DeviceHealthAttestationState) SetVirtualSecureMode(value *string)() {
    m.virtualSecureMode = value
}
func (m *DeviceHealthAttestationState) SetWindowsPE(value *string)() {
    m.windowsPE = value
}
