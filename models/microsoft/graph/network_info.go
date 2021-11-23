package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/callrecords"
)

// NetworkInfo 
type NetworkInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Fraction of the call that the media endpoint detected the available bandwidth or bandwidth policy was low enough to cause poor quality of the audio sent.
    bandwidthLowEventRatio *float32;
    // The wireless LAN basic service set identifier of the media endpoint used to connect to the network.
    basicServiceSetIdentifier *string;
    // Type of network used by the media endpoint. Possible values are: unknown, wired, wifi, mobile, tunnel, unknownFutureValue.
    connectionType *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.NetworkConnectionType;
    // Fraction of the call that the media endpoint detected the network delay was significant enough to impact the ability to have real-time two-way communication.
    delayEventRatio *float32;
    // DNS suffix associated with the network adapter of the media endpoint.
    dnsSuffix *string;
    // IP address of the media endpoint.
    ipAddress *string;
    // Link speed in bits per second reported by the network adapter used by the media endpoint.
    linkSpeed *int64;
    // The media access control (MAC) address of the media endpoint's network device.
    macAddress *string;
    // Network port number used by media endpoint.
    port *int32;
    // Fraction of the call that the media endpoint detected the network was causing poor quality of the audio received.
    receivedQualityEventRatio *float32;
    // IP address of the media endpoint as seen by the media relay server. This is typically the public internet IP address associated to the endpoint.
    reflexiveIPAddress *string;
    // IP address of the media relay server allocated by the media endpoint.
    relayIPAddress *string;
    // Network port number allocated on the media relay server by the media endpoint.
    relayPort *int32;
    // Fraction of the call that the media endpoint detected the network was causing poor quality of the audio sent.
    sentQualityEventRatio *float32;
    // Subnet used for media stream by the media endpoint.
    subnet *string;
    // WiFi band used by the media endpoint. Possible values are: unknown, frequency24GHz, frequency50GHz, frequency60GHz, unknownFutureValue.
    wifiBand *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiBand;
    // Estimated remaining battery charge in percentage reported by the media endpoint.
    wifiBatteryCharge *int32;
    // WiFi channel used by the media endpoint.
    wifiChannel *int32;
    // Name of the Microsoft WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
    wifiMicrosoftDriver *string;
    // Version of the Microsoft WiFi driver used by the media endpoint.
    wifiMicrosoftDriverVersion *string;
    // Type of WiFi radio used by the media endpoint. Possible values are: unknown, wifi80211a, wifi80211b, wifi80211g, wifi80211n, wifi80211ac, wifi80211ax, unknownFutureValue.
    wifiRadioType *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiRadioType;
    // WiFi signal strength in percentage reported by the media endpoint.
    wifiSignalStrength *int32;
    // Name of the WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
    wifiVendorDriver *string;
    // Version of the WiFi driver used by the media endpoint.
    wifiVendorDriverVersion *string;
}
// NewNetworkInfo instantiates a new networkInfo and sets the default values.
func NewNetworkInfo()(*NetworkInfo) {
    m := &NetworkInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *NetworkInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetBandwidthLowEventRatio gets the bandwidthLowEventRatio property value. Fraction of the call that the media endpoint detected the available bandwidth or bandwidth policy was low enough to cause poor quality of the audio sent.
func (m *NetworkInfo) GetBandwidthLowEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.bandwidthLowEventRatio
    }
}
// GetBasicServiceSetIdentifier gets the basicServiceSetIdentifier property value. The wireless LAN basic service set identifier of the media endpoint used to connect to the network.
func (m *NetworkInfo) GetBasicServiceSetIdentifier()(*string) {
    if m == nil {
        return nil
    } else {
        return m.basicServiceSetIdentifier
    }
}
// GetConnectionType gets the connectionType property value. Type of network used by the media endpoint. Possible values are: unknown, wired, wifi, mobile, tunnel, unknownFutureValue.
func (m *NetworkInfo) GetConnectionType()(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.NetworkConnectionType) {
    if m == nil {
        return nil
    } else {
        return m.connectionType
    }
}
// GetDelayEventRatio gets the delayEventRatio property value. Fraction of the call that the media endpoint detected the network delay was significant enough to impact the ability to have real-time two-way communication.
func (m *NetworkInfo) GetDelayEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.delayEventRatio
    }
}
// GetDnsSuffix gets the dnsSuffix property value. DNS suffix associated with the network adapter of the media endpoint.
func (m *NetworkInfo) GetDnsSuffix()(*string) {
    if m == nil {
        return nil
    } else {
        return m.dnsSuffix
    }
}
// GetIpAddress gets the ipAddress property value. IP address of the media endpoint.
func (m *NetworkInfo) GetIpAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.ipAddress
    }
}
// GetLinkSpeed gets the linkSpeed property value. Link speed in bits per second reported by the network adapter used by the media endpoint.
func (m *NetworkInfo) GetLinkSpeed()(*int64) {
    if m == nil {
        return nil
    } else {
        return m.linkSpeed
    }
}
// GetMacAddress gets the macAddress property value. The media access control (MAC) address of the media endpoint's network device.
func (m *NetworkInfo) GetMacAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.macAddress
    }
}
// GetPort gets the port property value. Network port number used by media endpoint.
func (m *NetworkInfo) GetPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.port
    }
}
// GetReceivedQualityEventRatio gets the receivedQualityEventRatio property value. Fraction of the call that the media endpoint detected the network was causing poor quality of the audio received.
func (m *NetworkInfo) GetReceivedQualityEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.receivedQualityEventRatio
    }
}
// GetReflexiveIPAddress gets the reflexiveIPAddress property value. IP address of the media endpoint as seen by the media relay server. This is typically the public internet IP address associated to the endpoint.
func (m *NetworkInfo) GetReflexiveIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.reflexiveIPAddress
    }
}
// GetRelayIPAddress gets the relayIPAddress property value. IP address of the media relay server allocated by the media endpoint.
func (m *NetworkInfo) GetRelayIPAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.relayIPAddress
    }
}
// GetRelayPort gets the relayPort property value. Network port number allocated on the media relay server by the media endpoint.
func (m *NetworkInfo) GetRelayPort()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.relayPort
    }
}
// GetSentQualityEventRatio gets the sentQualityEventRatio property value. Fraction of the call that the media endpoint detected the network was causing poor quality of the audio sent.
func (m *NetworkInfo) GetSentQualityEventRatio()(*float32) {
    if m == nil {
        return nil
    } else {
        return m.sentQualityEventRatio
    }
}
// GetSubnet gets the subnet property value. Subnet used for media stream by the media endpoint.
func (m *NetworkInfo) GetSubnet()(*string) {
    if m == nil {
        return nil
    } else {
        return m.subnet
    }
}
// GetWifiBand gets the wifiBand property value. WiFi band used by the media endpoint. Possible values are: unknown, frequency24GHz, frequency50GHz, frequency60GHz, unknownFutureValue.
func (m *NetworkInfo) GetWifiBand()(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiBand) {
    if m == nil {
        return nil
    } else {
        return m.wifiBand
    }
}
// GetWifiBatteryCharge gets the wifiBatteryCharge property value. Estimated remaining battery charge in percentage reported by the media endpoint.
func (m *NetworkInfo) GetWifiBatteryCharge()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.wifiBatteryCharge
    }
}
// GetWifiChannel gets the wifiChannel property value. WiFi channel used by the media endpoint.
func (m *NetworkInfo) GetWifiChannel()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.wifiChannel
    }
}
// GetWifiMicrosoftDriver gets the wifiMicrosoftDriver property value. Name of the Microsoft WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
func (m *NetworkInfo) GetWifiMicrosoftDriver()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiMicrosoftDriver
    }
}
// GetWifiMicrosoftDriverVersion gets the wifiMicrosoftDriverVersion property value. Version of the Microsoft WiFi driver used by the media endpoint.
func (m *NetworkInfo) GetWifiMicrosoftDriverVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiMicrosoftDriverVersion
    }
}
// GetWifiRadioType gets the wifiRadioType property value. Type of WiFi radio used by the media endpoint. Possible values are: unknown, wifi80211a, wifi80211b, wifi80211g, wifi80211n, wifi80211ac, wifi80211ax, unknownFutureValue.
func (m *NetworkInfo) GetWifiRadioType()(*i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiRadioType) {
    if m == nil {
        return nil
    } else {
        return m.wifiRadioType
    }
}
// GetWifiSignalStrength gets the wifiSignalStrength property value. WiFi signal strength in percentage reported by the media endpoint.
func (m *NetworkInfo) GetWifiSignalStrength()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.wifiSignalStrength
    }
}
// GetWifiVendorDriver gets the wifiVendorDriver property value. Name of the WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
func (m *NetworkInfo) GetWifiVendorDriver()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiVendorDriver
    }
}
// GetWifiVendorDriverVersion gets the wifiVendorDriverVersion property value. Version of the WiFi driver used by the media endpoint.
func (m *NetworkInfo) GetWifiVendorDriverVersion()(*string) {
    if m == nil {
        return nil
    } else {
        return m.wifiVendorDriverVersion
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *NetworkInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["bandwidthLowEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBandwidthLowEventRatio(val)
        }
        return nil
    }
    res["basicServiceSetIdentifier"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetBasicServiceSetIdentifier(val)
        }
        return nil
    }
    res["connectionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.ParseNetworkConnectionType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.NetworkConnectionType)
            m.SetConnectionType(&cast)
        }
        return nil
    }
    res["delayEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDelayEventRatio(val)
        }
        return nil
    }
    res["dnsSuffix"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDnsSuffix(val)
        }
        return nil
    }
    res["ipAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetIpAddress(val)
        }
        return nil
    }
    res["linkSpeed"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt64Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetLinkSpeed(val)
        }
        return nil
    }
    res["macAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetMacAddress(val)
        }
        return nil
    }
    res["port"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPort(val)
        }
        return nil
    }
    res["receivedQualityEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReceivedQualityEventRatio(val)
        }
        return nil
    }
    res["reflexiveIPAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetReflexiveIPAddress(val)
        }
        return nil
    }
    res["relayIPAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelayIPAddress(val)
        }
        return nil
    }
    res["relayPort"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetRelayPort(val)
        }
        return nil
    }
    res["sentQualityEventRatio"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetFloat32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSentQualityEventRatio(val)
        }
        return nil
    }
    res["subnet"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSubnet(val)
        }
        return nil
    }
    res["wifiBand"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.ParseWifiBand)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiBand)
            m.SetWifiBand(&cast)
        }
        return nil
    }
    res["wifiBatteryCharge"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiBatteryCharge(val)
        }
        return nil
    }
    res["wifiChannel"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiChannel(val)
        }
        return nil
    }
    res["wifiMicrosoftDriver"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiMicrosoftDriver(val)
        }
        return nil
    }
    res["wifiMicrosoftDriverVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiMicrosoftDriverVersion(val)
        }
        return nil
    }
    res["wifiRadioType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.ParseWifiRadioType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiRadioType)
            m.SetWifiRadioType(&cast)
        }
        return nil
    }
    res["wifiSignalStrength"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiSignalStrength(val)
        }
        return nil
    }
    res["wifiVendorDriver"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiVendorDriver(val)
        }
        return nil
    }
    res["wifiVendorDriverVersion"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetWifiVendorDriverVersion(val)
        }
        return nil
    }
    return res
}
func (m *NetworkInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *NetworkInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteFloat32Value("bandwidthLowEventRatio", m.GetBandwidthLowEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("basicServiceSetIdentifier", m.GetBasicServiceSetIdentifier())
        if err != nil {
            return err
        }
    }
    if m.GetConnectionType() != nil {
        cast := m.GetConnectionType().String()
        err := writer.WriteStringValue("connectionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("delayEventRatio", m.GetDelayEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("dnsSuffix", m.GetDnsSuffix())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("ipAddress", m.GetIpAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt64Value("linkSpeed", m.GetLinkSpeed())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("macAddress", m.GetMacAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("port", m.GetPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("receivedQualityEventRatio", m.GetReceivedQualityEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("reflexiveIPAddress", m.GetReflexiveIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("relayIPAddress", m.GetRelayIPAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("relayPort", m.GetRelayPort())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteFloat32Value("sentQualityEventRatio", m.GetSentQualityEventRatio())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("subnet", m.GetSubnet())
        if err != nil {
            return err
        }
    }
    if m.GetWifiBand() != nil {
        cast := m.GetWifiBand().String()
        err := writer.WriteStringValue("wifiBand", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wifiBatteryCharge", m.GetWifiBatteryCharge())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wifiChannel", m.GetWifiChannel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiMicrosoftDriver", m.GetWifiMicrosoftDriver())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiMicrosoftDriverVersion", m.GetWifiMicrosoftDriverVersion())
        if err != nil {
            return err
        }
    }
    if m.GetWifiRadioType() != nil {
        cast := m.GetWifiRadioType().String()
        err := writer.WriteStringValue("wifiRadioType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteInt32Value("wifiSignalStrength", m.GetWifiSignalStrength())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiVendorDriver", m.GetWifiVendorDriver())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("wifiVendorDriverVersion", m.GetWifiVendorDriverVersion())
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
func (m *NetworkInfo) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetBandwidthLowEventRatio sets the bandwidthLowEventRatio property value. Fraction of the call that the media endpoint detected the available bandwidth or bandwidth policy was low enough to cause poor quality of the audio sent.
func (m *NetworkInfo) SetBandwidthLowEventRatio(value *float32)() {
    m.bandwidthLowEventRatio = value
}
// SetBasicServiceSetIdentifier sets the basicServiceSetIdentifier property value. The wireless LAN basic service set identifier of the media endpoint used to connect to the network.
func (m *NetworkInfo) SetBasicServiceSetIdentifier(value *string)() {
    m.basicServiceSetIdentifier = value
}
// SetConnectionType sets the connectionType property value. Type of network used by the media endpoint. Possible values are: unknown, wired, wifi, mobile, tunnel, unknownFutureValue.
func (m *NetworkInfo) SetConnectionType(value *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.NetworkConnectionType)() {
    m.connectionType = value
}
// SetDelayEventRatio sets the delayEventRatio property value. Fraction of the call that the media endpoint detected the network delay was significant enough to impact the ability to have real-time two-way communication.
func (m *NetworkInfo) SetDelayEventRatio(value *float32)() {
    m.delayEventRatio = value
}
// SetDnsSuffix sets the dnsSuffix property value. DNS suffix associated with the network adapter of the media endpoint.
func (m *NetworkInfo) SetDnsSuffix(value *string)() {
    m.dnsSuffix = value
}
// SetIpAddress sets the ipAddress property value. IP address of the media endpoint.
func (m *NetworkInfo) SetIpAddress(value *string)() {
    m.ipAddress = value
}
// SetLinkSpeed sets the linkSpeed property value. Link speed in bits per second reported by the network adapter used by the media endpoint.
func (m *NetworkInfo) SetLinkSpeed(value *int64)() {
    m.linkSpeed = value
}
// SetMacAddress sets the macAddress property value. The media access control (MAC) address of the media endpoint's network device.
func (m *NetworkInfo) SetMacAddress(value *string)() {
    m.macAddress = value
}
// SetPort sets the port property value. Network port number used by media endpoint.
func (m *NetworkInfo) SetPort(value *int32)() {
    m.port = value
}
// SetReceivedQualityEventRatio sets the receivedQualityEventRatio property value. Fraction of the call that the media endpoint detected the network was causing poor quality of the audio received.
func (m *NetworkInfo) SetReceivedQualityEventRatio(value *float32)() {
    m.receivedQualityEventRatio = value
}
// SetReflexiveIPAddress sets the reflexiveIPAddress property value. IP address of the media endpoint as seen by the media relay server. This is typically the public internet IP address associated to the endpoint.
func (m *NetworkInfo) SetReflexiveIPAddress(value *string)() {
    m.reflexiveIPAddress = value
}
// SetRelayIPAddress sets the relayIPAddress property value. IP address of the media relay server allocated by the media endpoint.
func (m *NetworkInfo) SetRelayIPAddress(value *string)() {
    m.relayIPAddress = value
}
// SetRelayPort sets the relayPort property value. Network port number allocated on the media relay server by the media endpoint.
func (m *NetworkInfo) SetRelayPort(value *int32)() {
    m.relayPort = value
}
// SetSentQualityEventRatio sets the sentQualityEventRatio property value. Fraction of the call that the media endpoint detected the network was causing poor quality of the audio sent.
func (m *NetworkInfo) SetSentQualityEventRatio(value *float32)() {
    m.sentQualityEventRatio = value
}
// SetSubnet sets the subnet property value. Subnet used for media stream by the media endpoint.
func (m *NetworkInfo) SetSubnet(value *string)() {
    m.subnet = value
}
// SetWifiBand sets the wifiBand property value. WiFi band used by the media endpoint. Possible values are: unknown, frequency24GHz, frequency50GHz, frequency60GHz, unknownFutureValue.
func (m *NetworkInfo) SetWifiBand(value *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiBand)() {
    m.wifiBand = value
}
// SetWifiBatteryCharge sets the wifiBatteryCharge property value. Estimated remaining battery charge in percentage reported by the media endpoint.
func (m *NetworkInfo) SetWifiBatteryCharge(value *int32)() {
    m.wifiBatteryCharge = value
}
// SetWifiChannel sets the wifiChannel property value. WiFi channel used by the media endpoint.
func (m *NetworkInfo) SetWifiChannel(value *int32)() {
    m.wifiChannel = value
}
// SetWifiMicrosoftDriver sets the wifiMicrosoftDriver property value. Name of the Microsoft WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
func (m *NetworkInfo) SetWifiMicrosoftDriver(value *string)() {
    m.wifiMicrosoftDriver = value
}
// SetWifiMicrosoftDriverVersion sets the wifiMicrosoftDriverVersion property value. Version of the Microsoft WiFi driver used by the media endpoint.
func (m *NetworkInfo) SetWifiMicrosoftDriverVersion(value *string)() {
    m.wifiMicrosoftDriverVersion = value
}
// SetWifiRadioType sets the wifiRadioType property value. Type of WiFi radio used by the media endpoint. Possible values are: unknown, wifi80211a, wifi80211b, wifi80211g, wifi80211n, wifi80211ac, wifi80211ax, unknownFutureValue.
func (m *NetworkInfo) SetWifiRadioType(value *i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.WifiRadioType)() {
    m.wifiRadioType = value
}
// SetWifiSignalStrength sets the wifiSignalStrength property value. WiFi signal strength in percentage reported by the media endpoint.
func (m *NetworkInfo) SetWifiSignalStrength(value *int32)() {
    m.wifiSignalStrength = value
}
// SetWifiVendorDriver sets the wifiVendorDriver property value. Name of the WiFi driver used by the media endpoint. Value may be localized based on the language used by endpoint.
func (m *NetworkInfo) SetWifiVendorDriver(value *string)() {
    m.wifiVendorDriver = value
}
// SetWifiVendorDriverVersion sets the wifiVendorDriverVersion property value. Version of the WiFi driver used by the media endpoint.
func (m *NetworkInfo) SetWifiVendorDriverVersion(value *string)() {
    m.wifiVendorDriverVersion = value
}
