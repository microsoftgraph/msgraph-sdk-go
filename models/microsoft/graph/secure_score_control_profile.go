package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SecureScoreControlProfile 
type SecureScoreControlProfile struct {
    Entity
    // Control action type (Config, Review, Behavior).
    actionType *string;
    // URL to where the control can be actioned.
    actionUrl *string;
    // GUID string for tenant ID.
    azureTenantId *string;
    // The collection of compliance information associated with secure score control
    complianceInformation []ComplianceInformationable;
    // Control action category (Identity, Data, Device, Apps, Infrastructure).
    controlCategory *string;
    // Flag to indicate where the tenant has marked a control (ignore, thirdParty, reviewed) (supports update).
    controlStateUpdates []SecureScoreControlStateUpdateable;
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
    vendorInformation SecurityVendorInformationable;
}
// NewSecureScoreControlProfile instantiates a new secureScoreControlProfile and sets the default values.
func NewSecureScoreControlProfile()(*SecureScoreControlProfile) {
    m := &SecureScoreControlProfile{
        Entity: *NewEntity(),
    }
    return m
}
// CreateSecureScoreControlProfileFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecureScoreControlProfileFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSecureScoreControlProfile(), nil
}
// GetActionType gets the actionType property value. Control action type (Config, Review, Behavior).
func (m *SecureScoreControlProfile) GetActionType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionType
    }
}
// GetActionUrl gets the actionUrl property value. URL to where the control can be actioned.
func (m *SecureScoreControlProfile) GetActionUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionUrl
    }
}
// GetAzureTenantId gets the azureTenantId property value. GUID string for tenant ID.
func (m *SecureScoreControlProfile) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
// GetComplianceInformation gets the complianceInformation property value. The collection of compliance information associated with secure score control
func (m *SecureScoreControlProfile) GetComplianceInformation()([]ComplianceInformationable) {
    if m == nil {
        return nil
    } else {
        return m.complianceInformation
    }
}
// GetControlCategory gets the controlCategory property value. Control action category (Identity, Data, Device, Apps, Infrastructure).
func (m *SecureScoreControlProfile) GetControlCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.controlCategory
    }
}
// GetControlStateUpdates gets the controlStateUpdates property value. Flag to indicate where the tenant has marked a control (ignore, thirdParty, reviewed) (supports update).
func (m *SecureScoreControlProfile) GetControlStateUpdates()([]SecureScoreControlStateUpdateable) {
    if m == nil {
        return nil
    } else {
        return m.controlStateUpdates
    }
}
// GetDeprecated gets the deprecated property value. Flag to indicate if a control is depreciated.
func (m *SecureScoreControlProfile) GetDeprecated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deprecated
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecureScoreControlProfile) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["actionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionType(val)
        }
        return nil
    }
    res["actionUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActionUrl(val)
        }
        return nil
    }
    res["azureTenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAzureTenantId(val)
        }
        return nil
    }
    res["complianceInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComplianceInformationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceInformationable, len(val))
            for i, v := range val {
                res[i] = v.(ComplianceInformationable)
            }
            m.SetComplianceInformation(res)
        }
        return nil
    }
    res["controlCategory"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetControlCategory(val)
        }
        return nil
    }
    res["controlStateUpdates"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateSecureScoreControlStateUpdateFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SecureScoreControlStateUpdateable, len(val))
            for i, v := range val {
                res[i] = v.(SecureScoreControlStateUpdateable)
            }
            m.SetControlStateUpdates(res)
        }
        return nil
    }
    res["deprecated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDeprecated(val)
        }
        return nil
    }
    res["implementationCost"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImplementationCost(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["maxScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMaxScore(val)
        }
        return nil
    }
    res["rank"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRank(val)
        }
        return nil
    }
    res["remediation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediation(val)
        }
        return nil
    }
    res["remediationImpact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRemediationImpact(val)
        }
        return nil
    }
    res["service"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetService(val)
        }
        return nil
    }
    res["threats"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetThreats(res)
        }
        return nil
    }
    res["tier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTier(val)
        }
        return nil
    }
    res["title"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTitle(val)
        }
        return nil
    }
    res["userImpact"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetUserImpact(val)
        }
        return nil
    }
    res["vendorInformation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
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
func (m *SecureScoreControlProfile) GetImplementationCost()(*string) {
    if m == nil {
        return nil
    } else {
        return m.implementationCost
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. Time at which the control profile entity was last modified. The Timestamp type represents date and time
func (m *SecureScoreControlProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetMaxScore gets the maxScore property value. max attainable score for the control.
func (m *SecureScoreControlProfile) GetMaxScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maxScore
    }
}
// GetRank gets the rank property value. Microsoft's stack ranking of control.
func (m *SecureScoreControlProfile) GetRank()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rank
    }
}
// GetRemediation gets the remediation property value. Description of what the control will help remediate.
func (m *SecureScoreControlProfile) GetRemediation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediation
    }
}
// GetRemediationImpact gets the remediationImpact property value. Description of the impact on users of the remediation.
func (m *SecureScoreControlProfile) GetRemediationImpact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediationImpact
    }
}
// GetService gets the service property value. Service that owns the control (Exchange, Sharepoint, Azure AD).
func (m *SecureScoreControlProfile) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
// GetThreats gets the threats property value. List of threats the control mitigates (accountBreach,dataDeletion,dataExfiltration,dataSpillage,
func (m *SecureScoreControlProfile) GetThreats()([]string) {
    if m == nil {
        return nil
    } else {
        return m.threats
    }
}
// GetTier gets the tier property value. Control tier (Core, Defense in Depth, Advanced.)
func (m *SecureScoreControlProfile) GetTier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tier
    }
}
// GetTitle gets the title property value. Title of the control.
func (m *SecureScoreControlProfile) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// GetUserImpact gets the userImpact property value. User impact of implementing control (low, moderate, high).
func (m *SecureScoreControlProfile) GetUserImpact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userImpact
    }
}
// GetVendorInformation gets the vendorInformation property value. 
func (m *SecureScoreControlProfile) GetVendorInformation()(SecurityVendorInformationable) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
// Serialize serializes information the current object
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
    if m.GetComplianceInformation() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetComplianceInformation()))
        for i, v := range m.GetComplianceInformation() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetControlStateUpdates()))
        for i, v := range m.GetControlStateUpdates() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
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
    if m != nil {
        m.actionType = value
    }
}
// SetActionUrl sets the actionUrl property value. URL to where the control can be actioned.
func (m *SecureScoreControlProfile) SetActionUrl(value *string)() {
    if m != nil {
        m.actionUrl = value
    }
}
// SetAzureTenantId sets the azureTenantId property value. GUID string for tenant ID.
func (m *SecureScoreControlProfile) SetAzureTenantId(value *string)() {
    if m != nil {
        m.azureTenantId = value
    }
}
// SetComplianceInformation sets the complianceInformation property value. The collection of compliance information associated with secure score control
func (m *SecureScoreControlProfile) SetComplianceInformation(value []ComplianceInformationable)() {
    if m != nil {
        m.complianceInformation = value
    }
}
// SetControlCategory sets the controlCategory property value. Control action category (Identity, Data, Device, Apps, Infrastructure).
func (m *SecureScoreControlProfile) SetControlCategory(value *string)() {
    if m != nil {
        m.controlCategory = value
    }
}
// SetControlStateUpdates sets the controlStateUpdates property value. Flag to indicate where the tenant has marked a control (ignore, thirdParty, reviewed) (supports update).
func (m *SecureScoreControlProfile) SetControlStateUpdates(value []SecureScoreControlStateUpdateable)() {
    if m != nil {
        m.controlStateUpdates = value
    }
}
// SetDeprecated sets the deprecated property value. Flag to indicate if a control is depreciated.
func (m *SecureScoreControlProfile) SetDeprecated(value *bool)() {
    if m != nil {
        m.deprecated = value
    }
}
// SetImplementationCost sets the implementationCost property value. Resource cost of implemmentating control (low, moderate, high).
func (m *SecureScoreControlProfile) SetImplementationCost(value *string)() {
    if m != nil {
        m.implementationCost = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. Time at which the control profile entity was last modified. The Timestamp type represents date and time
func (m *SecureScoreControlProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetMaxScore sets the maxScore property value. max attainable score for the control.
func (m *SecureScoreControlProfile) SetMaxScore(value *float64)() {
    if m != nil {
        m.maxScore = value
    }
}
// SetRank sets the rank property value. Microsoft's stack ranking of control.
func (m *SecureScoreControlProfile) SetRank(value *int32)() {
    if m != nil {
        m.rank = value
    }
}
// SetRemediation sets the remediation property value. Description of what the control will help remediate.
func (m *SecureScoreControlProfile) SetRemediation(value *string)() {
    if m != nil {
        m.remediation = value
    }
}
// SetRemediationImpact sets the remediationImpact property value. Description of the impact on users of the remediation.
func (m *SecureScoreControlProfile) SetRemediationImpact(value *string)() {
    if m != nil {
        m.remediationImpact = value
    }
}
// SetService sets the service property value. Service that owns the control (Exchange, Sharepoint, Azure AD).
func (m *SecureScoreControlProfile) SetService(value *string)() {
    if m != nil {
        m.service = value
    }
}
// SetThreats sets the threats property value. List of threats the control mitigates (accountBreach,dataDeletion,dataExfiltration,dataSpillage,
func (m *SecureScoreControlProfile) SetThreats(value []string)() {
    if m != nil {
        m.threats = value
    }
}
// SetTier sets the tier property value. Control tier (Core, Defense in Depth, Advanced.)
func (m *SecureScoreControlProfile) SetTier(value *string)() {
    if m != nil {
        m.tier = value
    }
}
// SetTitle sets the title property value. Title of the control.
func (m *SecureScoreControlProfile) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
// SetUserImpact sets the userImpact property value. User impact of implementing control (low, moderate, high).
func (m *SecureScoreControlProfile) SetUserImpact(value *string)() {
    if m != nil {
        m.userImpact = value
    }
}
// SetVendorInformation sets the vendorInformation property value. 
func (m *SecureScoreControlProfile) SetVendorInformation(value SecurityVendorInformationable)() {
    if m != nil {
        m.vendorInformation = value
    }
}
