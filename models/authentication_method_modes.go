package models
import (
    "errors"
    "strings"
)
// 
type AuthenticationMethodModes int

const (
    PASSWORD_AUTHENTICATIONMETHODMODES AuthenticationMethodModes = iota
    VOICE_AUTHENTICATIONMETHODMODES
    HARDWAREOATH_AUTHENTICATIONMETHODMODES
    SOFTWAREOATH_AUTHENTICATIONMETHODMODES
    SMS_AUTHENTICATIONMETHODMODES
    FIDO2_AUTHENTICATIONMETHODMODES
    WINDOWSHELLOFORBUSINESS_AUTHENTICATIONMETHODMODES
    MICROSOFTAUTHENTICATORPUSH_AUTHENTICATIONMETHODMODES
    DEVICEBASEDPUSH_AUTHENTICATIONMETHODMODES
    TEMPORARYACCESSPASSONETIME_AUTHENTICATIONMETHODMODES
    TEMPORARYACCESSPASSMULTIUSE_AUTHENTICATIONMETHODMODES
    EMAIL_AUTHENTICATIONMETHODMODES
    X509CERTIFICATESINGLEFACTOR_AUTHENTICATIONMETHODMODES
    X509CERTIFICATEMULTIFACTOR_AUTHENTICATIONMETHODMODES
    FEDERATEDSINGLEFACTOR_AUTHENTICATIONMETHODMODES
    FEDERATEDMULTIFACTOR_AUTHENTICATIONMETHODMODES
    UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODMODES
)

func (i AuthenticationMethodModes) String() string {
    var values []string
    for p := AuthenticationMethodModes(1); p <= UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODMODES; p <<= 1 {
        if i&p == p {
            values = append(values, []string{"password", "voice", "hardwareOath", "softwareOath", "sms", "fido2", "windowsHelloForBusiness", "microsoftAuthenticatorPush", "deviceBasedPush", "temporaryAccessPassOneTime", "temporaryAccessPassMultiUse", "email", "x509CertificateSingleFactor", "x509CertificateMultiFactor", "federatedSingleFactor", "federatedMultiFactor", "unknownFutureValue"}[p])
        }
    }
    return strings.Join(values, ",")
}
func ParseAuthenticationMethodModes(v string) (any, error) {
    var result AuthenticationMethodModes
    values := strings.Split(v, ",")
    for _, str := range values {
        switch str {
            case "password":
                result |= PASSWORD_AUTHENTICATIONMETHODMODES
            case "voice":
                result |= VOICE_AUTHENTICATIONMETHODMODES
            case "hardwareOath":
                result |= HARDWAREOATH_AUTHENTICATIONMETHODMODES
            case "softwareOath":
                result |= SOFTWAREOATH_AUTHENTICATIONMETHODMODES
            case "sms":
                result |= SMS_AUTHENTICATIONMETHODMODES
            case "fido2":
                result |= FIDO2_AUTHENTICATIONMETHODMODES
            case "windowsHelloForBusiness":
                result |= WINDOWSHELLOFORBUSINESS_AUTHENTICATIONMETHODMODES
            case "microsoftAuthenticatorPush":
                result |= MICROSOFTAUTHENTICATORPUSH_AUTHENTICATIONMETHODMODES
            case "deviceBasedPush":
                result |= DEVICEBASEDPUSH_AUTHENTICATIONMETHODMODES
            case "temporaryAccessPassOneTime":
                result |= TEMPORARYACCESSPASSONETIME_AUTHENTICATIONMETHODMODES
            case "temporaryAccessPassMultiUse":
                result |= TEMPORARYACCESSPASSMULTIUSE_AUTHENTICATIONMETHODMODES
            case "email":
                result |= EMAIL_AUTHENTICATIONMETHODMODES
            case "x509CertificateSingleFactor":
                result |= X509CERTIFICATESINGLEFACTOR_AUTHENTICATIONMETHODMODES
            case "x509CertificateMultiFactor":
                result |= X509CERTIFICATEMULTIFACTOR_AUTHENTICATIONMETHODMODES
            case "federatedSingleFactor":
                result |= FEDERATEDSINGLEFACTOR_AUTHENTICATIONMETHODMODES
            case "federatedMultiFactor":
                result |= FEDERATEDMULTIFACTOR_AUTHENTICATIONMETHODMODES
            case "unknownFutureValue":
                result |= UNKNOWNFUTUREVALUE_AUTHENTICATIONMETHODMODES
            default:
                return 0, errors.New("Unknown AuthenticationMethodModes value: " + v)
        }
    }
    return &result, nil
}
func SerializeAuthenticationMethodModes(values []AuthenticationMethodModes) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i AuthenticationMethodModes) isMultiValue() bool {
    return true
}
