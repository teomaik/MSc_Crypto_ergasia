// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The following enables go generate to generate the doc.go file.
//go:generate go run v.io/x/lib/cmdline/gendoc . -help

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"v.io/v23/vdl"
	"v.io/v23/vdl/vdltest"
	"v.io/v23/vom"
	"v.io/x/lib/cmdline"
	"v.io/x/lib/textutil"
	"v.io/x/ref/lib/vdl/codegen"
	"v.io/x/ref/lib/vdl/codegen/vdlgen"
)

var cmdGen = &cmdline.Command{
	Runner: cmdline.RunnerFunc(runGen),
	Name:   "vomtestgen",
	Short:  "generates test data for the vomtest package",
	Long: `
Command vomtestgen generates test cases for the vomtest package.  The following
file is generated:

   pass81_gen.vdl - Golden file containing passing test cases for version 81.

Do not run this tool manually.  Instead invoke it via:
   $ go generate v.io/v23/vom/vomtest
`,
}

func main() {
	cmdline.Main(cmdGen)
}

func runGen(_ *cmdline.Env, _ []string) error {
	allCanonical := vdltest.AllPassFunc(func(e vdltest.Entry) bool {
		return e.IsCanonical
	})
	writePassFile("pass81", vdltest.ToEntryValues(allCanonical))
	return nil
}

// This tool is only used to generate test cases for the vdltest package, so the
// strategy is to panic on any error, to make the code simpler.
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func createFile(name string) (*os.File, func()) {
	file, err := os.Create(name)
	panicOnError(err)
	return file, func() { panicOnError(file.Close()) }
}

func writef(w io.Writer, format string, args ...interface{}) {
	_, err := fmt.Fprintf(w, format, args...)
	panicOnError(err)
}

func writeHeader(w io.Writer) {
	writef(w, `// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by v.io/v23/vom/vomtest/internal/vomtestgen
// Run the following to re-generate:
//   $ go generate v.io/v23/vom/vomtest

package vomtest

`)
}

func writePassFile(constName string, entries []vdltest.EntryValue) {
	const version vom.Version = vom.Version81
	fileName := constName + "_gen.vdl"
	// We skip all random entries for now, since they may contain multiple set and
	// map elements, which would cause non-deterministic output.
	//
	// TODO(toddw): Determine a better strategy.
	var filtered []vdltest.EntryValue
	for _, e := range entries {
		if !strings.HasPrefix(e.Label, "Random") {
			filtered = append(filtered, e)
		}
	}
	skipped := len(entries) - len(filtered)
	entries = filtered

	fmt.Printf("Writing %s:\t%d entries (%d skipped)\n", fileName, len(entries), skipped)
	file, cleanup := createFile(fileName)
	defer cleanup()
	writeHeader(file)
	imports := codegen.Imports{
		{Path: "v.io/v23/vdl/vdltest", Local: "vdltest"},
	}
	writef(file, "%s\n\n", vdlgen.Imports(imports))
	comment := textutil.PrefixLineWriter(file, "// ")
	panicOnError(vdltest.PrintEntryStats(comment, entries...))
	panicOnError(comment.Flush())
	writef(file, "\nconst %[1]s = []vdlEntry {\n", constName)
	for _, e := range entries {
		source := vdlgen.TypedConst(e.Source, "v.io/v23/vom/vomtest", imports)
		vomDump, hexType, hexValue := toVomHex(version, e.Source)
		writef(file, `%[1]s
	{
		%#[2]q,
		%#[3]q,
		%[3]s,
		0x%[4]x, %[5]q, %[6]q,
	},
`, vomDump, e.Label, source, byte(version), hexType, hexValue)
	}
	writef(file, "}\n")
}

func toVomHex(version vom.Version, value *vdl.Value) (_, _, _ string) {
	var buf, bufType bytes.Buffer
	encType := vom.NewVersionedTypeEncoder(version, &bufType)
	enc := vom.NewVersionedEncoderWithTypeEncoder(version, &buf, encType)
	panicOnError(enc.Encode(value))
	versionByte, _ := buf.ReadByte() // Read the version byte.
	if bufType.Len() > 0 {
		bufType.ReadByte() //nolint:errcheck // Remove the version byte.
	}
	vomBytes := append([]byte{versionByte}, bufType.Bytes()...)
	vomBytes = append(vomBytes, buf.Bytes()...)
	dump, err := vom.Dump(vomBytes)
	panicOnError(err)
	const pre = "\t// "
	vomDump := pre + strings.ReplaceAll(dump, "\n", "\n"+pre)
	if strings.HasSuffix(vomDump, "\n"+pre) {
		vomDump = vomDump[:len(vomDump)-len("\n"+pre)]
	}
	return vomDump, fmt.Sprintf("%x", bufType.Bytes()), fmt.Sprintf("%x", buf.Bytes())
}
