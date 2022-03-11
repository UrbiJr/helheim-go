#!/bin/zsh
# You have to adjust the different paths according to your setup/system. This is for working locally on macos
export PY_DIR='/Library/Frameworks/Python.framework/Versions/3.7'
export PY_VERSION='3.7'
CGO_LDFLAGS='-Wl,-rpath /Library/Frameworks/Python.framework/Versions/3.7 -L/Library/Frameworks/Python.framework/Versions/3.7/lib -lpython3.7 -lhelheim_cffi' CGO_CFLAGS="-I/Library/Frameworks/Python.framework/Versions/3.7/include/python3.7" go build main.go

# Specify where application finds the cffi lib file on runtime
# DYLD_LIBRARY_PATH="/Library/Frameworks/Python.framework/Versions/3.7/lib" ./main

# We assume here that you copied the cffi lib file into the applications working directory and we do not need to define the path for the dyld
# ./main


# You have to adjust the different paths according to your setup/system. This is for working locally on ubuntu
# CGO_LDFLAGS='-Wl,-rpath /usr/include/python3.7 -L/usr/include/python3.7 -lpython3.7 -lhelheim_cffi' CGO_CFLAGS="I/usr/include/python3.7" go build main.go

# Specify where application finds the cffi lib file on runtime
# LD_LIBRARY_PATH="/usr/include/python3.7" ./main

# We assume here that you copied the cffi lib file into the applications working directory and we do not need to define the path for the dyld
# ./main
