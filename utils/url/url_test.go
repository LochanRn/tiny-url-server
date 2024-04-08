package url

import (
	"testing"

	. "github.com/onsi/gomega"
)

func Test_GetDomainFromURL(t *testing.T) {
	g := NewGomegaWithT(t)
	type args struct {
		urlStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Valid url 1",
			args: args{
				urlStr: "https://www.google.com/search?q=golang",
			},
			want: "https://www.google.com/",
		},
		{
			name: "Valid url 2",
			args: args{
				urlStr: "https://www.google.com/search/v1/list/noddes/randomstringforevernotsurehow?q=golang",
			},
			want: "https://www.google.com/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetDomainFromURL(tt.args.urlStr)
			g.Expect(got).To(Equal(tt.want))
		})
	}
}

func Test_GetPathFromURL(t *testing.T) {
	g := NewGomegaWithT(t)
	type args struct {
		urlStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Valid url 1",
			args: args{
				urlStr: "https://www.google.com/search",
			},
			want: "/search",
		},
		{
			name: "Valid url 2",
			args: args{
				urlStr: "https://www.google.com/search/v1/list/noddes/randomstringforevernotsurehow",
			},
			want: "/search/v1/list/noddes/randomstringforevernotsurehow",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetPathFromURL(tt.args.urlStr)
			g.Expect(got).To(Equal(tt.want))
		})
	}
}
