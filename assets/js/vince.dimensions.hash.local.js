!function(){"use strict";var o=window.location,c=window.document,d=c.currentScript,s=d.getAttribute("data-api")||new URL(d.src).origin+"/api/event";function e(e,t){try{if("true"===window.localStorage.vince_ignore)return void console.warn("Ignoring Event: localStorage flag")}catch(e){}var n={};n.n=e,n.u=o.href,n.d=d.getAttribute("data-domain"),n.r=c.referrer||null,n.w=window.innerWidth,t&&t.meta&&(n.m=JSON.stringify(t.meta)),t&&t.props&&(n.p=t.props);var i=d.getAttributeNames().filter(function(e){return"event-"===e.substring(0,6)}),r=n.p||{};i.forEach(function(e){var t=e.replace("event-",""),n=d.getAttribute(e);r[t]=r[t]||n}),n.p=r,n.h=1;var a=new XMLHttpRequest;a.open("POST",s,!0),a.setRequestHeader("Content-Type","text/plain"),a.send(JSON.stringify(n)),a.onreadystatechange=function(){4===a.readyState&&t&&t.callback&&t.callback()}}var t=window.vince&&window.vince.q||[];window.vince=e;for(var n,i=0;i<t.length;i++)e.apply(this,t[i]);function r(){n=o.pathname,e("pageview")}window.addEventListener("hashchange",r),"prerender"===c.visibilityState?c.addEventListener("visibilitychange",function(){n||"visible"!==c.visibilityState||r()}):r()}();