package fonts

import (
	"github.com/pkg/errors"
	re "golang.org/x/sys/windows/registry"
	"strings"
)

func deleteValues(keyName string, values []string) error {

	k, err := re.OpenKey(re.LOCAL_MACHINE, keyName, re.ALL_ACCESS)
	if err != nil {
		return errors.Wrapf(err, "cannot open key %s", keyName)
	}
	defer k.Close()

	for _, name := range values {

		_, _, err := k.GetStringValue(name)
		if err == re.ErrNotExist {
			continue
		}
		if err != nil {
			return errors.Wrapf(err, "cannot read value '%s'", name)
		}

		err = k.DeleteValue(name)
		if err != nil {
			return errors.Wrapf(err, "cannot delete value '%s'", name)
		}
	}

	return nil
}

func Delete(values []string) error {
	err := deleteValues(regKeyWin, values)
	if err != nil {
		return err
	}
	return deleteValues(regKeyWow, values)
}

func DeleteWin81() error {

	parts := strings.Split(FontNamesWindows81, "\n")
	names := make([]string, 0, len(parts))
	for _, val := range parts {
		names = append(names, strings.TrimSpace(val))
	}

	return Delete(names)
}
