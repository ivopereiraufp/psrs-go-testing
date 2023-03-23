package divide

import "testing"

func TestDivideBy3(t *testing.T) {
	result := DivideBy3(6)
	if result != "DivBy3" {
		t.Errorf(`Result is incorrect, we got %v, 
							but we wanted: "DivBy3"`, result)
	}
}

func BenchmarkDivideBy3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DivideBy3(i)
	}
}

func FuzzDivideBy3(f *testing.F) {
	f.Add(3)
	f.Fuzz(func(t *testing.T, a int) {
		DivideBy3(a)
	})
}

func TestDivideBy3Skipped(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping this test in short mode")
	}
	result := DivideBy3(3)
	if result != "DivBy3" {
		t.Errorf(`Result is incorrect, we got %v, 
							but we wanted: "DivBy3"`, result)
	}
}

func TestDivideBy3TableDriven(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		want  string
	}{
		{"9 should be DivBy3", 9, "DivBy3"},
		{"3 should be DivBy3", 3, "DivBy3"},
		{"1 is not DivBy3", 1, "1"},
		{"0 should by DivBy3", 0, "DivBy3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := DivideBy3(tt.input)
			if res != tt.want {
				t.Errorf("DivideBy3() = %v, want %v", res, tt.want)
			}
		})
	}
}

func TestDivideBy3TableDriven2(t *testing.T) {
	type args struct {
		input int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"9 should be DivBy3", args{9}, "DivBy3"},
		{"3 should be DivBy3", args{3}, "DivBy3"},
		{"1 is not DivBy3", args{1}, "1"},
		{"0 should by DivBy3", args{0}, "DivBy3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DivideBy3(tt.args.input); got != tt.want {
				t.Errorf("DivideBy3() = %v, want %v", got, tt.want)
			}
		})
	}
}
