!function(){"use strict";var e,t,n,w=window.location,d=window.document,s=d.getElementById("vince"),u=s.getAttribute("data-api")||(e=s.src.split("/"),t=e[0],n=e[2],t+"//"+n+"/api/event");function p(e){console.warn("Ignoring Event: "+e)}function i(e,t){if(/^localhost$|^127(\.[0-9]+){0,2}\.[0-9]+$|^\[::1?\]$/.test(w.hostname)||"file:"===w.protocol)return p("localhost");if(!(window._phantom||window.__nightmare||window.navigator.webdriver||window.Cypress)){try{if("true"===window.localStorage.vince_ignore)return p("localStorage flag")}catch(e){}var n=s&&s.getAttribute("data-include"),i=s&&s.getAttribute("data-exclude");if("pageview"===e){var a=!n||n&&n.split(",").some(l),r=i&&i.split(",").some(l);if(!a||r)return p("exclusion rule")}var o={};o.n=e,o.u=t&&t.u?t.u:w.href,o.d=s.getAttribute("data-domain"),o.r=d.referrer||null,o.w=window.innerWidth,t&&t.meta&&(o.m=JSON.stringify(t.meta)),t&&t.props&&(o.p=t.props);var c=new XMLHttpRequest;c.open("POST",u,!0),c.setRequestHeader("Content-Type","text/plain"),c.send(JSON.stringify(o)),c.onreadystatechange=function(){4===c.readyState&&t&&t.callback&&t.callback()}}function l(e){return w.pathname.match(new RegExp("^"+e.trim().replace(/\*\*/g,".*").replace(/([^\.])\*/g,"$1[^\\s/]*")+"/?$"))}}var a=window.vince&&window.vince.q||[];window.vince=i;for(var r=0;r<a.length;r++)i.apply(this,a[r])}();