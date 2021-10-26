package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type DirectoryAudit struct {
    Entity
    // Indicates the date and time the activity was performed. The Timestamp type is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Indicates the activity name or the operation name (examples: 'Create User' and 'Add member to group'). For full list, see Azure AD activity list.
    activityDisplayName *string;
    // Indicates additional details on the activity.
    additionalDetails []KeyValue;
    // Indicates which resource category that's targeted by the activity. (For example: User Management, Group Management etc..)
    category *string;
    // Indicates a unique ID that helps correlate activities that span across various services. Can be used to trace logs across services.
    correlationId *string;
    // 
    initiatedBy *AuditActivityInitiator;
    // Indicates information on which service initiated the activity (For example: Self-service Password Management, Core Directory, B2C, Invited Users, Microsoft Identity Manager, Privileged Identity Management.
    loggedByService *string;
    // 
    operationType *string;
    // Indicates the result of the activity. Possible values are: success, failure, timeout, unknownFutureValue.
    result *OperationResult;
    // Indicates the reason for failure if the result is failure or timeout.
    resultReason *string;
    // Indicates information on which resource was changed due to the activity. Target Resource Type can be User, Device, Directory, App, Role, Group, Policy or Other.
    targetResources []TargetResource;
}
// Instantiates a new directoryAudit and sets the default values.
func NewDirectoryAudit()(*DirectoryAudit) {
    m := &DirectoryAudit{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the activityDateTime property value. Indicates the date and time the activity was performed. The Timestamp type is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *DirectoryAudit) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
// Gets the activityDisplayName property value. Indicates the activity name or the operation name (examples: 'Create User' and 'Add member to group'). For full list, see Azure AD activity list.
func (m *DirectoryAudit) GetActivityDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityDisplayName
    }
}
// Gets the additionalDetails property value. Indicates additional details on the activity.
func (m *DirectoryAudit) GetAdditionalDetails()([]KeyValue) {
    if m == nil {
        return nil
    } else {
        return m.additionalDetails
    }
}
// Gets the category property value. Indicates which resource category that's targeted by the activity. (For example: User Management, Group Management etc..)
func (m *DirectoryAudit) GetCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
// Gets the correlationId property value. Indicates a unique ID that helps correlate activities that span across various services. Can be used to trace logs across services.
func (m *DirectoryAudit) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
// Gets the initiatedBy property value. 
func (m *DirectoryAudit) GetInitiatedBy()(*AuditActivityInitiator) {
    if m == nil {
        return nil
    } else {
        return m.initiatedBy
    }
}
// Gets the loggedByService property value. Indicates information on which service initiated the activity (For example: Self-service Password Management, Core Directory, B2C, Invited Users, Microsoft Identity Manager, Privileged Identity Management.
func (m *DirectoryAudit) GetLoggedByService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.loggedByService
    }
}
// Gets the operationType property value. 
func (m *DirectoryAudit) GetOperationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operationType
    }
}
// Gets the result property value. Indicates the result of the activity. Possible values are: success, failure, timeout, unknownFutureValue.
func (m *DirectoryAudit) GetResult()(*OperationResult) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
// Gets the resultReason property value. Indicates the reason for failure if the result is failure or timeout.
func (m *DirectoryAudit) GetResultReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resultReason
    }
}
// Gets the targetResources property value. Indicates information on which resource was changed due to the activity. Target Resource Type can be User, Device, Directory, App, Role, Group, Policy or Other.
func (m *DirectoryAudit) GetTargetResources()([]TargetResource) {
    if m == nil {
        return nil
    } else {
        return m.targetResources
    }
}
// The deserialization information for the current model
func (m *DirectoryAudit) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["activityDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetActivityDateTime(val)
        return nil
    }
    res["activityDisplayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetActivityDisplayName(val)
        return nil
    }
    res["additionalDetails"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewKeyValue() })
        if err != nil {
            return err
        }
        res := make([]KeyValue, len(val))
        for i, v := range val {
            res[i] = *(v.(*KeyValue))
        }
        m.SetAdditionalDetails(res)
        return nil
    }
    res["category"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCategory(val)
        return nil
    }
    res["correlationId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCorrelationId(val)
        return nil
    }
    res["initiatedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAuditActivityInitiator() })
        if err != nil {
            return err
        }
        m.SetInitiatedBy(val.(*AuditActivityInitiator))
        return nil
    }
    res["loggedByService"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetLoggedByService(val)
        return nil
    }
    res["operationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOperationType(val)
        return nil
    }
    res["result"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationResult)
        if err != nil {
            return err
        }
        cast := val.(OperationResult)
        m.SetResult(&cast)
        return nil
    }
    res["resultReason"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResultReason(val)
        return nil
    }
    res["targetResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTargetResource() })
        if err != nil {
            return err
        }
        res := make([]TargetResource, len(val))
        for i, v := range val {
            res[i] = *(v.(*TargetResource))
        }
        m.SetTargetResources(res)
        return nil
    }
    return res
}
func (m *DirectoryAudit) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *DirectoryAudit) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
        err = writer.WriteStringValue("activityDisplayName", m.GetActivityDisplayName())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAdditionalDetails()))
        for i, v := range m.GetAdditionalDetails() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("additionalDetails", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("category", m.GetCategory())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("correlationId", m.GetCorrelationId())
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
        err = writer.WriteStringValue("loggedByService", m.GetLoggedByService())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("operationType", m.GetOperationType())
        if err != nil {
            return err
        }
    }
    if m.GetResult() != nil {
        cast := m.GetResult().String()
        err = writer.WriteStringValue("result", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resultReason", m.GetResultReason())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTargetResources()))
        for i, v := range m.GetTargetResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("targetResources", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the activityDateTime property value. Indicates the date and time the activity was performed. The Timestamp type is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the activityDateTime property.
func (m *DirectoryAudit) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
// Sets the activityDisplayName property value. Indicates the activity name or the operation name (examples: 'Create User' and 'Add member to group'). For full list, see Azure AD activity list.
// Parameters:
//  - value : Value to set for the activityDisplayName property.
func (m *DirectoryAudit) SetActivityDisplayName(value *string)() {
    m.activityDisplayName = value
}
// Sets the additionalDetails property value. Indicates additional details on the activity.
// Parameters:
//  - value : Value to set for the additionalDetails property.
func (m *DirectoryAudit) SetAdditionalDetails(value []KeyValue)() {
    m.additionalDetails = value
}
// Sets the category property value. Indicates which resource category that's targeted by the activity. (For example: User Management, Group Management etc..)
// Parameters:
//  - value : Value to set for the category property.
func (m *DirectoryAudit) SetCategory(value *string)() {
    m.category = value
}
// Sets the correlationId property value. Indicates a unique ID that helps correlate activities that span across various services. Can be used to trace logs across services.
// Parameters:
//  - value : Value to set for the correlationId property.
func (m *DirectoryAudit) SetCorrelationId(value *string)() {
    m.correlationId = value
}
// Sets the initiatedBy property value. 
// Parameters:
//  - value : Value to set for the initiatedBy property.
func (m *DirectoryAudit) SetInitiatedBy(value *AuditActivityInitiator)() {
    m.initiatedBy = value
}
// Sets the loggedByService property value. Indicates information on which service initiated the activity (For example: Self-service Password Management, Core Directory, B2C, Invited Users, Microsoft Identity Manager, Privileged Identity Management.
// Parameters:
//  - value : Value to set for the loggedByService property.
func (m *DirectoryAudit) SetLoggedByService(value *string)() {
    m.loggedByService = value
}
// Sets the operationType property value. 
// Parameters:
//  - value : Value to set for the operationType property.
func (m *DirectoryAudit) SetOperationType(value *string)() {
    m.operationType = value
}
// Sets the result property value. Indicates the result of the activity. Possible values are: success, failure, timeout, unknownFutureValue.
// Parameters:
//  - value : Value to set for the result property.
func (m *DirectoryAudit) SetResult(value *OperationResult)() {
    m.result = value
}
// Sets the resultReason property value. Indicates the reason for failure if the result is failure or timeout.
// Parameters:
//  - value : Value to set for the resultReason property.
func (m *DirectoryAudit) SetResultReason(value *string)() {
    m.resultReason = value
}
// Sets the targetResources property value. Indicates information on which resource was changed due to the activity. Target Resource Type can be User, Device, Directory, App, Role, Group, Policy or Other.
// Parameters:
//  - value : Value to set for the targetResources property.
func (m *DirectoryAudit) SetTargetResources(value []TargetResource)() {
    m.targetResources = value
}
