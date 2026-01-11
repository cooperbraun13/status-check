package main

import (
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://charm.sh/"

type model struct {
	status int
	err    error
}

func checkServer() tea.Msg {

	// create an HTTP client and make a GET request
	c := &http.Client{Timeout: 10 * time.Second}
	res, err := c.Get(url)

	if err != nil {
		// there was an error making our request. wrap the error we received
		// in a message and return it
		return errMsg{err}
	}
	// we received a response from the server. return the HTTP status code
	// as a message
	return statusMsg(res.StatusCode)
}

type statusMsg int

type errMsg struct{ err error }

// for messages that contain errors it's often handy to also implement the
// error interface on the message
func (e errMsg) Error() string {
	return e.err.Error()
}

func (m model) Init() tea.Cmd {
	return checkServer
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case statusMsg:
		// the server returned a status message. save it to our model. also
		// tell the Bubble Tea runtime we want to exit because we have nothing
		// else to do. we'll still be able to render a final view with our
		// status message
		m.status = int(msg)
		return m, tea.Quit

	case errMsg:
		// there was an error. note it in the model. and tell the runtime
		// we're done and want to quit
		m.err = msg
		return m, tea.Quit

	case tea.KeyMsg:
		// Ctrl+c exists. even with short running programs it's good to have
		// a quit key, just in case your logic is off. users will be very
		// annoyed if they can't exit
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}
	}

	// if we happen to get any other messages, don't do anything
	return m, nil
}

func (m model) View() string {

}

func main() {

}
