//拖入打开md
document.addEventListener('drop', function(e){
      e.preventDefault();
      e.stopPropagation();

      var reader = new FileReader();
      reader.onload = function(e){
        editor.setValue(e.target.result);
      };

      reader.readAsText(e.dataTransfer.files[0]);
    }, false);

//保存为markdown
 function saveAsMarkdown(){
      save(editor.getValue(), "newuntitled.md");
}

//保存为html
function saveAsHtml() {
      save(document.getElementById('out').innerHTML, "untitled.html");
}

//监听点击保存事件
 document.getElementById('saveas-markdown').addEventListener('click', function() {
      saveAsMarkdown();
      hideMenu();
 });
 
 document.getElementById('saveas-html').addEventListener('click', function() {
      saveAsHtml();
      hideMenu();
    });
	
//保存函数
function save(code, name){
      var blob = new Blob([code], { type: 'text/plain' });
      if(window.saveAs){
        window.saveAs(blob, name);
      }else if(navigator.saveBlob){
        navigator.saveBlob(blob, name);
      }else{
        url = URL.createObjectURL(blob);
        var link = document.createElement("a");
        link.setAttribute("href",url);
        link.setAttribute("download",name);
        var event = document.createEvent('MouseEvents');
        event.initMouseEvent('click', true, true, window, 1, 0, 0, 0, 0, false, false, false, false, 0, null);
        link.dispatchEvent(event);
      }
}

//监听关闭菜单按钮事件
var menuVisible = false;
var menu = document.getElementById('menu');
document.getElementById('close-menu').addEventListener('click', function(){
      hideMenu();
    });

//显示菜单
 function showMenu() {
      menuVisible = true;
      menu.style.display = 'block';
    }

//隐藏菜单
function hideMenu() {
      menuVisible = false;
      menu.style.display = 'none';
    }

//监听按键	
document.addEventListener('keydown', function(e){
      if(e.keyCode == 83 && (e.ctrlKey || e.metaKey)){
        e.shiftKey ? showMenu() : saveAsMarkdown();

        e.preventDefault();
        return false;
      }

      if(e.keyCode === 27 && menuVisible){
        hideMenu();

        e.preventDefault();
        return false;
      }
    });
	
//初始化编辑器=============

var URL = window.URL || window.webkitURL || window.mozURL || window.msURL;
    navigator.saveBlob = navigator.saveBlob || navigator.msSaveBlob || navigator.mozSaveBlob || navigator.webkitSaveBlob;
    window.saveAs = window.saveAs || window.webkitSaveAs || window.mozSaveAs || window.msSaveAs;

    // Because highlight.js is a bit awkward at times
var languageOverrides = {
      js: 'javascript',
      html: 'xml'
    };

    emojify.setConfig({ img_dir: 'emoji' });

    var md = markdownit({
      html: true,
      linkify: true,
      highlight: function(code, lang){
        if(languageOverrides[lang]) lang = languageOverrides[lang];
        if(lang && hljs.getLanguage(lang)){
          try {
            return hljs.highlight(lang, code).value;
          }catch(e){}
        }
        return '';
      }
    }).use(markdownitFootnote);


    var hashto;

    function update(e){
      setOutput(e.getValue());

      clearTimeout(hashto);
      hashto = setTimeout(updateHash, 1000);
    }

    function setOutput(val){
      val = val.replace(/<equation>((.*?\n)*?.*?)<\/equation>/ig, function(a, b){
        return '<img src="http://latex.codecogs.com/png.latex?' + encodeURIComponent(b) + '" />';
      });

      var out = document.getElementById('out');
      var old = out.cloneNode(true);
      out.innerHTML = md.render(val);
      emojify.run(out);

      var allold = old.getElementsByTagName("*");
      if (allold === undefined) return;

      var allnew = out.getElementsByTagName("*");
      if (allnew === undefined) return;

      for (var i = 0, max = Math.min(allold.length, allnew.length); i < max; i++) {
        if (!allold[i].isEqualNode(allnew[i])) {
          out.scrollTop = allnew[i].offsetTop;
          return;
        } 
      }
    }

    var editor = CodeMirror.fromTextArea(document.getElementById('code'), {
      mode: 'gfm',
      lineNumbers:true,
      matchBrackets: true,
      lineWrapping: true,
      theme: 'base16-light',
      extraKeys: {"Enter": "newlineAndIndentContinueMarkdownList"}
    });

    editor.on('change', update);
	
	function updateHash(){
      window.location.hash = btoa( // base64 so url-safe
        RawDeflate.deflate( // gzip
          unescape(encodeURIComponent( // convert to utf8
            editor.getValue()
          ))
        )
      );
    }

    if(window.location.hash){
      var h = window.location.hash.replace(/^#/, '');
      if(h.slice(0,5) == 'view:'){
        setOutput(decodeURIComponent(escape(RawDeflate.inflate(atob(h.slice(5))))));
        document.body.className = 'view';
      }else{
        editor.setValue(
          decodeURIComponent(escape(
            RawDeflate.inflate(
              atob(
                h
              )
            )
          ))
        );
        update(editor);
        editor.focus();
      }
    }else{
      update(editor);
      editor.focus();
    }
	
//编辑器工具栏==========================================
var cm=editor;
document.getElementById('homepage').addEventListener('click', function() {
	  location.href='/';
});
document.getElementById('editormd-bold1').addEventListener('click', function() {
	  h1();
});

document.getElementById('editormd-bold2').addEventListener('click', function() {
	  h2();
});

document.getElementById('fa-bold').addEventListener('click', function() {
	  bold();
});

document.getElementById('fa-italic').addEventListener('click', function() {
	  italic();
});

//新建文件
document.getElementById('fa-file-o').addEventListener('click', function() {	
	NewFile();
	 
});

//保存文件
document.getElementById('fa-save').addEventListener('click', function() {	
	saveMd();
	 
});

//打开文件
$("[type=\"file\"]").bind("change",readFileContent);
document.getElementById('fa-folder-open-o').addEventListener('click', function() {
	   $("#fa-folder-open-oinput").trigger("click");
});

//读取文件,拖拽或打开文件读取
function readFileContent(events) {
      events.preventDefault();
      events.stopPropagation();
	  
	  var r = (events.dataTransfer || events.target).files;  //dataTransfer:拖拽读取,target:打开文件读取

      var reader = new FileReader();
      reader.onload = function(e){
        editor.setValue(e.target.result);
      };
      reader.readAsText(r[0]);

 }



function h1(){
	//申明多个变量
	var e =cm,
	t = e.getCursor(e),  //返回光标所在元素 line:行,ch：列。
	i = e.getSelection(e);//得到选中的文字，如果未选，则为空
	console.log(i);
    0 !== t.ch ? (e.setCursor(t.line, 0), e.replaceSelection("# " + i), e.setCursor(t.line, t.ch + 2)) : e.replaceSelection("# " + i)
}

function h2() {
                var e = this.cm,
                t = e.getCursor(),
                i = e.getSelection();
                0 !== t.ch ? (e.setCursor(t.line, 0), e.replaceSelection("## " + i), e.setCursor(t.line, t.ch + 3)) : e.replaceSelection("## " + i)
}

//加粗
function bold() {
                var e = this.cm,
                t = e.getCursor(),
                i = e.getSelection();
                e.replaceSelection("**" + i + "**"),
                "" === i && e.setCursor(t.line, t.ch + 2)
}

//斜体
 function italic() {
                var e = this.cm,
                t = e.getCursor(),
                i = e.getSelection();
                e.replaceSelection("*" + i + "*"),
                "" === i && e.setCursor(t.line, t.ch + 1)
}
			
//替换文本
function  replaceSelection(e){
    return cm.replaceSelection(e);
}

function setCursor(e) {
    return cm.setCursor(e);
}

//保存
function saveMd(){
	
	    var md =cm.getValue();
		var title = cm.lineInfo(0);
		var id = parseInt($("#id").val())?parseInt($("#id").val()):0;
		if(!md){
			alert("不能为空");
			return false;
		}
		$.ajax({
			     type: "POST",
	             url: "/md/submitmd",
	             data:{id:id,content:md,title:title.text,submit:1},
	             dataType: "json",
	             success: function(data){
		          alert(data.msg);
				  $("#id").val(data.data.id);
				  setTimeout("window.location.reload()",1000);
				 }
		})	
}


//新建文件
function NewFile(){
	//window.location.href="/editor";
	window.open("/editor", "_blank");
}

                
