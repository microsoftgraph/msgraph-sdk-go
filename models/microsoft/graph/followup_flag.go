package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// FollowupFlag 
type FollowupFlag struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The date and time that the follow-up was finished.
    completedDateTime DateTimeTimeZoneable;
    // The date and time that the follow up is to be finished. Note: To set the due date, you must also specify the startDateTime; otherwise, you will get a 400 Bad Request response.
    dueDateTime DateTimeTimeZoneable;
    // The status for follow-up for an item. Possible values are notFlagged, complete, and flagged.
    flagStatus *FollowupFlagStatus;
    // The date and time that the follow-up is to begin.
    startDateTime DateTimeTimeZoneable;
}
// NewFollowupFlag instantiates a new followupFlag and sets the default values.
func NewFollowupFlag()(*FollowupFlag) {
    m := &FollowupFlag{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateFollowupFlagFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateFollowupFlagFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewFollowupFlag(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FollowupFlag) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCompletedDateTime gets the completedDateTime property value. The date and time that the follow-up was finished.
func (m *FollowupFlag) GetCompletedDateTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.completedDateTime
    }
}
// GetDueDateTime gets the dueDateTime property value. The date and time that the follow up is to be finished. Note: To set the due date, you must also specify the startDateTime; otherwise, you will get a 400 Bad Request response.
func (m *FollowupFlag) GetDueDateTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.dueDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *FollowupFlag) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["completedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCompletedDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["dueDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDueDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    res["flagStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFollowupFlagStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFlagStatus(val.(*FollowupFlagStatus))
        }
        return nil
    }
    res["startDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateDateTimeTimeZoneFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDateTime(val.(DateTimeTimeZoneable))
        }
        return nil
    }
    return res
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
func (m *FollowupFlag) GetStartDateTime()(DateTimeTimeZoneable) {
    if m == nil {
        return nil
    } else {
        return m.startDateTime
    }
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
        cast := (*m.GetFlagStatus()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *FollowupFlag) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCompletedDateTime sets the completedDateTime property value. The date and time that the follow-up was finished.
func (m *FollowupFlag) SetCompletedDateTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.completedDateTime = value
    }
}
// SetDueDateTime sets the dueDateTime property value. The date and time that the follow up is to be finished. Note: To set the due date, you must also specify the startDateTime; otherwise, you will get a 400 Bad Request response.
func (m *FollowupFlag) SetDueDateTime(value DateTimeTimeZoneable)() {
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
func (m *FollowupFlag) SetStartDateTime(value DateTimeTimeZoneable)() {
    if m != nil {
        m.startDateTime = value
    }
}
