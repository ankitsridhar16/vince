!function(){"use strict";var o=window.location,c=window.document,s=c.currentScript,p=s.getAttribute("data-api")||new URL(s.src).origin+"/api/event";function t(t,e){try{if("true"===window.localStorage.vince_ignore)return void console.warn("Ignoring Event: localStorage flag")}catch(t){}var n={};n.n=t,n.u=o.href,n.d=s.getAttribute("data-domain"),n.r=c.referrer||null,n.w=window.innerWidth,e&&e.meta&&(n.m=JSON.stringify(e.meta)),e&&e.props&&(n.p=e.props);var i=s.getAttributeNames().filter(function(t){return"event-"===t.substring(0,6)}),a=n.p||{};i.forEach(function(t){var e=t.replace("event-",""),n=s.getAttribute(t);a[e]=a[e]||n}),n.p=a;var r=new XMLHttpRequest;r.open("POST",p,!0),r.setRequestHeader("Content-Type","text/plain"),r.send(JSON.stringify(n)),r.onreadystatechange=function(){4===r.readyState&&e&&e.callback&&e.callback()}}var e=window.vince&&window.vince.q||[];window.vince=t;for(var n,i=0;i<e.length;i++)t.apply(this,e[i]);function a(){n!==o.pathname&&(n=o.pathname,t("pageview"))}var r,u=window.history;u.pushState&&(r=u.pushState,u.pushState=function(){r.apply(this,arguments),a()},window.addEventListener("popstate",a)),"prerender"===c.visibilityState?c.addEventListener("visibilitychange",function(){n||"visible"!==c.visibilityState||a()}):a()}();