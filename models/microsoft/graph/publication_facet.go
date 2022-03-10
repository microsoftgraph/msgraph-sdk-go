package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PublicationFacet provides operations to manage the drive singleton.
type PublicationFacet struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The state of publication for this document. Either published or checkout. Read-only.
    level *string;
    // The unique identifier for the version that is visible to the current caller. Read-only.
    versionId *string;
}
// NewPublicationFacet instantiates a new publicationFacet and sets the default values.
func NewPublicationFacet()(*PublicationFacet) {
    m := &PublicationFacet{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreatePublicationFacetFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreatePublicationFacetFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewPublicationFacet(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PublicationFacet) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PublicationFacet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["level"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLevel(val)
        }
        return nil
    }
    res["versionId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetVersionId(val)
        }
        return nil
    }
    return res
}
// GetLevel gets the level property value. The state of publication for this document. Either published or checkout. Read-only.
func (m *PublicationFacet) GetLevel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.level
    }
}
// GetVersionId gets the versionId property value. The unique identifier for the version that is visible to the current caller. Read-only.
func (m *PublicationFacet) GetVersionId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.versionId
    }
}
func (m *PublicationFacet) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *PublicationFacet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("level", m.GetLevel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("versionId", m.GetVersionId())
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
func (m *PublicationFacet) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetLevel sets the level property value. The state of publication for this document. Either published or checkout. Read-only.
func (m *PublicationFacet) SetLevel(value *string)() {
    if m != nil {
        m.level = value
    }
}
// SetVersionId sets the versionId property value. The unique identifier for the version that is visible to the current caller. Read-only.
func (m *PublicationFacet) SetVersionId(value *string)() {
    if m != nil {
        m.versionId = value
    }
}
