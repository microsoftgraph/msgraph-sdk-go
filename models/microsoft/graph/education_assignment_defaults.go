package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationAssignmentDefaults 
type EducationAssignmentDefaults struct {
    Entity
    // Class-level default behavior for handling students who are added after the assignment is published. Possible values are: none, assignIfOpen.
    addedStudentAction *EducationAddedStudentAction;
    // Class-level default value for due time field. Default value is 23:59:00.
    dueTime *string;
    // Default Teams channel to which notifications will be sent. Default value is null.
    notificationChannelUrl *string;
}
// NewEducationAssignmentDefaults instantiates a new educationAssignmentDefaults and sets the default values.
func NewEducationAssignmentDefaults()(*EducationAssignmentDefaults) {
    m := &EducationAssignmentDefaults{
        Entity: *NewEntity(),
    }
    return m
}
// GetAddedStudentAction gets the addedStudentAction property value. Class-level default behavior for handling students who are added after the assignment is published. Possible values are: none, assignIfOpen.
func (m *EducationAssignmentDefaults) GetAddedStudentAction()(*EducationAddedStudentAction) {
    if m == nil {
        return nil
    } else {
        return m.addedStudentAction
    }
}
// GetDueTime gets the dueTime property value. Class-level default value for due time field. Default value is 23:59:00.
func (m *EducationAssignmentDefaults) GetDueTime()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dueTime
    }
}
// GetNotificationChannelUrl gets the notificationChannelUrl property value. Default Teams channel to which notifications will be sent. Default value is null.
func (m *EducationAssignmentDefaults) GetNotificationChannelUrl()(*string) {
    if m == nil {
        return nil
    } else {
        return m.notificationChannelUrl
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationAssignmentDefaults) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["addedStudentAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseEducationAddedStudentAction)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(EducationAddedStudentAction)
            m.SetAddedStudentAction(&cast)
        }
        return nil
    }
    res["dueTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueTime(val)
        }
        return nil
    }
    res["notificationChannelUrl"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNotificationChannelUrl(val)
        }
        return nil
    }
    return res
}
func (m *EducationAssignmentDefaults) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EducationAssignmentDefaults) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAddedStudentAction() != nil {
        cast := m.GetAddedStudentAction().String()
        err = writer.WriteStringValue("addedStudentAction", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("dueTime", m.GetDueTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("notificationChannelUrl", m.GetNotificationChannelUrl())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAddedStudentAction sets the addedStudentAction property value. Class-level default behavior for handling students who are added after the assignment is published. Possible values are: none, assignIfOpen.
func (m *EducationAssignmentDefaults) SetAddedStudentAction(value *EducationAddedStudentAction)() {
    if m != nil {
        m.addedStudentAction = value
    }
}
// SetDueTime sets the dueTime property value. Class-level default value for due time field. Default value is 23:59:00.
func (m *EducationAssignmentDefaults) SetDueTime(value *string)() {
    if m != nil {
        m.dueTime = value
    }
}
// SetNotificationChannelUrl sets the notificationChannelUrl property value. Default Teams channel to which notifications will be sent. Default value is null.
func (m *EducationAssignmentDefaults) SetNotificationChannelUrl(value *string)() {
    if m != nil {
        m.notificationChannelUrl = value
    }
}
