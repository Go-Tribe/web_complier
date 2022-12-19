package docker

import (
	"testing"
)

func TestDockerRun(t *testing.T) {
	type args struct {
		lang        string
		code        string
		dir         string
		ext         string
		cmd         string
		langTimeout int64
		memory      int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"python3",
			args{lang: "python3", dir: "/tmp", ext: "py", cmd: "python3 filename.py", langTimeout: 5, memory: 104857600, code: "def print_welcome(name):\n    print(\"Welcome\", name)\n \nprint_welcome(\"gotribe\")"},
			"Welcome gotribe\n",
		},
		{
			"rust",
			args{lang: "rust", dir: "/tmp", ext: "rs", cmd: "rustc filename.rs -o filename\nif test -f \".filename\"; then\n.filename\nfi", langTimeout: 5, memory: 104857600, code: "fn main() {\n    println!(\"Hello, gotribe!\");}"},
			"Hello, gotribe!\n",
		},
		{
			"python3-2",
			args{lang: "python3", dir: "/tmp", ext: "py", cmd: "python3 filename.py", langTimeout: 1, memory: 104857600, code: "while True:\n   print(\"111\")\n\n\n# fn main() {\n#     println!(\"Hello, world!\");\n# }"},
			"execute timeout",
		},
		{
			"golang",
			args{lang: "golang", dir: "/go", ext: "go", cmd: "go run filename.go", langTimeout: 5, memory: 104857600, code: "package main\n\nimport (\n\t\"fmt\"\n\t\"time\"\n)\n\nfunc main() {\n  \n\tfor i := 1; i < 2; i++ {\n\t\tfmt.Print(i)\n\t\ttime.Sleep(time.Duration(0*time.Second))\n\t}\n}"},
			"1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DockerRun(tt.args.lang, tt.args.code, tt.args.dir, tt.args.cmd, tt.args.langTimeout, tt.args.memory, tt.args.ext); got != tt.want {
				t.Errorf("DockerRun() = %v, want %v", got, tt.want)
			}
		})
	}
}
