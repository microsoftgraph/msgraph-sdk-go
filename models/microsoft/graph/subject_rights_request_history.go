package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type SubjectRightsRequestHistory struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Identity of the user who changed the  subject rights request.
    changedBy *IdentitySet;
    // Data and time when the entity was changed.
    eventDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The stage when the entity was changed. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue.
    stage *SubjectRightsRequestStage;
    // The status of the stage when the entity was changed. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
    stageStatus *SubjectRightsRequestStageStatus;
    // Type of history.
    type_escaped *string;
}
// Instantiates a new subjectRightsRequestHistory and sets the default values.
func NewSubjectRightsRequestHistory()(*SubjectRightsRequestHistory) {
    m := &SubjectRightsRequestHistory{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubjectRightsRequestHistory) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the changedBy property value. Identity of the user who changed the  subject rights request.
func (m *SubjectRightsRequestHistory) GetChangedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.changedBy
    }
}
// Gets the eventDateTime property value. Data and time when the entity was changed.
func (m *SubjectRightsRequestHistory) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
// Gets the stage property value. The stage when the entity was changed. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue.
func (m *SubjectRightsRequestHistory) GetStage()(*SubjectRightsRequestStage) {
    if m == nil {
        return nil
    } else {
        return m.stage
    }
}
// Gets the stageStatus property value. The status of the stage when the entity was changed. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
func (m *SubjectRightsRequestHistory) GetStageStatus()(*SubjectRightsRequestStageStatus) {
    if m == nil {
        return nil
    } else {
        return m.stageStatus
    }
}
// Gets the type_escaped property value. Type of history.
func (m *SubjectRightsRequestHistory) GetType_escaped()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// The deserialization information for the current model
func (m *SubjectRightsRequestHistory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["changedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        m.SetChangedBy(val.(*IdentitySet))
        return nil
    }
    res["eventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetEventDateTime(val)
        return nil
    }
    res["stage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStage)
        if err != nil {
            return err
        }
        cast := val.(SubjectRightsRequestStage)
        m.SetStage(&cast)
        return nil
    }
    res["stageStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStageStatus)
        if err != nil {
            return err
        }
        cast := val.(SubjectRightsRequestStageStatus)
        m.SetStageStatus(&cast)
        return nil
    }
    res["type_escaped"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetType_escaped(val)
        return nil
    }
    return res
}
func (m *SubjectRightsRequestHistory) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *SubjectRightsRequestHistory) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("changedBy", m.GetChangedBy())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("eventDateTime", m.GetEventDateTime())
        if err != nil {
            return err
        }
    }
    if m.GetStage() != nil {
        cast := m.GetStage().String()
        err := writer.WriteStringValue("stage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStageStatus() != nil {
        cast := m.GetStageStatus().String()
        err := writer.WriteStringValue("stageStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type_escaped", m.GetType_escaped())
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
func (m *SubjectRightsRequestHistory) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the changedBy property value. Identity of the user who changed the  subject rights request.
// Parameters:
//  - value : Value to set for the changedBy property.
func (m *SubjectRightsRequestHistory) SetChangedBy(value *IdentitySet)() {
    m.changedBy = value
}
// Sets the eventDateTime property value. Data and time when the entity was changed.
// Parameters:
//  - value : Value to set for the eventDateTime property.
func (m *SubjectRightsRequestHistory) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.eventDateTime = value
}
// Sets the stage property value. The stage when the entity was changed. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue.
// Parameters:
//  - value : Value to set for the stage property.
func (m *SubjectRightsRequestHistory) SetStage(value *SubjectRightsRequestStage)() {
    m.stage = value
}
// Sets the stageStatus property value. The status of the stage when the entity was changed. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
// Parameters:
//  - value : Value to set for the stageStatus property.
func (m *SubjectRightsRequestHistory) SetStageStatus(value *SubjectRightsRequestStageStatus)() {
    m.stageStatus = value
}
// Sets the type_escaped property value. Type of history.
// Parameters:
//  - value : Value to set for the type_escaped property.
func (m *SubjectRightsRequestHistory) SetType_escaped(value *string)() {
    m.type_escaped = value
}
