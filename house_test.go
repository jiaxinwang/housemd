package house

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestIPAddresses(t *testing.T) {
	validate := validator.New()
	tests := []struct {
		name    string
		want    IPInfo
		wantErr bool
	}{
		{"case-1", IPInfo{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IPAddresses()
			if (err != nil) != tt.wantErr {
				t.Errorf("IPAddresses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			errs := validate.Var(got.IP, "required,ip")
			assert.NoError(t, errs)
		})
	}
}

// func TestGFW(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		wantErr bool
// 	}{
// 		{"case-1", true},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if err := GFW(); (err != nil) != tt.wantErr {
// 				t.Errorf("GFW() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
