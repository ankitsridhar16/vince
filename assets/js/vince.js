!function(){"use strict";var a=window.location,r=window.document,o=r.currentScript,s=o.getAttribute("data-api")||new URL(o.src).origin+"/api/event";function w(t){console.warn("Ignoring Event: "+t)}function t(t,e){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(a.hostname)||"file:"===a.protocol)return w("localhost");if(!(window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)){try{if("true"===window.localStorage.vince_ignore)return w("localStorage flag")}catch(t){}var n={};n.n=t,n.u=a.href,n.d=o.getAttribute("data-domain"),n.r=r.referrer||null,n.w=window.innerWidth,e&&e.meta&&(n.m=JSON.stringify(e.meta)),e&&e.props&&(n.p=e.props);var i=new XMLHttpRequest;i.open("POST",s,!0),i.setRequestHeader("Content-Type","text/plain"),i.send(JSON.stringify(n)),i.onreadystatechange=function(){4===i.readyState&&e&&e.callback&&e.callback()}}}var e=window.vince&&window.vince.q||[];window.vince=t;for(var n,i=0;i<e.length;i++)t.apply(this,e[i]);function c(){n!==a.pathname&&(n=a.pathname,t("pageview"))}var d,p=window.history;p.pushState&&(d=p.pushState,p.pushState=function(){d.apply(this,arguments),c()},window.addEventListener("popstate",c)),"prerender"===r.visibilityState?r.addEventListener("visibilitychange",function(){n||"visible"!==r.visibilityState||c()}):c()}();