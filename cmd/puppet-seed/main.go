package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type LanguageSeed struct {
	Name           string
	Version        string
	ImageName      string
	FileName       string
	CompileCommand string
	RunCommand     string
}

var languages = []LanguageSeed{
	{"C", "11", "gcc:latest", "main.c", "gcc main.c -o main", "./main"},
	{"C++", "17", "gcc:latest", "main.cpp", "g++ main.cpp -o main", "./main"},
	{"Python", "3.11", "python:3.11", "main.py", "", "python3 main.py"},
	{"Go", "1.21", "golang:1.21", "main.go", "", "go run main.go"},
	{"Rust", "1.70", "rust:1.70", "main.rs", "rustc main.rs", "./main"},
	{"Java", "17", "openjdk:17", "Main.java", "javac Main.java", "java Main"},
	{"Kotlin", "1.8", "openjdk:17", "Main.kt", "kotlinc Main.kt -include-runtime -d Main.jar", "java -jar Main.jar"},
	{"JavaScript", "18", "node:18", "main.js", "", "node main.js"},
	{"TypeScript", "5", "node:18", "main.ts", "tsc main.ts", "node main.js"},
	{"PHP", "8.2", "php:8.2", "main.php", "", "php main.php"},
	{"Ruby", "3.2", "ruby:3.2", "main.rb", "", "ruby main.rb"},
	{"Swift", "5.8", "swift:5.8", "main.swift", "swiftc main.swift -o main", "./main"},
	{"Perl", "5.36", "perl:5.36", "main.pl", "", "perl main.pl"},
	{"Haskell", "9.4", "haskell:9.4", "main.hs", "ghc main.hs", "./main"},
	{"Scala", "3", "hseeberger/scala-sbt:11.0.16_1.7.1_3.2.2", "Main.scala", "scalac Main.scala", "scala Main"},
	{"Bash", "5.1", "bash:5.1", "main.sh", "chmod +x main.sh", "./main.sh"},
	{"Dart", "3.1", "dart:stable", "main.dart", "", "dart main.dart"},
	{"Elixir", "1.15", "elixir:1.15", "main.ex", "", "elixir main.ex"},
	{"C#", "7.0", "mcr.microsoft.com/dotnet/sdk:7.0", "Program.cs", "dotnet build -o out", "dotnet out/Program.dll"},
	{"Julia", "1.9", "julia:1.9", "main.jl", "", "julia main.jl"},
	{"OCaml", "4.14", "ocaml/opam", "main.ml", "ocamlc -o main main.ml", "./main"},
	{"Nim", "2.0", "nimlang/nim:2.0.0", "main.nim", "nim c -r main.nim", "./main"},
	{"Crystal", "1.9", "crystallang/crystal:1.9.2", "main.cr", "crystal build main.cr", "./main"},
	{"Vlang", "0.4", "vlang/vlang:latest", "main.v", "v run main.v", "./main"},
	{"Zig", "0.11", "ziglang/zig:0.11.0", "main.zig", "zig build-exe main.zig", "./main"},
	{"Fortran", "10", "gcc:latest", "main.f90", "gfortran main.f90 -o main", "./main"},
	{"COBOL", "gnu-cobol", "leopardslab/gnu-cobol", "main.cob", "cobc -x main.cob", "./main"},
	{"Erlang", "26", "erlang:26", "main.erl", "erlc main.erl", "erl -noshell -s main start -s init stop"},
	{"F#", "7.0", "mcr.microsoft.com/dotnet/sdk:7.0", "Program.fs", "dotnet fsi Program.fs", "dotnet run"},
	{"Smalltalk", "2023", "hpi-swa/smalltalk", "main.st", "", "gst main.st"},
	{"Scheme", "9.5", "racket/racket:latest", "main.scm", "", "racket main.scm"},
	{"R", "4.3", "r-base:4.3", "main.r", "", "Rscript main.r"},
	{"Groovy", "4.0", "groovy:4.0", "main.groovy", "", "groovy main.groovy"},
	{"Lua", "5.4", "lua:5.4", "main.lua", "", "lua main.lua"},
	{"Pascal", "3.2", "fpc:3.2", "main.pas", "fpc main.pas", "./main"},
	{"Prolog", "1.4", "esolang/prolog", "main.pl", "", "swipl -q -f main.pl"},
	{"Objective-C", "clang", "clang:latest", "main.m", "clang main.m -o main -lobjc -framework Foundation", "./main"},
	{"Assembly", "nasm", "ubuntu", "main.asm", "nasm -f elf64 main.asm && ld -o main main.o", "./main"},
	{"AWK", "5.1", "awk:latest", "main.awk", "", "awk -f main.awk"},
	{"Brainfuck", "2.7", "esolang/brainfuck", "main.bf", "", "bf main.bf"},
}

func main() {
	dbURL := "postgres://puppet_user:puppet_pass@localhost:5432/puppet?sslmode=disable"

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	stmt := `
	INSERT INTO languages (
		name, version, image_name, file_name, compile_command, run_command, installed, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, false, $7, $7)
	ON CONFLICT (image_name) DO NOTHING;
	`

	now := time.Now()
	for _, lang := range languages {
		_, err := db.Exec(stmt, lang.Name, lang.Version, lang.ImageName, lang.FileName, lang.CompileCommand, lang.RunCommand, now)
		if err != nil {
			fmt.Printf("Failed to insert %s: %v\n", lang.Name, err)
		} else {
			fmt.Printf("Seeded: %s %s\n", lang.Name, lang.Version)
		}
	}
}
