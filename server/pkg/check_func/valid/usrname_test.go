package valid

import "testing"

func TestCheckUsrname(t *testing.T) {

	tests := []struct {
		name   string
		strIn  string
		lenOut int
		want   bool
	}{
		{
			name:   "EnglishStr",
			strIn:  "User",
			lenOut: 4,
			want:   true,
		},
		{
			name:   "RussianStr",
			strIn:  "Юзер",
			lenOut: 4,
			want:   true,
		},
		{
			name:   "LenStr",
			strIn:  "Юз",
			lenOut: 2,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckUsrname(tt.strIn); got != tt.want {
				t.Errorf("CheckUsrname() = %v, want %v", got, tt.want)
			}
		})
	}
}
