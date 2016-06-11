
//播放器的参数
var parameters = {
	src : "",
	autoPlay : "false",
	verbose : true,
	controlBarAutoHide : "true",
	controlBarPosition : "bottom",
	poster : "/static/player/images/poster.png",
	plugin_hls :"/static/player/swf/wmp_plugin_hls.swf"
}


function init_player(width,height) {
	// Embed the player SWF:
	swfobject.embedSWF("static/player/swf/woan_wmp.swf", "player_div", width,
			height, "10.1.0","/static/player/swf/expressInstall.swf",
			parameters, {
				allowFullScreen : "true",
				wmode : "opaque"
			}, {
				name : "Woan_Player"
			});
}

function init_copy(taskId) {
	ZeroClipboard.setDefaults({
		moviePath : getContextPath() + "/views/js_copy/ZeroClipboard.swf"
	});
//	console.info($("#" + taskId))
	var clip = new ZeroClipboard($("#" + taskId));
	clip.on("load", function(client) {
		debugstr("Flash  loaded and ready.");
	});
	clip.setHandCursor(true);
	clip.addEventListener("mouseUp", function(client) {
		alert("复制成功");
	});

	/*clip.on("click", function(client) {
		debugstr("Flash movie loaded and ready.1111");
		client.on("complete", function(client, args) {
			if (args && args.text.length > 0)
				alert("地址已复制！");
		});
	});*/
	clip.on("noFlash", function(client) {
		debugstr("Your browser has no Flash.");
	});

	clip.on("wrongFlash", function(client, args) {
		debugstr("Flash 10.0.0+ is required but you are running Flash "
				+ args.flashVersion.replace(/,/g, "."));
	});

	function debugstr(text) {
		console.log(text);
	}
	return clip;
}

