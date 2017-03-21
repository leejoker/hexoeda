	$("#myTextarea").markdown({
		autofocus:false,
		savable:false,
		height:700,
		onShow: function(e){
		 	var date=formatDate(new Date(),"yyyy-MM-dd HH:mm:ss");
			$("#myTextarea").text("---\ntitle: "+"此处添加标题"+
			"  \ndate: "+date.toLocaleString()+"  \ntags:  \n---" );
			$(".md-editor").addClass("md-fullscreen-mode");
			$('body').addClass('md-nooverflow');
		}
		,onFocus: function(e){
			return;
		}
		,additionalButtons: [
		    [{
		        name: "deploy",
		        data: [{
		            name: "deploy",
		            toggle: true, // this param only take effect if you load bootstrap.js
		            title: "deploy",
		            btnText: '发布',
					btnClass: 'btn btn-primary btn-sm',
		            callback: function(e){
						alert("Maybe you will wait a few minutes.....");
						$.ajax({
							url:'/obtainContent',
							type: 'POST',
							data: $("#contentForm").serialize(),
							success: function(e){
								alert("Deploy Success");
							},
							error: function(e){
								alert("Deploy Faild");
							}
						});

						// var contentForm = document.getElementById("contentForm");
						// contentForm.action = "/obtainContent";
						// contentForm.method =  "post";
						// contentForm.submit();
		            }
		          }]
		    }
		    ,{
		        name: "cancel",
		        data: [{
		            name: "cancel",
		            toggle: true, // this param only take effect if you load bootstrap.js
		            title: "cancel",
		            btnText: '取消',
					btnClass: 'btn btn-primary btn-sm',
		            callback: function(e){
						// alert("点击取消发布");
						window.location.href="/";
		            }
		        }]
		    }]
		]
	});