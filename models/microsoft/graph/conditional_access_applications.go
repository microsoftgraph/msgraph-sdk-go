package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConditionalAccessApplications struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The list of application IDs explicitly excluded from the policy.
    excludeApplications []string;
    // The list of application IDs the policy applies to, unless explicitly excluded (in excludeApplications). Can also be set to All.
    includeApplications []string;
    // Authentication context class references include. Supported values are c1 through c25.
    includeAuthenticationContextClassReferences []string;
    // User actions to include. Supported values are urn:user:registersecurityinfo and urn:user:registerdevice
    includeUserActions []string;
}
// Instantiates a new conditionalAccessApplications and sets the default values.
func NewConditionalAccessApplications()(*ConditionalAccessApplications) {
    m := &ConditionalAccessApplications{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessApplications) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the excludeApplications property value. The list of application IDs explicitly excluded from the policy.
func (m *ConditionalAccessApplications) GetExcludeApplications()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeApplications
    }
}
// Gets the includeApplications property value. The list of application IDs the policy applies to, unless explicitly excluded (in excludeApplications). Can also be set to All.
func (m *ConditionalAccessApplications) GetIncludeApplications()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeApplications
    }
}
// Gets the includeAuthenticationContextClassReferences property value. Authentication context class references include. Supported values are c1 through c25.
func (m *ConditionalAccessApplications) GetIncludeAuthenticationContextClassReferences()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeAuthenticationContextClassReferences
    }
}
// Gets the includeUserActions property value. User actions to include. Supported values are urn:user:registersecurityinfo and urn:user:registerdevice
func (m *ConditionalAccessApplications) GetIncludeUserActions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeUserActions
    }
}
// The deserialization information for the current model
func (m *ConditionalAccessApplications) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["excludeApplications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetExcludeApplications(res)
        return nil
    }
    res["includeApplications"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIncludeApplications(res)
        return nil
    }
    res["includeAuthenticationContextClassReferences"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIncludeAuthenticationContextClassReferences(res)
        return nil
    }
    res["includeUserActions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIncludeUserActions(res)
        return nil
    }
    return res
}
func (m *ConditionalAccessApplications) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ConditionalAccessApplications) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("excludeApplications", m.GetExcludeApplications())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includeApplications", m.GetIncludeApplications())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includeAuthenticationContextClassReferences", m.GetIncludeAuthenticationContextClassReferences())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includeUserActions", m.GetIncludeUserActions())
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
func (m *ConditionalAccessApplications) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the excludeApplications property value. The list of application IDs explicitly excluded from the policy.
// Parameters:
//  - value : Value to set for the excludeApplications property.
func (m *ConditionalAccessApplications) SetExcludeApplications(value []string)() {
    m.excludeApplications = value
}
// Sets the includeApplications property value. The list of application IDs the policy applies to, unless explicitly excluded (in excludeApplications). Can also be set to All.
// Parameters:
//  - value : Value to set for the includeApplications property.
func (m *ConditionalAccessApplications) SetIncludeApplications(value []string)() {
    m.includeApplications = value
}
// Sets the includeAuthenticationContextClassReferences property value. Authentication context class references include. Supported values are c1 through c25.
// Parameters:
//  - value : Value to set for the includeAuthenticationContextClassReferences property.
func (m *ConditionalAccessApplications) SetIncludeAuthenticationContextClassReferences(value []string)() {
    m.includeAuthenticationContextClassReferences = value
}
// Sets the includeUserActions property value. User actions to include. Supported values are urn:user:registersecurityinfo and urn:user:registerdevice
// Parameters:
//  - value : Value to set for the includeUserActions property.
func (m *ConditionalAccessApplications) SetIncludeUserActions(value []string)() {
    m.includeUserActions = value
}
