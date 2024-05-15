package jsonstore

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/precompile"
	"github.com/ethereum/go-ethereum/precompile/mocks"
)

const (
	ShortJson     = `{"forma":true}`
	LongJson      = `{"menu":{"id":"file","value":"File","popup":{"menuitem":[{"value":"New","onclick":"CreateNewDoc()"},{"value":"Open","onclick":"OpenDoc()"},{"value":"Close","onclick":"CloseDoc()"}]}}}`
	SuperLongJson = `{"web-app":{"servlet":[{"servlet-name":"cofaxCDS","servlet-class":"org.cofax.cds.CDSServlet","init-param":{"configGlossary:installationAt":"Philadelphia, PA","configGlossary:adminEmail":"ksm@pobox.com","configGlossary:poweredBy":"Cofax","configGlossary:poweredByIcon":"/images/cofax.gif","configGlossary:staticPath":"/content/static","templateProcessorClass":"org.cofax.WysiwygTemplate","templateLoaderClass":"org.cofax.FilesTemplateLoader","templatePath":"templates","templateOverridePath":"","defaultListTemplate":"listTemplate.htm","defaultFileTemplate":"articleTemplate.htm","useJSP":false,"jspListTemplate":"listTemplate.jsp","jspFileTemplate":"articleTemplate.jsp","cachePackageTagsTrack":200,"cachePackageTagsStore":200,"cachePackageTagsRefresh":60,"cacheTemplatesTrack":100,"cacheTemplatesStore":50,"cacheTemplatesRefresh":15,"cachePagesTrack":200,"cachePagesStore":100,"cachePagesRefresh":10,"cachePagesDirtyRead":10,"searchEngineListTemplate":"forSearchEnginesList.htm","searchEngineFileTemplate":"forSearchEngines.htm","searchEngineRobotsDb":"WEB-INF/robots.db","useDataStore":true,"dataStoreClass":"org.cofax.SqlDataStore","redirectionClass":"org.cofax.SqlRedirection","dataStoreName":"cofax","dataStoreDriver":"com.microsoft.jdbc.sqlserver.SQLServerDriver","dataStoreUrl":"jdbc:microsoft:sqlserver://LOCALHOST:1433;DatabaseName=goon","dataStoreUser":"sa","dataStorePassword":"dataStoreTestQuery","dataStoreTestQuery":"SET NOCOUNT ON;select test='test';","dataStoreLogFile":"/usr/local/tomcat/logs/datastore.log","dataStoreInitConns":10,"dataStoreMaxConns":100,"dataStoreConnUsageLimit":100,"dataStoreLogLevel":"debug","maxUrlLength":500}},{"servlet-name":"cofaxEmail","servlet-class":"org.cofax.cds.EmailServlet","init-param":{"mailHost":"mail1","mailHostOverride":"mail2"}},{"servlet-name":"cofaxAdmin","servlet-class":"org.cofax.cds.AdminServlet"},{"servlet-name":"fileServlet","servlet-class":"org.cofax.cds.FileServlet"},{"servlet-name":"cofaxTools","servlet-class":"org.cofax.cms.CofaxToolsServlet","init-param":{"templatePath":"toolstemplates/","log":1,"logLocation":"/usr/local/tomcat/logs/CofaxTools.log","logMaxSize":"","dataLog":1,"dataLogLocation":"/usr/local/tomcat/logs/dataLog.log","dataLogMaxSize":"","removePageCache":"/content/admin/remove?cache=pages&id=","removeTemplateCache":"/content/admin/remove?cache=templates&id=","fileTransferFolder":"/usr/local/tomcat/webapps/content/fileTransferFolder","lookInContext":1,"adminGroupID":4,"betaServer":true}}],"servlet-mapping":{"cofaxCDS":"/","cofaxEmail":"/cofaxutil/aemail/*","cofaxAdmin":"/admin/*","fileServlet":"/static/*","cofaxTools":"/tools/*"},"taglib":{"taglib-uri":"cofax.tld","taglib-location":"/WEB-INF/tlds/cofax.tld"}}}`
)

var (
	zero = big.NewInt(0)

	Slot0 = common.HexToHash("0x00")
	Slot1 = common.HexToHash("0x01")
	Slot2 = common.HexToHash("0x02")

	SelfAddrForTest  = common.HexToAddress("0x1000")
	MsgSenderForTest = common.BytesToAddress([]byte("0xMsgSender"))
)

func TestExtendedStorage(t *testing.T) {
	stateDB := mocks.NewMockStateDB()
	jsonstore := NewJsonStore()

	ctx := precompile.NewStatefulContext(stateDB, SelfAddrForTest, MsgSenderForTest, zero)

	_, err := jsonstore.Set(ctx, Slot0, SuperLongJson)
	if err != nil {
		t.Fatalf("Set slot0 failed with error: %v", err)
	}

	res, err := jsonstore.Get(ctx, MsgSenderForTest, Slot0)
	if err != nil {
		t.Fatalf("Get slot0 failed with error: %v", err)
	}

	if res != SuperLongJson {
		t.Fatalf("Get slot0 failed to return expected result.\n  Expected %v\n  got %v", SuperLongJson, res)
	}

	_, err = jsonstore.Set(ctx, Slot1, LongJson)
	if err != nil {
		t.Fatalf("Set slot1 failed with error: %v", err)
	}

	res, err = jsonstore.Get(ctx, MsgSenderForTest, Slot1)
	if err != nil {
		t.Fatalf("Get slot1 failed with error: %v", err)
	}

	if res != LongJson {
		t.Fatalf("Get slot1 failed to return expected result.\n  Expected %v\n  got %v", LongJson, res)
	}

	_, err = jsonstore.Set(ctx, Slot2, ShortJson)
	if err != nil {
		t.Fatalf("Set slot2 failed with error: %v", err)
	}

	res, err = jsonstore.Get(ctx, MsgSenderForTest, Slot2)
	if err != nil {
		t.Fatalf("Get slot2 failed with error: %v", err)
	}

	if res != ShortJson {
		t.Fatalf("Get slot2 failed to return expected result.\n  Expected %v\n  got %v", ShortJson, res)
	}
}
