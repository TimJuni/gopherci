	"bytes"
	"sync"
const webhookSecret = "ede9aa6b6e04fafd53f7460fb75644302e249177"

	var (
		wg sync.WaitGroup
		c  = make(chan interface{})
	)
	queue := queue.NewMemoryQueue(context.Background(), &wg, c)
	g, err := New(&mockAnalyser{}, memDB, queue, 1, integrationKey, webhookSecret)
func TestWebhookHandler(t *testing.T) {
	tests := []struct {
		signature  string
		event      string
		expectCode int
	}{
		{"sha1=d1e100e3f17e8399b73137382896ff1536c59457", "goci-invalid", http.StatusBadRequest},
		{"sha1=d1e100e3f17e8399b73137382896ff1536c59457", "push", http.StatusOK},
		{"sha1=aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "push", http.StatusBadRequest},
	}

	for _, test := range tests {
		g, _ := setup(t)
		body := bytes.NewBufferString(`{"key":"value"}`)
		r, err := http.NewRequest("POST", "https://example.com", body)
		if err != nil {
			t.Fatal(err)
		}
		r.Header.Add("X-GitHub-Event", test.event)
		r.Header.Add("X-Hub-Signature", test.signature)
		w := httptest.NewRecorder()
		g.WebHookHandler(w, r)

		if w.Code != test.expectCode {
			t.Fatalf("have code: %v, want: %v, test: %+v", w.Code, test.expectCode, test)
		}
	}
}

		accountID      = 3
		senderID       = 4
			ID:    github.Int(senderID),
	want := &db.GHInstallation{
		InstallationID: installationID,
		AccountID:      accountID,
		SenderID:       senderID,
	}

	have, _ := memDB.GetGHInstallation(installationID)
	if !reflect.DeepEqual(have, want) {
		t.Errorf("\nhave: %#v\nwant: %#v", have, want)
	have, _ = memDB.GetGHInstallation(installationID)
	if have != nil {
		t.Errorf("got: %#v, expected nil", have)
		expectedOwner   = "owner"
		expectedRepo    = "repo"
		expectedPR      = 3
		case fmt.Sprintf("/repos/%v/%v/pulls/%v/comments", expectedOwner, expectedRepo, expectedPR):
		accountID      = 3
		senderID       = 4
	_ = memDB.AddGHInstallation(installationID, accountID, senderID)
	memDB.EnableGHInstallation(installationID)
		Number: github.Int(expectedPR),
			HTMLURL:     github.String("https://github.com.au/owner/repo/pulls/3"),
					Name:     github.String(expectedRepo),
					Owner: &github.User{
						Login: github.String(expectedOwner),
					},
		},
		Installation: &github.Installation{
			ID: github.Int(installationID),

func TestPullRequestEvent_noInstall(t *testing.T) {
	g, _ := setup(t)

	const installationID = 2
	event := &github.PullRequestEvent{
		Action: github.String("opened"),
		Number: github.Int(1),
		Installation: &github.Installation{
			ID: github.Int(installationID),
		},
	}

	err := g.PullRequestEvent(event)
	if want := errors.New("could not find installation with ID 2"); err.Error() != want.Error() {
		t.Errorf("expected error %q have %q", want, err)
	}
}

func TestPullRequestEvent_disabled(t *testing.T) {
	g, memDB := setup(t)

	const installationID = 2

	// Added but not enabled
	_ = memDB.AddGHInstallation(installationID, 3, 4)

	event := &github.PullRequestEvent{
		Action: github.String("opened"),
		Number: github.Int(1),
		Installation: &github.Installation{
			ID: github.Int(installationID),
		},
	}

	err := g.PullRequestEvent(event)
	if want := errors.New("could not find installation with ID 2"); err.Error() != want.Error() {
		t.Errorf("expected error %q have %q", want, err)
	}
}