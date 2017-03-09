	"context"
	ExecuteErr []error
func (a *mockAnalyser) NewExecuter(_ context.Context, _ string) (Executer, error) {
func (a *mockAnalyser) Execute(_ context.Context, args []string) (out []byte, err error) {
	err, a.ExecuteErr = a.ExecuteErr[0], a.ExecuteErr[1:]
func (a *mockAnalyser) Stop(_ context.Context) error {
		{Name: "Name2", Path: "tool3"},
			[]byte("file is not generated"),              // isFileGenerated
			[]byte("file is not generated"),              // isFileGenerated
			[]byte("main.go:1: error3"),                  // tool 3 tested a generated file
			[]byte("file is generated"),                  // isFileGenerated
		},
		ExecuteErr: []error{
			nil, // git clone
			nil, // git fetch
			nil, // git diff
			nil, // install-deps.sh
			nil, // pwd
			nil, // tool 1
			&NonZeroError{ExitCode: 1}, // isFileGenerated - not generated
			nil, // tool 2 output abs paths
			&NonZeroError{ExitCode: 1}, // isFileGenerated - not generated
			nil, // tool 3 tested a generated file
			nil, // isFileGenerated - generated
	issues, err := Analyse(context.Background(), analyser, tools, cfg)
		{"isFileGenerated", "/go/src/gopherci", "main.go"},
		{"isFileGenerated", "/go/src/gopherci", "main.go"},
		{"tool3"},
		{"isFileGenerated", "/go/src/gopherci", "main.go"},
		{Name: "Name2", Path: "tool3"},
			[]byte("file is not generated"),              // isFileGenerated
			[]byte("file is not generated"),              // isFileGenerated
			[]byte("main.go:1: error3"),                  // tool 3 tested a generated file
			[]byte("file is generated"),                  // isFileGenerated
		},
		ExecuteErr: []error{
			nil, // git clone
			nil, // git checkout
			nil, // git diff
			nil, // install-deps.sh
			nil, // pwd
			nil, // tool 1
			&NonZeroError{ExitCode: 1}, // isFileGenerated - not generated
			nil, // tool 2 output abs paths
			&NonZeroError{ExitCode: 1}, // isFileGenerated - not generated
			nil, // tool 3 tested a generated file
			nil, // isFileGenerated - generated
	issues, err := Analyse(context.Background(), analyser, tools, cfg)
		{"isFileGenerated", "/go/src/gopherci", "main.go"},
		{"isFileGenerated", "/go/src/gopherci", "main.go"},
		{"tool3"},
		{"isFileGenerated", "/go/src/gopherci", "main.go"},
	_, err := Analyse(context.Background(), analyser, nil, cfg)