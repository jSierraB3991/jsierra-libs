package eliotlibs

import (
	"log"
	"testing"
)

type EpaycoRequestConfirmation struct {
	Signature     string `query:"x_signature"`
	CustIDCliente string `query:"x_cust_id_cliente"`
	RefPayco      string `query:"x_ref_payco"`
}

func TestXxx(t *testing.T) {
	var requ EpaycoRequestConfirmation
	iterator(&requ, "x_signature=sina&x_cust_id_cliente=clien_1&x_ref_payco=invoice")
	log.Println("final")
	log.Println(requ)
}
