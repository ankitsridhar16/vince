!function(){"use strict";var e,t,n,a=window.location,r=window.document,o=r.getElementById("vince"),c=o.getAttribute("data-api")||(e=o.src.split("/"),t=e[0],n=e[2],t+"//"+n+"/api/event");function i(e,t){try{if("true"===window.localStorage.vince_ignore)return void console.warn("Ignoring Event: localStorage flag")}catch(e){}var n={};n.n=e,n.u=t&&t.u?t.u:a.href,n.d=o.getAttribute("data-domain"),n.r=r.referrer||null,n.w=window.innerWidth,t&&t.meta&&(n.m=JSON.stringify(t.meta)),t&&t.props&&(n.p=t.props),n.h=1;var i=new XMLHttpRequest;i.open("POST",c,!0),i.setRequestHeader("Content-Type","text/plain"),i.send(JSON.stringify(n)),i.onreadystatechange=function(){4===i.readyState&&t&&t.callback&&t.callback()}}var d=window.vince&&window.vince.q||[];window.vince=i;for(var l=0;l<d.length;l++)i.apply(this,d[l])}();