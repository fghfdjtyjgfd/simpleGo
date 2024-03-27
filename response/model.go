package response

import "encoding/json"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// MarshalJSON custom content json
func (r Response) MarshalJSON(Message string) ([]byte, error) {
	type Alias Response
	iModel := &struct {
		*Alias
	}{
		Alias: (*Alias)(&r),
	}

	return json.Marshal(iModel)
}
