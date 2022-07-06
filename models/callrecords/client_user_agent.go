package callrecords

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClientUserAgent 
type ClientUserAgent struct {
    UserAgent
    // Identifies the platform used by this endpoint. Possible values are: unknown, windows, macOS, iOS, android, web, ipPhone, roomSystem, surfaceHub, holoLens, unknownFutureValue.
    platform *ClientPlatform
    // Identifies the family of application software used by this endpoint. Possible values are: unknown, teams, skypeForBusiness, lync, unknownFutureValue, azureCommunicationServices.  Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: azureCommunicationServices.
    productFamily *ProductFamily
}
// NewClientUserAgent instantiates a new ClientUserAgent and sets the default values.
func NewClientUserAgent()(*ClientUserAgent) {
    m := &ClientUserAgent{
        UserAgent: *NewUserAgent(),
    }
    return m
}
// CreateClientUserAgentFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClientUserAgentFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClientUserAgent(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ClientUserAgent) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.UserAgent.GetFieldDeserializers()
    res["platform"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseClientPlatform)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPlatform(val.(*ClientPlatform))
        }
        return nil
    }
    res["productFamily"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProductFamily)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProductFamily(val.(*ProductFamily))
        }
        return nil
    }
    return res
}
// GetPlatform gets the platform property value. Identifies the platform used by this endpoint. Possible values are: unknown, windows, macOS, iOS, android, web, ipPhone, roomSystem, surfaceHub, holoLens, unknownFutureValue.
func (m *ClientUserAgent) GetPlatform()(*ClientPlatform) {
    if m == nil {
        return nil
    } else {
        return m.platform
    }
}
// GetProductFamily gets the productFamily property value. Identifies the family of application software used by this endpoint. Possible values are: unknown, teams, skypeForBusiness, lync, unknownFutureValue, azureCommunicationServices.  Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: azureCommunicationServices.
func (m *ClientUserAgent) GetProductFamily()(*ProductFamily) {
    if m == nil {
        return nil
    } else {
        return m.productFamily
    }
}
// Serialize serializes information the current object
func (m *ClientUserAgent) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.UserAgent.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetPlatform() != nil {
        cast := (*m.GetPlatform()).String()
        err = writer.WriteStringValue("platform", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetProductFamily() != nil {
        cast := (*m.GetProductFamily()).String()
        err = writer.WriteStringValue("productFamily", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetPlatform sets the platform property value. Identifies the platform used by this endpoint. Possible values are: unknown, windows, macOS, iOS, android, web, ipPhone, roomSystem, surfaceHub, holoLens, unknownFutureValue.
func (m *ClientUserAgent) SetPlatform(value *ClientPlatform)() {
    if m != nil {
        m.platform = value
    }
}
// SetProductFamily sets the productFamily property value. Identifies the family of application software used by this endpoint. Possible values are: unknown, teams, skypeForBusiness, lync, unknownFutureValue, azureCommunicationServices.  Note that you must use the Prefer: include-unknown-enum-members request header to get the following value(s) in this evolvable enum: azureCommunicationServices.
func (m *ClientUserAgent) SetProductFamily(value *ProductFamily)() {
    if m != nil {
        m.productFamily = value
    }
}
