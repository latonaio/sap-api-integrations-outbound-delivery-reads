package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-outbound-delivery-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetOutboundDelivery(deliveryDocument, sDDocument, partnerFunction, deliveryDocumentItem string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "Header":
			func() {
				c.Header(deliveryDocument)
				wg.Done()
			}()
		case "HeaderPartner":
			func() {
				c.HeaderPartner(sDDocument, partnerFunction)
				wg.Done()
			}()
		case "PartnerAddress":
			func() {
				c.PartnerAddress(partnerFunction, sDDocument)
				wg.Done()
			}()
		case "Item":
			func() {
				c.Item(deliveryDocument, deliveryDocumentItem)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) Header(deliveryDocument string) {
	headerData, err := c.callOutboundDeliverySrvAPIRequirementHeader("A_OutbDeliveryHeader", deliveryDocument)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(headerData)
	}

	headerPartnerData, err := c.callToHeaderPartner(headerData[0].ToHeaderPartner)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(headerPartnerData)
	}
	
	partnerAddressData, err := c.callToPartnerAddress(headerPartnerData[0].ToPartnerAddress)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(partnerAddressData)
	}
	
	itemData, err := c.callToItem(headerData[0].ToItem)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemData)
	}
	
	itemDocumentFlowData, err := c.callToItemDocumentFlow(itemData[0].ToItemDocumentFlow)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemDocumentFlowData)
	}
	return
}

func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementHeader(api, deliveryDocument string) ([]sap_api_output_formatter.Header, error) {
	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV;v=0002", api}, "/")
	param := c.getQueryWithHeader(map[string]string{}, deliveryDocument)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeader(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToHeaderPartner(url string) ([]sap_api_output_formatter.ToHeaderPartner, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToHeaderPartner(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToPartnerAddress(url string) (*sap_api_output_formatter.ToPartnerAddress, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPartnerAddress(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItem(url string) ([]sap_api_output_formatter.ToItem, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemDocumentFlow(url string) ([]sap_api_output_formatter.ToItemDocumentFlow, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemDocumentFlow(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) HeaderPartner(sDDocument, partnerFunction string) {
	headerPartnerData, err := c.callOutboundDeliverySrvAPIRequirementHeaderPartner("A_OutbDeliveryHeader('%s')/to_DeliveryDocumentPartner", sDDocument, partnerFunction)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(headerPartnerData)
	}

	partnerAddressData, err := c.callToPartnerAddress2(headerPartnerData[0].ToPartnerAddress)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(partnerAddressData)
	}
	return	
}


func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementHeaderPartner(api, sDDocument, partnerFunction string) ([]sap_api_output_formatter.HeaderPartner, error) {
	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV;v=0002", api}, "/")
	param := c.getQueryWithHeaderPartner(map[string]string{}, sDDocument, partnerFunction)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToHeaderPartner(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}


func (c *SAPAPICaller) callToPartnerAddress2(url string) (*sap_api_output_formatter.ToPartnerAddress, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToPartnerAddress(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}


func (c *SAPAPICaller) PartnerAddress(partnerFunction, sDDocument string) {
	data, err := c.callOutboundDeliverySrvAPIRequirementPartnerAddress("A_OutbDeliveryPartner(PartnerFunction='%s',SDDocument='%s')/to_Address2", partnerFunction, sDDocument)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementPartnerAddress(api, partnerFunction, sDDocument string) (*sap_api_output_formatter.PartnerAddress, error) {
	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV;v=0002", api}, "/")
	param := c.getQueryWithPartnerAddress(map[string]string{}, partnerFunction, sDDocument)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToPartnerAddress(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) Item(deliveryDocument, deliveryDocumentItem string) {
	itemData, err := c.callOutboundDeliverySrvAPIRequirementItem("A_OutbDeliveryItem", deliveryDocument, deliveryDocumentItem)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemData)

	itemDocumentFlowData, err := c.callToItemDocumentFlow2(itemData[0].ToItemDocumentFlow)
	if err != nil {
		c.log.Error(err)
	} else {
		c.log.Info(itemDocumentFlowData)
	}
	return	
}


func (c *SAPAPICaller) callOutboundDeliverySrvAPIRequirementItem(api, deliveryDocument, deliveryDocumentItem string) ([]sap_api_output_formatter.Item, error) {
	url := strings.Join([]string{c.baseURL, "API_OUTBOUND_DELIVERY_SRV;v=0002", api}, "/")
	param := c.getQueryWithItem(map[string]string{}, deliveryDocument, deliveryDocumentItem)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToItem(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToItemDocumentFlow2(url string) ([]sap_api_output_formatter.ToItemDocumentFlow, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToItemDocumentFlow(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}


func (c *SAPAPICaller) getQueryWithHeader(params map[string]string, deliveryDocument string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("DeliveryDocument eq '%s'", deliveryDocument)
	return params
}

func (c *SAPAPICaller) getQueryWithHeaderPartner(params map[string]string, sDDocument, partnerFunction string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("SDDocument eq '%s' and PartnerFunction eq '%s'", sDDocument, partnerFunction)
	return params
}

func (c *SAPAPICaller) getQueryWithPartnerAddress(params map[string]string, partnerFunction, sDDocument string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("PartnerFunction eq '%s' and SDDocument eq '%s'", partnerFunction, sDDocument)
	return params
}

func (c *SAPAPICaller) getQueryWithItem(params map[string]string, deliveryDocument, deliveryDocumentItem string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("DeliveryDocument eq '%s' and DeliveryDocumentItem eq '%s'", deliveryDocument, deliveryDocumentItem)
	return params
}
