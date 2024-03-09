#/bin/bash 

it=$(./bin/itmat1)
rec=$(./bin/recmat1)

if [ "$it" != "$rec" ]; then
    >&2 echo "Cpp matrix multipliers inconsistent"
    exit 1
fi