!function(){"use strict";var t,e,n,a=window.location,o=window.document,r=o.getElementById("vince"),s=r.getAttribute("data-api")||(t=r.src.split("/"),e=t[0],n=t[2],e+"//"+n+"/api/event");function w(t){console.warn("Ignoring Event: "+t)}function i(t,e){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(a.hostname)||"file:"===a.protocol)return w("localhost");if(!(window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)){try{if("true"===window.localStorage.vince_ignore)return w("localStorage flag")}catch(t){}var n={};n.n=t,n.u=a.href,n.d=r.getAttribute("data-domain"),n.r=o.referrer||null,n.w=window.innerWidth,e&&e.meta&&(n.m=JSON.stringify(e.meta)),e&&e.props&&(n.p=e.props);var i=new XMLHttpRequest;i.open("POST",s,!0),i.setRequestHeader("Content-Type","text/plain"),i.send(JSON.stringify(n)),i.onreadystatechange=function(){4===i.readyState&&e&&e.callback&&e.callback()}}}var c=window.vince&&window.vince.q||[];window.vince=i;for(var d,l=0;l<c.length;l++)i.apply(this,c[l]);function p(){d!==a.pathname&&(d=a.pathname,i("pageview"))}var v,u=window.history;u.pushState&&(v=u.pushState,u.pushState=function(){v.apply(this,arguments),p()},window.addEventListener("popstate",p)),"prerender"===o.visibilityState?o.addEventListener("visibilitychange",function(){d||"visible"!==o.visibilityState||p()}):p()}();