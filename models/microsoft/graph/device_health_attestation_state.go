package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DeviceHealthAttestationState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // TWhen an Attestation Identity Key (AIK) is present on a device, it indicates that the device has an endorsement key (EK) certificate.
    attestationIdentityKey *string;
    // On or Off of BitLocker Drive Encryption
    bitLockerStatus *string;
    // The security version number of the Boot Application
    bootAppSecurityVersion *string;
    // When bootDebugging is enabled, the device is used in development and testing
    bootDebugging *string;
    // The security version number of the Boot Application
    bootManagerSecurityVersion *string;
    // The version of the Boot Manager
    bootManagerVersion *string;
    // The Boot Revision List that was loaded during initial boot on the attested device
    bootRevisionListInfo *string;
    // When code integrity is enabled, code execution is restricted to integrity verified code
    codeIntegrity *string;
    // The version of the Boot Manager
    codeIntegrityCheckVersion *string;
    // The Code Integrity policy that is controlling the security of the boot environment
    codeIntegrityPolicy *string;
    // The DHA report version. (Namespace version)
    contentNamespaceUrl *string;
    // The HealthAttestation state schema version
    contentVersion *string;
    // DEP Policy defines a set of hardware and software technologies that perform additional checks on memory
    dataExcutionPolicy *string;
    // The DHA report version. (Namespace version)
    deviceHealthAttestationStatus *string;
    // ELAM provides protection for the computers in your network when they start up
    earlyLaunchAntiMalwareDriverProtection *string;
    // This attribute indicates if DHA is supported for the device
    healthAttestationSupportedStatus *string;
    // This attribute appears if DHA-Service detects an integrity issue
    healthStatusMismatchInfo *string;
    // The DateTime when device was evaluated or issued to MDM
    issuedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Timestamp of the last update.
    lastUpdateDateTime *string;
    // When operatingSystemKernelDebugging is enabled, the device is used in development and testing
    operatingSystemKernelDebugging *string;
    // The Operating System Revision List that was loaded during initial boot on the attested device
    operatingSystemRevListInfo *string;
    // The measurement that is captured in PCR[0]
    pcr0 *string;
    // Informational attribute that identifies the HASH algorithm that was used by TPM
    pcrHashAlgorithm *string;
    // The number of times a PC device has hibernated or resumed
    resetCount *int64;
    // The number of times a PC device has rebooted
    restartCount *int64;
    // Safe mode is a troubleshooting option for Windows that starts your computer in a limited state
    safeMode *string;
    // When Secure Boot is enabled, the core components must have the correct cryptographic signatures
    secureBoot *string;
    // Fingerprint of the Custom Secure Boot Configuration Policy
    secureBootConfigurationPolicyFingerPrint *string;
    // When test signing is allowed, the device does not enforce signature validation during boot
    testSigning *string;
    // The security version number of the Boot Application
    tpmVersion *string;
    // VSM is a container that protects high value assets from a compromised kernel
    virtualSecureMode *string;
    // Operating system running with limited services that is used to prepare a computer for Windows
    windowsPE *string;
}
// Instantiates a new deviceHealthAttestationState and sets the default values.
func NewDeviceHealthAttestationState()(*DeviceHealthAttestationState) {
    m := &DeviceHealthAttestationState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *DeviceHealthAttestationState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the attestationIdentityKey property value. TWhen an Attestation Identity Key (AIK) is present on a device, it indicates that the device has an endorsement key (EK) certificate.
func (m *DeviceHealthAttestationState) GetAttestationIdentityKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.attestationIdentityKey
    }
}
// Gets the bitLockerStatus property value. On or Off of BitLocker Drive Encryption
func (m *DeviceHealthAttestationState) GetBitLockerStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bitLockerStatus
    }
}
// Gets the bootAppSecurityVersion property value. The security version number of the Boot Application
func (m *DeviceHealthAttestationState) GetBootAppSecurityVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bootAppSecurityVersion
    }
}
// Gets the bootDebugging property value. When bootDebugging is enabled, the device is used in development and testing
func (m *DeviceHealthAttestationState) GetBootDebugging()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bootDebugging
    }
}
// Gets the bootManagerSecurityVersion property value. The security version number of the Boot Application
func (m *DeviceHealthAttestationState) GetBootManagerSecurityVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bootManagerSecurityVersion
    }
}
// Gets the bootManagerVersion property value. The version of the Boot Manager
func (m *DeviceHealthAttestationState) GetBootManagerVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bootManagerVersion
    }
}
// Gets the bootRevisionListInfo property value. The Boot Revision List that was loaded during initial boot on the attested device
func (m *DeviceHealthAttestationState) GetBootRevisionListInfo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.bootRevisionListInfo
    }
}
// Gets the codeIntegrity property value. When code integrity is enabled, code execution is restricted to integrity verified code
func (m *DeviceHealthAttestationState) GetCodeIntegrity()(*string) {
    if m == nil {
        return nil
    } else {
        return m.codeIntegrity
    }
}
// Gets the codeIntegrityCheckVersion property value. The version of the Boot Manager
func (m *DeviceHealthAttestationState) GetCodeIntegrityCheckVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.codeIntegrityCheckVersion
    }
}
// Gets the codeIntegrityPolicy property value. The Code Integrity policy that is controlling the security of the boot environment
func (m *DeviceHealthAttestationState) GetCodeIntegrityPolicy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.codeIntegrityPolicy
    }
}
// Gets the contentNamespaceUrl property value. The DHA report version. (Namespace version)
func (m *DeviceHealthAttestationState) GetContentNamespaceUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentNamespaceUrl
    }
}
// Gets the contentVersion property value. The HealthAttestation state schema version
func (m *DeviceHealthAttestationState) GetContentVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.contentVersion
    }
}
// Gets the dataExcutionPolicy property value. DEP Policy defines a set of hardware and software technologies that perform additional checks on memory
func (m *DeviceHealthAttestationState) GetDataExcutionPolicy()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dataExcutionPolicy
    }
}
// Gets the deviceHealthAttestationStatus property value. The DHA report version. (Namespace version)
func (m *DeviceHealthAttestationState) GetDeviceHealthAttestationStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.deviceHealthAttestationStatus
    }
}
// Gets the earlyLaunchAntiMalwareDriverProtection property value. ELAM provides protection for the computers in your network when they start up
func (m *DeviceHealthAttestationState) GetEarlyLaunchAntiMalwareDriverProtection()(*string) {
    if m == nil {
        return nil
    } else {
        return m.earlyLaunchAntiMalwareDriverProtection
    }
}
// Gets the healthAttestationSupportedStatus property value. This attribute indicates if DHA is supported for the device
func (m *DeviceHealthAttestationState) GetHealthAttestationSupportedStatus()(*string) {
    if m == nil {
        return nil
    } else {
        return m.healthAttestationSupportedStatus
    }
}
// Gets the healthStatusMismatchInfo property value. This attribute appears if DHA-Service detects an integrity issue
func (m *DeviceHealthAttestationState) GetHealthStatusMismatchInfo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.healthStatusMismatchInfo
    }
}
// Gets the issuedDateTime property value. The DateTime when device was evaluated or issued to MDM
func (m *DeviceHealthAttestationState) GetIssuedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.issuedDateTime
    }
}
// Gets the lastUpdateDateTime property value. The Timestamp of the last update.
func (m *DeviceHealthAttestationState) GetLastUpdateDateTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.lastUpdateDateTime
    }
}
// Gets the operatingSystemKernelDebugging property value. When operatingSystemKernelDebugging is enabled, the device is used in development and testing
func (m *DeviceHealthAttestationState) GetOperatingSystemKernelDebugging()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemKernelDebugging
    }
}
// Gets the operatingSystemRevListInfo property value. The Operating System Revision List that was loaded during initial boot on the attested device
func (m *DeviceHealthAttestationState) GetOperatingSystemRevListInfo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operatingSystemRevListInfo
    }
}
// Gets the pcr0 property value. The measurement that is captured in PCR[0]
func (m *DeviceHealthAttestationState) GetPcr0()(*string) {
    if m == nil {
        return nil
    } else {
        return m.pcr0
    }
}
// Gets the pcrHashAlgorithm property value. Informational attribute that identifies the HASH algorithm that was used by TPM
func (m *DeviceHealthAttestationState) GetPcrHashAlgorithm()(*string) {
    if m == nil {
        return nil
    } else {
        return m.pcrHashAlgorithm
    }
}
// Gets the resetCount property value. The number of times a PC device has hibernated or resumed
func (m *DeviceHealthAttestationState) GetResetCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.resetCount
    }
}
// Gets the restartCount property value. The number of times a PC device has rebooted
func (m *DeviceHealthAttestationState) GetRestartCount()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.restartCount
    }
}
// Gets the safeMode property value. Safe mode is a troubleshooting option for Windows that starts your computer in a limited state
func (m *DeviceHealthAttestationState) GetSafeMode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.safeMode
    }
}
// Gets the secureBoot property value. When Secure Boot is enabled, the core components must have the correct cryptographic signatures
func (m *DeviceHealthAttestationState) GetSecureBoot()(*string) {
    if m == nil {
        return nil
    } else {
        return m.secureBoot
    }
}
// Gets the secureBootConfigurationPolicyFingerPrint property value. Fingerprint of the Custom Secure Boot Configuration Policy
func (m *DeviceHealthAttestationState) GetSecureBootConfigurationPolicyFingerPrint()(*string) {
    if m == nil {
        return nil
    } else {
        return m.secureBootConfigurationPolicyFingerPrint
    }
}
// Gets the testSigning property value. When test signing is allowed, the device does not enforce signature validation during boot
func (m *DeviceHealthAttestationState) GetTestSigning()(*string) {
    if m == nil {
        return nil
    } else {
        return m.testSigning
    }
}
// Gets the tpmVersion property value. The security version number of the Boot Application
func (m *DeviceHealthAttestationState) GetTpmVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tpmVersion
    }
}
// Gets the virtualSecureMode property value. VSM is a container that protects high value assets from a compromised kernel
func (m *DeviceHealthAttestationState) GetVirtualSecureMode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.virtualSecureMode
    }
}
// Gets the windowsPE property value. Operating system running with limited services that is used to prepare a computer for Windows
func (m *DeviceHealthAttestationState) GetWindowsPE()(*string) {
    if m == nil {
        return nil
    } else {
        return m.windowsPE
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *DeviceHealthAttestationState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the attestationIdentityKey property value. TWhen an Attestation Identity Key (AIK) is present on a device, it indicates that the device has an endorsement key (EK) certificate.
// Parameters:
//  - value : Value to set for the attestationIdentityKey property.
func (m *DeviceHealthAttestationState) SetAttestationIdentityKey(value *string)() {
    m.attestationIdentityKey = value
}
// Sets the bitLockerStatus property value. On or Off of BitLocker Drive Encryption
// Parameters:
//  - value : Value to set for the bitLockerStatus property.
func (m *DeviceHealthAttestationState) SetBitLockerStatus(value *string)() {
    m.bitLockerStatus = value
}
// Sets the bootAppSecurityVersion property value. The security version number of the Boot Application
// Parameters:
//  - value : Value to set for the bootAppSecurityVersion property.
func (m *DeviceHealthAttestationState) SetBootAppSecurityVersion(value *string)() {
    m.bootAppSecurityVersion = value
}
// Sets the bootDebugging property value. When bootDebugging is enabled, the device is used in development and testing
// Parameters:
//  - value : Value to set for the bootDebugging property.
func (m *DeviceHealthAttestationState) SetBootDebugging(value *string)() {
    m.bootDebugging = value
}
// Sets the bootManagerSecurityVersion property value. The security version number of the Boot Application
// Parameters:
//  - value : Value to set for the bootManagerSecurityVersion property.
func (m *DeviceHealthAttestationState) SetBootManagerSecurityVersion(value *string)() {
    m.bootManagerSecurityVersion = value
}
// Sets the bootManagerVersion property value. The version of the Boot Manager
// Parameters:
//  - value : Value to set for the bootManagerVersion property.
func (m *DeviceHealthAttestationState) SetBootManagerVersion(value *string)() {
    m.bootManagerVersion = value
}
// Sets the bootRevisionListInfo property value. The Boot Revision List that was loaded during initial boot on the attested device
// Parameters:
//  - value : Value to set for the bootRevisionListInfo property.
func (m *DeviceHealthAttestationState) SetBootRevisionListInfo(value *string)() {
    m.bootRevisionListInfo = value
}
// Sets the codeIntegrity property value. When code integrity is enabled, code execution is restricted to integrity verified code
// Parameters:
//  - value : Value to set for the codeIntegrity property.
func (m *DeviceHealthAttestationState) SetCodeIntegrity(value *string)() {
    m.codeIntegrity = value
}
// Sets the codeIntegrityCheckVersion property value. The version of the Boot Manager
// Parameters:
//  - value : Value to set for the codeIntegrityCheckVersion property.
func (m *DeviceHealthAttestationState) SetCodeIntegrityCheckVersion(value *string)() {
    m.codeIntegrityCheckVersion = value
}
// Sets the codeIntegrityPolicy property value. The Code Integrity policy that is controlling the security of the boot environment
// Parameters:
//  - value : Value to set for the codeIntegrityPolicy property.
func (m *DeviceHealthAttestationState) SetCodeIntegrityPolicy(value *string)() {
    m.codeIntegrityPolicy = value
}
// Sets the contentNamespaceUrl property value. The DHA report version. (Namespace version)
// Parameters:
//  - value : Value to set for the contentNamespaceUrl property.
func (m *DeviceHealthAttestationState) SetContentNamespaceUrl(value *string)() {
    m.contentNamespaceUrl = value
}
// Sets the contentVersion property value. The HealthAttestation state schema version
// Parameters:
//  - value : Value to set for the contentVersion property.
func (m *DeviceHealthAttestationState) SetContentVersion(value *string)() {
    m.contentVersion = value
}
// Sets the dataExcutionPolicy property value. DEP Policy defines a set of hardware and software technologies that perform additional checks on memory
// Parameters:
//  - value : Value to set for the dataExcutionPolicy property.
func (m *DeviceHealthAttestationState) SetDataExcutionPolicy(value *string)() {
    m.dataExcutionPolicy = value
}
// Sets the deviceHealthAttestationStatus property value. The DHA report version. (Namespace version)
// Parameters:
//  - value : Value to set for the deviceHealthAttestationStatus property.
func (m *DeviceHealthAttestationState) SetDeviceHealthAttestationStatus(value *string)() {
    m.deviceHealthAttestationStatus = value
}
// Sets the earlyLaunchAntiMalwareDriverProtection property value. ELAM provides protection for the computers in your network when they start up
// Parameters:
//  - value : Value to set for the earlyLaunchAntiMalwareDriverProtection property.
func (m *DeviceHealthAttestationState) SetEarlyLaunchAntiMalwareDriverProtection(value *string)() {
    m.earlyLaunchAntiMalwareDriverProtection = value
}
// Sets the healthAttestationSupportedStatus property value. This attribute indicates if DHA is supported for the device
// Parameters:
//  - value : Value to set for the healthAttestationSupportedStatus property.
func (m *DeviceHealthAttestationState) SetHealthAttestationSupportedStatus(value *string)() {
    m.healthAttestationSupportedStatus = value
}
// Sets the healthStatusMismatchInfo property value. This attribute appears if DHA-Service detects an integrity issue
// Parameters:
//  - value : Value to set for the healthStatusMismatchInfo property.
func (m *DeviceHealthAttestationState) SetHealthStatusMismatchInfo(value *string)() {
    m.healthStatusMismatchInfo = value
}
// Sets the issuedDateTime property value. The DateTime when device was evaluated or issued to MDM
// Parameters:
//  - value : Value to set for the issuedDateTime property.
func (m *DeviceHealthAttestationState) SetIssuedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.issuedDateTime = value
}
// Sets the lastUpdateDateTime property value. The Timestamp of the last update.
// Parameters:
//  - value : Value to set for the lastUpdateDateTime property.
func (m *DeviceHealthAttestationState) SetLastUpdateDateTime(value *string)() {
    m.lastUpdateDateTime = value
}
// Sets the operatingSystemKernelDebugging property value. When operatingSystemKernelDebugging is enabled, the device is used in development and testing
// Parameters:
//  - value : Value to set for the operatingSystemKernelDebugging property.
func (m *DeviceHealthAttestationState) SetOperatingSystemKernelDebugging(value *string)() {
    m.operatingSystemKernelDebugging = value
}
// Sets the operatingSystemRevListInfo property value. The Operating System Revision List that was loaded during initial boot on the attested device
// Parameters:
//  - value : Value to set for the operatingSystemRevListInfo property.
func (m *DeviceHealthAttestationState) SetOperatingSystemRevListInfo(value *string)() {
    m.operatingSystemRevListInfo = value
}
// Sets the pcr0 property value. The measurement that is captured in PCR[0]
// Parameters:
//  - value : Value to set for the pcr0 property.
func (m *DeviceHealthAttestationState) SetPcr0(value *string)() {
    m.pcr0 = value
}
// Sets the pcrHashAlgorithm property value. Informational attribute that identifies the HASH algorithm that was used by TPM
// Parameters:
//  - value : Value to set for the pcrHashAlgorithm property.
func (m *DeviceHealthAttestationState) SetPcrHashAlgorithm(value *string)() {
    m.pcrHashAlgorithm = value
}
// Sets the resetCount property value. The number of times a PC device has hibernated or resumed
// Parameters:
//  - value : Value to set for the resetCount property.
func (m *DeviceHealthAttestationState) SetResetCount(value *int64)() {
    m.resetCount = value
}
// Sets the restartCount property value. The number of times a PC device has rebooted
// Parameters:
//  - value : Value to set for the restartCount property.
func (m *DeviceHealthAttestationState) SetRestartCount(value *int64)() {
    m.restartCount = value
}
// Sets the safeMode property value. Safe mode is a troubleshooting option for Windows that starts your computer in a limited state
// Parameters:
//  - value : Value to set for the safeMode property.
func (m *DeviceHealthAttestationState) SetSafeMode(value *string)() {
    m.safeMode = value
}
// Sets the secureBoot property value. When Secure Boot is enabled, the core components must have the correct cryptographic signatures
// Parameters:
//  - value : Value to set for the secureBoot property.
func (m *DeviceHealthAttestationState) SetSecureBoot(value *string)() {
    m.secureBoot = value
}
// Sets the secureBootConfigurationPolicyFingerPrint property value. Fingerprint of the Custom Secure Boot Configuration Policy
// Parameters:
//  - value : Value to set for the secureBootConfigurationPolicyFingerPrint property.
func (m *DeviceHealthAttestationState) SetSecureBootConfigurationPolicyFingerPrint(value *string)() {
    m.secureBootConfigurationPolicyFingerPrint = value
}
// Sets the testSigning property value. When test signing is allowed, the device does not enforce signature validation during boot
// Parameters:
//  - value : Value to set for the testSigning property.
func (m *DeviceHealthAttestationState) SetTestSigning(value *string)() {
    m.testSigning = value
}
// Sets the tpmVersion property value. The security version number of the Boot Application
// Parameters:
//  - value : Value to set for the tpmVersion property.
func (m *DeviceHealthAttestationState) SetTpmVersion(value *string)() {
    m.tpmVersion = value
}
// Sets the virtualSecureMode property value. VSM is a container that protects high value assets from a compromised kernel
// Parameters:
//  - value : Value to set for the virtualSecureMode property.
func (m *DeviceHealthAttestationState) SetVirtualSecureMode(value *string)() {
    m.virtualSecureMode = value
}
// Sets the windowsPE property value. Operating system running with limited services that is used to prepare a computer for Windows
// Parameters:
//  - value : Value to set for the windowsPE property.
func (m *DeviceHealthAttestationState) SetWindowsPE(value *string)() {
    m.windowsPE = value
}
