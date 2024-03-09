#/bin/bash 

make
$(./bin/itmat1) | diff -y <($(./bin/recmat1)) -
ret0=$?
if [ $ret0 -ne 0 ]; then
    >&2 echo "Go matrix multipliers inconsistent"
    exit 1
fi