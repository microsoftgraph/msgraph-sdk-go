package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ProvisioningObjectSummary struct {
    Entity
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    changeId *string;
    cycleId *string;
    durationInMilliseconds *int32;
    initiatedBy *Initiator;
    jobId *string;
    modifiedProperties []ModifiedProperty;
    provisioningAction *ProvisioningAction;
    provisioningStatusInfo *ProvisioningStatusInfo;
    provisioningSteps []ProvisioningStep;
    servicePrincipal *ProvisioningServicePrincipal;
    sourceIdentity *ProvisionedIdentity;
    sourceSystem *ProvisioningSystem;
    targetIdentity *ProvisionedIdentity;
    targetSystem *ProvisioningSystem;
    tenantId *string;
}
func NewProvisioningObjectSummary()(*ProvisioningObjectSummary) {
    m := &ProvisioningObjectSummary{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ProvisioningObjectSummary) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
func (m *ProvisioningObjectSummary) GetChangeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeId
    }
}
func (m *ProvisioningObjectSummary) GetCycleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cycleId
    }
}
func (m *ProvisioningObjectSummary) GetDurationInMilliseconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInMilliseconds
    }
}
func (m *ProvisioningObjectSummary) GetInitiatedBy()(*Initiator) {
    if m == nil {
        return nil
    } else {
        return m.initiatedBy
    }
}
func (m *ProvisioningObjectSummary) GetJobId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jobId
    }
}
func (m *ProvisioningObjectSummary) GetModifiedProperties()([]ModifiedProperty) {
    if m == nil {
        return nil
    } else {
        return m.modifiedProperties
    }
}
func (m *ProvisioningObjectSummary) GetProvisioningAction()(*ProvisioningAction) {
    if m == nil {
        return nil
    } else {
        return m.provisioningAction
    }
}
func (m *ProvisioningObjectSummary) GetProvisioningStatusInfo()(*ProvisioningStatusInfo) {
    if m == nil {
        return nil
    } else {
        return m.provisioningStatusInfo
    }
}
func (m *ProvisioningObjectSummary) GetProvisioningSteps()([]ProvisioningStep) {
    if m == nil {
        return nil
    } else {
        return m.provisioningSteps
    }
}
func (m *ProvisioningObjectSummary) GetServicePrincipal()(*ProvisioningServicePrincipal) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipal
    }
}
func (m *ProvisioningObjectSummary) GetSourceIdentity()(*ProvisionedIdentity) {
    if m == nil {
        return nil
    } else {
        return m.sourceIdentity
    }
}
func (m *ProvisioningObjectSummary) GetSourceSystem()(*ProvisioningSystem) {
    if m == nil {
        return nil
    } else {
        return m.sourceSystem
    }
}
func (m *ProvisioningObjectSummary) GetTargetIdentity()(*ProvisionedIdentity) {
    if m == nil {
        return nil
    } else {
        return m.targetIdentity
    }
}
func (m *ProvisioningObjectSummary) GetTargetSystem()(*ProvisioningSystem) {
    if m == nil {
        return nil
    } else {
        return m.targetSystem
    }
}
func (m *ProvisioningObjectSummary) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
func (m *ProvisioningObjectSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetActivityDateTime(val)
        return nil
    }
    res["changeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetChangeId(val)
        return nil
    }
    res["cycleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCycleId(val)
        return nil
    }
    res["durationInMilliseconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDurationInMilliseconds(val)
        return nil
    }
    res["initiatedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInitiator() })
        if err != nil {
            return err
        }
        m.SetInitiatedBy(val.(*Initiator))
        return nil
    }
    res["jobId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetJobId(val)
        return nil
    }
    res["modifiedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewModifiedProperty() })
        if err != nil {
            return err
        }
        res := make([]ModifiedProperty, len(val))
        for i, v := range val {
            res[i] = *(v.(*ModifiedProperty))
        }
        m.SetModifiedProperties(res)
        return nil
    }
    res["provisioningAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningAction)
        if err != nil {
            return err
        }
        cast := val.(ProvisioningAction)
        m.SetProvisioningAction(&cast)
        return nil
    }
    res["provisioningStatusInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisioningStatusInfo() })
        if err != nil {
            return err
        }
        m.SetProvisioningStatusInfo(val.(*ProvisioningStatusInfo))
        return nil
    }
    res["provisioningSteps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisioningStep() })
        if err != nil {
            return err
        }
        res := make([]ProvisioningStep, len(val))
        for i, v := range val {
            res[i] = *(v.(*ProvisioningStep))
        }
        m.SetProvisioningSteps(res)
        return nil
    }
    res["servicePrincipal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisioningServicePrincipal() })
        if err != nil {
            return err
        }
        m.SetServicePrincipal(val.(*ProvisioningServicePrincipal))
        return nil
    }
    res["sourceIdentity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisionedIdentity() })
        if err != nil {
            return err
        }
        m.SetSourceIdentity(val.(*ProvisionedIdentity))
        return nil
    }
    res["sourceSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisioningSystem() })
        if err != nil {
            return err
        }
        m.SetSourceSystem(val.(*ProvisioningSystem))
        return nil
    }
    res["targetIdentity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisionedIdentity() })
        if err != nil {
            return err
        }
        m.SetTargetIdentity(val.(*ProvisionedIdentity))
        return nil
    }
    res["targetSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisioningSystem() })
        if err != nil {
            return err
        }
        m.SetTargetSystem(val.(*ProvisioningSystem))
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetTenantId(val)
        return nil
    }
    return res
}
func (m *ProvisioningObjectSummary) IsNil()(bool) {
    return m == nil
}
func (m *ProvisioningObjectSummary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("activityDateTime", m.GetActivityDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("changeId", m.GetChangeId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("cycleId", m.GetCycleId())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteInt32Value("durationInMilliseconds", m.GetDurationInMilliseconds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("initiatedBy", m.GetInitiatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("jobId", m.GetJobId())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetModifiedProperties()))
        for i, v := range m.GetModifiedProperties() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("modifiedProperties", cast)
        if err != nil {
            return err
        }
    }
    if m.GetProvisioningAction() != nil {
        cast := m.GetProvisioningAction().String()
        err = writer.WriteStringValue("provisioningAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("provisioningStatusInfo", m.GetProvisioningStatusInfo())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetProvisioningSteps()))
        for i, v := range m.GetProvisioningSteps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("provisioningSteps", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("servicePrincipal", m.GetServicePrincipal())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceIdentity", m.GetSourceIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("sourceSystem", m.GetSourceSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targetIdentity", m.GetTargetIdentity())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("targetSystem", m.GetTargetSystem())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("tenantId", m.GetTenantId())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *ProvisioningObjectSummary) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
func (m *ProvisioningObjectSummary) SetChangeId(value *string)() {
    m.changeId = value
}
func (m *ProvisioningObjectSummary) SetCycleId(value *string)() {
    m.cycleId = value
}
func (m *ProvisioningObjectSummary) SetDurationInMilliseconds(value *int32)() {
    m.durationInMilliseconds = value
}
func (m *ProvisioningObjectSummary) SetInitiatedBy(value *Initiator)() {
    m.initiatedBy = value
}
func (m *ProvisioningObjectSummary) SetJobId(value *string)() {
    m.jobId = value
}
func (m *ProvisioningObjectSummary) SetModifiedProperties(value []ModifiedProperty)() {
    m.modifiedProperties = value
}
func (m *ProvisioningObjectSummary) SetProvisioningAction(value *ProvisioningAction)() {
    m.provisioningAction = value
}
func (m *ProvisioningObjectSummary) SetProvisioningStatusInfo(value *ProvisioningStatusInfo)() {
    m.provisioningStatusInfo = value
}
func (m *ProvisioningObjectSummary) SetProvisioningSteps(value []ProvisioningStep)() {
    m.provisioningSteps = value
}
func (m *ProvisioningObjectSummary) SetServicePrincipal(value *ProvisioningServicePrincipal)() {
    m.servicePrincipal = value
}
func (m *ProvisioningObjectSummary) SetSourceIdentity(value *ProvisionedIdentity)() {
    m.sourceIdentity = value
}
func (m *ProvisioningObjectSummary) SetSourceSystem(value *ProvisioningSystem)() {
    m.sourceSystem = value
}
func (m *ProvisioningObjectSummary) SetTargetIdentity(value *ProvisionedIdentity)() {
    m.targetIdentity = value
}
func (m *ProvisioningObjectSummary) SetTargetSystem(value *ProvisioningSystem)() {
    m.targetSystem = value
}
func (m *ProvisioningObjectSummary) SetTenantId(value *string)() {
    m.tenantId = value
}
