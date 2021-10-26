package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type IdentityUserFlowAttributeAssignment struct {
    Entity
    // The display name of the identityUserFlowAttribute within a user flow.
    displayName *string;
    // Determines whether the identityUserFlowAttribute is optional. true means the user doesn't have to provide a value. false means the user cannot complete sign-up without providing a value.
    isOptional *bool;
    // Determines whether the identityUserFlowAttribute requires verification. This is only used for verifying the user's phone number or email address.
    requiresVerification *bool;
    // The user attribute that you want to add to your user flow.
    userAttribute *IdentityUserFlowAttribute;
    // The input options for the user flow attribute. Only applicable when the userInputType is radioSingleSelect, dropdownSingleSelect, or checkboxMultiSelect.
    userAttributeValues []UserAttributeValuesItem;
    // The input type of the user flow attribute. Possible values are: textBox, dateTimeDropdown, radioSingleSelect, dropdownSingleSelect, emailBox, checkboxMultiSelect.
    userInputType *IdentityUserFlowAttributeInputType;
}
// Instantiates a new identityUserFlowAttributeAssignment and sets the default values.
func NewIdentityUserFlowAttributeAssignment()(*IdentityUserFlowAttributeAssignment) {
    m := &IdentityUserFlowAttributeAssignment{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the displayName property value. The display name of the identityUserFlowAttribute within a user flow.
func (m *IdentityUserFlowAttributeAssignment) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the isOptional property value. Determines whether the identityUserFlowAttribute is optional. true means the user doesn't have to provide a value. false means the user cannot complete sign-up without providing a value.
func (m *IdentityUserFlowAttributeAssignment) GetIsOptional()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isOptional
    }
}
// Gets the requiresVerification property value. Determines whether the identityUserFlowAttribute requires verification. This is only used for verifying the user's phone number or email address.
func (m *IdentityUserFlowAttributeAssignment) GetRequiresVerification()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.requiresVerification
    }
}
// Gets the userAttribute property value. The user attribute that you want to add to your user flow.
func (m *IdentityUserFlowAttributeAssignment) GetUserAttribute()(*IdentityUserFlowAttribute) {
    if m == nil {
        return nil
    } else {
        return m.userAttribute
    }
}
// Gets the userAttributeValues property value. The input options for the user flow attribute. Only applicable when the userInputType is radioSingleSelect, dropdownSingleSelect, or checkboxMultiSelect.
func (m *IdentityUserFlowAttributeAssignment) GetUserAttributeValues()([]UserAttributeValuesItem) {
    if m == nil {
        return nil
    } else {
        return m.userAttributeValues
    }
}
// Gets the userInputType property value. The input type of the user flow attribute. Possible values are: textBox, dateTimeDropdown, radioSingleSelect, dropdownSingleSelect, emailBox, checkboxMultiSelect.
func (m *IdentityUserFlowAttributeAssignment) GetUserInputType()(*IdentityUserFlowAttributeInputType) {
    if m == nil {
        return nil
    } else {
        return m.userInputType
    }
}
// The deserialization information for the current model
func (m *IdentityUserFlowAttributeAssignment) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["isOptional"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsOptional(val)
        return nil
    }
    res["requiresVerification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetRequiresVerification(val)
        return nil
    }
    res["userAttribute"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewIdentityUserFlowAttribute() })
        if err != nil {
            return err
        }
        m.SetUserAttribute(val.(*IdentityUserFlowAttribute))
        return nil
    }
    res["userAttributeValues"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewUserAttributeValuesItem() })
        if err != nil {
            return err
        }
        res := make([]UserAttributeValuesItem, len(val))
        for i, v := range val {
            res[i] = *(v.(*UserAttributeValuesItem))
        }
        m.SetUserAttributeValues(res)
        return nil
    }
    res["userInputType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseIdentityUserFlowAttributeInputType)
        if err != nil {
            return err
        }
        cast := val.(IdentityUserFlowAttributeInputType)
        m.SetUserInputType(&cast)
        return nil
    }
    return res
}
func (m *IdentityUserFlowAttributeAssignment) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *IdentityUserFlowAttributeAssignment) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isOptional", m.GetIsOptional())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("requiresVerification", m.GetRequiresVerification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("userAttribute", m.GetUserAttribute())
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetUserAttributeValues()))
        for i, v := range m.GetUserAttributeValues() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("userAttributeValues", cast)
        if err != nil {
            return err
        }
    }
    if m.GetUserInputType() != nil {
        cast := m.GetUserInputType().String()
        err = writer.WriteStringValue("userInputType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the displayName property value. The display name of the identityUserFlowAttribute within a user flow.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *IdentityUserFlowAttributeAssignment) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the isOptional property value. Determines whether the identityUserFlowAttribute is optional. true means the user doesn't have to provide a value. false means the user cannot complete sign-up without providing a value.
// Parameters:
//  - value : Value to set for the isOptional property.
func (m *IdentityUserFlowAttributeAssignment) SetIsOptional(value *bool)() {
    m.isOptional = value
}
// Sets the requiresVerification property value. Determines whether the identityUserFlowAttribute requires verification. This is only used for verifying the user's phone number or email address.
// Parameters:
//  - value : Value to set for the requiresVerification property.
func (m *IdentityUserFlowAttributeAssignment) SetRequiresVerification(value *bool)() {
    m.requiresVerification = value
}
// Sets the userAttribute property value. The user attribute that you want to add to your user flow.
// Parameters:
//  - value : Value to set for the userAttribute property.
func (m *IdentityUserFlowAttributeAssignment) SetUserAttribute(value *IdentityUserFlowAttribute)() {
    m.userAttribute = value
}
// Sets the userAttributeValues property value. The input options for the user flow attribute. Only applicable when the userInputType is radioSingleSelect, dropdownSingleSelect, or checkboxMultiSelect.
// Parameters:
//  - value : Value to set for the userAttributeValues property.
func (m *IdentityUserFlowAttributeAssignment) SetUserAttributeValues(value []UserAttributeValuesItem)() {
    m.userAttributeValues = value
}
// Sets the userInputType property value. The input type of the user flow attribute. Possible values are: textBox, dateTimeDropdown, radioSingleSelect, dropdownSingleSelect, emailBox, checkboxMultiSelect.
// Parameters:
//  - value : Value to set for the userInputType property.
func (m *IdentityUserFlowAttributeAssignment) SetUserInputType(value *IdentityUserFlowAttributeInputType)() {
    m.userInputType = value
}
