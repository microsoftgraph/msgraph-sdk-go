package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ModifiedProperty 
type ModifiedProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Name of property that was modified.
    displayName *string
    // New property value.
    newValue *string
    // The OdataType property
    odataType *string
    // Old property value.
    oldValue *string
}
// NewModifiedProperty instantiates a new modifiedProperty and sets the default values.
func NewModifiedProperty()(*ModifiedProperty) {
    m := &ModifiedProperty{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    odataTypeValue := "#microsoft.graph.modifiedProperty";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateModifiedPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateModifiedPropertyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewModifiedProperty(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ModifiedProperty) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisplayName gets the displayName property value. Name of property that was modified.
func (m *ModifiedProperty) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ModifiedProperty) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["newValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetNewValue(val)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["oldValue"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldValue(val)
        }
        return nil
    }
    return res
}
// GetNewValue gets the newValue property value. New property value.
func (m *ModifiedProperty) GetNewValue()(*string) {
    return m.newValue
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *ModifiedProperty) GetOdataType()(*string) {
    return m.odataType
}
// GetOldValue gets the oldValue property value. Old property value.
func (m *ModifiedProperty) GetOldValue()(*string) {
    return m.oldValue
}
// Serialize serializes information the current object
func (m *ModifiedProperty) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("newValue", m.GetNewValue())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("oldValue", m.GetOldValue())
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
func (m *ModifiedProperty) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisplayName sets the displayName property value. Name of property that was modified.
func (m *ModifiedProperty) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetNewValue sets the newValue property value. New property value.
func (m *ModifiedProperty) SetNewValue(value *string)() {
    m.newValue = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *ModifiedProperty) SetOdataType(value *string)() {
    m.odataType = value
}
// SetOldValue sets the oldValue property value. Old property value.
func (m *ModifiedProperty) SetOldValue(value *string)() {
    m.oldValue = value
}
