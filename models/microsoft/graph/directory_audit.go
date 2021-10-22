package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type DirectoryAudit struct {
    Entity
    activityDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    activityDisplayName *string;
    additionalDetails []KeyValue;
    category *string;
    correlationId *string;
    initiatedBy *AuditActivityInitiator;
    loggedByService *string;
    operationType *string;
    result *OperationResult;
    resultReason *string;
    targetResources []TargetResource;
}
func NewDirectoryAudit()(*DirectoryAudit) {
    m := &DirectoryAudit{
        Entity: *NewEntity(),
    }
    return m
}
func (m *DirectoryAudit) GetActivityDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.activityDateTime
    }
}
func (m *DirectoryAudit) GetActivityDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.activityDisplayName
    }
}
func (m *DirectoryAudit) GetAdditionalDetails()([]KeyValue) {
    if m == nil {
        return nil
    } else {
        return m.additionalDetails
    }
}
func (m *DirectoryAudit) GetCategory()(*string) {
    if m == nil {
        return nil
    } else {
        return m.category
    }
}
func (m *DirectoryAudit) GetCorrelationId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.correlationId
    }
}
func (m *DirectoryAudit) GetInitiatedBy()(*AuditActivityInitiator) {
    if m == nil {
        return nil
    } else {
        return m.initiatedBy
    }
}
func (m *DirectoryAudit) GetLoggedByService()(*string) {
    if m == nil {
        return nil
    } else {
        return m.loggedByService
    }
}
func (m *DirectoryAudit) GetOperationType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.operationType
    }
}
func (m *DirectoryAudit) GetResult()(*OperationResult) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
func (m *DirectoryAudit) GetResultReason()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resultReason
    }
}
func (m *DirectoryAudit) GetTargetResources()([]TargetResource) {
    if m == nil {
        return nil
    } else {
        return m.targetResources
    }
}
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
func (m *DirectoryAudit) SetActivityDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.activityDateTime = value
}
func (m *DirectoryAudit) SetActivityDisplayName(value *string)() {
    m.activityDisplayName = value
}
func (m *DirectoryAudit) SetAdditionalDetails(value []KeyValue)() {
    m.additionalDetails = value
}
func (m *DirectoryAudit) SetCategory(value *string)() {
    m.category = value
}
func (m *DirectoryAudit) SetCorrelationId(value *string)() {
    m.correlationId = value
}
func (m *DirectoryAudit) SetInitiatedBy(value *AuditActivityInitiator)() {
    m.initiatedBy = value
}
func (m *DirectoryAudit) SetLoggedByService(value *string)() {
    m.loggedByService = value
}
func (m *DirectoryAudit) SetOperationType(value *string)() {
    m.operationType = value
}
func (m *DirectoryAudit) SetResult(value *OperationResult)() {
    m.result = value
}
func (m *DirectoryAudit) SetResultReason(value *string)() {
    m.resultReason = value
}
func (m *DirectoryAudit) SetTargetResources(value []TargetResource)() {
    m.targetResources = value
}
