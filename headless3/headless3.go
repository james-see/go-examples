package headless3

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/chromedp/chromedp"
	cdp "github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/runner"
	"github.com/raff/godet"
)

var chromeapp string

func lowhangingfruits(username string) (yesno bool) {
	client := &http.Client{}
	fmt.Println(username)
	fmt.Printf("https://www.instagram.com/%s/", username)
	req, err := http.NewRequest("GET", fmt.Sprintf("https://www.instagram.com/%s/", username), nil)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")
	gotcha, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(gotcha.StatusCode)
	fmt.Println("lol")
	if gotcha.StatusCode != 200 {
		return false
	} else {
		return true
	}
}

func createRequest() *http.Client {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://httpbin.org/user-agent", nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
	fmt.Println(resp.StatusCode)
	return client
}

// LoadJS : This is it!
func LoadJS(username string, remote *godet.RemoteDebugger) {
	remote.RuntimeEvents(true)
	remote.NetworkEvents(true)
	remote.PageEvents(true)
	remote.DOMEvents(true)
	remote.LogEvents(true)
	fmt.Printf("https://www.instagram.com/%s/", username)
	fmt.Println()
	_, _ = remote.Navigate(fmt.Sprintf("https://www.instagram.com/%s/", username))
	_ = remote.SaveScreenshot("screenshot.png", 0644, 0, true)

}

// Headless3 : It works!
func Headless3(defaultuser string) {
	switch runtime.GOOS {
	case "darwin":
		chromeapp = `open "/Applications/Google Chrome Canary.app" --args`
	case "linux":
		chromeapp = "chromium-browser"
	}
	if chromeapp != "" {
		chromeapp = " --headless --remote-debugging-port=9222 --hide-scrollbars"
	}
	exec.Command(chromeapp).Run()
	//exec.Command("open -a '/Applications/Google\\ Chrome\\ Canary.app' --args --headless --remote-debugging-port=9222 --hide-scrollbars --remote-debugging-address=0.0.0.0").Run()
	port := "localhost:9222"
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for i := range nums {
		if i > 0 {
			time.Sleep(500 * time.Millisecond)
		}
		remote, err := godet.Connect(port, false)
		if err == nil { // connection succeeded
			break
		}
		log.Println("connect", err, remote)

		if err != nil {
			log.Println("cannot connect to browser")
		}
	}
	remote, err := godet.Connect(port, false)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	remote.RuntimeEvents(true)
	remote.NetworkEvents(true)
	remote.PageEvents(true)
	remote.DOMEvents(true)
	remote.LogEvents(true)
	if defaultuser != "defaultuser" {
		LoadJS(defaultuser, remote)
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter instagram username to check and screenshot: ")
		text, _ := reader.ReadString('\n')
		fmt.Printf("Username: %s", text)
		LoadJS(text, remote)
	}

}

//Headless4 is using the preferred chromedp package for headless chrome interaction
func Headless4(instagramuser string) (texter string) {
	var err error
	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()
	// create chrome instance
	c, err := cdp.New(ctxt, cdp.WithRunnerOptions(
		runner.Flag("headless", true),
		runner.Flag("disable-gpu", true),
		runner.Flag("hide-scrollbars", true)))
	if err != nil {
		log.Fatal(err)
	}
	// run task list
	var buf []byte // holder for the image data from the screenshot call
	url := fmt.Sprintf("https://www.instagram.com/%s/", instagramuser)
	err = c.Run(ctxt, screenshot(url, `react-root`, &buf, &texter))
	if err != nil {
		log.Fatal(err)
	}
	// shutdown chrome
	// err = c.Shutdown(ctxt)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // wait for chrome to finish
	// err = c.Wait()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// write screenshot to file
	err = ioutil.WriteFile("instagram-profile.png", buf, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// print captured text from page
	fmt.Println(texter)
	return texter
}

func screenshot(urlstr, sel string, res *[]byte, textdata *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Sleep(2 * time.Second),
		//chromedp.WaitReady(sel),
		chromedp.Text(sel, textdata, chromedp.NodeVisible, chromedp.ByID),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID),
	}
}
