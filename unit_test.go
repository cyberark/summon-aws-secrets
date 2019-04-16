package main

import (
	"reflect"
	"testing"
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/secretsmanager"
	. "github.com/smartystreets/goconvey/convey"
)

func AreEqualJSON(s1, s2 string) bool {
	var o1 interface{}
	var o2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		panic(err)
	}

	return reflect.DeepEqual(o1, o2)
}

func Test_RetrieveSecret(t *testing.T) {
	Convey("Given a valid JSON format stored secret", t, func() {
		secret := `
{
 "a": 1,
 "b": "xyz"
}
`

		Convey("retrieve via secret name", func() {
			secretBytes, _ := RetrieveSecret("production/secret")
			So(AreEqualJSON(string(secretBytes), secret), ShouldBeTrue)
		})

		Convey("retrieve via secret key path", func() {
			secretBytes, _ := RetrieveSecret("production/secret#b")
			So(string(secretBytes), ShouldEqual, "xyz")
		})

		Convey("errors on not found", func() {
			_, err  := RetrieveSecret("production/no-secret")
			So(err.Error(), ShouldStartWith, secretsmanager.ErrCodeResourceNotFoundException)
		})
	})
}

func Test_getValueByKey(t *testing.T) {
	Convey("Given a valid JSON format stored secret", t, func() {
		secret := `{ "username": "USERNAME", "password": "PASSWORD", "port": 8000 }`

		Convey("Returns the value of the key", func() {
			stdout, err := getValueByKey("username", []byte(secret))
			So(err, ShouldBeNil)
			So(string(stdout), ShouldEqual, "USERNAME")
		})

		Convey("Always returns as a string", func() {
			stdout, err := getValueByKey("port", []byte(secret))
			So(err, ShouldBeNil)
			So(string(stdout), ShouldEqual, "8000")
		})
	})
}
