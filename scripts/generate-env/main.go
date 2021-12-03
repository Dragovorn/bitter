package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	argsLen := len(args)

	if argsLen == 0 {
		fmt.Println("No environment provided!")
	}

	environment := args[0]

	var isProd bool

	if argsLen == 2 {
		if result, err := strconv.ParseBool(args[1]); err != nil {
			panic(err)
		} else {
			isProd = result
		}
	}

	fmt.Println("Using environment: " + environment)

	var frontendDir string

	if tWorkDir, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		frontendDir = tWorkDir
	}

	baseDir := filepath.Dir(frontendDir + "/../")
	envFilePath := baseDir + "/.env." + environment
	angularEnvDir := frontendDir + "/src/environments"
	baseEnvironmentTs := angularEnvDir + "/environment.ts"
	environmentTs := angularEnvDir + "/environment." + environment + ".ts"

	fmt.Println("Frontend directory: " + frontendDir)
	fmt.Println("Base directory: " + baseDir)
	fmt.Println(".env file: " + envFilePath)
	fmt.Println("Angular environment directory: " + angularEnvDir)
	fmt.Println("Base environment.ts file: " + baseEnvironmentTs)
	fmt.Println("Provided environment.ts file: " + environmentTs)

	var envFile *os.File

	env := make(map[string]string)

	if tFile, err := os.Open(envFilePath); err != nil {
		panic(err)
	} else {
		envFile = tFile
	}

	defer envFile.Close()

	fmt.Println("Reading: " + envFilePath)
	fmt.Println()

	scanner := bufio.NewScanner(envFile)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.SplitN(line, "=", 2)
		key := split[0]
		value := split[1]

		fmt.Println(key + " = " + value)
		env[key] = value
	}

	fmt.Println()

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if _, err := os.Stat(angularEnvDir); os.IsNotExist(err) {
		if err2 := os.MkdirAll(angularEnvDir, os.ModePerm); err2 != nil {
			panic(err2)
		}
	}

	//baseEnvironmentTs = baseEnvironmentTs + ".test"

	deleteIfExists(baseEnvironmentTs)

	genEnvironmentTs(env, baseEnvironmentTs, isProd)
	genEnvironmentTs(env, environmentTs, isProd)
}

func deleteIfExists(path string) {
	if _, err := os.Stat(path); err != nil {
		if err2 := os.Remove(path); err2 != nil && !os.IsNotExist(err2) {
			panic(err2)
		}

		fmt.Println("Deleted: " + path)
	}
}

func genEnvironmentTs(env map[string]string, path string, isProd bool) {
	var file *os.File

	defer file.Close()

	if tFile, err := os.Create(path); err != nil {
		panic(err)
	} else {
		file = tFile
	}

	fmt.Println("Generating: " + path)

	writer := bufio.NewWriter(file)

	const prodLine = "  production: "

	writer.WriteString("export const environment = {\n")
	writer.WriteString(prodLine + strconv.FormatBool(isProd) + ",\n")

	for k, v := range env {
		writer.WriteString("  " + strings.ToLower(k) + ": " + v + ",\n")
	}

	writer.WriteString("};\n")
	writer.Flush()
}
