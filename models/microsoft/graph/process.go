package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type Process struct {
    // User account identifier (user account context the process ran under) for example, AccountName, SID, and so on.
    accountName *string;
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The full process invocation commandline including all parameters.
    commandLine *string;
    // Time at which the process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    createdDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // Complex type containing file hashes (cryptographic and location-sensitive).
    fileHash *FileHash;
    // The integrity level of the process. Possible values are: unknown, untrusted, low, medium, high, system.
    integrityLevel *ProcessIntegrityLevel;
    // True if the process is elevated.
    isElevated *bool;
    // The name of the process' Image file.
    name *string;
    // DateTime at which the parent process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
    parentProcessCreatedDateTime *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time;
    // The Process ID (PID) of the parent process.
    parentProcessId *int32;
    // The name of the image file of the parent process.
    parentProcessName *string;
    // Full path, including filename.
    path *string;
    // The Process ID (PID) of the process.
    processId *int32;
}
// Instantiates a new process and sets the default values.
func NewProcess()(*Process) {
    m := &Process{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the accountName property value. User account identifier (user account context the process ran under) for example, AccountName, SID, and so on.
func (m *Process) GetAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountName
    }
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Process) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the commandLine property value. The full process invocation commandline including all parameters.
func (m *Process) GetCommandLine()(*string) {
    if m == nil {
        return nil
    } else {
        return m.commandLine
    }
}
// Gets the createdDateTime property value. Time at which the process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Process) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// Gets the fileHash property value. Complex type containing file hashes (cryptographic and location-sensitive).
func (m *Process) GetFileHash()(*FileHash) {
    if m == nil {
        return nil
    } else {
        return m.fileHash
    }
}
// Gets the integrityLevel property value. The integrity level of the process. Possible values are: unknown, untrusted, low, medium, high, system.
func (m *Process) GetIntegrityLevel()(*ProcessIntegrityLevel) {
    if m == nil {
        return nil
    } else {
        return m.integrityLevel
    }
}
// Gets the isElevated property value. True if the process is elevated.
func (m *Process) GetIsElevated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isElevated
    }
}
// Gets the name property value. The name of the process' Image file.
func (m *Process) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// Gets the parentProcessCreatedDateTime property value. DateTime at which the parent process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Process) GetParentProcessCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.parentProcessCreatedDateTime
    }
}
// Gets the parentProcessId property value. The Process ID (PID) of the parent process.
func (m *Process) GetParentProcessId()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.parentProcessId
    }
}
// Gets the parentProcessName property value. The name of the image file of the parent process.
func (m *Process) GetParentProcessName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentProcessName
    }
}
// Gets the path property value. Full path, including filename.
func (m *Process) GetPath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.path
    }
}
// Gets the processId property value. The Process ID (PID) of the process.
func (m *Process) GetProcessId()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.processId
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the accountName property value. User account identifier (user account context the process ran under) for example, AccountName, SID, and so on.
// Parameters:
//  - value : Value to set for the accountName property.
func (m *Process) SetAccountName(value *string)() {
    m.accountName = value
}
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *Process) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the commandLine property value. The full process invocation commandline including all parameters.
// Parameters:
//  - value : Value to set for the commandLine property.
func (m *Process) SetCommandLine(value *string)() {
    m.commandLine = value
}
// Sets the createdDateTime property value. Time at which the process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the createdDateTime property.
func (m *Process) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.createdDateTime = value
}
// Sets the fileHash property value. Complex type containing file hashes (cryptographic and location-sensitive).
// Parameters:
//  - value : Value to set for the fileHash property.
func (m *Process) SetFileHash(value *FileHash)() {
    m.fileHash = value
}
// Sets the integrityLevel property value. The integrity level of the process. Possible values are: unknown, untrusted, low, medium, high, system.
// Parameters:
//  - value : Value to set for the integrityLevel property.
func (m *Process) SetIntegrityLevel(value *ProcessIntegrityLevel)() {
    m.integrityLevel = value
}
// Sets the isElevated property value. True if the process is elevated.
// Parameters:
//  - value : Value to set for the isElevated property.
func (m *Process) SetIsElevated(value *bool)() {
    m.isElevated = value
}
// Sets the name property value. The name of the process' Image file.
// Parameters:
//  - value : Value to set for the name property.
func (m *Process) SetName(value *string)() {
    m.name = value
}
// Sets the parentProcessCreatedDateTime property value. DateTime at which the parent process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
// Parameters:
//  - value : Value to set for the parentProcessCreatedDateTime property.
func (m *Process) SetParentProcessCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    m.parentProcessCreatedDateTime = value
}
// Sets the parentProcessId property value. The Process ID (PID) of the parent process.
// Parameters:
//  - value : Value to set for the parentProcessId property.
func (m *Process) SetParentProcessId(value *int32)() {
    m.parentProcessId = value
}
// Sets the parentProcessName property value. The name of the image file of the parent process.
// Parameters:
//  - value : Value to set for the parentProcessName property.
func (m *Process) SetParentProcessName(value *string)() {
    m.parentProcessName = value
}
// Sets the path property value. Full path, including filename.
// Parameters:
//  - value : Value to set for the path property.
func (m *Process) SetPath(value *string)() {
    m.path = value
}
// Sets the processId property value. The Process ID (PID) of the process.
// Parameters:
//  - value : Value to set for the processId property.
func (m *Process) SetProcessId(value *int32)() {
    m.processId = value
}
