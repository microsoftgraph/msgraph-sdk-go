package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SubjectRightsRequestStageDetail 
type SubjectRightsRequestStageDetail struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Describes the error, if any, for the current stage.
    error PublicErrorable;
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
// CreateSubjectRightsRequestStageDetailFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSubjectRightsRequestStageDetailFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSubjectRightsRequestStageDetail(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubjectRightsRequestStageDetail) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetError gets the error property value. Describes the error, if any, for the current stage.
func (m *SubjectRightsRequestStageDetail) GetError()(PublicErrorable) {
    if m == nil {
        return nil
    } else {
        return m.error
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SubjectRightsRequestStageDetail) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["error"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreatePublicErrorFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetError(val.(PublicErrorable))
        }
        return nil
    }
    res["stage"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStage)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStage(val.(*SubjectRightsRequestStage))
        }
        return nil
    }
    res["status"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseSubjectRightsRequestStageStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*SubjectRightsRequestStageStatus))
        }
        return nil
    }
    return res
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
// Serialize serializes information the current object
func (m *SubjectRightsRequestStageDetail) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("error", m.GetError())
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
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SubjectRightsRequestStageDetail) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetError sets the error property value. Describes the error, if any, for the current stage.
func (m *SubjectRightsRequestStageDetail) SetError(value PublicErrorable)() {
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
