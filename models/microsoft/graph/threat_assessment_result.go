package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ThreatAssessmentResult struct {
    Entity
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    message *string;
    resultType *ThreatAssessmentResultType;
}
func NewThreatAssessmentResult()(*ThreatAssessmentResult) {
    m := &ThreatAssessmentResult{
        Entity: *NewEntity(),
    }
    return m
}
func (m *ThreatAssessmentResult) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *ThreatAssessmentResult) GetMessage()(*string) {
    if m == nil {
        return nil
    } else {
        return m.message
    }
}
func (m *ThreatAssessmentResult) GetResultType()(*ThreatAssessmentResultType) {
    if m == nil {
        return nil
    } else {
        return m.resultType
    }
}
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
func (m *ThreatAssessmentResult) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *ThreatAssessmentResult) SetMessage(value *string)() {
    m.message = value
}
func (m *ThreatAssessmentResult) SetResultType(value *ThreatAssessmentResultType)() {
    m.resultType = value
}
