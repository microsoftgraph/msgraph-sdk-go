package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PersonOrGroupColumn provides operations to manage the educationRoot singleton.
type PersonOrGroupColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether multiple values can be selected from the source.
    allowMultipleSelection *bool;
    // Whether to allow selection of people only, or people and groups. Must be one of peopleAndGroups or peopleOnly.
    chooseFromType *string;
    // How to display the information about the person or group chosen. See below.
    displayAs *string;
}
// NewPersonOrGroupColumn instantiates a new personOrGroupColumn and sets the default values.
func NewPersonOrGroupColumn()(*PersonOrGroupColumn) {
    m := &PersonOrGroupColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePersonOrGroupColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePersonOrGroupColumnFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPersonOrGroupColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PersonOrGroupColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowMultipleSelection gets the allowMultipleSelection property value. Indicates whether multiple values can be selected from the source.
func (m *PersonOrGroupColumn) GetAllowMultipleSelection()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowMultipleSelection
    }
}
// GetChooseFromType gets the chooseFromType property value. Whether to allow selection of people only, or people and groups. Must be one of peopleAndGroups or peopleOnly.
func (m *PersonOrGroupColumn) GetChooseFromType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.chooseFromType
    }
}
// GetDisplayAs gets the displayAs property value. How to display the information about the person or group chosen. See below.
func (m *PersonOrGroupColumn) GetDisplayAs()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayAs
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PersonOrGroupColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowMultipleSelection"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMultipleSelection(val)
        }
        return nil
    }
    res["chooseFromType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetChooseFromType(val)
        }
        return nil
    }
    res["displayAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayAs(val)
        }
        return nil
    }
    return res
}
func (m *PersonOrGroupColumn) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PersonOrGroupColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleSelection", m.GetAllowMultipleSelection())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("chooseFromType", m.GetChooseFromType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayAs", m.GetDisplayAs())
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
func (m *PersonOrGroupColumn) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowMultipleSelection sets the allowMultipleSelection property value. Indicates whether multiple values can be selected from the source.
func (m *PersonOrGroupColumn) SetAllowMultipleSelection(value *bool)() {
    if m != nil {
        m.allowMultipleSelection = value
    }
}
// SetChooseFromType sets the chooseFromType property value. Whether to allow selection of people only, or people and groups. Must be one of peopleAndGroups or peopleOnly.
func (m *PersonOrGroupColumn) SetChooseFromType(value *string)() {
    if m != nil {
        m.chooseFromType = value
    }
}
// SetDisplayAs sets the displayAs property value. How to display the information about the person or group chosen. See below.
func (m *PersonOrGroupColumn) SetDisplayAs(value *string)() {
    if m != nil {
        m.displayAs = value
    }
}
