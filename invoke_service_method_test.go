package enova_test

import (
	"encoding/json"
	"fmt"
	"testing"

	enova "github.com/omniboost/go-enova"
)

func TestInvokeServiceMethod(t *testing.T) {
	req := client.NewInvokeServiceMethodRequest()
	req.RequestBody().InvokerParams.DatabaseHandle = enova.DatabaseHandle(client.DBName())
	req.RequestBody().InvokerParams.MethodName = enova.MethodName("Update")
	req.RequestBody().InvokerParams.Operator = enova.Operator(client.DBUsername())
	req.RequestBody().InvokerParams.Password = enova.Password(client.DBPassword())
	req.RequestBody().InvokerParams.MethodArgs = enova.MethodArgs{
		"TableName":      client.TableName(),
		"SchemaName":     client.SchemaName(),
		"ExternalSystem": client.ExternalSystemGUID(),
		// "Data": enova.UpdateParamsData{
		// 	Rows: enova.Rows{enova.Row{XML: "FML"}},
		// },
		"Data": `
              <UpdateParamsData>
                <Rows>
                  <Row>
                    <Xml>
                      <Dokument>
                        <invoice_fiscalcode>112</invoice_fiscalcode>
                        <Podmiot>
                          <invoice_customername>Smith &amp;amp;amp; Nephew Sp. z o.o.</invoice_customername>
                          <invoice_customertaxid>526 28 25 537</invoice_customertaxid>
                          <invoice_customeraddress>ul. Osma&amp;#x144;ska 12 02-823 Warszawa</invoice_customeraddress>
                          <invoice_protel_customerid>60024</invoice_protel_customerid>
                        </Podmiot>
                        <Naglowek>
                          <invoice_protel_reservationno>1103664</invoice_protel_reservationno>
                          <invoice_cancelled>True</invoice_cancelled>
                          <invoice_cancell_dt>2021-06-30</invoice_cancell_dt>
                          <invoice_type>1</invoice_type>
                          <invoice_nostr>25103/05/2021</invoice_nostr>
                          <invoice_saledate>2021-05-18</invoice_saledate>
                          <invoice_issuedate>2021-05-19</invoice_issuedate>
                          <invoice_id>691442</invoice_id>
                          <invoice_note1>Chmarek Jacek</invoice_note1>
                          <invoice_note2>Chmarek Jacek1</invoice_note2>
                          <invoice_note3/>
                          <invoice_property_name>Puro Hotel Warszawa</invoice_property_name>
                          <invoice_propertyid>9</invoice_propertyid>
                          <invoice_gtu/>
                          <invoice_dfeature/>
                          <invoice_protel_datefrom>2021-05-18</invoice_protel_datefrom>
                          <invoice_protel_dateto>2021-05-19</invoice_protel_dateto>
                        </Naglowek>
                        <LinkedDocument>
                          <CorrectedDocuments>
                            <CorrectedDocument>
                              <invoice_originalinvoicenostr>39943/02/2021</invoice_originalinvoicenostr>
                            </CorrectedDocument>
                          </CorrectedDocuments>
                        </LinkedDocument>
                        <VAT>
                          <VATTable>
                            <VATTableItem>
                              <VATRate>
                                <invoicevat_vatcode>B</invoicevat_vatcode>
                                <invoicevat_vatrate>8,00</invoicevat_vatrate>
                              </VATRate>
                              <invoicevat_vatamount>27,41</invoicevat_vatamount>
                              <invoicevat_grossamount>370,00</invoicevat_grossamount>
                              <invoicedtl_currency>PLN</invoicedtl_currency>
                              <invoicevat_vatcurrencyamount>27,41</invoicevat_vatcurrencyamount>
                              <invoicedtl_grosscurrencyamount>370,00</invoicedtl_grosscurrencyamount>
                              <invoicevat_id>961519</invoicevat_id>
                              <invoicevat_invoiceid>691442</invoicevat_invoiceid>
                              <invoicevat_itemtype>2</invoicevat_itemtype>
                              <invoicevat_vatratetxt>8.00</invoicevat_vatratetxt>
                            </VATTableItem>
                            <VATTableItem>
                              <VATRate>
                                <invoicevat_vatcode>A</invoicevat_vatcode>
                                <invoicevat_vatrate>23,00</invoicevat_vatrate>
                              </VATRate>
                              <invoicevat_vatamount>18,70</invoicevat_vatamount>
                              <invoicevat_grossamount>100,00</invoicevat_grossamount>
                              <invoicedtl_currency>PLN</invoicedtl_currency>
                              <invoicevat_vatcurrencyamount>18,70</invoicevat_vatcurrencyamount>
                              <invoicedtl_grosscurrencyamount>100,00</invoicedtl_grosscurrencyamount>
                              <invoicevat_id>961520</invoicevat_id>
                              <invoicevat_invoiceid>691442</invoicevat_invoiceid>
                              <invoicevat_itemtype>2</invoicevat_itemtype>
                              <invoicevat_vatratetxt>23.00</invoicevat_vatratetxt>
                            </VATTableItem>
                          </VATTable>
                        </VAT>
                        <payments>
                          <Payment>
                            <invoicedtl_currency>PLN</invoicedtl_currency>
                            <invoicedtl_grosscurrencyamount>470</invoicedtl_grosscurrencyamount>
                            <invoice_duedate>2021-05-19</invoice_duedate>
                            <invoicedtl_postingtext>Karta kredytowa</invoicedtl_postingtext>
                            <invoicedtl_postdate>2021-05-19</invoicedtl_postdate>
                            <invoice_cc/>
                            <invoicedtl_realdate>2021-05-19</invoicedtl_realdate>
                            <invoicedtl_reservationid>1103664</invoicedtl_reservationid>
                            <invoicedtl_id>4672384</invoicedtl_id>
                            <invoicedtl_deposit>false</invoicedtl_deposit>
                            <invoicedtl_payment>true</invoicedtl_payment>
                            <invoicedtl_invoiceid>691442</invoicedtl_invoiceid>
                          </Payment>
                        </payments>
                        <OpisAnalityczny>
                          <ElementOpisu>
                            <invoicedtl_vatamount>18,7</invoicedtl_vatamount>
                            <invoicedtl_currency>PLN</invoicedtl_currency>
                            <invoicedtl_grosscurrencyamount>100</invoicedtl_grosscurrencyamount>
                            <invoicedtl_grossamount>100</invoicedtl_grossamount>
                            <invoicedtl_originid>1103664</invoicedtl_originid>
                            <invoicedtl_paymentaccount/>
                            <invoicedtl_realdate>2021-05-18</invoicedtl_realdate>
                            <invoicedtl_postingaddtext/>
                            <invoicedtl_postingtext>Parking (Parking)</invoicedtl_postingtext>
                            <invoicedtl_protel_paymentname/>
                            <invoicedtl_protel_transname>Parking (Parking)</invoicedtl_protel_transname>
                            <invoicedtl_quantity>1</invoicedtl_quantity>
                            <invoicedtl_reservationid>1103664</invoicedtl_reservationid>
                            <invoicedtl_transaccount/>
                            <invoicedtl_vatcode>A</invoicedtl_vatcode>
                            <invoicedtl_vatcurrencyamount>18,7</invoicedtl_vatcurrencyamount>
                            <invoicedtl_id>4672381</invoicedtl_id>
                            <invoicedtl_sourcetransactionid>6</invoicedtl_sourcetransactionid>
                            <invoicedtl_postdate>2021-05-18</invoicedtl_postdate>
                            <invoicedtl_deposit>false</invoicedtl_deposit>
                            <invoicedtl_payment>false</invoicedtl_payment>
                            <invoicedtl_invoiceid>691442</invoicedtl_invoiceid>
                          </ElementOpisu>
                          <ElementOpisu>
                            <invoicedtl_vatamount>2,96</invoicedtl_vatamount>
                            <invoicedtl_currency>PLN</invoicedtl_currency>
                            <invoicedtl_grosscurrencyamount>40</invoicedtl_grosscurrencyamount>
                            <invoicedtl_grossamount>40</invoicedtl_grossamount>
                            <invoicedtl_originid>1103664</invoicedtl_originid>
                            <invoicedtl_paymentaccount/>
                            <invoicedtl_realdate>2021-05-18</invoicedtl_realdate>
                            <invoicedtl_postingaddtext>18/05/21/#418</invoicedtl_postingaddtext>
                            <invoicedtl_postingtext>Us&amp;#x142;uga noclegowa (Accommodation)</invoicedtl_postingtext>
                            <invoicedtl_protel_paymentname/>
                            <invoicedtl_protel_transname>&amp;#x15A;niadanie (Breakfast)</invoicedtl_protel_transname>
                            <invoicedtl_quantity>1</invoicedtl_quantity>
                            <invoicedtl_reservationid>1103664</invoicedtl_reservationid>
                            <invoicedtl_transaccount/>
                            <invoicedtl_vatcode>B</invoicedtl_vatcode>
                            <invoicedtl_vatcurrencyamount>2,96</invoicedtl_vatcurrencyamount>
                            <invoicedtl_id>4672382</invoicedtl_id>
                            <invoicedtl_sourcetransactionid>132</invoicedtl_sourcetransactionid>
                            <invoicedtl_postdate>2021-05-18</invoicedtl_postdate>
                            <invoicedtl_deposit>false</invoicedtl_deposit>
                            <invoicedtl_payment>false</invoicedtl_payment>
                            <invoicedtl_invoiceid>691442</invoicedtl_invoiceid>
                          </ElementOpisu>
                          <ElementOpisu>
                            <invoicedtl_vatamount>24,45</invoicedtl_vatamount>
                            <invoicedtl_currency>PLN</invoicedtl_currency>
                            <invoicedtl_grosscurrencyamount>330</invoicedtl_grosscurrencyamount>
                            <invoicedtl_grossamount>330</invoicedtl_grossamount>
                            <invoicedtl_originid>1103664</invoicedtl_originid>
                            <invoicedtl_paymentaccount/>
                            <invoicedtl_realdate>2021-05-18</invoicedtl_realdate>
                            <invoicedtl_postingaddtext>18/05/21/#418</invoicedtl_postingaddtext>
                            <invoicedtl_postingtext>Us&amp;#x142;uga noclegowa (Accommodation)</invoicedtl_postingtext>
                            <invoicedtl_protel_paymentname/>
                            <invoicedtl_protel_transname>Us&amp;#x142;uga noclegowa (Accommodation)</invoicedtl_protel_transname>
                            <invoicedtl_quantity>1</invoicedtl_quantity>
                            <invoicedtl_reservationid>1103664</invoicedtl_reservationid>
                            <invoicedtl_transaccount/>
                            <invoicedtl_vatcode>B</invoicedtl_vatcode>
                            <invoicedtl_vatcurrencyamount>24,44</invoicedtl_vatcurrencyamount>
                            <invoicedtl_id>4672383</invoicedtl_id>
                            <invoicedtl_sourcetransactionid>1</invoicedtl_sourcetransactionid>
                            <invoicedtl_postdate>2021-05-18</invoicedtl_postdate>
                            <invoicedtl_deposit>false</invoicedtl_deposit>
                            <invoicedtl_payment>false</invoicedtl_payment>
                            <invoicedtl_invoiceid>691442</invoicedtl_invoiceid>
                          </ElementOpisu>
                        </OpisAnalityczny>
                        <Attachments>
                          <Attachment>
                            <invoice_no>25103</invoice_no>
                            <pdf_link>Hotel 9\112\00025103.pdf</pdf_link>
                          </Attachment>
                        </Attachments>
                      </Dokument>
                    </Xml>
                  </Row>
                </Rows>
              </UpdateParamsData>`,
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
