package getpresencesbyuserid

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// GetPresencesByUserIdResponse provides operations to call the getPresencesByUserId method.
type GetPresencesByUserIdResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The value property
    value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Presenceable
}
// NewGetPresencesByUserIdResponse instantiates a new getPresencesByUserIdResponse and sets the default values.
func NewGetPresencesByUserIdResponse()(*GetPresencesByUserIdResponse) {
    m := &GetPresencesByUserIdResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetPresencesByUserIdResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetPresencesByUserIdResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetPresencesByUserIdResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetPresencesByUserIdResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetPresencesByUserIdResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePresenceFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *GetPresencesByUserIdResponse) GetValue()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Presenceable) {
    return m.value
}
// Serialize serializes information the current object
func (m *GetPresencesByUserIdResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetValue() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetValue())
        err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *GetPresencesByUserIdResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *GetPresencesByUserIdResponse) SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.Presenceable)() {
    m.value = value
}
