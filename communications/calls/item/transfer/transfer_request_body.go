package transfer

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// TransferRequestBody provides operations to call the transfer method.
type TransferRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The transferee property
    transferee iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParticipantInfoable
    // The transferTarget property
    transferTarget iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InvitationParticipantInfoable
}
// NewTransferRequestBody instantiates a new transferRequestBody and sets the default values.
func NewTransferRequestBody()(*TransferRequestBody) {
    m := &TransferRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTransferRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTransferRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewTransferRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TransferRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TransferRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["transferee"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransferee(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParticipantInfoable))
        }
        return nil
    }
    res["transferTarget"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateInvitationParticipantInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTransferTarget(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InvitationParticipantInfoable))
        }
        return nil
    }
    return res
}
// GetTransferee gets the transferee property value. The transferee property
func (m *TransferRequestBody) GetTransferee()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParticipantInfoable) {
    if m == nil {
        return nil
    } else {
        return m.transferee
    }
}
// GetTransferTarget gets the transferTarget property value. The transferTarget property
func (m *TransferRequestBody) GetTransferTarget()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InvitationParticipantInfoable) {
    if m == nil {
        return nil
    } else {
        return m.transferTarget
    }
}
// Serialize serializes information the current object
func (m *TransferRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("transferee", m.GetTransferee())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("transferTarget", m.GetTransferTarget())
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
func (m *TransferRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetTransferee sets the transferee property value. The transferee property
func (m *TransferRequestBody) SetTransferee(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ParticipantInfoable)() {
    if m != nil {
        m.transferee = value
    }
}
// SetTransferTarget sets the transferTarget property value. The transferTarget property
func (m *TransferRequestBody) SetTransferTarget(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.InvitationParticipantInfoable)() {
    if m != nil {
        m.transferTarget = value
    }
}
