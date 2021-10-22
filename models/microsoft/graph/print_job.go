package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PrintJob struct {
    Entity
    configuration *PrintJobConfiguration;
    createdBy *UserIdentity;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    documents []PrintDocument;
    isFetchable *bool;
    redirectedFrom *string;
    redirectedTo *string;
    status *PrintJobStatus;
    tasks []PrintTask;
}
func NewPrintJob()(*PrintJob) {
    m := &PrintJob{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PrintJob) GetConfiguration()(*PrintJobConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.configuration
    }
}
func (m *PrintJob) GetCreatedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
func (m *PrintJob) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *PrintJob) GetDocuments()([]PrintDocument) {
    if m == nil {
        return nil
    } else {
        return m.documents
    }
}
func (m *PrintJob) GetIsFetchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFetchable
    }
}
func (m *PrintJob) GetRedirectedFrom()(*string) {
    if m == nil {
        return nil
    } else {
        return m.redirectedFrom
    }
}
func (m *PrintJob) GetRedirectedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.redirectedTo
    }
}
func (m *PrintJob) GetStatus()(*PrintJobStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
func (m *PrintJob) GetTasks()([]PrintTask) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
func (m *PrintJob) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintJobConfiguration() })
        if err != nil {
            return err
        }
        m.SetConfiguration(val.(*PrintJobConfiguration))
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        m.SetCreatedBy(val.(*UserIdentity))
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["documents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintDocument() })
        if err != nil {
            return err
        }
        res := make([]PrintDocument, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintDocument))
        }
        m.SetDocuments(res)
        return nil
    }
    res["isFetchable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsFetchable(val)
        return nil
    }
    res["redirectedFrom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRedirectedFrom(val)
        return nil
    }
    res["redirectedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRedirectedTo(val)
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintJobStatus() })
        if err != nil {
            return err
        }
        m.SetStatus(val.(*PrintJobStatus))
        return nil
    }
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintTask() })
        if err != nil {
            return err
        }
        res := make([]PrintTask, len(val))
        for i, v := range val {
            res[i] = *(v.(*PrintTask))
        }
        m.SetTasks(res)
        return nil
    }
    return res
}
func (m *PrintJob) IsNil()(bool) {
    return m == nil
}
func (m *PrintJob) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("configuration", m.GetConfiguration())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("createdBy", m.GetCreatedBy())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetDocuments()))
        for i, v := range m.GetDocuments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("documents", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isFetchable", m.GetIsFetchable())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("redirectedFrom", m.GetRedirectedFrom())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("redirectedTo", m.GetRedirectedTo())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("status", m.GetStatus())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTasks()))
        for i, v := range m.GetTasks() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("tasks", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PrintJob) SetConfiguration(value *PrintJobConfiguration)() {
    m.configuration = value
}
func (m *PrintJob) SetCreatedBy(value *UserIdentity)() {
    m.createdBy = value
}
func (m *PrintJob) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *PrintJob) SetDocuments(value []PrintDocument)() {
    m.documents = value
}
func (m *PrintJob) SetIsFetchable(value *bool)() {
    m.isFetchable = value
}
func (m *PrintJob) SetRedirectedFrom(value *string)() {
    m.redirectedFrom = value
}
func (m *PrintJob) SetRedirectedTo(value *string)() {
    m.redirectedTo = value
}
func (m *PrintJob) SetStatus(value *PrintJobStatus)() {
    m.status = value
}
func (m *PrintJob) SetTasks(value []PrintTask)() {
    m.tasks = value
}
