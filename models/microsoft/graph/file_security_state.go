package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type FileSecurityState struct {
    additionalData map[string]interface{};
    fileHash *FileHash;
    name *string;
    path *string;
    riskScore *string;
}
func NewFileSecurityState()(*FileSecurityState) {
    m := &FileSecurityState{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *FileSecurityState) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *FileSecurityState) GetFileHash()(*FileHash) {
    if m == nil {
        return nil
    } else {
        return m.fileHash
    }
}
func (m *FileSecurityState) GetName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.name
    }
}
func (m *FileSecurityState) GetPath()(*string) {
    if m == nil {
        return nil
    } else {
        return m.path
    }
}
func (m *FileSecurityState) GetRiskScore()(*string) {
    if m == nil {
        return nil
    } else {
        return m.riskScore
    }
}
func (m *FileSecurityState) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["fileHash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewFileHash() })
        if err != nil {
            return err
        }
        m.SetFileHash(val.(*FileHash))
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
    res["path"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPath(val)
        return nil
    }
    res["riskScore"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetRiskScore(val)
        return nil
    }
    return res
}
func (m *FileSecurityState) IsNil()(bool) {
    return m == nil
}
func (m *FileSecurityState) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("fileHash", m.GetFileHash())
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
        err := writer.WriteStringValue("path", m.GetPath())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("riskScore", m.GetRiskScore())
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
func (m *FileSecurityState) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *FileSecurityState) SetFileHash(value *FileHash)() {
    m.fileHash = value
}
func (m *FileSecurityState) SetName(value *string)() {
    m.name = value
}
func (m *FileSecurityState) SetPath(value *string)() {
    m.path = value
}
func (m *FileSecurityState) SetRiskScore(value *string)() {
    m.riskScore = value
}
