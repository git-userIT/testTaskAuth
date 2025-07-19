package valid

import "testing"

func TestCheckPassLen(t *testing.T) {

	tests := []struct {
		name string
		pass string
		want bool
	}{
		{
			name: "LenDownSix",
			pass: "Pass",
			want: false,
		},
		{
			name: "LenUpSix",
			pass: "Passwo",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPassLen(tt.pass); got != tt.want {
				t.Errorf("CheckPassLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
