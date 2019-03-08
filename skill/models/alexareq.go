package models

type AlexaSimpleRequest struct {
	Version string `json:"version"`
	Request struct {
		Type   string `json:"type"`
		Time   string `json:"timestamp"`
		Intent struct {
			Name               string `json:"name"`
			ConfirmationStatus string `json:"confirmationstatus"`
		} `json:"intent"`
	} `json:"request"`
}

type AlexaComplexRequest struct {
	Version string `json:"version"`
	Request struct {
		Type   string `json:"type"`
		Time   string `json:"timestamp"`
		Intent struct {
			Name               string `json:"name"`
			ConfirmationStatus string `json:"confirmationstatus"`
			Slots              struct {
				Name struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"name"`
				Type struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"type"`
				Quantity struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"quantity"`
				QuantityRequested struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"quantityRequested"`
				Wgt struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"wgt"`
				Kind struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"kind"`
				LogType struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"logType"`
				Duration struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"durationRequest"`
				TimeUnit struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"timeUnit"`
			} `json:"slots"`
		} `json:"intent"`
	} `json:"request"`
}

type AlexaResponse struct {
	Version  string `json:"version"`
	Response struct {
		OutputSpeech struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"outputSpeech"`
	} `json:"response"`
}

func CreateResponse() *AlexaResponse {
	var resp AlexaResponse
	resp.Version = "1.0"
	resp.Response.OutputSpeech.Type = "PlainText"
	resp.Response.OutputSpeech.Text = "Hello.  Please override this default output."
	return &resp
}

func (resp *AlexaResponse) Say(text string) {
	resp.Response.OutputSpeech.Text = text
}
