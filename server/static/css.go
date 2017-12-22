// Code generated by go generate; DO NOT EDIT.
// 2017-12-22 11:25:01.96382557 -0800 PST m=+0.028793332

package static

var Stylesheets = map[string]string{
	"black":  `body{background:#222;color:#efefef}h1,h2,h3{color:#aaa}a{color:#aaa}a:focus,a:hover{color:#ddd}.header li{border-color:#333}.header a{color:#ddd;font-weight:400}.header .active a{font-weight:400;color:#9b9494}.header a:focus,.header a:hover{color:rgba(82,168,236,.85)}.page-header h1{border-color:#333}.logo a:hover span{color:#555}table,th,td{border:1px solid #555}th{background:#333;color:#aaa;font-weight:400}tr:hover{background-color:#333;color:#aaa}input[type=url],input[type=password],input[type=text]{border:1px solid #555;background:#333;color:#ccc}input[type=url]:focus,input[type=password]:focus,input[type=text]:focus{color:#efefef;border-color:rgba(82,168,236,.8);box-shadow:0 0 8px rgba(82,168,236,.6)}.button-primary{border-color:#444;background:#333;color:#efefef}.button-primary:hover,.button-primary:focus{border-color:#888;background:#555}.alert,.alert-success,.alert-error,.alert-info,.alert-normal{color:#efefef;background-color:#333;border-color:#444}.panel{background:#333;border-color:#555;color:#9b9b9b}.unread-counter{color:#bbb}.category{color:#efefef;background-color:#333;border-color:#444}.category a{color:#999}.category a:hover,.category a:focus{color:#aaa}.pagination a{color:#aaa}.pagination-bottom{border-color:#333}.item{border-color:#666;padding:4px}.item.current-item{border-width:2px;border-color:rgba(82,168,236,.8);box-shadow:0 0 8px rgba(82,168,236,.6)}.item-title a{font-weight:400}.item-status-read .item-title a{color:#666}.item-status-read .item-title a:focus,.item-status-read .item-title a:hover{color:rgba(82,168,236,.6)}.item-meta a:hover,.item-meta a:focus{color:#aaa}.item-meta li:after{color:#ddd}.entry header{border-color:#333}.entry header h1 a{color:#bbb}.entry-content,.entry-content p,ul{color:#999}.entry-content pre,.entry-content code{color:#fff;background:#555;border-color:#888}.entry-enclosure{border-color:#333}`,
	"common": `*{margin:0;padding:0;box-sizing:border-box}body{font-family:helvetica neue,Helvetica,Arial,sans-serif;text-rendering:optimizeLegibility}main{padding-left:5px;padding-right:5px}a{color:#36c}a:focus{outline:0;color:red;text-decoration:none;border:1px dotted #aaa}a:hover{color:#333;text-decoration:none}.header{margin-top:10px;margin-bottom:20px}.header nav ul{display:none}.header li{cursor:pointer;padding-left:10px;line-height:2.1em;font-size:1.2em;border-bottom:1px dotted #ddd}.header li:hover a{color:#888}.header a{font-size:.9em;color:#444;text-decoration:none;border:0}.header .active a{font-weight:600}.header a:hover,.header a:focus{color:#888}.page-header{margin-bottom:25px}.page-header h1{font-weight:500;border-bottom:1px dotted #ddd}.page-header ul{margin-left:25px}.page-header li{list-style-type:circle;line-height:1.8em}.logo{cursor:pointer;text-align:center}.logo a{color:#000;letter-spacing:1px}.logo a:hover{color:#396}.logo a span{color:#396}.logo a:hover span{color:#000}@media(min-width:600px){body{margin:auto;max-width:750px}.logo{text-align:left;float:left;margin-right:15px}.header nav ul{display:block}.header li{display:inline;padding:0;padding-right:15px;line-height:normal;border:0;font-size:1em}.page-header ul{margin-left:0}.page-header li{display:inline;padding-right:15px}}table{width:100%;border-collapse:collapse}table,th,td{border:1px solid #ddd}th,td{padding:5px;text-align:left}td{vertical-align:top}th{background:#fcfcfc}tr:hover{background-color:#f9f9f9}.column-40{width:40%}.column-25{width:25%}.column-20{width:20%}label{cursor:pointer;display:block}.radio-group{line-height:1.9em}div.radio-group label{display:inline-block}select{margin-bottom:15px}input[type=url],input[type=password],input[type=text]{border:1px solid #ccc;padding:3px;line-height:20px;width:250px;font-size:99%;margin-bottom:10px;margin-top:5px;-webkit-appearance:none}input[type=url]:focus,input[type=password]:focus,input[type=text]:focus{color:#000;border-color:rgba(82,168,236,.8);outline:0;box-shadow:0 0 8px rgba(82,168,236,.6)}input[type=checkbox]{margin-bottom:15px}::-moz-placeholder,::-ms-input-placeholder,::-webkit-input-placeholder{color:#ddd;padding-top:2px}.form-help{font-size:.9em;color:brown;margin-bottom:15px}.form-section{border-left:2px dotted #ddd;padding-left:20px;margin-left:10px}a.button{text-decoration:none}.button{display:inline-block;-webkit-appearance:none;-moz-appearance:none;font-size:1.1em;cursor:pointer;padding:3px 10px;border:1px solid;border-radius:unset}.button-primary{border-color:#3079ed;background:#4d90fe;color:#fff}.button-primary:hover,.button-primary:focus{border-color:#2f5bb7;background:#357ae8}.button-danger{border-color:#b0281a;background:#d14836;color:#fff}.button-danger:hover,.button-danger:focus{color:#fff;background:#c53727}.button:disabled{color:#ccc;background:#f7f7f7;border-color:#ccc}.buttons{margin-top:10px;margin-bottom:20px}.alert{padding:8px 35px 8px 14px;margin-bottom:20px;color:#c09853;background-color:#fcf8e3;border:1px solid #fbeed5;border-radius:4px;overflow:auto}.alert h3{margin-top:0;margin-bottom:15px}.alert-success{color:#468847;background-color:#dff0d8;border-color:#d6e9c6}.alert-error{color:#b94a48;background-color:#f2dede;border-color:#eed3d7}.alert-error a{color:#b94a48}.alert-info{color:#3a87ad;background-color:#d9edf7;border-color:#bce8f1}.panel{color:#333;background-color:#fcfcfc;border:1px solid #ddd;border-radius:5px;padding:10px;margin-bottom:15px}.panel h3{font-weight:500;margin-top:0;margin-bottom:20px}.panel ul{margin-left:30px}.login-form{margin:50px auto 0;max-width:280px}.unread-counter{font-size:.8em;font-weight:300;color:#666}.category{font-size:.75em;background-color:#fffcd7;border:1px solid #d5d458;border-radius:5px;margin-left:.25em;padding:1px .4em;white-space:nowrap}.category a{color:#555;text-decoration:none}.category a:hover,.category a:focus{color:#000}.pagination{font-size:1.1em;display:flex;align-items:center;padding-top:8px}.pagination-bottom{border-top:1px dotted #ddd;margin-bottom:15px;margin-top:50px}.pagination>div{flex:1}.pagination-next{text-align:right}.pagination-prev:before{content:"« "}.pagination-next:after{content:" »"}.pagination a{color:#333}.pagination a:hover,.pagination a:focus{text-decoration:none}.item{border:1px dotted #ddd;margin-bottom:20px;padding:5px;overflow:hidden}.item.current-item{border:3px solid #bce;padding:3px}.item-title a{text-decoration:none;font-weight:600}.item-status-read .item-title a{color:#777}.item-meta{color:#777;font-size:.8em}.item-meta a{color:#777;text-decoration:none}.item-meta a:hover,.item-meta a:focus{color:#333}.item-meta ul{margin-top:5px}.item-meta li{display:inline}.item-meta li:after{content:"|";color:#aaa}.item-meta li:last-child:after{content:""}.hide-read-items .item-status-read{display:none}.entry header{padding-bottom:5px;border-bottom:1px dotted #ddd}.entry header h1{font-size:2em;line-height:1.25em;margin:30px 0}.entry header h1 a{text-decoration:none;color:#333}.entry header h1 a:hover,.entry header h1 a:focus{color:#666}.entry-actions{margin-bottom:20px}.entry-actions li{display:inline}.entry-actions li:not(:last-child):after{content:"|"}.entry-meta{font-size:.95em;margin:0 0 20px;color:#666;overflow-wrap:break-word}.entry-website img{vertical-align:top}.entry-website a{color:#666;vertical-align:top;text-decoration:none}.entry-website a:hover,.entry-website a:focus{text-decoration:underline}.entry-date{font-size:.65em;font-style:italic;color:#555}.entry-content{padding-top:15px;font-size:1.2em;font-weight:300;font-family:Georgia,times new roman,Times,serif;color:#555;line-height:1.4em;overflow-wrap:break-word}.entry-content h1,h2,h3,h4,h5,h6{margin-top:15px;margin-bottom:10px}.entry-content iframe,.entry-content video,.entry-content img{max-width:100%}.entry-content figure{margin-top:15px;margin-bottom:15px}.entry-content figure img{border:1px solid #000}.entry-content figcaption{font-size:.75em;text-transform:uppercase;color:#777}.entry-content p{margin-top:10px;margin-bottom:15px}.entry-content a{overflow-wrap:break-word}.entry-content a:visited{color:purple}.entry-content dt{font-weight:500;margin-top:15px;color:#555}.entry-content dd{margin-left:15px;margin-top:5px;padding-left:20px;border-left:3px solid #ddd;color:#777;font-weight:300;line-height:1.4em}.entry-content blockquote{border-left:4px solid #ddd;padding-left:25px;margin-left:20px;margin-top:20px;margin-bottom:20px;color:#888;line-height:1.4em;font-family:Georgia,serif}.entry-content blockquote+p{color:#555;font-style:italic;font-weight:200}.entry-content q{color:purple;font-family:Georgia,serif;font-style:italic}.entry-content q:before{content:"“"}.entry-content q:after{content:"”"}.entry-content pre{padding:5px;background:#f0f0f0;border:1px solid #ddd;overflow:scroll;overflow-wrap:initial}.entry-content table{table-layout:fixed;max-width:100%}.entry-content ul,.entry-content ol{margin-left:30px}.entry-content ul{list-style-type:square}.entry-enclosures h3{font-weight:500}.entry-enclosure{border:1px dotted #ddd;padding:5px;margin-top:10px;max-width:100%}.entry-enclosure-download{font-size:.85em;overflow-wrap:break-word}.enclosure-video video,.enclosure-image img{max-width:100%}.confirm{font-weight:500;color:#ed2d04}.confirm a{color:#ed2d04}.loading{font-style:italic}.bookmarklet{border:1px dashed #ccc;border-radius:5px;padding:15px;margin:15px;text-align:center}.bookmarklet a{font-weight:600;text-decoration:none;font-size:1.2em}`,
}

var StylesheetsChecksums = map[string]string{
	"black":  "a4d90d8c24bed43c38ebfa1054f7f5623423c16005843637bec59b3d98b0bcba",
	"common": "ee1b00f84c731ba5834b6a899fe56e0a09bf771ebc172884262d2b4eac854100",
}
