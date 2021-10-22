package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ConditionalAccessLocations struct {
    additionalData map[string]interface{};
    excludeLocations []string;
    includeLocations []string;
}
func NewConditionalAccessLocations()(*ConditionalAccessLocations) {
    m := &ConditionalAccessLocations{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ConditionalAccessLocations) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ConditionalAccessLocations) GetExcludeLocations()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeLocations
    }
}
func (m *ConditionalAccessLocations) GetIncludeLocations()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeLocations
    }
}
func (m *ConditionalAccessLocations) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["excludeLocations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetExcludeLocations(res)
        return nil
    }
    res["includeLocations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIncludeLocations(res)
        return nil
    }
    return res
}
func (m *ConditionalAccessLocations) IsNil()(bool) {
    return m == nil
}
func (m *ConditionalAccessLocations) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("excludeLocations", m.GetExcludeLocations())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includeLocations", m.GetIncludeLocations())
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
func (m *ConditionalAccessLocations) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ConditionalAccessLocations) SetExcludeLocations(value []string)() {
    m.excludeLocations = value
}
func (m *ConditionalAccessLocations) SetIncludeLocations(value []string)() {
    m.includeLocations = value
}
