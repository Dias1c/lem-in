package lemin

import (
	"bufio"
	"lem-in/lemin/anthill"
)

// getResult - returns result, nil if shortest disjoint paths was found
func getResult(scanner *bufio.Scanner) (*anthill.Result, error) {
	terrain := anthill.CreateAnthill()
	for scanner.Scan() {
		err := terrain.ReadDataFromLine(scanner.Text())
		if err != nil {
			return nil, errInvalidDataFormat(err)
		}
	}
	err := terrain.Match()
	// fmt.Printf("%+v\n", terrain)
	// fmt.Printf("%+v\n", terrain.Result)
	if err != nil {
		return nil, errPaths(err)
	}
	return terrain.Result, nil
}
