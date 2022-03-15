package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// TermColumn provides operations to manage the drive singleton.
type TermColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies whether the column will allow more than one value.
    allowMultipleValues *bool;
    // Specifies whether to display the entire term path or only the term label.
    showFullyQualifiedName *bool;
}
// NewTermColumn instantiates a new termColumn and sets the default values.
func NewTermColumn()(*TermColumn) {
    m := &TermColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateTermColumnFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateTermColumnFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewTermColumn(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TermColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAllowMultipleValues gets the allowMultipleValues property value. Specifies whether the column will allow more than one value.
func (m *TermColumn) GetAllowMultipleValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowMultipleValues
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *TermColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowMultipleValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAllowMultipleValues(val)
        }
        return nil
    }
    res["showFullyQualifiedName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShowFullyQualifiedName(val)
        }
        return nil
    }
    return res
}
// GetShowFullyQualifiedName gets the showFullyQualifiedName property value. Specifies whether to display the entire term path or only the term label.
func (m *TermColumn) GetShowFullyQualifiedName()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showFullyQualifiedName
    }
}
func (m *TermColumn) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *TermColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleValues", m.GetAllowMultipleValues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("showFullyQualifiedName", m.GetShowFullyQualifiedName())
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
func (m *TermColumn) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAllowMultipleValues sets the allowMultipleValues property value. Specifies whether the column will allow more than one value.
func (m *TermColumn) SetAllowMultipleValues(value *bool)() {
    if m != nil {
        m.allowMultipleValues = value
    }
}
// SetShowFullyQualifiedName sets the showFullyQualifiedName property value. Specifies whether to display the entire term path or only the term label.
func (m *TermColumn) SetShowFullyQualifiedName(value *bool)() {
    if m != nil {
        m.showFullyQualifiedName = value
    }
}
