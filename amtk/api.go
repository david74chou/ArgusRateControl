package amtk

import (
	"fmt"
	"strconv"

	"gopkg.in/resty.v0"
)

const AMTK_CGI = "/param.cgi"

type AMTKAPIParams struct {
	APIHostURL  string
	APIUser     string
	APIPassword string
}

type amtkAPI struct {
	client *resty.Client
}

func New(params *AMTKAPIParams) *amtkAPI {
	api := &amtkAPI{
		client: resty.New(),
	}
	api.client.SetHostURL(fmt.Sprintf("http://%s", params.APIHostURL))
	api.client.SetBasicAuth(params.APIUser, params.APIPassword)
	api.client.SetDisableWarn(true)
	return api
}

func (api *amtkAPI) SetRateControl(mode string, targetBitrate, maxTargetBitrate int) error {
	response, err := api.client.R().
		SetQueryParam("action", "update").
		SetQueryParam("Image.I0.RateControl2.Mode", mode).
		SetQueryParam("Image.I0.RateControl2.TargetBitrate", strconv.Itoa(targetBitrate)).
		SetQueryParam("Image.I0.RateControl2.MaxTargetBitrate", strconv.Itoa(maxTargetBitrate)).
		Get(AMTK_CGI)
	if err != nil {
		return err
	}

	if response.String() != "OK" {
		return fmt.Errorf("amtkapi: fail to set rate control params")
	}

	return nil
}

func (api *amtkAPI) SetCompression(compression int) error {
	response, err := api.client.R().
		SetQueryParam("action", "update").
		SetQueryParam("Image.I0.Appearance2.Compression", strconv.Itoa(compression)).
		Get(AMTK_CGI)
	if err != nil {
		return err
	}

	if response.String() != "OK" {
		return fmt.Errorf("amtkapi: fail to set compression params")
	}

	return nil
}

func (api *amtkAPI) SetResolution(resolution string) error {
	response, err := api.client.R().
		SetQueryParam("action", "update").
		SetQueryParam("Image.I0.Appearance2.Resolution", resolution).
		Get(AMTK_CGI)
	if err != nil {
		return err
	}

	if response.String() != "OK" {
		return fmt.Errorf("amtkapi: fail to set resolution params")
	}

	return nil
}
