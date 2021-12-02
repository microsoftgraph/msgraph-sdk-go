package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SubjectRightsRequestStageDetail 
type SubjectRightsRequestStageDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Describes the error, if any, for the current stage.
    error *PublicError;
    // The stage of the subject rights request. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue.
    stage *SubjectRightsRequestStage;
    // Status of the current stage. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
    status *SubjectRightsRequestStageStatus;
}
// NewSubjectRightsRequestStageDetail instantiates a new subjectRightsRequestStageDetail and sets the default values.
func NewSubjectRightsRequestStageDetail()(*SubjectRightsRequestStageDetail) {
    m := &SubjectRightsRequestStageDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubjectRightsRequestStageDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetError gets the error property value. Describes the error, if any, for the current stage.
func (m *SubjectRightsRequestStageDetail) GetError()(*PublicError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetStage gets the stage property value. The stage of the subject rights request. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue.
func (m *SubjectRightsRequestStageDetail) GetStage()(*SubjectRightsRequestStage) {
    if m == nil {
        return nil
    } else {
        return m.stage
    }
}
// GetStatus gets the status property value. Status of the current stage. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
func (m *SubjectRightsRequestStageDetail) GetStatus()(*SubjectRightsRequestStageStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubjectRightsRequestStageDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublicError() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(*PublicError))
        }
        return nil
    }
    res["stage"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStage)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SubjectRightsRequestStage)
            m.SetStage(&cast)
        }
        return nil
    }
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStageStatus)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(SubjectRightsRequestStageStatus)
            m.SetStatus(&cast)
        }
        return nil
    }
    return res
}
func (m *SubjectRightsRequestStageDetail) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SubjectRightsRequestStageDetail) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("error", m.GetError())
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
    if m.GetStatus() != nil {
        cast := m.GetStatus().String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *SubjectRightsRequestStageDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetError sets the error property value. Describes the error, if any, for the current stage.
func (m *SubjectRightsRequestStageDetail) SetError(value *PublicError)() {
    if m != nil {
        m.error = value
    }
}
// SetStage sets the stage property value. The stage of the subject rights request. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue.
func (m *SubjectRightsRequestStageDetail) SetStage(value *SubjectRightsRequestStage)() {
    if m != nil {
        m.stage = value
    }
}
// SetStatus sets the status property value. Status of the current stage. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
func (m *SubjectRightsRequestStageDetail) SetStatus(value *SubjectRightsRequestStageStatus)() {
    if m != nil {
        m.status = value
    }
}
