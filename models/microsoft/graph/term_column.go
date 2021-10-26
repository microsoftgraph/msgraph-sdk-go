package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TermColumn struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Specifies whether the column will allow more than one value.
    allowMultipleValues *bool;
    // 
    parentTerm *Term;
    // Specifies whether to display the entire term path or only the term label.
    showFullyQualifiedName *bool;
    // 
    termSet *Set;
}
// Instantiates a new termColumn and sets the default values.
func NewTermColumn()(*TermColumn) {
    m := &TermColumn{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *TermColumn) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the allowMultipleValues property value. Specifies whether the column will allow more than one value.
func (m *TermColumn) GetAllowMultipleValues()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.allowMultipleValues
    }
}
// Gets the parentTerm property value. 
func (m *TermColumn) GetParentTerm()(*Term) {
    if m == nil {
        return nil
    } else {
        return m.parentTerm
    }
}
// Gets the showFullyQualifiedName property value. Specifies whether to display the entire term path or only the term label.
func (m *TermColumn) GetShowFullyQualifiedName()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.showFullyQualifiedName
    }
}
// Gets the termSet property value. 
func (m *TermColumn) GetTermSet()(*Set) {
    if m == nil {
        return nil
    } else {
        return m.termSet
    }
}
// The deserialization information for the current model
func (m *TermColumn) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["allowMultipleValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetAllowMultipleValues(val)
        return nil
    }
    res["parentTerm"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTerm() })
        if err != nil {
            return err
        }
        m.SetParentTerm(val.(*Term))
        return nil
    }
    res["showFullyQualifiedName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetShowFullyQualifiedName(val)
        return nil
    }
    res["termSet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSet() })
        if err != nil {
            return err
        }
        m.SetTermSet(val.(*Set))
        return nil
    }
    return res
}
func (m *TermColumn) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TermColumn) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("allowMultipleValues", m.GetAllowMultipleValues())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("parentTerm", m.GetParentTerm())
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
        err := writer.WriteObjectValue("termSet", m.GetTermSet())
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
func (m *TermColumn) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the allowMultipleValues property value. Specifies whether the column will allow more than one value.
// Parameters:
//  - value : Value to set for the allowMultipleValues property.
func (m *TermColumn) SetAllowMultipleValues(value *bool)() {
    m.allowMultipleValues = value
}
// Sets the parentTerm property value. 
// Parameters:
//  - value : Value to set for the parentTerm property.
func (m *TermColumn) SetParentTerm(value *Term)() {
    m.parentTerm = value
}
// Sets the showFullyQualifiedName property value. Specifies whether to display the entire term path or only the term label.
// Parameters:
//  - value : Value to set for the showFullyQualifiedName property.
func (m *TermColumn) SetShowFullyQualifiedName(value *bool)() {
    m.showFullyQualifiedName = value
}
// Sets the termSet property value. 
// Parameters:
//  - value : Value to set for the termSet property.
func (m *TermColumn) SetTermSet(value *Set)() {
    m.termSet = value
}
