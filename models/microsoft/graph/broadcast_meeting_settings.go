package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new broadcastMeetingSettings and sets the default values.
func NewBroadcastMeetingSettings()(*BroadcastMeetingSettings) {
    m := &BroadcastMeetingSettings{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *BroadcastMeetingSettings) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowedAudience property value. Defines who can join the Teams live event. Possible values are listed in the following table.
func (m *BroadcastMeetingSettings) GetAllowedAudience()(*BroadcastMeetingAudience) {
    if m == nil {
        return nil
    } else {
        return m.allowedAudience
    }
}
// Gets the isAttendeeReportEnabled property value. Indicates whether attendee report is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) GetIsAttendeeReportEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAttendeeReportEnabled
    }
}
// Gets the isQuestionAndAnswerEnabled property value. Indicates whether Q&A is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) GetIsQuestionAndAnswerEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isQuestionAndAnswerEnabled
    }
}
// Gets the isRecordingEnabled property value. Indicates whether recording is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) GetIsRecordingEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRecordingEnabled
    }
}
// Gets the isVideoOnDemandEnabled property value. Indicates whether video on demand is enabled for this Teams live event. Default value is false.
func (m *BroadcastMeetingSettings) GetIsVideoOnDemandEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isVideoOnDemandEnabled
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *BroadcastMeetingSettings) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowedAudience property value. Defines who can join the Teams live event. Possible values are listed in the following table.
// Parameters:
//  - value : Value to set for the allowedAudience property.
func (m *BroadcastMeetingSettings) SetAllowedAudience(value *BroadcastMeetingAudience)() {
    m.allowedAudience = value
}
// Sets the isAttendeeReportEnabled property value. Indicates whether attendee report is enabled for this Teams live event. Default value is false.
// Parameters:
//  - value : Value to set for the isAttendeeReportEnabled property.
func (m *BroadcastMeetingSettings) SetIsAttendeeReportEnabled(value *bool)() {
    m.isAttendeeReportEnabled = value
}
// Sets the isQuestionAndAnswerEnabled property value. Indicates whether Q&A is enabled for this Teams live event. Default value is false.
// Parameters:
//  - value : Value to set for the isQuestionAndAnswerEnabled property.
func (m *BroadcastMeetingSettings) SetIsQuestionAndAnswerEnabled(value *bool)() {
    m.isQuestionAndAnswerEnabled = value
}
// Sets the isRecordingEnabled property value. Indicates whether recording is enabled for this Teams live event. Default value is false.
// Parameters:
//  - value : Value to set for the isRecordingEnabled property.
func (m *BroadcastMeetingSettings) SetIsRecordingEnabled(value *bool)() {
    m.isRecordingEnabled = value
}
// Sets the isVideoOnDemandEnabled property value. Indicates whether video on demand is enabled for this Teams live event. Default value is false.
// Parameters:
//  - value : Value to set for the isVideoOnDemandEnabled property.
func (m *BroadcastMeetingSettings) SetIsVideoOnDemandEnabled(value *bool)() {
    m.isVideoOnDemandEnabled = value
}
