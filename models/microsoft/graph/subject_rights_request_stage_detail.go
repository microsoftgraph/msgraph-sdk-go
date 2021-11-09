package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
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
// Instantiates a new subjectRightsRequestStageDetail and sets the default values.
func NewSubjectRightsRequestStageDetail()(*SubjectRightsRequestStageDetail) {
    m := &SubjectRightsRequestStageDetail{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubjectRightsRequestStageDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the error property value. Describes the error, if any, for the current stage.
func (m *SubjectRightsRequestStageDetail) GetError()(*PublicError) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// Gets the stage property value. The stage of the subject rights request. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue.
func (m *SubjectRightsRequestStageDetail) GetStage()(*SubjectRightsRequestStage) {
    if m == nil {
        return nil
    } else {
        return m.stage
    }
}
// Gets the status property value. Status of the current stage. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
func (m *SubjectRightsRequestStageDetail) GetStatus()(*SubjectRightsRequestStageStatus) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// The deserialization information for the current model
func (m *SubjectRightsRequestStageDetail) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["error"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewPublicError() })
        if err != nil {
            return err
        }
        m.SetError(val.(*PublicError))
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
    res["status"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStageStatus)
        if err != nil {
            return err
        }
        cast := val.(SubjectRightsRequestStageStatus)
        m.SetStatus(&cast)
        return nil
    }
    return res
}
func (m *SubjectRightsRequestStageDetail) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *SubjectRightsRequestStageDetail) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the error property value. Describes the error, if any, for the current stage.
// Parameters:
//  - value : Value to set for the error property.
func (m *SubjectRightsRequestStageDetail) SetError(value *PublicError)() {
    m.error = value
}
// Sets the stage property value. The stage of the subject rights request. Possible values are: contentRetrieval, contentReview, generateReport, contentDeletion, caseResolved, unknownFutureValue.
// Parameters:
//  - value : Value to set for the stage property.
func (m *SubjectRightsRequestStageDetail) SetStage(value *SubjectRightsRequestStage)() {
    m.stage = value
}
// Sets the status property value. Status of the current stage. Possible values are: notStarted, current, completed, failed, unknownFutureValue.
// Parameters:
//  - value : Value to set for the status property.
func (m *SubjectRightsRequestStageDetail) SetStatus(value *SubjectRightsRequestStageStatus)() {
    m.status = value
}
