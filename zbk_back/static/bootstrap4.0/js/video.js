
//var pageHeight=window.screen.availHeight;
//var getPageWidth=window.screen.availWidth;
var pageHeight=0
var pageHeight=0
$(function(){


    
     pageHeight=$(document).height();
     pageWidth=$(document).width()
    //alert("屏幕总高度"+pageHeight)
    //alert("网页可用高度"+ph1)
	//alert(pageHeight);
	//alert(getPageWidth);


     //   // 竖屏
     // var video_width=360
     //  var video_height=640
     //  var video_url="http://s3.cn-north-1.amazonaws.com.cn/xvs/pub/ZBK_WEIXIN/20160618/16/02/m3u8/aws-cn_north_1-5-48f5eba3681d9be5.m3u8"

       // 竖屏
     // var video_width=360
     //  var video_height=480
     //  var video_url="http://s3.cn-north-1.amazonaws.com.cn/xvs/pub/ZBK_WEIXIN/20160618/16/42/m3u8/aws-cn_north_1-5-cd6f5fa2438ce145.m3u8"

       // 竖屏
     var video_width=480
      var video_height=640
      var video_url="http://s3.cn-north-1.amazonaws.com.cn/xvs/pub/ZBK_WEIXIN/20160618/16/48/m3u8/aws-cn_north_1-5-cad561a6343858cd.m3u8"

      

      //横屏
     //  var video_width=640
     // var video_height=360
     // var video_url="http://s3.cn-north-1.amazonaws.com.cn/xvs/pub/ZBK_WEIXIN/20160618/16/04/m3u8/aws-cn_north_1-5-93cd6760454a6092.m3u8"
    //横屏
     //  var video_width=480
     // var video_height=360
     // var video_url="http://s3.cn-north-1.amazonaws.com.cn/xvs/pub/ZBK_WEIXIN/20160618/16/30/m3u8/aws-cn_north_1-5-7cef012d17eaaefe.m3u8"

    // var video_width=640
    // var video_height=480
    // var video_url="http://s3.cn-north-1.amazonaws.com.cn/xvs/pub/ZBK_WEIXIN/20160618/15/59/m3u8/aws-cn_north_1-5-bc24cd1c759a9f70.m3u8"

    //判断如果不是移动设备
    console.info(browser)
    //alert(browser.versions.mobile)
    if(!browser.versions.mobile){
        $("#video_container").empty();
        //$("#video_container").height(video_height)
        $("#video_container").css({"text-align":"center"})
        $("#video_container").append("<div  id=\"video\"></div>");
        parameters.src=video_url;
        parameters.poster=video_url+".jpg"
        init_player(video_width,video_height)
    }
    if(browser.versions.mobile){
        reflash_all_page()
        if(browser.versions.iPhone||browser.versions.android){
             $("#video_container").parent().css({"margin-top":$("#nav_header").height()+10})
            $("#video_container").css({"padding":"0px"})
             $("#video_container").css({"margin":"0px"})
            // $("#video_container").css({"text-align":"center"})
            if(video_width>video_height){
                //横屏充满
                horizontal_full(video_width,video_height,pageWidth)
            }else{
                //竖屏充满屏幕2/3
                vertical_full()
            }
            var video_tag="<video controls='controls' poster="+video_url+".jpg  width='100%' height='100%'  webkit-playsinline ><source src="+video_url+" type=\"video/mp4\"></video> " 
            $("#video_container").append(video_tag);
        }else{
            $("#video_container").css({"text-align":"center"})
             $("#video_container").css({"margin":"0px"})
             $("#video_container").css({"padding":"0px"})
             var video_tag="<video controls='controls' poster="+video_url+".jpg  width="+video_width+" height="+video_height+"  webkit-playsinline ><source src="+video_url+" type=\"video/mp4\"></video> " 
            $("#video_container").append(video_tag);
        }

    }

    $('#nav_column a').click(function (e) {
        e.preventDefault()
        $(this).tab('show')
    })
})
//视频横屏拍摄模式
function horizontal_full(video_width,video_height,pageWidth){
     $("#video_container").height(video_height*(pageWidth/video_width))
     var tempHeight=$("#nav_header").height()+$("#nav_bottom").height()+$("#nav_column").height()+$("#video_container").height()
     $("#video_descrption").height(pageHeight-tempHeight)
    
}
function vertical_full(){
      $("#video_container").height(pageHeight*2/3*0.9)
      var tempHeight=$("#nav_header").height()+$("#nav_bottom").height()+$("#nav_column").height()+$("#video_container").height()
    $("#video_descrption").height(pageHeight-tempHeight)
}





//下拉刷新
function reflash_all_page(){
    var $statu = $('.loading-warp .text');
    var pullRefresh = $('.allpage').pPullRefresh({
            $el: $('.allpage'),
            $loadingEl: $('.loading-warp'),
            sendData: null,
            url: 'http://www.baidu.com',
            startPX: 6, 
            autoHide: true,
            callbacks: {
                pullStart: function(){
                    $statu.text('松开开始刷新');
                },
                start: function(){
                    $statu.text('数据刷新中···');
                },
                success: function(response){
                    $statu.text('数据刷新成功！');
                },
                end: function(){
                    $statu.text('下拉刷新结束');
                     location.reload()
                    //window.location.replace(window.location.href + "&ts=" + new Date().getTime());
                },
                error: function(){
                    //window.location.replace(window.location.href + "&ts=" + new Date().getTime());
                    $statu.text('找不到请求地址,数据刷新失败');
                }
            }
        });
}



//浏览器类型
var browser = {
    versions: function() {
        var u = navigator.userAgent,
        app = navigator.appVersion;
        return {
            mobile: !!u.match(/AppleWebKit.*Mobile.*/),
            ios: !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/),
            android: u.indexOf("Android") > -1 || u.indexOf("Linux") > -1,
            iPhone: u.indexOf("iPhone") > -1,
            iPad: u.indexOf("iPad") > -1
        };
    }(),
    language: (navigator.browserLanguage || navigator.language).toLowerCase()
};
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
    swfobject.embedSWF("static/player/swf/woan_wmp.swf", "video", width,
        height, "10.1.0","/static/player/swf/expressInstall.swf",
        parameters, {
        allowFullScreen : "true",
        wmode : "opaque"
        }, {
            name : "Woan_Player"
        });
}
//下载软件地址
function downloadjump() {
       window.location.href = 'http://a.app.qq.com/o/simple.jsp?pkgname=cn.com.zhiboke_weixin&g_f=991653';
}