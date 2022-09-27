package getactivitiesbyinterval

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// GetActivitiesByIntervalResponse provides operations to call the getActivitiesByInterval method.
type GetActivitiesByIntervalResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The value property
    value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable
}
// NewGetActivitiesByIntervalResponse instantiates a new getActivitiesByIntervalResponse and sets the default values.
func NewGetActivitiesByIntervalResponse()(*GetActivitiesByIntervalResponse) {
    m := &GetActivitiesByIntervalResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetActivitiesByIntervalResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetActivitiesByIntervalResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetActivitiesByIntervalResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetActivitiesByIntervalResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetActivitiesByIntervalResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["value"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateItemActivityStatFromDiscriminatorValue , m.SetValue)
    return res
}
// GetValue gets the value property value. The value property
func (m *GetActivitiesByIntervalResponse) GetValue()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable) {
    return m.value
}
// Serialize serializes information the current object
func (m *GetActivitiesByIntervalResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
func (m *GetActivitiesByIntervalResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetValue sets the value property value. The value property
func (m *GetActivitiesByIntervalResponse) SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.ItemActivityStatable)() {
    m.value = value
}
