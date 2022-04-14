package onenotepatchcontent

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// OnenotePatchContentRequestBody provides operations to call the onenotePatchContent method.
type OnenotePatchContentRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The commands property
    commands []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePatchContentCommandable
}
// NewOnenotePatchContentRequestBody instantiates a new onenotePatchContentRequestBody and sets the default values.
func NewOnenotePatchContentRequestBody()(*OnenotePatchContentRequestBody) {
    m := &OnenotePatchContentRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateOnenotePatchContentRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateOnenotePatchContentRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewOnenotePatchContentRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *OnenotePatchContentRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCommands gets the commands property value. The commands property
func (m *OnenotePatchContentRequestBody) GetCommands()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePatchContentCommandable) {
    if m == nil {
        return nil
    } else {
        return m.commands
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *OnenotePatchContentRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["commands"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateOnenotePatchContentCommandFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePatchContentCommandable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePatchContentCommandable)
            }
            m.SetCommands(res)
        }
        return nil
    }
    return res
}
// Serialize serializes information the current object
func (m *OnenotePatchContentRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetCommands() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetCommands()))
        for i, v := range m.GetCommands() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("commands", cast)
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
func (m *OnenotePatchContentRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCommands sets the commands property value. The commands property
func (m *OnenotePatchContentRequestBody) SetCommands(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.OnenotePatchContentCommandable)() {
    if m != nil {
        m.commands = value
    }
}
