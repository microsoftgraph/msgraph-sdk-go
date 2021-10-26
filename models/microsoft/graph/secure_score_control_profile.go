package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SecureScoreControlProfile struct {
    Entity
    // Control action type (Config, Review, Behavior).
    actionType *string;
    // URL to where the control can be actioned.
    actionUrl *string;
    // GUID string for tenant ID.
    azureTenantId *string;
    // The collection of compliance information associated with secure score control
    complianceInformation []ComplianceInformation;
    // Control action category (Identity, Data, Device, Apps, Infrastructure).
    controlCategory *string;
    // Flag to indicate where the tenant has marked a control (ignore, thirdParty, reviewed) (supports update).
    controlStateUpdates []SecureScoreControlStateUpdate;
    // Flag to indicate if a control is depreciated.
    deprecated *bool;
    // Resource cost of implemmentating control (low, moderate, high).
    implementationCost *string;
    // Time at which the control profile entity was last modified. The Timestamp type represents date and time
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // max attainable score for the control.
    maxScore *float64;
    // Microsoft's stack ranking of control.
    rank *int32;
    // Description of what the control will help remediate.
    remediation *string;
    // Description of the impact on users of the remediation.
    remediationImpact *string;
    // Service that owns the control (Exchange, Sharepoint, Azure AD).
    service *string;
    // List of threats the control mitigates (accountBreach,dataDeletion,dataExfiltration,dataSpillage,
    threats []string;
    // Control tier (Core, Defense in Depth, Advanced.)
    tier *string;
    // Title of the control.
    title *string;
    // User impact of implementing control (low, moderate, high).
    userImpact *string;
    // 
    vendorInformation *SecurityVendorInformation;
}
// Instantiates a new secureScoreControlProfile and sets the default values.
func NewSecureScoreControlProfile()(*SecureScoreControlProfile) {
    m := &SecureScoreControlProfile{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the actionType property value. Control action type (Config, Review, Behavior).
func (m *SecureScoreControlProfile) GetActionType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionType
    }
}
// Gets the actionUrl property value. URL to where the control can be actioned.
func (m *SecureScoreControlProfile) GetActionUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionUrl
    }
}
// Gets the azureTenantId property value. GUID string for tenant ID.
func (m *SecureScoreControlProfile) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// Gets the complianceInformation property value. The collection of compliance information associated with secure score control
func (m *SecureScoreControlProfile) GetComplianceInformation()([]ComplianceInformation) {
    if m == nil {
        return nil
    } else {
        return m.complianceInformation
    }
}
// Gets the controlCategory property value. Control action category (Identity, Data, Device, Apps, Infrastructure).
func (m *SecureScoreControlProfile) GetControlCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.controlCategory
    }
}
// Gets the controlStateUpdates property value. Flag to indicate where the tenant has marked a control (ignore, thirdParty, reviewed) (supports update).
func (m *SecureScoreControlProfile) GetControlStateUpdates()([]SecureScoreControlStateUpdate) {
    if m == nil {
        return nil
    } else {
        return m.controlStateUpdates
    }
}
// Gets the deprecated property value. Flag to indicate if a control is depreciated.
func (m *SecureScoreControlProfile) GetDeprecated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deprecated
    }
}
// Gets the implementationCost property value. Resource cost of implemmentating control (low, moderate, high).
func (m *SecureScoreControlProfile) GetImplementationCost()(*string) {
    if m == nil {
        return nil
    } else {
        return m.implementationCost
    }
}
// Gets the lastModifiedDateTime property value. Time at which the control profile entity was last modified. The Timestamp type represents date and time
func (m *SecureScoreControlProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the maxScore property value. max attainable score for the control.
func (m *SecureScoreControlProfile) GetMaxScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maxScore
    }
}
// Gets the rank property value. Microsoft's stack ranking of control.
func (m *SecureScoreControlProfile) GetRank()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rank
    }
}
// Gets the remediation property value. Description of what the control will help remediate.
func (m *SecureScoreControlProfile) GetRemediation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediation
    }
}
// Gets the remediationImpact property value. Description of the impact on users of the remediation.
func (m *SecureScoreControlProfile) GetRemediationImpact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediationImpact
    }
}
// Gets the service property value. Service that owns the control (Exchange, Sharepoint, Azure AD).
func (m *SecureScoreControlProfile) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
// Gets the threats property value. List of threats the control mitigates (accountBreach,dataDeletion,dataExfiltration,dataSpillage,
func (m *SecureScoreControlProfile) GetThreats()([]string) {
    if m == nil {
        return nil
    } else {
        return m.threats
    }
}
// Gets the tier property value. Control tier (Core, Defense in Depth, Advanced.)
func (m *SecureScoreControlProfile) GetTier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tier
    }
}
// Gets the title property value. Title of the control.
func (m *SecureScoreControlProfile) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// Gets the userImpact property value. User impact of implementing control (low, moderate, high).
func (m *SecureScoreControlProfile) GetUserImpact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userImpact
    }
}
// Gets the vendorInformation property value. 
func (m *SecureScoreControlProfile) GetVendorInformation()(*SecurityVendorInformation) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
// The deserialization information for the current model
func (m *SecureScoreControlProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActionType(val)
        return nil
    }
    res["actionUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActionUrl(val)
        return nil
    }
    res["azureTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAzureTenantId(val)
        return nil
    }
    res["complianceInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewComplianceInformation() })
        if err != nil {
            return err
        }
        res := make([]ComplianceInformation, len(val))
        for i, v := range val {
            res[i] = *(v.(*ComplianceInformation))
        }
        m.SetComplianceInformation(res)
        return nil
    }
    res["controlCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetControlCategory(val)
        return nil
    }
    res["controlStateUpdates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecureScoreControlStateUpdate() })
        if err != nil {
            return err
        }
        res := make([]SecureScoreControlStateUpdate, len(val))
        for i, v := range val {
            res[i] = *(v.(*SecureScoreControlStateUpdate))
        }
        m.SetControlStateUpdates(res)
        return nil
    }
    res["deprecated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetDeprecated(val)
        return nil
    }
    res["implementationCost"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetImplementationCost(val)
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetLastModifiedDateTime(val)
        return nil
    }
    res["maxScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        m.SetMaxScore(val)
        return nil
    }
    res["rank"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetRank(val)
        return nil
    }
    res["remediation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRemediation(val)
        return nil
    }
    res["remediationImpact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRemediationImpact(val)
        return nil
    }
    res["service"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetService(val)
        return nil
    }
    res["threats"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetThreats(res)
        return nil
    }
    res["tier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTier(val)
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTitle(val)
        return nil
    }
    res["userImpact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetUserImpact(val)
        return nil
    }
    res["vendorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSecurityVendorInformation() })
        if err != nil {
            return err
        }
        m.SetVendorInformation(val.(*SecurityVendorInformation))
        return nil
    }
    return res
}
func (m *SecureScoreControlProfile) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SecureScoreControlProfile) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetComplianceInformation()))
        for i, v := range m.GetComplianceInformation() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetControlStateUpdates()))
        for i, v := range m.GetControlStateUpdates() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
        err = writer.WriteFloat64Value("maxScore", m.GetMaxScore())
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
    {
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
// Sets the actionType property value. Control action type (Config, Review, Behavior).
// Parameters:
//  - value : Value to set for the actionType property.
func (m *SecureScoreControlProfile) SetActionType(value *string)() {
    m.actionType = value
}
// Sets the actionUrl property value. URL to where the control can be actioned.
// Parameters:
//  - value : Value to set for the actionUrl property.
func (m *SecureScoreControlProfile) SetActionUrl(value *string)() {
    m.actionUrl = value
}
// Sets the azureTenantId property value. GUID string for tenant ID.
// Parameters:
//  - value : Value to set for the azureTenantId property.
func (m *SecureScoreControlProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
// Sets the complianceInformation property value. The collection of compliance information associated with secure score control
// Parameters:
//  - value : Value to set for the complianceInformation property.
func (m *SecureScoreControlProfile) SetComplianceInformation(value []ComplianceInformation)() {
    m.complianceInformation = value
}
// Sets the controlCategory property value. Control action category (Identity, Data, Device, Apps, Infrastructure).
// Parameters:
//  - value : Value to set for the controlCategory property.
func (m *SecureScoreControlProfile) SetControlCategory(value *string)() {
    m.controlCategory = value
}
// Sets the controlStateUpdates property value. Flag to indicate where the tenant has marked a control (ignore, thirdParty, reviewed) (supports update).
// Parameters:
//  - value : Value to set for the controlStateUpdates property.
func (m *SecureScoreControlProfile) SetControlStateUpdates(value []SecureScoreControlStateUpdate)() {
    m.controlStateUpdates = value
}
// Sets the deprecated property value. Flag to indicate if a control is depreciated.
// Parameters:
//  - value : Value to set for the deprecated property.
func (m *SecureScoreControlProfile) SetDeprecated(value *bool)() {
    m.deprecated = value
}
// Sets the implementationCost property value. Resource cost of implemmentating control (low, moderate, high).
// Parameters:
//  - value : Value to set for the implementationCost property.
func (m *SecureScoreControlProfile) SetImplementationCost(value *string)() {
    m.implementationCost = value
}
// Sets the lastModifiedDateTime property value. Time at which the control profile entity was last modified. The Timestamp type represents date and time
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *SecureScoreControlProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the maxScore property value. max attainable score for the control.
// Parameters:
//  - value : Value to set for the maxScore property.
func (m *SecureScoreControlProfile) SetMaxScore(value *float64)() {
    m.maxScore = value
}
// Sets the rank property value. Microsoft's stack ranking of control.
// Parameters:
//  - value : Value to set for the rank property.
func (m *SecureScoreControlProfile) SetRank(value *int32)() {
    m.rank = value
}
// Sets the remediation property value. Description of what the control will help remediate.
// Parameters:
//  - value : Value to set for the remediation property.
func (m *SecureScoreControlProfile) SetRemediation(value *string)() {
    m.remediation = value
}
// Sets the remediationImpact property value. Description of the impact on users of the remediation.
// Parameters:
//  - value : Value to set for the remediationImpact property.
func (m *SecureScoreControlProfile) SetRemediationImpact(value *string)() {
    m.remediationImpact = value
}
// Sets the service property value. Service that owns the control (Exchange, Sharepoint, Azure AD).
// Parameters:
//  - value : Value to set for the service property.
func (m *SecureScoreControlProfile) SetService(value *string)() {
    m.service = value
}
// Sets the threats property value. List of threats the control mitigates (accountBreach,dataDeletion,dataExfiltration,dataSpillage,
// Parameters:
//  - value : Value to set for the threats property.
func (m *SecureScoreControlProfile) SetThreats(value []string)() {
    m.threats = value
}
// Sets the tier property value. Control tier (Core, Defense in Depth, Advanced.)
// Parameters:
//  - value : Value to set for the tier property.
func (m *SecureScoreControlProfile) SetTier(value *string)() {
    m.tier = value
}
// Sets the title property value. Title of the control.
// Parameters:
//  - value : Value to set for the title property.
func (m *SecureScoreControlProfile) SetTitle(value *string)() {
    m.title = value
}
// Sets the userImpact property value. User impact of implementing control (low, moderate, high).
// Parameters:
//  - value : Value to set for the userImpact property.
func (m *SecureScoreControlProfile) SetUserImpact(value *string)() {
    m.userImpact = value
}
// Sets the vendorInformation property value. 
// Parameters:
//  - value : Value to set for the vendorInformation property.
func (m *SecureScoreControlProfile) SetVendorInformation(value *SecurityVendorInformation)() {
    m.vendorInformation = value
}
