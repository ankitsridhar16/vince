!function(){"use strict";var e,t,n,a=window.location,o=window.document,r=o.getElementById("vince"),c=r.getAttribute("data-api")||(e=r.src.split("/"),t=e[0],n=e[2],t+"//"+n+"/api/event");function d(e){console.warn("Ignoring Event: "+e)}function i(e,t){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(a.hostname)||"file:"===a.protocol)return d("localhost");if(!(window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)){try{if("true"===window.localStorage.vince_ignore)return d("localStorage flag")}catch(e){}var n={};n.n=e,n.u=a.href,n.d=r.getAttribute("data-domain"),n.r=o.referrer||null,n.w=window.innerWidth,t&&t.meta&&(n.m=JSON.stringify(t.meta)),t&&t.props&&(n.p=t.props),n.h=1;var i=new XMLHttpRequest;i.open("POST",c,!0),i.setRequestHeader("Content-Type","text/plain"),i.send(JSON.stringify(n)),i.onreadystatechange=function(){4===i.readyState&&t&&t.callback&&t.callback()}}}var w=window.vince&&window.vince.q||[];window.vince=i;for(var l,s=0;s<w.length;s++)i.apply(this,w[s]);function v(){l=a.pathname,i("pageview")}window.addEventListener("hashchange",v),"prerender"===o.visibilityState?o.addEventListener("visibilitychange",function(){l||"visible"!==o.visibilityState||v()}):v()}();