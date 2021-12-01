package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FollowupFlag 
type FollowupFlag struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The date and time that the follow-up was finished.
    completedDateTime *DateTimeTimeZone;
    // The date and time that the follow up is to be finished. Note: To set the due date, you must also specify the startDateTime; otherwise, you will get a 400 Bad Request response.
    dueDateTime *DateTimeTimeZone;
    // The status for follow-up for an item. Possible values are notFlagged, complete, and flagged.
    flagStatus *FollowupFlagStatus;
    // The date and time that the follow-up is to begin.
    startDateTime *DateTimeTimeZone;
}
// NewFollowupFlag instantiates a new followupFlag and sets the default values.
func NewFollowupFlag()(*FollowupFlag) {
    m := &FollowupFlag{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FollowupFlag) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCompletedDateTime gets the completedDateTime property value. The date and time that the follow-up was finished.
func (m *FollowupFlag) GetCompletedDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
// GetDueDateTime gets the dueDateTime property value. The date and time that the follow up is to be finished. Note: To set the due date, you must also specify the startDateTime; otherwise, you will get a 400 Bad Request response.
func (m *FollowupFlag) GetDueDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
// GetFlagStatus gets the flagStatus property value. The status for follow-up for an item. Possible values are notFlagged, complete, and flagged.
func (m *FollowupFlag) GetFlagStatus()(*FollowupFlagStatus) {
    if m == nil {
        return nil
    } else {
        return m.flagStatus
    }
}
// GetStartDateTime gets the startDateTime property value. The date and time that the follow-up is to begin.
func (m *FollowupFlag) GetStartDateTime()(*DateTimeTimeZone) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FollowupFlag) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["flagStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFollowupFlagStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(FollowupFlagStatus)
            m.SetFlagStatus(&cast)
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewDateTimeTimeZone() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(*DateTimeTimeZone))
        }
        return nil
    }
    return res
}
func (m *FollowupFlag) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *FollowupFlag) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("completedDateTime", m.GetCompletedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("dueDateTime", m.GetDueDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetFlagStatus() != nil {
        cast := m.GetFlagStatus().String()
        err := writer.WriteStringValue("flagStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("startDateTime", m.GetStartDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FollowupFlag) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCompletedDateTime sets the completedDateTime property value. The date and time that the follow-up was finished.
func (m *FollowupFlag) SetCompletedDateTime(value *DateTimeTimeZone)() {
    if m != nil {
        m.completedDateTime = value
    }
}
// SetDueDateTime sets the dueDateTime property value. The date and time that the follow up is to be finished. Note: To set the due date, you must also specify the startDateTime; otherwise, you will get a 400 Bad Request response.
func (m *FollowupFlag) SetDueDateTime(value *DateTimeTimeZone)() {
    if m != nil {
        m.dueDateTime = value
    }
}
// SetFlagStatus sets the flagStatus property value. The status for follow-up for an item. Possible values are notFlagged, complete, and flagged.
func (m *FollowupFlag) SetFlagStatus(value *FollowupFlagStatus)() {
    if m != nil {
        m.flagStatus = value
    }
}
// SetStartDateTime sets the startDateTime property value. The date and time that the follow-up is to begin.
func (m *FollowupFlag) SetStartDateTime(value *DateTimeTimeZone)() {
    if m != nil {
        m.startDateTime = value
    }
}
