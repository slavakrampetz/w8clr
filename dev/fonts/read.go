package fonts

import (
	"errors"
	re "golang.org/x/sys/windows/registry"
	"sort"
)

func readValues(keyName string) ([]string, error) {

	k, err := re.OpenKey(re.LOCAL_MACHINE, keyName, re.QUERY_VALUE)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	names, err := k.ReadValueNames(-1)
	if err != nil {
		return nil, errors.New("reading value names of " + keyName + " failed: " + err.Error())
	}
	return names, nil
}

func Read() ([]string, error) {

	names1, err := readValues(regKeyWin)
	if err != nil {
		return nil, err
	}
	names2, err := readValues(regKeyWow)
	if err != nil {
		return nil, err
	}

	all := append(names1, names2...)
	sort.Strings(all)

	empty := struct{}{}
	names := make([]string,0,len(names1))
	m := make(map[string]struct{})
	for _, name := range all {
		if _, ok := m[name]; !ok {
			m[name] = empty
			names = append(names, name)
		}
	}

	return names, nil
}
