package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/7045kHz/http-client/core"
	"github.com/7045kHz/http-client/gohttp"
	"github.com/7045kHz/http-client/gomime"
)

type ConfigBenchmark struct {
	Name              string `json:"Name"`
	Description       string `json:"Description"`
	EndpointFile      string `json:"EndpointFile"`
	TlsInsecureVerify bool   `json:"TlsInsecureVerify"`
	Proxy             string `json:"Proxy"`
}
type Event struct {
	Events []Endpoints `json:"Events"`
}
type Endpoints struct {
	Name              string           `json:"Name"`
	Description       string           `json:"Description"`
	Url               string           `json:"Url"`
	PayloadFile       string           `json:"PayloadFile"`
	Method            string           `json:"Method"`
	SummaryOutputFile string           `json:"SummaryOutputFile"`
	BodyOutputFile    string           `json:"BodyOutputFile"`
	Headers           []EndpointHeader `json:"Headers"`
}

type EndpointHeader struct {
	Header_Key   string `json:"Header_Key"`
	Header_Value string `json:"Header_Value"`
}

var (
	httpClient = getHttpClient()
)

func main() {

	configFile := "config.json"
	C, err := ReadConfig(configFile)
	if err != nil {
		log.Printf("Config file read of file %s failed with error %v\n", configFile, err)
	}

	log.Println("=================================================================================================")
	log.Println()
	log.Printf("Name:            %s\n", C.Name)
	log.Printf("Description:     %s\n", C.Description)
	log.Printf("Endpoint File:   %s\n", C.EndpointFile)
	log.Printf("Insecure Verify: %v\n", C.TlsInsecureVerify)
	log.Printf("Proxy Server:    %s\n", C.Proxy)
	log.Println()
	log.Println("=================================================================================================")
	log.Println()

	E, _ := ReadEndpointFile(C.EndpointFile)
	RunEndPoints(E)

}
func RunEndPoints(E *Event) error {

	for i := 0; i < len(E.Events); i++ {

		fs, err := os.Create(E.Events[i].SummaryOutputFile)
		if err != nil {
			log.Printf("Could not create %s, %v\n", E.Events[i].SummaryOutputFile, err)
			log.Fatal("Exiting")
		}
		defer fs.Close()
		log.Println("--------------------------------------------------------------------------------------------------")
		log.Printf("Running test on endpoint %s\n", E.Events[i].Name)
		log.Printf("\tSee %s for summary output\n", E.Events[i].SummaryOutputFile)
		log.Printf("\tSee %s for Body output\n", E.Events[i].BodyOutputFile)
		fmt.Fprintf(fs, "%s Testing Event EndPoint: %s\n", time.Now().UTC().String(), E.Events[i].Name)

		fmt.Fprintln(fs, "=================================================================================================")
		fmt.Fprintf(fs, "Name:             %s\n", E.Events[i].Name)
		fmt.Fprintf(fs, "Description:      %s\n", E.Events[i].Description)
		fmt.Fprintf(fs, "URL:              %s\n", E.Events[i].Url)
		fmt.Fprintf(fs, "Method:           %s\n", E.Events[i].Method)
		fmt.Fprintf(fs, "Summary File:     %s\n", E.Events[i].SummaryOutputFile)
		fmt.Fprintf(fs, "Body Output File: %s\n", E.Events[i].BodyOutputFile)
		fmt.Fprintf(fs, "Payload File: %s\n", E.Events[i].PayloadFile)
		fmt.Fprintf(fs, "Custom Headers: [%d]\n", len(E.Events[i].Headers))
		_, err = url.ParseRequestURI(E.Events[i].Url)
		if err != nil {
			fmt.Fprintf(fs, "Invalid URL defined for the End Point (%s): (%s)\n", E.Events[i].Name, E.Events[i].Url)
			return err
		}
		headers := make(http.Header)
		for ik := 0; ik < len(E.Events[i].Headers); ik++ {

			log.Printf("\tLoading Header: %s = %s\n", E.Events[i].Headers[ik].Header_Key, E.Events[i].Headers[ik].Header_Value)

			headers.Set(E.Events[i].Headers[ik].Header_Key, E.Events[i].Headers[ik].Header_Value)
		}
		fmt.Fprintln(fs, "=================================================================================================")
		//log.Printf("\tHeaders: %v\n", headers)

		/*
			SENDING DATA TO SERVER
		*/

		fmt.Fprintf(fs, "%s EVENT [%s] Client Sending Url: %s\n", time.Now().UTC().String(), E.Events[i].Name, E.Events[i].Url)

		response := func() *core.Response {
			switch M := E.Events[i].Method; M {
			case "GET":
				response, err := httpClient.Get(E.Events[i].Url, headers)
				if err != nil {
					// Deal with the error as you need:
					fmt.Fprintf(fs, "%s Error Returned from Url: %s [%v]\n", time.Now().UTC().String(), E.Events[i].Url, err)
					return nil
				}
				fmt.Fprintf(fs, "%s EVENT [%s] Data Returned from Url: %s\n", time.Now().UTC().String(), E.Events[i].Name, E.Events[i].Url)
				return response

			case "POST":
				body, err := ReadPayload(string(E.Events[i].PayloadFile))
				if err != nil {
					// Deal with the error as you need:
					fmt.Fprintf(fs, "%s Error Returned sourcing file: %s [%v]\n", time.Now().UTC().String(), E.Events[i].PayloadFile, err)
					return nil
				}
				response, err := httpClient.Post(E.Events[i].Url, body, headers)
				if err != nil {
					// Deal with the error as you need:
					fmt.Fprintf(fs, "%s Error Returned from Url: %s [%v]\n", time.Now().UTC().String(), E.Events[i].Url, err)
					return nil
				}
				fmt.Fprintf(fs, "%s EVENT [%s] Data Returned from Url: %s\n", time.Now().UTC().String(), E.Events[i].Name, E.Events[i].Url)
				return response

			case "PUT":
				body, err := ReadPayload(string(E.Events[i].PayloadFile))
				if err != nil {
					// Deal with the error as you need:
					fmt.Fprintf(fs, "%s Error Returned sourcing file: %s [%v]\n", time.Now().UTC().String(), E.Events[i].PayloadFile, err)
					return nil
				}
				response, err := httpClient.Put(E.Events[i].Url, body, headers)
				if err != nil {
					// Deal with the error as you need:
					fmt.Fprintf(fs, "%s Error Returned from Url: %s [%v]\n", time.Now().UTC().String(), E.Events[i].Url, err)
					return nil
				}
				fmt.Fprintf(fs, "%s EVENT [%s] Data Returned from Url: %s\n", time.Now().UTC().String(), E.Events[i].Name, E.Events[i].Url)
				return response
			}
			return nil
		}()

		//fmt.Println(response)
		if response == nil {
			fmt.Fprintf(fs, "%s Request Failed.\n", time.Now().UTC().String())
			return err
		}

		/*
			END SENDING DATA TO SERVER
		*/
		fmt.Fprintf(fs, "%s Status Code: %d\n", time.Now().UTC().String(), response.StatusCode)
		fmt.Fprintf(fs, "%s Status: %s\n", time.Now().UTC().String(), response.Status)
		//fmt.Println(fmt.Sprintf("Body: %s\n", response.String()))
		if len(E.Events[i].BodyOutputFile) > 1 {
			fmt.Fprintf(fs, "%s Creating and logging body content in %s\n", time.Now().UTC().String(), E.Events[i].BodyOutputFile)
			fb, err := os.Create(E.Events[i].BodyOutputFile)

			if err != nil {
				fmt.Fprintf(fs, "%s Error creating %s: %v\n", time.Now().UTC().String(), E.Events[i].BodyOutputFile, err)
				log.Printf("Could not create %s, %v\n", E.Events[i].BodyOutputFile, err)
				log.Fatal("Exiting")
			}
			defer fb.Close()
			fmt.Fprintf(fb, "%s Event Test: %s\n", time.Now().UTC().String(), E.Events[i].Name)
			fmt.Fprintf(fb, "%s TARGET URL: %s\n", time.Now().UTC().String(), E.Events[i].Url)
			fmt.Fprintf(fb, "%s RESPONSE Status Code: %d\n", time.Now().UTC().String(), response.StatusCode)
			fmt.Fprintf(fb, "%s RESPONSE Headers: %s\n", time.Now().UTC().String(), response.Headers)
			fmt.Fprintf(fb, "%s RESPONSE Status: %s\n", time.Now().UTC().String(), response.Status)
			fmt.Fprintf(fb, "%s RESPONSE Body: %v\n", time.Now().UTC().String(), string(response.Body))
			fmt.Fprintf(fb, "%s Closing Test: %s\n", time.Now().UTC().String(), E.Events[i].Name)
		}
	}

	return nil
}

func ReadConfig(configFile string) (*ConfigBenchmark, error) {
	var C ConfigBenchmark
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Printf("File %s reading error %v\n", configFile, err)
		return &C, err
	}

	err = json.Unmarshal(data, &C)
	if err != nil {
		log.Printf("File %s Unmarshal of ConfigBenchmark Struct Error %v\n", configFile, err)
		return &C, err
	}

	return &C, nil
}
func ReadPayload(file string) (interface{}, error) {
	//fmt.Printf("Passed file %s\n", file)
	var C interface{}
	_, err := os.Stat(file)
	if err != nil {
		fmt.Printf("Missing file %s\n", file)
		return C, err
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("File %s reading error %v\n", file, err)
		return &C, err
	}

	C = string(data)
	//fmt.Printf("Body: %T  %v\n", C, C)
	return C, nil
}
func ReadEndpointFile(file string) (*Event, error) {
	var C Event
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("File %s reading error %v\n", file, err)
		return &C, err
	}

	err = json.Unmarshal(data, &C)
	if err != nil {
		log.Printf("File %s Unmarshal of ConfigBenchmark Struct Error %v\n", file, err)
		return &C, err
	}
	return &C, nil
}

func getHttpClient() gohttp.Client {
	headers := make(http.Header)
	headers.Set(gomime.HeaderContentType, gomime.ContentTypeJson)
	//log.Println("Starting new client")
	client := gohttp.NewBuilder().
		SetHeaders(headers).
		SetConnectionTimeout(2 * time.Second).
		SetResponseTimeout(3 * time.Second).
		SetUserAgent("Test-Computer").
		SetNtlm(true).
		SetTlsInsecureVerify(true).
		SetProxyServer("").
		Build()
	return client
}
