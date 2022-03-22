package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SortProperty 
type SortProperty struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // True if the sort order is descending. Default is false, with the sort order as ascending. Optional.
    isDescending *bool;
    // The name of the property to sort on. Required.
    name *string;
}
// NewSortProperty instantiates a new sortProperty and sets the default values.
func NewSortProperty()(*SortProperty) {
    m := &SortProperty{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSortPropertyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSortPropertyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewSortProperty(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SortProperty) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SortProperty) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isDescending"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsDescending(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    return res
}
// GetIsDescending gets the isDescending property value. True if the sort order is descending. Default is false, with the sort order as ascending. Optional.
func (m *SortProperty) GetIsDescending()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isDescending
    }
}
// GetName gets the name property value. The name of the property to sort on. Required.
func (m *SortProperty) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Serialize serializes information the current object
func (m *SortProperty) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isDescending", m.GetIsDescending())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
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
func (m *SortProperty) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetIsDescending sets the isDescending property value. True if the sort order is descending. Default is false, with the sort order as ascending. Optional.
func (m *SortProperty) SetIsDescending(value *bool)() {
    if m != nil {
        m.isDescending = value
    }
}
// SetName sets the name property value. The name of the property to sort on. Required.
func (m *SortProperty) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
