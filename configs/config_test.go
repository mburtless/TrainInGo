package configs_test

import(
	"testing"
	"os"
	"os/exec"
	"github.com/mburtless/trainingo/configs"
)

func TestInitSvcCode(t *testing.T) {
	svcCode := configs.InitSvcCode()
	valid := map[string]bool{"SUN": true, "SAT": true, "WKD": true}
	if valid[svcCode] {
		return
	} else{
		t.Errorf("InitSvcCode() returned %q but expecting either SUN, SAT or WKD", svcCode)
	}
}

func TestInitCredentials(t *testing.T) {
	var value string
	key := "TEST"

	//If test key is unset, set it for testing and defer unsettings it
	if value = os.Getenv(key); value == "" {
		os.Setenv(key, "TESTING")
		defer func() { os.Unsetenv(key) }()
	}

	returnVal := configs.InitCredentials(key)

	if returnVal != "TESTING" {
		t.Errorf("InitCredentials() didn't return appropriate test val from env %q", key)
	}

	return
}

func TestInitCredentialsErr(t *testing.T) {
	var value string
	key := "TEST"

	//If test key is set, unset it for testing and defer setting it back to orig val
	if value = os.Getenv(key); value != "" {
		os.Unsetenv(key)
		defer func() { os.Setenv(key, value)  }()
	}

	if os.Getenv("TEST_INIT_CREDENTIALS") == "1" {
		configs.InitCredentials(key)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestInitCredentials")
	cmd.Env = append(os.Environ(), "TEST_INIT_CREDENTIALS=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	// If didn't hit return above, we must not have hit fatal on unset env var
	t.Errorf("InitCredentials() failed to fatal if env %q is not present", key)
}

func TestInitLineFeeds(t *testing.T) {
	lineFeeds := configs.InitLineFeeds("TestKey")
	if len(lineFeeds) < 7 {
		t.Errorf("InitLineFeeds failed to return a map of expected length 7")
	}

	return
}
