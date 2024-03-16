#!/bin/sh

cd cpp
cmake .
make

for program in "recmat1" "itmat1"; do  
    ./bin/${program}
    pprof --pdf prof/${program}.cprof > ${program}.pdf
done

cd ..
cd go
make 

for program in "recmat1" "itmat1"; do 
    ./bin/${program} -cpuprofile=prof/${program}.cprof
    go tool pprof --pdf prof/${program}.cprof > ${program}.pdf
done 
