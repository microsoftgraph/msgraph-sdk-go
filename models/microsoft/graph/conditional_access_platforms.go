package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ConditionalAccessPlatforms 
type ConditionalAccessPlatforms struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue.
    excludePlatforms []ConditionalAccessDevicePlatform;
    // Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue.
    includePlatforms []ConditionalAccessDevicePlatform;
}
// NewConditionalAccessPlatforms instantiates a new conditionalAccessPlatforms and sets the default values.
func NewConditionalAccessPlatforms()(*ConditionalAccessPlatforms) {
    m := &ConditionalAccessPlatforms{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessPlatforms) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetExcludePlatforms gets the excludePlatforms property value. Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue.
func (m *ConditionalAccessPlatforms) GetExcludePlatforms()([]ConditionalAccessDevicePlatform) {
    if m == nil {
        return nil
    } else {
        return m.excludePlatforms
    }
}
// GetIncludePlatforms gets the includePlatforms property value. Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue.
func (m *ConditionalAccessPlatforms) GetIncludePlatforms()([]ConditionalAccessDevicePlatform) {
    if m == nil {
        return nil
    } else {
        return m.includePlatforms
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ConditionalAccessPlatforms) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["excludePlatforms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseConditionalAccessDevicePlatform)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessDevicePlatform, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConditionalAccessDevicePlatform))
            }
            m.SetExcludePlatforms(res)
        }
        return nil
    }
    res["includePlatforms"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfEnumValues(ParseConditionalAccessDevicePlatform)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConditionalAccessDevicePlatform, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConditionalAccessDevicePlatform))
            }
            m.SetIncludePlatforms(res)
        }
        return nil
    }
    return res
}
func (m *ConditionalAccessPlatforms) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ConditionalAccessPlatforms) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("excludePlatforms", SerializeConditionalAccessDevicePlatform(m.GetExcludePlatforms()))
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includePlatforms", SerializeConditionalAccessDevicePlatform(m.GetIncludePlatforms()))
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessPlatforms) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetExcludePlatforms sets the excludePlatforms property value. Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue.
func (m *ConditionalAccessPlatforms) SetExcludePlatforms(value []ConditionalAccessDevicePlatform)() {
    if m != nil {
        m.excludePlatforms = value
    }
}
// SetIncludePlatforms sets the includePlatforms property value. Possible values are: android, iOS, windows, windowsPhone, macOS, all, unknownFutureValue.
func (m *ConditionalAccessPlatforms) SetIncludePlatforms(value []ConditionalAccessDevicePlatform)() {
    if m != nil {
        m.includePlatforms = value
    }
}
