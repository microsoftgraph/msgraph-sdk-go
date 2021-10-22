package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Process struct {
    accountName *string;
    additionalData map[string]interface{};
    commandLine *string;
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    fileHash *FileHash;
    integrityLevel *ProcessIntegrityLevel;
    isElevated *bool;
    name *string;
    parentProcessCreatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    parentProcessId *int32;
    parentProcessName *string;
    path *string;
    processId *int32;
}
func NewProcess()(*Process) {
    m := &Process{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Process) GetAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountName
    }
}
func (m *Process) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Process) GetCommandLine()(*string) {
    if m == nil {
        return nil
    } else {
        return m.commandLine
    }
}
func (m *Process) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
func (m *Process) GetFileHash()(*FileHash) {
    if m == nil {
        return nil
    } else {
        return m.fileHash
    }
}
func (m *Process) GetIntegrityLevel()(*ProcessIntegrityLevel) {
    if m == nil {
        return nil
    } else {
        return m.integrityLevel
    }
}
func (m *Process) GetIsElevated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isElevated
    }
}
func (m *Process) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *Process) GetParentProcessCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.parentProcessCreatedDateTime
    }
}
func (m *Process) GetParentProcessId()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.parentProcessId
    }
}
func (m *Process) GetParentProcessName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentProcessName
    }
}
func (m *Process) GetPath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.path
    }
}
func (m *Process) GetProcessId()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.processId
    }
}
func (m *Process) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accountName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAccountName(val)
        return nil
    }
    res["commandLine"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCommandLine(val)
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetCreatedDateTime(val)
        return nil
    }
    res["fileHash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFileHash() })
        if err != nil {
            return err
        }
        m.SetFileHash(val.(*FileHash))
        return nil
    }
    res["integrityLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProcessIntegrityLevel)
        if err != nil {
            return err
        }
        cast := val.(ProcessIntegrityLevel)
        m.SetIntegrityLevel(&cast)
        return nil
    }
    res["isElevated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsElevated(val)
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetName(val)
        return nil
    }
    res["parentProcessCreatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        m.SetParentProcessCreatedDateTime(val)
        return nil
    }
    res["parentProcessId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetParentProcessId(val)
        return nil
    }
    res["parentProcessName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetParentProcessName(val)
        return nil
    }
    res["path"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPath(val)
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
    return res
}
func (m *Process) IsNil()(bool) {
    return m == nil
}
func (m *Process) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("accountName", m.GetAccountName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("commandLine", m.GetCommandLine())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("createdDateTime", m.GetCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteObjectValue("fileHash", m.GetFileHash())
        if err != nil {
            return err
        }
    }
    if m.GetIntegrityLevel() != nil {
        cast := m.GetIntegrityLevel().String()
        err := writer.WriteStringValue("integrityLevel", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isElevated", m.GetIsElevated())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("name", m.GetName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteTimeValue("parentProcessCreatedDateTime", m.GetParentProcessCreatedDateTime())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("parentProcessId", m.GetParentProcessId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("parentProcessName", m.GetParentProcessName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("path", m.GetPath())
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
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *Process) SetAccountName(value *string)() {
    m.accountName = value
}
func (m *Process) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Process) SetCommandLine(value *string)() {
    m.commandLine = value
}
func (m *Process) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
func (m *Process) SetFileHash(value *FileHash)() {
    m.fileHash = value
}
func (m *Process) SetIntegrityLevel(value *ProcessIntegrityLevel)() {
    m.integrityLevel = value
}
func (m *Process) SetIsElevated(value *bool)() {
    m.isElevated = value
}
func (m *Process) SetName(value *string)() {
    m.name = value
}
func (m *Process) SetParentProcessCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.parentProcessCreatedDateTime = value
}
func (m *Process) SetParentProcessId(value *int32)() {
    m.parentProcessId = value
}
func (m *Process) SetParentProcessName(value *string)() {
    m.parentProcessName = value
}
func (m *Process) SetPath(value *string)() {
    m.path = value
}
func (m *Process) SetProcessId(value *int32)() {
    m.processId = value
}
