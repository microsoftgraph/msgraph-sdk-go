package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// AlterationResponse 
type AlterationResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Defines the original user query string.
    originalQueryString *string;
    // Defines the details of alteration information for the spelling correction.
    queryAlteration *SearchAlteration;
    // Defines the type of the spelling correction. Possible values are suggestion, modification.
    queryAlterationType *SearchAlterationType;
}
// NewAlterationResponse instantiates a new alterationResponse and sets the default values.
func NewAlterationResponse()(*AlterationResponse) {
    m := &AlterationResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AlterationResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetOriginalQueryString gets the originalQueryString property value. Defines the original user query string.
func (m *AlterationResponse) GetOriginalQueryString()(*string) {
    if m == nil {
        return nil
    } else {
        return m.originalQueryString
    }
}
// GetQueryAlteration gets the queryAlteration property value. Defines the details of alteration information for the spelling correction.
func (m *AlterationResponse) GetQueryAlteration()(*SearchAlteration) {
    if m == nil {
        return nil
    } else {
        return m.queryAlteration
    }
}
// GetQueryAlterationType gets the queryAlterationType property value. Defines the type of the spelling correction. Possible values are suggestion, modification.
func (m *AlterationResponse) GetQueryAlterationType()(*SearchAlterationType) {
    if m == nil {
        return nil
    } else {
        return m.queryAlterationType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AlterationResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["originalQueryString"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOriginalQueryString(val)
        }
        return nil
    }
    res["queryAlteration"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSearchAlteration() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryAlteration(val.(*SearchAlteration))
        }
        return nil
    }
    res["queryAlterationType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseSearchAlterationType)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQueryAlterationType(val.(*SearchAlterationType))
        }
        return nil
    }
    return res
}
func (m *AlterationResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *AlterationResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("originalQueryString", m.GetOriginalQueryString())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("queryAlteration", m.GetQueryAlteration())
        if err != nil {
            return err
        }
    }
    if m.GetQueryAlterationType() != nil {
        cast := (*m.GetQueryAlterationType()).String()
        err := writer.WriteStringValue("queryAlterationType", &cast)
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
func (m *AlterationResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetOriginalQueryString sets the originalQueryString property value. Defines the original user query string.
func (m *AlterationResponse) SetOriginalQueryString(value *string)() {
    if m != nil {
        m.originalQueryString = value
    }
}
// SetQueryAlteration sets the queryAlteration property value. Defines the details of alteration information for the spelling correction.
func (m *AlterationResponse) SetQueryAlteration(value *SearchAlteration)() {
    if m != nil {
        m.queryAlteration = value
    }
}
// SetQueryAlterationType sets the queryAlterationType property value. Defines the type of the spelling correction. Possible values are suggestion, modification.
func (m *AlterationResponse) SetQueryAlterationType(value *SearchAlterationType)() {
    if m != nil {
        m.queryAlterationType = value
    }
}
