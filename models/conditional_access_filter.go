package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ConditionalAccessFilter 
type ConditionalAccessFilter struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // Mode to use for the filter. Possible values are include or exclude.
    mode *FilterMode
    // Rule syntax is similar to that used for membership rules for groups in Azure Active Directory. For details, see rules with multiple expressions
    rule *string
}
// NewConditionalAccessFilter instantiates a new conditionalAccessFilter and sets the default values.
func NewConditionalAccessFilter()(*ConditionalAccessFilter) {
    m := &ConditionalAccessFilter{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateConditionalAccessFilterFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateConditionalAccessFilterFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewConditionalAccessFilter(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessFilter) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessFilter) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["mode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseFilterMode)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMode(val.(*FilterMode))
        }
        return nil
    }
    res["rule"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRule(val)
        }
        return nil
    }
    return res
}
// GetMode gets the mode property value. Mode to use for the filter. Possible values are include or exclude.
func (m *ConditionalAccessFilter) GetMode()(*FilterMode) {
    if m == nil {
        return nil
    } else {
        return m.mode
    }
}
// GetRule gets the rule property value. Rule syntax is similar to that used for membership rules for groups in Azure Active Directory. For details, see rules with multiple expressions
func (m *ConditionalAccessFilter) GetRule()(*string) {
    if m == nil {
        return nil
    } else {
        return m.rule
    }
}
// Serialize serializes information the current object
func (m *ConditionalAccessFilter) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetMode() != nil {
        cast := (*m.GetMode()).String()
        err := writer.WriteStringValue("mode", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("rule", m.GetRule())
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
func (m *ConditionalAccessFilter) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetMode sets the mode property value. Mode to use for the filter. Possible values are include or exclude.
func (m *ConditionalAccessFilter) SetMode(value *FilterMode)() {
    if m != nil {
        m.mode = value
    }
}
// SetRule sets the rule property value. Rule syntax is similar to that used for membership rules for groups in Azure Active Directory. For details, see rules with multiple expressions
func (m *ConditionalAccessFilter) SetRule(value *string)() {
    if m != nil {
        m.rule = value
    }
}
