package domain

import (
	"encoding/json"

	"github.com/satori/go.uuid"
)

//https://github.com/cockroachdb/cockroach/blob/master/pkg/util/uuid/uuid.go
type UID string

func (id UID) String() string {
	suid := string(id)
	return suid
}

func ToString(id UID) string {
	suid := id.String()
	return suid
}

func StringToID(s string) (*UID, error) {
	IDObj := UID(s)
	return &IDObj, nil
}

//Marshaller Interface
func (i UID) MarshalJSON() ([]byte, error) {

	out := i.String()
	outbytes, err := json.Marshal(out)
	return outbytes, err
}

//UnMarshaller Interface
func (i *UID) UnmarshalJSON(jsonb []byte) error {

	var uuidString string
	if err := json.Unmarshal(jsonb, &uuidString); err != nil {
		return err
	}
	uuid, err := StringToID(uuidString)
	*i = *uuid
	return err
}

func NewUID() UID {
	id, _ := uuid.NewV4()
	return UID(id.String())
}

func IsNil(uid UID) bool {
	if uid == "" {
		return true
	} else {
		return false
	}

}
