package lwapi

import (
	"net/http"
	"testing"

	"gopkg.in/jarcoal/httpmock.v1"
)

func TestGetVirtualServersList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", virtualServersAPI,
		httpmock.NewStringResponder(http.StatusOK, `{"virtualServers":[{"id":"222903","reference":"Web server","customerId":"1301178860","dataCenter":"AMS-01","cloudServerId":null,"state":"STOPPED","firewallState":"DISABLED","template":"Ubuntu 14.04 64 40 20140707T1340","serviceOffering":"S","sla":"Bronze","contract":{"id":"30000778","startsAt":"2016-02-01T00:00:00+0200","endsAt":"2017-01-31T00:00:00+0200","billingCycle":12,"billingFrequency":"MONTH","pricePerFrequency":4.7,"currency":"EUR"},"hardware":{"cpu":{"cores":1},"memory":{"unit":"MB","amount":1024},"storage":{"unit":"GB","amount":40}},"iso":null,"ips":[{"ip":"10.11.116.130","version":"4","type":"PUBLIC"}]},{"id":"301708","reference":null,"customerId":"1301178860","dataCenter":"AMS-01","cloudServerId":null,"state":"STOPPED","firewallState":"ENABLED","template":"CentOS 7.0 64 60 20140711T1039","serviceOffering":"M","sla":"Bronze","contract":{"id":"30000779","startsAt":"2016-02-01T00:00:00+0200","endsAt":"2017-01-31T00:00:00+0200","billingCycle":12,"billingFrequency":"MONTH","pricePerFrequency":4.7,"currency":"EUR"},"hardware":{"cpu":{"cores":2},"memory":{"unit":"MB","amount":2048},"storage":{"unit":"GB","amount":60}},"iso":{"id":"9eadbe14-69be-4dee-8f56-5ebb23bb3c33","name":"Knoppix","displayName":"Knoppix"},"ips":[{"ip":"10.11.116.132","version":"4","type":"PUBLIC"}]}],"_metadata":{"totalCount":2,"offset":0,"limit":10}}`))
	lw := LW{}
	_, err := lw.GetVirtualServersList()
	checkErr(t, err)

}

func TestGetVirtualServer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", virtualServersAPI+"/222903",
		httpmock.NewStringResponder(http.StatusOK, `{"id":"222903","reference":"Web server","customerId":"1301178860","dataCenter":"AMS-01","cloudServerId":null,"state":"STOPPED","firewallState":"DISABLED","template":"Ubuntu 14.04 64 40 20140707T1340","serviceOffering":"S","sla":"Bronze","contract":{"id":"30000778","startsAt":"2016-02-01T00:00:00+0200","endsAt":"2017-01-31T00:00:00+0200","billingCycle":12,"billingFrequency":"MONTH","pricePerFrequency":4.7,"currency":"EUR"},"hardware":{"cpu":{"cores":1},"memory":{"unit":"MB","amount":1024},"storage":{"unit":"GB","amount":40}},"iso":{"id":"9eadbe14-69be-4dee-8f56-5ebb23bb3c33","name":"Knoppix","displayName":"Knoppix"},"ips":[{"ip":"10.11.116.130","version":"4","type":"PUBLIC"}]}`))
	lw := LW{}
	_, err := lw.GetVirtualServer("222903")
	checkErr(t, err)

	httpmock.RegisterResponder("GET", virtualServersAPI+"/222903222903222903",
		httpmock.NewStringResponder(http.StatusOK, `{"errorCode":404,"errorMessage":"Resource '222903222903222903' of type 'VIRTUAL_SERVER' was not found","reference":"https://www.leaseweb.com/contact","userMessage":"The requested resource was not found"}`))
	_, err = lw.GetVirtualServer("222903222903222903")

}

func TestUpdateServerReference(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("PUT", virtualServersAPI+"/222903",
		httpmock.NewStringResponder(http.StatusOK, `{"id":"222903","reference":"Caching Server","customerId":"1301178860","dataCenter":"AMS-01","cloudServerId":null,"state":"STOPPED","firewallState":"DISABLED","template":"Ubuntu 14.04 64 40 20140707T1340","serviceOffering":"S","sla":"Bronze","contract":{"id":"30000778","startsAt":"2016-02-01T00:00:00+0200","endsAt":"2017-01-31T00:00:00+0200","billingCycle":12,"billingFrequency":"MONTH","pricePerFrequency":4.7,"currency":"EUR"},"hardware":{"cpu":{"cores":1},"memory":{"unit":"MB","amount":1024},"storage":{"unit":"GB","amount":40}},"iso":{"id":"9eadbe14-69be-4dee-8f56-5ebb23bb3c33","name":"Knoppix","displayName":"Knoppix"},"ips":[{"ip":"10.11.116.130","version":"4","type":"PUBLIC"}]}`))
	lw := LW{}
	_, err := lw.UpdateServerReference("222903", "Caching Server")
	checkErr(t, err)

}

func TestPowerControlPowerON(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", virtualServersAPI+"/222903/powerOn",
		httpmock.NewStringResponder(http.StatusAccepted, `{"id":"cs01.237daad0-2aed-4260-b0e4-488d9cd55607","name":"virtualServers.powerOn","status":"PENDING","createdAt":"2016-12-31T01:00:59+00:00"}`))
	lw := LW{}
	_, err := lw.PowerControl("222903", POWERON)
	checkErr(t, err)

}

func TestPowerControlPowerOff(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", virtualServersAPI+"/222903/powerOff",
		httpmock.NewStringResponder(http.StatusAccepted, `{"id":"cs01.237daad0-2aed-4260-b0e4-488d9cd55607","name":"virtualServers.powerOff","status":"PENDING","createdAt":"2016-12-31T01:00:59+00:00"}`))
	lw := LW{}
	_, err := lw.PowerControl("222903", POWEROFF)
	checkErr(t, err)
}

func TestPowerControlReboot(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", virtualServersAPI+"/222903/reboot",
		httpmock.NewStringResponder(http.StatusAccepted, `{"id":"cs01.237daad0-2aed-4260-b0e4-488d9cd55607","name":"virtualServers.reboot","status":"PENDING","createdAt":"2016-12-31T01:00:59+00:00"}`))
	lw := LW{}
	_, err := lw.PowerControl("222903", REBOOT)
	checkErr(t, err)
}

func TestReinstallVirtualServer(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", virtualServersAPI+"/222903/reinstall",
		httpmock.NewStringResponder(http.StatusAccepted, `{"id":"cs01.237daad0-2aed-4260-b0e4-488d9cd55607","name":"virtualServers.reinstall","status":"PENDING","createdAt":"2016-12-31T01:00:59+00:00"}`))
	lw := LW{}
	_, err := lw.ReinstallVirtualServer("222903")
	checkErr(t, err)

}
func TestGetCredentialsList(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", virtualServersAPI+"/222903/credentials/OPERATING_SYSTEM",
		httpmock.NewStringResponder(http.StatusOK, `{"_metadata":{"limit":1,"offset":0,"totalCount":1},"credentials":[{"domain":null,"password":null,"type":"OPERATING_SYSTEM","username":"root"}]}`))
	lw := LW{}
	_, err := lw.GetCredentialsList("222903", OPERATINGSYSTEM)
	checkErr(t, err)

}

func TestGetCredentials(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", virtualServersAPI+"/222903/credentials/OPERATING_SYSTEM",
		httpmock.NewStringResponder(http.StatusOK, `{"domain":null,"password":"PASSWORD","type":"OPERATING_SYSTEM","username":"root"}`))
	lw := LW{}
	_, err := lw.GetCredentials("222903", OPERATINGSYSTEM, "root")
	checkErr(t, err)
}

func TestGetTrafficDataMetrics(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", virtualServersAPI+"/222903/metrics/datatraffic?aggregation=SUM&from=2017-07-01T00%3A00%3A00%2B00%3A00&to=2016-01-31T23%3A59%3A59%2B00%3A00",
		httpmock.NewStringResponder(http.StatusOK, `{"_metadata":{"from":"2016-01-01T00:00:00+00:00","to":"2016-01-31T23:59:59+00:00","granularity":"DAY","aggregation":"SUM"},"metrics":{"DATATRAFFIC_UP":{"unit":"B","values":[{"timestamp":"2016-01-01T23:59:59+00:00","value":900},{"timestamp":"2016-01-31T23:59:59+00:00","value":2500}]},"DATATRAFFIC_DOWN":{"unit":"B","values":[{"timestamp":"2016-01-01T23:59:59+00:00","value":90},{"timestamp":"2016-01-31T23:59:59+00:00","value":250}]}}}`))
	lw := LW{}
	_, err := lw.GetTrafficDataMetrics("222903", "SUM", "2017-07-01T00:00:00+00:00", "2016-01-31T23:59:59+00:00")
	checkErr(t, err)
}
func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
