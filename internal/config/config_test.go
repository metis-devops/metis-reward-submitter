package config

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestParse(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name    string
		args    args
		envs    map[string]string
		want    *Config
		wantErr bool
	}{
		{
			name: "json",
			args: args{"testdata/config.json"},
			envs: map[string]string{"METIS_RPC": "https://test.metis.io", "THEMIS_REST": "https://rest.metis.io"},
			want: &Config{
				Eth: &Eth{
					RPC:          "http://eth.example.com",
					LockingPool:  common.HexToAddress("0x293E8b379FD574ec5f64013307359ff2863C1bEC"),
					LockcingInfo: common.HexToAddress("0x9F8f709002907cd1D734d16632a18F5E5735d8DC"),
				},
				Metis: &Metis{
					RPC:    "https://test.metis.io",
					Themis: "https://rest.metis.io",
					SeqSet: common.HexToAddress("0x2F5F158996dab39cfD35f580e6f4547FDF159D1d"),
				},
			},
		},
		{
			name: "yaml",
			args: args{"testdata/config.yaml"},
			envs: map[string]string{"ETH_RPC": "https://test.ethereum.org"},
			want: &Config{
				Eth: &Eth{
					RPC:          "https://test.ethereum.org",
					LockingPool:  common.HexToAddress("0x8308918Ff3D37B94d82c83ec8f58fFe7bCA2b904"),
					LockcingInfo: common.HexToAddress("0x5f1bd39614d18017b9f4ac3B2bFa34cA4BaCd08A"),
				},
				Metis: &Metis{
					RPC:    "http://example-metis.com",
					Themis: "http://example-themis.com",
					SeqSet: common.HexToAddress("0x06C6c78E1d5b6278C54F1AA0A9C1a524B4143ef6"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.envs {
				t.Setenv(key, value)
			}

			got, err := Parse(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
