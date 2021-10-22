package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type RegistryKeyState struct {
    additionalData map[string]interface{};
    hive *RegistryHive;
    key *string;
    oldKey *string;
    oldValueData *string;
    oldValueName *string;
    operation *RegistryOperation;
    processId *int32;
    valueData *string;
    valueName *string;
    valueType *RegistryValueType;
}
func NewRegistryKeyState()(*RegistryKeyState) {
    m := &RegistryKeyState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *RegistryKeyState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *RegistryKeyState) GetHive()(*RegistryHive) {
    if m == nil {
        return nil
    } else {
        return m.hive
    }
}
func (m *RegistryKeyState) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
func (m *RegistryKeyState) GetOldKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oldKey
    }
}
func (m *RegistryKeyState) GetOldValueData()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oldValueData
    }
}
func (m *RegistryKeyState) GetOldValueName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oldValueName
    }
}
func (m *RegistryKeyState) GetOperation()(*RegistryOperation) {
    if m == nil {
        return nil
    } else {
        return m.operation
    }
}
func (m *RegistryKeyState) GetProcessId()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.processId
    }
}
func (m *RegistryKeyState) GetValueData()(*string) {
    if m == nil {
        return nil
    } else {
        return m.valueData
    }
}
func (m *RegistryKeyState) GetValueName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.valueName
    }
}
func (m *RegistryKeyState) GetValueType()(*RegistryValueType) {
    if m == nil {
        return nil
    } else {
        return m.valueType
    }
}
func (m *RegistryKeyState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegistryHive)
        if err != nil {
            return err
        }
        cast := val.(RegistryHive)
        m.SetHive(&cast)
        return nil
    }
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetKey(val)
        return nil
    }
    res["oldKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOldKey(val)
        return nil
    }
    res["oldValueData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOldValueData(val)
        return nil
    }
    res["oldValueName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetOldValueName(val)
        return nil
    }
    res["operation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegistryOperation)
        if err != nil {
            return err
        }
        cast := val.(RegistryOperation)
        m.SetOperation(&cast)
        return nil
    }
    res["processId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetProcessId(val)
        return nil
    }
    res["valueData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValueData(val)
        return nil
    }
    res["valueName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetValueName(val)
        return nil
    }
    res["valueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegistryValueType)
        if err != nil {
            return err
        }
        cast := val.(RegistryValueType)
        m.SetValueType(&cast)
        return nil
    }
    return res
}
func (m *RegistryKeyState) IsNil()(bool) {
    return m == nil
}
func (m *RegistryKeyState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetHive() != nil {
        cast := m.GetHive().String()
        err := writer.WriteStringValue("hive", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("key", m.GetKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("oldKey", m.GetOldKey())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("oldValueData", m.GetOldValueData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("oldValueName", m.GetOldValueName())
        if err != nil {
            return err
        }
    }
    if m.GetOperation() != nil {
        cast := m.GetOperation().String()
        err := writer.WriteStringValue("operation", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("processId", m.GetProcessId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("valueData", m.GetValueData())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("valueName", m.GetValueName())
        if err != nil {
            return err
        }
    }
    if m.GetValueType() != nil {
        cast := m.GetValueType().String()
        err := writer.WriteStringValue("valueType", &cast)
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
func (m *RegistryKeyState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *RegistryKeyState) SetHive(value *RegistryHive)() {
    m.hive = value
}
func (m *RegistryKeyState) SetKey(value *string)() {
    m.key = value
}
func (m *RegistryKeyState) SetOldKey(value *string)() {
    m.oldKey = value
}
func (m *RegistryKeyState) SetOldValueData(value *string)() {
    m.oldValueData = value
}
func (m *RegistryKeyState) SetOldValueName(value *string)() {
    m.oldValueName = value
}
func (m *RegistryKeyState) SetOperation(value *RegistryOperation)() {
    m.operation = value
}
func (m *RegistryKeyState) SetProcessId(value *int32)() {
    m.processId = value
}
func (m *RegistryKeyState) SetValueData(value *string)() {
    m.valueData = value
}
func (m *RegistryKeyState) SetValueName(value *string)() {
    m.valueName = value
}
func (m *RegistryKeyState) SetValueType(value *RegistryValueType)() {
    m.valueType = value
}
