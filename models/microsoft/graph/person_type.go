package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type PersonType struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The type of data source, such as Person.
    class *string;
    // The secondary type of data source, such as OrganizationUser.
    subclass *string;
}
// Instantiates a new personType and sets the default values.
func NewPersonType()(*PersonType) {
    m := &PersonType{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PersonType) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the class property value. The type of data source, such as Person.
func (m *PersonType) GetClass()(*string) {
    if m == nil {
        return nil
    } else {
        return m.class
    }
}
// Gets the subclass property value. The secondary type of data source, such as OrganizationUser.
func (m *PersonType) GetSubclass()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subclass
    }
}
// The deserialization information for the current model
func (m *PersonType) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["class"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetClass(val)
        return nil
    }
    res["subclass"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSubclass(val)
        return nil
    }
    return res
}
func (m *PersonType) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *PersonType) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *PersonType) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the class property value. The type of data source, such as Person.
// Parameters:
//  - value : Value to set for the class property.
func (m *PersonType) SetClass(value *string)() {
    m.class = value
}
// Sets the subclass property value. The secondary type of data source, such as OrganizationUser.
// Parameters:
//  - value : Value to set for the subclass property.
func (m *PersonType) SetSubclass(value *string)() {
    m.subclass = value
}
