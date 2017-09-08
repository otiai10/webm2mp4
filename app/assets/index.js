(function(){

  Promise.prototype.progress = function(onprogress) {
    this.onprogress = onprogress;
    return this;
  };

  var ui = {
    buttons: document.querySelector('div#buttons'),
    submit:  document.querySelector('a#submit'),
    source:  document.querySelector('input[type=file]'),
    message: document.querySelector('div#message'),
    error: function(msg) {
      ui.message.innerHTML = document.querySelector('script#message-danger').innerHTML.replace('#{message}', msg);
    },
    reset: function() {
      ui.message.innerHTML = '';
      var a = document.querySelector('a#download');
      if (a) a.remove();
    }
  };
  ui.submit.startLoading = function() {
    ui.submit.setAttribute('disabled', true);
    ui.submit.className += ' is-loading';
  };
  ui.submit.endLoading = function() {
    ui.submit.removeAttribute('disabled');
    ui.submit.className = ui.submit.className.replace(/[ ]*is-loading[ ]*/, '');
  };

  var api = {
    fetch: function(url, opt) {
      var xhr = new XMLHttpRequest();
      xhr.upload.onprogress = function(ev) {
        if (typeof p.onprogress == 'function') p.onprogress(ev);
      };
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
    if (ui.source.files.length == 0) return ui.error("No file specified")
    ui.reset();
    ui.submit.startLoading();
    var file = ui.source.files[0];
    api.convert(file).progress(function(ev) {
      console.log('Progress', ev);
    }).then(function(blob) {
      ui.submit.endLoading();
      var url = URL.createObjectURL(blob);
      var a = document.createElement('a');
      a.className = "button is-info";
      a.download = "result.mp4";
      a.href = url;
      a.innerText = "Download";
      a.id = "download";
      a.addEventListener('click', function() { a.remove(); });
      ui.buttons.appendChild(a);
    }).catch(function(err) {
      ui.submit.endLoading();
      console.log("ERROR", err);
      ui.error(err.message || JSON.stringify(err, null, "\t"));
    });
  });

})();
