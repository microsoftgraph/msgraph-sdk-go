package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// registryKeyState 
type RegistryKeyState struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A Windows registry hive : HKEY_CURRENT_CONFIG HKEY_CURRENT_USER HKEY_LOCAL_MACHINE/SAM HKEY_LOCAL_MACHINE/Security HKEY_LOCAL_MACHINE/Software HKEY_LOCAL_MACHINE/System HKEY_USERS/.Default. Possible values are: unknown, currentConfig, currentUser, localMachineSam, localMachineSecurity, localMachineSoftware, localMachineSystem, usersDefault.
    hive *RegistryHive;
    // Current (i.e. changed) registry key (excludes HIVE).
    key *string;
    // Previous (i.e. before changed) registry key (excludes HIVE).
    oldKey *string;
    // Previous (i.e. before changed) registry key value data (contents).
    oldValueData *string;
    // Previous (i.e. before changed) registry key value name.
    oldValueName *string;
    // Operation that changed the registry key name and/or value. Possible values are: unknown, create, modify, delete.
    operation *RegistryOperation;
    // Process ID (PID) of the process that modified the registry key (process details will appear in the alert 'processes' collection).
    processId *int32;
    // Current (i.e. changed) registry key value data (contents).
    valueData *string;
    // Current (i.e. changed) registry key value name
    valueName *string;
    // Registry key value type REG_BINARY REG_DWORD REG_DWORD_LITTLE_ENDIAN REG_DWORD_BIG_ENDIANREG_EXPAND_SZ REG_LINK REG_MULTI_SZ REG_NONE REG_QWORD REG_QWORD_LITTLE_ENDIAN REG_SZ Possible values are: unknown, binary, dword, dwordLittleEndian, dwordBigEndian, expandSz, link, multiSz, none, qword, qwordlittleEndian, sz.
    valueType *RegistryValueType;
}
// NewRegistryKeyState instantiates a new registryKeyState and sets the default values.
func NewRegistryKeyState()(*RegistryKeyState) {
    m := &RegistryKeyState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RegistryKeyState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetHive gets the hive property value. A Windows registry hive : HKEY_CURRENT_CONFIG HKEY_CURRENT_USER HKEY_LOCAL_MACHINE/SAM HKEY_LOCAL_MACHINE/Security HKEY_LOCAL_MACHINE/Software HKEY_LOCAL_MACHINE/System HKEY_USERS/.Default. Possible values are: unknown, currentConfig, currentUser, localMachineSam, localMachineSecurity, localMachineSoftware, localMachineSystem, usersDefault.
func (m *RegistryKeyState) GetHive()(*RegistryHive) {
    if m == nil {
        return nil
    } else {
        return m.hive
    }
}
// GetKey gets the key property value. Current (i.e. changed) registry key (excludes HIVE).
func (m *RegistryKeyState) GetKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.key
    }
}
// GetOldKey gets the oldKey property value. Previous (i.e. before changed) registry key (excludes HIVE).
func (m *RegistryKeyState) GetOldKey()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oldKey
    }
}
// GetOldValueData gets the oldValueData property value. Previous (i.e. before changed) registry key value data (contents).
func (m *RegistryKeyState) GetOldValueData()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oldValueData
    }
}
// GetOldValueName gets the oldValueName property value. Previous (i.e. before changed) registry key value name.
func (m *RegistryKeyState) GetOldValueName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.oldValueName
    }
}
// GetOperation gets the operation property value. Operation that changed the registry key name and/or value. Possible values are: unknown, create, modify, delete.
func (m *RegistryKeyState) GetOperation()(*RegistryOperation) {
    if m == nil {
        return nil
    } else {
        return m.operation
    }
}
// GetProcessId gets the processId property value. Process ID (PID) of the process that modified the registry key (process details will appear in the alert 'processes' collection).
func (m *RegistryKeyState) GetProcessId()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.processId
    }
}
// GetValueData gets the valueData property value. Current (i.e. changed) registry key value data (contents).
func (m *RegistryKeyState) GetValueData()(*string) {
    if m == nil {
        return nil
    } else {
        return m.valueData
    }
}
// GetValueName gets the valueName property value. Current (i.e. changed) registry key value name
func (m *RegistryKeyState) GetValueName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.valueName
    }
}
// GetValueType gets the valueType property value. Registry key value type REG_BINARY REG_DWORD REG_DWORD_LITTLE_ENDIAN REG_DWORD_BIG_ENDIANREG_EXPAND_SZ REG_LINK REG_MULTI_SZ REG_NONE REG_QWORD REG_QWORD_LITTLE_ENDIAN REG_SZ Possible values are: unknown, binary, dword, dwordLittleEndian, dwordBigEndian, expandSz, link, multiSz, none, qword, qwordlittleEndian, sz.
func (m *RegistryKeyState) GetValueType()(*RegistryValueType) {
    if m == nil {
        return nil
    } else {
        return m.valueType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *RegistryKeyState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hive"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegistryHive)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RegistryHive)
            m.SetHive(&cast)
        }
        return nil
    }
    res["key"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetKey(val)
        }
        return nil
    }
    res["oldKey"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldKey(val)
        }
        return nil
    }
    res["oldValueData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldValueData(val)
        }
        return nil
    }
    res["oldValueName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOldValueName(val)
        }
        return nil
    }
    res["operation"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegistryOperation)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RegistryOperation)
            m.SetOperation(&cast)
        }
        return nil
    }
    res["processId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProcessId(val)
        }
        return nil
    }
    res["valueData"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueData(val)
        }
        return nil
    }
    res["valueName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValueName(val)
        }
        return nil
    }
    res["valueType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseRegistryValueType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(RegistryValueType)
            m.SetValueType(&cast)
        }
        return nil
    }
    return res
}
func (m *RegistryKeyState) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *RegistryKeyState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetHive sets the hive property value. A Windows registry hive : HKEY_CURRENT_CONFIG HKEY_CURRENT_USER HKEY_LOCAL_MACHINE/SAM HKEY_LOCAL_MACHINE/Security HKEY_LOCAL_MACHINE/Software HKEY_LOCAL_MACHINE/System HKEY_USERS/.Default. Possible values are: unknown, currentConfig, currentUser, localMachineSam, localMachineSecurity, localMachineSoftware, localMachineSystem, usersDefault.
func (m *RegistryKeyState) SetHive(value *RegistryHive)() {
    m.hive = value
}
// SetKey sets the key property value. Current (i.e. changed) registry key (excludes HIVE).
func (m *RegistryKeyState) SetKey(value *string)() {
    m.key = value
}
// SetOldKey sets the oldKey property value. Previous (i.e. before changed) registry key (excludes HIVE).
func (m *RegistryKeyState) SetOldKey(value *string)() {
    m.oldKey = value
}
// SetOldValueData sets the oldValueData property value. Previous (i.e. before changed) registry key value data (contents).
func (m *RegistryKeyState) SetOldValueData(value *string)() {
    m.oldValueData = value
}
// SetOldValueName sets the oldValueName property value. Previous (i.e. before changed) registry key value name.
func (m *RegistryKeyState) SetOldValueName(value *string)() {
    m.oldValueName = value
}
// SetOperation sets the operation property value. Operation that changed the registry key name and/or value. Possible values are: unknown, create, modify, delete.
func (m *RegistryKeyState) SetOperation(value *RegistryOperation)() {
    m.operation = value
}
// SetProcessId sets the processId property value. Process ID (PID) of the process that modified the registry key (process details will appear in the alert 'processes' collection).
func (m *RegistryKeyState) SetProcessId(value *int32)() {
    m.processId = value
}
// SetValueData sets the valueData property value. Current (i.e. changed) registry key value data (contents).
func (m *RegistryKeyState) SetValueData(value *string)() {
    m.valueData = value
}
// SetValueName sets the valueName property value. Current (i.e. changed) registry key value name
func (m *RegistryKeyState) SetValueName(value *string)() {
    m.valueName = value
}
// SetValueType sets the valueType property value. Registry key value type REG_BINARY REG_DWORD REG_DWORD_LITTLE_ENDIAN REG_DWORD_BIG_ENDIANREG_EXPAND_SZ REG_LINK REG_MULTI_SZ REG_NONE REG_QWORD REG_QWORD_LITTLE_ENDIAN REG_SZ Possible values are: unknown, binary, dword, dwordLittleEndian, dwordBigEndian, expandSz, link, multiSz, none, qword, qwordlittleEndian, sz.
func (m *RegistryKeyState) SetValueType(value *RegistryValueType)() {
    m.valueType = value
}
