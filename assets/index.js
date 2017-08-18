(function(){

  var model = {
    target: { loaded: 0, total:0, name: ''},
  };

  var ui = {
    source: document.querySelector('input[type=file]#source'),
    submit: document.querySelector('button#submit'),
    progress: {
      bar:     document.querySelector('div#progress-bar'),
      message: document.querySelector('div#progress-message'),
    },
    update: function() {
      this.progress.message.innerText = "(" + model.target.loaded + "/" + model.target.total + ")" + " " + model.target.name;
      var r = ((model.target.total) ? (model.target.loaded / model.target.total) : 0)
      this.progress.bar.style.width = (r * 100) + "%";
    },
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
      xhr.upload.onprogress = function(ev) {
        model.target.loaded = ev.loaded;
      };
      xhr.open(opt.method, url, true)
      if (opt.type) xhr.responseType = opt.type;
      opt.body ? xhr.send(opt.body) : xhr.send();
      return p;
    },
    test: function() {
      return this.fetch("/upload", {});
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
    if (source.files.length == 0) return alert('No source file is set');
    var file = source.files[0];
    api.convert(file).then(function(blob) {
      var url = URL.createObjectURL(blob);
      window.open(url);
    }).catch(function(err) {
      console.log("ERROR", err);
    });
  });

  ui.source.addEventListener('change', function() {
    var file = source.files[0];
    model.target = {
      total: file.size,
      loaded: 0,
      name: file.name,
    };
    ui.update();
  });

})();
