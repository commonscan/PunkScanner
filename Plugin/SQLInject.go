package Plugin

import (
	"net/url"
	"strings"
	"github.com/valyala/fasthttp"
	"regexp"
	"fmt"
)

type SQLInjectionPlugin struct {
}

func (sqlinj SQLInjectionPlugin) GenPayload(url url.URL) string {
	var q = url.Query()
	for _, i := range strings.Split(url.RawQuery, "&") {
		data := strings.Split(i, "=")
		key := data[0]
		if len(data) == 2 {
			q.Set(key, data[1]+"'")
		} else {
			return ""
		}
	}
	url.RawQuery = q.Encode()
	return url.String()
}
func (sqlinj SQLInjectionPlugin) ParserResponse(response *fasthttp.Response) bool {
	str := response.Body()
	var errorRegexp, _ = regexp.Compile("Active.Server|ADODB.Field|An.illegal|An.unexpected|ASP.NET.is.configured.to.show.verbose.error.messages|ASP.NET_SessionId|A.syntax|Can't.connect|CLI.Driver|Custom.Error|data.source|DB2.Driver|DB2.Error|DB2.ODBC|detected.an|Died.at|Disallowed.Parent|Error.converting|Error.Diagnostic|Error.Message|Error.Report|Fatal.error|include_path|Incorrect.syntax|Index.of|Internal.Server|Invalid.Path|Invalid.procedure|invalid.query|Invision.Power|is.not.allowed.to.access|JDBC.Driver|JDBC.Error|JDBC.MySQL|JDBC.Oracle|JDBC.SQL|Microsoft.OLE|Microsoft.VBScript|missing.expression|MySQL.Driver|mysql.error|mySQL.error|MySQL.Error|MySQL.ODBC|ODBC.DB2|ODBC.Driver|ODBC.Error|ODBC.Microsoft|ODBC.Oracle|ODBC.SQL|OLE/DB.provider|on.MySQL|ORA-0|ORA-1|Oracle.DB2|Oracle.Driver|Oracle.Error|Oracle.ODBC|Parent.Directory|Permission.denied|PHP.Error|PHP.Parse|PHP.Warning|PostgreSQL.query|server.at|server.object|SQL.command|SQLException|SQL.Server|supplied.argument|Supplied.argument|Syntax.error|The.error.occurred.in|The.script.whose.uid.is|Type.mismatch|Unable.to.jump.to.row|Unclosed.quotation|unexpected.end|Unterminated.string|Warning..Cannot|Warning..mysql_query|Warning..pg_connect|Warning..Supplied|You.have.an.error.in.your.SQL.syntax|invalid.in.the.select.list.because")
	return errorRegexp.Match(str)
}
func (plugin SQLInjectionPlugin) GenInfo(url url.URL) string {
	return fmt.Sprintf("SQLINJ %s", plugin.GenPayload(url))
}

func (SQLInjectionPlugin) GetName() string {
	return "SQLINJ"
}
