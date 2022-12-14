package security

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// IpEvidence 
type IpEvidence struct {
    AlertEvidence
    // The countryLetterCode property
    countryLetterCode *string
    // The ipAddress property
    ipAddress *string
}
// NewIpEvidence instantiates a new IpEvidence and sets the default values.
func NewIpEvidence()(*IpEvidence) {
    m := &IpEvidence{
        AlertEvidence: *NewAlertEvidence(),
    }
    return m
}
// CreateIpEvidenceFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateIpEvidenceFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewIpEvidence(), nil
}
// GetCountryLetterCode gets the countryLetterCode property value. The countryLetterCode property
func (m *IpEvidence) GetCountryLetterCode()(*string) {
    return m.countryLetterCode
}
// GetFieldDeserializers the deserialization information for the current model
func (m *IpEvidence) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.AlertEvidence.GetFieldDeserializers()
    res["countryLetterCode"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCountryLetterCode(val)
        }
        return nil
    }
    res["ipAddress"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    return res
}
// GetIpAddress gets the ipAddress property value. The ipAddress property
func (m *IpEvidence) GetIpAddress()(*string) {
    return m.ipAddress
}
// Serialize serializes information the current object
func (m *IpEvidence) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.AlertEvidence.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("countryLetterCode", m.GetCountryLetterCode())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCountryLetterCode sets the countryLetterCode property value. The countryLetterCode property
func (m *IpEvidence) SetCountryLetterCode(value *string)() {
    m.countryLetterCode = value
}
// SetIpAddress sets the ipAddress property value. The ipAddress property
func (m *IpEvidence) SetIpAddress(value *string)() {
    m.ipAddress = value
}
