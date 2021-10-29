package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ConditionalAccessUsers struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Group IDs excluded from scope of policy.
    excludeGroups []string;
    // Role IDs excluded from scope of policy.
    excludeRoles []string;
    // User IDs excluded from scope of policy and/or GuestsOrExternalUsers.
    excludeUsers []string;
    // Group IDs in scope of policy unless explicitly excluded, or All.
    includeGroups []string;
    // Role IDs in scope of policy unless explicitly excluded, or All.
    includeRoles []string;
    // User IDs in scope of policy unless explicitly excluded, or None or All or GuestsOrExternalUsers.
    includeUsers []string;
}
// Instantiates a new conditionalAccessUsers and sets the default values.
func NewConditionalAccessUsers()(*ConditionalAccessUsers) {
    m := &ConditionalAccessUsers{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ConditionalAccessUsers) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the excludeGroups property value. Group IDs excluded from scope of policy.
func (m *ConditionalAccessUsers) GetExcludeGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeGroups
    }
}
// Gets the excludeRoles property value. Role IDs excluded from scope of policy.
func (m *ConditionalAccessUsers) GetExcludeRoles()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeRoles
    }
}
// Gets the excludeUsers property value. User IDs excluded from scope of policy and/or GuestsOrExternalUsers.
func (m *ConditionalAccessUsers) GetExcludeUsers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeUsers
    }
}
// Gets the includeGroups property value. Group IDs in scope of policy unless explicitly excluded, or All.
func (m *ConditionalAccessUsers) GetIncludeGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeGroups
    }
}
// Gets the includeRoles property value. Role IDs in scope of policy unless explicitly excluded, or All.
func (m *ConditionalAccessUsers) GetIncludeRoles()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeRoles
    }
}
// Gets the includeUsers property value. User IDs in scope of policy unless explicitly excluded, or None or All or GuestsOrExternalUsers.
func (m *ConditionalAccessUsers) GetIncludeUsers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeUsers
    }
}
// The deserialization information for the current model
func (m *ConditionalAccessUsers) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["excludeGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetExcludeGroups(res)
        return nil
    }
    res["excludeRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetExcludeRoles(res)
        return nil
    }
    res["excludeUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetExcludeUsers(res)
        return nil
    }
    res["includeGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIncludeGroups(res)
        return nil
    }
    res["includeRoles"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIncludeRoles(res)
        return nil
    }
    res["includeUsers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetIncludeUsers(res)
        return nil
    }
    return res
}
func (m *ConditionalAccessUsers) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ConditionalAccessUsers) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteCollectionOfStringValues("excludeGroups", m.GetExcludeGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("excludeRoles", m.GetExcludeRoles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("excludeUsers", m.GetExcludeUsers())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includeGroups", m.GetIncludeGroups())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includeRoles", m.GetIncludeRoles())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("includeUsers", m.GetIncludeUsers())
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
func (m *ConditionalAccessUsers) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the excludeGroups property value. Group IDs excluded from scope of policy.
// Parameters:
//  - value : Value to set for the excludeGroups property.
func (m *ConditionalAccessUsers) SetExcludeGroups(value []string)() {
    m.excludeGroups = value
}
// Sets the excludeRoles property value. Role IDs excluded from scope of policy.
// Parameters:
//  - value : Value to set for the excludeRoles property.
func (m *ConditionalAccessUsers) SetExcludeRoles(value []string)() {
    m.excludeRoles = value
}
// Sets the excludeUsers property value. User IDs excluded from scope of policy and/or GuestsOrExternalUsers.
// Parameters:
//  - value : Value to set for the excludeUsers property.
func (m *ConditionalAccessUsers) SetExcludeUsers(value []string)() {
    m.excludeUsers = value
}
// Sets the includeGroups property value. Group IDs in scope of policy unless explicitly excluded, or All.
// Parameters:
//  - value : Value to set for the includeGroups property.
func (m *ConditionalAccessUsers) SetIncludeGroups(value []string)() {
    m.includeGroups = value
}
// Sets the includeRoles property value. Role IDs in scope of policy unless explicitly excluded, or All.
// Parameters:
//  - value : Value to set for the includeRoles property.
func (m *ConditionalAccessUsers) SetIncludeRoles(value []string)() {
    m.includeRoles = value
}
// Sets the includeUsers property value. User IDs in scope of policy unless explicitly excluded, or None or All or GuestsOrExternalUsers.
// Parameters:
//  - value : Value to set for the includeUsers property.
func (m *ConditionalAccessUsers) SetIncludeUsers(value []string)() {
    m.includeUsers = value
}
