package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// WorkbooksItemCheckinPostRequestBody provides operations to call the checkin method.
type WorkbooksItemCheckinPostRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The checkInAs property
    checkInAs *string
    // The comment property
    comment *string
}
// NewWorkbooksItemCheckinPostRequestBody instantiates a new WorkbooksItemCheckinPostRequestBody and sets the default values.
func NewWorkbooksItemCheckinPostRequestBody()(*WorkbooksItemCheckinPostRequestBody) {
    m := &WorkbooksItemCheckinPostRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateWorkbooksItemCheckinPostRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateWorkbooksItemCheckinPostRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewWorkbooksItemCheckinPostRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *WorkbooksItemCheckinPostRequestBody) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetCheckInAs gets the checkInAs property value. The checkInAs property
func (m *WorkbooksItemCheckinPostRequestBody) GetCheckInAs()(*string) {
    return m.checkInAs
}
// GetComment gets the comment property value. The comment property
func (m *WorkbooksItemCheckinPostRequestBody) GetComment()(*string) {
    return m.comment
}
// GetFieldDeserializers the deserialization information for the current model
func (m *WorkbooksItemCheckinPostRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["checkInAs"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCheckInAs(val)
        }
        return nil
    }
    res["comment"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetComment(val)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *WorkbooksItemCheckinPostRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("checkInAs", m.GetCheckInAs())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("comment", m.GetComment())
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
func (m *WorkbooksItemCheckinPostRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCheckInAs sets the checkInAs property value. The checkInAs property
func (m *WorkbooksItemCheckinPostRequestBody) SetCheckInAs(value *string)() {
    m.checkInAs = value
}
// SetComment sets the comment property value. The comment property
func (m *WorkbooksItemCheckinPostRequestBody) SetComment(value *string)() {
    m.comment = value
}
