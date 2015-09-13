/*	Autor: Martin roberto Mondragon Sotelo
	Licencia. GPL
	Url: mygnet.com
	martin@mygnet.com
*/

mUtil = {
	isIE:navigator.userAgent.indexOf('MSIE')>=0?true:false, 
	isFF:navigator.userAgent.indexOf('Firefox')>=0?true:false,		
	isNS:navigator.userAgent.indexOf('Netscape')>=0?true:false,
	isCH:navigator.userAgent.indexOf('Chrome')>=0?true:false,
	isSF:false,
	isOP:window.opera?true:false,
	isMAC:navigator.platform.indexOf('Mac')>=0?true:false,
	isLIN:navigator.platform.indexOf('Linux')>=0?true:false,
	isWIN:navigator.platform.indexOf('Win')>=0?true:false,
	isREXP:(window.RegExp)?true:false,
	isVer:navigator.appVersion,		
	isIE4:false,
	isIE5:false,
	isIE6:false,
	isIE7:false,
	isIE8:false,
	isIE9:false,	
	init:function(){
		mUtil.isVer = mUtil.getVersion(); 
		mUtil.isSF	= navigator.userAgent.indexOf('Safari')>=0?(mUtil.isCH?false:true):false;
		mUtil.isIE4 = mUtil.isIE?(mUtil.isVer>=4.0 && mUtil.isVer<5.0?true:false):false;
		mUtil.isIE5 = mUtil.isIE?(mUtil.isVer>=5.0 && mUtil.isVer<6.0?true:false):false;
		mUtil.isIE6 = mUtil.isIE?(mUtil.isVer>=6.0 && mUtil.isVer<7.0?true:false):false;
		mUtil.isIE7 = mUtil.isIE?(mUtil.isVer>=7.0 && mUtil.isVer<8.0?true:false):false;
		mUtil.isIE8 = mUtil.isIE?(mUtil.isVer>=8.0 && mUtil.isVer<9.0?true:false):false;
		mUtil.isIE9 = mUtil.isIE?(mUtil.isVer>=9.0 && mUtil.isVer<10.0?true:false):false;
	},	
	getVersion:function(){
		var isVer = navigator.appVersion;
		if(isVer+"" != "NaN" && mUtil.isREXP)
		{ 	if(mUtil.isIE){	isVer.match(/(MSIE)(\s*)([0-9].[0-9]+)/ig);	isVer = RegExp.$3;	} 
			else if(mUtil.isSF){	isVer.match(/(Version\/)(\s*)([0-9].[0-9]+)/ig); isVer = RegExp.$3;	} 
			else if(mUtil.isCH){	isVer.match(/(Chrome\/)(\s*)([0-9].[0-9]+)/ig);	isVer = RegExp.$3;	} 
		}
		return isVer;
	},
	getBrowser:function()
	{	return mUtil.isIE?'Iternet Explorer':(mUtil.isOP?'Opera':(mUtil.isFF?'Firefox':(mUtil.isNS?'Netscape':(mUtil.isSF?'Safari':(mUtil.isCH?'Chrome':'Otro')))));
	},
	addLoad:function(func)
	{ 	var oldonload = window.onload;  
		if(typeof window.onload != 'function')window.onload = func;
		else window.onload = function(){ if(oldonload)oldonload(); func(); }
	},
	addUnLoad:function(func)
	{ 	var oldonload = window.onunload;  
		if(typeof window.onunload != 'function')window.onunload = func;
		else window.onunload = function(){ if(oldonload)oldonload(); func(); }
	},
	setNull:function(){ },
	cancelEvent:function(e)
	{	if (!e || mUtil.isIE){ window.event.returnValue = false;  window.event.cancelBubble = true; return window.event; }
		else{ e.stopPropagation(); e.preventDefault(); 	return e; }
	},
	win:function(url,width,height)
	{	var left = parseInt((screen.availWidth/2) - (width/2));
		var top = parseInt((screen.availHeight/2) - (height/2));
		var windowFeatures = "width=" + width + ",height=" + height +",status=0,resizable=0,left=" + left + ",top=" + top + ",screenX=" + left + ",screenY=" + top + ",titlebar=0,toolbar=0,scrollbars=1";
		win = window.open(url, null, windowFeatures);	
		win.focus();
	},
	winLeft:function(url,width,height,noresize,title)
	{	noresize = typeof(noresize) == 'undefined'?1:noresize;	
		title = typeof(title) == 'undefined'?null:title;
		var windowFeatures = "width=" + width + ",height=" + height +",status=0,resizable="+noresize+",left=0,top=0,screenX=0,screenY=0,titlebar=0,toolbar=0,scrollbars=1";
		win = window.open(url, title, windowFeatures);	
		win.focus();
	},
	ltrim:function(s) {  return s.replace(/^\s+/, "");},
	rtrim:function(s) {  return s.replace(/\s+$/, "");},
	trim:function(s) {  return mUtil.rtrim(mUtil.ltrim(s)); },
	obj:function(id)
	{  	if (document.getElementById)return document.getElementById(id);
		else if(document.all)return document.all[id];
		else if(document.layers[id]) return document.layers[id];
		else return null;
	},
	set:function(id,text)
	{	if (document.getElementById || document.all ){ 
				mUtil.obj(id).innerHTML = text;
			}
			else if (document.layers){
				var idLayer = document.layers[id];
				idLayer.document.open();
				idLayer.document.write(text);
				idLayer.document.close();
			}	
		},	
	get:function(id){ return mUtil.obj(id).innerHTML; },		
	code:function(text,ty)
	{	switch(ty)
		{ case 'html':	text = text.replace(/&ntilde;/ig,'ñ');
						text = text.replace(/&amp;/,'&');
						text = text.replace(/&iacute;/,'í');
						return text.replace(/&([a-zA-Z])(uml|acute|grave|circ|tilde);/,'$1');	
				break;
		  case 'text':	var car = new Array('&acute;','&Acute;','&ecute;','&Ecute;','&icute;','&Icute;','&ocute;','&Ocute;','&ucute;','&Ucute;','&ntilde;','&Ntilde;');
						var val = new Array('á'		 ,'Á'	   ,'é',	  'É',		'í',	  'Í',		'ó',	  'Ó',		'ú',	  'Ú',		'ñ',	   'Ñ');
						for(i=0; i<car.length; i++)text = text.replace(car[i],val[i]);	
						return text.replace(/&([a-zA-Z])(uml|grave|circ|tilde);/,'$1');	
				break;   	
		}
	}
};

mUtil.init();

mAjax = {
	status:		0,
	img:		'/template/skin/loadingBarra.gif',
	url:		'idx.ajax.php',
	top:		'40%',
	
	loading:function(id)
	{	mUtil.set(id,'<div class="mLoad" style="top='+mAjax.top+'; '+(mUtil.isFF?'height:0px;':'')+'"><img src="' + mAjax.img + '" alt="Cargando" style="float:none;"/></div>');
	},
	
	getXMLHttp:function()
	{	if(typeof(XMLHttpRequest) != 'undefined') return new XMLHttpRequest();
		var onjA=['Msxml2.XMLHTTP.6.0', 'Msxml2.XMLHTTP.4.0','Msxml2.XMLHTTP.3.0', 'Msxml2.XMLHTTP', 'Microsoft.XMLHTTP'];
		for(var i=0; i<onjA.length; i++) try{ return new ActiveXObject(onjA[i]); } catch(e){ }
		return null;
	},	
	
	send:function(id,elements,func)
	{	var xmlHttp		=	mAjax.getXMLHttp();	
		var method		=	'POST';
		if(!xmlHttp)return;
		try{ xmlHttp.open(method,mAjax.url+'/?idAjax=' + id + '&hash=' + Math.random() * 123456789 + (method=='GET'?'&'+elements:''),true ); 
			 xmlHttp.onreadystatechange=function()
				 {	switch(xmlHttp.readyState)
					{ 	case 1:	break; 
						case 2: break; 
						case 3: break; 
						case 4: if(xmlHttp.status != 200)
								{	if(typeof(func)=='function')func('<!ERROR!>',xmlHttp.status); 
									else mUtil.set(id,'ERROR DE CONEXIÓN',xmlHttp.statusText);
								}
								else
								{	if(xmlHttp.responseText.indexOf('RELOAD')!=-1){ window.location.reload();  return; }
									else if(typeof(func)=='function')func(id,xmlHttp.responseText); 
									else mUtil.set(id,xmlHttp.responseText);
								}
							break;					
					}					 
				 }	
		   }
        catch(excepcion){ return; }
        if(method=='POST'){ xmlHttp.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded'); }
		try{ xmlHttp.send(method=='POST'?elements:null); } catch(excepcion){ return; }
		if(id)mAjax.loading(id);
	}		
};