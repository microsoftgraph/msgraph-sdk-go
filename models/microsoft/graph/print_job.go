package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PrintJob 
type PrintJob struct {
    Entity
    // 
    configuration *PrintJobConfiguration;
    // Read-only. Nullable.
    createdBy *UserIdentity;
    // The DateTimeOffset when the job was created. Read-only.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Read-only.
    documents []PrintDocument;
    // If true, document can be fetched by printer.
    isFetchable *bool;
    // Contains the source job URL, if the job has been redirected from another printer.
    redirectedFrom *string;
    // Contains the destination job URL, if the job has been redirected to another printer.
    redirectedTo *string;
    // 
    status *PrintJobStatus;
    // A list of printTasks that were triggered by this print job.
    tasks []PrintTask;
}
// NewPrintJob instantiates a new printJob and sets the default values.
func NewPrintJob()(*PrintJob) {
    m := &PrintJob{
        Entity: *NewEntity(),
    }
    return m
}
// GetConfiguration gets the configuration property value. 
func (m *PrintJob) GetConfiguration()(*PrintJobConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.configuration
    }
}
// GetCreatedBy gets the createdBy property value. Read-only. Nullable.
func (m *PrintJob) GetCreatedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The DateTimeOffset when the job was created. Read-only.
func (m *PrintJob) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDocuments gets the documents property value. Read-only.
func (m *PrintJob) GetDocuments()([]PrintDocument) {
    if m == nil {
        return nil
    } else {
        return m.documents
    }
}
// GetIsFetchable gets the isFetchable property value. If true, document can be fetched by printer.
func (m *PrintJob) GetIsFetchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFetchable
    }
}
// GetRedirectedFrom gets the redirectedFrom property value. Contains the source job URL, if the job has been redirected from another printer.
func (m *PrintJob) GetRedirectedFrom()(*string) {
    if m == nil {
        return nil
    } else {
        return m.redirectedFrom
    }
}
// GetRedirectedTo gets the redirectedTo property value. Contains the destination job URL, if the job has been redirected to another printer.
func (m *PrintJob) GetRedirectedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.redirectedTo
    }
}
// GetStatus gets the status property value. 
func (m *PrintJob) GetStatus()(*PrintJobStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTasks gets the tasks property value. A list of printTasks that were triggered by this print job.
func (m *PrintJob) GetTasks()([]PrintTask) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PrintJob) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["configuration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintJobConfiguration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConfiguration(val.(*PrintJobConfiguration))
        }
        return nil
    }
    res["createdBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserIdentity() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedBy(val.(*UserIdentity))
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["documents"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintDocument() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintDocument, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintDocument))
            }
            m.SetDocuments(res)
        }
        return nil
    }
    res["isFetchable"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsFetchable(val)
        }
        return nil
    }
    res["redirectedFrom"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedirectedFrom(val)
        }
        return nil
    }
    res["redirectedTo"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRedirectedTo(val)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintJobStatus() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*PrintJobStatus))
        }
        return nil
    }
    res["tasks"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPrintTask() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]PrintTask, len(val))
            for i, v := range val {
                res[i] = *(v.(*PrintTask))
            }
            m.SetTasks(res)
        }
        return nil
    }
    return res
}
func (m *PrintJob) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetConfiguration sets the configuration property value. 
func (m *PrintJob) SetConfiguration(value *PrintJobConfiguration)() {
    if m != nil {
        m.configuration = value
    }
}
// SetCreatedBy sets the createdBy property value. Read-only. Nullable.
func (m *PrintJob) SetCreatedBy(value *UserIdentity)() {
    if m != nil {
        m.createdBy = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The DateTimeOffset when the job was created. Read-only.
func (m *PrintJob) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDocuments sets the documents property value. Read-only.
func (m *PrintJob) SetDocuments(value []PrintDocument)() {
    if m != nil {
        m.documents = value
    }
}
// SetIsFetchable sets the isFetchable property value. If true, document can be fetched by printer.
func (m *PrintJob) SetIsFetchable(value *bool)() {
    if m != nil {
        m.isFetchable = value
    }
}
// SetRedirectedFrom sets the redirectedFrom property value. Contains the source job URL, if the job has been redirected from another printer.
func (m *PrintJob) SetRedirectedFrom(value *string)() {
    if m != nil {
        m.redirectedFrom = value
    }
}
// SetRedirectedTo sets the redirectedTo property value. Contains the destination job URL, if the job has been redirected to another printer.
func (m *PrintJob) SetRedirectedTo(value *string)() {
    if m != nil {
        m.redirectedTo = value
    }
}
// SetStatus sets the status property value. 
func (m *PrintJob) SetStatus(value *PrintJobStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetTasks sets the tasks property value. A list of printTasks that were triggered by this print job.
func (m *PrintJob) SetTasks(value []PrintTask)() {
    if m != nil {
        m.tasks = value
    }
}
