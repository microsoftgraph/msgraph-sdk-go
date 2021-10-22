package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type File struct {
    additionalData map[string]interface{};
    hashes *Hashes;
    mimeType *string;
    processingMetadata *bool;
}
func NewFile()(*File) {
    m := &File{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *File) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *File) GetHashes()(*Hashes) {
    if m == nil {
        return nil
    } else {
        return m.hashes
    }
}
func (m *File) GetMimeType()(*string) {
    if m == nil {
        return nil
    } else {
        return m.mimeType
    }
}
func (m *File) GetProcessingMetadata()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.processingMetadata
    }
}
func (m *File) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["hashes"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewHashes() })
        if err != nil {
            return err
        }
        m.SetHashes(val.(*Hashes))
        return nil
    }
    res["mimeType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetMimeType(val)
        return nil
    }
    res["processingMetadata"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetProcessingMetadata(val)
        return nil
    }
    return res
}
func (m *File) IsNil()(bool) {
    return m == nil
}
func (m *File) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("hashes", m.GetHashes())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("mimeType", m.GetMimeType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("processingMetadata", m.GetProcessingMetadata())
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
func (m *File) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *File) SetHashes(value *Hashes)() {
    m.hashes = value
}
func (m *File) SetMimeType(value *string)() {
    m.mimeType = value
}
func (m *File) SetProcessingMetadata(value *bool)() {
    m.processingMetadata = value
}
