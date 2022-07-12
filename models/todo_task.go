package models

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// TodoTask provides operations to manage the admin singleton.
type TodoTask struct {
    Entity
    // The task body that typically contains information about the task.
    body ItemBodyable
    // The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
    bodyLastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The categories associated with the task. Each category corresponds to the displayName property of an outlookCategory that the user has defined.
    categories []string
    // A collection of smaller subtasks linked to the more complex parent task.
    checklistItems []ChecklistItemable
    // The date in the specified time zone that the task was finished.
    completedDateTime DateTimeTimeZoneable
    // The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // The date in the specified time zone that the task is to be finished.
    dueDateTime DateTimeTimeZoneable
    // The collection of open extensions defined for the task. Nullable.
    extensions []Extensionable
    // The importance property
    importance *Importance
    // Set to true if an alert is set to remind the user of the task.
    isReminderOn *bool
    // The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
    lastModifiedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time
    // A collection of resources linked to the task.
    linkedResources []LinkedResourceable
    // The recurrence pattern for the task.
    recurrence PatternedRecurrenceable
    // The date and time for a reminder alert of the task to occur.
    reminderDateTime DateTimeTimeZoneable
    // The status property
    status *TaskStatus
    // A brief description of the task.
    title *string
}
// NewTodoTask instantiates a new todoTask and sets the default values.
func NewTodoTask()(*TodoTask) {
    m := &TodoTask{
        Entity: *NewEntity(),
    }
    return m
}
// CreateTodoTaskFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTodoTaskFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTodoTask(), nil
}
// GetBody gets the body property value. The task body that typically contains information about the task.
func (m *TodoTask) GetBody()(ItemBodyable) {
    if m == nil {
        return nil
    } else {
        return m.body
    }
}
// GetBodyLastModifiedDateTime gets the bodyLastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *TodoTask) GetBodyLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.bodyLastModifiedDateTime
    }
}
// GetCategories gets the categories property value. The categories associated with the task. Each category corresponds to the displayName property of an outlookCategory that the user has defined.
func (m *TodoTask) GetCategories()([]string) {
    if m == nil {
        return nil
    } else {
        return m.categories
    }
}
// GetChecklistItems gets the checklistItems property value. A collection of smaller subtasks linked to the more complex parent task.
func (m *TodoTask) GetChecklistItems()([]ChecklistItemable) {
    if m == nil {
        return nil
    } else {
        return m.checklistItems
    }
}
// GetCompletedDateTime gets the completedDateTime property value. The date in the specified time zone that the task was finished.
func (m *TodoTask) GetCompletedDateTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
// GetCreatedDateTime gets the createdDateTime property value. The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *TodoTask) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetDueDateTime gets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *TodoTask) GetDueDateTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
// GetExtensions gets the extensions property value. The collection of open extensions defined for the task. Nullable.
func (m *TodoTask) GetExtensions()([]Extensionable) {
    if m == nil {
        return nil
    } else {
        return m.extensions
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TodoTask) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["body"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateItemBodyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBody(val.(ItemBodyable))
        }
        return nil
    }
    res["bodyLastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBodyLastModifiedDateTime(val)
        }
        return nil
    }
    res["categories"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetCategories(res)
        }
        return nil
    }
    res["checklistItems"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateChecklistItemFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ChecklistItemable, len(val))
            for i, v := range val {
                res[i] = v.(ChecklistItemable)
            }
            m.SetChecklistItems(res)
        }
        return nil
    }
    res["completedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["createdDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["dueDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["extensions"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateExtensionFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Extensionable, len(val))
            for i, v := range val {
                res[i] = v.(Extensionable)
            }
            m.SetExtensions(res)
        }
        return nil
    }
    res["importance"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseImportance)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImportance(val.(*Importance))
        }
        return nil
    }
    res["isReminderOn"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsReminderOn(val)
        }
        return nil
    }
    res["lastModifiedDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLastModifiedDateTime(val)
        }
        return nil
    }
    res["linkedResources"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateLinkedResourceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]LinkedResourceable, len(val))
            for i, v := range val {
                res[i] = v.(LinkedResourceable)
            }
            m.SetLinkedResources(res)
        }
        return nil
    }
    res["recurrence"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePatternedRecurrenceFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRecurrence(val.(PatternedRecurrenceable))
        }
        return nil
    }
    res["reminderDateTime"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReminderDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseTaskStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*TaskStatus))
        }
        return nil
    }
    res["title"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetImportance gets the importance property value. The importance property
func (m *TodoTask) GetImportance()(*Importance) {
    if m == nil {
        return nil
    } else {
        return m.importance
    }
}
// GetIsReminderOn gets the isReminderOn property value. Set to true if an alert is set to remind the user of the task.
func (m *TodoTask) GetIsReminderOn()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isReminderOn
    }
}
// GetLastModifiedDateTime gets the lastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *TodoTask) GetLastModifiedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.lastModifiedDateTime
    }
}
// GetLinkedResources gets the linkedResources property value. A collection of resources linked to the task.
func (m *TodoTask) GetLinkedResources()([]LinkedResourceable) {
    if m == nil {
        return nil
    } else {
        return m.linkedResources
    }
}
// GetRecurrence gets the recurrence property value. The recurrence pattern for the task.
func (m *TodoTask) GetRecurrence()(PatternedRecurrenceable) {
    if m == nil {
        return nil
    } else {
        return m.recurrence
    }
}
// GetReminderDateTime gets the reminderDateTime property value. The date and time for a reminder alert of the task to occur.
func (m *TodoTask) GetReminderDateTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.reminderDateTime
    }
}
// GetStatus gets the status property value. The status property
func (m *TodoTask) GetStatus()(*TaskStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetTitle gets the title property value. A brief description of the task.
func (m *TodoTask) GetTitle()(*string) {
    if m == nil {
        return nil
    } else {
        return m.title
    }
}
// Serialize serializes information the current object
func (m *TodoTask) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
    if m.GetCategories() != nil {
        err = writer.WriteCollectionOfStringValues("categories", m.GetCategories())
        if err != nil {
            return err
        }
    }
    if m.GetChecklistItems() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetChecklistItems()))
        for i, v := range m.GetChecklistItems() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("checklistItems", cast)
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
    if m.GetExtensions() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetExtensions()))
        for i, v := range m.GetExtensions() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("extensions", cast)
        if err != nil {
            return err
        }
    }
    if m.GetImportance() != nil {
        cast := (*m.GetImportance()).String()
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
    if m.GetLinkedResources() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetLinkedResources()))
        for i, v := range m.GetLinkedResources() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
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
        cast := (*m.GetStatus()).String()
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
// SetBody sets the body property value. The task body that typically contains information about the task.
func (m *TodoTask) SetBody(value ItemBodyable)() {
    if m != nil {
        m.body = value
    }
}
// SetBodyLastModifiedDateTime sets the bodyLastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *TodoTask) SetBodyLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.bodyLastModifiedDateTime = value
    }
}
// SetCategories sets the categories property value. The categories associated with the task. Each category corresponds to the displayName property of an outlookCategory that the user has defined.
func (m *TodoTask) SetCategories(value []string)() {
    if m != nil {
        m.categories = value
    }
}
// SetChecklistItems sets the checklistItems property value. A collection of smaller subtasks linked to the more complex parent task.
func (m *TodoTask) SetChecklistItems(value []ChecklistItemable)() {
    if m != nil {
        m.checklistItems = value
    }
}
// SetCompletedDateTime sets the completedDateTime property value. The date in the specified time zone that the task was finished.
func (m *TodoTask) SetCompletedDateTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.completedDateTime = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. The date and time when the task was created. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *TodoTask) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetDueDateTime sets the dueDateTime property value. The date in the specified time zone that the task is to be finished.
func (m *TodoTask) SetDueDateTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.dueDateTime = value
    }
}
// SetExtensions sets the extensions property value. The collection of open extensions defined for the task. Nullable.
func (m *TodoTask) SetExtensions(value []Extensionable)() {
    if m != nil {
        m.extensions = value
    }
}
// SetImportance sets the importance property value. The importance property
func (m *TodoTask) SetImportance(value *Importance)() {
    if m != nil {
        m.importance = value
    }
}
// SetIsReminderOn sets the isReminderOn property value. Set to true if an alert is set to remind the user of the task.
func (m *TodoTask) SetIsReminderOn(value *bool)() {
    if m != nil {
        m.isReminderOn = value
    }
}
// SetLastModifiedDateTime sets the lastModifiedDateTime property value. The date and time when the task was last modified. By default, it is in UTC. You can provide a custom time zone in the request header. The property value uses ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2020 would look like this: '2020-01-01T00:00:00Z'.
func (m *TodoTask) SetLastModifiedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.lastModifiedDateTime = value
    }
}
// SetLinkedResources sets the linkedResources property value. A collection of resources linked to the task.
func (m *TodoTask) SetLinkedResources(value []LinkedResourceable)() {
    if m != nil {
        m.linkedResources = value
    }
}
// SetRecurrence sets the recurrence property value. The recurrence pattern for the task.
func (m *TodoTask) SetRecurrence(value PatternedRecurrenceable)() {
    if m != nil {
        m.recurrence = value
    }
}
// SetReminderDateTime sets the reminderDateTime property value. The date and time for a reminder alert of the task to occur.
func (m *TodoTask) SetReminderDateTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.reminderDateTime = value
    }
}
// SetStatus sets the status property value. The status property
func (m *TodoTask) SetStatus(value *TaskStatus)() {
    if m != nil {
        m.status = value
    }
}
// SetTitle sets the title property value. A brief description of the task.
func (m *TodoTask) SetTitle(value *string)() {
    if m != nil {
        m.title = value
    }
}
