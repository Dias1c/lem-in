#!/bin/bash
PathToFiles="examples/audit/*"
for file in $PathToFiles
do
    if [ -f "$file" ]
    then
        echo "Executing $file"
        time go run . "$file"
        echo ""
    else
        echo "Wrong $file"
    fi
done

# echo ""
# echo "#*_*#"
# echo ""

# wget "https://raw.githubusercontent.com/elijahkash/lemin/master/test_maps/mars_4000_20_20_95_180_5_no_z"
# time go run . "mars_4000_20_20_95_180_5_no_z"