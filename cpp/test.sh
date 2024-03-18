#/bin/bash 

it=$(./bin/itmat1)
rec=$(./bin/recmat1)

if [ "$it" != "$rec" ]; then
    >&2 echo "Cpp matrix multipliers inconsistent"
    exit 1
fi

quick=$(./bin/quicksort1)
merge=$(./bin/mergesort1)

if [ "$quick" != "$merge"]; then
    >&2 echo "Cpp sorting algorithms inconsistent"
    exit 1
fi