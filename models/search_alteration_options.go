package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SearchAlterationOptions 
type SearchAlterationOptions struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether spelling modifications are enabled. If enabled, the user will get the search results for the corrected query in case of no results for the original query with typos. The response will also include the spelling modification information in the queryAlterationResponse property. Optional.
    enableModification *bool;
    // Indicates whether spelling suggestions are enabled. If enabled, the user will get the search results for the original search query and suggestions for spelling correction in the queryAlterationResponse property of the response for the typos in the query. Optional.
    enableSuggestion *bool;
}
// NewSearchAlterationOptions instantiates a new searchAlterationOptions and sets the default values.
func NewSearchAlterationOptions()(*SearchAlterationOptions) {
    m := &SearchAlterationOptions{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateSearchAlterationOptionsFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSearchAlterationOptionsFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSearchAlterationOptions(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *SearchAlterationOptions) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetEnableModification gets the enableModification property value. Indicates whether spelling modifications are enabled. If enabled, the user will get the search results for the corrected query in case of no results for the original query with typos. The response will also include the spelling modification information in the queryAlterationResponse property. Optional.
func (m *SearchAlterationOptions) GetEnableModification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableModification
    }
}
// GetEnableSuggestion gets the enableSuggestion property value. Indicates whether spelling suggestions are enabled. If enabled, the user will get the search results for the original search query and suggestions for spelling correction in the queryAlterationResponse property of the response for the typos in the query. Optional.
func (m *SearchAlterationOptions) GetEnableSuggestion()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enableSuggestion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SearchAlterationOptions) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["enableModification"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnableModification(val)
        }
        return nil
    }
    res["enableSuggestion"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// Serialize serializes information the current object
func (m *SearchAlterationOptions) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetEnableModification sets the enableModification property value. Indicates whether spelling modifications are enabled. If enabled, the user will get the search results for the corrected query in case of no results for the original query with typos. The response will also include the spelling modification information in the queryAlterationResponse property. Optional.
func (m *SearchAlterationOptions) SetEnableModification(value *bool)() {
    if m != nil {
        m.enableModification = value
    }
}
// SetEnableSuggestion sets the enableSuggestion property value. Indicates whether spelling suggestions are enabled. If enabled, the user will get the search results for the original search query and suggestions for spelling correction in the queryAlterationResponse property of the response for the typos in the query. Optional.
func (m *SearchAlterationOptions) SetEnableSuggestion(value *bool)() {
    if m != nil {
        m.enableSuggestion = value
    }
}
