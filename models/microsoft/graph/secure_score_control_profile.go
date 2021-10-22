package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type SecureScoreControlProfile struct {
    Entity
    actionType *string;
    actionUrl *string;
    azureTenantId *string;
    complianceInformation []ComplianceInformation;
    controlCategory *string;
    controlStateUpdates []SecureScoreControlStateUpdate;
    deprecated *bool;
    implementationCost *string;
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    maxScore *float64;
    rank *int32;
    remediation *string;
    remediationImpact *string;
    service *string;
    threats []string;
    tier *string;
    title *string;
    userImpact *string;
    vendorInformation *SecurityVendorInformation;
}
func NewSecureScoreControlProfile()(*SecureScoreControlProfile) {
    m := &SecureScoreControlProfile{
        Entity: *NewEntity(),
    }
    return m
}
func (m *SecureScoreControlProfile) GetActionType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionType
    }
}
func (m *SecureScoreControlProfile) GetActionUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.actionUrl
    }
}
func (m *SecureScoreControlProfile) GetAzureTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.azureTenantId
    }
}
func (m *SecureScoreControlProfile) GetComplianceInformation()([]ComplianceInformation) {
    if m == nil {
        return nil
    } else {
        return m.complianceInformation
    }
}
func (m *SecureScoreControlProfile) GetControlCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.controlCategory
    }
}
func (m *SecureScoreControlProfile) GetControlStateUpdates()([]SecureScoreControlStateUpdate) {
    if m == nil {
        return nil
    } else {
        return m.controlStateUpdates
    }
}
func (m *SecureScoreControlProfile) GetDeprecated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.deprecated
    }
}
func (m *SecureScoreControlProfile) GetImplementationCost()(*string) {
    if m == nil {
        return nil
    } else {
        return m.implementationCost
    }
}
func (m *SecureScoreControlProfile) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
func (m *SecureScoreControlProfile) GetMaxScore()(*float64) {
    if m == nil {
        return nil
    } else {
        return m.maxScore
    }
}
func (m *SecureScoreControlProfile) GetRank()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.rank
    }
}
func (m *SecureScoreControlProfile) GetRemediation()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediation
    }
}
func (m *SecureScoreControlProfile) GetRemediationImpact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.remediationImpact
    }
}
func (m *SecureScoreControlProfile) GetService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.service
    }
}
func (m *SecureScoreControlProfile) GetThreats()([]string) {
    if m == nil {
        return nil
    } else {
        return m.threats
    }
}
func (m *SecureScoreControlProfile) GetTier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tier
    }
}
func (m *SecureScoreControlProfile) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
func (m *SecureScoreControlProfile) GetUserImpact()(*string) {
    if m == nil {
        return nil
    } else {
        return m.userImpact
    }
}
func (m *SecureScoreControlProfile) GetVendorInformation()(*SecurityVendorInformation) {
    if m == nil {
        return nil
    } else {
        return m.vendorInformation
    }
}
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
func (m *SecureScoreControlProfile) SetActionType(value *string)() {
    m.actionType = value
}
func (m *SecureScoreControlProfile) SetActionUrl(value *string)() {
    m.actionUrl = value
}
func (m *SecureScoreControlProfile) SetAzureTenantId(value *string)() {
    m.azureTenantId = value
}
func (m *SecureScoreControlProfile) SetComplianceInformation(value []ComplianceInformation)() {
    m.complianceInformation = value
}
func (m *SecureScoreControlProfile) SetControlCategory(value *string)() {
    m.controlCategory = value
}
func (m *SecureScoreControlProfile) SetControlStateUpdates(value []SecureScoreControlStateUpdate)() {
    m.controlStateUpdates = value
}
func (m *SecureScoreControlProfile) SetDeprecated(value *bool)() {
    m.deprecated = value
}
func (m *SecureScoreControlProfile) SetImplementationCost(value *string)() {
    m.implementationCost = value
}
func (m *SecureScoreControlProfile) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
func (m *SecureScoreControlProfile) SetMaxScore(value *float64)() {
    m.maxScore = value
}
func (m *SecureScoreControlProfile) SetRank(value *int32)() {
    m.rank = value
}
func (m *SecureScoreControlProfile) SetRemediation(value *string)() {
    m.remediation = value
}
func (m *SecureScoreControlProfile) SetRemediationImpact(value *string)() {
    m.remediationImpact = value
}
func (m *SecureScoreControlProfile) SetService(value *string)() {
    m.service = value
}
func (m *SecureScoreControlProfile) SetThreats(value []string)() {
    m.threats = value
}
func (m *SecureScoreControlProfile) SetTier(value *string)() {
    m.tier = value
}
func (m *SecureScoreControlProfile) SetTitle(value *string)() {
    m.title = value
}
func (m *SecureScoreControlProfile) SetUserImpact(value *string)() {
    m.userImpact = value
}
func (m *SecureScoreControlProfile) SetVendorInformation(value *SecurityVendorInformation)() {
    m.vendorInformation = value
}
