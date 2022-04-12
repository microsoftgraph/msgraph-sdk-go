package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// PersonType 
type PersonType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The type of data source, such as Person.
    class *string;
    // The secondary type of data source, such as OrganizationUser.
    subclass *string;
}
// NewPersonType instantiates a new personType and sets the default values.
func NewPersonType()(*PersonType) {
    m := &PersonType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePersonTypeFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonTypeFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewPersonType(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PersonType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetClass gets the class property value. The type of data source, such as Person.
func (m *PersonType) GetClass()(*string) {
    if m == nil {
        return nil
    } else {
        return m.class
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonType) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["class"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetClass(val)
        }
        return nil
    }
    res["subclass"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubclass(val)
        }
        return nil
    }
    return res
}
// GetSubclass gets the subclass property value. The secondary type of data source, such as OrganizationUser.
func (m *PersonType) GetSubclass()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subclass
    }
}
// Serialize serializes information the current object
func (m *PersonType) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("class", m.GetClass())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subclass", m.GetSubclass())
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
func (m *PersonType) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetClass sets the class property value. The type of data source, such as Person.
func (m *PersonType) SetClass(value *string)() {
    if m != nil {
        m.class = value
    }
}
// SetSubclass sets the subclass property value. The secondary type of data source, such as OrganizationUser.
func (m *PersonType) SetSubclass(value *string)() {
    if m != nil {
        m.subclass = value
    }
}
