	function GetQueryString(name) {
	  var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
	  var r = window.location.search.substr(1).match(reg);
	  if (r != null) return unescape(r[2]);
	  return null;
	}

	$("#myTextarea").markdown({
	  autofocus: false,
	  savable: false,
	  height: 700,
	  onShow: function(e) {
	    var date = formatDate(new Date(), "yyyy-MM-dd HH:mm:ss");
	    var headContent = "---\ntitle: " + "此处添加标题" +
	      "  \ndate: " + date.toLocaleString() + "  \ntags:  \n---";

	    var fileName = decodeURI(decodeURI(GetQueryString("name")));
	    if (fileName != "" && fileName != null) {
	      //根据明文地址访问后台获取对应内容追加到编辑器中
	      $.ajax({
	        url: '/findContent',
	        type: 'POST',
	        data: {
	          name: fileName
	        },
	        success: function(e) {
	          headContent = e.data;
	          $("#myTextarea").text(headContent);
	        },
	        error: function(e) {
	          alert("Get Cotent faild");
	        },
	        complete: function(e) {
	          $("#myTextarea").text(headContent);
	          $(".md-editor").addClass("md-fullscreen-mode");
	          $('body').addClass('md-nooverflow');
	        }
	      });
	    }else{
				$("#myTextarea").text(headContent);
				$(".md-editor").addClass("md-fullscreen-mode");
				$('body').addClass('md-nooverflow');
			}
	  },
	  onFocus: function(e) {
	    return;
	  },
	  additionalButtons: [
	    [{
	      name: "deploy",
	      data: [{
	        name: "deploy",
	        toggle: true,
	        title: "deploy",
	        btnText: '发布',
	        btnClass: 'btn btn-primary btn-sm',
	        callback: function(e) {
	          alert("Maybe you will wait a few minutes.....");
	          $.ajax({
	            url: '/obtainContent',
	            type: 'POST',
	            data: $("#contentForm").serialize(),
	            success: function(e) {
	              alert("Deploy Success");
	            },
	            error: function(e) {
	              alert("Deploy Faild");
	            }
	          });
	        }
	      }]
	    }, {
	      name: "cancel",
	      data: [{
	        name: "cancel",
	        toggle: true, // this param only take effect if you load bootstrap.js
	        title: "cancel",
	        btnText: '取消',
	        btnClass: 'btn btn-primary btn-sm',
	        callback: function(e) {
	          // alert("点击取消发布");
	          window.location.href = "/";
	        }
	      }]
	    }]
	  ]
	});
