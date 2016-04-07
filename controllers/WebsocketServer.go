package controllers
/*
import (
	//"io"
	"log"
	//"encoding/json"
	//"net/http"
	"html/template"
	"t_client/pub"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
)

type WebSocketController struct {
	beego.Controller
}

// Get method handles GET requests for WebSocketController.
func (this *WebSocketController) Get() {
	this.EnableRender = false
	this.TplName = ""
	c, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
			var (
				msg pub.Msg1001
				head pub.MsgHead
				body []pub.StreamInfo
			)
			body =[]pub.StreamInfo{
				{1,"1号桌_全景","rtsp://1","rtmp://1",},
				{2,"2号桌_全景","rtsp://2","rtmp://2",},
			}
			head = pub.MsgHead{
                1,
				2,
				3,
            }
			msg = pub.Msg1001{
				head,
				body,
			}
		//err = c.WriteMessage(mt, message)
		err = c.WriteJSON(msg)
		log.Printf("write:", msg)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

var homeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output"></div>
</td></tr></table>
</body>
</html>
`))
*/