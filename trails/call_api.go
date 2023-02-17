package trails

import (
	"bytes"
	"encoding/json"
	"errors"
	"go_starter/logs"
	"io/ioutil"
	"net/http"
)

func CallApi(queryUrl string, request interface{}) ([]byte, error) {
	httpClient := http.Client{}
	marshal, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	response, err := http.NewRequest("POST", queryUrl, bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}
	response.Header.Add("Content-Type", "application/json")
	res, err := httpClient.Do(response)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	defer res.Body.Close()
	readAll, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if res.StatusCode == 200 {

		return readAll, nil
	}
	catchErr := map[string]interface{}{}
	err = json.Unmarshal([]byte(readAll), &catchErr)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	logs.Error(catchErr["error"].(string))
	return nil, errors.New(catchErr["error"].(string))

}

func CallApiIpro(queryUrl string, request interface{}) ([]byte, error) {
	httpClient := http.Client{}
	marshal, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	response, err := http.NewRequest("POST", queryUrl, bytes.NewBuffer(marshal))
	if err != nil {
		return nil, err
	}
	response.Header.Add("Content-Type", "application/json")
	res, err := httpClient.Do(response)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	defer res.Body.Close()
	readAll, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	if res.StatusCode == 200 {
		decodeError := map[string]interface{}{}
		err = json.Unmarshal([]byte(readAll), &decodeError)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		responseStatusCode := decodeError["status"].(string)
		responseDescription := decodeError["description"].(string)
		if responseStatusCode != "1" {
			logs.Error(errors.New(responseDescription))
			return nil, errors.New(responseDescription)
		}

		return readAll, nil
	}

	catchErr := map[string]interface{}{}
	err = json.Unmarshal([]byte(readAll), &catchErr)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	logs.Error(catchErr["error"].(string))
	return nil, errors.New(catchErr["error"].(string))

}
