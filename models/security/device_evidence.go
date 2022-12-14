package security

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// DeviceEvidence 
type DeviceEvidence struct {
    AlertEvidence
    // The azureAdDeviceId property
    azureAdDeviceId *string
    // The defenderAvStatus property
    defenderAvStatus *DefenderAvStatus
    // The deviceDnsName property
    deviceDnsName *string
    // The firstSeenDateTime property
    firstSeenDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The healthStatus property
    healthStatus *DeviceHealthStatus
    // The loggedOnUsers property
    loggedOnUsers []LoggedOnUserable
    // The mdeDeviceId property
    mdeDeviceId *string
    // The onboardingStatus property
    onboardingStatus *OnboardingStatus
    // The osBuild property
    osBuild *int64
    // The osPlatform property
    osPlatform *string
    // The rbacGroupId property
    rbacGroupId *int32
    // The rbacGroupName property
    rbacGroupName *string
    // The riskScore property
    riskScore *DeviceRiskScore
    // The version property
    version *string
    // The vmMetadata property
    vmMetadata VmMetadataable
}
// NewDeviceEvidence instantiates a new DeviceEvidence and sets the default values.
func NewDeviceEvidence()(*DeviceEvidence) {
    m := &DeviceEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateDeviceEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateDeviceEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDeviceEvidence(), nil
}
// GetAzureAdDeviceId gets the azureAdDeviceId property value. The azureAdDeviceId property
func (m *DeviceEvidence) GetAzureAdDeviceId()(*string) {
    return m.azureAdDeviceId
}
// GetDefenderAvStatus gets the defenderAvStatus property value. The defenderAvStatus property
func (m *DeviceEvidence) GetDefenderAvStatus()(*DefenderAvStatus) {
    return m.defenderAvStatus
}
// GetDeviceDnsName gets the deviceDnsName property value. The deviceDnsName property
func (m *DeviceEvidence) GetDeviceDnsName()(*string) {
    return m.deviceDnsName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *DeviceEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["azureAdDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureAdDeviceId(val)
        }
        return nil
    }
    res["defenderAvStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDefenderAvStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDefenderAvStatus(val.(*DefenderAvStatus))
        }
        return nil
    }
    res["deviceDnsName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeviceDnsName(val)
        }
        return nil
    }
    res["firstSeenDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFirstSeenDateTime(val)
        }
        return nil
    }
    res["healthStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceHealthStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHealthStatus(val.(*DeviceHealthStatus))
        }
        return nil
    }
    res["loggedOnUsers"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLoggedOnUserFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LoggedOnUserable, len(val))
            for i, v := range val {
                res[i] = v.(LoggedOnUserable)
            }
            m.SetLoggedOnUsers(res)
        }
        return nil
    }
    res["mdeDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMdeDeviceId(val)
        }
        return nil
    }
    res["onboardingStatus"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseOnboardingStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOnboardingStatus(val.(*OnboardingStatus))
        }
        return nil
    }
    res["osBuild"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsBuild(val)
        }
        return nil
    }
    res["osPlatform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOsPlatform(val)
        }
        return nil
    }
    res["rbacGroupId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRbacGroupId(val)
        }
        return nil
    }
    res["rbacGroupName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRbacGroupName(val)
        }
        return nil
    }
    res["riskScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseDeviceRiskScore)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRiskScore(val.(*DeviceRiskScore))
        }
        return nil
    }
    res["version"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersion(val)
        }
        return nil
    }
    res["vmMetadata"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateVmMetadataFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVmMetadata(val.(VmMetadataable))
        }
        return nil
    }
    return res
}
// GetFirstSeenDateTime gets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *DeviceEvidence) GetFirstSeenDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    return m.firstSeenDateTime
}
// GetHealthStatus gets the healthStatus property value. The healthStatus property
func (m *DeviceEvidence) GetHealthStatus()(*DeviceHealthStatus) {
    return m.healthStatus
}
// GetLoggedOnUsers gets the loggedOnUsers property value. The loggedOnUsers property
func (m *DeviceEvidence) GetLoggedOnUsers()([]LoggedOnUserable) {
    return m.loggedOnUsers
}
// GetMdeDeviceId gets the mdeDeviceId property value. The mdeDeviceId property
func (m *DeviceEvidence) GetMdeDeviceId()(*string) {
    return m.mdeDeviceId
}
// GetOnboardingStatus gets the onboardingStatus property value. The onboardingStatus property
func (m *DeviceEvidence) GetOnboardingStatus()(*OnboardingStatus) {
    return m.onboardingStatus
}
// GetOsBuild gets the osBuild property value. The osBuild property
func (m *DeviceEvidence) GetOsBuild()(*int64) {
    return m.osBuild
}
// GetOsPlatform gets the osPlatform property value. The osPlatform property
func (m *DeviceEvidence) GetOsPlatform()(*string) {
    return m.osPlatform
}
// GetRbacGroupId gets the rbacGroupId property value. The rbacGroupId property
func (m *DeviceEvidence) GetRbacGroupId()(*int32) {
    return m.rbacGroupId
}
// GetRbacGroupName gets the rbacGroupName property value. The rbacGroupName property
func (m *DeviceEvidence) GetRbacGroupName()(*string) {
    return m.rbacGroupName
}
// GetRiskScore gets the riskScore property value. The riskScore property
func (m *DeviceEvidence) GetRiskScore()(*DeviceRiskScore) {
    return m.riskScore
}
// GetVersion gets the version property value. The version property
func (m *DeviceEvidence) GetVersion()(*string) {
    return m.version
}
// GetVmMetadata gets the vmMetadata property value. The vmMetadata property
func (m *DeviceEvidence) GetVmMetadata()(VmMetadataable) {
    return m.vmMetadata
}
// Serialize serializes information the current object
func (m *DeviceEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("azureAdDeviceId", m.GetAzureAdDeviceId())
        if err != nil {
            return err
        }
    }
    if m.GetDefenderAvStatus() != nil {
        cast := (*m.GetDefenderAvStatus()).String()
        err = writer.WriteStringValue("defenderAvStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("deviceDnsName", m.GetDeviceDnsName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("firstSeenDateTime", m.GetFirstSeenDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetHealthStatus() != nil {
        cast := (*m.GetHealthStatus()).String()
        err = writer.WriteStringValue("healthStatus", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetLoggedOnUsers() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLoggedOnUsers()))
        for i, v := range m.GetLoggedOnUsers() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("loggedOnUsers", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("mdeDeviceId", m.GetMdeDeviceId())
        if err != nil {
            return err
        }
    }
    if m.GetOnboardingStatus() != nil {
        cast := (*m.GetOnboardingStatus()).String()
        err = writer.WriteStringValue("onboardingStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt64Value("osBuild", m.GetOsBuild())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("osPlatform", m.GetOsPlatform())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("rbacGroupId", m.GetRbacGroupId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("rbacGroupName", m.GetRbacGroupName())
        if err != nil {
            return err
        }
    }
    if m.GetRiskScore() != nil {
        cast := (*m.GetRiskScore()).String()
        err = writer.WriteStringValue("riskScore", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("version", m.GetVersion())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vmMetadata", m.GetVmMetadata())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAzureAdDeviceId sets the azureAdDeviceId property value. The azureAdDeviceId property
func (m *DeviceEvidence) SetAzureAdDeviceId(value *string)() {
    m.azureAdDeviceId = value
}
// SetDefenderAvStatus sets the defenderAvStatus property value. The defenderAvStatus property
func (m *DeviceEvidence) SetDefenderAvStatus(value *DefenderAvStatus)() {
    m.defenderAvStatus = value
}
// SetDeviceDnsName sets the deviceDnsName property value. The deviceDnsName property
func (m *DeviceEvidence) SetDeviceDnsName(value *string)() {
    m.deviceDnsName = value
}
// SetFirstSeenDateTime sets the firstSeenDateTime property value. The firstSeenDateTime property
func (m *DeviceEvidence) SetFirstSeenDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.firstSeenDateTime = value
}
// SetHealthStatus sets the healthStatus property value. The healthStatus property
func (m *DeviceEvidence) SetHealthStatus(value *DeviceHealthStatus)() {
    m.healthStatus = value
}
// SetLoggedOnUsers sets the loggedOnUsers property value. The loggedOnUsers property
func (m *DeviceEvidence) SetLoggedOnUsers(value []LoggedOnUserable)() {
    m.loggedOnUsers = value
}
// SetMdeDeviceId sets the mdeDeviceId property value. The mdeDeviceId property
func (m *DeviceEvidence) SetMdeDeviceId(value *string)() {
    m.mdeDeviceId = value
}
// SetOnboardingStatus sets the onboardingStatus property value. The onboardingStatus property
func (m *DeviceEvidence) SetOnboardingStatus(value *OnboardingStatus)() {
    m.onboardingStatus = value
}
// SetOsBuild sets the osBuild property value. The osBuild property
func (m *DeviceEvidence) SetOsBuild(value *int64)() {
    m.osBuild = value
}
// SetOsPlatform sets the osPlatform property value. The osPlatform property
func (m *DeviceEvidence) SetOsPlatform(value *string)() {
    m.osPlatform = value
}
// SetRbacGroupId sets the rbacGroupId property value. The rbacGroupId property
func (m *DeviceEvidence) SetRbacGroupId(value *int32)() {
    m.rbacGroupId = value
}
// SetRbacGroupName sets the rbacGroupName property value. The rbacGroupName property
func (m *DeviceEvidence) SetRbacGroupName(value *string)() {
    m.rbacGroupName = value
}
// SetRiskScore sets the riskScore property value. The riskScore property
func (m *DeviceEvidence) SetRiskScore(value *DeviceRiskScore)() {
    m.riskScore = value
}
// SetVersion sets the version property value. The version property
func (m *DeviceEvidence) SetVersion(value *string)() {
    m.version = value
}
// SetVmMetadata sets the vmMetadata property value. The vmMetadata property
func (m *DeviceEvidence) SetVmMetadata(value VmMetadataable)() {
    m.vmMetadata = value
}
