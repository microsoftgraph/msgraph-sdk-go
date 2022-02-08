package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SubjectRightsRequestHistory 
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
// NewSubjectRightsRequestHistory instantiates a new subjectRightsRequestHistory and sets the default values.
func NewSubjectRightsRequestHistory()(*SubjectRightsRequestHistory) {
    m := &SubjectRightsRequestHistory{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubjectRightsRequestHistory) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetChangedBy gets the changedBy property value. Identity of the user who changed the  subject rights request.
func (m *SubjectRightsRequestHistory) GetChangedBy()(*IdentitySet) {
    if m == nil {
        return nil
    } else {
        return m.changedBy
    }
}
// GetEventDateTime gets the eventDateTime property value. Data and time when the entity was changed.
func (m *SubjectRightsRequestHistory) GetEventDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.eventDateTime
    }
}
// GetStage gets the stage property value. The stage when the entity was changed. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue.
func (m *SubjectRightsRequestHistory) GetStage()(*SubjectRightsRequestStage) {
    if m == nil {
        return nil
    } else {
        return m.stage
    }
}
// GetStageStatus gets the stageStatus property value. The status of the stage when the entity was changed. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
func (m *SubjectRightsRequestHistory) GetStageStatus()(*SubjectRightsRequestStageStatus) {
    if m == nil {
        return nil
    } else {
        return m.stageStatus
    }
}
// GetType gets the type property value. Type of history.
func (m *SubjectRightsRequestHistory) GetType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.type_escaped
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubjectRightsRequestHistory) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["changedBy"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentitySet() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChangedBy(val.(*IdentitySet))
        }
        return nil
    }
    res["eventDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEventDateTime(val)
        }
        return nil
    }
    res["stage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStage(val.(*SubjectRightsRequestStage))
        }
        return nil
    }
    res["stageStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStageStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStageStatus(val.(*SubjectRightsRequestStageStatus))
        }
        return nil
    }
    res["type"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetType(val)
        }
        return nil
    }
    return res
}
func (m *SubjectRightsRequestHistory) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetStage()).String()
        err := writer.WriteStringValue("stage", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetStageStatus() != nil {
        cast := (*m.GetStageStatus()).String()
        err := writer.WriteStringValue("stageStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("type", m.GetType())
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
func (m *SubjectRightsRequestHistory) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetChangedBy sets the changedBy property value. Identity of the user who changed the  subject rights request.
func (m *SubjectRightsRequestHistory) SetChangedBy(value *IdentitySet)() {
    if m != nil {
        m.changedBy = value
    }
}
// SetEventDateTime sets the eventDateTime property value. Data and time when the entity was changed.
func (m *SubjectRightsRequestHistory) SetEventDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.eventDateTime = value
    }
}
// SetStage sets the stage property value. The stage when the entity was changed. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue.
func (m *SubjectRightsRequestHistory) SetStage(value *SubjectRightsRequestStage)() {
    if m != nil {
        m.stage = value
    }
}
// SetStageStatus sets the stageStatus property value. The status of the stage when the entity was changed. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
func (m *SubjectRightsRequestHistory) SetStageStatus(value *SubjectRightsRequestStageStatus)() {
    if m != nil {
        m.stageStatus = value
    }
}
// SetType sets the type property value. Type of history.
func (m *SubjectRightsRequestHistory) SetType(value *string)() {
    if m != nil {
        m.type_escaped = value
    }
}
