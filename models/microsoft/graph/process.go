package graph

import (
    i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e "time"
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Process provides operations to manage the security singleton.
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
    fileHash FileHashable;
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
// NewProcess instantiates a new process and sets the default values.
func NewProcess()(*Process) {
    m := &Process{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateProcessFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProcessFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewProcess(), nil
}
// GetAccountName gets the accountName property value. User account identifier (user account context the process ran under) for example, AccountName, SID, and so on.
func (m *Process) GetAccountName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.accountName
    }
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Process) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCommandLine gets the commandLine property value. The full process invocation commandline including all parameters.
func (m *Process) GetCommandLine()(*string) {
    if m == nil {
        return nil
    } else {
        return m.commandLine
    }
}
// GetCreatedDateTime gets the createdDateTime property value. Time at which the process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Process) GetCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.createdDateTime
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Process) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["accountName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAccountName(val)
        }
        return nil
    }
    res["commandLine"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCommandLine(val)
        }
        return nil
    }
    res["createdDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCreatedDateTime(val)
        }
        return nil
    }
    res["fileHash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateFileHashFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetFileHash(val.(FileHashable))
        }
        return nil
    }
    res["integrityLevel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseProcessIntegrityLevel)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIntegrityLevel(val.(*ProcessIntegrityLevel))
        }
        return nil
    }
    res["isElevated"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIsElevated(val)
        }
        return nil
    }
    res["name"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetName(val)
        }
        return nil
    }
    res["parentProcessCreatedDateTime"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetTimeValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentProcessCreatedDateTime(val)
        }
        return nil
    }
    res["parentProcessId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentProcessId(val)
        }
        return nil
    }
    res["parentProcessName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetParentProcessName(val)
        }
        return nil
    }
    res["path"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPath(val)
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
    return res
}
// GetFileHash gets the fileHash property value. Complex type containing file hashes (cryptographic and location-sensitive).
func (m *Process) GetFileHash()(FileHashable) {
    if m == nil {
        return nil
    } else {
        return m.fileHash
    }
}
// GetIntegrityLevel gets the integrityLevel property value. The integrity level of the process. Possible values are: unknown, untrusted, low, medium, high, system.
func (m *Process) GetIntegrityLevel()(*ProcessIntegrityLevel) {
    if m == nil {
        return nil
    } else {
        return m.integrityLevel
    }
}
// GetIsElevated gets the isElevated property value. True if the process is elevated.
func (m *Process) GetIsElevated()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isElevated
    }
}
// GetName gets the name property value. The name of the process' Image file.
func (m *Process) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
// GetParentProcessCreatedDateTime gets the parentProcessCreatedDateTime property value. DateTime at which the parent process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Process) GetParentProcessCreatedDateTime()(*i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time) {
    if m == nil {
        return nil
    } else {
        return m.parentProcessCreatedDateTime
    }
}
// GetParentProcessId gets the parentProcessId property value. The Process ID (PID) of the parent process.
func (m *Process) GetParentProcessId()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.parentProcessId
    }
}
// GetParentProcessName gets the parentProcessName property value. The name of the image file of the parent process.
func (m *Process) GetParentProcessName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.parentProcessName
    }
}
// GetPath gets the path property value. Full path, including filename.
func (m *Process) GetPath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.path
    }
}
// GetProcessId gets the processId property value. The Process ID (PID) of the process.
func (m *Process) GetProcessId()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.processId
    }
}
func (m *Process) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
        cast := (*m.GetIntegrityLevel()).String()
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
// SetAccountName sets the accountName property value. User account identifier (user account context the process ran under) for example, AccountName, SID, and so on.
func (m *Process) SetAccountName(value *string)() {
    if m != nil {
        m.accountName = value
    }
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Process) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCommandLine sets the commandLine property value. The full process invocation commandline including all parameters.
func (m *Process) SetCommandLine(value *string)() {
    if m != nil {
        m.commandLine = value
    }
}
// SetCreatedDateTime sets the createdDateTime property value. Time at which the process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Process) SetCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.createdDateTime = value
    }
}
// SetFileHash sets the fileHash property value. Complex type containing file hashes (cryptographic and location-sensitive).
func (m *Process) SetFileHash(value FileHashable)() {
    if m != nil {
        m.fileHash = value
    }
}
// SetIntegrityLevel sets the integrityLevel property value. The integrity level of the process. Possible values are: unknown, untrusted, low, medium, high, system.
func (m *Process) SetIntegrityLevel(value *ProcessIntegrityLevel)() {
    if m != nil {
        m.integrityLevel = value
    }
}
// SetIsElevated sets the isElevated property value. True if the process is elevated.
func (m *Process) SetIsElevated(value *bool)() {
    if m != nil {
        m.isElevated = value
    }
}
// SetName sets the name property value. The name of the process' Image file.
func (m *Process) SetName(value *string)() {
    if m != nil {
        m.name = value
    }
}
// SetParentProcessCreatedDateTime sets the parentProcessCreatedDateTime property value. DateTime at which the parent process was started. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
func (m *Process) SetParentProcessCreatedDateTime(value *i336074805fc853987abe6f7fe3ad97a6a6f3077a16391fec744f671a015fbd7e.Time)() {
    if m != nil {
        m.parentProcessCreatedDateTime = value
    }
}
// SetParentProcessId sets the parentProcessId property value. The Process ID (PID) of the parent process.
func (m *Process) SetParentProcessId(value *int32)() {
    if m != nil {
        m.parentProcessId = value
    }
}
// SetParentProcessName sets the parentProcessName property value. The name of the image file of the parent process.
func (m *Process) SetParentProcessName(value *string)() {
    if m != nil {
        m.parentProcessName = value
    }
}
// SetPath sets the path property value. Full path, including filename.
func (m *Process) SetPath(value *string)() {
    if m != nil {
        m.path = value
    }
}
// SetProcessId sets the processId property value. The Process ID (PID) of the process.
func (m *Process) SetProcessId(value *int32)() {
    if m != nil {
        m.processId = value
    }
}
