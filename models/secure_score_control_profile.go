package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e "github.com/microsoft/kiota-abstractions-go/store"
)

type SecureScoreControlProfile struct {
    Entity
}
// SecureScoreControlProfile_SecureScoreControlProfile_maxScore composed type wrapper for classes float64, ReferenceNumeric, string
type SecureScoreControlProfile_SecureScoreControlProfile_maxScore struct {
    // Stores model information.
    backingStore ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore
}
// NewSecureScoreControlProfile_SecureScoreControlProfile_maxScore instantiates a new SecureScoreControlProfile_SecureScoreControlProfile_maxScore and sets the default values.
func NewSecureScoreControlProfile_SecureScoreControlProfile_maxScore()(*SecureScoreControlProfile_SecureScoreControlProfile_maxScore) {
    m := &SecureScoreControlProfile_SecureScoreControlProfile_maxScore{
    }
    m.backingStore = ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStoreFactoryInstance();
    return m
}
// CreateSecureScoreControlProfile_SecureScoreControlProfile_maxScoreFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecureScoreControlProfile_SecureScoreControlProfile_maxScoreFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    result := NewSecureScoreControlProfile_SecureScoreControlProfile_maxScore()
    if parseNode != nil {
        mappingValueNode, err := parseNode.GetChildNode("")
        if err != nil {
            return nil, err
        }
        if mappingValueNode != nil {
            mappingValue, err := mappingValueNode.GetStringValue()
            if err != nil {
                return nil, err
            }
            if mappingValue != nil {
            }
        }
    }
    if val, err := parseNode.GetEnumValue(ParseReferenceNumeric); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetReferenceNumeric(val)
    } else if val, err := parseNode.GetFloat64Value(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetDouble(val)
    } else if val, err := parseNode.GetStringValue(); val != nil {
        if err != nil {
            return nil, err
        }
        result.SetString(val)
    }
    return result, nil
}
// GetBackingStore gets the BackingStore property value. Stores model information.
// returns a BackingStore when successful
func (m *SecureScoreControlProfile_SecureScoreControlProfile_maxScore) GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore) {
    return m.backingStore
}
// GetDouble gets the double property value. Composed type representation for type float64
// returns a *float64 when successful
func (m *SecureScoreControlProfile_SecureScoreControlProfile_maxScore) GetDouble()(*float64) {
    val, err := m.GetBackingStore().Get("double")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*float64)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecureScoreControlProfile_SecureScoreControlProfile_maxScore) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    return make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
}
// GetIsComposedType determines if the current object is a wrapper around a composed type
// returns a bool when successful
func (m *SecureScoreControlProfile_SecureScoreControlProfile_maxScore) GetIsComposedType()(bool) {
    return true
}
// GetReferenceNumeric gets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
// returns a *ReferenceNumeric when successful
func (m *SecureScoreControlProfile_SecureScoreControlProfile_maxScore) GetReferenceNumeric()(*ReferenceNumeric) {
    val, err := m.GetBackingStore().Get("referenceNumeric")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*ReferenceNumeric)
    }
    return nil
}
// GetString gets the string property value. Composed type representation for type string
// returns a *string when successful
func (m *SecureScoreControlProfile_SecureScoreControlProfile_maxScore) GetString()(*string) {
    val, err := m.GetBackingStore().Get("string")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SecureScoreControlProfile_SecureScoreControlProfile_maxScore) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetReferenceNumeric() != nil {
        cast := (*m.GetReferenceNumeric()).String()
        err := writer.WriteStringValue("", &cast)
        if err != nil {
            return err
        }
    } else if m.GetDouble() != nil {
        err := writer.WriteFloat64Value("", m.GetDouble())
        if err != nil {
            return err
        }
    } else if m.GetString() != nil {
        err := writer.WriteStringValue("", m.GetString())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetBackingStore sets the BackingStore property value. Stores model information.
func (m *SecureScoreControlProfile_SecureScoreControlProfile_maxScore) SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)() {
    m.backingStore = value
}
// SetDouble sets the double property value. Composed type representation for type float64
func (m *SecureScoreControlProfile_SecureScoreControlProfile_maxScore) SetDouble(value *float64)() {
    err := m.GetBackingStore().Set("double", value)
    if err != nil {
        panic(err)
    }
}
// SetReferenceNumeric sets the ReferenceNumeric property value. Composed type representation for type ReferenceNumeric
func (m *SecureScoreControlProfile_SecureScoreControlProfile_maxScore) SetReferenceNumeric(value *ReferenceNumeric)() {
    err := m.GetBackingStore().Set("referenceNumeric", value)
    if err != nil {
        panic(err)
    }
}
// SetString sets the string property value. Composed type representation for type string
func (m *SecureScoreControlProfile_SecureScoreControlProfile_maxScore) SetString(value *string)() {
    err := m.GetBackingStore().Set("string", value)
    if err != nil {
        panic(err)
    }
}
type SecureScoreControlProfile_SecureScoreControlProfile_maxScoreable interface {
    ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackedModel
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetBackingStore()(ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)
    GetDouble()(*float64)
    GetReferenceNumeric()(*ReferenceNumeric)
    GetString()(*string)
    SetBackingStore(value ie8677ce2c7e1b4c22e9c3827ecd078d41185424dd9eeb92b7d971ed2d49a392e.BackingStore)()
    SetDouble(value *float64)()
    SetReferenceNumeric(value *ReferenceNumeric)()
    SetString(value *string)()
}
// NewSecureScoreControlProfile instantiates a new SecureScoreControlProfile and sets the default values.
func NewSecureScoreControlProfile()(*SecureScoreControlProfile) {
    m := &SecureScoreControlProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecureScoreControlProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateSecureScoreControlProfileFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecureScoreControlProfile(), nil
}
// GetActionType gets the actionType property value. Control action type (Config, Review, Behavior).
// returns a *string when successful
func (m *SecureScoreControlProfile) GetActionType()(*string) {
    val, err := m.GetBackingStore().Get("actionType")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetActionUrl gets the actionUrl property value. URL to where the control can be actioned.
// returns a *string when successful
func (m *SecureScoreControlProfile) GetActionUrl()(*string) {
    val, err := m.GetBackingStore().Get("actionUrl")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetAzureTenantId gets the azureTenantId property value. GUID string for tenant ID.
// returns a *string when successful
func (m *SecureScoreControlProfile) GetAzureTenantId()(*string) {
    val, err := m.GetBackingStore().Get("azureTenantId")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetComplianceInformation gets the complianceInformation property value. The collection of compliance information associated with secure score control
// returns a []ComplianceInformationable when successful
func (m *SecureScoreControlProfile) GetComplianceInformation()([]ComplianceInformationable) {
    val, err := m.GetBackingStore().Get("complianceInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]ComplianceInformationable)
    }
    return nil
}
// GetControlCategory gets the controlCategory property value. Control action category (Identity, Data, Device, Apps, Infrastructure).
// returns a *string when successful
func (m *SecureScoreControlProfile) GetControlCategory()(*string) {
    val, err := m.GetBackingStore().Get("controlCategory")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetControlStateUpdates gets the controlStateUpdates property value. Flag to indicate where the tenant has marked a control (ignored, thirdParty, reviewed) (supports update).
// returns a []SecureScoreControlStateUpdateable when successful
func (m *SecureScoreControlProfile) GetControlStateUpdates()([]SecureScoreControlStateUpdateable) {
    val, err := m.GetBackingStore().Get("controlStateUpdates")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]SecureScoreControlStateUpdateable)
    }
    return nil
}
// GetDeprecated gets the deprecated property value. Flag to indicate if a control is depreciated.
// returns a *bool when successful
func (m *SecureScoreControlProfile) GetDeprecated()(*bool) {
    val, err := m.GetBackingStore().Get("deprecated")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*bool)
    }
    return nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *SecureScoreControlProfile) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionType"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionType(val)
        }
        return nil
    }
    res["actionUrl"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionUrl(val)
        }
        return nil
    }
    res["azureTenantId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
    res["complianceInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComplianceInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceInformationable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(ComplianceInformationable)
                }
            }
            m.SetComplianceInformation(res)
        }
        return nil
    }
    res["controlCategory"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControlCategory(val)
        }
        return nil
    }
    res["controlStateUpdates"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoreControlStateUpdateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoreControlStateUpdateable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(SecureScoreControlStateUpdateable)
                }
            }
            m.SetControlStateUpdates(res)
        }
        return nil
    }
    res["deprecated"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeprecated(val)
        }
        return nil
    }
    res["implementationCost"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImplementationCost(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["maxScore"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecureScoreControlProfile_SecureScoreControlProfile_maxScoreFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxScore(val.(*SecureScoreControlProfile_SecureScoreControlProfile_maxScore))
        }
        return nil
    }
    res["rank"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRank(val)
        }
        return nil
    }
    res["remediation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediation(val)
        }
        return nil
    }
    res["remediationImpact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationImpact(val)
        }
        return nil
    }
    res["service"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    res["threats"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = *(v.(*string))
                }
            }
            m.SetThreats(res)
        }
        return nil
    }
    res["tier"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTier(val)
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["userImpact"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserImpact(val)
        }
        return nil
    }
    res["vendorInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateSecurityVendorInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVendorInformation(val.(SecurityVendorInformationable))
        }
        return nil
    }
    return res
}
// GetImplementationCost gets the implementationCost property value. Resource cost of implemmentating control (low, moderate, high).
// returns a *string when successful
func (m *SecureScoreControlProfile) GetImplementationCost()(*string) {
    val, err := m.GetBackingStore().Get("implementationCost")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Time at which the control profile entity was last modified. The Timestamp type represents date and time
// returns a *Time when successful
func (m *SecureScoreControlProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    val, err := m.GetBackingStore().Get("lastModifiedDateTime")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    }
    return nil
}
// GetMaxScore gets the maxScore property value. max attainable score for the control.
// returns a SecureScoreControlProfile_SecureScoreControlProfile_maxScoreable when successful
func (m *SecureScoreControlProfile) GetMaxScore()(SecureScoreControlProfile_SecureScoreControlProfile_maxScoreable) {
    val, err := m.GetBackingStore().Get("maxScore")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SecureScoreControlProfile_SecureScoreControlProfile_maxScoreable)
    }
    return nil
}
// GetRank gets the rank property value. Microsoft's stack ranking of control.
// returns a *int32 when successful
func (m *SecureScoreControlProfile) GetRank()(*int32) {
    val, err := m.GetBackingStore().Get("rank")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*int32)
    }
    return nil
}
// GetRemediation gets the remediation property value. Description of what the control will help remediate.
// returns a *string when successful
func (m *SecureScoreControlProfile) GetRemediation()(*string) {
    val, err := m.GetBackingStore().Get("remediation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetRemediationImpact gets the remediationImpact property value. Description of the impact on users of the remediation.
// returns a *string when successful
func (m *SecureScoreControlProfile) GetRemediationImpact()(*string) {
    val, err := m.GetBackingStore().Get("remediationImpact")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetService gets the service property value. Service that owns the control (Exchange, Sharepoint, Microsoft Entra ID).
// returns a *string when successful
func (m *SecureScoreControlProfile) GetService()(*string) {
    val, err := m.GetBackingStore().Get("service")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetThreats gets the threats property value. List of threats the control mitigates (accountBreach, dataDeletion, dataExfiltration, dataSpillage, elevationOfPrivilege, maliciousInsider, passwordCracking, phishingOrWhaling, spoofing).
// returns a []string when successful
func (m *SecureScoreControlProfile) GetThreats()([]string) {
    val, err := m.GetBackingStore().Get("threats")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]string)
    }
    return nil
}
// GetTier gets the tier property value. Control tier (Core, Defense in Depth, Advanced.)
// returns a *string when successful
func (m *SecureScoreControlProfile) GetTier()(*string) {
    val, err := m.GetBackingStore().Get("tier")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetTitle gets the title property value. Title of the control.
// returns a *string when successful
func (m *SecureScoreControlProfile) GetTitle()(*string) {
    val, err := m.GetBackingStore().Get("title")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetUserImpact gets the userImpact property value. User impact of implementing control (low, moderate, high).
// returns a *string when successful
func (m *SecureScoreControlProfile) GetUserImpact()(*string) {
    val, err := m.GetBackingStore().Get("userImpact")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(*string)
    }
    return nil
}
// GetVendorInformation gets the vendorInformation property value. Complex type containing details about the security product/service vendor, provider, and subprovider (for example, vendor=Microsoft; provider=SecureScore). Required.
// returns a SecurityVendorInformationable when successful
func (m *SecureScoreControlProfile) GetVendorInformation()(SecurityVendorInformationable) {
    val, err := m.GetBackingStore().Get("vendorInformation")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.(SecurityVendorInformationable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *SecureScoreControlProfile) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("actionType", m.GetActionType())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("actionUrl", m.GetActionUrl())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("azureTenantId", m.GetAzureTenantId())
        if err != nil {
            return err
        }
    }
    if m.GetComplianceInformation() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetComplianceInformation()))
        for i, v := range m.GetComplianceInformation() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("complianceInformation", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("controlCategory", m.GetControlCategory())
        if err != nil {
            return err
        }
    }
    if m.GetControlStateUpdates() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetControlStateUpdates()))
        for i, v := range m.GetControlStateUpdates() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("controlStateUpdates", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("deprecated", m.GetDeprecated())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("implementationCost", m.GetImplementationCost())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("lastModifiedDateTime", m.GetLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("maxScore", m.GetMaxScore())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("rank", m.GetRank())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remediation", m.GetRemediation())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("remediationImpact", m.GetRemediationImpact())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("service", m.GetService())
        if err != nil {
            return err
        }
    }
    if m.GetThreats() != nil {
        err = writer.WriteCollectionOfStringValues("threats", m.GetThreats())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tier", m.GetTier())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("title", m.GetTitle())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("userImpact", m.GetUserImpact())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("vendorInformation", m.GetVendorInformation())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetActionType sets the actionType property value. Control action type (Config, Review, Behavior).
func (m *SecureScoreControlProfile) SetActionType(value *string)() {
    err := m.GetBackingStore().Set("actionType", value)
    if err != nil {
        panic(err)
    }
}
// SetActionUrl sets the actionUrl property value. URL to where the control can be actioned.
func (m *SecureScoreControlProfile) SetActionUrl(value *string)() {
    err := m.GetBackingStore().Set("actionUrl", value)
    if err != nil {
        panic(err)
    }
}
// SetAzureTenantId sets the azureTenantId property value. GUID string for tenant ID.
func (m *SecureScoreControlProfile) SetAzureTenantId(value *string)() {
    err := m.GetBackingStore().Set("azureTenantId", value)
    if err != nil {
        panic(err)
    }
}
// SetComplianceInformation sets the complianceInformation property value. The collection of compliance information associated with secure score control
func (m *SecureScoreControlProfile) SetComplianceInformation(value []ComplianceInformationable)() {
    err := m.GetBackingStore().Set("complianceInformation", value)
    if err != nil {
        panic(err)
    }
}
// SetControlCategory sets the controlCategory property value. Control action category (Identity, Data, Device, Apps, Infrastructure).
func (m *SecureScoreControlProfile) SetControlCategory(value *string)() {
    err := m.GetBackingStore().Set("controlCategory", value)
    if err != nil {
        panic(err)
    }
}
// SetControlStateUpdates sets the controlStateUpdates property value. Flag to indicate where the tenant has marked a control (ignored, thirdParty, reviewed) (supports update).
func (m *SecureScoreControlProfile) SetControlStateUpdates(value []SecureScoreControlStateUpdateable)() {
    err := m.GetBackingStore().Set("controlStateUpdates", value)
    if err != nil {
        panic(err)
    }
}
// SetDeprecated sets the deprecated property value. Flag to indicate if a control is depreciated.
func (m *SecureScoreControlProfile) SetDeprecated(value *bool)() {
    err := m.GetBackingStore().Set("deprecated", value)
    if err != nil {
        panic(err)
    }
}
// SetImplementationCost sets the implementationCost property value. Resource cost of implemmentating control (low, moderate, high).
func (m *SecureScoreControlProfile) SetImplementationCost(value *string)() {
    err := m.GetBackingStore().Set("implementationCost", value)
    if err != nil {
        panic(err)
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Time at which the control profile entity was last modified. The Timestamp type represents date and time
func (m *SecureScoreControlProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    err := m.GetBackingStore().Set("lastModifiedDateTime", value)
    if err != nil {
        panic(err)
    }
}
// SetMaxScore sets the maxScore property value. max attainable score for the control.
func (m *SecureScoreControlProfile) SetMaxScore(value SecureScoreControlProfile_SecureScoreControlProfile_maxScoreable)() {
    err := m.GetBackingStore().Set("maxScore", value)
    if err != nil {
        panic(err)
    }
}
// SetRank sets the rank property value. Microsoft's stack ranking of control.
func (m *SecureScoreControlProfile) SetRank(value *int32)() {
    err := m.GetBackingStore().Set("rank", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediation sets the remediation property value. Description of what the control will help remediate.
func (m *SecureScoreControlProfile) SetRemediation(value *string)() {
    err := m.GetBackingStore().Set("remediation", value)
    if err != nil {
        panic(err)
    }
}
// SetRemediationImpact sets the remediationImpact property value. Description of the impact on users of the remediation.
func (m *SecureScoreControlProfile) SetRemediationImpact(value *string)() {
    err := m.GetBackingStore().Set("remediationImpact", value)
    if err != nil {
        panic(err)
    }
}
// SetService sets the service property value. Service that owns the control (Exchange, Sharepoint, Microsoft Entra ID).
func (m *SecureScoreControlProfile) SetService(value *string)() {
    err := m.GetBackingStore().Set("service", value)
    if err != nil {
        panic(err)
    }
}
// SetThreats sets the threats property value. List of threats the control mitigates (accountBreach, dataDeletion, dataExfiltration, dataSpillage, elevationOfPrivilege, maliciousInsider, passwordCracking, phishingOrWhaling, spoofing).
func (m *SecureScoreControlProfile) SetThreats(value []string)() {
    err := m.GetBackingStore().Set("threats", value)
    if err != nil {
        panic(err)
    }
}
// SetTier sets the tier property value. Control tier (Core, Defense in Depth, Advanced.)
func (m *SecureScoreControlProfile) SetTier(value *string)() {
    err := m.GetBackingStore().Set("tier", value)
    if err != nil {
        panic(err)
    }
}
// SetTitle sets the title property value. Title of the control.
func (m *SecureScoreControlProfile) SetTitle(value *string)() {
    err := m.GetBackingStore().Set("title", value)
    if err != nil {
        panic(err)
    }
}
// SetUserImpact sets the userImpact property value. User impact of implementing control (low, moderate, high).
func (m *SecureScoreControlProfile) SetUserImpact(value *string)() {
    err := m.GetBackingStore().Set("userImpact", value)
    if err != nil {
        panic(err)
    }
}
// SetVendorInformation sets the vendorInformation property value. Complex type containing details about the security product/service vendor, provider, and subprovider (for example, vendor=Microsoft; provider=SecureScore). Required.
func (m *SecureScoreControlProfile) SetVendorInformation(value SecurityVendorInformationable)() {
    err := m.GetBackingStore().Set("vendorInformation", value)
    if err != nil {
        panic(err)
    }
}
type SecureScoreControlProfileable interface {
    Entityable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetActionType()(*string)
    GetActionUrl()(*string)
    GetAzureTenantId()(*string)
    GetComplianceInformation()([]ComplianceInformationable)
    GetControlCategory()(*string)
    GetControlStateUpdates()([]SecureScoreControlStateUpdateable)
    GetDeprecated()(*bool)
    GetImplementationCost()(*string)
    GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)
    GetMaxScore()(SecureScoreControlProfile_SecureScoreControlProfile_maxScoreable)
    GetRank()(*int32)
    GetRemediation()(*string)
    GetRemediationImpact()(*string)
    GetService()(*string)
    GetThreats()([]string)
    GetTier()(*string)
    GetTitle()(*string)
    GetUserImpact()(*string)
    GetVendorInformation()(SecurityVendorInformationable)
    SetActionType(value *string)()
    SetActionUrl(value *string)()
    SetAzureTenantId(value *string)()
    SetComplianceInformation(value []ComplianceInformationable)()
    SetControlCategory(value *string)()
    SetControlStateUpdates(value []SecureScoreControlStateUpdateable)()
    SetDeprecated(value *bool)()
    SetImplementationCost(value *string)()
    SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)()
    SetMaxScore(value SecureScoreControlProfile_SecureScoreControlProfile_maxScoreable)()
    SetRank(value *int32)()
    SetRemediation(value *string)()
    SetRemediationImpact(value *string)()
    SetService(value *string)()
    SetThreats(value []string)()
    SetTier(value *string)()
    SetTitle(value *string)()
    SetUserImpact(value *string)()
    SetVendorInformation(value SecurityVendorInformationable)()
}
