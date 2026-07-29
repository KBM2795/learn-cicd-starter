package auth

import (
    "reflect"
    "testing"
	"net/http"
)

func TestAuth(t *testing.T){
	headers := http.Header{
		"Authorization": []string{"ApiKey mykey"},
	}
	got,err := GetAPIKey(headers)
	want := "mykey"

    if err != nil{
		t.Errorf("got error %v, want %v", err, nil)
    }

    if !reflect.DeepEqual(got, want){
        t.Errorf("got %v, want %v", got, want)
    }

}

