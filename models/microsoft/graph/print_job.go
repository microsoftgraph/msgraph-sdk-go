package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new printJob and sets the default values.
func NewPrintJob()(*PrintJob) {
    m := &PrintJob{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the configuration property value. 
func (m *PrintJob) GetConfiguration()(*PrintJobConfiguration) {
    if m == nil {
        return nil
    } else {
        return m.configuration
    }
}
// Gets the createdBy property value. Read-only. Nullable.
func (m *PrintJob) GetCreatedBy()(*UserIdentity) {
    if m == nil {
        return nil
    } else {
        return m.createdBy
    }
}
// Gets the createdDateTime property value. The DateTimeOffset when the job was created. Read-only.
func (m *PrintJob) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the documents property value. Read-only.
func (m *PrintJob) GetDocuments()([]PrintDocument) {
    if m == nil {
        return nil
    } else {
        return m.documents
    }
}
// Gets the isFetchable property value. If true, document can be fetched by printer.
func (m *PrintJob) GetIsFetchable()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFetchable
    }
}
// Gets the redirectedFrom property value. Contains the source job URL, if the job has been redirected from another printer.
func (m *PrintJob) GetRedirectedFrom()(*string) {
    if m == nil {
        return nil
    } else {
        return m.redirectedFrom
    }
}
// Gets the redirectedTo property value. Contains the destination job URL, if the job has been redirected to another printer.
func (m *PrintJob) GetRedirectedTo()(*string) {
    if m == nil {
        return nil
    } else {
        return m.redirectedTo
    }
}
// Gets the status property value. 
func (m *PrintJob) GetStatus()(*PrintJobStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the tasks property value. A list of printTasks that were triggered by this print job.
func (m *PrintJob) GetTasks()([]PrintTask) {
    if m == nil {
        return nil
    } else {
        return m.tasks
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the configuration property value. 
// Parameters:
//  - value : Value to set for the configuration property.
func (m *PrintJob) SetConfiguration(value *PrintJobConfiguration)() {
    m.configuration = value
}
// Sets the createdBy property value. Read-only. Nullable.
// Parameters:
//  - value : Value to set for the createdBy property.
func (m *PrintJob) SetCreatedBy(value *UserIdentity)() {
    m.createdBy = value
}
// Sets the createdDateTime property value. The DateTimeOffset when the job was created. Read-only.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *PrintJob) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the documents property value. Read-only.
// Parameters:
//  - value : Value to set for the documents property.
func (m *PrintJob) SetDocuments(value []PrintDocument)() {
    m.documents = value
}
// Sets the isFetchable property value. If true, document can be fetched by printer.
// Parameters:
//  - value : Value to set for the isFetchable property.
func (m *PrintJob) SetIsFetchable(value *bool)() {
    m.isFetchable = value
}
// Sets the redirectedFrom property value. Contains the source job URL, if the job has been redirected from another printer.
// Parameters:
//  - value : Value to set for the redirectedFrom property.
func (m *PrintJob) SetRedirectedFrom(value *string)() {
    m.redirectedFrom = value
}
// Sets the redirectedTo property value. Contains the destination job URL, if the job has been redirected to another printer.
// Parameters:
//  - value : Value to set for the redirectedTo property.
func (m *PrintJob) SetRedirectedTo(value *string)() {
    m.redirectedTo = value
}
// Sets the status property value. 
// Parameters:
//  - value : Value to set for the status property.
func (m *PrintJob) SetStatus(value *PrintJobStatus)() {
    m.status = value
}
// Sets the tasks property value. A list of printTasks that were triggered by this print job.
// Parameters:
//  - value : Value to set for the tasks property.
func (m *PrintJob) SetTasks(value []PrintTask)() {
    m.tasks = value
}
