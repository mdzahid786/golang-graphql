package custom

import (
	"fmt"
	"io"
	"strconv"
)

type IntID int

// UnmarshalGQL implements the graphql scalar unmarshaler
func (i *IntID) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		n, err := strconv.Atoi(v)
		if err != nil {
			return err
		}
		*i = IntID(n)
		return nil
	case int:
		*i = IntID(v)
		return nil
	default:
		return fmt.Errorf("unexpected type for ID: %T", v)
	}
}

// MarshalGQL implements the graphql scalar marshaler
func (i IntID) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Itoa(int(i)))
}
