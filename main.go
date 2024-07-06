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
	"strings"

	asciistring "github.com/Com1Software/Go-ASCII-String-Package"
)

var xip = fmt.Sprintf("%s", GetOutboundIP())

// ----------------------------------------------------------------
// ------------------------- (c) 1992-2024 Com1 Software Development
// ----------------------------------------------------------------
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
		http.HandleFunc("/fieldstringdisplay", func(w http.ResponseWriter, r *http.Request) {
			fieldinfo := r.FormValue("field")
			xdata := FieldStringDisplay(fieldinfo, xip)
			fmt.Fprint(w, xdata)

		})

		//--------------------------------------------------
		http.HandleFunc("/mapvalidatereport", func(w http.ResponseWriter, r *http.Request) {
			mapinfo := r.FormValue("map")
			fmt.Println(mapinfo)
			xdata := MapValidateReport(mapinfo, xip)
			fmt.Fprint(w, xdata)

		})

		http.HandleFunc("/mapvalidatereportupload", uploadMapFile)

		//--------------------------------------------------
		http.HandleFunc("/mapvalidate", func(w http.ResponseWriter, r *http.Request) {
			xdata := MapValidate(xip)
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

func InitPage(xip string) string {
	//---------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	xxip := ""
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Go Web Server Start Page</title>"
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<center>"
	xdata = xdata + "<H1>Map Utility</H1>"
	//---------
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
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/fieldstringdisplay'> [ Field String Display ] </A> <BR><BR> "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/mapvalidate'> [ Map Validate ] </A>  "

	xdata = xdata + "<BR><BR>Map Utility"
	//------------------------------------------------------------------------
	xdata = xdata + "</center>"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata
}

// ----------------------------------------------------------------
func AboutPage(xip string) string {
	//---------------------------------------------------------------------------

	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>About Page</title>"
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)

	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: white;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: black;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p {"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------

	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<center>"
	xdata = xdata + "<p>Map Utility</p>"
	xdata = xdata + "<div id='txtdt'></div>"
	xdata = xdata + "<BR>"
	xdata = xdata + "(c) 1992-2024 Com1 Software Development<BR><BR>"
	xdata = xdata + "  <A HREF='http://com1software.com'> [Com1 Software Web Site ] </A><BR><BR>  "
	xdata = xdata + "  <A HREF='https://github.com/Com1Software/Map-Utility'> [ Map Utility GitHub Repository ] </A> <BR><BR> "
	xdata = xdata + "  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "Map Utility"
	xdata = xdata + "</center>"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

// ----------------------------------------------------------------
func MapValidate(xip string) string {
	//---------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Map Validate</title>"
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: white;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: black;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p1 {"
	xdata = xdata + "color: green;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	p2 {"
	xdata = xdata + "color: red;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	div {"
	xdata = xdata + "color: black;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<H1>Map Validate</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	//---------
	xdata = xdata + "<center>"
	xdata = xdata + "<p1>Map Utility</p1>"
	xdata = xdata + "<BR>"
	xdata = xdata + "<p2>Map Validate</p2>"
	xdata = xdata + "<BR><BR>"

	xdata = xdata + "  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "

	xdata = xdata + "<BR><BR><BR><BR>"
	xdata = xdata + " Select a Map file to Upload and Validate <BR><BR>"
	xdata = xdata + "<form  enctype='multipart/form-data' action='/mapvalidatereportupload' method='post'>"
	xdata = xdata + "<input type='file' name='mapFile'/>"
	xdata = xdata + "<input type='submit' value='Submit'/>"
	xdata = xdata + "</form>"
	xdata = xdata + "<BR><BR><BR>"
	//------------------------------------------------------------------------
	xdata = xdata + " Cut and Paste Map to Validate<BR><BR>"
	xdata = xdata + "<form action='/mapvalidatereport' method='post'>"
	xdata = xdata + "<textarea id='map' name='map' rows='20' cols='150'></textarea>"
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "<input type='submit' value='Submit'/>"
	xdata = xdata + "</form>"
	xdata = xdata + "<BR><BR>"

	//------------------------------------------------------------------------
	xdata = xdata + "</center>"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

// ----------------------------------------------------------------
func MapValidateReport(mapinfo string, xip string) string {
	//---------------------------------------------------------------------------
	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Map Validation Report</title>"
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: white;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: black;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p1 {"
	xdata = xdata + "color: green;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	p2 {"
	xdata = xdata + "color: red;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	div {"
	xdata = xdata + "color: black;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<H1>Map Validation Report</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	//---------
	xdata = xdata + "<center>"
	xdata = xdata + "<p1>Map Utility</p1>"
	xdata = xdata + "<BR>"
	xdata = xdata + "<p2>Map Validate</p2>"
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "  <A HREF='http://" + xip + ":8080/mapvalidate'> [ Return to Map Validate ] </A>  "
	xdata = xdata + "<BR><BR>"

	xdata = xdata + "<BR><BR>"
	xdata = xdata + MapValidation(mapinfo, xip)
	//------------------------------------------------------------------------
	xdata = xdata + "</center>"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

}

// ----------------------------------------------------------------
func MapValidation(mapinfo string, xip string) string {
	xdata := ""
	tdata := ""
	lc := 0
	chr := ""
	ascval := 0
	for x := 0; x < len(mapinfo); x++ {
		chr = mapinfo[x : x+1]
		ascval = asciistring.StringToASCII(chr)
		switch {
		case ascval == 13:
			lc++
		}
	}
	xdata = xdata + "The configuration contains " + strconv.Itoa(lc) + " lines.<BR>"
	fc := 0
	for x := 0; x < len(mapinfo); x++ {
		chr = mapinfo[x : x+1]
		ascval = asciistring.StringToASCII(chr)
		if ascval == 13 {
			tmp := strings.Split(strings.ToUpper(tdata), "=")
			switch {
			case tmp[0] == "X12TRANSACTIONTYPE":
				xdata = xdata + "The X12 trasnaction type is an " + tmp[1] + " .<BR>"

			case tmp[0] == "FIELD":
				fc++
			}
			tdata = ""

		} else {
			if ascval != 10 {
				tdata = tdata + chr

			}

		}

	}
	xdata = xdata + "The configuration contains " + strconv.Itoa(fc) + " fields.<BR>"

	return xdata

}

// ----------------------------------------------------------------
func FieldStringDisplay(fieldinfo string, xip string) string {
	//---------------------------------------------------------------------------
	msg := ""
	flda := ""
	fldb := ""
	fldc := ""
	fldd := ""
	flde := ""
	fldf := ""
	fldg := ""
	fldh := ""
	fldi := ""
	fldj := ""
	fldk := ""
	fldl := ""
	fldm := ""
	fldn := ""
	fldo := ""
	fldp := ""
	fldq := ""
	fldr := ""
	flds := ""
	fi := strings.TrimLeft(fieldinfo, " ")
	if len(fi) > 6 {
		if strings.ToLower(fi[0:6]) == "field=" {
			tmp := strings.Split(fi[6:len(fi)], ",")
			for x := 0; x < len(tmp); x++ {
				switch {
				case x == 0:
					flda = tmp[x]
				case x == 1:
					fldb = tmp[x]
				case x == 2:
					fldc = tmp[x]
				case x == 3:
					fldd = tmp[x]
				case x == 4:
					flde = tmp[x]
				case x == 5:
					fldf = tmp[x]
				case x == 6:
					fldg = tmp[x]
				case x == 7:
					fldh = tmp[x]
				case x == 8:
					fldi = tmp[x]
				case x == 9:
					fldj = tmp[x]
				case x == 10:
					fldk = tmp[x]
				case x == 11:
					fldl = tmp[x]
				case x == 12:
					fldm = tmp[x]
				case x == 13:
					fldn = tmp[x]
				case x == 14:
					fldo = tmp[x]
				case x == 15:
					fldp = tmp[x]
				case x == 16:
					fldq = tmp[x]
				case x == 17:
					fldr = tmp[x]
				case x == 18:
					flds = tmp[x]

				}

			}

			msg = fieldinfo + "<BR><BR>Successfully Loaded.<BR><BR>Enter Another Field Line Sring and Submit"
		} else {
			msg = "Invalid Format"
		}
	} else {
		msg = "Enter a Field Line Sring and Submit"
	}
	//----------------------------------------------------------------------------
	xdata := "<!DOCTYPE html>"
	xdata = xdata + "<html>"
	xdata = xdata + "<head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<title>Field String Display</title>"
	//------------------------------------------------------------------------
	xdata = DateTimeDisplay(xdata)
	xdata = xdata + "<style>"
	xdata = xdata + "body {"
	xdata = xdata + "    background-color: white;"
	xdata = xdata + "}"
	xdata = xdata + "	h1 {"
	xdata = xdata + "	color: black;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "	p1 {"
	xdata = xdata + "color: green;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	p2 {"
	xdata = xdata + "color: red;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "}"
	xdata = xdata + "	div {"
	xdata = xdata + "color: black;"
	xdata = xdata + "font-family: verdana;"
	xdata = xdata + "	font-size: 20px;"
	xdata = xdata + "	text-align: center;"
	xdata = xdata + "}"
	xdata = xdata + "</style>"
	xdata = xdata + "</head>"
	//------------------------------------------------------------------------
	xdata = xdata + "<body onload='startTime()'>"
	xdata = xdata + "<H1>Field String Display</H1>"
	xdata = xdata + "<div id='txtdt'></div>"
	//---------
	xdata = xdata + "<center>"
	xdata = xdata + "<p1>Map Utility</p1>"
	xdata = xdata + "<BR>"
	xdata = xdata + "<p2>Field String Display</p2>"
	xdata = xdata + "<BR><BR>"

	xdata = xdata + "  <A HREF='http://" + xip + ":8080'> [ Return to Start Page ] </A>  "

	xdata = xdata + "<BR><BR>"
	//------------------------------------------------------------------------
	xdata = xdata + "<form action='/fieldstringdisplay' method='post'>"
	xdata = xdata + "Enter Field String: <input type='text' name='field' /><br/>"
	xdata = xdata + "<BR><BR>"
	xdata = xdata + "<input type='submit' value='Submit'/>"
	xdata = xdata + "</form>"

	xdata = xdata + "<BR><BR>"
	xdata = xdata + "FIELD<BR>"
	xdata = xdata + "Syntax  FIELD=a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s<BR><BR>"

	xdata = xdata + "A - Field Name : " + flda + "<BR>"
	xdata = xdata + "B - Field Width : " + fldb + "<BR>"
	xdata = xdata + "C - Segment Type : " + fldc + "<BR>"
	xdata = xdata + "D - Segment Qualifier : " + fldd + "<BR>"
	xdata = xdata + "E - Qualifier Position : " + flde + "<BR>"
	xdata = xdata + "F - Element Position : " + fldf + "<BR>"
	xdata = xdata + "G - Sub Element Position : " + fldg + "<BR>"
	xdata = xdata + "H - Conditional Field : " + fldh + "<BR>"
	xdata = xdata + "I -Occurance : " + fldi + "<BR>"
	xdata = xdata + "J - Level : " + fldj + "<BR>"
	xdata = xdata + "K - Service Line Number : " + fldk + "<BR>"
	xdata = xdata + "L - Field Description : " + fldl + "<BR>"
	xdata = xdata + "M - Data Starting Position within the Field : " + fldm + "<BR>"
	xdata = xdata + "N - Length of Data : " + fldn + "<BR>"
	xdata = xdata + "O = Function : " + fldo + "<BR>"
	xdata = xdata + "P - Previous Segment : " + fldp + "<BR>"
	xdata = xdata + "Q - Previous Segment Qualifier : " + fldq + "<BR>"
	xdata = xdata + "R - Previous Segment Qualofier Position : " + fldr + "<BR>"
	xdata = xdata + "S - Toggle Off Field : " + flds + "<BR>"

	xdata = xdata + "<BR><BR>" + msg + "<BR>"
	//------------------------------------------------------------------------
	xdata = xdata + "</center>"
	xdata = xdata + " </body>"
	xdata = xdata + " </html>"
	return xdata

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

func uploadMapFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("mapFile")
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
	mapString := string(fileBytes[:])
	xdata := MapValidateReport(mapString, xip)
	fmt.Fprint(w, xdata)
}
