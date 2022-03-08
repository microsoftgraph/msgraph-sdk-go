package archive

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// ArchiveRequestBody provides operations to call the archive method.
type ArchiveRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    shouldSetSpoSiteReadOnlyForMembers *bool;
}
// NewArchiveRequestBody instantiates a new archiveRequestBody and sets the default values.
func NewArchiveRequestBody()(*ArchiveRequestBody) {
    m := &ArchiveRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateArchiveRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateArchiveRequestBodyFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewArchiveRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ArchiveRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ArchiveRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["shouldSetSpoSiteReadOnlyForMembers"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetShouldSetSpoSiteReadOnlyForMembers(val)
        }
        return nil
    }
    return res
}
// GetShouldSetSpoSiteReadOnlyForMembers gets the shouldSetSpoSiteReadOnlyForMembers property value. 
func (m *ArchiveRequestBody) GetShouldSetSpoSiteReadOnlyForMembers()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.shouldSetSpoSiteReadOnlyForMembers
    }
}
func (m *ArchiveRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *ArchiveRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("shouldSetSpoSiteReadOnlyForMembers", m.GetShouldSetSpoSiteReadOnlyForMembers())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ArchiveRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetShouldSetSpoSiteReadOnlyForMembers sets the shouldSetSpoSiteReadOnlyForMembers property value. 
func (m *ArchiveRequestBody) SetShouldSetSpoSiteReadOnlyForMembers(value *bool)() {
    if m != nil {
        m.shouldSetSpoSiteReadOnlyForMembers = value
    }
}
