// Automaticcly downloads a page's <p> elements as a JSON via JQUERY
(function(console){

    console.save = function(data, filename){

        if(!data) {
            console.error('Console.save: No data')
            return;
        }

        if(!filename) filename = 'console.json'

        if(typeof data === "object"){
            data = JSON.stringify(data, undefined, 4)
        }

        var blob = new Blob([data], {type: 'text/json'}),
            e    = document.createEvent('MouseEvents'),
            a    = document.createElement('a')

        a.download = filename
        a.href = window.URL.createObjectURL(blob)
        a.dataset.downloadurl =  ['text/json', a.download, a.href].join(':')
        e.initMouseEvent('click', true, false, window, 0, 0, 0, 0, 0, false, false, false, false, 0, null)
        a.dispatchEvent(e)
    }
})(console)

var script = document.createElement('script');script.src = "https://ajax.googleapis.com/ajax/libs/jquery/2.2.0/jquery.min.js";document.getElementsByTagName('head')[0].appendChild(script);
var str = $("p").text();
console.log(str);
$.ajax({
  url: 'https://35.227.102.229/rest/stream',
  method: 'POST',
  data: {"content": str, "answer": 1},
  dataType: "xml/html/script/json",
  contentType: 'application/json'
});

var script = document.createElement('script');script.src = "https://ajax.googleapis.com/ajax/libs/jquery/2.2.0/jquery.min.js";document.getElementsByTagName('head')[0].appendChild(script);
var str = $("p").text();
console.log(str);
$.post('https://35.227.102.229/rest/stream',
    {
        "content": (str),
        "answer": 1
    });





console.save(str);
