package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type Hashes struct {
    additionalData map[string]interface{};
    crc32Hash *string;
    quickXorHash *string;
    sha1Hash *string;
    sha256Hash *string;
}
func NewHashes()(*Hashes) {
    m := &Hashes{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *Hashes) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *Hashes) GetCrc32Hash()(*string) {
    if m == nil {
        return nil
    } else {
        return m.crc32Hash
    }
}
func (m *Hashes) GetQuickXorHash()(*string) {
    if m == nil {
        return nil
    } else {
        return m.quickXorHash
    }
}
func (m *Hashes) GetSha1Hash()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sha1Hash
    }
}
func (m *Hashes) GetSha256Hash()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sha256Hash
    }
}
func (m *Hashes) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["crc32Hash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetCrc32Hash(val)
        return nil
    }
    res["quickXorHash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetQuickXorHash(val)
        return nil
    }
    res["sha1Hash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSha1Hash(val)
        return nil
    }
    res["sha256Hash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetSha256Hash(val)
        return nil
    }
    return res
}
func (m *Hashes) IsNil()(bool) {
    return m == nil
}
func (m *Hashes) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("crc32Hash", m.GetCrc32Hash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("quickXorHash", m.GetQuickXorHash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha1Hash", m.GetSha1Hash())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("sha256Hash", m.GetSha256Hash())
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
func (m *Hashes) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *Hashes) SetCrc32Hash(value *string)() {
    m.crc32Hash = value
}
func (m *Hashes) SetQuickXorHash(value *string)() {
    m.quickXorHash = value
}
func (m *Hashes) SetSha1Hash(value *string)() {
    m.sha1Hash = value
}
func (m *Hashes) SetSha256Hash(value *string)() {
    m.sha256Hash = value
}
