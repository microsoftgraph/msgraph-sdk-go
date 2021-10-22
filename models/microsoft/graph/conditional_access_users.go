package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type ConditionalAccessUsers struct {
    additionalData map[string]interface{};
    excludeGroups []string;
    excludeRoles []string;
    excludeUsers []string;
    includeGroups []string;
    includeRoles []string;
    includeUsers []string;
}
func NewConditionalAccessUsers()(*ConditionalAccessUsers) {
    m := &ConditionalAccessUsers{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *ConditionalAccessUsers) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *ConditionalAccessUsers) GetExcludeGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeGroups
    }
}
func (m *ConditionalAccessUsers) GetExcludeRoles()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeRoles
    }
}
func (m *ConditionalAccessUsers) GetExcludeUsers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.excludeUsers
    }
}
func (m *ConditionalAccessUsers) GetIncludeGroups()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeGroups
    }
}
func (m *ConditionalAccessUsers) GetIncludeRoles()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeRoles
    }
}
func (m *ConditionalAccessUsers) GetIncludeUsers()([]string) {
    if m == nil {
        return nil
    } else {
        return m.includeUsers
    }
}
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
func (m *ConditionalAccessUsers) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *ConditionalAccessUsers) SetExcludeGroups(value []string)() {
    m.excludeGroups = value
}
func (m *ConditionalAccessUsers) SetExcludeRoles(value []string)() {
    m.excludeRoles = value
}
func (m *ConditionalAccessUsers) SetExcludeUsers(value []string)() {
    m.excludeUsers = value
}
func (m *ConditionalAccessUsers) SetIncludeGroups(value []string)() {
    m.includeGroups = value
}
func (m *ConditionalAccessUsers) SetIncludeRoles(value []string)() {
    m.includeRoles = value
}
func (m *ConditionalAccessUsers) SetIncludeUsers(value []string)() {
    m.includeUsers = value
}
