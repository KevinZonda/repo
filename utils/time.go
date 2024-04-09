package utils

import "time"

type Date time.Time

func (d *Date) UnmarshalJSON(bytes []byte) error {
	dd, err := time.Parse(`"2006-01-02"`, string(bytes))
	if err != nil {
		return err
	}
	*d = Date(dd)

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(d).Format(`2006-01-02`) + `"`), nil
}
