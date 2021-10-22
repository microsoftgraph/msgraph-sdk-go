package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type FileHash struct {
    additionalData map[string]interface{};
    hashType *FileHashType;
    hashValue *string;
}
func NewFileHash()(*FileHash) {
    m := &FileHash{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *FileHash) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *FileHash) GetHashType()(*FileHashType) {
    if m == nil {
        return nil
    } else {
        return m.hashType
    }
}
func (m *FileHash) GetHashValue()(*string) {
    if m == nil {
        return nil
    } else {
        return m.hashValue
    }
}
func (m *FileHash) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hashType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseFileHashType)
        if err != nil {
            return err
        }
        cast := val.(FileHashType)
        m.SetHashType(&cast)
        return nil
    }
    res["hashValue"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetHashValue(val)
        return nil
    }
    return res
}
func (m *FileHash) IsNil()(bool) {
    return m == nil
}
func (m *FileHash) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    if m.GetHashType() != nil {
        cast := m.GetHashType().String()
        err := writer.WriteStringValue("hashType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("hashValue", m.GetHashValue())
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
func (m *FileHash) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *FileHash) SetHashType(value *FileHashType)() {
    m.hashType = value
}
func (m *FileHash) SetHashValue(value *string)() {
    m.hashValue = value
}
