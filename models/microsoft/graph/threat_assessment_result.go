package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ThreatAssessmentResult struct {
    Entity
    // The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The result message for each threat assessment.
    message *string;
    // The threat assessment result type. Possible values are: checkPolicy, rescan.
    resultType *ThreatAssessmentResultType;
}
// Instantiates a new threatAssessmentResult and sets the default values.
func NewThreatAssessmentResult()(*ThreatAssessmentResult) {
    m := &ThreatAssessmentResult{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *ThreatAssessmentResult) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the message property value. The result message for each threat assessment.
func (m *ThreatAssessmentResult) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
// Gets the resultType property value. The threat assessment result type. Possible values are: checkPolicy, rescan.
func (m *ThreatAssessmentResult) GetResultType()(*ThreatAssessmentResultType) {
    if m == nil {
        return nil
    } else {
        return m.resultType
    }
}
// The deserialization information for the current model
func (m *ThreatAssessmentResult) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["message"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMessage(val)
        return nil
    }
    res["resultType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseThreatAssessmentResultType)
        if err != nil {
            return err
        }
        cast := val.(ThreatAssessmentResultType)
        m.SetResultType(&cast)
        return nil
    }
    return res
}
func (m *ThreatAssessmentResult) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ThreatAssessmentResult) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("message", m.GetMessage())
        if err != nil {
            return err
        }
    }
    if m.GetResultType() != nil {
        cast := m.GetResultType().String()
        err = writer.WriteStringValue("resultType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the createdDateTime property value. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *ThreatAssessmentResult) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the message property value. The result message for each threat assessment.
// Parameters:
//  - value : Value to set for the message property.
func (m *ThreatAssessmentResult) SetMessage(value *string)() {
    m.message = value
}
// Sets the resultType property value. The threat assessment result type. Possible values are: checkPolicy, rescan.
// Parameters:
//  - value : Value to set for the resultType property.
func (m *ThreatAssessmentResult) SetResultType(value *ThreatAssessmentResultType)() {
    m.resultType = value
}
