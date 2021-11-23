package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// hashes 
type Hashes struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The CRC32 value of the file in little endian (if available). Read-only.
    crc32Hash *string;
    // A proprietary hash of the file that can be used to determine if the contents of the file have changed (if available). Read-only.
    quickXorHash *string;
    // SHA1 hash for the contents of the file (if available). Read-only.
    sha1Hash *string;
    // SHA256 hash for the contents of the file (if available). Read-only.
    sha256Hash *string;
}
// NewHashes instantiates a new hashes and sets the default values.
func NewHashes()(*Hashes) {
    m := &Hashes{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Hashes) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCrc32Hash gets the crc32Hash property value. The CRC32 value of the file in little endian (if available). Read-only.
func (m *Hashes) GetCrc32Hash()(*string) {
    if m == nil {
        return nil
    } else {
        return m.crc32Hash
    }
}
// GetQuickXorHash gets the quickXorHash property value. A proprietary hash of the file that can be used to determine if the contents of the file have changed (if available). Read-only.
func (m *Hashes) GetQuickXorHash()(*string) {
    if m == nil {
        return nil
    } else {
        return m.quickXorHash
    }
}
// GetSha1Hash gets the sha1Hash property value. SHA1 hash for the contents of the file (if available). Read-only.
func (m *Hashes) GetSha1Hash()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sha1Hash
    }
}
// GetSha256Hash gets the sha256Hash property value. SHA256 hash for the contents of the file (if available). Read-only.
func (m *Hashes) GetSha256Hash()(*string) {
    if m == nil {
        return nil
    } else {
        return m.sha256Hash
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Hashes) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["crc32Hash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCrc32Hash(val)
        }
        return nil
    }
    res["quickXorHash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetQuickXorHash(val)
        }
        return nil
    }
    res["sha1Hash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha1Hash(val)
        }
        return nil
    }
    res["sha256Hash"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSha256Hash(val)
        }
        return nil
    }
    return res
}
func (m *Hashes) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Hashes) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetCrc32Hash sets the crc32Hash property value. The CRC32 value of the file in little endian (if available). Read-only.
func (m *Hashes) SetCrc32Hash(value *string)() {
    m.crc32Hash = value
}
// SetQuickXorHash sets the quickXorHash property value. A proprietary hash of the file that can be used to determine if the contents of the file have changed (if available). Read-only.
func (m *Hashes) SetQuickXorHash(value *string)() {
    m.quickXorHash = value
}
// SetSha1Hash sets the sha1Hash property value. SHA1 hash for the contents of the file (if available). Read-only.
func (m *Hashes) SetSha1Hash(value *string)() {
    m.sha1Hash = value
}
// SetSha256Hash sets the sha256Hash property value. SHA256 hash for the contents of the file (if available). Read-only.
func (m *Hashes) SetSha256Hash(value *string)() {
    m.sha256Hash = value
}
