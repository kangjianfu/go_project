{{define "header"}}
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
     <title>直播客</title>
	<!-- core CSS -->
    <link href="/static/bootstrap4.0/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/bootstrap4.0/css/font-awesome.css" rel="stylesheet">
 	<link href="/static/bootstrap4.0/css/normalize.css" rel="stylesheet">
	<script  type="text/javascript" src="/static/bootstrap4.0/js/jquery-1.11.min.js" ></script>
	<script  type="text/javascript" src="/static/bootstrap4.0/js/bootstrap.js"></script>
    <script type="text/javascript" src="/static/bootstrap4.0/js/touche.js"></script>
    <script type="text/javascript" src="/static/bootstrap4.0/js/p-pull-refresh.js"></script>
    <script type="text/javascript" src="/static/player/js/swfobject.js"></script>
    <script src="/static/bootstrap4.0/js/video.js"></script> 

    <style type="text/css">
        .nav-tabs .nav-link{
                border-top:1px solid transparent;
                border-radius:0px;
                bottom: 0px;
                color:black;
        }
        #nav_column>li.nav-item.active{
             background-color:#e5e5e5;border-radius:0px; border-color:#e5e5e5 transparent #e5e5e5
        }
       .preloader.visible {
        top: 0;
        opacity: 1;
       }
       .allpage{overflow: hidden; min-height:100%}
       .loading-warp{ display: table; width: 100%; margin-top: -5.9rem; }
       .loading-warp .box{ width: 100%; padding-top: 1rem; padding-bottom: 1rem;
        display: table-cell; text-align: center; vertical-align: middle; }
       .loading-warp .box img{ display: block; width: 2rem; height: 2rem; margin: 0 auto; }
       .loading-warp .box .text{ display: block; text-align: center; font-size: 0.5rem; 
        line-height: 0.5rem; opacity: 0.7; margin-top: 1.4rem; }
    .bottom-span-font{
        font-size: 14px
    }
    #download-btn{
        background-color:#ED0043;border-radius:19px ;font-size:15px;border-color:white
    }
    </style>
{{end}}