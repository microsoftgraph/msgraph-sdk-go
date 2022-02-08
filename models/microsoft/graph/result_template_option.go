package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ResultTemplateOption 
type ResultTemplateOption struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether search display layouts are enabled. If enabled, the user will get the result template to render the search results content in the resultTemplates property of the response. The result template is based on Adaptive Cards. This property is optional.
    enableResultTemplate *bool;
}
// NewResultTemplateOption instantiates a new resultTemplateOption and sets the default values.
func NewResultTemplateOption()(*ResultTemplateOption) {
    m := &ResultTemplateOption{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ResultTemplateOption) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnableResultTemplate gets the enableResultTemplate property value. Indicates whether search display layouts are enabled. If enabled, the user will get the result template to render the search results content in the resultTemplates property of the response. The result template is based on Adaptive Cards. This property is optional.
func (m *ResultTemplateOption) GetEnableResultTemplate()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableResultTemplate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ResultTemplateOption) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enableResultTemplate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableResultTemplate(val)
        }
        return nil
    }
    return res
}
func (m *ResultTemplateOption) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ResultTemplateOption) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enableResultTemplate", m.GetEnableResultTemplate())
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
func (m *ResultTemplateOption) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnableResultTemplate sets the enableResultTemplate property value. Indicates whether search display layouts are enabled. If enabled, the user will get the result template to render the search results content in the resultTemplates property of the response. The result template is based on Adaptive Cards. This property is optional.
func (m *ResultTemplateOption) SetEnableResultTemplate(value *bool)() {
    if m != nil {
        m.enableResultTemplate = value
    }
}
