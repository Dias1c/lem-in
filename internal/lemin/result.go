package lemin

import (
	"bufio"

	"github.com/Dias1c/lem-in/internal/lemin/anthill"
)

// GetResult - returns result, nil if shortest disjoint paths was found
func GetResult(scanner *bufio.Scanner) (*anthill.Result, error) {
	terrain := anthill.CreateAnthill()
	var err error
	for scanner.Scan() {
		err = terrain.ReadDataFromLine(scanner.Text())
		if err != nil {
			return nil, errInvalidDataFormat(err)
		}
	}
	err = terrain.ValidateByFieldInfo()
	if err != nil {
		return nil, errInvalidDataFormat(err)
	}
	err = terrain.Match()
	if err != nil {
		return nil, errPaths(err)
	}
	return terrain.Result, nil
}
