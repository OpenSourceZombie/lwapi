package lwapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//virtualSErverAPI leaseweb API endpoint
const virtualServersAPI = "https://api.leaseweb.com/cloud/v2/virtualServers"
const (
	//POWERON The power on call is an asynchronous call, meaning the power on will be performed as soon as possible.
	POWERON = iota
	//POWEROFF The power off call is an asynchronous call, meaning the power off will be performed as soon as possible.
	POWEROFF = iota
	//REBOOT The reboot call is an asynchronous call, meaning the reboot will be performed as soon as possible.
	REBOOT = iota
)
const (
	//OPERATINGSYSTEM credentials, you will only be able to retrieve the last password that we configured in your server or control panel. If you changed it, the information retrieved by this API call will not work.
	OPERATINGSYSTEM = iota
	//CONTROLPANEL credentials
	CONTROLPANEL = iota
)

//LW Main struct you will use to make requests
type LW struct {
	AuthToken  string
	httpclient http.Client
}

//GetVirtualServersList Get a list of all virtual servers
func (lw LW) GetVirtualServersList() (VirtualServerList, error) {
	vsl := VirtualServerList{}
	req, _ := http.NewRequest("GET", virtualServersAPI, nil)
	req.Header.Add("X-Lsw-Auth", lw.AuthToken)
	body, err := lw.sendHTTPRequest(req, http.StatusOK)
	if err != nil {
		return vsl, err
	}
	err = json.Unmarshal(body, &vsl)
	if err != nil {
		log.Println(err)
	}
	return vsl, err

}

//GetVirtualServer Get information about one server by providing the server's id
func (lw LW) GetVirtualServer(id string) (VirtualServer, error) {
	vs := VirtualServer{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", virtualServersAPI, id), nil)
	req.Header.Add("X-Lsw-Auth", lw.AuthToken)
	body, err := lw.sendHTTPRequest(req, http.StatusOK)
	if err != nil {
		return vs, errors.New("Error " + err.Error())
	}
	err = json.Unmarshal(body, &vs)
	if err != nil {
		log.Println(err)
	}
	return vs, err
}

//UpdateServerReference Update the current server reference
func (lw LW) UpdateServerReference(id, reference string) (VirtualServer, error) {
	vs := VirtualServer{}
	var newReference = []byte(fmt.Sprintf(`{"reference":"%s"}`, reference))
	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/%s", virtualServersAPI, id), bytes.NewBuffer(newReference))
	req.Header.Add("X-Lsw-Auth", lw.AuthToken)
	body, err := lw.sendHTTPRequest(req, http.StatusOK)
	if err != nil {
		return vs, errors.New("Error " + err.Error())
	}
	err = json.Unmarshal(body, &vs)
	if err != nil {
		log.Println(err)
	}
	return vs, err
}

//PowerControl PowerOn, PowerOff or reboot a virtual server, powerOn = 0, powerOff =1, reboot=2
//You can also use lwapi.POWERON, lwapi.POWEROFF, lwapi.REBOOT as an alternative
//This call is an asynchronous call, meaning the power off will be performed as soon as possible.
func (lw LW) PowerControl(id string, powerState int) (AsyncResponse, error) {
	var state string
	asyncResp := AsyncResponse{}

	switch powerState {
	case 0:
		state = "powerOn"
	case 1:
		state = "powerOff"
	case 2:
		state = "reboot"
	default:
		return asyncResp, errors.New("Invalid Power State")
	}
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/%s/%s", virtualServersAPI, id, state), nil)
	req.Header.Add("X-Lsw-Auth", lw.AuthToken)
	body, err := lw.sendHTTPRequest(req, http.StatusAccepted)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(body, &asyncResp)
	if err != nil {
		return asyncResp, errors.New("Error " + err.Error())
	}
	return asyncResp, err
}

//ReinstallVirtualServer reinstall a virtual serven providing it's virtual server's id
//This call is an asynchronous call, meaning the power off will be performed as soon as possible.
func (lw LW) ReinstallVirtualServer(id string) (AsyncResponse, error) {

	asyncResp := AsyncResponse{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/%s/%s", virtualServersAPI, id, "reinstall"), nil)
	req.Header.Add("X-Lsw-Auth", lw.AuthToken)
	body, err := lw.sendHTTPRequest(req, http.StatusAccepted)
	if err != nil {
		return asyncResp, errors.New("Something went wrong: " + err.Error())
	}
	err = json.Unmarshal(body, &asyncResp)
	if err != nil {
		log.Println(err)
	}
	return asyncResp, err
}

//GetCredentialsList Get all the credentials related to a user or a virtualserver
// credentialType can be lwapi.OPERATINGSYSTEM  = 0 or lwapi.CONTROL_PANEL = 1
func (lw LW) GetCredentialsList(id string, credentialType int) (CredentialsList, error) {
	var credType string
	credslist := CredentialsList{}
	switch credentialType {
	case 0:
		credType = "OPERATING_SYSTEM"
	case 1:
		credType = "CONTROL_PANEL"
	default:
		return credslist, errors.New("Invalid Credential Type")
	}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s/credentials/%s", virtualServersAPI, id, credType), nil)
	req.Header.Add("X-Lsw-Auth", lw.AuthToken)
	body, err := lw.sendHTTPRequest(req, http.StatusOK)
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(body, &credslist)
	if err != nil {
		log.Println(err)
	}
	return credslist, err
}

//GetCredentials Get a single user credentials
func (lw LW) GetCredentials(id string, credentialType int, username string) (Credentials, error) {
	credslist, err := lw.GetCredentialsList(id, credentialType)
	if err != nil {
		return Credentials{}, err
	}
	for _, creds := range credslist.Credentials {
		if creds.Username == username {
			return creds, err
		}
	}
	return Credentials{}, err
}

//GetTrafficDataMetrics Retrieve a list of all of your datatraffic metrics,
func (lw LW) GetTrafficDataMetrics(id, aggregation, from, to string) (Metrics, error) {
	metrics := Metrics{}
	rawurl := fmt.Sprintf("%s/%s/metrics/datatraffic", virtualServersAPI, id)
	url, err := url.Parse(rawurl)
	if err != nil {
		log.Println(err)
	}
	q := url.Query()
	q.Set("from", from)
	q.Set("to", to)
	q.Set("aggregation", aggregation)
	url.RawQuery = q.Encode()
	req, _ := http.NewRequest("GET", url.String(), nil)
	req.Header.Add("X-Lsw-Auth", lw.AuthToken)
	body, err := lw.sendHTTPRequest(req, http.StatusOK)
	if err != nil {
		return metrics, errors.New("Something went wrong: " + err.Error())
	}
	err = json.Unmarshal(body, &metrics)
	if err != nil {
		log.Println(err)
	}
	return metrics, err
}
func (lw LW) sendHTTPRequest(req *http.Request, ExpectedHTTPStatus int) ([]byte, error) {
	resp, err := lw.httpclient.Do(req)
	if err != nil {
		log.Println(err)
	}
	if resp.StatusCode != ExpectedHTTPStatus {
		return []byte(""), errors.New("Something went wrong: " + resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
