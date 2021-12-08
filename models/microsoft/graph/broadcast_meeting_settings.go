package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// BroadcastMeetingSettings 
type BroadcastMeetingSettings struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Defines who can join the Teams live event. Possible values are listed in the following table.
    allowedAudience *BroadcastMeetingAudience;
    // Indicates whether attendee report is enabled for this Teams live event. Default value is false.
    isAttendeeReportEnabled *bool;
    // Indicates whether Q&A is enabled for this Teams live event. Default value is false.
    isQuestionAndAnswerEnabled *bool;
    // Indicates whether recording is enabled for this Teams live event. Default value is false.
    isRecordingEnabled *bool;
    // Indicates whether video on demand is enabled for this Teams live event. Default value is false.
    isVideoOnDemandEnabled *bool;
}
// NewBroadcastMeetingSettings instantiates a new broadcastMeetingSettings and sets the default values.
func NewBroadcastMeetingSettings()(*BroadcastMeetingSettings) {
    m := &BroadcastMeetingSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BroadcastMeetingSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowedAudience gets the allowedAudience property value. Defines who can join the Teams live event. Possible values are listed in the following table.
func (m *BroadcastMeetingSettings) GetAllowedAudience()(*BroadcastMeetingAudience) {
    if m == nil {
        return nil
    } else {
        return m.allowedAudience
    }
}
// GetIsAttendeeReportEnabled gets the isAttendeeReportEnabled property value. Indicates whether attendee report is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) GetIsAttendeeReportEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAttendeeReportEnabled
    }
}
// GetIsQuestionAndAnswerEnabled gets the isQuestionAndAnswerEnabled property value. Indicates whether Q&A is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) GetIsQuestionAndAnswerEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isQuestionAndAnswerEnabled
    }
}
// GetIsRecordingEnabled gets the isRecordingEnabled property value. Indicates whether recording is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) GetIsRecordingEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRecordingEnabled
    }
}
// GetIsVideoOnDemandEnabled gets the isVideoOnDemandEnabled property value. Indicates whether video on demand is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) GetIsVideoOnDemandEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVideoOnDemandEnabled
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *BroadcastMeetingSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowedAudience"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseBroadcastMeetingAudience)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(BroadcastMeetingAudience)
            m.SetAllowedAudience(&cast)
        }
        return nil
    }
    res["isAttendeeReportEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsAttendeeReportEnabled(val)
        }
        return nil
    }
    res["isQuestionAndAnswerEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsQuestionAndAnswerEnabled(val)
        }
        return nil
    }
    res["isRecordingEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsRecordingEnabled(val)
        }
        return nil
    }
    res["isVideoOnDemandEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsVideoOnDemandEnabled(val)
        }
        return nil
    }
    return res
}
func (m *BroadcastMeetingSettings) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *BroadcastMeetingSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetAllowedAudience() != nil {
        cast := m.GetAllowedAudience().String()
        err := writer.WriteStringValue("allowedAudience", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isAttendeeReportEnabled", m.GetIsAttendeeReportEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isQuestionAndAnswerEnabled", m.GetIsQuestionAndAnswerEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRecordingEnabled", m.GetIsRecordingEnabled())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isVideoOnDemandEnabled", m.GetIsVideoOnDemandEnabled())
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
func (m *BroadcastMeetingSettings) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowedAudience sets the allowedAudience property value. Defines who can join the Teams live event. Possible values are listed in the following table.
func (m *BroadcastMeetingSettings) SetAllowedAudience(value *BroadcastMeetingAudience)() {
    if m != nil {
        m.allowedAudience = value
    }
}
// SetIsAttendeeReportEnabled sets the isAttendeeReportEnabled property value. Indicates whether attendee report is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) SetIsAttendeeReportEnabled(value *bool)() {
    if m != nil {
        m.isAttendeeReportEnabled = value
    }
}
// SetIsQuestionAndAnswerEnabled sets the isQuestionAndAnswerEnabled property value. Indicates whether Q&A is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) SetIsQuestionAndAnswerEnabled(value *bool)() {
    if m != nil {
        m.isQuestionAndAnswerEnabled = value
    }
}
// SetIsRecordingEnabled sets the isRecordingEnabled property value. Indicates whether recording is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) SetIsRecordingEnabled(value *bool)() {
    if m != nil {
        m.isRecordingEnabled = value
    }
}
// SetIsVideoOnDemandEnabled sets the isVideoOnDemandEnabled property value. Indicates whether video on demand is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) SetIsVideoOnDemandEnabled(value *bool)() {
    if m != nil {
        m.isVideoOnDemandEnabled = value
    }
}
