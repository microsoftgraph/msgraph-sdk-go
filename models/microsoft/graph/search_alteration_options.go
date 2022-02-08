package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// SearchAlterationOptions 
type SearchAlterationOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether spelling modifications are enabled. If enabled, user will get the search results for corrected query when there are no results for the original query with typos and get the spelling modification information in queryAlterationResponse property of the response. Optional.
    enableModification *bool;
    // Indicates whether spelling suggestions are enabled. If enabled, user will get the search results for original search query and suggesting spelling correction in queryAlterationResponse property of the response for typos in query. Optional.
    enableSuggestion *bool;
}
// NewSearchAlterationOptions instantiates a new searchAlterationOptions and sets the default values.
func NewSearchAlterationOptions()(*SearchAlterationOptions) {
    m := &SearchAlterationOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchAlterationOptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnableModification gets the enableModification property value. Indicates whether spelling modifications are enabled. If enabled, user will get the search results for corrected query when there are no results for the original query with typos and get the spelling modification information in queryAlterationResponse property of the response. Optional.
func (m *SearchAlterationOptions) GetEnableModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableModification
    }
}
// GetEnableSuggestion gets the enableSuggestion property value. Indicates whether spelling suggestions are enabled. If enabled, user will get the search results for original search query and suggesting spelling correction in queryAlterationResponse property of the response for typos in query. Optional.
func (m *SearchAlterationOptions) GetEnableSuggestion()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableSuggestion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchAlterationOptions) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["enableModification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableModification(val)
        }
        return nil
    }
    res["enableSuggestion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableSuggestion(val)
        }
        return nil
    }
    return res
}
func (m *SearchAlterationOptions) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *SearchAlterationOptions) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("enableModification", m.GetEnableModification())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("enableSuggestion", m.GetEnableSuggestion())
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
func (m *SearchAlterationOptions) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetEnableModification sets the enableModification property value. Indicates whether spelling modifications are enabled. If enabled, user will get the search results for corrected query when there are no results for the original query with typos and get the spelling modification information in queryAlterationResponse property of the response. Optional.
func (m *SearchAlterationOptions) SetEnableModification(value *bool)() {
    if m != nil {
        m.enableModification = value
    }
}
// SetEnableSuggestion sets the enableSuggestion property value. Indicates whether spelling suggestions are enabled. If enabled, user will get the search results for original search query and suggesting spelling correction in queryAlterationResponse property of the response for typos in query. Optional.
func (m *SearchAlterationOptions) SetEnableSuggestion(value *bool)() {
    if m != nil {
        m.enableSuggestion = value
    }
}
