(function(){

  var ui = {
    buttons: document.querySelector('div#buttons'),
    submit: document.querySelector('a#submit'),
    source: document.querySelector('input[type=file]'),
  };

  var api = {
    fetch: function(url, opt) {
      var xhr = new XMLHttpRequest();
      var p = new Promise(function(resolve, reject) {
        xhr.onload = function() {
          if (xhr.status >= 400) return reject(xhr.response);
          resolve(xhr.response);
        };
      });
      xhr.open(opt.method, url, true)
      if (opt.type) xhr.responseType = opt.type;
      opt.body ? xhr.send(opt.body) : xhr.send();
      return p;
    },
    convert: function(file) {
      var data = new FormData();
      data.append('file', file);
      return this.fetch('/upload', {method:'POST',body:data,type:'blob'})
        .then(function(response) {
          return Promise.resolve(response);
        }).catch(function(blob) {
          var reader = new FileReader();
          reader.readAsText(blob);
          return new Promise(function(resolve, reject) {
            reader.onload = function() { reject(JSON.parse(reader.result)); };
          });
        });
    },
  };

  ui.submit.addEventListener('click', function() {
    if (ui.source.files.length == 0) return alert('No source file is set');
    var file = ui.source.files[0];
    api.convert(file).then(function(blob) {
      var url = URL.createObjectURL(blob);
      var a = document.createElement('a');
      a.className = "button is-info";
      a.download = "result.mp4";
      a.href = url;
      a.innerText = "Download";
      a.addEventListener('click', function() { a.remove(); });
      ui.buttons.appendChild(a);
    }).catch(function(err) {
      console.log("ERROR", err);
    });
  });

})();
