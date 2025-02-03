package a

import (
	"os"

	"github.com/beati/stacked"
)

func callExternalPackage() error {
	var name string

	name, err := os.Hostname() // want "err is not wrapped with stacked"
	if err != nil {
		return err
	}

	name, err = os.Hostname()
	err = stacked.Wrap(err)
	if err != nil {
		return err
	}

	if err := os.Chmod(name, 0777); err != nil { // want "err is not wrapped with stacked"
		return err
	}

	err = os.Chmod(name, 0777) // want "err is not wrapped with stacked"
	name = "test"
	err = stacked.Wrap(err)
	if err != nil {
		return err
	}

	err = stacked.Wrap(os.Chmod("test", 0777))
	if err != nil {
		return err
	}

	err = os.Chmod("test", 0777)
	err = stacked.Wrap(err)
	if err != nil {
		return err
	}

	err = os.Chmod("test", 0777) // want "err is not wrapped with stacked"
	if err != nil {
		return stacked.Wrap(err)
	}

	f, err := os.Open(name) // want "err is not wrapped with stacked"
	if err != nil {
		return err
	}

	err = f.Close() // want "err is not wrapped with stacked"
	if err != nil {
		return err
	}

	err = stacked.Wrap(f.Close())
	if err != nil {
		return err
	}

	err = f.Close()
	err = stacked.Wrap(err)
	if err != nil {
		return err
	}

	return nil
}
