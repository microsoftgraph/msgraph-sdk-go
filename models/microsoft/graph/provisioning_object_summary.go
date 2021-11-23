package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// provisioningObjectSummary 
type ProvisioningObjectSummary struct {
    Entity
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Unique ID of this change in this cycle.
    changeId *string;
    // Unique ID per job iteration.
    cycleId *string;
    // Indicates how long this provisioning action took to finish. Measured in milliseconds.
    durationInMilliseconds *int32;
    // Details of who initiated this provisioning.
    initiatedBy *Initiator;
    // The unique ID for the whole provisioning job.
    jobId *string;
    // Details of each property that was modified in this provisioning action on this object.
    modifiedProperties []ModifiedProperty;
    // Indicates the activity name or the operation name. Possible values are: create, update, delete, stageddelete, disable, other and unknownFutureValue. For a list of activities logged, refer to Azure AD activity list.
    provisioningAction *ProvisioningAction;
    // Details of provisioning status.
    provisioningStatusInfo *ProvisioningStatusInfo;
    // Details of each step in provisioning.
    provisioningSteps []ProvisioningStep;
    // Represents the service principal used for provisioning.
    servicePrincipal *ProvisioningServicePrincipal;
    // Details of source object being provisioned.
    sourceIdentity *ProvisionedIdentity;
    // Details of source system of the object being provisioned.
    sourceSystem *ProvisioningSystem;
    // Details of target object being provisioned.
    targetIdentity *ProvisionedIdentity;
    // Details of target system of the object being provisioned.
    targetSystem *ProvisioningSystem;
    // Unique Azure AD tenant ID.
    tenantId *string;
}
// NewProvisioningObjectSummary instantiates a new provisioningObjectSummary and sets the default values.
func NewProvisioningObjectSummary()(*ProvisioningObjectSummary) {
    m := &ProvisioningObjectSummary{
        Entity: *NewEntity(),
    }
    return m
}
// GetActivityDateTime gets the activityDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ProvisioningObjectSummary) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
// GetChangeId gets the changeId property value. Unique ID of this change in this cycle.
func (m *ProvisioningObjectSummary) GetChangeId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.changeId
    }
}
// GetCycleId gets the cycleId property value. Unique ID per job iteration.
func (m *ProvisioningObjectSummary) GetCycleId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.cycleId
    }
}
// GetDurationInMilliseconds gets the durationInMilliseconds property value. Indicates how long this provisioning action took to finish. Measured in milliseconds.
func (m *ProvisioningObjectSummary) GetDurationInMilliseconds()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.durationInMilliseconds
    }
}
// GetInitiatedBy gets the initiatedBy property value. Details of who initiated this provisioning.
func (m *ProvisioningObjectSummary) GetInitiatedBy()(*Initiator) {
    if m == nil {
        return nil
    } else {
        return m.initiatedBy
    }
}
// GetJobId gets the jobId property value. The unique ID for the whole provisioning job.
func (m *ProvisioningObjectSummary) GetJobId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.jobId
    }
}
// GetModifiedProperties gets the modifiedProperties property value. Details of each property that was modified in this provisioning action on this object.
func (m *ProvisioningObjectSummary) GetModifiedProperties()([]ModifiedProperty) {
    if m == nil {
        return nil
    } else {
        return m.modifiedProperties
    }
}
// GetProvisioningAction gets the provisioningAction property value. Indicates the activity name or the operation name. Possible values are: create, update, delete, stageddelete, disable, other and unknownFutureValue. For a list of activities logged, refer to Azure AD activity list.
func (m *ProvisioningObjectSummary) GetProvisioningAction()(*ProvisioningAction) {
    if m == nil {
        return nil
    } else {
        return m.provisioningAction
    }
}
// GetProvisioningStatusInfo gets the provisioningStatusInfo property value. Details of provisioning status.
func (m *ProvisioningObjectSummary) GetProvisioningStatusInfo()(*ProvisioningStatusInfo) {
    if m == nil {
        return nil
    } else {
        return m.provisioningStatusInfo
    }
}
// GetProvisioningSteps gets the provisioningSteps property value. Details of each step in provisioning.
func (m *ProvisioningObjectSummary) GetProvisioningSteps()([]ProvisioningStep) {
    if m == nil {
        return nil
    } else {
        return m.provisioningSteps
    }
}
// GetServicePrincipal gets the servicePrincipal property value. Represents the service principal used for provisioning.
func (m *ProvisioningObjectSummary) GetServicePrincipal()(*ProvisioningServicePrincipal) {
    if m == nil {
        return nil
    } else {
        return m.servicePrincipal
    }
}
// GetSourceIdentity gets the sourceIdentity property value. Details of source object being provisioned.
func (m *ProvisioningObjectSummary) GetSourceIdentity()(*ProvisionedIdentity) {
    if m == nil {
        return nil
    } else {
        return m.sourceIdentity
    }
}
// GetSourceSystem gets the sourceSystem property value. Details of source system of the object being provisioned.
func (m *ProvisioningObjectSummary) GetSourceSystem()(*ProvisioningSystem) {
    if m == nil {
        return nil
    } else {
        return m.sourceSystem
    }
}
// GetTargetIdentity gets the targetIdentity property value. Details of target object being provisioned.
func (m *ProvisioningObjectSummary) GetTargetIdentity()(*ProvisionedIdentity) {
    if m == nil {
        return nil
    } else {
        return m.targetIdentity
    }
}
// GetTargetSystem gets the targetSystem property value. Details of target system of the object being provisioned.
func (m *ProvisioningObjectSummary) GetTargetSystem()(*ProvisioningSystem) {
    if m == nil {
        return nil
    } else {
        return m.targetSystem
    }
}
// GetTenantId gets the tenantId property value. Unique Azure AD tenant ID.
func (m *ProvisioningObjectSummary) GetTenantId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.tenantId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningObjectSummary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetActivityDateTime(val)
        }
        return nil
    }
    res["changeId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangeId(val)
        }
        return nil
    }
    res["cycleId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCycleId(val)
        }
        return nil
    }
    res["durationInMilliseconds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationInMilliseconds(val)
        }
        return nil
    }
    res["initiatedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewInitiator() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetInitiatedBy(val.(*Initiator))
        }
        return nil
    }
    res["jobId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetJobId(val)
        }
        return nil
    }
    res["modifiedProperties"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewModifiedProperty() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ModifiedProperty, len(val))
            for i, v := range val {
                res[i] = *(v.(*ModifiedProperty))
            }
            m.SetModifiedProperties(res)
        }
        return nil
    }
    res["provisioningAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningAction)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(ProvisioningAction)
            m.SetProvisioningAction(&cast)
        }
        return nil
    }
    res["provisioningStatusInfo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisioningStatusInfo() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisioningStatusInfo(val.(*ProvisioningStatusInfo))
        }
        return nil
    }
    res["provisioningSteps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisioningStep() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ProvisioningStep, len(val))
            for i, v := range val {
                res[i] = *(v.(*ProvisioningStep))
            }
            m.SetProvisioningSteps(res)
        }
        return nil
    }
    res["servicePrincipal"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisioningServicePrincipal() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServicePrincipal(val.(*ProvisioningServicePrincipal))
        }
        return nil
    }
    res["sourceIdentity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisionedIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceIdentity(val.(*ProvisionedIdentity))
        }
        return nil
    }
    res["sourceSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisioningSystem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSourceSystem(val.(*ProvisioningSystem))
        }
        return nil
    }
    res["targetIdentity"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisionedIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetIdentity(val.(*ProvisionedIdentity))
        }
        return nil
    }
    res["targetSystem"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewProvisioningSystem() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTargetSystem(val.(*ProvisioningSystem))
        }
        return nil
    }
    res["tenantId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTenantId(val)
        }
        return nil
    }
    return res
}
func (m *ProvisioningObjectSummary) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetActivityDateTime sets the activityDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
func (m *ProvisioningObjectSummary) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
// SetChangeId sets the changeId property value. Unique ID of this change in this cycle.
func (m *ProvisioningObjectSummary) SetChangeId(value *string)() {
    m.changeId = value
}
// SetCycleId sets the cycleId property value. Unique ID per job iteration.
func (m *ProvisioningObjectSummary) SetCycleId(value *string)() {
    m.cycleId = value
}
// SetDurationInMilliseconds sets the durationInMilliseconds property value. Indicates how long this provisioning action took to finish. Measured in milliseconds.
func (m *ProvisioningObjectSummary) SetDurationInMilliseconds(value *int32)() {
    m.durationInMilliseconds = value
}
// SetInitiatedBy sets the initiatedBy property value. Details of who initiated this provisioning.
func (m *ProvisioningObjectSummary) SetInitiatedBy(value *Initiator)() {
    m.initiatedBy = value
}
// SetJobId sets the jobId property value. The unique ID for the whole provisioning job.
func (m *ProvisioningObjectSummary) SetJobId(value *string)() {
    m.jobId = value
}
// SetModifiedProperties sets the modifiedProperties property value. Details of each property that was modified in this provisioning action on this object.
func (m *ProvisioningObjectSummary) SetModifiedProperties(value []ModifiedProperty)() {
    m.modifiedProperties = value
}
// SetProvisioningAction sets the provisioningAction property value. Indicates the activity name or the operation name. Possible values are: create, update, delete, stageddelete, disable, other and unknownFutureValue. For a list of activities logged, refer to Azure AD activity list.
func (m *ProvisioningObjectSummary) SetProvisioningAction(value *ProvisioningAction)() {
    m.provisioningAction = value
}
// SetProvisioningStatusInfo sets the provisioningStatusInfo property value. Details of provisioning status.
func (m *ProvisioningObjectSummary) SetProvisioningStatusInfo(value *ProvisioningStatusInfo)() {
    m.provisioningStatusInfo = value
}
// SetProvisioningSteps sets the provisioningSteps property value. Details of each step in provisioning.
func (m *ProvisioningObjectSummary) SetProvisioningSteps(value []ProvisioningStep)() {
    m.provisioningSteps = value
}
// SetServicePrincipal sets the servicePrincipal property value. Represents the service principal used for provisioning.
func (m *ProvisioningObjectSummary) SetServicePrincipal(value *ProvisioningServicePrincipal)() {
    m.servicePrincipal = value
}
// SetSourceIdentity sets the sourceIdentity property value. Details of source object being provisioned.
func (m *ProvisioningObjectSummary) SetSourceIdentity(value *ProvisionedIdentity)() {
    m.sourceIdentity = value
}
// SetSourceSystem sets the sourceSystem property value. Details of source system of the object being provisioned.
func (m *ProvisioningObjectSummary) SetSourceSystem(value *ProvisioningSystem)() {
    m.sourceSystem = value
}
// SetTargetIdentity sets the targetIdentity property value. Details of target object being provisioned.
func (m *ProvisioningObjectSummary) SetTargetIdentity(value *ProvisionedIdentity)() {
    m.targetIdentity = value
}
// SetTargetSystem sets the targetSystem property value. Details of target system of the object being provisioned.
func (m *ProvisioningObjectSummary) SetTargetSystem(value *ProvisioningSystem)() {
    m.targetSystem = value
}
// SetTenantId sets the tenantId property value. Unique Azure AD tenant ID.
func (m *ProvisioningObjectSummary) SetTenantId(value *string)() {
    m.tenantId = value
}
