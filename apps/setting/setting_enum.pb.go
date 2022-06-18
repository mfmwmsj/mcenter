// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package setting

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseSMS_PROVIDERFromString Parse SMS_PROVIDER from string
func ParseSMS_PROVIDERFromString(str string) (SMS_PROVIDER, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := SMS_PROVIDER_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown SMS_PROVIDER: %s", str)
	}

	return SMS_PROVIDER(v), nil
}

// Equal type compare
func (t SMS_PROVIDER) Equal(target SMS_PROVIDER) bool {
	return t == target
}

// IsIn todo
func (t SMS_PROVIDER) IsIn(targets ...SMS_PROVIDER) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t SMS_PROVIDER) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *SMS_PROVIDER) UnmarshalJSON(b []byte) error {
	ins, err := ParseSMS_PROVIDERFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseNOTIFY_TYPEFromString Parse NOTIFY_TYPE from string
func ParseNOTIFY_TYPEFromString(str string) (NOTIFY_TYPE, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := NOTIFY_TYPE_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown NOTIFY_TYPE: %s", str)
	}

	return NOTIFY_TYPE(v), nil
}

// Equal type compare
func (t NOTIFY_TYPE) Equal(target NOTIFY_TYPE) bool {
	return t == target
}

// IsIn todo
func (t NOTIFY_TYPE) IsIn(targets ...NOTIFY_TYPE) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t NOTIFY_TYPE) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *NOTIFY_TYPE) UnmarshalJSON(b []byte) error {
	ins, err := ParseNOTIFY_TYPEFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
