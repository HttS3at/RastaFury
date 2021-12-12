/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"RastaShellGenerator/donut"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var moduleName string
var url string
var entropy int
var arch string
var bypass int
var dstFile string
var format int
var oepString string
var action int
var classname string
var appdomain string
var method string
var params string
var wflag bool
var runtime string
var tflag bool
var zflag int
var srcfile string
var verbose bool
var err error

// ShellCodeGeneratorCmd represents the ShellCodeGenerator command
var ShellCodeGeneratorCmd = &cobra.Command{
	Use:   "shellcodegenerator",
	Short: "Disect a PE file and Generate a ShellCode from it",
	Long:  `Convert a VBS/JS or PE/.NET EXE/DLL to shellcode`,

	Run: func(cmd *cobra.Command, args []string) {
		if srcfile == "" {
			log.Println(err)
			return
		}
		oep := uint64(0)
		if oepString != "" {
			oep, err = strconv.ParseUint(oepString, 16, 64)
			if err != nil {
				log.Println("Invalid OEP: " + err.Error())
				return
			}
		}

		var donutArch donut.DonutArch
		switch strings.ToLower(arch) {
		case "x32", "386":
			donutArch = donut.X32
		case "x64", "amd64":
			donutArch = donut.X64
		case "x84":
			donutArch = donut.X84
		default:
			log.Fatal("Unknown architecture provided")
		}

		config := new(donut.DonutConfig)
		config.Arch = donutArch
		config.Entropy = uint32(entropy)
		config.OEP = oep

		if url == "" {
			config.InstType = donut.DONUT_INSTANCE_PIC
		} else {
			config.InstType = donut.DONUT_INSTANCE_URL
		}

		config.Parameters = params
		config.Runtime = runtime
		config.URL = url
		config.Class = classname
		config.Method = method
		config.Domain = appdomain
		config.Bypass = bypass
		config.ModuleName = moduleName
		config.Compress = uint32(zflag)
		config.Format = uint32(format)
		config.Verbose = verbose

		if tflag {
			config.Thread = 1
		}
		if wflag { // convert command line to unicode? only applies to unmanaged DLL function
			config.Unicode = 1
		}
		config.ExitOpt = uint32(action)

		if srcfile == "" {
			if url == "" {
				log.Fatal("No source URL or file provided")
			}
			payload, err := donut.ShellcodeFromURL(url, config)
			if err == nil {
				err = ioutil.WriteFile(dstFile, payload.Bytes(), 0644)
			}
		} else {
			payload, err := donut.ShellcodeFromFile(srcfile, config)
			//fmt.Println(base64.StdEncoding.EncodeToString([]byte(payload.String())))
			if err == nil {
				f, err := os.Create(dstFile)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()
				if _, err = payload.WriteTo(f); err != nil {
					log.Fatal(err)
				}
			}
		}

		if err != nil {
			log.Println(err)
		} else {
			log.Println("Done!")
		}
	},
}

func init() {
	rootCmd.AddCommand(ShellCodeGeneratorCmd)

	//Flags
	// -MODULE OPTIONS-
	ShellCodeGeneratorCmd.Flags().StringVarP(&moduleName, "module", "n", "", "Module name. Randomly generated by default with entropy enabled.")
	ShellCodeGeneratorCmd.Flags().StringVarP(&url, "url", "u", "", "HTTP server that will host the donut module.")
	ShellCodeGeneratorCmd.Flags().IntVarP(&entropy, "entropy", "e", 3, "Entropy. 1=disable, 2=use random names, 3=random names + symmetric encryption (default)")

	//  -PIC/SHELLCODE OPTIONS-
	ShellCodeGeneratorCmd.Flags().StringVarP(&arch, "arch", "a", "x84", "Target Architecture: x32, x64, or x84")
	ShellCodeGeneratorCmd.Flags().IntVarP(&bypass, "bypass", "b", 3, "Bypass AMSI/WLDP : 1=skip, 2=abort on fail, 3=continue on fail.")
	ShellCodeGeneratorCmd.Flags().StringVarP(&dstFile, "out", "o", "defaultc2client.bin", "Output file.")
	ShellCodeGeneratorCmd.Flags().IntVarP(&format, "format", "f", 1, "Output format. 1=raw, 2=base64, 3=c, 4=ruby, 5=python, 6=powershell, 7=C#, 8=hex")
	ShellCodeGeneratorCmd.Flags().StringVarP(&oepString, "oep", "y", "", "Create a new thread for loader. Optionally execute original entrypoint of host process.")
	ShellCodeGeneratorCmd.Flags().IntVarP(&action, "exit", "x", 1, "Exiting. 1=exit thread, 2=exit process")

	//  -FILE OPTIONS-
	ShellCodeGeneratorCmd.Flags().StringVarP(&classname, "class", "c", "", "Optional class name.  (required for .NET DLL)")
	ShellCodeGeneratorCmd.Flags().StringVarP(&appdomain, "domain", "d", "", "AppDomain name to create for .NET.  Randomly generated by default with entropy enabled.")
	ShellCodeGeneratorCmd.Flags().StringVarP(&method, "method", "m", "", "Optional method or API name for DLL. (a method is required for .NET DLL)")
	ShellCodeGeneratorCmd.Flags().StringVarP(&params, "params", "p", "", "Optional parameters/command line inside quotations for DLL method/function or EXE.")
	ShellCodeGeneratorCmd.Flags().BoolVarP(&wflag, "unicode", "w", false, "Command line is passed to unmanaged DLL function in UNICODE format. (default is ANSI)")
	ShellCodeGeneratorCmd.Flags().StringVarP(&runtime, "runtime", "r", "", "CLR runtime version. This will override the auto-detected version.")
	ShellCodeGeneratorCmd.Flags().BoolVarP(&tflag, "thread", "t", false, "Create new thread for entrypoint of unmanaged EXE.")
	ShellCodeGeneratorCmd.Flags().IntVarP(&zflag, "compress", "z", 1, "Pack/Compress file. 1=disable, 2=LZNT1, 3=Xpress, 4=Xpress Huffman")

	// go-donut only flags
	ShellCodeGeneratorCmd.Flags().StringVarP(&srcfile, "in", "i", "", ".NET assembly, EXE, DLL, VBS, JS or XSL file to execute in-memory.")
	ShellCodeGeneratorCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Show verbose output.")
	//No Optional Flags
	_ = ShellCodeGeneratorCmd.MarkFlagRequired("in")
}
