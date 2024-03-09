#/bin/bash 

it=$(./bin/itmat1)
rec=$(./bin/recmat1)

if [ "$it" != "$rec" ]; then
    >&2 echo "Go matrix multipliers inconsistent"
    exit 1
fi

gc1=$(./bin/gcgo1)
gc2=$(./bin/gcgo2)
if [ "$gc1" != "$gc2" ]; then 
    >&2 echo "Go gc inconsistent"
    exit 1
fi