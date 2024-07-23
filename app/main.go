package main

// todo, indicate on each exercices if some modules are authorised

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Categories and their corresponding subjects
var categories = map[string][]string{
	"10%":  {"DisplayAlpham", "DisplayAlrevm", "DisplayFirstparam", "DisplayLastparam", "ParamCount", "PrintDigits", "CheckNumber"},
	"20%":  {"CountDown", "StrLen", "WdMatch", "PrintIf", "PrintIfNot", "RectPerimeter"},
	"30%":  {"FirstRune", "LastRune", "LastWord", "ReduceInt", "HashCode"},
	"40%":  {"AlphaMirror", "Chunk", "Doop", "FindPrevPrime", "FoldInt", "SearchReplace", "TabMult", "RetainFirstHalf"},
	"50%":  {"Inter", "IsPowerof2", "PigLatin", "PrintBits", "ReverseBits", "RomanNumbers", "SwapBits", "Union"},
	"60%":  {"Gcd", "PrintHex", "RepeatAlpha", "SortWordArr", "ZipString"},
	"70%":  {"AddPrimeSum", "Fprime", "Hiddenp", "Revwstr", "Rostring", "Split"},
	"80%":  {"AtoiBase", "Itoa"},
	"90%":  {"Brackets", "ListSize", "Options", "PrintMemory", "RpnCalc"},
	"100%": {"Brainfuck", "ItoaBase", "ListRemoveIf"},
	// Add more categories and subjects as needed
}

// const subjectDir = "app/subjects/"
const subjectDir = "./app/.subjects/"
const targetDir = "./test/"

func main() {
	var selectedSubjects []string

	// Create target directory if it does not exist
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	// Remove any existing files in the target directory
	if err := removeFilesInDir(targetDir); err != nil {
		fmt.Printf("Error removing files in target directory: %v\n", err)
		return
	}

	// Squidward?
	SquidwardWisdom := []string{
		"Another day, another migraine.",
		"Oh, my aching tentacles!",
		"It all started when I was born.",
		"Here, please hit me as hard as you can.",
		"Too bad that didn't kill me.",
		"Being dead, or anything else!",
		"I'm not just ready to go to work, I'm ready to go home.",
		"Happiness is overrated.",
		"Happiness? Is that something you eat?",
		"Happiness is fleeting, but misery is forever.",
		"Why is it that whenever I'm having fun, it's wrong?",
		"Do you have barnacles for brains?",
		"I don't know what's worse: your breath or your ideas.",
		"Just get me out of here!",
	}
	randomSentence := chooseRandomString(SquidwardWisdom)
	printAscii(randomSentence)

	// Process each category in order and copy a randomly selected subject file
	for _, category := range getOrderedCategories() {
		subjects := categories[category]
		selectedSubject := subjects[rand.Intn(len(subjects))]
		fmt.Printf("%s - %s\n", category, selectedSubject)
		selectedSubjects = append(selectedSubjects, selectedSubject)
		// copy the go file
		copyFile(filepath.Join(subjectDir, strings.ToLower(selectedSubject)+".go"), filepath.Join(targetDir, category+"_"+strings.ToLower(selectedSubject)+".go"))
		// copy the markdown file
		copyFile(filepath.Join(subjectDir, strings.ToLower(selectedSubject)+".md"), filepath.Join(targetDir, category+"_"+strings.ToLower(selectedSubject)+".md"))
	}

	// Initialize go.mod in the target directory
	if err := initGoMod(targetDir); err != nil {
		fmt.Printf("Error initializing go.mod: %v\n", err)
		return
	}

	// Create main.go in the target directory
	createMainGo(targetDir, selectedSubjects)

	// Run go mod tidy to ensure all dependencies are resolved
	if err := tidyGoMod(targetDir); err != nil {
		fmt.Printf("Error running go mod tidy: %v\n", err)
		return
	}

	// Launch VSCode pointing to the target directory
	openWithDefaultCodeEditor(targetDir)
}

// getOrderedCategories returns the categories in a fixed order
func getOrderedCategories() []string {
	return []string{
		"10%", "20%", "30%", "40%", "50%", "60%", "70%", "80%", "90%", "100%",
	}
}

// removeFilesInDir removes all files in the specified directory
func removeFilesInDir(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}

	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) {
	sourceFile, err := os.Open(src)
	if err != nil {
		fmt.Printf("Error opening source file: %v\n", err)
		return
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		fmt.Printf("Error creating destination file: %v\n", err)
		return
	}
	defer destinationFile.Close()

	if _, err := io.Copy(destinationFile, sourceFile); err != nil {
		fmt.Printf("Error copying file: %v\n", err)
		return
	}
}

// initGoMod initializes a Go module in the specified directory and adds the required import
func initGoMod(dir string) error {
	cmd := exec.Command("go", "mod", "init", "test")
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to initialize go.mod: %v, output: %s", err, string(output))
	}

	// Add the required import to go.mod
	modFilePath := filepath.Join(dir, "go.mod")
	f, err := os.OpenFile(modFilePath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("failed to open go.mod: %v", err)
	}
	defer f.Close()

	if _, err := f.WriteString("\nrequire github.com/01-edu/z01 v0.1.0\n"); err != nil {
		return fmt.Errorf("failed to write to go.mod: %v", err)
	}

	return nil
}

// createMainGo creates a main.go file in the specified directory that calls functions from the copied files
func createMainGo(dir string, subjects []string) {
	mainFilePath := filepath.Join(dir, "main.go")
	mainFile, err := os.Create(mainFilePath)
	if err != nil {
		fmt.Printf("Error creating main.go: %v\n", err)
		return
	}
	defer mainFile.Close()

	// var imports []string
	var functionCalls []string

	for _, subject := range subjects {
		// importPath := fmt.Sprintf("\"./%s\"", subject)
		functionName := fmt.Sprintf("//%s()", strings.Title(subject))
		// imports = append(imports, importPath)
		functionCalls = append(functionCalls, functionName)
	}

	mainContent := fmt.Sprintf(`package main

import	"github.com/01-edu/z01"

func main() {
	z01.PrintRune('\n') // here to initialize the z01 package
	%s
}
`, strings.Join(functionCalls, "\n\t"))
	// `, strings.Join(imports, "\n\t"), strings.Join(functionCalls, "\n\t"))

	if _, err := mainFile.WriteString(mainContent); err != nil {
		fmt.Printf("Error writing to main.go: %v\n", err)
		return
	}
}

// tidyGoMod runs `go mod tidy` in the specified directory
func tidyGoMod(dir string) error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("failed to run go mod tidy: %v, output: %s", err, string(output))
	}
	return nil
}

// openWithDefaultCodeEditor opens the specified directory with the default code editor
func openWithDefaultCodeEditor(dir string) {
	var cmd *exec.Cmd

	// Attempt to use Visual Studio Code if installed
	if path, err := exec.LookPath("code"); err == nil {
		cmd = exec.Command(path, dir)
	} else if path, err := exec.LookPath("codium"); err == nil {
		cmd = exec.Command(path, dir)
	} else {
		// Use default editors based on the operating system
		switch runtime.GOOS {
		case "windows":
			cmd = exec.Command("notepad.exe", filepath.Join(dir, ""))
		case "darwin":
			cmd = exec.Command("open", "-a", "TextEdit", dir)
		case "linux":
			cmd = exec.Command("xdg-open", dir)
		default:
			fmt.Printf("Unsupported platform: %s\n", runtime.GOOS)
			return
		}
	}

	if err := cmd.Start(); err != nil {
		fmt.Printf("Error opening directory with default editor: %v\n", err)
	}
}

// because we all need some motivation
func printAscii(str string) {
	fmt.Printf(`
⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⡀⠠⠤⠀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣀⢤⡒⠉⠁⠀⠒⢂⡀⠀⠀⠀⠈⠉⣒⠤⣀⠀⠀⠀⠀
⠀⠀⣠⠾⠅⠈⠀⠙⠀⠀⠀⠈⠀⠀⢀⣀⣓⡀⠉⠀⠬⠕⢄⠀⠀
⠀⣰⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡤⠶⢦⡀⠑⠀⠀⠀⠀⠈⢧⠀
⠀⡇⠀⠀⠀⠀⠀⢤⣀⣀⣀⣀⡀⢀⣀⣀⠙⠀⠀⠀⠀⠀⠀⢸⡄
⠀⢹⡀⠀⠀⠀⠀⡜⠁⠀⠀⠙⡴⠁⠀⠀⠱⡄⠀⠀⠀⠀⠀⣸⠀
⠀⠀⠱⢄⡀⠀⢰⣁⣒⣒⣂⣰⣃⣀⣒⣒⣂⢣⠀⠀⠀⢀⡴⠁⠀
⠀⠀⠀⠀⠙⠲⢼⡀⠀⠙⠀⢠⡇⠀⠛⠀⠀⣌⣀⡤⠖⠉⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢸⡗⢄⣀⡠⠊⠈⢦⣀⣀⠔⡏⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠈⡇⠀⢰⠁⠀⠀⠀⢣⠀⠀⣷⠀⠀⠀⠀%v
⠀⠀⠀⠀⣠⠔⠊⠉⠁⡏⠀⠀⠀⠀⠘⡆⠤⠿⣄⣀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣧⠸⠒⣚⡩⡇⠀⠀⠀⠀⠀⣏⣙⠒⢴⠈⡇⠀⠀⠀⠀
⠀⠀⠀⠀⠈⠋⠉⠀⠀⢳⡀⠀⠀⠀⣸⠁⠈⠉⠓⠚⠁⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠓⠛⠛

`, str)
}

func chooseRandomString(choices []string) string {
	// Seed the random number generator to ensure different results each run

	// Generate a random index in the range [0, len(choices)-1]
	randomIndex := rand.Intn(len(choices))

	// Return the string at the random index
	return choices[randomIndex]
}
