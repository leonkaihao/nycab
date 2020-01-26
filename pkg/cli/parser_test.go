package cli

import "testing"

import "strings"

import "fmt"

type testHandler struct {
	argsNum int
}

func (t *testHandler) Handle(args []string) (resp string, err error) {
	if len(args) != t.argsNum {
		return "", fmt.Errorf("expect %v args, got %v", t.argsNum, len(args))
	}
	return strings.Join(args, " "), nil
}

func Test_parser_Parse(t *testing.T) {
	type fields struct {
		cmds []*Cmdline
	}
	type args struct {
		cmd string
	}
	cmdFields := fields{
		cmds: []*Cmdline{
			&Cmdline{"cmd1", "", "", &testHandler{1}},
			&Cmdline{"cmd2", "", "", &testHandler{2}},
			&Cmdline{"cmd3", "", "", &testHandler{3}},
			&Cmdline{"cmd4", "", "", &testHandler{4}},
		},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// positive test
		{name: "normal1", fields: cmdFields, args: args{"cmd1 arg1=v"}, wantErr: false},
		{name: "normal2", fields: cmdFields, args: args{"cmd2 arg1=v arg2=v"}, wantErr: false},
		{name: "normal3", fields: cmdFields, args: args{"cmd3 arg1=v arg2=v arg3=v"}, wantErr: false},
		{name: "normal4", fields: cmdFields, args: args{"cmd4 arg1=v arg2=v arg3=v arg4=v"}, wantErr: false},
		{name: "normal5", fields: cmdFields, args: args{"cmd4 arg1=v arg2=v arg3=v arg4=v    "}, wantErr: false},
		// neg
		{name: "null command", fields: cmdFields, args: args{""}, wantErr: true},
		{name: "unknown command", fields: cmdFields, args: args{"cmd5"}, wantErr: true},
		{name: "argnum1", fields: cmdFields, args: args{"cmd4"}, wantErr: true},
		{name: "argnum2", fields: cmdFields, args: args{"cmd4 arg1=v"}, wantErr: true},
		{name: "argnum3", fields: cmdFields, args: args{"cmd4 arg1=v arg2=v"}, wantErr: true},
		{name: "argnum4", fields: cmdFields, args: args{"cmd4 arg1=v arg2=v arg3=v"}, wantErr: true},
		{name: "argnum5", fields: cmdFields, args: args{"cmd4 arg1=v arg2=v arg3=v arg4=v arg5=v"}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &parser{
				cmds: tt.fields.cmds,
			}
			if err := p.Parse(tt.args.cmd); (err != nil) != tt.wantErr {
				t.Errorf("parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
