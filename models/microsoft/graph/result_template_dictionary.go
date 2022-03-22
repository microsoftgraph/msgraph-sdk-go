package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ResultTemplateDictionary 
type ResultTemplateDictionary struct {
    Dictionary
}
// NewResultTemplateDictionary instantiates a new resultTemplateDictionary and sets the default values.
func NewResultTemplateDictionary()(*ResultTemplateDictionary) {
    m := &ResultTemplateDictionary{
        Dictionary: *NewDictionary(),
    }
    return m
}
// CreateResultTemplateDictionaryFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateResultTemplateDictionaryFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewResultTemplateDictionary(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResultTemplateDictionary) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Dictionary.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *ResultTemplateDictionary) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Dictionary.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
