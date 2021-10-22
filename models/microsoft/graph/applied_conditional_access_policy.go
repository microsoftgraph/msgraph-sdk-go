package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type AppliedConditionalAccessPolicy struct {
    additionalData map[string]interface{};
    displayName *string;
    enforcedGrantControls []string;
    enforcedSessionControls []string;
    id *string;
    result *AppliedConditionalAccessPolicyResult;
}
func NewAppliedConditionalAccessPolicy()(*AppliedConditionalAccessPolicy) {
    m := &AppliedConditionalAccessPolicy{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *AppliedConditionalAccessPolicy) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *AppliedConditionalAccessPolicy) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
func (m *AppliedConditionalAccessPolicy) GetEnforcedGrantControls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enforcedGrantControls
    }
}
func (m *AppliedConditionalAccessPolicy) GetEnforcedSessionControls()([]string) {
    if m == nil {
        return nil
    } else {
        return m.enforcedSessionControls
    }
}
func (m *AppliedConditionalAccessPolicy) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
func (m *AppliedConditionalAccessPolicy) GetResult()(*AppliedConditionalAccessPolicyResult) {
    if m == nil {
        return nil
    } else {
        return m.result
    }
}
func (m *AppliedConditionalAccessPolicy) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["enforcedGrantControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetEnforcedGrantControls(res)
        return nil
    }
    res["enforcedSessionControls"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetEnforcedSessionControls(res)
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetId(val)
        return nil
    }
    res["result"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAppliedConditionalAccessPolicyResult)
        if err != nil {
            return err
        }
        cast := val.(AppliedConditionalAccessPolicyResult)
        m.SetResult(&cast)
        return nil
    }
    return res
}
func (m *AppliedConditionalAccessPolicy) IsNil()(bool) {
    return m == nil
}
func (m *AppliedConditionalAccessPolicy) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("enforcedGrantControls", m.GetEnforcedGrantControls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("enforcedSessionControls", m.GetEnforcedSessionControls())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
        if err != nil {
            return err
        }
    }
    if m.GetResult() != nil {
        cast := m.GetResult().String()
        err := writer.WriteStringValue("result", &cast)
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
func (m *AppliedConditionalAccessPolicy) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *AppliedConditionalAccessPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
func (m *AppliedConditionalAccessPolicy) SetEnforcedGrantControls(value []string)() {
    m.enforcedGrantControls = value
}
func (m *AppliedConditionalAccessPolicy) SetEnforcedSessionControls(value []string)() {
    m.enforcedSessionControls = value
}
func (m *AppliedConditionalAccessPolicy) SetId(value *string)() {
    m.id = value
}
func (m *AppliedConditionalAccessPolicy) SetResult(value *AppliedConditionalAccessPolicyResult)() {
    m.result = value
}
