//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./static"))))
var socket = null;

addr.value='ws://d.newb.xyz:1234/c';
btnConnect.onclick=function(){
    socket=new WebSocket(addr.value)
    //socket.onopen=function(event){
    //    socket.send('ping');
    //}
    socket.onmessage=function(event){
	    console.log(event)
        show(eval('(' + event.data + ')'))
    }
}



// tester
function tester(){
	str='{"Timestamp":1438930391,"Codeline":"/go/src/caidanhezi-go/httpServer.go 19","Level":4,"Info":"Service start @port :9090","Detail":null}'
	show(eval('(' + str + ')'))
}

//WS

//View
function show(logObj){
    logObj.Timestamp = (new Date(logObj.Timestamp* 1000)).toLocaleString();
    console.log(logObj)
        appending =hereDoc(singleLog).replace('[[[time]]]',logObj.Timestamp)
        .replace('[[[level]]]',logObj.Level).replace('[[[info]]]',logObj.Info)
        .replace('[[[Codeline]]]',logObj.Codeline);
    if (logObj.Detail == null) {
        appending=appending.replace('[[[Detail]]]','null');
    }else{
        appending=appending.replace('[[[Detail]]]',logObj.Detail);
    }
    $('#logtable')[0].innerHTML+=appending;
}


function singleLog(){/*
        <tr >
          <td style="vertical-align: top;">[[[time]]]</td>
          <td>[[[level]]]</td>
          <td>[[[Codeline]]]</td>
          <td>[[[info]]]</td>
          <td>[[[Detail]]]</td>
        </tr>
*/}

function hereDoc(func) {
return func.toString().split(/\n/).slice(1, -1).join('\n');
}

