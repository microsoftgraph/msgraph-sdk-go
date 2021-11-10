package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TodoTask struct {
    Entity
    // The task body that typically contains information about the task.
    body *ItemBody;
    // The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
    bodyLastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The date in the specified time zone that the task was finished.
    completedDateTime *DateTimeTimeZone;
    // The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The date in the specified time zone that the task is to be finished.
    dueDateTime *DateTimeTimeZone;
    // The collection of open extensions defined for the task. Nullable.
    extensions []Extension;
    // The importance of the task. Possible values are: low, normal, high.
    importance *Importance;
    // Set to true if an alert is set to remind the user of the task.
    isReminderOn *bool;
    // The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // A collection of resources linked to the task.
    linkedResources []LinkedResource;
    // The recurrence pattern for the task.
    recurrence *PatternedRecurrence;
    // The date and time for a reminder alert of the task to occur.
    reminderDateTime *DateTimeTimeZone;
    // Indicates the state or progress of the task. Possible values are: notStarted, inProgress, completed, waitingOnOthers, deferred.
    status *TaskStatus;
    // A brief description of the task.
    title *string;
}
// Instantiates a new todoTask and sets the default values.
func NewTodoTask()(*TodoTask) {
    m := &TodoTask{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the body property value. The task body that typically contains information about the task.
func (m *TodoTask) GetBody()(*ItemBody) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
// Gets the bodyLastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *TodoTask) GetBodyLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.bodyLastModifiedDateTime
    }
}
// Gets the completedDateTime property value. The date in the specified time zone that the task was finished.
func (m *TodoTask) GetCompletedDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
// Gets the createdDateTime property value. The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *TodoTask) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *TodoTask) GetDueDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
// Gets the extensions property value. The collection of open extensions defined for the task. Nullable.
func (m *TodoTask) GetExtensions()([]Extension) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// Gets the importance property value. The importance of the task. Possible values are: low, normal, high.
func (m *TodoTask) GetImportance()(*Importance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// Gets the isReminderOn property value. Set to true if an alert is set to remind the user of the task.
func (m *TodoTask) GetIsReminderOn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReminderOn
    }
}
// Gets the lastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *TodoTask) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// Gets the linkedResources property value. A collection of resources linked to the task.
func (m *TodoTask) GetLinkedResources()([]LinkedResource) {
    if m == nil {
        return nil
    } else {
        return m.linkedResources
    }
}
// Gets the recurrence property value. The recurrence pattern for the task.
func (m *TodoTask) GetRecurrence()(*PatternedRecurrence) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// Gets the reminderDateTime property value. The date and time for a reminder alert of the task to occur.
func (m *TodoTask) GetReminderDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.reminderDateTime
    }
}
// Gets the status property value. Indicates the state or progress of the task. Possible values are: notStarted, inProgress, completed, waitingOnOthers, deferred.
func (m *TodoTask) GetStatus()(*TaskStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Gets the title property value. A brief description of the task.
func (m *TodoTask) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// The deserialization information for the current model
func (m *TodoTask) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["body"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewItemBody() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(*ItemBody))
        }
        return nil
    }
    res["bodyLastModifiedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyLastModifiedDateTime(val)
        }
        return nil
    }
    res["completedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val.(*DateTimeTimeZone))
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
    res["dueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    res["extensions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewExtension() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extension, len(val))
            for i, v := range val {
                res[i] = *(v.(*Extension))
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["importance"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(Importance)
            m.SetImportance(&cast)
        }
        return nil
    }
    res["isReminderOn"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReminderOn(val)
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
    res["linkedResources"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewLinkedResource() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LinkedResource, len(val))
            for i, v := range val {
                res[i] = *(v.(*LinkedResource))
            }
            m.SetLinkedResources(res)
        }
        return nil
    }
    res["recurrence"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPatternedRecurrence() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(*PatternedRecurrence))
        }
        return nil
    }
    res["reminderDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminderDateTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseTaskStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(TaskStatus)
            m.SetStatus(&cast)
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
    return res
}
func (m *TodoTask) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TodoTask) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteObjectValue("body", m.GetBody())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteTimeValue("bodyLastModifiedDateTime", m.GetBodyLastModifiedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("completedDateTime", m.GetCompletedDateTime())
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
        err = writer.WriteObjectValue("dueDateTime", m.GetDueDateTime())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := m.GetImportance().String()
        err = writer.WriteStringValue("importance", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isReminderOn", m.GetIsReminderOn())
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
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetLinkedResources()))
        for i, v := range m.GetLinkedResources() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("linkedResources", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("recurrence", m.GetRecurrence())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("reminderDateTime", m.GetReminderDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err = writer.WriteStringValue("status", &cast)
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
    return nil
}
// Sets the body property value. The task body that typically contains information about the task.
// Parameters:
//  - value : Value to set for the body property.
func (m *TodoTask) SetBody(value *ItemBody)() {
    m.body = value
}
// Sets the bodyLastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
// Parameters:
//  - value : Value to set for the bodyLastModifiedDateTime property.
func (m *TodoTask) SetBodyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.bodyLastModifiedDateTime = value
}
// Sets the completedDateTime property value. The date in the specified time zone that the task was finished.
// Parameters:
//  - value : Value to set for the completedDateTime property.
func (m *TodoTask) SetCompletedDateTime(value *DateTimeTimeZone)() {
    m.completedDateTime = value
}
// Sets the createdDateTime property value. The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *TodoTask) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
// Parameters:
//  - value : Value to set for the dueDateTime property.
func (m *TodoTask) SetDueDateTime(value *DateTimeTimeZone)() {
    m.dueDateTime = value
}
// Sets the extensions property value. The collection of open extensions defined for the task. Nullable.
// Parameters:
//  - value : Value to set for the extensions property.
func (m *TodoTask) SetExtensions(value []Extension)() {
    m.extensions = value
}
// Sets the importance property value. The importance of the task. Possible values are: low, normal, high.
// Parameters:
//  - value : Value to set for the importance property.
func (m *TodoTask) SetImportance(value *Importance)() {
    m.importance = value
}
// Sets the isReminderOn property value. Set to true if an alert is set to remind the user of the task.
// Parameters:
//  - value : Value to set for the isReminderOn property.
func (m *TodoTask) SetIsReminderOn(value *bool)() {
    m.isReminderOn = value
}
// Sets the lastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
// Parameters:
//  - value : Value to set for the lastModifiedDateTime property.
func (m *TodoTask) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.lastModifiedDateTime = value
}
// Sets the linkedResources property value. A collection of resources linked to the task.
// Parameters:
//  - value : Value to set for the linkedResources property.
func (m *TodoTask) SetLinkedResources(value []LinkedResource)() {
    m.linkedResources = value
}
// Sets the recurrence property value. The recurrence pattern for the task.
// Parameters:
//  - value : Value to set for the recurrence property.
func (m *TodoTask) SetRecurrence(value *PatternedRecurrence)() {
    m.recurrence = value
}
// Sets the reminderDateTime property value. The date and time for a reminder alert of the task to occur.
// Parameters:
//  - value : Value to set for the reminderDateTime property.
func (m *TodoTask) SetReminderDateTime(value *DateTimeTimeZone)() {
    m.reminderDateTime = value
}
// Sets the status property value. Indicates the state or progress of the task. Possible values are: notStarted, inProgress, completed, waitingOnOthers, deferred.
// Parameters:
//  - value : Value to set for the status property.
func (m *TodoTask) SetStatus(value *TaskStatus)() {
    m.status = value
}
// Sets the title property value. A brief description of the task.
// Parameters:
//  - value : Value to set for the title property.
func (m *TodoTask) SetTitle(value *string)() {
    m.title = value
}
