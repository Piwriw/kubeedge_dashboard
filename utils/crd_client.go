package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type CRDClient struct {
	Host  string
	Token string
	Port  string
}

func CreateCRDClient(token, host, port string) *CRDClient {
	return &CRDClient{Token: token, Host: host, Port: port}
}
func (cli *CRDClient) DeleteCR(crName, group, version, namespace, plural string) (bool, error) {
	//  忽略验证ssl
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}
	params := url.Values{}
	params.Set("pretty", "true")
	apiURL := fmt.Sprintf("https://%s:%s/apis/%s/%s/namespaces/%s/%s/%s", cli.Host, cli.Port, group, version, namespace, plural, crName)
	apiURL += "?" + params.Encode()

	req, _ := http.NewRequest("DELETE", apiURL, nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cli.Token))

	response, err := client.Do(req)
	if err != nil {
		return false, err
	}
	if response.StatusCode != 200 {
		return false, errors.New(response.Status)
	}
	return true, err
}

// PatchCR 更新CR
func (cli *CRDClient) PatchCR(cr interface{}, group, apiVersion, namespace, plural, crName string) ([]byte, error) {
	//  忽略验证ssl
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}
	params := url.Values{}
	params.Set("pretty", "true")
	apiURL := fmt.Sprintf("https://%s:%s/apis/%s/%s/namespaces/%s/%s/%s", cli.Host, cli.Port, group, apiVersion, namespace, plural, crName)
	apiURL += "?" + params.Encode()
	crbytes, err := json.Marshal(cr)
	if err != nil {
		return nil, err
	}
	crbody := bytes.NewBuffer(crbytes)

	req, _ := http.NewRequest("PATCH", apiURL, crbody)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cli.Token))

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	s, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return s, nil

}

// UpdateCR 更新CR
func (cli *CRDClient) UpdateCR(cr interface{}, group, version, namespace, plural, crName string) ([]byte, error) {
	//  忽略验证ssl
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}
	params := url.Values{}
	params.Set("pretty", "true")
	apiURL := fmt.Sprintf("https://%s:%s/apis/%s/%s/namespaces/%s/%s/%s", cli.Host, cli.Port, group, version, namespace, plural, crName)
	apiURL += "?" + params.Encode()
	crbytes, err := json.Marshal(cr)
	if err != nil {
		return nil, err
	}
	crbody := bytes.NewBuffer(crbytes)

	req, _ := http.NewRequest("PUT", apiURL, crbody)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cli.Token))

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	s, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return s, nil

}

func (cli *CRDClient) CreateCR(cr interface{}, group, version, namespace, plural string) ([]byte, error) {
	//  忽略验证ssl
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}
	params := url.Values{}
	params.Set("pretty", "true")
	//POST https://10.10.124.19:6443/apis/rules.kubeedge.io/v1/namespaces/default/ruleendpoints
	apiURL := fmt.Sprintf("https://%s:%s/apis/%s/%s/namespaces/%s/%s", cli.Host, cli.Port, group, version, namespace, plural)
	apiURL += "?" + params.Encode()
	crbytes, err := json.Marshal(cr)
	if err != nil {
		return nil, err
	}
	crio := bytes.NewBuffer(crbytes)

	req, _ := http.NewRequest("POST", apiURL, crio)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cli.Token))

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 && response.StatusCode != 201 {
		return nil, errors.New(response.Status)
	}
	s, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetCR 获取一个CR 对象
func (cli *CRDClient) GetCR(crName, group, version,
	namespace, plural string) ([]byte, error) {

	//  忽略验证ssl
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}
	params := url.Values{}
	params.Set("pretty", "true")
	apiURL := fmt.Sprintf("https://%s:%s/apis/%s/%s/namespaces/%s/%s/%s", cli.Host, cli.Port, group, version, namespace, plural, crName)
	apiURL += "?" + params.Encode()
	req, _ := http.NewRequest("GET", apiURL, nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cli.Token))

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}
	s, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetCRList 获取CR List
func (cli *CRDClient) GetCRList(group, version, namespace, plural string) ([]byte, error) {

	//  忽略验证ssl
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}
	params := url.Values{}
	params.Set("pretty", "true")
	apiURL := fmt.Sprintf("https://%s:%s/apis/%s/%s/namespaces/%s/%s", cli.Host, cli.Port, group, version, namespace, plural)
	apiURL += "?" + params.Encode()
	req, _ := http.NewRequest("GET", apiURL, nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cli.Token))

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 && response.StatusCode != 204 {
		return nil, errors.New(response.Status)
	}
	s, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return s, nil
}
