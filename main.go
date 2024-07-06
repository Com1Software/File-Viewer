package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	asciistring "github.com/Com1Software/Go-ASCII-String-Package"
)

var xip = fmt.Sprintf("%s", GetOutboundIP())

// ----------------------------------------------------------------
// ------------------------- (c) 1992-2024 Com1 Software Development
// ----------------------------------------------------------------

// Openbrowser : Opens default web browser to specified url
func Openbrowser(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start msedge"}
	case "linux":
		cmd = "chromium-browser"
		args = []string{""}

	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

// ------------------------------------------------------------------------
func DateTimeDisplay(xdata string) string {
	//------------------------------------------------------------------------
	xdata = xdata + "<script>"
	xdata = xdata + "function startTime() {"
	xdata = xdata + "  var today = new Date();"
	xdata = xdata + "  var d = today.getDay();"
	xdata = xdata + "  var h = today.getHours();"
	xdata = xdata + "  var m = today.getMinutes();"
	xdata = xdata + "  var s = today.getSeconds();"
	xdata = xdata + "  var ampm = h >= 12 ? 'pm' : 'am';"
	xdata = xdata + "  var mo = today.getMonth();"
	xdata = xdata + "  var dm = today.getDate();"
	xdata = xdata + "  var yr = today.getFullYear();"
	xdata = xdata + "  m = checkTimeMS(m);"
	xdata = xdata + "  s = checkTimeMS(s);"
	xdata = xdata + "  h = checkTimeH(h);"
	//------------------------------------------------------------------------
	xdata = xdata + "  switch (d) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       day = 'Sunday';"
	xdata = xdata + "    break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "    day = 'Monday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "        day = 'Tuesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "        day = 'Wednesday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "        day = 'Thursday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "        day = 'Friday';"
	xdata = xdata + "        break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "        day = 'Saturday';"
	xdata = xdata + "}"
	//------------------------------------------------------------------------------------
	xdata = xdata + "  switch (mo) {"
	xdata = xdata + "    case 0:"
	xdata = xdata + "       month = 'January';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 1:"
	xdata = xdata + "       month = 'Febuary';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 2:"
	xdata = xdata + "       month = 'March';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 3:"
	xdata = xdata + "       month = 'April';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 4:"
	xdata = xdata + "       month = 'May';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 5:"
	xdata = xdata + "       month = 'June';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 6:"
	xdata = xdata + "       month = 'July';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 7:"
	xdata = xdata + "       month = 'August';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 8:"
	xdata = xdata + "       month = 'September';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 9:"
	xdata = xdata + "       month = 'October';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 10:"
	xdata = xdata + "       month = 'November';"
	xdata = xdata + "       break;"
	xdata = xdata + "    case 11:"
	xdata = xdata + "       month = 'December';"
	xdata = xdata + "       break;"
	xdata = xdata + "}"
	//  -------------------------------------------------------------------
	xdata = xdata + "  document.getElementById('txtdt').innerHTML = ' '+h + ':' + m + ':' + s+' '+ampm+' - '+day+', '+month+' '+dm+', '+yr;"
	xdata = xdata + "  var t = setTimeout(startTime, 500);"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeMS(i) {"
	xdata = xdata + "  if (i < 10) {i = '0' + i};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	//----------
	xdata = xdata + "function checkTimeH(i) {"
	xdata = xdata + "  if (i > 12) {i = i -12};"
	xdata = xdata + "  return i;"
	xdata = xdata + "}"
	xdata = xdata + "</script>"
	return xdata

}

func LoopDisplay(xdata string) string {
	//------------------------------------------------------------------------
	xdata = xdata + "<script>"
	xdata = xdata + "function startLoop() {"
	//  -------------------------------------------------------------------
	xdata = xdata + "  document.getElementById('txtloop').innerHTML = Math.random();"
	xdata = xdata + "  var t = setTimeout(startLoop, 500);"
	xdata = xdata + "}"
	xdata = xdata + "</script>"
	return xdata

}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func InitPage(xip string) string {
	xxip := ""
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	xdata = xdata + "<title>Ville Viewer</title>"
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "</head>"
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<center>"
	xdata = xdata + "<H1>File Viewer</H1>"
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			xxip = fmt.Sprintf("%s", ipv4)
		}
	}
	xdata = xdata + "<div id='txtdt'></div>"
	xdata = xdata + "<p> Host Port IP : " + xip
	xdata = xdata + "<BR> Machine IP : " + xxip + "</p>"
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/about'> [ About ] </A> <BR><BR> "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/viewcsv'> [ View CSV File ] </A> <BR><BR> "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/viewxml'> [ View XML File ] </A>  "
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "</center>"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata
}

// ----------------------------------------------------------------
func AboutPage(xip string) string {
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	xdata = xdata + "<title>About Page</title>"
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "</head>"
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<center>"
	xdata = xdata + "<H1>File Viewer</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	xdata = xdata + "<BR>"
	xdata = xdata + "(c) 1992-2024 Com1 Software Development<BR><BR>"
	xdata = xdata + "  <A HREF='http://com1software.com'> [Com1 Software Web Site ] </A><BR><BR>  "
	xdata = xdata + "  <A HREF='https://github.com/Com1Software/Map-Utility'> [ Map Utility GitHub Repository ] </A> <BR><BR> "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "File Viewer"
	xdata = xdata + "</center>"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

// ----------------------------------------------------------------
func ViewCSV() string {
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	xdata = xdata + "<title>View CSV </title>"
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "</head>"
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<center>"
	xdata = xdata + "<H1>View CSV</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	xdata = xdata + "  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "
	xdata = xdata + "<BR><BR><BR><BR>"
	xdata = xdata + " Select a CSV file to Upload and View <BR><BR>"
	xdata = xdata + "<form  enctype='multipart/form-data' action='/csvfileupload' method='post'>"
	xdata = xdata + "<input type='file' name='csvFile'/>"
	xdata = xdata + "<input type='submit' value='Submit'/>"
	xdata = xdata + "</form>"
	xdata = xdata + "<BR><BR><BR>"
	xdata = xdata + " Cut and Paste Map to Validate<BR><BR>"
	xdata = xdata + "<form action='/mapvalidatereport' method='post'>"
	xdata = xdata + "<textarea id='map' name='map' rows='20' cols='150'></textarea>"
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "<input type='submit' value='Submit'/>"
	xdata = xdata + "</form>"
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "</center>"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata
}

// ----------------------------------------------------------------
func CSVViewer(csvinfo string) string {
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	xdata = xdata + "<title>CSV Viewer</title>"
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "</head>"
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<center>"
	xdata = xdata + "<h1>CSV Viewer</h1>"
	xdata = xdata + "<BR>"
	xdata = xdata + "<div id='txtdt'></div>"
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/viewcsv'> [ Return to CSV View ] </A>  "
	xdata = xdata + "<BR><BR>"
	xdata = xdata + ParseCSV(csvinfo)

	xdata = xdata + "<BR><BR>"
	xdata = xdata + "</center>"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

// ----------------------------------------------------------------
func uploadCSVFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("csvFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	csvString := string(fileBytes[:])
	xdata := CSVViewer(csvString)
	fmt.Fprint(w, xdata)
}

// ----------------------------------------------------------------
func ParseCSV(csvinfo string) string {
	xdata := ""
	ascval := 0
	chr := ""
	cc := 0
	rc := 0
	ccheck := true
	for x := 0; x < len(csvinfo); x++ {
		chr = csvinfo[x : x+1]
		ascval = asciistring.StringToASCII(chr)
		switch {
		case ascval == 13:
		case ascval == 10:
			rc++
			ccheck = false
		case ascval == 44:
			if ccheck {
				cc++
			}
		}
	}
	xdata = xdata + "Columns " + strconv.Itoa(cc) + "<BR>"
	xdata = xdata + "Rows " + strconv.Itoa(rc) + "<BR>"

	//	xdata = xdata + csvinfo
	return xdata

}

func main() {
	fmt.Println("Map Utility")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)

	port := "8080"
	switch {
	//-------------------------------------------------------------
	case len(os.Args) == 2:

		fmt.Println("Not")

		//-------------------------------------------------------------
	default:

		fmt.Println("Server running....")
		fmt.Println("Listening on " + xip + ":" + port)

		fmt.Println("")
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			xdata := InitPage(xip)
			fmt.Fprint(w, xdata)
		})
		//------------------------------------------------ About Page Handler
		http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
			xdata := AboutPage(xip)
			fmt.Fprint(w, xdata)
		})
		//--------------------------------------------------
		http.HandleFunc("/csvviewer", func(w http.ResponseWriter, r *http.Request) {
			mapinfo := r.FormValue("map")
			fmt.Println(mapinfo)
			xdata := CSVViewer(mapinfo)
			fmt.Fprint(w, xdata)

		})
		http.HandleFunc("/csvfileupload", uploadCSVFile)
		//--------------------------------------------------
		http.HandleFunc("/viewcsv", func(w http.ResponseWriter, r *http.Request) {
			xdata := ViewCSV()
			fmt.Fprint(w, xdata)

		})
		//------------------------------------------------- Static Handler Handler
		fs := http.FileServer(http.Dir("static/"))
		http.Handle("/static/", http.StripPrefix("/static/", fs))
		//------------------------------------------------- Start Server
		Openbrowser(xip + ":" + port)
		if err := http.ListenAndServe(xip+":"+port, nil); err != nil {
			panic(err)
		}
	}
}
