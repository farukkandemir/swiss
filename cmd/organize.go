package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func organizeFolder(cmd *cobra.Command, args []string) {
	dir := args[0]

	files, err := os.ReadDir(dir)

	if err != nil {
		fmt.Println("Error happened", err)
		os.Exit(1)
	}

	folders := map[string]string{
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

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		ext := filepath.Ext(file.Name())

		folder, exists := folders[ext]

		if !exists {
			continue
		}

		folderPath := filepath.Join(dir, folder)

		err := os.MkdirAll(folderPath, os.ModePerm)

		if err != nil {
			fmt.Println("Error while creating a dir", err)
			os.Exit(1)
		}

		oldPath := filepath.Join(dir, file.Name())
		err = os.Rename(oldPath, filepath.Join(folderPath, file.Name()))

		if err != nil {
			fmt.Println("Error while moving file to new folder", err)
			os.Exit(1)
		}

		fmt.Println("Successful")

	}

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
