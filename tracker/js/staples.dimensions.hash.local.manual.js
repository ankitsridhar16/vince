(function(){"use strict";var t,n,i=window.location,s=window.document,e=s.currentScript,a=e.getAttribute("data-api")||c(e);function r(e){console.warn("Ignoring Event: "+e)}function c(e){return new URL(e.src).origin+"/api/event"}function o(t,n){try{if(window.localStorage.vince_ignore==="true")return r("localStorage flag")}catch{}var c,l,d,o={};o.n=t,o.u=n&&n.u?n.u:i.href,o.d=e.getAttribute("data-domain"),o.r=s.referrer||null,o.w=window.innerWidth,n&&n.meta&&(o.m=JSON.stringify(n.meta)),n&&n.props&&(o.p=n.props),d=e.getAttributeNames().filter(function(e){return e.substring(0,6)==="event-"}),l=o.p||{},d.forEach(function(t){var n=t.replace("event-",""),s=e.getAttribute(t);l[n]=l[n]||s}),o.p=l,o.h=1,c=new XMLHttpRequest,c.open("POST",a,!0),c.setRequestHeader("Content-Type","text/plain"),c.send(JSON.stringify(o)),c.onreadystatechange=function(){c.readyState===4&&n&&n.callback&&n.callback()}}n=window.vince&&window.vince.q||[],window.vince=o;for(t=0;t<n.length;t++)o.apply(this,n[t])})()