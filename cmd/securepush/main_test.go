package main

import "testing"

func Test_handleapiRequest(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handleapiRequest(); got != tt.want {
				t.Errorf("handleapiRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_executeShellScript(t *testing.T) {
	type args struct {
		username string
		password string
		reponame string
		tarfile  string
		provfile string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := executeShellScript(tt.args.username, tt.args.password, tt.args.reponame, tt.args.tarfile, tt.args.provfile); got != tt.want {
				t.Errorf("executeShellScript() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
