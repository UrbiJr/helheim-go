# helheim-go

## Prerequisites
This go library was created with
* go 1.15 
* helheim 0.8.6 py38

## Warning
`char *helheimVersion();` is not in the default cffi lib you can build out of the examples.
Please check helheim discord server how to add this function and then be able to call `Version()` on the helheimClient.

## Install Python

### MacOS
Install Python in a version which is compatible with helheim.
I used this: https://www.python.org/ftp/python/3.8.5/python-3.8.5-macosx10.9.pkg

## Windows
TBD (Did not test that on windows)

## Linux

### Ubuntu
Install Python in a version which is compatible with helheim.
I used this:
```
apt-get update 
apt-get upgrade

apt-get install python3 python3-dev python3-pip

# Download Go 1.15 and install it to /usr/local/go
curl -s https://storage.googleapis.com/golang/go1.15.linux-amd64.tar.gz| tar -v -C /usr/local -xz

python3 -m pip install -U --force-reinstall pip
pip install -U helheim-0.8.2-py38-linux.x86_64.tar.gz
```

## Build
Build the cffi library. It is very important that you download from discord the tar.gz file with the **correct** version and for your system you are building your go app.
In my case with the above installed python version on MacOS it is `helheim-0.8.2-py38-darwin.x86_64.tar.gz`
On Ubuntu (amd64) it is `helheim-0.8.2-py38-linux.x86_64.tar.gz`.

Install helheim by running `python setup.py install` in the projects root directory or just `pip install -U helheim-0.8.2-py38-linux.x86_64.tar.gz` (with the correct package name)

Now build the cffi library as described in python example 3. Navigate inside the directory and run:  
`python build-cffi.py`. In my case `python3 build-cffi.py`

This will create the following files for you:
* `helheim_cffi.c`
* `helheim_cffi.o`
* `helheim_cffi.dylib` (MacOS)
* `helheim_cffi.so` (Linux)
* `helheim_cffi.dll` (Windows)

## Copy cffi library
Copy the `helheim_cffi` library (`dylib, so or dll`) into your pythons lib directory.
In my case it is: `/Library/Frameworks/Python.framework/Versions/3.8/lib`.
Under Ubuntu it is: `/usr/include/python3.8`

**Note:** On Linux and MacOS create a duplicate of the cffi library file and name it `libhelheim_cffi.dylib` (MacOS) or `libhelheim_cffi.so` (Linux). I have to keep **both** files in the directory to get the example running.

In addition copy the `helheim_cffi.dylib` file in your working directory of your running app or use the following env variable to start your application:
```
# Specify where application finds the cffi lib file on runtime
# MacOS
# DYLD_LIBRARY_PATH="/Library/Frameworks/Python.framework/Versions/3.8/lib" ./yourCompiledAppBinary
# Linux
# LD_LIBRARY_PATH="/usr/include/python3.8" ./yourCompiledAppBinary

# We assume here that you copied the cffi lib file into the applications working directory and we do not need to define the path for the dyld
# ./yourCompiledAppBinary
```

## How to build golang app
Due to the fact that we are loading a C dynamic library we need to build our app with cgo.
That is done by adding a few flags and env vars to our `go build` command.
Check `./example/build.sh` for the correct `go build` command. This script will build the example application located in `./example/main.go`

**You have to adjust the different paths according to your setup/system in `build.sh`**

## Quick Usage Example
```go
    package main
    
    import(
	helheim_go "github.com/bogdanfinn/helheim-go"
        "log"
        "net/http"
    )

    func main() {
        // NewClient() returns each time a new instance which is requesting helheims auth() endpoint
        // helheimClient, err := helheim_go.NewClient("YOUR_API_KEY", false, nil)
        
        // Provide() is creating one helheim client instance and returning the same authenticated instance on every Provide() call
        helheimClient, err := helheim_go.ProvideClient("YOUR_API_KEY", false, true, nil)
        
        if err != nil {
            log.Println(err)
            return
        }
        
        log.Println("helheim client initiated")
        
        // check possible options in the python examples
        options := helheim_go.CreateSessionOptions{
            Browser: helheim_go.BrowserOptions{
                Browser:  "chrome",
                Mobile:   false,
                Platform: "windows",
            },
            Captcha: helheim_go.CaptchaOptions{
                Provider: "vanaheim",
            },
        }
        
        session, err := helheimClient.NewSession(options)
        if err != nil {
            log.Println(err)
            return
        }
        
        log.Println("session initiated")
        
        reqOpts := helheim_go.RequestOptions{
            Method:  http.MethodGet,
            Url:     "https://www.genx.co.nz/iuam/",
            Options: make(map[string]string),
        }
        
        resp, err := session.Request(reqOpts)
        if err != nil {
            log.Println(err)
            return
        }
    
        log.Println("response:")
        log.Println(resp)
    }
```
For more examples check `./example/main.go`

### C Types in Go cheat sheet
https://gist.github.com/zchee/b9c99695463d8902cd33