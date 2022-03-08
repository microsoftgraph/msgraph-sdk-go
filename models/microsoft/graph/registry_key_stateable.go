package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// RegistryKeyStateable 
type RegistryKeyStateable interface {
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.AdditionalDataHolder
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetHive()(*RegistryHive)
    GetKey()(*string)
    GetOldKey()(*string)
    GetOldValueData()(*string)
    GetOldValueName()(*string)
    GetOperation()(*RegistryOperation)
    GetProcessId()(*int32)
    GetValueData()(*string)
    GetValueName()(*string)
    GetValueType()(*RegistryValueType)
    SetHive(value *RegistryHive)()
    SetKey(value *string)()
    SetOldKey(value *string)()
    SetOldValueData(value *string)()
    SetOldValueName(value *string)()
    SetOperation(value *RegistryOperation)()
    SetProcessId(value *int32)()
    SetValueData(value *string)()
    SetValueName(value *string)()
    SetValueType(value *RegistryValueType)()
}
