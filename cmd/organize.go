package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var folders = map[string]string{
	// Documents
	".pdf":  "PDFs",
	".doc":  "Documents",
	".docx": "Documents",
	".xls":  "Spreadsheets",
	".xlsx": "Spreadsheets",
	".ppt":  "Presentations",
	".pptx": "Presentations",
	".txt":  "Texts",
	".csv":  "Spreadsheets",
	".md":   "Documents",

	// Images
	".png":  "Images",
	".jpg":  "Images",
	".jpeg": "Images",
	".gif":  "Images",
	".svg":  "Images",
	".webp": "Images",
	".ico":  "Images",
	".bmp":  "Images",

	// Videos
	".mp4":  "Videos",
	".mov":  "Videos",
	".avi":  "Videos",
	".mkv":  "Videos",
	".webm": "Videos",

	// Audio
	".mp3":  "Audio",
	".wav":  "Audio",
	".flac": "Audio",
	".aac":  "Audio",
	".ogg":  "Audio",

	// Archives
	".zip": "Archives",
	".rar": "Archives",
	".7z":  "Archives",
	".tar": "Archives",
	".gz":  "Archives",

	// Code
	".go":   "Code",
	".js":   "Code",
	".ts":   "Code",
	".py":   "Code",
	".html": "Code",
	".css":  "Code",
	".json": "Code",

	// Fonts
	".ttf":  "Fonts",
	".otf":  "Fonts",
	".woff": "Fonts",

	// Installers
	".exe": "Executables",
	".msi": "Executables",
	".dmg": "Executables",
}

func exit(err string) {
	fmt.Println(err)
	os.Exit(1)
}

func organizeFolder(cmd *cobra.Command, args []string) {
	dir := args[0]

	files, err := os.ReadDir(dir)

	if err != nil {
		exit("Error reading the file")
	}

	countFile := 0

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := filepath.Ext(file.Name())

		folder, exists := folders[ext]

		if !exists {
			continue
		}

		newFolderPath := filepath.Join(dir, folder)
		err := os.MkdirAll(newFolderPath, os.ModePerm)

		if err != nil {
			exit("Error while creating folder")
		}

		oldPath := filepath.Join(dir, file.Name())
		newPath := filepath.Join(newFolderPath, file.Name())

		err = os.Rename(oldPath, newPath)

		if err != nil {
			exit("error while moving to new folder")
		}

		countFile++
	}

	fmt.Printf("%d files has been successfully categorized\n", countFile)

}

var organizeCmd = &cobra.Command{
	Use:   "organize [path]",
	Short: "Organize files in a directory by type",
	Long:  "Scan a directory and automatically sort files into folders based on their file extension (e.g., Images, Documents, Videos).",
	Args:  cobra.ExactArgs(1),
	Run:   organizeFolder,
}

func init() {
	rootCmd.AddCommand(organizeCmd)
}
